# RIBERDOCES.PCMETA — dicionário

Gerado por `gateway inventory` em 2026-10-25. 53 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODIGO | NUMBER(8) | N |  |
| 2 | CODFILIAL | VARCHAR2(2) | N |  |
| 3 | CODUSUR | NUMBER(4) | N |  |
| 4 | TIPOMETA | VARCHAR2(2) | N |  |
| 5 | DATA | DATE | N |  |
| 6 | VLVENDAPREV | NUMBER(12,2) | S |  |
| 7 | QTVENDAPREV | NUMBER(20,6) | S |  |
| 8 | QTPESOPREV | NUMBER(12,2) | S |  |
| 9 | MIXPREV | NUMBER(10,4) | S |  |
| 10 | CLIPOSPREV | NUMBER(10,4) | S |  |
| 11 | VOLUMEPREV | NUMBER(20,8) | S |  |
| 12 | CODPROMOCAO | NUMBER(6) | S |  |
| 13 | CLICADPREV | NUMBER(10,4) | S | Indica a meta de clientes cadastrados. |
| 14 | MARGEMPREV | NUMBER(8,4) | S | Indica o percentual lucro meta. |
| 15 | EFETIVIDADEPREV | NUMBER(12,6) | S | Indica meta de efetividade. |
| 16 | EFETIVIDADEZEROPREV | NUMBER(12,6) | S | Indica meta sem efetividade (resto ou desconto de meta. |
| 17 | EFETIVIDADESEM1PREV | NUMBER(12,6) | S | Indica a meta de efetividade com 1 semana. |
| 18 | EFETIVIDADESEM2PREV | NUMBER(12,6) | S | Indica a efetividade 2 semanas(quinzenal). |
| 19 | EFETIVIDADESEM3PREV | NUMBER(12,6) | S | Indica a meta de efetividade com 3 semanas. |
| 20 | EFETIVIDADESEM4PREV | NUMBER(12,6) | S | Indica a efetividade 4 semanas(semanal). |
| 21 | PEDIDOSPREV | NUMBER(8,4) | S | Indica a meta de pedidos/dia. |
| 22 | QTMEDIAITENS | NUMBER(8,4) | S | Indica a meta de média de itens por pedido. |
| 23 | VLMEDIOITEM | NUMBER(12,6) | S | Indica a meta de valor médio de item. |
| 24 | VLMEDIOPEDIDO | NUMBER(12,2) | S | Indica a meta de valor médio de pedido. |
| 25 | VLMEDIOPEDIDOSDIA | NUMBER(12,2) | S | Indica a meta de valor médio de pedidos/dia. |
| 26 | DESCVLVENDAPREV | NUMBER(7,4) | S | Indica o % desconto na meta de valor de venda. |
| 27 | QTMETROSPREV | NUMBER(12,6) | S | Indica a quantidade de metros previstos. |
| 28 | PERVLDEVOLPREV | NUMBER(8,4) | S | Indica o percentual de devolução esperado. |
| 29 | PERINADIMPPREV | NUMBER(8,4) | S | Indica o percentual de inadimplência. |
| 30 | VLMINVENDAPOS | NUMBER(10,2) | S | Indica o valor mínimode venda. |
| 31 | PRAZOMEDIOPREV | NUMBER(4) | S | Indica o prazo médio esperado. |
| 32 | PRECOMEDIOPREV | NUMBER(18,6) | S | Indica o preço médio esperado. |
| 33 | PRECOMEDIOKGPREV | NUMBER(18,6) | S | Indica o preço médio/Kg esperado. |
| 34 | QTVENDAINVESTPREV | NUMBER(20,6) | S | Indica a meta de quant. venda por cliente. |
| 35 | VLVENDAINVESTPREV | NUMBER(14,2) | S | Indica a meta de valor venda por cliente investimento. |
| 36 | QTVENDAFOCOPREV | NUMBER(20,6) | S | Indica a meta de quant. venda por cliente foco. |
| 37 | VLVENDAFOCOPREV | NUMBER(14,2) | S | Indica a meta de valor venda por cliente foco. |
| 38 | PERQTDEVOLPREV | NUMBER(8,4) | S | Indica o percentual de quant. devolvida esperado. |
| 39 | MEDIAITENSCLIPREV | NUMBER(6,2) | S | Indica a média de itens/cliente esperada. |
| 40 | EFETIVIDADEDIARIA | NUMBER(12,6) | S | Indica a freqüência de visita diária esperada. |
| 41 | EFETIVIDADEBISSEMANAL | NUMBER(12,6) | S | Indica a freqüência de visita bissemanal esperada. |
| 42 | SEGURADA | VARCHAR2(1) | S | Indica carga segurada. |
| 43 | PERCLIPOSPREV | NUMBER(6,2) | S | Percentual de clientes positivados. |
| 44 | QTDCLIENTESATIVO | NUMBER(12,4) | S | Qtd. De clientes ativos na data da inclusão. |
| 45 | AMORTIZACAOEFETIVIDADE | NUMBER(12,6) | S | Percentual lançado para amortizar carteira efetividade semanal. |
| 46 | PERVLDESCONTOPREV | NUMBER(8,4) | S | Percentual de desconto previsto. |
| 47 | MIXMEDIOPREV | NUMBER(10,4) | S | Quantidade do mix medio previsto. |
| 48 | ROTINALANC | NUMBER(6) | S | Rotina lançamento. |
| 49 | CODIGO2 | NUMBER(8) | N | Indica o segundo código da meta relacionada. |
| 50 | PERCRATEIO | NUMBER(8,4) | S | Percentual de rateio entre RCA. |
| 51 | CODMETAC | NUMBER(8) | S | Indica o valor da chave da tabela PCMETAC |
| 52 | EFETIVIDADESEM5PREV | NUMBER(12,6) | S | Indica a efetividade 5 semanas(semanal). |
| 53 | LITRAGEM | NUMBER(18,6) | S | Indica a meta de litragem vendida |

## Chaves e índices

- Constraint PCMETA_PK (P): CODFILIAL, CODIGO, CODIGO2, CODUSUR, TIPOMETA, DATA
- Índice PCMETA_IDX1: CODUSUR
- Índice PCMETA_IDX1_CODIGO: CODIGO
- Índice PCMETA_PK UNIQUE: CODFILIAL, CODIGO, CODIGO2, CODUSUR, TIPOMETA, DATA
