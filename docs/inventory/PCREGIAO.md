# RIBERDOCES.PCREGIAO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 33 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMREGIAO | NUMBER(4) | N |  |
| 2 | REGIAO | VARCHAR2(40) | N |  |
| 3 | PERFRETE | NUMBER(8,4) | N |  |
| 4 | UF | VARCHAR2(2) | S |  |
| 5 | VLFRETEKG | NUMBER(10,4) | S |  |
| 6 | CODFILIAL | VARCHAR2(2) | S |  |
| 7 | SEGMENTO | VARCHAR2(1) | S |  |
| 8 | PERFRETETERCEIROS | NUMBER(8,4) | S |  |
| 9 | PERFRETEESPECIAL | NUMBER(8,4) | S |  |
| 10 | STATUS | VARCHAR2(1) | S |  |
| 11 | TAREPF | VARCHAR2(1) | S |  |
| 12 | REGIAOZFM | VARCHAR2(1) | S |  |
| 13 | PERFRETECONHEC | NUMBER(8,4) | S |  |
| 14 | VLMINFATBK | NUMBER(14,6) | S |  |
| 15 | VLMINFATCH | NUMBER(14,6) | S |  |
| 16 | EXPORTAFV | NUMBER(1) | S |  |
| 17 | NUMTABELA | VARCHAR2(20) | S |  |
| 18 | CODESTABELECIMENTO | VARCHAR2(3) | S |  |
| 19 | ATUALIZAF11 | VARCHAR2(1) | S |  |
| 20 | OBS | VARCHAR2(60) | S | Campo para informar observações importantes sobre a Região.  |
| 21 | VLFRETEKGPADRAO | NUMBER(10,4) | S | Valor de Frete Kg por Região, que será gravado na Tabela de Preço (PCTABPR.VLACRESFRETEKG).  |
| 22 | VLFRETEKGVENDA | NUMBER(10,4) | S | Valor de Venda Kg por Região.  |
| 23 | CODREPRESENTANTE | VARCHAR2(4) | S |  |
| 24 | VLMINVENDA | NUMBER(18,6) | S | Valor mínimo de venda para cálculo da taxa de entrega |
| 25 | VLTXENT | NUMBER(18,6) | S | Valor da taxa de entrega |
| 26 | USAECOMMERCE | VARCHAR2(1) | S | Define se os preços da região serão exportados para o E-Commerce. |
| 27 | DTALTERC5 | TIMESTAMP(6) | S | Data de alteração |
| 28 | DTALTERACAO | TIMESTAMP(6) | S | Data e hora de alteração preenchido caso algum registro seja alterado ou excluído na tabela |
| 29 | DTCRIACAO | TIMESTAMP(6) | S | Data e hora de criação de registro na tabela. |
| 30 | TIPOEMPRESA | VARCHAR2(4) | S | Indica o tipo da empresa. |
| 31 | ORGAOPUB | VARCHAR2(1) | S | Tipo de Órgão Público |
| 32 | CONTRIBUINTE | VARCHAR2(1) | S | Contribuinte |
| 33 | CONSUMIDORFINAL | VARCHAR2(1) | S | Consumidor Final |

## Chaves e índices

- Constraint PCREGIAO_PK (P): NUMREGIAO
- Índice PCREGIAO_PK UNIQUE: NUMREGIAO
