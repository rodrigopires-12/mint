# RIBERDOCES.PCCATEGORIA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 8 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODSEC | NUMBER(6) | N |  |
| 2 | CODCATEGORIA | NUMBER(6) | N |  |
| 3 | CATEGORIA | VARCHAR2(40) | S |  |
| 4 | IDINTEGRACAOCIASHOP | VARCHAR2(250) | S | Referência da categoria no E-commerce CiaShop |
| 5 | ENVIAECOMMERCE | VARCHAR2(1) | S | Identifica se o produto será enviado ao e-commerce |
| 6 | DTULTALTER | DATE | S | Data da ultima alteração |
| 7 | DTCADASTRO | DATE | S | Data de cadastro |
| 8 | DTALTERC5 | TIMESTAMP(6) | S | Data alteracao |

## Chaves e índices

- Constraint PCCATEGORIA_PK (P): CODSEC, CODCATEGORIA
- Índice PCCATEGORIA_PK UNIQUE: CODSEC, CODCATEGORIA
