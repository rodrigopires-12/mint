# RIBERDOCES.PCESTCOM — dicionário

Gerado por `gateway inventory` em 2026-10-25. 18 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMTRANSENT | NUMBER(10) | N |  |
| 2 | VLDEVOLUCAO | NUMBER(14,2) | N |  |
| 3 | VLESTORNO | NUMBER(14,2) | N |  |
| 4 | DTESTORNO | DATE | N |  |
| 5 | CODUSUR | NUMBER(4) | N |  |
| 6 | CODFUNC | NUMBER(8) | S |  |
| 7 | NUMTRANSVENDA | NUMBER(10) | S |  |
| 8 | HISTORICO | VARCHAR2(60) | S |  |
| 9 | DTLANC | DATE | S |  |
| 10 | VLESTORNOCMV | NUMBER(18,6) | S | Indica o valor estorno CMV. |
| 11 | CODUSUR2 | NUMBER(6) | S | Indica o código do primeiro profissional. |
| 12 | CODUSUR3 | NUMBER(6) | S | Indica o código do segundo profissional. |
| 13 | CODUSUR4 | NUMBER(6) | S | Indica o código do terceiro profissional. |
| 14 | VLESTORNO2 | NUMBER(14,2) | S | Data de pagamento comisão do primeiro profissional. |
| 15 | VLESTORNO3 | NUMBER(14,2) | S | Data de pagamento comisão do segundo profissional. |
| 16 | VLESTORNO4 | NUMBER(14,2) | S | Data de pagamento comisão do segundo profissional. |
| 17 | DTPAGCOMISSAO | DATE | S | Data de pagamento da comissão referente a devolução |
| 18 | DTPAGCOMISSAOOP | DATE | S | Fechamento de comissão do Operador |

## Chaves e índices

- Constraint PCESTCOM_PK (P): NUMTRANSENT
- Índice PCESTCOM_IDX1: NUMTRANSVENDA
- Índice PCESTCOM_IDX2: CODUSUR, VLESTORNO
- Índice PCESTCOM_IDX3: CODUSUR, DTPAGCOMISSAO, DTPAGCOMISSAOOP
- Índice PCESTCOM_IDX5: NUMTRANSVENDA, NUMTRANSENT
- Índice PCESTCOM_PK UNIQUE: NUMTRANSENT
