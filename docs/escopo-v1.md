# Escopo v1 — Modelo do DW (marts) para Winthor

Especificação das tabelas do DW na versão enxuta. Origem: as 42 consultas do ETL TARGIT atual e o modelo de dados (24 cubos), reorganizados em 23 tabelas, em 3 ondas.

## Decisões já tomadas

- O realizado das metas e o número principal de "Vendas" é o **faturamento líquido** (faturamento menos devoluções).
- Histórico carregado **desde 01/jan do ano anterior**, para permitir comparativos mês a mês e ano contra ano (janela configurável por tenant).
- Contas a pagar **mantém o filtro de grupos de conta por `PCLIB`** por enquanto.

## Convenções

- **`tenant_id` é a primeira coluna de todas as tabelas** e não é repetido nas tabelas abaixo.
- **Tipos por prefixo.** Os tipos finais saem do inventário de tabelas (E1.2) e prevalecem sobre esta regra.

| Prefixo / padrão | Tipo ClickHouse |
|---|---|
| `cod_*`, `num_*` | `Int64` (ou `String` quando a origem for alfanumérica) |
| `data`, `data_*` | `Date` |
| `valor_*` | `Decimal(18,2)` |
| `qtd_*`, `peso_*`, `volume_*` | `Decimal(18,3)` |
| `*_pct` | `Decimal(9,4)` |
| `eh_*` | `UInt8` |
| `nome_*`, `descricao_*`, `tipo_*`, `status` | `String` (`LowCardinality` quando houver poucos valores) |

- **Controle:** toda tabela tem `_atualizado_em DateTime`.
- **Fatos:** `PARTITION BY toYYYYMM(data)` e `ORDER BY (tenant_id, data, cod_filial, <chave do grão>)`. O grão de cada fato é a chave única dele.
- **Colunas que dependem da data atual (`SYSDATE`) não são gravadas.** Elas são calculadas na camada semântica (por exemplo status do cliente, dias de atraso, dias sem venda). Assim o valor nunca fica desatualizado.
- **Regras de negócio herdadas** do ETL atual (exclusões por condição de venda, CFOP, espécie etc.) ficam em um catálogo versionado de regras do dbt, e não espalhadas pelos modelos.
- **Dados pessoais:** CPF/CNPJ e e-mail de cliente ficam fora por padrão. Os e-mails de RCA, supervisor e gerente entram, porque servem para o envio agendado de relatórios.

## Visão geral

| Onda | Dimensões e pontes | Fatos |
|---|---|---|
| **1 — Comercial** | `dim_data`, `dim_filial`, `dim_cliente`, `dim_rca`, `dim_produto`, `dim_fornecedor`, `dim_motivo_devolucao`, `ponte_carteira` | `fato_faturamento`, `fato_devolucao`, `fato_venda`, `fato_meta` |
| **2 — Receber, estoque e nível de serviço** | `dim_regiao` | `fato_contas_receber`, `fato_estoque_snapshot`, `fato_corte`, `fato_falta` |
| **3 — Suprimentos e pagar** | `dim_parceiro`, `dim_centro_custo`, `ponte_lancamento_centro_custo` | `fato_compra`, `fato_pedido_compra`, `fato_contas_pagar` |

---

# Onda 1 — Comercial

## `dim_data`

Grão: um dia. Vai de 01/jan do ano anterior até 31/dez do ano seguinte. Gerada no dbt. O dia útil vem de `PCDATAS`.

| Coluna | Observação |
|---|---|
| `data` | chave |
| `ano`, `trimestre`, `mes`, `ano_mes` | `ano_mes` no formato AAAAMM |
| `nome_mes`, `nome_mes_abrev` | em pt-BR |
| `semana_iso`, `ano_semana_iso` | |
| `dia_mes`, `dia_semana`, `nome_dia_semana` | |
| `eh_dia_util` | `PCDATAS.DIAUTIL` diferente de 'N' |
| `seq_dia_util_mes` | n-ésimo dia útil do mês (nulo se não for útil). Base da projeção de fechamento |
| `qtd_dias_uteis_mes` | total de dias úteis do mês |

A mesma tabela serve para emissão, vencimento, pagamento e lançamento. A camada semântica cria os papéis.

## `dim_filial`

Origem: `PCFILIAL`. Chave: `cod_filial`.

| Coluna | Origem |
|---|---|
| `cod_filial` | `CODIGO` |
| `nome_filial` | `RAZAOSOCIAL` |
| `nome_fantasia` | `FANTASIA` |
| `descricao_filial` | código + razão social |
| `descricao_fantasia` | código + fantasia |
| `num_regiao_padrao` | `NUMREGIAOPADRAO`. Usada no preço do estoque (onda 2) |

## `dim_cliente`

Origem: `PCCLIENT` e cadastros. Chave: `cod_cliente`.

| Coluna | Origem |
|---|---|
| `cod_cliente`, `nome_cliente`, `nome_fantasia` | `CODCLI`, `CLIENTE`, `FANTASIA` |
| `descricao_cliente` | código + nome |
| `cod_cliente_principal`, `nome_cliente_principal` | `CODCLIPRINC` e o nome do cliente principal |
| `tipo_pessoa` | `TIPOFJ` |
| `cod_ramo_atividade`, `ramo_atividade` | `PCATIVI` (por `CODATV1`) |
| `cod_praca`, `praca` | `PCPRACA` |
| `cod_rota`, `rota` | `PCROTAEXP` (via praça) |
| `num_regiao`, `regiao`, `uf_regiao` | `PCREGIAO` (via praça) |
| `cod_cidade`, `cidade`, `uf_cidade`, `cod_ibge` | `PCCIDADE`. IBGE nulo vira 999, como no ETL atual |
| `bairro` | `BAIRROENT` |
| `data_cadastro`, `data_ultima_compra`, `data_exclusao` | `DTCADASTRO`, `DTULTCOMP`, `DTEXCLUSAO` |
| `dias_param_inatividade` | `PCCONSUM.NUMDIASCLIINATIV` |

