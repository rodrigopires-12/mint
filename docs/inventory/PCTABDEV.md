# RIBERDOCES.PCTABDEV — dicionário

Gerado por `gateway inventory` em 2026-10-25. 14 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODDEVOL | NUMBER(4) | N |  |
| 2 | MOTIVO | VARCHAR2(30) | N |  |
| 3 | PERCESTCOM | NUMBER(6,2) | S |  |
| 4 | ESTORNACOMISSAO | VARCHAR2(1) | S |  |
| 5 | TIPO | VARCHAR2(2) | S |  |
| 6 | CODCONTA | NUMBER(10) | S |  |
| 7 | CODMOTNESTLE | NUMBER(6) | S |  |
| 8 | CODMOTFORNECNESTLE | NUMBER(6) | S |  |
| 9 | PERCESTCOMMOT | NUMBER(6,2) | S | Percentual de estorno para comissão de motorista. Será aplicado só na geração do relatório de apuração de comissão para motorista.  |
| 10 | BLOQUEIACLIENTE | VARCHAR2(1) | S | Motivo da devolução determina uma possível condição de bloqueio do cliente. |
| 11 | MANTERDEBITODESCFIN | VARCHAR2(1) | S | Manter o débito relativo ao desconto financeiro. |
| 12 | MANTERCREDITO | VARCHAR2(1) | S | Manter o débito relativo ao desconto financeiro. |
| 13 | RETORNOVENDANAOENTREGUEAODEST | VARCHAR2(1) | S | Define é uma operação de Retorno de Venda não entregue ao destinatário |
| 14 | CODMOTDEVDORI | NUMBER(5) | S | Código do motivo de devolução com a integração DORI |

## Chaves e índices

- Constraint PCTABDEV_PK (P): CODDEVOL
- Índice PCTABDEV_IDX1: TIPO, MOTIVO
- Índice PCTABDEV_IDX2: TIPO, CODDEVOL
- Índice PCTABDEV_PK UNIQUE: CODDEVOL
