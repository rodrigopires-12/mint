# Inventário das tabelas de origem (E1.2)

Gerado por `gateway inventory` em 2026-10-25 19:56 contra Oracle Database 19c Standard Edition 2  19.0.0.0.0. Os arquivos por tabela estão em `sources/winthor/<TABELA>.yaml` (status `proposed` até a aprovação) e a lista completa de colunas de cada tabela em `docs/inventory/<TABELA>.md`.

## Resumo

| Tabela | Onda | Owner | Linhas (dicionário) | COUNT(*) | MB est. | Chave | Estratégia | Watermark | Partição | Riscos |
|---|---|---|---|---|---|---|---|---|---|---|
| PCCLIENT | 1 | RIBERDOCES | 283076 | 283076 | 517 | CODCLI [constraint:PCCLIENT_PK] | full |  |  | dictionary_has_lob_or_long, wide_table |
| PCUSUARI | 1 | RIBERDOCES | 397 | 397 | 0 | CODUSUR [constraint:PCUSUARI_PK] | full |  |  |  |
| PCUSURCLI | 1 | RIBERDOCES | 15 | 15 | 0 | CODUSUR, CODCLI [constraint:PCUSURCLI_PK] | full |  |  |  |
| PCSUPERV | 1 | RIBERDOCES | 50 | 50 | 0 | CODSUPERVISOR [constraint:PCSUPERV_PK] | full |  |  |  |
| PCGERENTE | 1 | RIBERDOCES | 3 | 3 | 0 | CODGERENTE [constraint:PCGERENTE_PK] | full |  |  |  |
| PCPRACA | 1 | RIBERDOCES | 1125 | 1125 | 0 | CODPRACA [constraint:PCPRACA_PK] | full |  |  |  |
| PCROTAEXP | 1 | RIBERDOCES | 147 | 147 | 0 | CODROTA [constraint:PCROTAEXP_PK] | full |  |  |  |
| PCREGIAO | 1 | RIBERDOCES | 23 | 23 | 0 | NUMREGIAO [constraint:PCREGIAO_PK] | full |  |  |  |
| PCATIVI | 1 | RIBERDOCES | 21 | 21 | 0 | CODATIV [constraint:PCATIVI_PK] | full |  |  |  |
| PCCIDADE | 1 | RIBERDOCES | 5568 | 5568 | 0 | CODCIDADE [constraint:PCCIDADE_PK] | full |  |  |  |
| PCCONSUM | 1 | RIBERDOCES | 1 | 1 | 0 | (nenhuma) | full |  |  | no_key, wide_table |
| PCFILIAL | 1 | RIBERDOCES | 12 | 12 | 0 | CODIGO [constraint:PCFILIAL_PK] | full |  |  | dictionary_has_lob_or_long, wide_table |
| PCFORNEC | 1 | RIBERDOCES | 11107 | 11107 | 10 | CODFORNEC [constraint:PCFORNEC_PK] | full |  |  | dictionary_has_lob_or_long, wide_table |
| PCPRODUT | 1 | RIBERDOCES | 18567 | 18567 | 37 | CODPROD [constraint:PCPRODUT_PK] | full |  |  | dictionary_has_lob_or_long, wide_table |
| PCDEPTO | 1 | RIBERDOCES | 25 | 25 | 0 | CODEPTO [constraint:PCDEPTO_PK] | full |  |  |  |
| PCSECAO | 1 | RIBERDOCES | 139 | 139 | 0 | CODSEC [constraint:PCSECAO_PK] | full |  |  |  |
| PCCATEGORIA | 1 | RIBERDOCES | 157 | 157 | 0 | CODSEC, CODCATEGORIA [constraint:PCCATEGORIA_PK] | full |  |  |  |
| PCSUBCATEGORIA | 1 | RIBERDOCES | 1365 | 1365 | 0 | CODSEC, CODCATEGORIA, CODSUBCATEGORIA [constraint:PCSUBCATEGORIA_PK] | full |  |  |  |
| PCMARCA | 1 | RIBERDOCES | 564 | 564 | 0 | CODMARCA [constraint:PCMARCA_PK] | full |  |  |  |
| PCLINHAPROD | 1 | RIBERDOCES | 0 | 0 | 0 | CODLINHA [constraint:PCLINHAPROD_PK] | full |  |  |  |
| PCTABDEV | 1 | RIBERDOCES | 129 | 129 | 0 | CODDEVOL [constraint:PCTABDEV_PK] | full |  |  |  |
| PCDATAS | 1 | RIBERDOCES | 7824 | 7824 | 0 | DATA [constraint:PCDATAS_PK] | full |  |  |  |
| PCNFSAID | 1 | RIBERDOCES | 7030552 | 7030552 | 16912 | NUMTRANSVENDA [constraint:PCNFSAID_PK] | window | DTSAIDA (date) | DTSAIDA (date) | dictionary_has_lob_or_long, wide_table |
| PCNFENT | 1 | RIBERDOCES | 535056 | 535056 | 1015 | NUMTRANSENT, CODCONT [constraint:PCNFENT_PK] | window | DTENT (date) | DTENT (date) | dictionary_has_lob_or_long, wide_table |
| PCESTCOM | 1 | RIBERDOCES | 137757 | 137757 | 10 | NUMTRANSENT [constraint:PCESTCOM_PK] | full |  |  |  |
| PCMOV | 1 | RIBERDOCES | 64017288 | 64017288 | 68860 | NUMTRANSITEM [constraint:PCMOV_UK] | window | DTMOV (date) | DTMOV (date) | nullable_key_column, dictionary_has_lob_or_long, wide_table, key_has_nulls |
| PCMOVCOMPLE | 1 | RIBERDOCES | 63250854 | 63250854 | 45455 | NUMTRANSITEM [constraint:PCMOVCOMPLE_PK] | window |  | NUMTRANSITEM (number) | incremental_without_watermark, wide_table |
| PCPEDC | 1 | RIBERDOCES | 6917841 | 6917849 | 5819 | NUMPED [constraint:PCPEDC_PK] | window | DATA (date) | DATA (date) | dictionary_has_lob_or_long, wide_table |
| PCPEDI | 1 | RIBERDOCES | 37589500 | 37589594 | 21368 | NUMPED, CODPROD, NUMSEQ [constraint:PCPEDI_PK] | window | DATA (date) | DATA (date) | dictionary_has_lob_or_long, wide_table |
| PCMETA | 1 | RIBERDOCES | 3385 | 3385 | 0 | CODFILIAL, CODIGO, CODIGO2, CODUSUR, TIPOMETA, DATA [constraint:PCMETA_PK] | full |  |  |  |
| PCMETARCA | 1 | RIBERDOCES | 296975 | 296975 | 13 | CODFILIAL, CODUSUR, DATA [constraint:PCMETARCA_PK] | full |  |  |  |
| PCPREST | 2 | RIBERDOCES | 10792794 | 10792794 | 8697 | NUMTRANSVENDA, PREST [constraint:PCPREST_PK] | open_items | DTEMISSAO (date) | DTEMISSAO (date) | wide_table |
| PCEST | 2 | RIBERDOCES | 195721 | 195721 | 63 | CODFILIAL, CODPROD [constraint:PCEST_PK] | full |  |  |  |
| PCTABPR | 2 | RIBERDOCES | 424644 | 424644 | 223 | NUMREGIAO, CODPROD [constraint:PCTABPR_PK] | full |  |  | dictionary_has_lob_or_long |
| PCPARAMFILIAL | 2 | RIBERDOCES | 16490 | 16490 | 1 | NOME, CODFILIAL [constraint:PCPARAMFILIAL_PK] | full |  |  |  |
| PCEMBALAGEM | 2 | RIBERDOCES | 257800 | 257800 | 71 | CODFILIAL, CODAUXILIAR [constraint:PCEMBALAGEM_PK] | full |  |  | dictionary_has_lob_or_long |
| PCCORTEI | 2 | RIBERDOCES | 207520 | 207520 | 20 | CODPROD, NUMCAR, NUMPED, NUMSEQ [constraint:PCCORTEI_PK] | window | DATA (date) | DATA (date) |  |
| PCFALTA | 2 | RIBERDOCES | 30967 | 30967 | 2 | NUMPED, CODPROD, NUMSEQ [constraint:PCFALTA1_PK] | full | DATA (date) |  | watermark_not_indexed |
| PCPEDIDO | 3 | RIBERDOCES | 134393 | 134399 | 58 | NUMPED [constraint:PCPEDIDO_PK] | window | DTEMISSAO (date) | DTEMISSAO (date) | dictionary_has_lob_or_long, wide_table |
| PCITEM | 3 | RIBERDOCES | 1375271 | 1375283 | 749 | NUMPED, CODPROD, NUMSEQ [constraint:PCITEM_PK] | window |  | NUMPED (number) | incremental_without_watermark, wide_table |
| PCLANC | 3 | RIBERDOCES | 3308388 | 3308391 | 1561 | RECNUM [constraint:PCLANC_PK] | open_items | DTLANC (date) | DTLANC (date) | dictionary_has_lob_or_long, wide_table, watermark_has_nulls |
| PCCONTA | 3 | RIBERDOCES | 504 | 504 | 0 | CODCONTA [constraint:PCCONTA_PK] | full |  |  |  |
| PCLIB | 3 | RIBERDOCES | 1252336 | 1252336 | 45 | CODFUNC, CODTABELA, CODIGOA, CODIGON [constraint:PCLIB_PK] | full |  |  |  |
| PCRATEIOCENTROCUSTO | 3 | RIBERDOCES | 12050 | 12050 | 2 | RECNUM, CODIGOCENTROCUSTO [proposed] | full |  |  | no_unique_key_in_dictionary, nullable_key_column |
| PCCENTROCUSTO | 3 | RIBERDOCES | 6 | 6 | 0 | CODIGOCENTROCUSTO [constraint:PK_PCCENTROCUSTO] | full |  |  |  |
| PCEMPR | 3 | RIBERDOCES | 2404 | 2404 | 2 | MATRICULA [constraint:PCEMPR_PK] | full |  |  | dictionary_has_lob_or_long |

## Respostas aos pontos do escopo que dependem só do dicionário

### #1 — PCDATAS tem uma linha por data ou por filial?

- Colunas: DATA DATE, AUXTMP NUMBER(15,2), DIAUTIL VARCHAR2(1)
- Chave primária: PCDATAS_PK (DATA)
- Coluna CODFILIAL existe: false
- COUNT(*) = 7824, COUNT(DISTINCT DATA) = 7824

**Resposta:** Uma linha por data: a data é única na tabela.

### #2 — Chaves reais: NUMTRANSITEM em PCMOV, NUMSEQ em PCPEDI e PCITEM, chave de PCSECAO

- PCMOV: sem PK; únicos: PCMOV_UK(NUMTRANSITEM); unicidade verificada da chave adotada (NUMTRANSITEM): true (644975 linhas com chave nula)
- PCPEDI: PK PCPEDI_PK (NUMPED, CODPROD, NUMSEQ); unicidade verificada da chave adotada (NUMPED, CODPROD, NUMSEQ): true
- PCITEM: PK PCITEM_PK (NUMPED, CODPROD, NUMSEQ); unicidade verificada da chave adotada (NUMPED, CODPROD, NUMSEQ): true
- PCSECAO: PK PCSECAO_PK (CODSEC); unicidade verificada da chave adotada (CODSEC): true

