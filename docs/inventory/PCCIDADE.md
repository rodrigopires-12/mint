# RIBERDOCES.PCCIDADE — dicionário

Gerado por `gateway inventory` em 2026-10-25. 10 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODCIDADE | NUMBER(6) | N |  |
| 2 | NOMECIDADE | VARCHAR2(80) | S |  |
| 3 | CODIBGE | NUMBER(10) | S |  |
| 4 | UF | VARCHAR2(2) | S |  |
| 5 | POPULACAO | NUMBER(8) | S | Indica o numero de habitantes da cidade. |
| 6 | CODMUNESTADUAL | NUMBER(8) | S | Indica o código do município no estado. |
| 7 | UTILIZAFRETETRANSP | VARCHAR2(1) | S | Identifica se irá utilizar o o processo de frete de transportadora para a cidade |
| 8 | CODMUNSIAFI | NUMBER(10) | S | Código Municipio SIAFI |
| 9 | LATITUDE | VARCHAR2(20) | S | Latitude geográfica da cidade |
| 10 | LONGITUDE | VARCHAR2(20) | S | Longitude geográfica da cidade |

## Chaves e índices

- Constraint PCCIDADE_PK (P): CODCIDADE
- Índice PCCIDADE_IDX1: NOMECIDADE
- Índice PCCIDADE_PK UNIQUE: CODCIDADE
