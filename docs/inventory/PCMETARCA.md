# RIBERDOCES.PCMETARCA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 13 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODFILIAL | VARCHAR2(2) | N |  |
| 2 | CODUSUR | NUMBER(4) | N |  |
| 3 | DATA | DATE | N |  |
| 4 | PERVENDAPREV | NUMBER(5,2) | S |  |
| 5 | VLVENDAPREV | NUMBER(14,2) | S |  |
| 6 | VLVENDA | NUMBER(14,2) | S |  |
| 7 | VLCUSTOREAL | NUMBER(14,2) | S |  |
| 8 | VLCUSTOFIN | NUMBER(14,2) | S |  |
| 9 | VLTABELA | NUMBER(14,2) | S |  |
| 10 | QTITENSPEDPREV | NUMBER(14,2) | S |  |
| 11 | QTPEDPREV | NUMBER(14,2) | S |  |
| 12 | QTMETROPREV | NUMBER(20,8) | S |  |
| 13 | NUMCLIPOS | NUMBER(20,8) | S | Indica a meta de clientes para positivação. |

## Chaves e índices

- Constraint PCMETARCA_PK (P): CODFILIAL, CODUSUR, DATA
- Índice PCMETARCA_PK UNIQUE: CODFILIAL, CODUSUR, DATA
