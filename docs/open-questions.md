# Perguntas abertas

Dúvidas registradas durante a Fase 1. Cada item traz a decisão provisória adotada enquanto não há resposta.

## Infraestrutura e plataforma

| # | Pergunta | Provisório | Aberta em |
|---|---|---|---|
| Q1 | Nome definitivo do módulo Go e do repositório GitHub. | `github.com/the12company/mint` | 2026-10-25 |
| Q2 | Hospedagem do ClickHouse em staging/prod: o ClickHouse Cloud só oferece `sa-east-1` como private region (Enterprise). Ver ADR 0002. | Dev usa ClickHouse OSS em EC2 `sa-east-1`; prod decide antes da Fase 2 | 2026-10-25 |
| Q3 | Clientes com Windows Server 2012 R2: o gateway (Go 1.27) exige 2016+. Usar modo direto para eles? | Fora do gateway na v1 | 2026-10-25 |
| Q4 | Assinatura de código do `gateway.exe` (certificado) e empacotamento MSI. | Zip + `install.ps1` sem assinatura | 2026-10-25 |
| Q5 | ~~Owner das tabelas `PC*`~~ Resolvido no E1.2: o `doctor` e o inventário resolvem o owner por sinônimo público; Oracle de teste é 19c Standard Edition 2. | Fechada | 2026-10-25 |

## Modelo e tipos

