# RIBERDOCES.PCPARAMFILIAL — dicionário

Gerado por `gateway inventory` em 2026-10-25. 3 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NOME | VARCHAR2(34) | N | Nome do parâmetro. |
| 2 | CODFILIAL | VARCHAR2(2) | N | Código da filial que o parâmetro está relacionado. Use 999 para valores por empresa, e não por filial. |
| 3 | VALOR | VARCHAR2(1000) | S | Valor do parâmetros. O valor do parâmetro deve ser convertido para varchar2 antes de salvar. |

## Chaves e índices

- Constraint PCPARAMFILIAL_PK (P): NOME, CODFILIAL
- Índice PCPARAMFILIAL_PK UNIQUE: NOME, CODFILIAL
