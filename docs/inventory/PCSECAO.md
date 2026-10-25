# RIBERDOCES.PCSECAO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 13 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODSEC | NUMBER(6) | N |  |
| 2 | DESCRICAO | VARCHAR2(40) | N |  |
| 3 | CODEPTO | NUMBER(6) | N |  |
| 4 | QTMAX | NUMBER(10,3) | S |  |
| 5 | TIPO | VARCHAR2(1) | S |  |
| 6 | CODSECNESTLE | NUMBER(6) | S |  |
| 7 | LINHA | VARCHAR2(20) | S | Linha do produto |
| 8 | DTEXCLUSAO | DATE | S | Data exclusão produto |
| 9 | IDINTEGRACAOCIASHOP | VARCHAR2(250) | S | Referência da seção no E-commerce CiaShop |
| 10 | ENVIAECOMMERCE | VARCHAR2(1) | S | Identifica se o produto será enviado ao e-commerce |
| 11 | DTULTALTER | DATE | S | Data da ultima alteração |
| 12 | DTCADASTRO | DATE | S | Data de cadastro |
| 13 | DTALTERC5 | TIMESTAMP(6) | S | Data alteracao |

## Chaves e índices

- Constraint PCSECAO_PK (P): CODSEC
- Índice PCSECAO_INDEX2: CODEPTO, DESCRICAO, CODSEC
- Índice PCSECAO_PK UNIQUE: CODSEC
