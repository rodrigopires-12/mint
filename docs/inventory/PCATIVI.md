# RIBERDOCES.PCATIVI — dicionário

Gerado por `gateway inventory` em 2026-10-25. 15 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODATIV | NUMBER(6) | N | Código atividade. |
| 2 | RAMO | VARCHAR2(40) | N | Ramo. |
| 3 | PERCDESC | NUMBER(5,2) | S | Percentual de desconto. |
| 4 | CODCP | VARCHAR2(6) | S | Código CP. |
| 5 | CODCATEGORIAKRAFT | NUMBER(2) | S | Código categoria KRAFT. |
| 6 | CODRAMONESTLE | NUMBER(6) | S | Código ramo Nestle. |
| 7 | CODCATEGORIABAUDU | VARCHAR2(10) | S | Código Categoria BAUDU. |
| 8 | ATACADISTA | VARCHAR2(1) | S | Campo para identificação se o ramo de atividade é atacadista ou não. |
| 9 | COR | VARCHAR2(30) | S | Cor. |
| 10 | CODATIVPRINC | NUMBER(6) | S | Indica o código do ramo de atividade principal.. |
| 11 | PERCREDALIQIPI | NUMBER(18,6) | S | Percentual de Redução da Alíquota de IPI |
| 12 | CALCULAST | VARCHAR2(1) | S | Define se irá calcular ST ou não para o ramo de atividade |
| 13 | CATVAREJODORI | VARCHAR2(5) | S | Categoria do varejo para integração com DORI |
| 14 | DATAALTERACAO | DATE | S | Indica a última vez que a tabela foi modificada |
| 15 | DATACADASTRO | DATE | S | Indica a data de cadastro |

## Chaves e índices

- Constraint PCATIVI_PK (P): CODATIV
- Índice PCATIVI_PK UNIQUE: CODATIV
