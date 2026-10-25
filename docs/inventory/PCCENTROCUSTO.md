# RIBERDOCES.PCCENTROCUSTO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 8 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | DESCRICAO | VARCHAR2(40) | S | Descrição do Centro de Custo. |
| 2 | CODCENTROCUSTO | VARCHAR2(40) | S |  |
| 3 | RECEBE_LANCTO | VARCHAR2(1) | S | Indica se o Nível do Centro de custos recebe lançamentos |
| 4 | ATIVO | VARCHAR2(1) | S | Indica se o Nível do Centro de custos está ativo. |
| 5 | CODIGOCENTROCUSTO | VARCHAR2(40) | N | Código do centro de custo |
| 6 | CODIGOCENTROCUSTOINTFOLHA | VARCHAR2(40) | S | Cód. Centro de Custo - RM |
| 7 | DTINCLUSAO | DATE | S | Data de Inclusão |
| 8 | DTALTERACAO | DATE | S | Data de Alteração |

## Chaves e índices

- Constraint PK_PCCENTROCUSTO (P): CODIGOCENTROCUSTO
- Índice PK_PCCENTROCUSTO UNIQUE: CODIGOCENTROCUSTO