Calculados na camada semântica (dependem da data atual):

- `status_cliente`: EXCLUÍDO se `data_exclusao` tem valor. INATIVO se os dias desde `data_ultima_compra` passam o parâmetro. Senão ATIVO.
- `dias_cadastro`
- `dias_ultima_compra`

## `dim_rca`

Origem: `PCUSUARI`, `PCSUPERV`, `PCGERENTE`. Chave: `cod_rca`.

| Coluna | Origem |
|---|---|
| `cod_rca`, `nome_rca`, `descricao_rca`, `email_rca` | `CODUSUR`, `NOME`, `EMAIL` |
| `tipo_rca` | `TIPOVEND`: R representante, P profissional, I interno, E externo, outros "NÃO DEFINIDO" |
| `situacao_rca` | ATIVO se `DTTERMINO` for nulo, senão INATIVO |
| `data_termino` | `DTTERMINO` |
| `cod_supervisor`, `nome_supervisor`, `descricao_supervisor`, `email_supervisor` | `PCSUPERV` |
| `cod_gerente`, `nome_gerente`, `descricao_gerente`, `email_gerente` | `PCGERENTE` |

## `ponte_carteira`

Grão: par (RCA, cliente). Origem: `PCCLIENT` (`CODUSUR1/2/3`) e `PCUSURCLI`.

| Coluna | Observação |
|---|---|
| `cod_rca`, `cod_cliente` | chave |
| `tipo_vinculo` | vínculo principal, com prioridade RCA 1, RCA 2, RCA 3, MÚLTIPLO |
| `eh_rca1`, `eh_rca2`, `eh_rca3`, `eh_multiplo` | marcadores. Evitam duplicar linhas quando o mesmo par aparece em mais de um vínculo |

No ETL atual o par pode se repetir (um por tipo de vínculo), o que duplica valores ao juntar com os fatos. Aqui o grão é o par. Esta tabela serve às métricas de carteira, e não deve ser juntada aos fatos por padrão.

## `dim_produto`

Origem: `PCPRODUT` mais departamento, seção, categoria, subcategoria, marca e linha, tudo achatado. Chave: `cod_produto`.

| Coluna | Origem |
|---|---|
| `cod_produto`, `nome_produto`, `descricao_produto` | `CODPROD`, `DESCRICAO` |
| `cod_departamento`, `departamento`, `descricao_departamento` | `PCDEPTO` |
| `tipo_mercadoria` | `PCDEPTO.TIPOMERC`. Usado nos filtros de compras e estoque (exclui CI e IM) |
| `cod_secao`, `secao`, `descricao_secao` | `PCSECAO` |
| `cod_categoria`, `categoria`, `descricao_categoria` | `PCCATEGORIA` (chave composta seção + categoria) |
| `cod_subcategoria`, `subcategoria`, `descricao_subcategoria` | `PCSUBCATEGORIA` (chave composta seção + categoria + subcategoria) |
| `cod_marca`, `marca` | `PCMARCA` |
| `cod_linha`, `linha` | `PCLINHAPROD` |
| `cod_fornecedor` | `PCPRODUT.CODFORNEC` |
| `qtd_unidades_caixa` | `QTUNITCX` |
| `peso_bruto_unit`, `peso_liquido_unit`, `volume_unit` | `PESOBRUTO`, `PESOLIQ`, `VOLUME` |
| `data_vencimento` | `DTVENC`. Usada na análise de validade do estoque |
| `data_exclusao` | `DTEXCLUSAO` |

## `dim_fornecedor`

Origem: `PCFORNEC`. Chave: `cod_fornecedor`.

| Coluna | Origem |
|---|---|
| `cod_fornecedor`, `nome_fornecedor`, `descricao_fornecedor` | `CODFORNEC`, `FORNECEDOR` |
| `cod_fornecedor_principal`, `nome_fornecedor_principal`, `descricao_fornecedor_principal` | `CODFORNECPRINC` |
| `cod_comprador` | `CODCOMPRADOR` |
| `prazo_entrega_dias` | `PRAZOENTREGA` |

## `dim_motivo_devolucao`

Origem: `PCTABDEV`. Colunas: `cod_motivo_devolucao`, `motivo_devolucao`.

## `fato_faturamento`

Grão: item de nota fiscal de saída (`PCMOV` com `PCNFSAID`). Chave: `num_transacao_item` (`NUMTRANSITEM`, a confirmar).

| Coluna | Origem / regra |
|---|---|
| `num_transacao_item` | `PCMOV.NUMTRANSITEM` |
| `data` | `PCNFSAID.DTSAIDA` |
| `cod_filial`, `cod_cliente`, `cod_rca`, `cod_produto` | `PCMOV` |
| `cod_fornecedor` | `PCPRODUT.CODFORNEC` |
| `num_nota_saida`, `num_transacao_venda`, `num_pedido_venda` | |
| `tipo_venda`, `condicao_venda`, `especie`, `serie` | `PCNFSAID` |
| `origem_pedido`, `tipo_pedido` | `PCPEDC` (ORIGEMPED, e a regra de brinde, bonificação, transferência etc. do `DIM_PEDIDOVENDA`) |
| `valor` | valor faturado (subtotal do item; senão qtd × preço com frete, outras despesas, IPI e ST) |
| `valor_tabela` | qtd × preço de tabela |
| `valor_bonificado` | itens de condição de venda 5, 6, 11 ou 12, ao preço de tabela |
| `valor_st` | ST × qtd |
| `valor_desconto_financeiro` | `PERCDESCFIN` × valor / 100 |
| `cmv` | qtd × `CUSTOFIN` |
| `qtd_itens`, `qtd_itens_bonificados`, `qtd_caixas` | |
| `peso_bruto`, `peso_liquido`, `volume` | |