**Resposta:** Chaves adotadas por tabela: PCMOV = (NUMTRANSITEM) [constraint:PCMOV_UK]; PCPEDI = (NUMPED, CODPROD, NUMSEQ) [constraint:PCPEDI_PK]; PCITEM = (NUMPED, CODPROD, NUMSEQ) [constraint:PCITEM_PK]; PCSECAO = (CODSEC) [constraint:PCSECAO_PK]

### #11 — Colunas de PCRATEIOCENTROCUSTO (valor ou percentual de rateio)

- Colunas: RECNUM NUMBER(8), CODCONTA NUMBER(10), CODCENTROCUSTO NUMBER(10), VALOR NUMBER(12,2), PERCRATEIO NUMBER(10,6), DTLANC DATE, CONTRAPARTIDA VARCHAR2(1), RECNUMPRINC NUMBER(10), CODFILIAL VARCHAR2(2), CODIGOCENTROCUSTO VARCHAR2(40), ROTINAUPDATE VARCHAR2(100), ROTINAINSERT VARCHAR2(100), CODFUNCAUTSUPLORC NUMBER(8), DATAAUTSUPLORC DATE, ORCAMENTOEXCEDIDO VARCHAR2(1), NUMTRANSPISCOFINS NUMBER(10), RECNUMREF NUMBER(10), PERCRATCONTREF NUMBER(10,6), TIPOTRIB VARCHAR2(20), LANCIMPRETIDO VARCHAR2(1)

**Resposta:** Colunas candidatas ao rateio: VALOR NUMBER(12,2), PERCRATEIO NUMBER(10,6), PERCRATCONTREF NUMBER(10,6). Confirmar a semântica com o negócio.

### #12 — Tipos reais das chaves e versão do Oracle

- Versão: Oracle Database 19c Standard Edition 2  19.0.0.0.0
- CODFILIAL: VARCHAR2(2) em 20 tabelas (PCUSUARI, PCREGIAO, PCPRODUT, PCLINHAPROD, PCNFSAID, PCNFENT, PCMOV, PCPEDC, PCMETA, PCMETARCA, PCPREST, PCEST, PCPARAMFILIAL, PCEMBALAGEM, PCCORTEI, PCFALTA, PCPEDIDO, PCLANC, PCRATEIOCENTROCUSTO, PCEMPR)
- CODCLI: NUMBER(9) em 12 tabelas (PCCLIENT, PCUSURCLI, PCFILIAL, PCFORNEC, PCNFSAID, PCMOV, PCPEDC, PCPEDI, PCPREST, PCEMBALAGEM, PCCORTEI, PCFALTA)
- CODPROD: NUMBER(6) em 9 tabelas (PCPRODUT, PCMOV, PCPEDI, PCEST, PCTABPR, PCEMBALAGEM, PCCORTEI, PCFALTA, PCITEM)
- CODUSUR: NUMBER(4) em 12 tabelas (PCUSUARI, PCUSURCLI, PCNFSAID, PCESTCOM, PCPEDC, PCPEDI, PCMETA, PCMETARCA, PCPREST, PCCORTEI, PCFALTA, PCEMPR); NUMBER(8) em 1 tabelas (PCMOV)
- NUMTRANSITEM: NUMBER(18) em 2 tabelas (PCMOV, PCMOVCOMPLE)
- NUMTRANSVENDA: NUMBER(10) em 7 tabelas (PCNFSAID, PCESTCOM, PCMOV, PCPEDC, PCPREST, PCPEDIDO, PCLANC)
- NUMPED: NUMBER(10) em 8 tabelas (PCNFSAID, PCMOV, PCPEDC, PCPEDI, PCPREST, PCCORTEI, PCFALTA, PCITEM); NUMBER(6) em 1 tabelas (PCPEDIDO)
- RECNUM: NUMBER(8) em 2 tabelas (PCLANC, PCRATEIOCENTROCUSTO)

**Resposta:** Tipos por coluna acima; cod_filial é texto quando CODFILIAL for VARCHAR2 em todas as tabelas.

### #13 — A nota de entrada guarda o pedido de compra de origem?

- PCNFENT: colunas com PED no nome: VLPEDAGIO NUMBER(12,2), TIPOEMBALAGEMPEDIDO VARCHAR2(1), NUMPEDPRINC NUMBER(8)
- PCMOV: colunas com PED no nome: NUMPED NUMBER(10), NUMSEQPED NUMBER(6), PLIQPED NUMBER(18,6), TIPOEMBALAGEMPEDIDO VARCHAR2(1)
- PCMOV.NUMPED preenchido em 58817426 de 64017288 linhas

**Resposta:** Ver as colunas candidatas e a taxa de preenchimento; confirmar com o negócio qual delas liga a entrada ao pedido de compra.

### #14 — O que PCPRODUT.DTVENC representa?

- DTVENC DATE, comentário no dicionário: ""
- Preenchida em 18 de 18567 produtos
- Mínimo 1899-12-30, máximo 2019-07-04

**Resposta:** O dicionário não define a semântica; a taxa de preenchimento e a faixa de datas indicam se é validade de produto (poucos valores, futuros) ou outra data. Confirmar com o negócio.

### #Q9 — Colunas de controle de alteração (candidatas a watermark secundário)

- PCCLIENT: DTULTALTER, DTULTALTERCLIESP, DTULTALTERSRVPRC, DTULTALTER1203, DTULTALTERCLASSE, DTULTALTERSOFITVIEW, RBULTALTER, DTALTERC5
- PCUSUARI: DTULTALTERACAO
- PCROTAEXP: DTULTALTERSOFITVIEW
- PCREGIAO: DTALTERC5, DTALTERACAO, DTCRIACAO
- PCATIVI: DATAALTERACAO
- PCCONSUM: DTATUALIZACAOBD
- PCFILIAL: DTULTATUALIZACAODRVINTER, DTULTALTER, DTALTERC5
- PCFORNEC: DTULTALTER, DTULTALTERSRVPRC, DTULTALTERSOFITVIEW
- PCPRODUT: DTULTALTER, DTULTALTERSRVPRC, DTULTALTERSOFITVIEW, DTALTERC5
- PCDEPTO: DTULTALTER, DTALTERC5
- PCSECAO: DTULTALTER, DTALTERC5
- PCCATEGORIA: DTULTALTER, DTALTERC5
- PCSUBCATEGORIA: DTULTALTER, DTALTERC5
- PCMARCA: DTULTALTER, DTALTERC5
- PCPEDI: DTULTLALTER, DTALTERACAOOS
- PCPREST: DTULTALTER, DTINCLUSAO, DTCRIACAO, DTINCLUSAOMANUAL
- PCEST: DTULTALTERSRVPRC, DTALTERACAO1107
- PCTABPR: DTULTALTERSRVPRC, DTALTERC5
- PCEMBALAGEM: DTULTALTERSRVPRC, DTALTERC5, DTULALTERINTEGRA
- PCLANC: DTULTALTER
- PCCENTROCUSTO: DTINCLUSAO, DTALTERACAO
- PCEMPR: DTULTALTERSOFITVIEW, DTULTALTER, DTALTERC5

**Resposta:** Só as tabelas listadas têm carimbo de alteração no dicionário; nas demais o watermark é a data de negócio mais a janela móvel. Resta validar, por teste controlado, se o carimbo muda a cada alteração.

## Detalhe por tabela

### PCCLIENT (onda 1) — Clientes (sem CPF/CNPJ, e-mail e telefone)

