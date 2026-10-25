# RIBERDOCES.PCSUBCATEGORIA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 9 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODSEC | NUMBER(6) | N |  |
| 2 | CODCATEGORIA | NUMBER(6) | N |  |
| 3 | CODSUBCATEGORIA | NUMBER(6) | N |  |
| 4 | SUBCATEGORIA | VARCHAR2(40) | S |  |
| 5 | IDINTEGRACAOCIASHOP | VARCHAR2(250) | S | Referência da sub-categoria no E-commerce CiaShop |
| 6 | ENVIAECOMMERCE | VARCHAR2(1) | S | Identifica se o produto será enviado ao e-commerce |
| 7 | DTULTALTER | DATE | S | Data da última alteração feita na subcategoria. |
| 8 | DTCADASTRO | DATE | S | Data do cadastro da subcategoria |
| 9 | DTALTERC5 | TIMESTAMP(6) | S | DATA ALTERACAO |

## Chaves e índices

- Constraint PCSUBCATEGORIA_PK (P): CODSEC, CODCATEGORIA, CODSUBCATEGORIA
- Índice PCSUBCATEGORIA_PK UNIQUE: CODSEC, CODCATEGORIA, CODSUBCATEGORIA