Exclusões herdadas do ETL atual: nota cancelada, `TIPOVENDA` SR ou DF, `TIPOITEM` fora de C e N, CFOP 0, 722, 532, 632 e 732, e condições de venda 4, 0, 8, 13, 20, 98 e 99.

## `fato_devolucao`

Grão: item de nota de entrada de devolução (`PCMOV` operação ED). Inclui a devolução vinculada a uma venda e a avulsa, separadas por `tipo_devolucao`. Chave: `num_transacao_item`.

| Coluna | Origem / regra |
|---|---|
| `num_transacao_item` | `PCMOV.NUMTRANSITEM` |
| `tipo_devolucao` | VINCULADA (tem transação de venda de origem) ou AVULSA |
| `data` | `PCNFENT.DTENT` (data da devolução) |
| `cod_filial`, `cod_cliente`, `cod_produto` | `PCMOV` |
| `cod_rca` | `PCNFENT.CODUSURDEVOL` |
| `cod_fornecedor` | `PCPRODUT.CODFORNEC` |
| `cod_motivo_devolucao` | `PCNFENT.CODDEVOL` |
| `num_nota_devolucao`, `num_transacao_entrada` | `PCNFENT` |
| `num_transacao_venda`, `num_pedido_venda` | da nota de venda de origem. Nulos na avulsa |
| `valor`, `valor_st`, `cmv` | mesma lógica de valor do faturamento (bonificação zerada) |
| `qtd_itens`, `qtd_caixas`, `peso_bruto`, `peso_liquido`, `volume` | |

## `fato_venda` (pedidos)

Grão: item de pedido de venda (`PCPEDI` com `PCPEDC`). Chave: (`num_pedido_venda`, `cod_produto`, `num_seq_item`), a confirmar.

| Coluna | Origem / regra |
|---|---|
| `data` | `PCPEDC.DATA` |
| `num_pedido_venda`, `cod_produto`, `num_seq_item` | |
| `cod_filial`, `cod_cliente`, `cod_rca` | |
| `cod_fornecedor` | `PCPRODUT.CODFORNEC` |
| `origem_pedido`, `situacao_pedido`, `tipo_pedido`, `condicao_venda` | `PCPEDC`. A situação vem de `POSICAO` (L, P, B, M, F) |
| `valor` | subtotal do item; zero quando bonificado |
| `valor_tabela`, `valor_bonificado` | |
| `cmv` | qtd × `VLCUSTOFIN` |
| `qtd_itens`, `qtd_itens_bonificados` | |
| `peso_bruto`, `peso_liquido`, `volume` | |

Exclusões herdadas: posição fora de L, M, B, P, F e C, condições de venda 4, 8, 13, 20, 98 e 99, e pedidos cancelados.

## `fato_meta`

Grão: uma meta por data, tipo, chave, filial e RCA. Origem: `PCMETA` e `PCMETARCA`.

| Coluna | Origem / regra |
|---|---|
| `data` | `PCMETA.DATA` |
| `tipo_meta` | A categoria, I cliente, D departamento, F fornecedor, W marca, P produto, M RCA, S seção. `PCMETARCA` entra como `RCA_353` |
| `cod_chave`, `cod_chave2` | `CODIGO`, `CODIGO2`. O significado depende do `tipo_meta` |
| `cod_filial`, `cod_rca` | |
| `valor_meta`, `qtd_venda_meta`, `qtd_pedidos_meta`, `positivacao_meta`, `mix_meta` | `VLVENDAPREV`, `QTVENDAPREV`, `PEDIDOSPREV`, `CLIPOSPREV`, `MIXPREV` |
| `peso_meta`, `volume_meta` | `QTPESOPREV`, `VOLUMEPREV` |
| `margem_meta_pct`, `inadimplencia_meta_pct`, `devolucao_meta_pct` | percentuais. **Não somar**: a camada semântica usa média ponderada |

Calculados na camada semântica:

- **Atingimento:** faturamento líquido ÷ `valor_meta`.
- **Projeção de fechamento:** realizado ÷ dias úteis decorridos × dias úteis do mês.

---

# Onda 2 — Receber, estoque e nível de serviço

## `dim_regiao`

Origem: `PCREGIAO`. Colunas: `num_regiao`, `regiao`, `uf`, `status`.

## `fato_contas_receber`

Grão: título (transação de venda + prestação). Origem: `PCPREST`. Chave: (`num_transacao_venda`, `num_prestacao`).

| Coluna | Origem / regra |
|---|---|
| `codigo_titulo` | `T` + transação + `P` + prestação |
| `num_transacao_venda`, `num_prestacao`, `num_duplicata` | |
| `cod_cliente`, `cod_rca`, `cod_filial` | |
| `cod_cobranca` | `CODCOB` |
| `data_emissao`, `data_vencimento`, `data_pagamento`, `data_baixa` | `DTEMISSAO`, `DTVENC`, `DTPAG`, `DTBAIXA` |
| `status` | ABERTO se `DTPAG` e `DTBAIXA` forem nulos, senão PAGO |
| `valor_original` | `VALOR` |
| `valor_recebido`, `valor_juros`, `valor_desconto` | `VPAGO`, `TXPERM`, `VALORDESC` |
| `valor_a_receber` | se não pago e a cobrança não for devolução, bonificação, crédito ou transferência: valor menos desconto. Senão 0 |