- Owner `RIBERDOCES`, 587 colunas no dicionário, 15 extraídas. Estratégia `full`.
- Chave: CODCLI [constraint:PCCLIENT_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 283076 (analisado em 2026-10-25), COUNT(*) 283076, tamanho estimado 517 MB, linha média 1288 bytes.
- Riscos: dictionary_has_lob_or_long (3, not extracted): OBS_ADIC CLOB, IDINTEGRACAOMYFROTA RAW, MOTIVOBLOQ CLOB; wide_table: 587 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODCLI | NUMBER(9) | Int64 |  | key | dim_cliente.cod_cliente |  |
| CLIENTE | VARCHAR2(60) | String | S | scope | dim_cliente.nome_cliente |  |
| FANTASIA | VARCHAR2(40) | String | S | scope | dim_cliente.nome_fantasia |  |
| CODCLIPRINC | NUMBER(9) | Int64 | S | scope | dim_cliente.cod_cliente_principal |  |
| TIPOFJ | VARCHAR2(1) | String | S | scope | dim_cliente.tipo_pessoa |  |
| CODATV1 | NUMBER(6) | Int64 | S | scope | dim_cliente.cod_ramo_atividade |  |
| CODPRACA | NUMBER(6) | Int64 |  | scope | dim_cliente.cod_praca |  |
| CODCIDADE | NUMBER(6) | Int64 | S | proposed | dim_cliente.cod_cidade | liga a PCCIDADE; o escopo cita a tabela mas não a coluna |
| BAIRROENT | VARCHAR2(40) | String | S | scope | dim_cliente.bairro |  |
| DTCADASTRO | DATE | DateTime64(0) | S | scope | dim_cliente.data_cadastro |  |
| DTULTCOMP | DATE | DateTime64(0) | S | scope | dim_cliente.data_ultima_compra |  |
| DTEXCLUSAO | DATE | DateTime64(0) | S | scope | dim_cliente.data_exclusao |  |
| CODUSUR1 | NUMBER(4) | Int64 | S | scope | ponte_carteira |  |
| CODUSUR2 | NUMBER(4) | Int64 | S | scope | ponte_carteira |  |
| CODUSUR3 | NUMBER(4) | Int64 | S | scope | ponte_carteira |  |

### PCUSUARI (onda 1) — RCAs (vendedores)

- Owner `RIBERDOCES`, 188 colunas no dicionário, 6 extraídas. Estratégia `full`.
- Chave: CODUSUR [constraint:PCUSUARI_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 397 (analisado em 2026-10-25), COUNT(*) 397, tamanho estimado 0 MB, linha média 382 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODUSUR | NUMBER(4) | Int64 |  | key | dim_rca.cod_rca |  |
| NOME | VARCHAR2(40) | String |  | scope | dim_rca.nome_rca |  |
| EMAIL | VARCHAR2(100) | String | S | scope | dim_rca.email_rca |  |
| TIPOVEND | VARCHAR2(2) | String | S | scope | dim_rca.tipo_rca |  |
| DTTERMINO | DATE | DateTime64(0) | S | scope | dim_rca.data_termino |  |
| CODSUPERVISOR | NUMBER(4) | Int64 |  | proposed | dim_rca.cod_supervisor | liga a PCSUPERV; o escopo cita a tabela mas não a coluna |

### PCUSURCLI (onda 1) — Vínculos adicionais RCA x cliente

- Owner `RIBERDOCES`, 5 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: CODUSUR, CODCLI [constraint:PCUSURCLI_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 15 (analisado em 2026-10-25), COUNT(*) 15, tamanho estimado 0 MB, linha média 10 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODUSUR | NUMBER(4) | Int64 |  | key | ponte_carteira.cod_rca |  |
| CODCLI | NUMBER(9) | Int64 |  | key | ponte_carteira.cod_cliente |  |

### PCSUPERV (onda 1) — Supervisores

- Owner `RIBERDOCES`, 18 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODSUPERVISOR [constraint:PCSUPERV_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 50 (analisado em 2026-10-25), COUNT(*) 50, tamanho estimado 0 MB, linha média 76 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODSUPERVISOR | NUMBER(4) | Int64 |  | key | dim_rca.cod_supervisor |  |
| NOME | VARCHAR2(40) | String |  | scope | dim_rca.nome_supervisor |  |
| EMAIL | VARCHAR2(100) | String | S | scope | dim_rca.email_supervisor |  |
| CODGERENTE | NUMBER(4) | Int64 | S | proposed | dim_rca.cod_gerente | liga a PCGERENTE; o escopo cita a tabela mas não a coluna |

### PCGERENTE (onda 1) — Gerentes

- Owner `RIBERDOCES`, 10 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: CODGERENTE [constraint:PCGERENTE_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 3 (analisado em 2026-10-25), COUNT(*) 3, tamanho estimado 0 MB, linha média 55 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODGERENTE | NUMBER(4) | Int64 |  | key | dim_rca.cod_gerente |  |
| NOMEGERENTE | VARCHAR2(40) | String | S | scope | dim_rca.nome_gerente | o escopo cita NOME; na base a coluna é NOMEGERENTE |
| EMAIL | VARCHAR2(100) | String | S | scope | dim_rca.email_gerente |  |

### PCPRACA (onda 1) — Praças

- Owner `RIBERDOCES`, 25 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODPRACA [constraint:PCPRACA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 1125 (analisado em 2026-10-25), COUNT(*) 1125, tamanho estimado 0 MB, linha média 52 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODPRACA | NUMBER(4) | Int64 |  | key | dim_cliente.cod_praca |  |
| PRACA | VARCHAR2(25) | String |  | scope | dim_cliente.praca |  |
| NUMREGIAO | NUMBER(4) | Int64 |  | proposed | dim_cliente.num_regiao | região via praça; o escopo cita o caminho mas não a coluna |
| ROTA | NUMBER(4) | Int64 |  | proposed | dim_cliente.cod_rota | rota via praça (PCROTAEXP); o escopo cita o caminho mas não a coluna |

### PCROTAEXP (onda 1) — Rotas

- Owner `RIBERDOCES`, 75 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: CODROTA [constraint:PCROTAEXP_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 147 (analisado em 2026-10-25), COUNT(*) 147, tamanho estimado 0 MB, linha média 78 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODROTA | NUMBER(4) | Int64 |  | key | dim_cliente.cod_rota |  |
| DESCRICAO | VARCHAR2(40) | String | S | proposed | dim_cliente.rota | nome da rota; o escopo não cita a coluna |

### PCREGIAO (onda 1) — Regiões

- Owner `RIBERDOCES`, 33 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: NUMREGIAO [constraint:PCREGIAO_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 23 (analisado em 2026-10-25), COUNT(*) 23, tamanho estimado 0 MB, linha média 72 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMREGIAO | NUMBER(4) | Int64 |  | key | dim_regiao.num_regiao |  |
| REGIAO | VARCHAR2(40) | String |  | scope | dim_regiao.regiao |  |
| UF | VARCHAR2(2) | String | S | proposed | dim_regiao.uf | o escopo pede uf na dim_regiao; confirmar o nome da coluna |
| STATUS | VARCHAR2(1) | String | S | proposed | dim_regiao.status | o escopo pede status na dim_regiao; confirmar o nome da coluna |

### PCATIVI (onda 1) — Ramos de atividade

- Owner `RIBERDOCES`, 15 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: CODATIV [constraint:PCATIVI_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 21 (analisado em 2026-10-25), COUNT(*) 21, tamanho estimado 0 MB, linha média 35 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODATIV | NUMBER(6) | Int64 |  | key | dim_cliente.cod_ramo_atividade |  |
| RAMO | VARCHAR2(40) | String |  | proposed | dim_cliente.ramo_atividade | descrição do ramo; o escopo não cita a coluna |

### PCCIDADE (onda 1) — Cidades

- Owner `RIBERDOCES`, 10 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODCIDADE [constraint:PCCIDADE_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 5568 (analisado em 2026-10-25), COUNT(*) 5568, tamanho estimado 0 MB, linha média 50 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODCIDADE | NUMBER(6) | Int64 |  | key | dim_cliente.cod_cidade |  |
| NOMECIDADE | VARCHAR2(80) | String | S | proposed | dim_cliente.cidade |  |
| UF | VARCHAR2(2) | String | S | proposed | dim_cliente.uf_cidade |  |
| CODIBGE | NUMBER(10) | Int64 | S | proposed | dim_cliente.cod_ibge |  |

### PCCONSUM (onda 1) — Parâmetros gerais (uma linha)

- Owner `RIBERDOCES`, 775 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: (nenhuma).
- Linhas: dicionário 1 (analisado em 2026-10-25), COUNT(*) 1, tamanho estimado 0 MB, linha média 1969 bytes.
- Riscos: no_key; wide_table: 775 columns in dictionary
- Pergunta: PCCONSUM não tem chave primária (tabela de uma linha). Tratar como singleton: full refresh e chave sintética.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMDIASCLIINATIV | NUMBER(4) | Int64 | S | scope | dim_cliente.dias_param_inatividade |  |
| BLOQUEIAVENDAESTPENDENTE | VARCHAR2(1) | String | S | proposed | fato_estoque_snapshot.qtd_disponivel | o escopo cita o parâmetro sem dizer a tabela; pode estar em PCCONSUM ou PCPARAMFILIAL |

### PCFILIAL (onda 1) — Filiais

- Owner `RIBERDOCES`, 472 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODIGO [constraint:PCFILIAL_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 12 (analisado em 2026-10-25), COUNT(*) 12, tamanho estimado 0 MB, linha média 983 bytes.
- Riscos: dictionary_has_lob_or_long (2, not extracted): LOGOMARCA BLOB, CERTIFICADOA1 BLOB; wide_table: 472 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODIGO | VARCHAR2(2) | String |  | key | dim_filial.cod_filial |  |
| RAZAOSOCIAL | VARCHAR2(40) | String |  | scope | dim_filial.nome_filial |  |
| FANTASIA | VARCHAR2(25) | String | S | scope | dim_filial.nome_fantasia |  |
| NUMREGIAOPADRAO | NUMBER(4) | Int64 | S | scope | dim_filial.num_regiao_padrao |  |

### PCFORNEC (onda 1) — Fornecedores

- Owner `RIBERDOCES`, 363 colunas no dicionário, 5 extraídas. Estratégia `full`.
- Chave: CODFORNEC [constraint:PCFORNEC_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 11107 (analisado em 2026-10-25), COUNT(*) 11107, tamanho estimado 10 MB, linha média 668 bytes.
- Riscos: dictionary_has_lob_or_long (1, not extracted): IDINTEGRACAOMYFROTA RAW; wide_table: 363 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODFORNEC | NUMBER(6) | Int64 |  | key | dim_fornecedor.cod_fornecedor |  |
| FORNECEDOR | VARCHAR2(60) | String |  | scope | dim_fornecedor.nome_fornecedor |  |
| CODFORNECPRINC | NUMBER(6) | Int64 | S | scope | dim_fornecedor.cod_fornecedor_principal |  |
| CODCOMPRADOR | NUMBER(8) | Int64 | S | scope | dim_fornecedor.cod_comprador |  |
| PRAZOENTREGA | NUMBER(4) | Int64 | S | scope | dim_fornecedor.prazo_entrega_dias |  |

### PCPRODUT (onda 1) — Produtos

- Owner `RIBERDOCES`, 805 colunas no dicionário, 16 extraídas. Estratégia `full`.
- Chave: CODPROD [constraint:PCPRODUT_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 18567 (analisado em 2026-10-25), COUNT(*) 18567, tamanho estimado 37 MB, linha média 1129 bytes.
- Riscos: dictionary_has_lob_or_long (2, not extracted): DADOSTECNICOS CLOB, IDINTEGRACAOMYFROTA RAW; wide_table: 805 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODPROD | NUMBER(6) | Int64 |  | key | dim_produto.cod_produto |  |
| DESCRICAO | VARCHAR2(40) | String |  | scope | dim_produto.nome_produto |  |
| CODEPTO | NUMBER(6) | Int64 |  | proposed | dim_produto.cod_departamento | liga a PCDEPTO |
| CODSEC | NUMBER(6) | Int64 |  | proposed | dim_produto.cod_secao | liga a PCSECAO |
| CODCATEGORIA | NUMBER(6) | Int64 | S | proposed | dim_produto.cod_categoria | liga a PCCATEGORIA (com CODSEC) |
| CODSUBCATEGORIA | NUMBER(6) | Int64 | S | proposed | dim_produto.cod_subcategoria | liga a PCSUBCATEGORIA |
| CODMARCA | NUMBER(8) | Int64 | S | proposed | dim_produto.cod_marca | liga a PCMARCA |
| CODLINHAPROD | NUMBER(6) | Int64 | S | proposed | dim_produto.cod_linha | liga a PCLINHAPROD; confirmar nome |
| CODFORNEC | NUMBER(6) | Int64 |  | scope | dim_produto.cod_fornecedor |  |
| QTUNITCX | NUMBER(8,2) | Decimal(8,2) | S | scope | dim_produto.qtd_unidades_caixa |  |
| PESOBRUTO | NUMBER(12,6) | Decimal(12,6) | S | scope | dim_produto.peso_bruto_unit |  |
| PESOLIQ | NUMBER(12,6) | Decimal(12,6) | S | scope | dim_produto.peso_liquido_unit |  |
| VOLUME | NUMBER(20,8) | Decimal(20,8) | S | scope | dim_produto.volume_unit |  |
| DTVENC | DATE | DateTime64(0) | S | scope | dim_produto.data_vencimento |  |
| DTEXCLUSAO | DATE | DateTime64(0) | S | scope | dim_produto.data_exclusao |  |
| TEMREPOS | NUMBER(4) | Int64 | S | scope | fato_estoque_snapshot.tempo_reposicao_dias |  |

### PCDEPTO (onda 1) — Departamentos

- Owner `RIBERDOCES`, 24 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: CODEPTO [constraint:PCDEPTO_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 25 (analisado em 2026-10-25), COUNT(*) 25, tamanho estimado 0 MB, linha média 55 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODEPTO | NUMBER(6) | Int64 |  | key | dim_produto.cod_departamento |  |
| DESCRICAO | VARCHAR2(25) | String |  | proposed | dim_produto.departamento |  |
| TIPOMERC | VARCHAR2(2) | String | S | scope | dim_produto.tipo_mercadoria |  |

### PCSECAO (onda 1) — Seções (chave a confirmar, ponto 2 do escopo)

- Owner `RIBERDOCES`, 13 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: CODSEC [constraint:PCSECAO_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 139 (analisado em 2026-10-25), COUNT(*) 139, tamanho estimado 0 MB, linha média 45 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODSEC | NUMBER(6) | Int64 |  | key | dim_produto.cod_secao |  |
| CODEPTO | NUMBER(6) | Int64 |  | proposed | dim_produto.cod_departamento | necessária se a chave for CODEPTO + CODSEC |
| DESCRICAO | VARCHAR2(40) | String |  | proposed | dim_produto.secao |  |

### PCCATEGORIA (onda 1) — Categorias (chave composta seção + categoria)

- Owner `RIBERDOCES`, 8 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: CODSEC, CODCATEGORIA [constraint:PCCATEGORIA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 157 (analisado em 2026-10-25), COUNT(*) 157, tamanho estimado 0 MB, linha média 47 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODSEC | NUMBER(6) | Int64 |  | key | dim_produto.cod_secao |  |
| CODCATEGORIA | NUMBER(6) | Int64 |  | key | dim_produto.cod_categoria |  |
| CATEGORIA | VARCHAR2(40) | String | S | proposed | dim_produto.categoria |  |

### PCSUBCATEGORIA (onda 1) — Subcategorias (chave composta seção + categoria + subcategoria)

- Owner `RIBERDOCES`, 9 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODSEC, CODCATEGORIA, CODSUBCATEGORIA [constraint:PCSUBCATEGORIA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 1365 (analisado em 2026-10-25), COUNT(*) 1365, tamanho estimado 0 MB, linha média 32 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODSEC | NUMBER(6) | Int64 |  | key | dim_produto.cod_secao |  |
| CODCATEGORIA | NUMBER(6) | Int64 |  | key | dim_produto.cod_categoria |  |
| CODSUBCATEGORIA | NUMBER(6) | Int64 |  | key | dim_produto.cod_subcategoria |  |
| SUBCATEGORIA | VARCHAR2(40) | String | S | proposed | dim_produto.subcategoria |  |

### PCMARCA (onda 1) — Marcas

- Owner `RIBERDOCES`, 11 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: CODMARCA [constraint:PCMARCA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 564 (analisado em 2026-10-25), COUNT(*) 564, tamanho estimado 0 MB, linha média 26 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODMARCA | NUMBER(8) | Int64 |  | key | dim_produto.cod_marca |  |
| MARCA | VARCHAR2(40) | String | S | proposed | dim_produto.marca |  |

### PCLINHAPROD (onda 1) — Linhas de produto

- Owner `RIBERDOCES`, 10 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: CODLINHA [constraint:PCLINHAPROD_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 0 (analisado em 2026-10-25), COUNT(*) 0, tamanho estimado 0 MB, linha média 0 bytes.
- Pergunta: A chave primária do dicionário (PCLINHAPROD_PK: CODLINHA) difere da esperada pelo escopo (CODLINHAPROD). O inventário adota a do dicionário.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODLINHAPROD | NUMBER(6) | Int64 | S | key | dim_produto.cod_linha |  |
| DESCRICAO | VARCHAR2(40) | String | S | proposed | dim_produto.linha |  |
| CODLINHA | NUMBER(6) | Int64 |  | key |  | added by inventory: required for key |

### PCTABDEV (onda 1) — Motivos de devolução

- Owner `RIBERDOCES`, 14 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: CODDEVOL [constraint:PCTABDEV_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 129 (analisado em 2026-10-25), COUNT(*) 129, tamanho estimado 0 MB, linha média 42 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODDEVOL | NUMBER(4) | Int64 |  | key | dim_motivo_devolucao.cod_motivo_devolucao |  |
| MOTIVO | VARCHAR2(30) | String |  | proposed | dim_motivo_devolucao.motivo_devolucao |  |

### PCDATAS (onda 1) — Calendário de dias úteis (ponto 1 do escopo)

- Owner `RIBERDOCES`, 3 colunas no dicionário, 2 extraídas. Estratégia `full`.
- Chave: DATA [constraint:PCDATAS_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 7824 (analisado em 2026-10-25), COUNT(*) 7824, tamanho estimado 0 MB, linha média 18 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| DATA | DATE | DateTime64(0) |  | key | dim_data.data |  |
| DIAUTIL | VARCHAR2(1) | String | S | scope | dim_data.eh_dia_util |  |

### PCNFSAID (onda 1) — Notas fiscais de saída (cabeçalho)

- Owner `RIBERDOCES`, 670 colunas no dicionário, 15 extraídas. Estratégia `window`, watermark `DTSAIDA` (date), reprocessamento de 45 dias.
- Chave: NUMTRANSVENDA [constraint:PCNFSAID_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 7030552 (analisado em 2026-10-25), COUNT(*) 7030552, tamanho estimado 16912 MB, linha média 1901 bytes.
- Watermark: mínimo 2009-10-02, máximo 2026-10-24, nulos 0.
- Distribuição por mês (watermark): 2023-10=53262, 2023-11=50072, 2023-12=51184, 2024-01=47431, 2024-02=47037, 2024-03=51198, 2024-04=51229, 2024-05=52258, 2024-06=49492, 2024-07=56059, 2024-08=54280, 2024-09=53063, 2024-10=58621, 2024-11=51720, 2024-12=56005, 2025-01=51358, 2025-02=50180, 2025-03=51279, 2025-04=56278, 2025-05=51616, 2025-06=50212, 2025-07=59817, 2025-08=54406, 2025-09=56875, 2025-10=62065, 2025-11=53248, 2025-12=60208, 2026-01=50791, 2026-02=49722, 2026-03=59757, 2026-04=56425, 2026-05=52627, 2026-06=54152, 2026-07=60820, 2026-08=57133, 2026-09=61811, 2026-10=48906
- Riscos: dictionary_has_lob_or_long (3, not extracted): OBSNFE CLOB, IDINTEGRACAOMYFROTA RAW, OBSNFEFISCO CLOB; wide_table: 670 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSVENDA | NUMBER(10) | Int64 |  | key | fato_faturamento.num_transacao_venda |  |
| DTSAIDA | DATE | DateTime64(0) |  | watermark | fato_faturamento.data |  |
| CODFILIAL | VARCHAR2(2) | String |  | proposed | fato_faturamento.cod_filial |  |
| CODCLI | NUMBER(9) | Int64 |  | proposed | fato_faturamento.cod_cliente |  |
| CODUSUR | NUMBER(4) | Int64 |  | proposed | fato_faturamento.cod_rca |  |
| NUMNOTA | NUMBER(10) | Int64 | S | proposed | fato_faturamento.num_nota_saida |  |
| NUMPED | NUMBER(10) | Int64 | S | proposed | fato_faturamento.num_pedido_venda |  |
| TIPOVENDA | VARCHAR2(2) | String |  | scope | fato_faturamento.tipo_venda |  |
| CONDVENDA | NUMBER(5) | Int64 | S | scope | fato_faturamento.condicao_venda |  |
| ESPECIE | VARCHAR2(2) | String | S | scope | fato_faturamento.especie |  |
| SERIE | VARCHAR2(5) | String | S | scope | fato_faturamento.serie |  |
| CODFISCAL | NUMBER(8) | Int64 |  | control |  | CFOP da nota, para as exclusões herdadas; confirmar se a regra usa o CFOP do item (PCMOV) |
| DTCANCEL | DATE | DateTime64(0) | S | control |  | nota cancelada (exclusão herdada) |
| PERDESCFIN | NUMBER(12,6) | Decimal(12,6) | S | proposed | fato_faturamento.valor_desconto_financeiro | percentual de desconto financeiro da nota; o item tem PCMOV.PERCDESCFIN |
| VLDESCFIN | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor_desconto_financeiro | valor do desconto financeiro da nota |

### PCNFENT (onda 1) — Notas fiscais de entrada (devoluções e compras)

- Owner `RIBERDOCES`, 532 colunas no dicionário, 14 extraídas. Estratégia `window`, watermark `DTENT` (date), reprocessamento de 45 dias.
- Chave: NUMTRANSENT, CODCONT [constraint:PCNFENT_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 535056 (analisado em 2026-10-25), COUNT(*) 535056, tamanho estimado 1015 MB, linha média 959 bytes.
- Watermark: mínimo 2009-09-17, máximo 2026-10-24, nulos 0.
- Distribuição por mês (watermark): 2023-10=4172, 2023-11=4057, 2023-12=4189, 2024-01=4037, 2024-02=3564, 2024-03=3795, 2024-04=4011, 2024-05=4077, 2024-06=3813, 2024-07=4292, 2024-08=4168, 2024-09=3990, 2024-10=4540, 2024-11=3618, 2024-12=3951, 2025-01=3629, 2025-02=3565, 2025-03=3521, 2025-04=3668, 2025-05=3553, 2025-06=3445, 2025-07=4268, 2025-08=3978, 2025-09=4302, 2025-10=4308, 2025-11=3755, 2025-12=4541, 2026-01=3916, 2026-02=3801, 2026-03=4425, 2026-04=4163, 2026-05=4110, 2026-06=3995, 2026-07=4389, 2026-08=4384, 2026-09=4575, 2026-10=3291
- Riscos: dictionary_has_lob_or_long (2, not extracted): OBSNFE CLOB, OBSNFEFISCO CLOB; wide_table: 532 columns in dictionary
- Pergunta: A chave primária do dicionário (PCNFENT_PK: NUMTRANSENT, CODCONT) difere da esperada pelo escopo (NUMTRANSENT). O inventário adota a do dicionário.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSENT | NUMBER(10) | Int64 |  | key | fato_devolucao.num_transacao_entrada |  |
| DTENT | DATE | DateTime64(0) |  | watermark | fato_devolucao.data |  |
| CODFILIAL | VARCHAR2(2) | String |  | proposed | fato_compra.cod_filial |  |
| CODFORNEC | NUMBER(9) | Int64 |  | proposed | fato_compra.cod_fornecedor |  |
| NUMNOTA | NUMBER(10) | Int64 |  | proposed | fato_devolucao.num_nota_devolucao |  |
| ESPECIE | VARCHAR2(2) | String |  | scope | fato_compra.especie |  |
| SERIE | VARCHAR2(3) | String | S | scope | fato_compra.serie |  |
| TIPODESCARGA | VARCHAR2(1) | String | S | scope | fato_compra.tipo_entrada |  |
| CODDEVOL | NUMBER(4) | Int64 | S | scope | fato_devolucao.cod_motivo_devolucao |  |
| CODUSURDEVOL | NUMBER(4) | Int64 | S | scope | fato_devolucao.cod_rca |  |
| NUMTRANSVENDAORIG | NUMBER(10) | Int64 | S | proposed | fato_devolucao.num_transacao_venda | venda de origem da devolução vinculada; o cliente vem de PCMOV.CODCLI ou da nota de origem |
| NUMPEDPRINC | NUMBER(8) | Int64 | S | proposed | fato_compra.num_pedido_compra | candidato a pedido de compra de origem (ponto 13); PCNFENT não tem NUMPED |
| DTCANCEL | DATE | DateTime64(0) | S | control |  | nota cancelada |
| CODCONT | NUMBER(10) | Int64 |  | key |  | faz parte da chave primária; candidato à regra de contabilização de fato_compra |

### PCESTCOM (onda 1) — Citada no escopo (onda 1) sem colunas; função a confirmar

- Owner `RIBERDOCES`, 18 colunas no dicionário, 1 extraídas. Estratégia `full`.
- Chave: NUMTRANSENT [constraint:PCESTCOM_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 137757 (analisado em 2026-10-25), COUNT(*) 137757, tamanho estimado 10 MB, linha média 58 bytes.
- Pergunta: PCESTCOM está na lista de tabelas da onda 1, mas nenhum mart cita suas colunas. Confirmar o uso (provável: dados complementares da nota de entrada) antes de extrair.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSENT | NUMBER(10) | Int64 |  | key |  | added by inventory: required for key |

### PCMOV (onda 1) — Itens de nota (saída, devolução e compra) e movimentos de estoque

- Owner `RIBERDOCES`, 524 colunas no dicionário, 33 extraídas. Estratégia `window`, watermark `DTMOV` (date), reprocessamento de 45 dias.
- Chave: NUMTRANSITEM [constraint:PCMOV_UK]; unicidade verificada: true (0 grupos duplicados, 644975 linhas com chave nula).
- Linhas: dicionário 64017288 (analisado em 2026-10-25), COUNT(*) 64017288, tamanho estimado 68860 MB, linha média 775 bytes.
- Watermark: mínimo 1899-12-30, máximo 2026-10-24, nulos 0.
- Distribuição por mês (watermark): 2023-10=640554, 2023-11=594935, 2023-12=616367, 2024-01=588814, 2024-02=559665, 2024-03=592099, 2024-04=593239, 2024-05=616082, 2024-06=558468, 2024-07=660884, 2024-08=645427, 2024-09=639641, 2024-10=697595, 2024-11=610079, 2024-12=674183, 2025-01=629664, 2025-02=602518, 2025-03=603890, 2025-04=647012, 2025-05=591229, 2025-06=561481, 2025-07=693883, 2025-08=635959, 2025-09=685479, 2025-10=739159, 2025-11=627508, 2025-12=739176, 2026-01=609940, 2026-02=583482, 2026-03=693201, 2026-04=660190, 2026-05=617500, 2026-06=615695, 2026-07=726853, 2026-08=688214, 2026-09=733477, 2026-10=567334
- Riscos: nullable_key_column: NUMTRANSITEM; dictionary_has_lob_or_long (3, not extracted): RETORNO_SERVICO_CBS CLOB, RETORNO_SERVICO_IBS CLOB, RETORNO_SERVICO_IS CLOB; wide_table: 524 columns in dictionary; key_has_nulls: 644975 rows with NULL in a key column

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSITEM | NUMBER(18) | Int64 | S | key | fato_faturamento.num_transacao_item |  |
| DTMOV | DATE | DateTime64(0) | S | watermark |  | data do movimento; a data do fato vem de PCNFSAID.DTSAIDA / PCNFENT.DTENT |
| NUMTRANSVENDA | NUMBER(10) | Int64 | S | proposed | fato_faturamento.num_transacao_venda |  |
| NUMTRANSENT | NUMBER(10) | Int64 | S | proposed | fato_compra.num_transacao_entrada |  |
| CODOPER | VARCHAR2(2) | String |  | control |  | operação (S saída, ED devolução, E/EB compra) |
| CODFILIAL | VARCHAR2(2) | String |  | scope | fato_faturamento.cod_filial |  |
| CODCLI | NUMBER(9) | Int64 | S | scope | fato_faturamento.cod_cliente |  |
| CODUSUR | NUMBER(8) | Int64 | S | scope | fato_faturamento.cod_rca |  |
| CODPROD | NUMBER(6) | Int64 |  | scope | fato_faturamento.cod_produto |  |
| CODFORNEC | NUMBER(9) | Int64 | S | proposed | fato_compra.cod_fornecedor |  |
| NUMNOTA | NUMBER(10) | Int64 | S | proposed | fato_faturamento.num_nota_saida |  |
| NUMPED | NUMBER(10) | Int64 | S | proposed | fato_faturamento.num_pedido_venda |  |
| NUMSEQ | NUMBER(20) | Decimal(20,0) | S | proposed |  | sequência do item |
| QT | NUMBER(20,6) | Decimal(20,6) | S | proposed | fato_faturamento.qtd_itens |  |
| PUNIT | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_compra.valor |  |
| PTABELA | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor_tabela |  |
| PUNITCONT | NUMBER(18,6) | Decimal(18,6) | S | proposed |  | preço unitário contábil; candidato ao valor faturado |
| ST | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor_st |  |
| VLFRETE | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor | parcela de frete do valor faturado |
| VLOUTRASDESP | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor | outras despesas |
| VLIPI | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor | IPI |
| VLDESCONTO | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_compra.valor_desconto |  |
| CUSTOFIN | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_faturamento.cmv |  |
| CODFISCAL | NUMBER(8) | Int64 | S | control |  | CFOP do item (exclusões herdadas) |
| TIPOITEM | VARCHAR2(1) | String | S | control |  | TIPOITEM fora de C e N é exclusão herdada |
| PERCDESCFIN | NUMBER(12,4) | Decimal(12,4) | S | scope | fato_faturamento.valor_desconto_financeiro |  |
| VLDESCFIN | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor_desconto_financeiro | valor do desconto financeiro do item |
| VLBONIFIC | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.valor_bonificado | bonificação no item; a condição de venda fica em PCNFSAID.CONDVENDA |
| STATUS | VARCHAR2(2) | String |  | control |  | item cancelado |
| PESOBRUTO | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_faturamento.peso_bruto | se não existir, vem de PCPRODUT |
| PESOLIQ | NUMBER(12,6) | Decimal(12,6) | S | proposed | fato_faturamento.peso_liquido | se não existir, vem de PCPRODUT |
| VOLUME | NUMBER(20,8) | Decimal(20,8) | S | proposed | fato_faturamento.volume | se não existir, vem de PCPRODUT |
| QTUNITCX | NUMBER(8,2) | Decimal(8,2) | S | proposed | fato_faturamento.qtd_caixas |  |

### PCMOVCOMPLE (onda 1) — Complemento do item de movimento

- Owner `RIBERDOCES`, 587 colunas no dicionário, 1 extraídas. Estratégia `window`.
- Chave: NUMTRANSITEM [constraint:PCMOVCOMPLE_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 63250854 (analisado em 2026-10-25), COUNT(*) 63250854, tamanho estimado 45455 MB, linha média 559 bytes.
- Riscos: incremental_without_watermark; wide_table: 587 columns in dictionary
- Pergunta: Nenhum mart cita colunas de PCMOVCOMPLE (TIPOITEM está em PCMOV). Confirmar se a tabela ainda é necessária.
- Pergunta: PCMOVCOMPLE não tem data própria: a janela precisa vir de PCMOV (mesma NUMTRANSITEM). Proposta: extrair por faixa de NUMTRANSITEM derivada da janela de PCMOV.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSITEM | NUMBER(18) | Int64 |  | key | fato_faturamento.num_transacao_item |  |

### PCPEDC (onda 1) — Pedidos de venda (cabeçalho)

- Owner `RIBERDOCES`, 524 colunas no dicionário, 11 extraídas. Estratégia `window`, watermark `DATA` (date), reprocessamento de 45 dias.
- Chave: NUMPED [constraint:PCPEDC_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 6917841 (analisado em 2026-10-25), COUNT(*) 6917849, tamanho estimado 5819 MB, linha média 547 bytes.
- Watermark: mínimo 2009-10-02, máximo 2026-10-25, nulos 0.
- Distribuição por mês (watermark): 2023-10=53216, 2023-11=49089, 2023-12=49189, 2024-01=47543, 2024-02=45939, 2024-03=49263, 2024-04=51201, 2024-05=50240, 2024-06=48636, 2024-07=56172, 2024-08=52085, 2024-09=53439, 2024-10=56966, 2024-11=49745, 2024-12=54906, 2025-01=50042, 2025-02=49160, 2025-03=51461, 2025-04=54934, 2025-05=49620, 2025-06=50408, 2025-07=58172, 2025-08=52648, 2025-09=57987, 2025-10=58478, 2025-11=52223, 2025-12=58923, 2026-01=49716, 2026-02=48663, 2026-03=60197, 2026-04=53479, 2026-05=51623, 2026-06=54617, 2026-07=58002, 2026-08=57523, 2026-09=59338, 2026-10=45277
- Riscos: dictionary_has_lob_or_long (1, not extracted): JSONTPD_PAGAMENTO BLOB; wide_table: 524 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMPED | NUMBER(10) | Int64 |  | key | fato_venda.num_pedido_venda |  |
| DATA | DATE | DateTime64(0) |  | watermark | fato_venda.data |  |
| CODFILIAL | VARCHAR2(2) | String |  | proposed | fato_venda.cod_filial |  |
| CODCLI | NUMBER(9) | Int64 |  | proposed | fato_venda.cod_cliente |  |
| CODUSUR | NUMBER(4) | Int64 |  | proposed | fato_venda.cod_rca |  |
| POSICAO | VARCHAR2(2) | String |  | scope | fato_venda.situacao_pedido |  |
| ORIGEMPED | VARCHAR2(1) | String | S | scope | fato_venda.origem_pedido |  |
| CONDVENDA | NUMBER(5) | Int64 | S | scope | fato_venda.condicao_venda |  |
| NUMTRANSVENDA | NUMBER(10) | Int64 | S | proposed |  | liga o pedido à nota faturada |
| DTCANCEL | DATE | DateTime64(0) | S | control |  | pedido cancelado |
| TIPOVENDA | VARCHAR2(2) | String | S | control |  | regra de brinde, bonificação e transferência (DIM_PEDIDOVENDA) |

### PCPEDI (onda 1) — Itens de pedido de venda

- Owner `RIBERDOCES`, 476 colunas no dicionário, 15 extraídas. Estratégia `window`, watermark `DATA` (date), reprocessamento de 45 dias.
- Chave: NUMPED, CODPROD, NUMSEQ [constraint:PCPEDI_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 37589500 (analisado em 2026-10-25), COUNT(*) 37589594, tamanho estimado 21368 MB, linha média 406 bytes.
- Watermark: mínimo 2009-10-02, máximo 2026-10-25, nulos 0.
- Distribuição por mês (watermark): 2023-10=323670, 2023-11=300516, 2023-12=300598, 2024-01=292748, 2024-02=277931, 2024-03=291787, 2024-04=301442, 2024-05=300105, 2024-06=281244, 2024-07=337758, 2024-08=307452, 2024-09=330317, 2024-10=346474, 2024-11=298085, 2024-12=340136, 2025-01=311253, 2025-02=300740, 2025-03=312400, 2025-04=326762, 2025-05=290604, 2025-06=294824, 2025-07=346713, 2025-08=316980, 2025-09=360828, 2025-10=356536, 2025-11=317587, 2025-12=372885, 2026-01=301690, 2026-02=291482, 2026-03=362640, 2026-04=319967, 2026-05=308066, 2026-06=321081, 2026-07=351584, 2026-08=356242, 2026-09=362826, 2026-10=274595
- Riscos: dictionary_has_lob_or_long (4, not extracted): ORIGEMDEPRECO CLOB, RETORNO_SERVICO_CBS CLOB, RETORNO_SERVICO_IBS CLOB, RETORNO_SERVICO_IS CLOB; wide_table: 476 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMPED | NUMBER(10) | Int64 |  | key | fato_venda.num_pedido_venda |  |
| CODPROD | NUMBER(6) | Int64 |  | key | fato_venda.cod_produto |  |
| NUMSEQ | NUMBER(20) | Decimal(20,0) |  | key | fato_venda.num_seq_item |  |
| DATA | DATE | DateTime64(0) |  | watermark | fato_venda.data |  |
| CODCLI | NUMBER(9) | Int64 |  | proposed | fato_venda.cod_cliente |  |
| CODUSUR | NUMBER(4) | Int64 |  | proposed | fato_venda.cod_rca |  |
| QT | NUMBER(20,6) | Decimal(20,6) |  | proposed | fato_venda.qtd_itens |  |
| PVENDA | NUMBER(18,6) | Decimal(18,6) |  | proposed | fato_venda.valor |  |
| PTABELA | NUMBER(18,6) | Decimal(18,6) |  | proposed | fato_venda.valor_tabela |  |
| VLCUSTOFIN | NUMBER(18,6) | Decimal(18,6) |  | scope | fato_venda.cmv |  |
| CONDVENDA | NUMBER(5) | Int64 | S | control |  | condição de venda do item (bonificado) |
| POSICAO | VARCHAR2(2) | String |  | control |  | posição do item |
| BONIFIC | VARCHAR2(1) | String | S | control |  | item bonificado (S/N) |
| PESOBRUTO | NUMBER(7,3) | Decimal(7,3) | S | proposed | fato_venda.peso_bruto |  |
| VLBONIFIC | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_venda.valor_bonificado |  |

### PCMETA (onda 1) — Metas por tipo (ponto 3 do escopo)

- Owner `RIBERDOCES`, 53 colunas no dicionário, 16 extraídas. Estratégia `full`.
- Chave: CODFILIAL, CODIGO, CODIGO2, CODUSUR, TIPOMETA, DATA [constraint:PCMETA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 3385 (analisado em 2026-10-25), COUNT(*) 3385, tamanho estimado 0 MB, linha média 72 bytes.
- Pergunta: Ponto 3: o tipo M de PCMETA é meta de RCA? PCMETARCA é outra rotina (353). Escolher a fonte oficial com o negócio.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| DATA | DATE | DateTime64(0) |  | key | fato_meta.data |  |
| TIPOMETA | VARCHAR2(2) | String |  | key | fato_meta.tipo_meta | o escopo chama de TIPO |
| CODIGO | NUMBER(8) | Int64 |  | key | fato_meta.cod_chave |  |
| CODIGO2 | NUMBER(8) | Int64 |  | key | fato_meta.cod_chave2 |  |
| CODFILIAL | VARCHAR2(2) | String |  | key | fato_meta.cod_filial |  |
| CODUSUR | NUMBER(4) | Int64 |  | key | fato_meta.cod_rca |  |
| VLVENDAPREV | NUMBER(12,2) | Decimal(12,2) | S | scope | fato_meta.valor_meta |  |
| QTVENDAPREV | NUMBER(20,6) | Decimal(20,6) | S | scope | fato_meta.qtd_venda_meta |  |
| PEDIDOSPREV | NUMBER(8,4) | Decimal(8,4) | S | scope | fato_meta.qtd_pedidos_meta |  |
| CLIPOSPREV | NUMBER(10,4) | Decimal(10,4) | S | scope | fato_meta.positivacao_meta |  |
| MIXPREV | NUMBER(10,4) | Decimal(10,4) | S | scope | fato_meta.mix_meta |  |
| QTPESOPREV | NUMBER(12,2) | Decimal(12,2) | S | scope | fato_meta.peso_meta |  |
| VOLUMEPREV | NUMBER(20,8) | Decimal(20,8) | S | scope | fato_meta.volume_meta |  |
| MARGEMPREV | NUMBER(8,4) | Decimal(8,4) | S | proposed | fato_meta.margem_meta_pct |  |
| PERINADIMPPREV | NUMBER(8,4) | Decimal(8,4) | S | proposed | fato_meta.inadimplencia_meta_pct |  |
| PERVLDEVOLPREV | NUMBER(8,4) | Decimal(8,4) | S | proposed | fato_meta.devolucao_meta_pct | percentual de devolução em valor; existe também PERQTDEVOLPREV (quantidade) |

### PCMETARCA (onda 1) — Metas de RCA (rotina 353)

- Owner `RIBERDOCES`, 13 colunas no dicionário, 8 extraídas. Estratégia `full`.
- Chave: CODFILIAL, CODUSUR, DATA [constraint:PCMETARCA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 296975 (analisado em 2026-10-25), COUNT(*) 296975, tamanho estimado 13 MB, linha média 30 bytes.
- Pergunta: PCMETARCA não tem QTVENDAPREV, PEDIDOSPREV, CLIPOSPREV, MIXPREV, QTPESOPREV nem VOLUMEPREV. As colunas propostas (QTPEDPREV, NUMCLIPOS, QTITENSPEDPREV, QTMETROPREV) precisam de confirmação; peso não tem candidata.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODFILIAL | VARCHAR2(2) | String |  | key | fato_meta.cod_filial |  |
| CODUSUR | NUMBER(4) | Int64 |  | key | fato_meta.cod_rca |  |
| DATA | DATE | DateTime64(0) |  | key | fato_meta.data |  |
| VLVENDAPREV | NUMBER(14,2) | Decimal(14,2) | S | scope | fato_meta.valor_meta |  |
| QTPEDPREV | NUMBER(14,2) | Decimal(14,2) | S | proposed | fato_meta.qtd_pedidos_meta | o escopo cita PEDIDOSPREV, que não existe |
| NUMCLIPOS | NUMBER(20,8) | Decimal(20,8) | S | proposed | fato_meta.positivacao_meta | o escopo cita CLIPOSPREV, que não existe |
| QTITENSPEDPREV | NUMBER(14,2) | Decimal(14,2) | S | proposed | fato_meta.mix_meta | candidato a mix (itens por pedido); o escopo cita MIXPREV, que não existe |
| QTMETROPREV | NUMBER(20,8) | Decimal(20,8) | S | proposed | fato_meta.volume_meta | candidato a volume; o escopo cita VOLUMEPREV, que não existe |

### PCPREST (onda 2) — Contas a receber (títulos)

- Owner `RIBERDOCES`, 309 colunas no dicionário, 16 extraídas. Estratégia `open_items`, watermark `DTEMISSAO` (date), reprocessamento de 45 dias, em aberto quando `DTPAG`, `DTBAIXA` nulo.
- Chave: NUMTRANSVENDA, PREST [constraint:PCPREST_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 10792794 (analisado em 2026-10-25), COUNT(*) 10792794, tamanho estimado 8697 MB, linha média 693 bytes.
- Watermark: mínimo 2009-10-02, máximo 2026-10-24, nulos 0.
- Distribuição por mês (watermark): 2023-10=82817, 2023-11=77928, 2023-12=80902, 2024-01=72860, 2024-02=73271, 2024-03=81654, 2024-04=78064, 2024-05=80607, 2024-06=76657, 2024-07=86639, 2024-08=83781, 2024-09=82817, 2024-10=91164, 2024-11=79447, 2024-12=87743, 2025-01=79016, 2025-02=77606, 2025-03=79713, 2025-04=87846, 2025-05=79433, 2025-06=76980, 2025-07=92518, 2025-08=84513, 2025-09=88518, 2025-10=96122, 2025-11=82340, 2025-12=96028, 2026-01=79913, 2026-02=78288, 2026-03=96510, 2026-04=87809, 2026-05=82565, 2026-06=84203, 2026-07=95206, 2026-08=90823, 2026-09=95749, 2026-10=73153
- Riscos: wide_table: 309 columns in dictionary

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMTRANSVENDA | NUMBER(10) | Int64 |  | key | fato_contas_receber.num_transacao_venda |  |
| PREST | VARCHAR2(2) | String |  | key | fato_contas_receber.num_prestacao |  |
| DUPLIC | NUMBER(10) | Int64 |  | proposed | fato_contas_receber.num_duplicata |  |
| DTEMISSAO | DATE | DateTime64(0) |  | watermark | fato_contas_receber.data_emissao |  |
| DTULTALTER | DATE | DateTime64(0) | S | proposed |  | data de alteração; candidata a watermark secundário (ponto Q9) |
| CODCLI | NUMBER(9) | Int64 |  | proposed | fato_contas_receber.cod_cliente |  |
| CODUSUR | NUMBER(4) | Int64 |  | proposed | fato_contas_receber.cod_rca |  |
| CODFILIAL | VARCHAR2(2) | String |  | proposed | fato_contas_receber.cod_filial |  |
| CODCOB | VARCHAR2(4) | String |  | scope | fato_contas_receber.cod_cobranca |  |
| DTVENC | DATE | DateTime64(0) |  | scope | fato_contas_receber.data_vencimento |  |
| DTPAG | DATE | DateTime64(0) | S | scope | fato_contas_receber.data_pagamento |  |
| DTBAIXA | DATE | DateTime64(0) | S | scope | fato_contas_receber.data_baixa |  |
| VALOR | NUMBER(10,2) | Decimal(10,2) |  | scope | fato_contas_receber.valor_original |  |
| VPAGO | NUMBER(10,2) | Decimal(10,2) | S | scope | fato_contas_receber.valor_recebido |  |
| TXPERM | NUMBER(10,2) | Decimal(10,2) | S | scope | fato_contas_receber.valor_juros |  |
| VALORDESC | NUMBER(10,2) | Decimal(10,2) | S | scope | fato_contas_receber.valor_desconto |  |

### PCEST (onda 2) — Posição atual de estoque por filial e produto

- Owner `RIBERDOCES`, 193 colunas no dicionário, 16 extraídas. Estratégia `full`.
- Chave: CODFILIAL, CODPROD [constraint:PCEST_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 195721 (analisado em 2026-10-25), COUNT(*) 195721, tamanho estimado 63 MB, linha média 292 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODPROD | NUMBER(6) | Int64 |  | key | fato_estoque_snapshot.cod_produto |  |
| CODFILIAL | VARCHAR2(2) | String |  | key | fato_estoque_snapshot.cod_filial |  |
| QTESTGER | NUMBER(22,8) | Decimal(22,8) | S | scope | fato_estoque_snapshot.qtd_gerencial |  |
| QTRESERV | NUMBER(22,8) | Decimal(22,8) | S | scope | fato_estoque_snapshot.qtd_reservada |  |
| QTBLOQUEADA | NUMBER(20,6) | Decimal(20,6) | S | scope | fato_estoque_snapshot.qtd_bloqueada |  |
| QTPENDENTE | NUMBER(16,3) | Decimal(16,3) | S | scope | fato_estoque_snapshot.qtd_pendente |  |
| QTINDENIZ | NUMBER(20,6) | Decimal(20,6) | S | scope | fato_estoque_snapshot.qtd_avaria |  |
| CUSTOREAL | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.custo_real_unit |  |
| CUSTOFIN | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.custo_fin_unit |  |
| CUSTOCONT | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.custo_cont_unit |  |
| CUSTOREP | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.custo_rep_unit |  |
| CUSTOULTENT | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.custo_ult_entrada_unit |  |
| VALORULTENT | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_estoque_snapshot.valor_ultima_entrada_unit |  |
| QTGIRODIA | NUMBER(16,3) | Decimal(16,3) | S | scope | fato_estoque_snapshot.giro_dia |  |
| DTULTENT | DATE | DateTime64(0) | S | scope | fato_estoque_snapshot.data_ultima_entrada |  |
| DTULTSAIDA | DATE | DateTime64(0) | S | scope | fato_estoque_snapshot.data_ultima_saida |  |

### PCTABPR (onda 2) — Tabela de preços por região

- Owner `RIBERDOCES`, 185 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: NUMREGIAO, CODPROD [constraint:PCTABPR_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 424644 (analisado em 2026-10-25), COUNT(*) 424644, tamanho estimado 223 MB, linha média 437 bytes.
- Riscos: dictionary_has_lob_or_long (7, not extracted): FORMULA CLOB, RETORNO_SERVICO_CBS CLOB, RETORNO_SERVICO_IBS CLOB, RETORNO_SERVICO_IS CLOB, RETORNO_SERVICO_CBS_TAB CLOB, RETORNO_SERVICO_IBS_TAB CLOB, RETORNO_SERVICO_IS_TAB CLOB

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMREGIAO | NUMBER(4) | Int64 |  | key |  |  |
| CODPROD | NUMBER(6) | Int64 |  | key |  |  |
| PVENDA | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_estoque_snapshot.preco_venda_unit | preço de tabela da região padrão |

### PCPARAMFILIAL (onda 2) — Parâmetros por filial (preço por embalagem)

- Owner `RIBERDOCES`, 3 colunas no dicionário, 3 extraídas. Estratégia `full`.
- Chave: NOME, CODFILIAL [constraint:PCPARAMFILIAL_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 16490 (analisado em 2026-10-25), COUNT(*) 16490, tamanho estimado 1 MB, linha média 26 bytes.
- Pergunta: Confirmar quais parâmetros de PCPARAMFILIAL o escopo usa (preço por embalagem e, talvez, BLOQUEIAVENDAESTPENDENTE).

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODFILIAL | VARCHAR2(2) | String |  | key |  |  |
| NOME | VARCHAR2(34) | String |  | key |  | nome do parâmetro |
| VALOR | VARCHAR2(1000) | String | S | proposed |  | valor do parâmetro; confirmar nome |

### PCEMBALAGEM (onda 2) — Embalagens e preço por embalagem

- Owner `RIBERDOCES`, 149 colunas no dicionário, 6 extraídas. Estratégia `full`.
- Chave: CODFILIAL, CODAUXILIAR [constraint:PCEMBALAGEM_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 257800 (analisado em 2026-10-25), COUNT(*) 257800, tamanho estimado 71 MB, linha média 197 bytes.
- Riscos: dictionary_has_lob_or_long (1, not extracted): INFOPRODWEB CLOB

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODFILIAL | VARCHAR2(2) | String |  | key |  |  |
| CODAUXILIAR | NUMBER(20) | Decimal(20,0) |  | key |  |  |
| CODPROD | NUMBER(6) | Int64 |  | scope | fato_estoque_snapshot.cod_produto |  |
| PVENDA | NUMBER(12,3) | Decimal(12,3) | S | proposed | fato_estoque_snapshot.preco_venda_unit |  |
| QTUNIT | NUMBER(18,6) | Decimal(18,6) | S | proposed |  | unidades por embalagem |
| DTINATIVO | DATE | DateTime64(0) | S | control |  | embalagem inativa |

### PCCORTEI (onda 2) — Itens cortados na separação

- Owner `RIBERDOCES`, 25 colunas no dicionário, 14 extraídas. Estratégia `window`, watermark `DATA` (date), reprocessamento de 45 dias.
- Chave: CODPROD, NUMCAR, NUMPED, NUMSEQ [constraint:PCCORTEI_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 207520 (analisado em 2026-10-25), COUNT(*) 207520, tamanho estimado 20 MB, linha média 80 bytes.
- Watermark: mínimo 2009-09-16, máximo 2026-10-23, nulos 0.
- Distribuição por mês (watermark): 2023-10=401, 2023-11=433, 2023-12=1079, 2024-01=154, 2024-02=181, 2024-03=495, 2024-04=186, 2024-05=402, 2024-06=170, 2024-07=240, 2024-08=259, 2024-09=302, 2024-10=477, 2024-11=352, 2024-12=962, 2025-01=611, 2025-02=305, 2025-03=478, 2025-04=625, 2025-05=789, 2025-06=540, 2025-07=973, 2025-08=541, 2025-09=813, 2025-10=681, 2025-11=661, 2025-12=1213, 2026-01=601, 2026-02=506, 2026-03=937, 2026-04=895, 2026-05=688, 2026-06=705, 2026-07=1190, 2026-08=895, 2026-09=604, 2026-10=531
- Pergunta: O grão de fato_corte no escopo é (pedido, produto, data, motivo); a chave real inclui NUMCAR e NUMSEQ. Ajustar o grão do mart na fase 2.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| DATA | DATE | DateTime64(0) |  | watermark | fato_corte.data |  |
| NUMPED | NUMBER(10) | Int64 |  | key | fato_corte.num_pedido_venda |  |
| CODPROD | NUMBER(6) | Int64 |  | key | fato_corte.cod_produto |  |
| NUMCAR | NUMBER(8) | Int64 |  | key |  | carregamento; parte da chave primária |
| NUMSEQ | NUMBER(20) | Decimal(20,0) |  | key |  | parte da chave primária |
| MOTIVO | VARCHAR2(80) | String | S | proposed | fato_corte.motivo_corte |  |
| TIPOCORTE | VARCHAR2(1) | String | S | control |  | tipo do corte |
| CODCLI | NUMBER(9) | Int64 | S | proposed | fato_corte.cod_cliente | permite a análise de clientes afetados sem join com PCPEDC |
| CODFILIAL | VARCHAR2(2) | String | S | proposed | fato_corte.cod_filial |  |
| CODUSUR | NUMBER(4) | Int64 | S | proposed | fato_corte.cod_rca |  |
| CODFUNC | NUMBER(8) | Int64 |  | proposed | fato_corte.cod_funcionario |  |
| QTCORTADA | NUMBER(8,2) | Decimal(8,2) |  | scope | fato_corte.qtd_cortada |  |
| QTSEPARADA | NUMBER(8,2) | Decimal(8,2) |  | scope | fato_corte.qtd_separada |  |
| PVENDA | NUMBER(12,3) | Decimal(12,3) | S | scope | fato_corte.valor_cortado |  |

### PCFALTA (onda 2) — Itens em falta (pequena e sem índice na data: full refresh)

- Owner `RIBERDOCES`, 9 colunas no dicionário, 9 extraídas. Estratégia `full`, watermark `DATA` (date).
- Chave: NUMPED, CODPROD, NUMSEQ [constraint:PCFALTA1_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 30967 (analisado em 2026-10-25), COUNT(*) 30967, tamanho estimado 2 MB, linha média 37 bytes.
- Watermark: mínimo 2009-10-13, máximo 2026-10-14, nulos 0.
- Distribuição por mês (watermark): 2023-10=3, 2023-11=4, 2023-12=93, 2024-01=27, 2024-02=20, 2024-03=3, 2024-04=3, 2024-05=5, 2024-06=5, 2024-07=78, 2024-08=13, 2024-09=8, 2024-10=23, 2024-11=15, 2024-12=78, 2025-01=35, 2025-02=7, 2025-03=5, 2025-04=7, 2025-05=4, 2025-06=7, 2025-07=35, 2025-08=29, 2025-09=45, 2025-10=38, 2025-11=45, 2025-12=151, 2026-01=56, 2026-02=4, 2026-03=40, 2026-04=35, 2026-05=6, 2026-06=24, 2026-07=18, 2026-08=18, 2026-09=8, 2026-10=4
- Riscos: watermark_not_indexed: DATA
- Pergunta: DATA de PCFALTA não tem índice e a tabela tem ~31 mil linhas: full refresh. O watermark fica registrado só para a distribuição por mês.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| DATA | DATE | DateTime64(0) |  | watermark | fato_falta.data |  |
| NUMPED | NUMBER(10) | Int64 |  | key | fato_falta.num_pedido_venda |  |
| CODPROD | NUMBER(6) | Int64 |  | key | fato_falta.cod_produto |  |
| NUMSEQ | NUMBER(20) | Decimal(20,0) |  | key |  | parte da chave primária |
| CODFILIAL | VARCHAR2(2) | String | S | proposed | fato_falta.cod_filial |  |
| CODCLI | NUMBER(9) | Int64 | S | proposed | fato_falta.cod_cliente |  |
| CODUSUR | NUMBER(4) | Int64 | S | proposed | fato_falta.cod_rca |  |
| QT | NUMBER(8,2) | Decimal(8,2) | S | scope | fato_falta.qtd_em_falta |  |
| PVENDA | NUMBER(12,3) | Decimal(12,3) | S | scope | fato_falta.valor_falta |  |

### PCPEDIDO (onda 3) — Pedidos de compra (cabeçalho)

- Owner `RIBERDOCES`, 242 colunas no dicionário, 6 extraídas. Estratégia `window`, watermark `DTEMISSAO` (date), reprocessamento de 90 dias.
- Chave: NUMPED [constraint:PCPEDIDO_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 134393 (analisado em 2026-10-25), COUNT(*) 134399, tamanho estimado 58 MB, linha média 363 bytes.
- Watermark: mínimo 2009-09-26, máximo 2026-10-25, nulos 0.
- Distribuição por mês (watermark): 2023-10=604, 2023-11=569, 2023-12=501, 2024-01=547, 2024-02=493, 2024-03=494, 2024-04=578, 2024-05=555, 2024-06=523, 2024-07=591, 2024-08=615, 2024-09=488, 2024-10=597, 2024-11=499, 2024-12=513, 2025-01=555, 2025-02=501, 2025-03=515, 2025-04=530, 2025-05=504, 2025-06=507, 2025-07=678, 2025-08=595, 2025-09=576, 2025-10=566, 2025-11=593, 2025-12=585, 2026-01=573, 2026-02=525, 2026-03=671, 2026-04=574, 2026-05=546, 2026-06=640, 2026-07=618, 2026-08=650, 2026-09=652, 2026-10=498
- Riscos: dictionary_has_lob_or_long (1, not extracted): OCORRENCIAS CLOB; wide_table: 242 columns in dictionary
- Pergunta: PCPEDIDO não tem POSICAO; a situação do pedido vem das quantidades pendentes de PCITEM, como no escopo.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMPED | NUMBER(6) | Int64 |  | key | fato_pedido_compra.num_pedido_compra |  |
| DTEMISSAO | DATE | DateTime64(0) | S | watermark | fato_pedido_compra.data_emissao |  |
| DTPREVENT | DATE | DateTime64(0) | S | proposed | fato_pedido_compra.data_entrega_prevista | o escopo cita DATAENTREGA, que no cabeçalho não existe (existe em PCITEM) |
| CODFILIAL | VARCHAR2(2) | String | S | proposed | fato_pedido_compra.cod_filial |  |
| CODFORNEC | NUMBER(6) | Int64 |  | proposed | fato_pedido_compra.cod_fornecedor |  |
| CODCOMPRADOR | NUMBER(8) | Int64 | S | proposed |  | comprador do pedido |

### PCITEM (onda 3) — Itens de pedido de compra

- Owner `RIBERDOCES`, 359 colunas no dicionário, 10 extraídas. Estratégia `window`.
- Chave: NUMPED, CODPROD, NUMSEQ [constraint:PCITEM_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 1375271 (analisado em 2026-10-25), COUNT(*) 1375283, tamanho estimado 749 MB, linha média 473 bytes.
- Riscos: incremental_without_watermark; wide_table: 359 columns in dictionary
- Pergunta: PCITEM pode não ter data própria: a janela viria de PCPEDIDO (mesmo NUMPED). Alternativa: full refresh se o volume for pequeno.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| NUMPED | NUMBER(10) | Int64 |  | key | fato_pedido_compra.num_pedido_compra |  |
| CODPROD | NUMBER(6) | Int64 |  | key | fato_pedido_compra.cod_produto |  |
| NUMSEQ | NUMBER(6) | Int64 |  | key | fato_pedido_compra.num_seq |  |
| QTPEDIDA | NUMBER(20,6) | Decimal(20,6) | S | proposed | fato_pedido_compra.qtd_pedida |  |
| QTENTREGUE | NUMBER(20,6) | Decimal(20,6) | S | proposed | fato_pedido_compra.qtd_entregue |  |
| PCOMPRA | NUMBER(18,6) | Decimal(18,6) | S | proposed | fato_pedido_compra.valor_pedido | preço da mercadoria; nome a confirmar |
| PERCDESC | NUMBER(12,4) | Decimal(12,4) | S | proposed | fato_pedido_compra.valor_desconto | percentual de desconto do item; não há valor de desconto |
| QTFATORCONVERSAO | NUMBER(6,2) | Decimal(6,2) | S | proposed |  | candidato a fator de embalagem |
| DATAENTREGA | DATE | DateTime64(0) | S | scope | fato_pedido_compra.data_entrega_prevista | data de entrega prevista do item |
| DTULTENT | DATE | DateTime64(0) | S | proposed |  | última entrada do item; candidato ao prazo realizado de entrega |

### PCLANC (onda 3) — Contas a pagar (lançamentos)

- Owner `RIBERDOCES`, 203 colunas no dicionário, 21 extraídas. Estratégia `open_items`, watermark `DTLANC` (date), reprocessamento de 45 dias, em aberto quando `DTPAGTO` nulo.
- Chave: RECNUM [constraint:PCLANC_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 3308388 (analisado em 2026-10-25), COUNT(*) 3308391, tamanho estimado 1561 MB, linha média 459 bytes.
- Watermark: mínimo 2009-09-26, máximo 2026-10-25, nulos 901.
- Distribuição por mês (watermark): 2023-10=27836, 2023-11=25760, 2023-12=26420, 2024-01=26848, 2024-02=23987, 2024-03=25797, 2024-04=29951, 2024-05=27786, 2024-06=26396, 2024-07=29736, 2024-08=28603, 2024-09=28120, 2024-10=37292, 2024-11=26991, 2024-12=28761, 2025-01=26908, 2025-02=24153, 2025-03=26504, 2025-04=29065, 2025-05=28228, 2025-06=26944, 2025-07=30504, 2025-08=28038, 2025-09=30818, 2025-10=31798, 2025-11=28353, 2025-12=33324, 2026-01=27329, 2026-02=24759, 2026-03=30193, 2026-04=31021, 2026-05=28876, 2026-06=31008, 2026-07=31501, 2026-08=29087, 2026-09=33504, 2026-10=22677
- Riscos: dictionary_has_lob_or_long (2, not extracted): JSONENVIOCP CLOB, JSONRETORNOCP CLOB; wide_table: 203 columns in dictionary; watermark_has_nulls: 901 rows

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| RECNUM | NUMBER(8) | Int64 |  | key | fato_contas_pagar.num_lancamento |  |
| RECNUMPRINC | NUMBER(8) | Int64 | S | scope | fato_contas_pagar.num_lancamento_principal |  |
| DTLANC | DATE | DateTime64(0) | S | watermark | fato_contas_pagar.data_lancamento |  |
| DTULTALTER | DATE | DateTime64(0) | S | proposed |  | data de alteração; candidata a watermark secundário (ponto Q9) |
| DTEMISSAO | DATE | DateTime64(0) | S | proposed | fato_contas_pagar.data_emissao |  |
| DTVENC | DATE | DateTime64(0) | S | proposed | fato_contas_pagar.data_vencimento |  |
| DTPAGTO | DATE | DateTime64(0) | S | scope | fato_contas_pagar.data_pagamento |  |
| PREST | VARCHAR2(2) | String | S | scope | fato_contas_pagar.num_prestacao |  |
| DUPLIC | VARCHAR2(1) | String | S | scope | fato_contas_pagar.num_duplicata |  |
| TIPOPAGTO | VARCHAR2(3) | String | S | scope | fato_contas_pagar.tipo_pagamento |  |
| TIPOLANC | VARCHAR2(1) | String | S | proposed | fato_contas_pagar.tipo_lancamento | C confirmado / P provisionado; nome a confirmar |
| TIPOPARCEIRO | VARCHAR2(1) | String | S | scope | fato_contas_pagar.cod_parceiro |  |
| CODFORNEC | NUMBER(9) | Int64 |  | scope | fato_contas_pagar.cod_fornecedor |  |
| CODFILIAL | VARCHAR2(2) | String |  | proposed | fato_contas_pagar.cod_filial |  |
| CODCONTA | NUMBER(10) | Int64 |  | proposed | fato_contas_pagar.cod_conta |  |
| VALOR | NUMBER(12,2) | Decimal(12,2) |  | scope | fato_contas_pagar.valor_original |  |
| VPAGO | NUMBER(12,2) | Decimal(12,2) | S | scope | fato_contas_pagar.valor_pago |  |
| DESCONTOFIN | NUMBER(14,2) | Decimal(14,2) | S | scope | fato_contas_pagar.valor_desconto |  |
| TXPERM | NUMBER(14,2) | Decimal(14,2) | S | scope | fato_contas_pagar.valor_juros |  |
| VALORDEV | NUMBER(12,2) | Decimal(12,2) | S | scope | fato_contas_pagar.valor_devolucao |  |
| VLVARIACAOCAMBIAL | NUMBER(18,6) | Decimal(18,6) | S | scope | fato_contas_pagar.variacao_cambial |  |

### PCCONTA (onda 3) — Plano de contas

- Owner `RIBERDOCES`, 20 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODCONTA [constraint:PCCONTA_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 504 (analisado em 2026-10-25), COUNT(*) 504, tamanho estimado 0 MB, linha média 60 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODCONTA | NUMBER(10) | Int64 |  | key | fato_contas_pagar.cod_conta |  |
| CONTA | VARCHAR2(40) | String |  | proposed |  | descrição da conta |
| GRUPOCONTA | NUMBER(4) | Int64 |  | proposed | fato_contas_pagar.cod_grupo_conta |  |
| TIPO | VARCHAR2(1) | String | S | proposed |  | tipo da conta |

### PCLIB (onda 3) — Liberações (grupos de conta do contas a pagar)

- Owner `RIBERDOCES`, 6 colunas no dicionário, 4 extraídas. Estratégia `full`.
- Chave: CODFUNC, CODTABELA, CODIGOA, CODIGON [constraint:PCLIB_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 1252336 (analisado em 2026-10-25), COUNT(*) 1252336, tamanho estimado 45 MB, linha média 27 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODFUNC | NUMBER(8) | Int64 |  | key |  |  |
| CODTABELA | NUMBER(4) | Int64 |  | key |  |  |
| CODIGOA | VARCHAR2(40) | String |  | key |  | código alfanumérico liberado |
| CODIGON | NUMBER(10) | Int64 |  | key |  | código numérico liberado: grupo de conta quando CODFUNC = 1 e CODTABELA = 6 |

### PCRATEIOCENTROCUSTO (onda 3) — Rateio de lançamentos por centro de custo (ponto 11 do escopo)

- Owner `RIBERDOCES`, 20 colunas no dicionário, 7 extraídas. Estratégia `full`.
- Chave: RECNUM, CODIGOCENTROCUSTO [proposed]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 12050 (analisado em 2026-10-25), COUNT(*) 12050, tamanho estimado 2 MB, linha média 67 bytes.
- Riscos: no_unique_key_in_dictionary; nullable_key_column: CODIGOCENTROCUSTO
- Pergunta: PCRATEIOCENTROCUSTO não tem chave única no dicionário; CODCENTROCUSTO (numérico) é sempre nulo na base de teste e (RECNUM, CODIGOCENTROCUSTO) é único. Confirmar em outros clientes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| RECNUM | NUMBER(8) | Int64 |  | key | ponte_lancamento_centro_custo.num_lancamento |  |
| CODIGOCENTROCUSTO | VARCHAR2(40) | String | S | key | ponte_lancamento_centro_custo.cod_centro_custo | código texto do centro de custo, chave de PCCENTROCUSTO; (RECNUM, CODIGOCENTROCUSTO) é único na base de teste |
| CODCENTROCUSTO | NUMBER(10) | Int64 | S | proposed |  | código numérico; sempre nulo na base de teste |
| VALOR | NUMBER(12,2) | Decimal(12,2) |  | proposed | ponte_lancamento_centro_custo.valor_rateio | ponto 11: a tabela tem valor e percentual |
| PERCRATEIO | NUMBER(10,6) | Decimal(10,6) |  | proposed | ponte_lancamento_centro_custo.perc_rateio | ponto 11: a tabela tem valor e percentual |
| CODCONTA | NUMBER(10) | Int64 |  | proposed |  | conta do rateio |
| DTLANC | DATE | DateTime64(0) |  | proposed |  | data do lançamento rateado |

### PCCENTROCUSTO (onda 3) — Centros de custo

- Owner `RIBERDOCES`, 8 colunas no dicionário, 5 extraídas. Estratégia `full`.
- Chave: CODIGOCENTROCUSTO [constraint:PK_PCCENTROCUSTO]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 6 (analisado em 2026-10-25), COUNT(*) 6, tamanho estimado 0 MB, linha média 32 bytes.

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| CODIGOCENTROCUSTO | VARCHAR2(40) | String |  | key | dim_centro_custo.cod_centro_custo |  |
| CODCENTROCUSTO | VARCHAR2(40) | String | S | proposed |  | código alternativo (texto); liga ao rateio quando CODIGOCENTROCUSTO não bater |
| DESCRICAO | VARCHAR2(40) | String | S | proposed | dim_centro_custo.centro_custo |  |
| ATIVO | VARCHAR2(1) | String | S | proposed | dim_centro_custo.ativo |  |
| DTALTERACAO | DATE | DateTime64(0) | S | proposed |  | data de alteração; candidata a watermark secundário |

### PCEMPR (onda 3) — Funcionários (motoristas e colaboradores da dim_parceiro)

- Owner `RIBERDOCES`, 193 colunas no dicionário, 8 extraídas. Estratégia `full`.
- Chave: MATRICULA [constraint:PCEMPR_PK]; unicidade verificada: true (0 grupos duplicados, 0 linhas com chave nula).
- Linhas: dicionário 2404 (analisado em 2026-10-25), COUNT(*) 2404, tamanho estimado 2 MB, linha média 435 bytes.
- Riscos: dictionary_has_lob_or_long (6, not extracted): IDINTEGRACAOMYFROTA RAW, DIGITALPOLEGAR BLOB, DIGITALINDICADOR BLOB, DIGITALMEDIO BLOB, DIGITALANELAR BLOB, DIGITALMINIMO BLOB

| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |
|---|---|---|---|---|---|---|
| MATRICULA | NUMBER(8) | Int64 |  | key | dim_parceiro.cod_origem |  |
| NOME | VARCHAR2(40) | String |  | proposed | dim_parceiro.nome_parceiro |  |
| CODFILIAL | VARCHAR2(2) | String | S | proposed |  |  |
| SITUACAO | VARCHAR2(1) | String | S | control |  | ativo/inativo |
| TIPO | VARCHAR2(1) | String | S | proposed | dim_parceiro.tipo_parceiro |  |
| TIPOMOTORISTA | VARCHAR2(1) | String | S | proposed | dim_parceiro.tipo_parceiro | distingue motorista de colaborador |
| FUNCAO | VARCHAR2(30) | String | S | proposed |  | função do colaborador |
| DTULTALTER | DATE | DateTime64(0) | S | proposed |  | data de alteração; candidata a watermark secundário |

