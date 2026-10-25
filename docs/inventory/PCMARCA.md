# RIBERDOCES.PCMARCA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 11 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODMARCA | NUMBER(8) | N | Indica o código da marca. |
| 2 | MARCA | VARCHAR2(40) | S | Indica o descrição da marca. |
| 3 | ATIVO | VARCHAR2(1) | S | Indica se a marca esta ativa ou inativa. |
| 4 | CODADWORDS | VARCHAR2(200) | S | Código do AdWords |
| 5 | DESCRICAOECOMMERCE | VARCHAR2(400) | S | Descrição da marca apresentado no ecommerce |
| 6 | CODCAMPLOMADEE | VARCHAR2(200) | S | Código campanha Lomadee |
| 7 | TITULO | VARCHAR2(200) | S | Titulo da marca apresentado no ecommerce |
| 8 | DTULTALTER | DATE | S | Data da ultima alteração do registro. |
| 9 | DTCADASTRO | DATE | S | Data de cadastro |
| 10 | CODCOMPRADOR | NUMBER(8) | S | Código do comprador |
| 11 | DTALTERC5 | TIMESTAMP(6) | S | DATA ALTERACAO |

## Chaves e índices

- Constraint PCMARCA_PK (P): CODMARCA
- Índice PCMARCA_PK UNIQUE: CODMARCA