Calculados na camada semântica (dependem da data atual): `em_atraso`, `inadimplente`, `valor_inadimplente`, `dias_atraso` e as faixas de aging.

**Carga:** além da janela por vencimento, carregar **todos os títulos em aberto**, independente da data. Senão a inadimplência de períodos anteriores some.

## `fato_estoque_snapshot`

Grão: foto diária por filial e produto. Origem: `PCEST`, com preço da `PCTABPR` na **região padrão da filial** e `PCPARAMFILIAL` (preço por embalagem). A posição atual é a foto mais recente (visão `estoque_atual`).

| Coluna | Origem / regra |
|---|---|
| `data_snapshot` | data da foto |
| `cod_filial`, `cod_produto`, `cod_fornecedor` | |
| `qtd_gerencial` | `QTESTGER` |
| `qtd_reservada`, `qtd_bloqueada`, `qtd_pendente`, `qtd_avaria` | `QTRESERV`, `QTBLOQUEADA`, `QTPENDENTE`, `QTINDENIZ` |
| `qtd_disponivel` | gerencial − reservada − bloqueada − pendente, quando o parâmetro `BLOQUEIAVENDAESTPENDENTE` for 'S' |
| `custo_real_unit`, `custo_fin_unit`, `custo_cont_unit`, `custo_rep_unit`, `custo_ult_entrada_unit` | `CUSTOREAL`, `CUSTOFIN`, `CUSTOCONT`, `CUSTOREP`, `CUSTOULTENT` |
| `valor_ultima_entrada_unit` | `VALORULTENT` |
| `preco_venda_unit` | preço de tabela da região padrão (ou preço por embalagem, conforme o parâmetro da filial) |
| `valor_custo_real`, `valor_custo_fin`, `valor_custo_cont`, `valor_custo_rep`, `valor_custo_ult_entrada`, `valor_ultima_entrada`, `valor_venda` | quantidade disponível × valor unitário |
| `peso_bruto_total`, `peso_liquido_total`, `volume_total` | quantidade disponível × atributo do produto |
| `giro_dia` | `QTGIRODIA` |
| `tempo_reposicao_dias` | `PCPRODUT.TEMREPOS` |
| `qtd_estoque_ideal` | giro × (reposição + prazo de entrega do fornecedor) |
| `data_ultima_entrada`, `data_ultima_saida` | `DTULTENT`, `DTULTSAIDA` |

Calculados na camada semântica: `dias_sem_compra`, `dias_sem_venda`, dias até o vencimento e faixa de validade, relativos a `data_snapshot` (assim o histórico fica correto). Também a cobertura em dias (`qtd_disponivel ÷ giro_dia`) e o estoque parado.

**Retenção:** foto diária. Reavaliar para mensal depois de 13 meses.

## `fato_corte`

Grão: item cortado. Origem: `PCCORTEI`. Chave: (`num_pedido_venda`, `cod_produto`, `data`, `motivo_corte`), a confirmar.

| Coluna | Origem |
|---|---|
| `data` | `DATA` (data do corte) |
| `num_pedido_venda`, `cod_produto`, `motivo_corte` | |
| `cod_filial`, `cod_rca`, `cod_funcionario` | |
| `cod_fornecedor` | via `PCPRODUT` |
| `qtd_cortada`, `qtd_separada` | `QTCORTADA`, `QTSEPARADA` |
| `valor_cortado` | `QTCORTADA × PVENDA` |

## `fato_falta`

Grão: item em falta. Origem: `PCFALTA` com `PCPEDC` (exclui a condição de venda 10). Chave: (`num_pedido_venda`, `cod_produto`, `data`), a confirmar.

| Coluna | Origem |
|---|---|
| `data` | `PCFALTA.DATA` (data da falta) |
| `data_venda` | `PCPEDC.DATA` |
| `num_pedido_venda`, `cod_produto` | |
| `cod_filial`, `cod_rca`, `cod_cliente` | `PCPEDC` |
| `qtd_em_falta` | `QT` |
| `valor_falta` | `QT × PVENDA` |

---

# Onda 3 — Suprimentos e pagar

## `dim_parceiro`

Origem: `PCFORNEC`, `PCCLIENT`, `PCUSUARI`, `PCEMPR`. Chave: `cod_parceiro`.

| Coluna | Observação |
|---|---|
| `cod_parceiro` | letra do tipo + 9 dígitos: F fornecedor, C cliente, R RCA, M motorista, L colaborador; `0` para outros |
| `tipo_parceiro` | FORNECEDOR, CLIENTE, RCA, MOTORISTA, COLABORADOR, OUTROS |
| `cod_origem` | código original no cadastro de origem |
| `nome_parceiro`, `descricao_parceiro` | |

## `dim_centro_custo`

Origem: `PCCENTROCUSTO`. Colunas: `cod_centro_custo`, `centro_custo`, `descricao_centro_custo`, `ativo`.

## `ponte_lancamento_centro_custo`

Grão: lançamento × centro de custo. Origem: `PCRATEIOCENTROCUSTO`.

| Coluna | Observação |
|---|---|
| `num_lancamento` | `RECNUM` |
| `cod_centro_custo` | |
| `perc_rateio` ou `valor_rateio` | **a confirmar** nas colunas da tabela de rateio |

Um lançamento pode ter vários centros de custo. Ligar direto ao fato duplicaria o valor. O rateio passa a ser aplicado pela ponte.

## `fato_compra`

Grão: item de nota de entrada de compra (`PCMOV` operações E e EB, com `PCNFENT`). Chave: `num_transacao_item`.

