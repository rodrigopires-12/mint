# RIBERDOCES.PCLIB — dicionário

Gerado por `gateway inventory` em 2026-10-25. 6 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODFUNC | NUMBER(8) | N |  |
| 2 | CODTABELA | NUMBER(4) | N |  |
| 3 | CODIGOA | VARCHAR2(40) | N | Código do cadastro. |
| 4 | CODIGON | NUMBER(10) | N |  |
| 5 | CODFUNC_LIB | NUMBER(8) | S |  |
| 6 | DATA_LIB | DATE | S |  |

## Chaves e índices

- Constraint PCLIB_PK (P): CODFUNC, CODTABELA, CODIGOA, CODIGON
- Índice IDX$$_D36F0004: CODIGON, CODFUNC, CODTABELA
- Índice PCLIB_IDX2: CODFUNC, CODTABELA, CODIGON
- Índice PCLIB_IDX3: CODTABELA, CODFUNC, CODIGOA
- Índice PCLIB_IDXPCLIB_IDX01: CODTABELA, CODFUNC, CODIGON
- Índice PCLIB_PK UNIQUE: CODFUNC, CODTABELA, CODIGOA, CODIGON
