# RIBERDOCES.PCROTAEXP — dicionário

Gerado por `gateway inventory` em 2026-10-25. 75 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODROTA | NUMBER(4) | N |  |
| 2 | DESCRICAO | VARCHAR2(40) | S |  |
| 3 | PERCOMMOT | NUMBER(6,3) | S | Descricao coluna PERCOMMOT |
| 4 | SITUACAO | VARCHAR2(1) | S |  |
| 5 | VLCOMVEICLEVE | NUMBER(10,2) | S |  |
| 6 | VLCOMVEICMEDIO | NUMBER(10,2) | S |  |
| 7 | VLCOMVEICPESADO | NUMBER(10,2) | S |  |
| 8 | VLKMADICIONAL | NUMBER(10,2) | S |  |
| 9 | KMROTA | NUMBER(10,2) | S |  |
| 10 | KMEXCLUIDO | NUMBER(10,2) | S |  |
| 11 | VLMINCARREG | NUMBER(10,2) | S |  |
| 12 | DTSITUACAO | DATE | S |  |
| 13 | VLFRETETON | NUMBER(10,4) | S |  |
| 14 | VLFRETEENTREGA | NUMBER(10,4) | S |  |
| 15 | PERCOMMOT2 | NUMBER(6,3) | S | Descricao coluna PERCOMMOT2 |
| 16 | PERCOMMOT3 | NUMBER(6,3) | S | Descricao coluna PERCOMMOT3 |
| 17 | DIASEG | VARCHAR2(1) | S | Descricao coluna DIASEG |
| 18 | DIATER | VARCHAR2(1) | S | Descricao coluna DIATER |
| 19 | DIAQUA | VARCHAR2(1) | S | Descricao coluna DIAQUA |
| 20 | DIAQUI | VARCHAR2(1) | S | Descricao coluna DIAQUI |
| 21 | DIASEX | VARCHAR2(1) | S | Descricao coluna DIASEX |
| 22 | DIASAB | VARCHAR2(1) | S | Descricao coluna DIASAB |
| 23 | DIADOM | VARCHAR2(1) | S | Descricao coluna DIADOM |
| 24 | PRAZOPREVENT | NUMBER(5,1) | S |  |
| 25 | QTENTREGA | NUMBER(4) | S |  |
| 26 | PERFRETEVALOR | NUMBER(5,2) | S |  |
| 27 | ALIQICMSFRETE | NUMBER(8,4) | S |  |
| 28 | TIPOCOMISSAO | VARCHAR2(1) | S |  |
| 29 | VALORCOMMOT | NUMBER(10,2) | S |  |
| 30 | SEQENTREGA | NUMBER(4) | S |  |
| 31 | VLDIARIA | NUMBER(10,4) | S |  |
| 32 | NUMDIARIA | NUMBER(5,1) | S | Descricao coluna NUMDIARIA |
| 33 | RASTREADA | VARCHAR2(1) | S | Indica se a rota possui rastreamento, caso a carga possua pelo menos uma rota rastreada, será exibida mensagem de alerta caso o veículo selecionado não seja rastreado.  |
| 34 | CODROTAPRINCIPAL | NUMBER(4) | S | Indica o código da rotina principal. |
| 35 | UTILIZAVLDESCARGAFIXO | VARCHAR2(1) | S | Define se utiliza valor fixo da descarga. |
| 36 | VLDESCARGAFIXO | NUMBER(18,6) | S | Define so valor fixo para descarga. |
| 37 | PERCDIARIA | NUMBER(5,2) | S | Indica o percentual comissão diária. |
| 38 | TIPODIARIA | VARCHAR2(1) | S | Indica o tipo de comissão diária (P)percentual (V)valor. |
| 39 | VLAJUDA | NUMBER(6,2) | S | Campo referente ao valor de ajuda de custo com motorista. |
| 40 | PERCAJUDA | NUMBER(5,2) | S | Percentual de ajuda de custo. |
| 41 | PERCMINCARREG | NUMBER(10,2) | S | Variação do valor minimo de carga |
| 42 | DIASENTREGA | NUMBER(5,1) | S | Dias de entrega |
| 43 | PROXNUMNFETV14 | NUMBER(9) | S | Próximo número de NF-e |
| 44 | SERIENFETV14 | NUMBER(3) | S | Série NFe |
| 45 | PROXNUMNFESCANTV14 | NUMBER(9) | S | Próximo número de NF-e |
| 46 | SERIENFESCANTV14 | NUMBER(3) | S | Série NFe |
| 47 | CODESTEIRA | NUMBER(6) | S | Relação com código tabela pcesteirawms |
| 48 | SEGMANHA | NUMBER(4) | S |  |
| 49 | SEGTARDE | NUMBER(4) | S |  |
| 50 | SEGNOITE | NUMBER(4) | S |  |
| 51 | TERMANHA | NUMBER(4) | S |  |
| 52 | TERTARDE | NUMBER(4) | S |  |
| 53 | TERNOITE | NUMBER(4) | S |  |
| 54 | QUAMANHA | NUMBER(4) | S |  |
| 55 | QUATARDE | NUMBER(4) | S |  |
| 56 | QUANOITE | NUMBER(4) | S |  |
| 57 | QUIMANHA | NUMBER(4) | S |  |
| 58 | QUITARDE | NUMBER(4) | S |  |
| 59 | QUINOITE | NUMBER(4) | S |  |
| 60 | SEXMANHA | NUMBER(4) | S |  |
| 61 | SEXTARDE | NUMBER(4) | S |  |
| 62 | SEXNOITE | NUMBER(4) | S |  |
| 63 | SABMANHA | NUMBER(4) | S |  |
| 64 | SABTARDE | NUMBER(4) | S |  |
| 65 | SABNOITE | NUMBER(4) | S |  |
| 66 | DOMMANHA | NUMBER(4) | S |  |
| 67 | DOMTARDE | NUMBER(4) | S |  |
| 68 | DOMNOITE | NUMBER(4) | S |  |
| 69 | IDSOFITVIEW | VARCHAR2(10) | S | Indica o código da rota na SofitView |
| 70 | DTULTALTERSOFITVIEW | TIMESTAMP(6) | S | Indica a data que a rota foi integrado com a SofitView |
| 71 | DESP_CAFEMANHA | NUMBER(6,2) | S | Despesa com café da manhã |
| 72 | DESP_ALMOCO | NUMBER(6,2) | S | Despesa com almoço |
| 73 | DESP_JANTAR | NUMBER(6,2) | S | Despesa com jantar |
| 74 | DESP_PERNOITE | NUMBER(6,2) | S | Despesa com pernoite |
| 75 | DESP_EXTRA | NUMBER(6,2) | S | 	Outras despesas |

## Chaves e índices

- Constraint PCROTAEXP_PK (P): CODROTA
- Índice PCROTAEXP_PK UNIQUE: CODROTA
