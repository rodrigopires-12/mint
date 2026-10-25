# RIBERDOCES.PCUSURCLI — dicionário

Gerado por `gateway inventory` em 2026-10-25. 5 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODUSUR | NUMBER(4) | N |  |
| 2 | CODCLI | NUMBER(9) | N |  |
| 3 | TIPOOPER | VARCHAR2(1) | S |  |
| 4 | CODUSURPADRAO | NUMBER(6) | S | Indica o Código do RCA padrão(iniciar venda) |
| 5 | ENVIAFV | VARCHAR2(1) | S | Limitar envio de informações par FV |

## Chaves e índices

- Constraint PCUSURCLI_PK (P): CODUSUR, CODCLI
- Índice PCUSURCLI_IDX1: CODCLI, CODUSUR
- Índice PCUSURCLI_PK UNIQUE: CODUSUR, CODCLI