| Coluna | Origem / regra |
|---|---|
| `num_transacao_item` | `PCMOV.NUMTRANSITEM` |
| `data` | `PCNFENT.DTENT` |
| `cod_filial`, `cod_fornecedor`, `cod_produto`, `cod_comprador` | |
| `num_nota_entrada`, `num_transacao_entrada`, `especie`, `serie`, `tipo_entrada` | |
| `num_pedido_compra` | pedido de origem, se `PCNFENT` ou `PCMOV` trouxer (a confirmar). Permite medir o prazo real de entrega do fornecedor |
| `valor` | qtd × `PUNIT` |
| `valor_custo` | qtd × `CUSTOFIN` |
| `valor_desconto` | qtd × `VLDESCONTO` |
| `qtd_itens`, `peso_bruto`, `peso_liquido`, `volume` | |

Exclusões herdadas: item cancelado, `TIPODESCARGA` fora de 1, 5 e I, contabilização diferente do parâmetro de fornecedor e departamentos CI e IM.

## `fato_pedido_compra`

Grão: item de pedido de compra (`PCITEM` com `PCPEDIDO`). Chave: (`num_pedido_compra`, `cod_produto`, `num_seq`).

| Coluna | Origem / regra |
|---|---|
| `data_emissao`, `data_entrega_prevista` | `DTEMISSAO`, `DATAENTREGA` |
| `num_pedido_compra`, `cod_produto`, `num_seq` | |
| `cod_filial`, `cod_fornecedor` | |
| `situacao` | FECHADO se `qtd_pendente` for 0, senão PENDENTE |
| `qtd_pedida`, `qtd_entregue`, `qtd_pendente` | em unidades do pedido (fator de embalagem aplicado) |
| `valor_pedido` | `qtd_pedida` × preço da mercadoria |
| `valor_entregue` | `qtd_entregue` × preço da mercadoria |
| `valor_pendente` | `qtd_pendente` × preço da mercadoria |
| `valor_desconto` | desconto do item |
| `prazo_entrega_dias` | `PCFORNEC.PRAZOENTREGA` |

As definições de valor aqui são uma simplificação. O ETL atual tem fórmulas diferentes para pendente e total.

## `fato_contas_pagar`

Grão: lançamento (`RECNUM`). Origem: `PCLANC` com `PCCONTA`, filtrado pelos grupos de conta liberados em `PCLIB` (`CODFUNC` = 1, `CODTABELA` = 6).

| Coluna | Origem / regra |
|---|---|
| `num_lancamento`, `num_lancamento_principal` | `RECNUM`, `RECNUMPRINC` |
| `num_prestacao`, `num_duplicata` | `PREST`, `DUPLIC` |
| `tipo_pagamento` | `TIPOPAGTO` |
| `tipo_lancamento` | Confirmado (C) ou Provisionado (P) |
| `status` | ABERTO se `DTPAGTO` for nulo, senão PAGO |
| `cod_parceiro` | letra do `TIPOPARCEIRO` + `CODFORNEC` com 9 dígitos |
| `cod_fornecedor`, `cod_filial`, `cod_conta`, `cod_grupo_conta` | |
| `data_emissao`, `data_vencimento`, `data_lancamento`, `data_pagamento` | |
| `valor_original`, `valor_pago`, `valor_desconto`, `valor_juros`, `valor_devolucao`, `variacao_cambial` | `VALOR`, `VPAGO`, `DESCONTOFIN`, `TXPERM`, `VALORDEV`, `VLVARIACAOCAMBIAL` |
| `valor_a_pagar` | se não pago: valor + juros − desconto − devolução + variação cambial. Senão 0 |

**Carga:** além da janela por data de lançamento, carregar **todos os lançamentos em aberto**.

---

# Calculado na camada semântica (por onda)

**Onda 1**
- Faturamento líquido, lucro, % lucro, ticket médio, % bonificado, % desconto, % devolução.
- Positivação (clientes distintos), mix (produtos distintos), nº de notas e de pedidos.
- Atingimento de meta, projeção de fechamento por dias úteis.
- Comparativos: mês anterior, mesmo mês do ano anterior e acumulado do ano.
- Cobertura da carteira (positivados ÷ clientes ativos).
- Curva ABC de produtos e clientes.

**Onda 2**
- Inadimplência em valor e %, aging, perdas de crédito (> 365 dias).
- Cobertura de estoque em dias, estoque parado.
- % de corte (valor cortado ÷ valor vendido).

**Onda 3**
- Valor líquido de contas a pagar.
- % de desconto em compras.
- Pedido de compra: valor pendente e atraso na entrega.

---

# Principais análises por onda

Cada análise responde a uma pergunta de negócio e termina em uma ação de alguém. Este catálogo tem dois usos: define os **dashboards padrão** de cada onda (as marcadas com ★) e serve de **conjunto de perguntas de referência** para validar a camada semântica e o Claude (fases 2 e 3). Os nomes de medidas e colunas são os deste documento.

## Onda 1 — Comercial

