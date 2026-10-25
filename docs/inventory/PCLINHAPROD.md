# RIBERDOCES.PCLINHAPROD — dicionário

Gerado por `gateway inventory` em 2026-10-25. 10 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODLINHA | NUMBER(6) | N |  |
| 2 | DESCRICAO | VARCHAR2(40) | S |  |
| 3 | CODFILIAL | VARCHAR2(2) | S |  |
| 4 | NUMDIASTRABSEMANA | NUMBER(4) | S |  |
| 5 | NUMTURNODIA | NUMBER(4) | S |  |
| 6 | HORASPORTURNO | NUMBER(4) | S |  |
| 7 | NUMPESSOASTURNO | NUMBER(4) | S |  |
| 8 | CODLINHAPROD | NUMBER(6) | S |  |
| 9 | PRAZOMAXIMO | NUMBER(4) | S |  |
| 10 | OBS | VARCHAR2(100) | S | Indica a observação da linha de produto. |

## Chaves e índices

- Constraint PCLINHAPROD_PK (P): CODLINHA
- Índice PCLINHAPROD_PK UNIQUE: CODLINHA