| # | Pergunta | Provisório | Aberta em |
|---|---|---|---|
| Q6 | `NUMBER` sem precisão declarada na `raw`: `String` decimal ou `Decimal(38,10)`? | Decidir no mapeamento de tipos do E1.4 | 2026-10-25 |
| Q7 | Pontos do escopo que dependem só do dicionário (evidência em `docs/tables-inventory.md`): **#1** `PCDATAS` tem uma linha por data (7.824 linhas = 7.824 datas, sem `CODFILIAL`). **#2** `PCMOV` não tem PK, só a unique `PCMOV_UK (NUMTRANSITEM)` com a coluna aceitando nulo (ver Q10); `PCPEDI` e `PCITEM` = (`NUMPED`, `CODPROD`, `NUMSEQ`), únicas; `PCSECAO` = `CODSEC`. **#11** `PCRATEIOCENTROCUSTO` tem `VALOR NUMBER(12,2)` e `PERCRATEIO NUMBER(10,6)` (ver Q11). **#12** `CODFILIAL` é `VARCHAR2(2)` em todas as tabelas; `CODCLI NUMBER(9)`, `CODPROD NUMBER(6)`, `CODUSUR NUMBER(4)` (`NUMBER(8)` em `PCMOV`), `NUMTRANSITEM NUMBER(18)`, `NUMTRANSVENDA NUMBER(10)`, `NUMPED NUMBER(10)` (`NUMBER(6)` em `PCPEDIDO`), `RECNUM NUMBER(8)`. **#13** `PCNFENT` não tem `NUMPED`; a única candidata é `NUMPEDPRINC NUMBER(8)`; `PCMOV.NUMPED` está preenchido em 92% das linhas. **#14** `PCPRODUT.DTVENC` está preenchida em 18 de 18.567 produtos, com máximo em 2019: não serve como validade. | #1, #2 e #12 fechados; #11 e #13 dependem do negócio; #14: a análise de validade precisa de outra fonte (lote) | 2026-10-25 |
| Q8 | Pontos #3 a #10 de `docs/escopo-v1.md` (dependem de especialista de negócio). | Seguem como no escopo | 2026-10-25 |
| Q9 | Carimbos de alteração existem no dicionário de 22 tabelas (`DTULTALTER` em `PCCLIENT`, `PCPRODUT`, `PCFORNEC`, `PCPREST`, `PCLANC`, cadastros de produto; `DTULTALTERACAO` em `PCUSUARI`; `DTULTLALTER` em `PCPEDI`), e não existem em `PCMOV`, `PCNFSAID`, `PCNFENT`, `PCPEDC`, `PCEST` (lista completa em `docs/tables-inventory.md`, #Q9). Falta validar por teste controlado se o carimbo muda a cada alteração pelo ERP. | Data de negócio + janela móvel (45 dias) em todas; `DTULTALTER` extraída em `PCPREST` e `PCLANC` para avaliar no E1.8 | 2026-10-25 |

## Inventário (E1.2)

| # | Pergunta | Provisório | Aberta em |
|---|---|---|---|
| Q10 | `PCMOV.NUMTRANSITEM` aceita nulo: 644.975 das 64,0 M linhas (1,0%) têm a chave nula e, portanto, não têm identidade para dedup e reconciliação. São antigas? Existem dentro da janela desde 01/jan do ano anterior? | Extrair só linhas com `NUMTRANSITEM` não nulo e medir quantas ficam de fora na janela (E1.4) | 2026-10-25 |
| Q11 | `PCRATEIOCENTROCUSTO` não tem chave única no dicionário. `CODCENTROCUSTO` (numérico) é nulo em todas as 12.050 linhas da base de teste; (`RECNUM`, `CODIGOCENTROCUSTO`) é único e sem nulos. Vale para os outros clientes? | Chave (`RECNUM`, `CODIGOCENTROCUSTO`); o inventário de cada cliente reverifica | 2026-10-25 |
| Q12 | `PCMOVCOMPLE` (63 M linhas) e `PCITEM` (1,4 M) não têm data própria. Nenhum mart cita colunas de `PCMOVCOMPLE` (o `TIPOITEM` está em `PCMOV`). | `PCITEM`: faixa de `NUMPED` derivada da janela de `PCPEDIDO`. `PCMOVCOMPLE`: confirmar se ainda é necessária; se for, faixa de `NUMTRANSITEM` derivada de `PCMOV` | 2026-10-25 |
| Q13 | `PCESTCOM` (estorno de comissão?) está na onda 1 sem nenhuma coluna citada pelo escopo. Qual o uso? | Extrair só a chave até definir | 2026-10-25 |
| Q14 | `PCMETARCA` não tem `QTVENDAPREV`, `PEDIDOSPREV`, `CLIPOSPREV`, `MIXPREV`, `QTPESOPREV` nem `VOLUMEPREV`. Propostas: `QTPEDPREV`, `NUMCLIPOS`, `QTITENSPEDPREV`, `QTMETROPREV`; peso sem candidata. Ponto #3 do escopo (fonte oficial de meta de RCA) continua com o negócio. | Extrair as propostas e decidir na fase 2 | 2026-10-25 |
| Q15 | `PCLANC.DTLANC` é nula em 901 lançamentos: eles escapam da janela por data. | `open_items` cobre os em aberto; para os demais avaliar `DTEMISSAO` ou `DTULTALTER` no E1.8 | 2026-10-25 |
| Q16 | Grãos do escopo diferem das chaves reais: `PCCORTEI` = (`CODPROD`, `NUMCAR`, `NUMPED`, `NUMSEQ`), `PCFALTA` = (`NUMPED`, `CODPROD`, `NUMSEQ`), `PCNFENT` = (`NUMTRANSENT`, `CODCONT`), `PCEMBALAGEM` = (`CODFILIAL`, `CODAUXILIAR`), `PCLIB` = (`CODFUNC`, `CODTABELA`, `CODIGOA`, `CODIGON`), `PCCENTROCUSTO` = `CODIGOCENTROCUSTO`, `PCLINHAPROD` = `CODLINHA`. | A raw adota as chaves reais; os grãos dos marts são ajustados na fase 2 | 2026-10-25 |
| Q17 | Colunas citadas pelo escopo que não existem: `PCGERENTE.NOME` (é `NOMEGERENTE`), `PCTABDEV.DESCRICAO` (é `MOTIVO`), `PCMETA.TIPO` (é `TIPOMETA`), `PCPEDIDO.DATAENTREGA` (existe em `PCITEM`; cabeçalho tem `DTPREVENT`), `PCCONTA.CODGRUPOCONTA` (é `GRUPOCONTA`), `PCMOV.CONDVENDA` (está em `PCNFSAID` e `PCPEDI`), `PCPEDI.PESOLIQ`/`VOLUME` (vêm de `PCPRODUT`). | Substituídas nos `<TABELA>.yaml` com `use: proposed` e nota | 2026-10-25 |
| Q18 | Qual parâmetro de `PCPARAMFILIAL` define o preço por embalagem, e onde está `BLOQUEIAVENDAESTPENDENTE` (`PCCONSUM` tem a coluna; `PCPARAMFILIAL` guarda `NOME`/`VALOR` por filial)? | Extrair `PCCONSUM.BLOQUEIAVENDAESTPENDENTE` e `PCPARAMFILIAL` inteira (16 mil linhas) | 2026-10-25 |
| Q19 | Dados pessoais nas colunas extraídas: `PCCLIENT.CLIENTE`, `FANTASIA`, `BAIRROENT`; `PCUSUARI.NOME`, `EMAIL`; `PCSUPERV`, `PCGERENTE`, `PCEMPR.NOME`. O escopo permite nomes e e-mails de RCA/supervisor/gerente; nomes de clientes e de funcionários entram na dimensão. Confirmar que é aceitável sob LGPD (dado de pessoa jurídica na maioria, mas há pessoa física). | Extrair como no escopo; CPF/CNPJ, telefone e e-mail de cliente ficam fora | 2026-10-25 |