| # | Análise | Pergunta de negócio | Medidas e dimensões | Insight e ação |
|---|---|---|---|---|
| 1 ★ | **Meta e ritmo do mês** | Quem vai bater a meta e quem precisa de ajuda agora? | `valor_meta`, faturamento líquido, `seq_dia_util_mes`, `qtd_dias_uteis_mes`. Por RCA, supervisor e filial | Mostra atingimento, projeção de fechamento e quanto falta vender por dia útil restante. O supervisor aciona os RCAs abaixo do ritmo enquanto ainda dá tempo de reverter. Bom candidato a relatório diário |
| 2 ★ | **Faturamento líquido e comparativos** | Estamos crescendo? Onde? | Faturamento líquido, mês anterior, mesmo mês do ano anterior, acumulado do ano. Por filial, fornecedor, categoria e região | Mostra onde a queda ou o crescimento está concentrado (fornecedor, categoria, praça), em vez de só "o total caiu". Base da reunião mensal de resultado |
| 3 ★ | **Positivação e cobertura da carteira** | Quantos clientes ativos compraram no mês? Quais não? | Positivação (clientes distintos), clientes ativos, `ponte_carteira`. Por RCA, rota e praça | Gera a lista de clientes ativos que ainda não compraram no mês, por RCA e rota, ordenada pelo valor histórico. O RCA sabe quem visitar primeiro |
| 4 | **Clientes em risco** | Quais clientes estão deixando de comprar? | Dias desde a última compra contra o intervalo habitual de compra do cliente, queda contra o mesmo período do ano anterior | Pega o cliente que passou do ritmo normal dele antes de virar inativo. Ação: contato do RCA ou do supervisor, com o histórico do cliente. Depende do histórico desde o ano anterior |
| 5 | **Rentabilidade** | Onde ganhamos e onde perdemos margem? | Lucro, % lucro, `cmv`. Por produto, cliente, RCA, fornecedor e filial | Revela itens, clientes e RCAs com margem baixa ou negativa. Ação: rever preço, renegociar condição ou limitar a venda |
| 6 | **Descontos e bonificação** | Quanto damos de desconto e de bonificação, e para quem? | `valor_tabela` − `valor`, `valor_bonificado`, % bonificado, `valor_desconto_financeiro`. Por RCA, cliente e fornecedor | Mostra RCAs que descontam muito acima da média e fornecedores com bonificação alta. Ação: definir alçada de desconto e cobrar contrapartida do fornecedor |
| 7 ★ | **Devoluções** | Por que devolvem e quanto isso custa? | `fato_devolucao`, % devolução, `cod_motivo_devolucao`. Por RCA, cliente, produto, fornecedor e filial | Separa motivos logísticos, comerciais e de qualidade. Ação: corrigir o processo na origem e cobrar do fornecedor o que for qualidade. A devolução vinculada permite voltar ao RCA e à nota de venda |
| 8 | **Mix e venda cruzada** | O que clientes parecidos compram que este cliente não compra? | Mix (produtos distintos), `ramo_atividade`, categoria, cliente | Compara o cliente com o perfil de compra do ramo de atividade dele e lista os produtos que ele ainda não compra. Ação: sugestão de itens para o RCA oferecer |
| 9 | **Curva ABC** | Quão concentrada é a venda? | Faturamento líquido acumulado. Por produto e cliente | Define o foco de abastecimento e de atendimento (itens e clientes A) e expõe a dependência de poucos clientes. Ação: proteger o estoque dos itens A, atendimento diferenciado e plano para a cauda |
| 10 | **Pedidos parados** | Quanto dinheiro está preso em pedidos que não faturam? | `fato_venda` por `situacao_pedido` (bloqueado, pendente etc.), data do pedido contra o faturamento | Mostra pedidos bloqueados ou pendentes há dias, em valor. Ação: liberar crédito ou estoque, ou cancelar. É venda feita que ainda não virou faturamento |
| 11 | **Canais de venda** | Por qual canal os pedidos entram? | `origem_pedido`, ticket médio, nº de pedidos. Por cliente e RCA | Revela clientes pequenos atendidos por canal caro e a adoção do canal digital. Ação: migrar clientes para um canal mais eficiente |

## Onda 2 — Receber, estoque e nível de serviço

| # | Análise | Pergunta de negócio | Medidas e dimensões | Insight e ação |
|---|---|---|---|---|
| 1 ★ | **Inadimplência e aging** | Quanto está atrasado e há quanto tempo? | `valor_inadimplente`, % inadimplência, faixas de atraso. Por cliente, RCA, filial e cobrança | Prioriza a cobrança pelos títulos de maior valor e mais antigos e mostra a carteira inadimplente de cada RCA. Ação: cobrar e suspender novas vendas a quem passou do limite. Relatório semanal ao financeiro e ao RCA |
| 2 ★ | **Recebimentos previstos** | Quanto entra nas próximas semanas? | `valor_a_receber` por `data_vencimento` futura. Por filial e cobrança | Agenda de entradas. Ação: planejar o caixa e antecipar a cobrança nas semanas fracas. Base do fluxo de caixa projetado |
| 3 | **Prazo médio e pontualidade** | Os clientes pagam no prazo combinado? | `data_pagamento` − `data_emissao` e `data_pagamento` − `data_vencimento`. Por cliente e cobrança | Separa bons e maus pagadores e compara o prazo real com o combinado. Ação: renegociar prazo, revisar limite de crédito e incentivar formas de cobrança melhores |
| 4 | **Perdas de crédito** | Quanto provavelmente não será recebido? | Títulos vencidos há mais de 365 dias (`valor_a_receber`). Por RCA e cliente | Dimensiona a provisão e mostra os maiores devedores antigos. Ação: protesto, negociação ou baixa |
| 5 ★ | **Ruptura iminente** | O que vai faltar antes da reposição chegar? | Cobertura (`qtd_disponivel` ÷ `giro_dia`) contra `tempo_reposicao_dias` + prazo do fornecedor. Por produto, fornecedor e filial | Lista os produtos cuja cobertura é menor que o prazo de reposição, com a quantidade a comprar por fornecedor. Ação: pedido de compra hoje, antes de perder venda |
| 6 ★ | **Excesso e estoque parado** | Onde está o dinheiro parado em estoque? | `valor_custo_fin`, dias sem venda, cobertura acima do ideal (`qtd_estoque_ideal`). Por fornecedor, categoria e filial | Mostra o valor imobilizado em itens sem venda ou com cobertura muito alta. Ação: promoção ou queima, devolução ao fornecedor, suspender a compra |
| 7 | **Validade** | O que vence antes de vender? | `data_vencimento` do produto, `qtd_disponivel`, giro. Por produto e filial | Itens que vencem antes de girar. Ação: priorizar a venda, promover ou negociar troca. **Depende de confirmar o que `DTVENC` representa** no Winthor do cliente |
| 8 | **Estoque indisponível** | Quanto do estoque não pode ser vendido? | `qtd_bloqueada`, `qtd_avaria`, `qtd_reservada` contra `qtd_gerencial`. Por produto e filial | Resolver os bloqueios (qualidade, fiscal, avaria) devolve estoque vendável. Ação: tratar primeiro os maiores valores |
| 9 | **Dias de estoque (tendência)** | O estoque está crescendo mais que a venda? | Valor do estoque na foto diária contra o CMV do período. Por filial e fornecedor | Mostra a evolução do giro e dos dias de estoque. Ação: ajustar o volume de compra. O histórico começa na implantação da foto diária |
| 10 ★ | **Nível de serviço: corte e falta** | Quanto deixamos de entregar e por quê? | % de corte, `motivo_corte`, `valor_falta`. Por produto, fornecedor e filial | Estima a "venda perdida" (valor cortado + valor em falta) e separa a causa: falta de compra, separação ou cadastro. Ação: corrigir a compra dos itens com falta recorrente e atacar o motivo de corte mais frequente |
| 11 | **Clientes mais afetados** | Quais clientes mais sofrem com corte e falta? | Corte e falta por cliente e RCA, cruzados com o faturamento | Cliente importante com muito corte é cliente em risco. Ação: priorizar a reposição e avisar o RCA para tratar |

