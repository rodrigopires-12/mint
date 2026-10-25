# RIBERDOCES.PCRATEIOCENTROCUSTO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 20 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | RECNUM | NUMBER(8) | N | Número do lançamento. |
| 2 | CODCONTA | NUMBER(10) | N | Código da Conta Gerencial. |
| 3 | CODCENTROCUSTO | NUMBER(10) | S | Código do Centro de Custo. |
| 4 | VALOR | NUMBER(12,2) | N | Valor do lançamento. |
| 5 | PERCRATEIO | NUMBER(10,6) | N | Percentual de Rateio do lançamento. |
| 6 | DTLANC | DATE | N | Data do lançamento. |
| 7 | CONTRAPARTIDA | VARCHAR2(1) | S | Lançamento teve contra-partida |
| 8 | RECNUMPRINC | NUMBER(10) | S | Número do lançamento principal |
| 9 | CODFILIAL | VARCHAR2(2) | S | Código da filial do centro de custo |
| 10 | CODIGOCENTROCUSTO | VARCHAR2(40) | S | Código do centro de custo |
| 11 | ROTINAUPDATE | VARCHAR2(100) | S | Campo usado pela trigger TRG_LOG_PCRATEIOCENTROCUSTO para acompanhar alterações na tabela |
| 12 | ROTINAINSERT | VARCHAR2(100) | S | Campo usado pela trigger TRG_LOG_PCRATEIOCENTROCUSTO para acompanhar alterações na tabela |
| 13 | CODFUNCAUTSUPLORC | NUMBER(8) | S | CODIGO FUNCIONARIO AUTORIZACAO SUPLEMENTACAO ORCAMENTARIA |
| 14 | DATAAUTSUPLORC | DATE | S | DATA AUTORIZACAO SUPLEMENTACAO ORCAMENTARIA |
| 15 | ORCAMENTOEXCEDIDO | VARCHAR2(1) | S | IDENTIFICA QUAL CENTRO DE CUSTO DO RATEIO EXCEDEU O ORÇAMENTO |
| 16 | NUMTRANSPISCOFINS | NUMBER(10) | S | Numero de transação do PISCONFINS |
| 17 | RECNUMREF | NUMBER(10) | S | Recnum ao que o imposto pertence |
| 18 | PERCRATCONTREF | NUMBER(10,6) | S | Percentual da conta ao que o imposto pertence |
| 19 | TIPOTRIB | VARCHAR2(20) | S | Tipo do tributo |
| 20 | LANCIMPRETIDO | VARCHAR2(1) | S | Indica se é um lançamento de imposto retido |

## Chaves e índices

- Índice PCRATEIOCENTROCUSTO_IDX1: RECNUM
- Índice PCRATEIOCENTROCUSTO_IDX2: RECNUM, CODCONTA, CODIGOCENTROCUSTO
