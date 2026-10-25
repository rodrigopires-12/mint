# RIBERDOCES.PCCORTEI — dicionário

Gerado por `gateway inventory` em 2026-10-25. 25 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODPROD | NUMBER(6) | N |  |
| 2 | QTSEPARADA | NUMBER(8,2) | N |  |
| 3 | QTCORTADA | NUMBER(8,2) | N |  |
| 4 | DATA | DATE | N |  |
| 5 | NUMCAR | NUMBER(8) | N |  |
| 6 | CODFUNC | NUMBER(8) | N |  |
| 7 | NUMPED | NUMBER(10) | N |  |
| 8 | PVENDA | NUMBER(12,3) | S |  |
| 9 | CODFUNCCONF | NUMBER(8) | S |  |
| 10 | CODFUNCSEP | NUMBER(8) | S |  |
| 11 | CODFILIAL | VARCHAR2(2) | S |  |
| 12 | QTORIG | NUMBER(20,6) | S |  |
| 13 | QTFALTA | NUMBER(8,2) | S |  |
| 14 | MOTIVO | VARCHAR2(80) | S | Indica o motivo do corte.  |
| 15 | DTFINALCHECKOUT | DATE | S |  |
| 16 | HORA | NUMBER(2) | S |  |
| 17 | MINUTO | NUMBER(2) | S |  |
| 18 | CODCLI | NUMBER(9) | S | Cliente do pedido onde foi lançado corte.  |
| 19 | CODUSUR | NUMBER(4) | S | RCA do pedido onde foi lançado corte.  |
| 20 | CODROTINA | NUMBER(6) | S | Indica a rotina que efetuou o processo de corte.  |
| 21 | CONDVENDA | NUMBER(5) | S | Indica a condição de venda.  |
| 22 | CODEMITENTEPED | NUMBER(8) | S |  |
| 23 | NUMLOTE | VARCHAR2(15) | S | numero do lote |
| 24 | NUMSEQ | NUMBER(20) | N | Número sequência |
| 25 | TIPOCORTE | VARCHAR2(1) | S | Gravar o itpo do corte: "C" - Corte ou "D" - Depuração |

## Chaves e índices

- Constraint PCCORTEI_PK (P): CODPROD, NUMCAR, NUMPED, NUMSEQ
- Índice PCCORTEI_IDX1: NUMPED, CODPROD
- Índice PCCORTEI_IDX2: DATA, QTCORTADA, NUMPED
- Índice PCCORTEI_PK UNIQUE: CODPROD, NUMCAR, NUMPED, NUMSEQ