## Onda 3 — Suprimentos e pagar

| # | Análise | Pergunta de negócio | Medidas e dimensões | Insight e ação |
|---|---|---|---|---|
| 1 ★ | **Custo de compra e variação** | O custo subiu? Já repassamos ao preço? | Custo unitário de compra (`valor_custo` ÷ `qtd_itens`) ao longo do tempo, contra o preço de venda. Por produto e fornecedor | Alerta os produtos cujo custo subiu e cuja margem caiu. Ação: reajustar o preço de venda ou negociar com o fornecedor |
| 2 | **Condições por fornecedor** | Quem oferece as melhores condições? | % de desconto em compras (`valor_desconto` ÷ `valor`), volume comprado. Por fornecedor | Compara fornecedores e embasa a negociação. Ação: concentrar a compra onde as condições são melhores |
| 3 ★ | **Pedidos de compra pendentes e atrasados** | O que foi comprado e ainda não chegou? | `qtd_pendente`, `valor_pendente`, `data_entrega_prevista` vencida. Por fornecedor e produto | Lista os pedidos atrasados por fornecedor, com o impacto nos itens em ruptura iminente (onda 2). Ação: cobrar o fornecedor ou comprar de outro |
| 4 | **Nível de atendimento do fornecedor** | O fornecedor entrega o que pedimos, no prazo? | `qtd_entregue` ÷ `qtd_pedida`, prazo prometido (`prazo_entrega_dias`) contra o prazo realizado. Por fornecedor | Ranking de confiabilidade. Ação: renegociar ou diversificar fornecedores. **O prazo realizado exige ligar a nota de entrada ao pedido de compra** (ver pontos a confirmar) |
| 5 ★ | **Agenda de pagamentos** | O que vence e quanto de caixa preciso? | `valor_a_pagar` por `data_vencimento`, títulos vencidos. Por fornecedor e filial | Agenda de saídas e títulos vencidos. Ação: priorizar pagamentos e evitar juros e protesto |
| 6 | **Juros e descontos nos pagamentos** | Pagamos juros por atraso ou aproveitamos descontos? | `valor_juros`, `valor_desconto`. Por fornecedor | Mostra o custo do atraso e os descontos obtidos. Ação: pagar em dia onde o juros supera o ganho de caixa e negociar antecipação onde há desconto |
| 7 | **Despesas por centro de custo** | Para onde vai o gasto? | `valor_original` rateado por `ponte_lancamento_centro_custo`. Por centro de custo e mês | Evolução e comparativo mês a mês. Ação: controle orçamentário e investigação de picos |

## Análises cruzadas (disponíveis quando as ondas indicadas estiverem no ar)

| # | Análise | Ondas | O que combina | Insight e ação |
|---|---|---|---|---|
| A ★ | **Qualidade da venda por RCA** | 1 + 2 | Faturamento líquido, devolução, desconto, inadimplência e corte, por RCA | O RCA que mais vende nem sempre é o que mais gera resultado. Ação: ajustar meta e comissão pelo resultado, e não só pelo volume |
| B | **Rentabilidade por fornecedor** | 1 + 2 (+ 3) | Margem, % bonificado, % devolução, giro, estoque imobilizado e desconto de compra | Pauta de negociação com cada indústria, com os números do próprio cliente |
| C ★ | **Fluxo de caixa operacional projetado** | 2 + 3 | Recebimentos previstos menos pagamentos previstos, por semana | Antecipa semanas de aperto. Ação: planejar o capital de giro. Limite: não inclui saldo bancário nem despesas fora do Winthor |
| D | **Capital imobilizado contra retorno** | 1 + 2 | Valor do estoque, margem e giro por fornecedor e categoria | Mostra onde cada real investido em estoque rende mais. Ação: realocar a compra |
| E | **Cliente em risco ampliado** | 1 + 2 | Queda de compra, corte sofrido e atraso no pagamento | Prioriza a retenção dos clientes que combinam sinais de perda. Ação: plano de contato do RCA e do supervisor |

