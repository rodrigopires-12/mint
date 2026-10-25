# RIBERDOCES.PCSUPERV — dicionário

Gerado por `gateway inventory` em 2026-10-25. 18 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODSUPERVISOR | NUMBER(4) | N |  |
| 2 | NOME | VARCHAR2(40) | N |  |
| 3 | REGIONAL | NUMBER(2) | N |  |
| 4 | COD_CADRCA | NUMBER(4) | S |  |
| 5 | POSICAO | VARCHAR2(1) | S |  |
| 6 | PERCPARTVENDAPREV | NUMBER(8,4) | S |  |
| 7 | PERCMARGEMPREV | NUMBER(8,4) | S |  |
| 8 | CODGERENTE | NUMBER(4) | S |  |
| 9 | TIPOSUPERVISOR | VARCHAR2(1) | S |  |
| 10 | PERCOMISSAO | NUMBER(4,2) | S |  |
| 11 | DTADMISSAO | DATE | S |  |
| 12 | DTDEMISSAO | DATE | S | Indica a data de emissão. |
| 13 | CPF | VARCHAR2(20) | S | Indica o numero do CPF. |
| 14 | CODCOORDENADOR | NUMBER(10) | S | Indica o Coordenador de venda. |
| 15 | EMAIL | VARCHAR2(100) | S | Email do supervisor |
| 16 | VLCORRENTE | NUMBER(22,6) | S | Valor do conta corrente |
| 17 | VLLIMCRED | NUMBER(22,6) | S | Valor do limite do crédito |
| 18 | USADEBCRED | VARCHAR2(1) | S | Usa débito e crédito |

## Chaves e índices

- Constraint PCSUPERV_PK (P): CODSUPERVISOR
- Índice PCSUPERV_PK UNIQUE: CODSUPERVISOR
