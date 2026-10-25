# RIBERDOCES.PCPEDIDO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 242 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMPED | NUMBER(6) | N | Numero do pedido. |
| 2 | DTEMISSAO | DATE | S | Data de emissao do pedido. |
| 3 | CODFORNEC | NUMBER(6) | N | Codigo do fornecedor. |
| 4 | VLTOTAL | NUMBER(12,2) | S | Valor total do pedido. |
| 5 | DTPREVENT | DATE | S | Data de previsao de entrada. |
| 6 | OBS | VARCHAR2(100) | S | Observacao linha 1. |
| 7 | CODFILIAL | VARCHAR2(2) | S | Codigo da filial. |
| 8 | TRANSPORTE | VARCHAR2(28) | S | Nome da transportadora. |
| 9 | FRETE | VARCHAR2(1) | S | Tipo de frete. |
| 10 | VLENTREGUE | NUMBER(12,2) | S | Valor entregue. |
| 11 | OBS2 | VARCHAR2(100) | S | Observacao linha 2. |
| 12 | OBS3 | VARCHAR2(100) | S | Observacao linha 3. |
| 13 | VLVERBA | NUMBER(14,2) | S | Valor da verba. |
| 14 | CODCOMPRADOR | NUMBER(8) | S | Codigo do comprador. |
| 15 | ENTANTECIP | VARCHAR2(1) | S | Permite entrada antes da data de previsao de entrada. |
| 16 | TIPOBONIFIC | VARCHAR2(2) | S | Tipo do pedido: Normal, bonificado, consignado, consumo proprio. |
| 17 | VLPRODUTO | NUMBER(14,2) | S | Valor do produto. |
| 18 | REBAIXARPR | VARCHAR2(1) | S | Se deve rebaixar o preco. |
| 19 | CODCONTA | NUMBER(10) | S | Codigo da conta. |
| 20 | NUMBONIFIC | NUMBER(8) | S | Não utilizado. |
| 21 | CODPLPAG | NUMBER(6) | S | Codigo do plano de pagamento rotina 256. |
| 22 | TIPOVENC | VARCHAR2(1) | S | Tipo de vencimento (Faturamento, Embarque, Entrega, Emissao Pedido, Dia Fixo, Bonificado). |
| 23 | DTFATUR | DATE | S | Data do faturamento. |
| 24 | DTEMBARQUE | DATE | S | Data de embarque |
| 25 | VLFRETE | NUMBER(12,2) | S | Valor do frete CIF. |
| 26 | PRAZO1 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 27 | PRAZO2 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 28 | PRAZO3 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 29 | PRAZO4 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 30 | PRAZO5 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 31 | PRAZO6 | NUMBER(4) | S | Nao utilizado(Descontinuado). |
| 32 | OBS4 | VARCHAR2(100) | S | Observacao linha 4. |
| 33 | OBS5 | VARCHAR2(100) | S | Observacao linha 5. |
| 34 | OBS6 | VARCHAR2(100) | S | Observacao linha 6. |
| 35 | OBS7 | VARCHAR2(100) | S | Observacao linha 7. |
| 36 | NUMVERBA | NUMBER(8) | S | Numero da verba em Mercadoria. |
| 37 | NUMVERBABONIFIC | NUMBER(8) | S | Numero da verba em Mercadoria pedido bonificado. |
| 38 | CALCIPICOMDESC | VARCHAR2(1) | S | Considera desconto para calculo do IPI. |
| 39 | NUMVERBADIN | NUMBER(8) | S | Numero da verba em dinheiro. |
| 40 | NUMNEGOCIACAO | NUMBER(12) | S | Numero da negociacao. |
| 41 | PALETIZADO | VARCHAR2(1) | S | Tipo da carga paletizada. |
| 42 | CALCIPICOMFRETENF | VARCHAR2(1) | S | Considera frete CIF para calculo do IPI. |
| 43 | TIPOEMBALAGEMPEDIDO | VARCHAR2(1) | S | Tipo de embalagem: Master ou Venda. |
| 44 | APLICAPERCFRETECAB | VARCHAR2(1) | S |  |
| 45 | PERCFRETE | NUMBER(7,4) | S | Porcentagem do frete CIF no pedido. |
| 46 | DEDUZIRSUFRAMACALCCREDICM | VARCHAR2(1) | S | Deduzir SUFRAMA credito ICMS. |
| 47 | CODFUNCLIBERA | NUMBER(8) | S | Codigo do funcionario que liberou o pedido (rotina 215). |
| 48 | DTLIBERA | DATE | S | Data da liberacao do pedido utilizado na rotina 215. |
| 49 | HORALIBERA | VARCHAR2(2) | S | Hora da liberacao do pedido utilizado na rotina 215. |
| 50 | MINUTOLIBERA | VARCHAR2(2) | S | Minuto da liberacao utilizado na rotina 215. |
| 51 | DEDUZIRSUFRAMACALCCREDPIS | VARCHAR2(1) | S | Deduzir SUFRAMA credito PIS/COFINS. |
| 52 | UTILIZAOUTRASDESPCALCICMS | VARCHAR2(1) | S | Considera Outras despesas e Seguro para calculo do ICMS. |
| 53 | CALCSUFRAMASOBREPLIQUIDO | VARCHAR2(1) | S | Deduzir descontos para calcular SUFRAMA. |
| 54 | DEDUZIRSUFRAMABCSTALIQ1 | VARCHAR2(1) | S | Deduzir SUFRAMA BC ST aliquota interna. |
| 55 | ALTERAPERCPRESUMIDO | VARCHAR2(1) | S |  |
| 56 | SOMAFRETEFOBNOSTFORANF | VARCHAR2(1) | S | Soma frete FOB no ST guia. |
| 57 | CALCULAPISCOFINSCOMIPI | VARCHAR2(1) | S | Incluir IPI para calcular PIS/COFINS. |
| 58 | DATAENVIOEDI | DATE | S | Data de envio EDI (rotina 2425). |
| 59 | CODFUNCENVIOEDI | NUMBER(8) | S | Codigo do funcionario que enviou arquivo EDI (rotina 2425). |
| 60 | INCPRODUTOMAISDEUMAVEZ | VARCHAR2(1) | S | Incluir o mesmo produto mais de uma vez. |
| 61 | APLICAPERCDESCCOMCAB | VARCHAR2(1) | S |  |
| 62 | CODFUNCINCLUSAO | NUMBER(8) | S | Codigo do funcionario da inclusao. |
| 63 | CODFUNCALTERACAO | NUMBER(8) | S | Codigo do funcionario da alteracao. |
| 64 | UTILIZAFRETECALCICMS | VARCHAR2(1) | S | Considera frete CIF para calculo do ICMS. |
| 65 | NFIMPORTACAO | CHAR(1) | S | Numero da nota fiscal de importacao (rotina 3002). |
| 66 | NUMTRANSVENDA | NUMBER(10) | S | Numero de transacao de venda. |
| 67 | UTILIZAOUTRASDESPCALCIPI | VARCHAR2(1) | S | Utilizado na rotina 209 e 1301, caso o valor seja igual a "S" será utilizado o valor de outras despesas para cálculo de IPI. |
| 68 | VLVERBAOUTRAS | NUMBER(12,2) | S | Valor de outras verbas. |
| 69 | CODFUNC | NUMBER(8) | S | Codigo do funcionario. |
| 70 | NUMVERBAOUTRAS | NUMBER(8) | S | Numero de outras verbas. |
| 71 | UTILIZADESCCALCST | VARCHAR2(1) | S | Esse campo será utilizada pelas rotinas de pedido de compra, fornecedor e entrada de mercadoria quando esse campo possuir o valor igual a [S] os descontos comerciais serão utilizados Substituição Tributária (ST). |
| 72 | IMPORTACAO | VARCHAR2(1) | S | Se o pedido foi importado (modulo 25). |
| 73 | UTILIZADESCCALCICMS | VARCHAR2(1) | S | Esse campo será utilizada pelas rotinas de pedido de compra, fornecedor e entrada de mercadoria quando esse campo possuir o valor igual a [S] os descontos comerciais serão utilizados ICMS. |
| 74 | EXPORTADO | VARCHAR2(1) | S | Se o pedido ja foi exportado (modulo 25). |
| 75 | CONSIPICALCBASEST | VARCHAR2(1) | S | Considera IPI para calculo da base ST. |
| 76 | CONSIPICALCBASECREPRES | VARCHAR2(1) | S | Considera IPI para calculo da base de credito presumido. |
| 77 | INCLUIRFRETECUSTOCPAGAR | VARCHAR2(1) | S |  |
| 78 | CODMOEDA | NUMBER(6) | S | Codigo da moeda utilizada para digitar pedido de compra.  |
| 79 | COTACAO | NUMBER(18,6) | S | Valor da cotação escolhida para digitação do pedido de compra.  |
| 80 | EREABASTECIMENTO | VARCHAR2(1) | S | Indica se é de reabastecimento.  |
| 81 | CODPARCELA | NUMBER(6) | S | Código da Parcela.  |
| 82 | TIPODESCARGA | VARCHAR2(1) | S | Tipo do pedido (descarga). Se informado, será validado na entrada de mercadoria.  |
| 83 | ROTINALANC | VARCHAR2(48) | S | . \|Campo do tipo caracter, de tamanho 48.  |
| 84 | FUNCLANC | VARCHAR2(30) | S | . \|Campo do tipo caracter, de tamanho 30.  |
| 85 | DATALANC | DATE | S | . \|Campo do tipo data, com valor default SYSDATE  .  |
| 86 | CALCIPICOMFRETEFOB | VARCHAR2(1) | S | Considera frete FOB para calculo do IPI. |
| 87 | DTCOTACAO | DATE | S | Campo para armazenar a data da cotação do pedido de compra. |
| 88 | TIPOFRETEFOB | VARCHAR2(2) | S | Indica o tipo do frete FOB.  |
| 89 | UTILIZAICMSDIFZERADO | VARCHAR2(1) | S | Indica se utiliza ICMS diferido com alíquota zero. |
| 90 | CALCSTGUIAALIQEXT | VARCHAR2(1) | S | Calcula ST guia alíquota externa. |
| 91 | CALCPISCOFINSBASERED | VARCHAR2(1) | S | Indica o código da transportadora. |
| 92 | CODFORNECFRETE | NUMBER(8) | S | Indica o código da transportadora. |
| 93 | CALCULARIPIPESOLIQ | VARCHAR2(1) | S | Calcula IPI com base no peso líquido. |
| 94 | CODFORNECSTGUIA | NUMBER(6) | S | Indica o código fornecedor ST guia. |
| 95 | ISENTOST | VARCHAR2(1) | S | Indica o fornecedor isento de ST. |
| 96 | CODFORNECOUTRASDESP | NUMBER(6) | S | Indica o código do fornecedor de outras despesas. |
| 97 | AUTOMATICO | VARCHAR2(1) | S | Indica se o pedido e automatico. |
| 98 | USAPERCICMSNAALIQEXTST | VARCHAR2(1) | S | Considera frete FOB para o calculo do ICMS antecipado. |
| 99 | CONSFRETEFOBCALCICMSANTECIP | VARCHAR2(1) | S | Considera frete FOB para o calculo do ICMS antecipado. |
| 100 | CODFILIALPRINC | VARCHAR2(2) | S | Número da filial principal pedido de compra. |
| 101 | NUMPEDPRINC | NUMBER(10) | S | Número do pedido de compra principal. |
| 102 | CODPARCELASTGUIA | NUMBER(6) | S | Indica o código parcela ST guia. |
| 103 | CONSSTNFPISCOFINS | VARCHAR2(1) | S | Considerar o ST Nota fiscal no PIS/COFINS. |
| 104 | NUMTRANSENTN | NUMBER(10) | S | Número de transação tipo N |
| 105 | UTILIZAOUTDESPCALCSUFRAMA | VARCHAR2(1) | S | Utiliza outras despesa para calculo do desconto SUFRAMA. |
| 106 | CONFIRMADO | VARCHAR2(1) | S | Define se o pedido foi confirmado pelo fornecedor após a digitação web dos preços da cotação utilizada. |
| 107 | NUMPROFORMA | VARCHAR2(20) | S | Número Proforma. |
| 108 | CODVIA | NUMBER(10) | S | Código da via de transporte. |
| 109 | CODFORNECRCA | NUMBER(6) | S | Código do representante comercial;. |
| 110 | CODFORNECIMPORT | NUMBER(6) | S | Código do Fornecedor do tipo importador. |
| 111 | CODFORNECFABRIC | NUMBER(6) | S | Código do Fornecedor do tipo fabricante. |
| 112 | CODFORNECPROD | NUMBER(6) | S | Código do Fornecedor do tipo produtor. |
| 113 | CODPAISORIGEM | NUMBER(6) | S | Código do país de origem. |
| 114 | CODPAISPROC | NUMBER(6) | S | Código do país de procedência. |
| 115 | CODPAISAQUISICAO | NUMBER(6) | S | Código do país de aquisição. |
| 116 | INCOTERM | VARCHAR2(5) | S | Sigla Incoterm. |
| 117 | ORCAMENTO | VARCHAR2(2) | S | valores possíveis "S" quando for um Pedido máster de importação; "N" quando não for; VALOR DEFAULT "N". |
| 118 | IDPEDIDOMASTER | VARCHAR2(20) | S | Identificação Pedido Máster. |
| 119 | PERCADIANTAMENTO | NUMBER(18,10) | S | Percentual sobre o valor total da compra que vai ser pago de modo antecipado. |
| 120 | PERCAVISTA | NUMBER(18,10) | S | Percentual sobre o valor total da compra que vai ser pago de modo à vista. |
| 121 | PERCAPRAZO | NUMBER(18,10) | S | Percentual sobre o valor total da compra que vai ser pago de modo a prazo. |
| 122 | DTVENC | DATE | S | Data de vencimento da parte que vai ser paga a vista. |
| 123 | CODTRANSITARIO | NUMBER(6) | S | Código do fornecedor do tipo Transitário. |
| 124 | CODDESPACHANTE | NUMBER(6) | S | Código do fornecedor do tipo Despachante. |
| 125 | CODSEGURADORA | NUMBER(6) | S | Código do fornecedor do tipo Seguradora. |
| 126 | DTNACIONALIZACAO | DATE | S | Data de nacionalização efetiva da mercadoria. |
| 127 | DTPREVEMBARQUE | DATE | S | Data de previsão de embarque |
| 128 | DTPREVCHEGADA | DATE | S | Data de previsão chegada |
| 129 | DTCHEGADA | DATE | S | Data de chegada |
| 130 | LOCALEMBARQUE | VARCHAR2(200) | S | Local de embarque |
| 131 | DTCOMPROVIMPORT | DATE | S | Data de emissão do comprovante de importação |
| 132 | DTEMISSAOCONHECEMBARQ | DATE | S | Data de emissão do conhecimento de embarque |
| 133 | NUMCONHECEMBARQ | VARCHAR2(20) | S | Número do conhecimento de embarque |
| 134 | DTLEMISSAOLICENCAIMPORT | DATE | S | Data de emissão da licença de importação |
| 135 | NUMLICENCAIMPORT | VARCHAR2(20) | S | Número da licença de importação |
| 136 | DTVENCLICENCAIMPORT | DATE | S | Data de vencimento da licença de importação |
| 137 | DTEMISSAOINVOCE | DATE | S | Data de emissão da fatura comercial. |
| 138 | NUMINVOCE | VARCHAR2(20) | S | Número da fatura comercial |
| 139 | CANALCONFERENCIA | NUMBER(2) | S | Canal de conferência da receita |
| 140 | DTEMISSAODOCIMPORT | DATE | S | Data do documento de importação |
| 141 | NUMDOCIMPORT | VARCHAR2(20) | S | Número do documento de importação |
| 142 | CODPORTOCHEGADA | NUMBER(10) | S | Código do porto de chegada |
| 143 | DTPREVNACIONALIZACAO | DATE | S | Data prevista para nacionalização |
| 144 | CODPORTONACIONALIZACAO | NUMBER(10) | S | Código do porto de nacionalização |
| 145 | CODTERMINALPORTUARIO | NUMBER(10) | S | Código do terminal portuário |
| 146 | DTEMISSAOPASTA | DATE | S | Data de emissão da pasta/ processo |
| 147 | NUMPASTA | VARCHAR2(20) | S | Número da pasta/processo |
| 148 | IDCONTROLEEMBARQUE | VARCHAR2(20) | S | Identificação do controle de embarque |
| 149 | TIPOVERBAOUTRAS | VARCHAR2(1) | S | Tipo de pagamento da verba outras |
| 150 | CONSMAIORICMSVLPAUTA | VARCHAR2(1) | S | Calcula ICMS sobre Vl.Pauta maior valor(Importação). |
| 151 | DEDFRETECIFCREDPRESICMS | VARCHAR2(1) | S | Considera frete FOB no calculo da base PIS/COFINS |
| 152 | DTPREVPRIMEMBARQUE | DATE | S | Data de previsão do primeiro embarque. |
| 153 | DTFECHACE | DATE | S | Data de fechamento do controle de embarque. |
| 154 | OCORRENCIAS | CLOB | S | Texto geral registros de ocorrências |
| 155 | MEIOTRANSPORTE | VARCHAR2(100) | S | Nome do meio de transporte |
| 156 | DTPREVENTRADAESTOQUE | DATE | S | Data de previsão de entrada no estoque.. |
| 157 | DTENTRADAESTOQUE | DATE | S | Data de entrada efetiva no estoque. |
| 158 | CODPLPAGAVISTA | NUMBER(6) | S | Código do plano de pagamento para modalidade à vista. |
| 159 | CODPLPAGADIANT | NUMBER(6) | S | Código do plano de pagamento para modalidade adiantamento. |
| 160 | USACREDPISCOFINSENTBONIFIC | VARCHAR2(1) | S | Apresenta se o pedido utiliza crédito PIS/COFINS em entrada bonificada |
| 161 | USAOUTRASDESPSEGUROPISCOFINS | VARCHAR2(1) | S | Somar outras despesas e seguros na base PIS/COFINS |
| 162 | CALCDESPFINPRAZOPGTO | VARCHAR2(1) | S | Calcular vendor por prazo de pagamento |
| 163 | CODMOEDACUSTO | NUMBER(6) | S | Código Moeda Estrangeira |
| 164 | COTACAOCUSTO | NUMBER(16,6) | S | Cotação Moeda Estrangeira |
| 165 | DTCOTACAOCUSTO | DATE | S | Data Cotação Moeda Estrangeira |
| 166 | CODPORTOORIGEM | NUMBER(10) | S | Porto de origem de embarque |
| 167 | CODCLIDEST | NUMBER(9) | S | Cliente destino do pedido |
| 168 | CODFORNECNOTIF | NUMBER(6) | S | Fornecedor a ser notificado do embarque |
| 169 | DOCIMPORTACAO | VARCHAR2(500) | S | Lista de documentos solicitados para embarque |
| 170 | MODALIDADEPGTO | VARCHAR2(100) | S | Modalidade de pagamento |
| 171 | GERACPPREV | VARCHAR2(1) | S | DEVE GERAR CONTAS A PAGAR PREVISTO |
| 172 | UTILIZAIPICALCICMS | VARCHAR2(1) | S | Utiliza IPI no calculo de ICMS |
| 173 | COTACAOAVISTA | NUMBER(14,6) | S | Cotação da moeda estrangeira para pagamento a vista |
| 174 | CODMOEDAAVISTA | NUMBER(6) | S | Código de moeda estrangeira para pagamento a vista |
| 175 | DTCOTACAOAVISTA | DATE | S | Data da cotação da moeda estrangeira |
| 176 | COTACAOPREVISTA | NUMBER(14,6) | S | Cotação prevista no momento da entrada |
| 177 | CODDESPACHANTE2 | NUMBER(6) | S | Códido de fornecedor despachante Courier |
| 178 | DTEXPEDI | DATE | S | Data de Exportação para EDI |
| 179 | MODALIDADEPGTOADIANT | VARCHAR2(100) | S | Descrição da modalidade de pagamento adiantamento |
| 180 | MODALIDADEPGTOPRAZO | VARCHAR2(100) | S | Descrição da modalidade de pagamento a prazo |
| 181 | NUMBONUS | NUMBER(6) | S | Numero do Bonus |
| 182 | AVALIACAOVENDA | VARCHAR2(2) | S | Período de Avaliação de venda para calculo da sugestão de compra |
| 183 | DIASAVALIACAOVENDA | NUMBER(4) | S | Período de Avaliação de venda para calculo da sugestão de compra |
| 184 | USADRAWBACK | VARCHAR2(1) | S | PIS/COFINS DrawBack |
| 185 | CONSSTGUIAPISCOFINS | VARCHAR2(1) | S | Considera ST Guia Base do PISCOFINS. |
| 186 | CODEDITAL | NUMBER(9) | S | Código do edital |
| 187 | GERAPEDBONIFIC | VARCHAR2(1) | S | Gera Pedido Bonificado Automático |
| 188 | DTLI | DATE | S | Data LI |
| 189 | DTDIFLI | DATE | S | Data diferimento da LI |
| 190 | DTENVIODESPACHANTE | DATE | S | Data de envio de documentos ao despachante |
| 191 | DTZONASECUNDARIA | DATE | S | Data da chegada na zona secundária |
| 192 | DTVISTORIAORGAOANUENTE | DATE | S | Data da vistoria do órgão anuente |
| 193 | DTREGISTROCANAL | DATE | S | Data de registro do canal |
| 194 | MULTIMPLOUNIDADEMASTER | VARCHAR2(1) | S | Calcula Sugestão em multiplos da Unidade Master |
| 195 | CONSIDERAESTPENDSUGCOMPRA | VARCHAR2(1) | S | Considera estoque pendente de venda na sugestão de compra |
| 196 | IMPORTADOXML | VARCHAR2(1) | S | Importado do xml |
| 197 | UTILIZAOUTDESPNFBASEST | VARCHAR2(1) | S | Considerar o valor Outras Despesas NF + Seguro na base de calculo ST |
| 198 | CODDOCIMP | NUMBER(6) | S | Código do documento de importação - Status |
| 199 | CODPARCELADIFALIQ | NUMBER(6) | S | Código parcela Diferencial Alíquota |
| 200 | CODPARCELAFRETEFOB | NUMBER(6) | S | Código parcela Frete FOB |
| 201 | CODPARCELAOUTDESP | NUMBER(6) | S | Código parcela Outras Desp.Fora NF |
| 202 | CODPARCICMSANTECIP | NUMBER(6) | S | Código parcela ICMS Antecipado |
| 203 | CONSFRETEFOBBASESTNF | VARCHAR2(1) | S | Considera Frete FOB base ST NF |
| 204 | GERACPPREVDESPFORANF | VARCHAR2(1) | S | Gera contas a pagar previsto para as despesas fora da NF. |
| 205 | DTFECHACAMBIO | DATE | S | Data fechamento do câmbio na Importação |
| 206 | PERCVENDOR | NUMBER(12,4) | S | Percentual de vendor |
| 207 | CALCFUNRURAL | VARCHAR2(1) | S | Indica se caluca o FUNRURAL de forma separada no contas a pagar. |
| 208 | CODPARCELAFUNRURAL | NUMBER(6) | S | Indica o código de parcela do funrural. |
| 209 | CODFORNECFUNRURAL | NUMBER(6) | S | Código fornecedor funrurall |
| 210 | CONSCAPATAZIAICMS | VARCHAR2(1) | S | Considera capatazia cálculo ICMS (importação) |
| 211 | DEDUZIRSUFRAMAALIQEXT | VARCHAR2(1) | S | Deduzir Suframa Créd. ICMS Aliq. Ext. ST |
| 212 | TIPOUNIDADEMASTER | VARCHAR2(1) | S | Tipo da unidade master (por fornecedor ou produto) informada pelo usuario |
| 213 | OBSBLOQUEIO | VARCHAR2(500) | S | Observação com o motivo do bloqueio |
| 214 | CONSIPICALCICMSANT | VARCHAR2(1) | S | Considera IPI no cálculo do ICMS antecipado |
| 215 | ORIGEMESTOQUEMASTER | VARCHAR2(1) | S |  |
| 216 | CODPARCELAORIGEM | NUMBER(6) | S | Cód. Parcela gravado origem |
| 217 | TIPOVENCORIGEM | VARCHAR2(1) | S | Tipo vencimento origem |
| 218 | CONSIDERADTVENCRET | VARCHAR2(1) | S | Considera data de vencimento retroativa |
| 219 | CONSIDERADTVENCRETOR | VARCHAR2(1) | S | Considera data de vencimento retroativa origem |
| 220 | NUMPEDORIG | NUMBER(8) | S | Número do pedido de origem |
| 221 | USAMULTUNIDITEM | VARCHAR2(1) | S | Usa multimplas unidades (Master ou Venda) por item nopedido de compra |
| 222 | DEDUZIRSUFRAMADIFALIQ | VARCHAR2(1) | S | Deduzir Suframa Diferença de Alíquotas |
| 223 | DEDUZIRICMSBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do ICMS da base de cálculo do PIS/COFINS |
| 224 | DEDUZIRFCPICMSDOFCPST | VARCHAR2(1) | S | Deduzir o FCP ICMS do valor do FCP ST |
| 225 | CONSIDERARPAUTAZEROSTGUIA | VARCHAR2(1) | S | Considerar pauta zero para cálculo do ST Guia |
| 226 | CALCCREDICMSBASEREDUZIDA | VARCHAR2(1) | S | Utiliza crédito de ICMS sobre base reduzida |
| 227 | CHAVENFE | VARCHAR2(44) | S | Chave NFe do XML que gerou o pedido de compra |
| 228 | CALCCREDFUNCEPCONT | VARCHAR2(1) | S | Calcular credito FUNCEP (FCP ICMS) Contabil |
| 229 | CALCCREDFUNCEP | VARCHAR2(1) | S | Calcular credito FUNCEP (FCP ICMS) Gerencial |
| 230 | ZERARALIQICMSPRESUMIDO | VARCHAR2(1) | S | Zerar Alíquota ICMS PRESUMIDO para produto Bonificado |
| 231 | DEDUZIRIPIBASEDIFALIQ | VARCHAR2(1) | S | Deduzir IPI da base Diferencial Alíquota |
| 232 | CONSFRETEFOBBASEDIFALIQ | VARCHAR2(1) | S | Deduzir IPI da base Diferencial Alíquota |
| 233 | DEDUZFUNCEPBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do FUNCEP da Base de Cálculo do Pis/Cofins |
| 234 | APLICAVERBAREBCUSTO | VARCHAR2(1) | S | Indica se utiliza aplicação de verba para rebaixa de custo |
| 235 | ALIMENTAFUNDOMULTIFILIAL | VARCHAR2(1) | S | Momento de que o fundo multifilial sera alimentado |
| 236 | CODFORNECDIFALIQ | NUMBER(6) | S | Código do fornecedor das despesas com diferença de alíquotas |
| 237 | CODFORNECICMSANT | NUMBER(6) | S | Código do fornecedor do ICMS de antecipação |
| 238 | DESCREDICMSFCP3010 | VARCHAR2(1) | S | Desconsiderar %Red ICMS para cálculo do FCP ICMS na importação 3010 |
| 239 | DEDUZIRCAPATAZIABASEII | VARCHAR2(1) | S | DEDUZIR CAPATAZIA DA BASE II |
| 240 | DEDUZIRCAPATAZIABASEPISCOFINS | VARCHAR2(1) | S | DEDUZIR CAPATAZIA BASE PIS/COFINS |
| 241 | DEDUZIRICMSDIFBASEICMSATECIP | VARCHAR2(1) | S | Deduzir ICMS Diferido da Base do ICMS Antecipado |
| 242 | CROSSDOCKING | VARCHAR2(1) | S | Indica que é um pedido com entrada para crossdocking |

## Chaves e índices

- Constraint PCPEDIDO_PK (P): NUMPED
- Índice PCPEDIDO_IDX1: DTEMISSAO
- Índice PCPEDIDO_IDX10: NUMVERBADIN
- Índice PCPEDIDO_IDX11: NUMVERBAOUTRAS
- Índice PCPEDIDO_IDX2: CODFORNEC
- Índice PCPEDIDO_IDX3: CODFILIAL
- Índice PCPEDIDO_IDX4: NUMTRANSVENDA
- Índice PCPEDIDO_IDX5: IDCONTROLEEMBARQUE
- Índice PCPEDIDO_IDX6: CHAVENFE
- Índice PCPEDIDO_IDX7: NUMVERBABONIFIC
- Índice PCPEDIDO_IDX8: NUMBONIFIC
- Índice PCPEDIDO_IDX9: NUMVERBA
- Índice PCPEDIDO_PK UNIQUE: NUMPED