## Pacote de relatórios agendados por perfil

| Perfil | Conteúdo | Frequência | Canal |
|---|---|---|---|
| Diretor ou dono | Faturamento líquido contra meta e contra o ano anterior, margem, inadimplência, estoque parado e fluxo de caixa de 30 dias | Semanal e fechamento mensal | E-mail (PDF) e resumo no WhatsApp |
| Gerente comercial | Meta e ritmo por supervisor, qualidade da venda por RCA, devoluções | Diária e semanal | WhatsApp e e-mail |
| Supervisor | Meta e ritmo dos RCAs, clientes sem compra, cortes | Diária | WhatsApp |
| RCA | Minha meta e ritmo, meus clientes sem compra, meus títulos vencidos | Diária (manhã) | WhatsApp |
| Financeiro | Aging, agenda de recebimentos e de pagamentos | Diária ou semanal | E-mail |
| Compras | Ruptura iminente, pedidos atrasados, estoque parado por fornecedor e variação de custo | Diária ou semanal | E-mail |

Os e-mails de RCA, supervisor e gerente vêm de `dim_rca`. **Telefone ou WhatsApp não vêm do Winthor neste escopo**: o cadastro de destinatários é feito na plataforma, com o consentimento (opt-in) de cada pessoa.

## Fora do escopo da v1 (candidatos futuros)

Logística de entrega e frete, comissões, histórico de políticas de preço e visitas da força de vendas.

---

# Pontos a confirmar

1. `PCDATAS` tem uma linha por data ou por filial? A consulta atual não filtra por filial.
2. Chaves reais: `NUMTRANSITEM` em `PCMOV`, `NUMSEQ` em `PCPEDI` e em `PCITEM`, e a chave de `PCSECAO` (só `CODSEC`, ou `CODEPTO` + `CODSEC`).
3. `PCMETA`: o tipo M é de RCA? Ele e `PCMETARCA` parecem metas de RCA de rotinas diferentes, então é preciso escolher uma fonte oficial. A coluna `DATA` é mensal?
4. Estoque: o ETL atual junta a tabela de preço sem filtrar a região, o que pode multiplicar o estoque. Aqui usei a região padrão da filial, e isso precisa ser validado contra os números atuais. Além disso, a quantidade disponível considera o parâmetro de pendente, mas os totais de custo sempre subtraem o pendente. É preciso definir uma regra só.
5. Cobrança '033' tem tolerância de 1 a 2 dias na inadimplência. É uma regra específica do cliente. O ideal é virar parâmetro por tenant.
6. Os sinais de "dias de inadimplência" e "dias até o vencimento" no ETL atual parecem invertidos. Aqui serão positivos (dias de atraso).
7. O "valor de desconto" tem sinais diferentes no faturamento e na venda. Aqui: tabela menos faturado.
8. Regras de exclusão (condição de venda, CFOP, espécie) variam entre os fatos (por exemplo, a condição 8 sai do faturamento mas não da devolução). Confirmar cada uma com o negócio.
9. Devolução e pedido de compra têm joins sem uso que funcionam como filtros ocultos. A paridade contra os números atuais do TARGIT mostrará o efeito.
10. `PCLIB` (grupos de conta no contas a pagar) é uma regra de configuração de cada cliente. Confirmar se vale para todos.
11. Colunas de `PCRATEIOCENTROCUSTO` (valor ou percentual de rateio).
12. Tipos reais das chaves (`cod_filial` costuma ser texto no Winthor) e a versão do Oracle Database de cada cliente.
13. A nota de entrada guarda o pedido de compra de origem? Sem isso, o prazo realizado de entrega do fornecedor (análise 4 da onda 3) não é calculável.
14. O que `PCPRODUT.DTVENC` representa no Winthor do cliente (validade do lote, do produto ou outra data)? A análise de validade da onda 2 depende disso.

---

# Carga: janela e tabelas de origem

**Janela:** `data >= 01/jan do ano anterior` (configurável por tenant), mais reprocessamento dos últimos N dias para pegar alterações em notas e pedidos antigos. As exceções são os títulos em aberto (receber e pagar), que entram sempre, e o estoque (posição atual, sem janela). Cadastros são carregados por inteiro.

**Tabelas de origem (raw), cerca de 45:**

- **Onda 1:** `PCCLIENT`, `PCUSUARI`, `PCUSURCLI`, `PCSUPERV`, `PCGERENTE`, `PCPRACA`, `PCROTAEXP`, `PCREGIAO`, `PCATIVI`, `PCCIDADE`, `PCCONSUM`, `PCFILIAL`, `PCFORNEC`, `PCPRODUT`, `PCDEPTO`, `PCSECAO`, `PCCATEGORIA`, `PCSUBCATEGORIA`, `PCMARCA`, `PCLINHAPROD`, `PCTABDEV`, `PCDATAS`, `PCNFSAID`, `PCNFENT`, `PCESTCOM`, `PCMOV`, `PCMOVCOMPLE`, `PCPEDC`, `PCPEDI`, `PCMETA`, `PCMETARCA`.
- **Onda 2:** `PCPREST`, `PCEST`, `PCTABPR`, `PCPARAMFILIAL`, `PCEMBALAGEM`, `PCCORTEI`, `PCFALTA`.
- **Onda 3:** `PCPEDIDO`, `PCITEM`, `PCLANC`, `PCCONTA`, `PCLIB`, `PCRATEIOCENTROCUSTO`, `PCCENTROCUSTO`, `PCEMPR`.

Nas tabelas largas (`PCCLIENT`, `PCPRODUT`, `PCMOV`, `PCPEDC`, `PCNFSAID`), extrair só as colunas usadas.