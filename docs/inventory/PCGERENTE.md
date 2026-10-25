# RIBERDOCES.PCGERENTE — dicionário

Gerado por `gateway inventory` em 2026-10-25. 10 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODGERENTE | NUMBER(4) | N |  |
| 2 | NOMEGERENTE | VARCHAR2(40) | S |  |
| 3 | VLCORRENTE | NUMBER(22,6) | S | Valor do conta corrente |
| 4 | VLLIMCRED | NUMBER(22,6) | S | Valor do limite do crédito |
| 5 | USADEBCRED | VARCHAR2(1) | S | Usa débito e crédito |
| 6 | COD_CADRCA | NUMBER(4) | S | Código de RCA que representa o gerente |
| 7 | CPF | VARCHAR2(20) | S | CPF do Gerente |
| 8 | EMAIL | VARCHAR2(100) | S | email do Gerente |
| 9 | TIPO | VARCHAR2(1) | S | Tipo de Gerente |
| 10 | CODGERENTESUPERIOR | NUMBER(4) | S | Código do Gerente Superior |

## Chaves e índices

- Constraint PCGERENTE_PK (P): CODGERENTE
- Índice PCGERENTE_PK UNIQUE: CODGERENTE
