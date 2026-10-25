# RIBERDOCES.PCFALTA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 9 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMPED | NUMBER(10) | N |  |
| 2 | DATA | DATE | N |  |
| 3 | CODPROD | NUMBER(6) | N |  |
| 4 | CODUSUR | NUMBER(4) | S |  |
| 5 | CODCLI | NUMBER(9) | S |  |
| 6 | QT | NUMBER(8,2) | S |  |
| 7 | PVENDA | NUMBER(12,3) | S |  |
| 8 | CODFILIAL | VARCHAR2(2) | S |  |
| 9 | NUMSEQ | NUMBER(20) | N |  |

## Chaves e índices

- Constraint PCFALTA1_PK (P): NUMPED, CODPROD, NUMSEQ
- Índice PCFALTA1_PK UNIQUE: NUMPED, CODPROD, NUMSEQ
- Índice PCFALTA_IDX1: CODPROD, DATA
- Índice PCFALTA_IDX2: CODCLI, DATA
- Índice PCFALTA_IDX3: CODUSUR, DATA
- Índice PCFALTA_IDX4: CODFILIAL
