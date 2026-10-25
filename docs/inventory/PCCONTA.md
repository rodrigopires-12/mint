# RIBERDOCES.PCCONTA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 20 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODCONTA | NUMBER(10) | N |  |
| 2 | CONTA | VARCHAR2(40) | N |  |
| 3 | GRUPOCONTA | NUMBER(4) | N |  |
| 4 | CODCONTAMASTER | NUMBER(10) | S |  |
| 5 | TIPO | VARCHAR2(1) | S |  |
| 6 | INVESTIMENTO | VARCHAR2(1) | N |  |
| 7 | BONIFIC | VARCHAR2(1) | S |  |
| 8 | VLORCAMES | NUMBER(14,2) | S |  |
| 9 | FIXAVARIAVEL | VARCHAR2(1) | S |  |
| 10 | GERAPROVLANCCONTAB | VARCHAR2(1) | S |  |
| 11 | CODCONTACONTRAPARTIDA | NUMBER(10) | S |  |
| 12 | CODCENTROCUSTOPADRAO | NUMBER(10) | S | Código do Centro de Custo Padrão |
| 13 | USARATEIOCENTROCUSTO | VARCHAR2(1) | N | Usa rateio por Centro de Custo? |
| 14 | VERBA | VARCHAR2(1) | S | Conta contábil de verba. |
| 15 | CONTACONTABIL | VARCHAR2(12) | S |  |
| 16 | CODEVENTOINTFOLHA | VARCHAR2(1000) | S | Cód. Evento folha - RM |
| 17 | CODSECAOINTFOLHA | VARCHAR2(50) | S | Código da seção integração folha |
| 18 | RESTRINGIRNOBALANCETE | VARCHAR2(1) | S | Campo que define a restricao de apresentacao da conta no balancete (rotina 124) |
| 19 | UTILIZACENTROCUSTORESTRITO | VARCHAR2(2) | S | Indica se a conta utiliza restrição de centro de custo |
| 20 | PRESTACAODECONTAS | VARCHAR2(1) | S | É conta de prestação de contas? |

## Chaves e índices

- Constraint PCCONTA_PK (P): CODCONTA
- Índice PCCONTA_IDX1: CODCONTA, GRUPOCONTA, INVESTIMENTO
- Índice PCCONTA_PK UNIQUE: CODCONTA
