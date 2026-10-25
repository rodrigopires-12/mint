# RIBERDOCES.PCPRACA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 25 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODPRACA | NUMBER(4) | N |  |
| 2 | PRACA | VARCHAR2(25) | N |  |
| 3 | NUMREGIAO | NUMBER(4) | N |  |
| 4 | ROTA | NUMBER(4) | N |  |
| 5 | SEQROTA | NUMBER(4) | N |  |
| 6 | POPULACAO | NUMBER(14) | S |  |
| 7 | PERFRETEPROGRESS | NUMBER(8,4) | S |  |
| 8 | CODPRACAORIG | NUMBER(6) | S |  |
| 9 | DISTANCIA | NUMBER(12,2) | S |  |
| 10 | VLPAUTAFRETE | NUMBER(12,2) | S |  |
| 11 | CODPRACAORIG2 | NUMBER(6) | S |  |
| 12 | CODPRACAORIG3 | NUMBER(6) | S |  |
| 13 | CODPRACAORIG4 | NUMBER(6) | S |  |
| 14 | CODMUNIC | NUMBER(8) | S |  |
| 15 | SITUACAO | VARCHAR2(1) | S |  |
| 16 | NUMREGIAO2 | NUMBER(4) | S |  |
| 17 | DTCADASTRO | DATE | S | Data em que foi efetuado o cadastro da praça. |
| 18 | CEPINICIAL | VARCHAR2(9) | S | Código do CEP Inicial da Praça. |
| 19 | CEPFINAL | VARCHAR2(9) | S | Código do CEP Final da Praça. |
| 20 | PRIORIDADEENTREGA | VARCHAR2(1) | S | Prioridade de entrega de pedidos na praça. |
| 21 | TIPOPRACA | VARCHAR2(1) | S | Tipo de Praça F ou P - Frete ou Praça. |
| 22 | OBS | VARCHAR2(300) | S | Indica as obsevações para o cadastro de praça. |
| 23 | CODPRACAPRINCIPAL | NUMBER(4) | S | Indica o código da praça principal. |
| 24 | VLMINCARREG | NUMBER(12,2) | S | Vl. Mínimo Montagem Carga. |
| 25 | PERCMINCARREG | NUMBER(10,2) | S | Variação do valor minimo de carga |

## Chaves e índices

- Constraint PCPRACA_PK (P): CODPRACA
- Índice PCPRACA_IDX1: ROTA
- Índice PCPRACA_IX1: ROTA, SEQROTA
- Índice PCPRACA_PK UNIQUE: CODPRACA
