# RIBERDOCES.PCPEDI — dicionário

Gerado por `gateway inventory` em 2026-10-25. 476 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMPED | NUMBER(10) | N |  |
| 2 | DATA | DATE | N |  |
| 3 | CODCLI | NUMBER(9) | N |  |
| 4 | CODPROD | NUMBER(6) | N |  |
| 5 | CODUSUR | NUMBER(4) | N |  |
| 6 | QT | NUMBER(20,6) | N |  |
| 7 | PVENDA | NUMBER(18,6) | N |  |
| 8 | PTABELA | NUMBER(18,6) | N |  |
| 9 | NUMCAR | NUMBER(8) | S |  |
| 10 | POSICAO | VARCHAR2(2) | N |  |
| 11 | ST | NUMBER(18,6) | N | Valor do ST. |
| 12 | VLCUSTOFIN | NUMBER(18,6) | N |  |
| 13 | VLCUSTOREAL | NUMBER(18,6) | N |  |
| 14 | PERCOM | NUMBER(8,4) | N |  |
| 15 | PERDESC | NUMBER(18,6) | N |  |
| 16 | QTFALTA | NUMBER(20,6) | S |  |
| 17 | NUMSEQ | NUMBER(20) | N |  |
| 18 | TIPOPESO | VARCHAR2(1) | S |  |
| 19 | PERCOMTAB | NUMBER(8,4) | S |  |
| 20 | PERDESCTAB | NUMBER(8,4) | S |  |
| 21 | CODMOTNAOCOMPRA | NUMBER(6) | S | Código do motivo de não compra. |
| 22 | VLDESCCUSTOCMV | NUMBER(12,4) | S |  |
| 23 | QTSEPARADA | NUMBER(20,6) | S |  |
| 24 | QTVENDAEMB | NUMBER(12,3) | S |  |
| 25 | PVENDAEMB | NUMBER(18,6) | S |  |
| 26 | VLOUTROS | NUMBER(16,3) | S |  |
| 27 | QTEMBALAGEM | NUMBER(12,3) | S |  |
| 28 | PVENDAEMBALAGEM | NUMBER(18,6) | S |  |
| 29 | CODAUXILIAR | NUMBER(20) | S |  |
| 30 | VLCUSTOREP | NUMBER(18,6) | S |  |
| 31 | VLCUSTOCONT | NUMBER(18,6) | S |  |
| 32 | CODCERTIFIC | NUMBER(8) | S | Código do certificado. |
| 33 | PVENDABASE | NUMBER(18,6) | S |  |
| 34 | NOMECONCORRENTE | VARCHAR2(60) | S | Nome do concorrente. |
| 35 | PRECO | NUMBER(10,2) | S | Preço do concorrente. |
| 36 | PRAZO | VARCHAR2(40) | S | Prazo do concorrente. |
| 37 | QTNAOCOMPRA | NUMBER(18,6) | S | Quantidade não comprada. |
| 38 | CODFILIALRETIRA | VARCHAR2(2) | S |  |
| 39 | NUMTIRA | NUMBER(10) | S | Indica o o número de tira. |
| 40 | CODFUNCSEP | NUMBER(8) | S |  |
| 41 | VLDESCSUFRAMA | NUMBER(18,6) | S |  |
| 42 | NUMLOTE | VARCHAR2(15) | S | Indica o número do lote ao qual o produto pertence. |
| 43 | VLDESCREPASSE | NUMBER(12,12) | S | Valor desconto repasse. |
| 44 | REFCOR | VARCHAR2(20) | S | Referencia cor do produto. |
| 45 | CODFUNCCONF | NUMBER(8) | S |  |
| 46 | DATACONF | DATE | S |  |
| 47 | VLDESCICMISENCAO | NUMBER(18,6) | S | Valor descontro ICMS isento. |
| 48 | QTORIGINAL | NUMBER(20,6) | S |  |
| 49 | VLDESCFORNEC | NUMBER(18,6) | S |  |
| 50 | VLFRETE | NUMBER(18,6) | S |  |
| 51 | VLIPI | NUMBER(18,6) | S |  |
| 52 | QTORIG | NUMBER(20,6) | S |  |
| 53 | QTSEPARARUN | NUMBER(18,6) | S |  |
| 54 | QTSEPARARCX | NUMBER(18,6) | S |  |
| 55 | CODST | NUMBER(4) | S |  |
| 56 | VLDESCFIN | NUMBER(18,6) | S |  |
| 57 | PERCIPI | NUMBER(12,4) | S |  |
| 58 | IVA | NUMBER(8,4) | S |  |
| 59 | ALIQICMS1 | NUMBER(8,4) | S |  |
| 60 | ALIQICMS2 | NUMBER(8,4) | S |  |
| 61 | PAUTA | NUMBER(8,4) | S | Valor de pauta. |
| 62 | PERCBASERED | NUMBER(8,4) | S |  |
| 63 | VLDESCCOM | NUMBER(18,6) | S |  |
| 64 | PERDESCCOM | NUMBER(12,4) | S |  |
| 65 | PERDESCFIN | NUMBER(12,4) | S |  |
| 66 | VLBONIFIC | NUMBER(18,6) | S |  |
| 67 | PERBONIFIC | NUMBER(12,4) | S |  |
| 68 | PORIGINAL | NUMBER(18,6) | S |  |
| 69 | VLREBAIXACMV | NUMBER(18,6) | S | Valor verba do fornecedor para rebaixa do CMV. |
| 70 | NUMAPLIC | NUMBER(8) | S |  |
| 71 | PERFRETECMV | NUMBER(8,4) | S |  |
| 72 | VLDESCRODAPE | NUMBER(18,6) | S |  |
| 73 | STCLIENTEGNRE | NUMBER(18,6) | S | Valor ST para cliente com GNRE. |
| 74 | IMPRIME | VARCHAR2(1) | S | Indica se imprime o item ou não. |
| 75 | COMPLEMENTO | VARCHAR2(40) | S |  |
| 76 | CUSTOFINEST | NUMBER(18,6) | S | Valor custo financeiro estoque. |
| 77 | PERCBASEREDSTFONTE | NUMBER(8,4) | S | Percentual base reduzida ST cliente fonte. |
| 78 | PERCBASEREDST | NUMBER(8,4) | S | Percentual da base reduzida ST. |
| 79 | PERDESCCUSTO | NUMBER(8,4) | S | Percentual desconto custo. |
| 80 | CODICMTAB | NUMBER(8,4) | S | Código ICM tributação tabela. |
| 81 | TXVENDA | NUMBER(8,6) | S |  |
| 82 | PERCOM2 | NUMBER(8,4) | S |  |
| 83 | PERCOM3 | NUMBER(8,4) | S |  |
| 84 | PERCISS | NUMBER(8,4) | S |  |
| 85 | VLISS | NUMBER(18,6) | S |  |
| 86 | NUMTRANSWMS | NUMBER(10) | S |  |
| 87 | CODPROMOCAO | VARCHAR2(10) | S |  |
| 88 | PRAZOMEDIO | NUMBER(4) | S |  |
| 89 | LOCALIZACAO | VARCHAR2(40) | S |  |
| 90 | VLREPASSE | NUMBER(18,6) | S |  |
| 91 | PBONIFIC | NUMBER(18,6) | S |  |
| 92 | PERCVENDA | NUMBER(5,2) | S |  |
| 93 | VLDESCPISSUFRAMA | NUMBER(18,6) | S |  |
| 94 | CODDEGUSTACAO | NUMBER(10) | S |  |
| 95 | QTLOCALIZADA | NUMBER(8,2) | S |  |
| 96 | PERDESCFLEX | NUMBER(18,6) | S |  |
| 97 | VLDESCFLEX | NUMBER(18,6) | S |  |
| 98 | PERREDCOMISS | NUMBER(18,6) | S | Percentual de redução comissão. |
| 99 | VLREDCOMISS | NUMBER(18,6) | S | Valor da redução comissão. |
| 100 | TIPODESCAPLICADO | VARCHAR2(2) | S | Tipo desconto aplicado. |
| 101 | PBASERCA | NUMBER(18,6) | S | Preço base RCA. |
| 102 | PESOBRUTO | NUMBER(7,3) | S |  |
| 103 | NUMVERBAREBCMV | NUMBER(6) | S |  |
| 104 | CONDVENDA | NUMBER(5) | S | Indica a condição de venda.  |
| 105 | CODPLPAG | NUMBER(4) | S |  |
| 106 | EANCODPROD | NUMBER(14) | S |  |
| 107 | BRINDE | VARCHAR2(1) | S |  |
| 108 | PERCOMSUP | NUMBER(8,4) | S |  |
| 109 | PERREDCOMISSSUP | NUMBER(18,6) | S |  |
| 110 | VLREDCOMISSSUP | NUMBER(18,6) | S |  |
| 111 | BASEICST | NUMBER(18,6) | S |  |
| 112 | NUMOP | NUMBER(8) | S | Indica o numero da ordem de produção. |
| 113 | QTCX | NUMBER(14,6) | S |  |
| 114 | QTPECAS | NUMBER(14,6) | S |  |
| 115 | CODECF | VARCHAR2(4) | S |  |
| 116 | LETRACOMISS | VARCHAR2(2) | S | Código classificação tipo comissão. |
| 117 | VLACRESCRODAPE | NUMBER(18,6) | S | Valor acrescimo rodapé. |
| 118 | NUMCONFERENCIA | NUMBER(2) | S |  |
| 119 | PERDESCISENTOICMS | NUMBER(4,2) | S | Percentual desconto isento ICMS. |
| 120 | PERCOMPROF | NUMBER(6,2) | S | Percentual comissão profissional. |
| 121 | NUMCARAUX | NUMBER(8) | S |  |
| 122 | PVENDA1 | NUMBER(18,6) | S | Valor do preço a vista, para realização de cálculo de comissão sobre preço a vista.  |
| 123 | PERCAGREGADORST | NUMBER(8,4) | S |  |
| 124 | VLVERBACMVCLI | NUMBER(18,6) | S | Indica o valor da verba CMV calculado a partir do % de verba e aplicado sobre o preço de venda do produto. Este valor será adicionado ao CMV. |
| 125 | VLOUTRASDESP | NUMBER(18,6) | S | Valor de Outras Despesas Rateadas no Item.  |
| 126 | EXPORTADOSERVINT | VARCHAR2(1) | S |  |
| 127 | DTEXPORTACAOSERVINT | DATE | S |  |
| 128 | DTIMPORTACAOSERVPRINC | DATE | S |  |
| 129 | IMPORTADOSERVPRINC | VARCHAR2(1) | S |  |
| 130 | CODVASILHAME | NUMBER(6) | S | Código do Vasilhame associado ao Produto.  |
| 131 | QTAPANHA | NUMBER(20,6) | S | Usado na separação via Coletor de Dados na rotina 947.  |
| 132 | QTUNITCX | NUMBER(8,2) | S | Quantidade unitária por caixa. Rotina 316. Quantidade, Unidade Master.  |
| 133 | TRUNCARITEM | VARCHAR2(1) | S | Para correção de diferenças de truncamento/arredondamento no subtotal do item das impressoras fiscais.  |
| 134 | ABASTECIDO | VARCHAR2(1) | S |  |
| 135 | QTIMEDIATA | NUMBER(16,3) | S | Quantidade de Entrega Imediata. |
| 136 | CODFUNCLANC | NUMBER(8) | S | Código do Funcionário que Incluiu o Item. |
| 137 | ROTINALANC | NUMBER(6) | S | Código da Rotina que sofreu alteração. |
| 138 | DTLANC | DATE | S | Data da inclusão do Item no Pedido. |
| 139 | CODFUNCULTALTER | NUMBER(8) | S | Código do Funcionário que Alterou o Item. |
| 140 | ROTINAULTLALTER | NUMBER(6) | S | Código da Rotina que sofreu alteração. |
| 141 | DTULTLALTER | DATE | S | Data daAlteração do Item no Pedido. |
| 142 | CODFUNCALTERACAOOS | NUMBER(8) | S | Código do funcionário que alterou a OS. |
| 143 | DTALTERACAOOS | DATE | S | Data/Hora de alteração da OS . |
| 144 | QTPENDOS | NUMBER(10,4) | S | Quantidade pendente de itens da O.S. |
| 145 | VLDIFALIQUOTAS | NUMBER(18,6) | S | Valor da diferença de aliquotas. |
| 146 | BASEDIFALIQUOTAS | NUMBER(18,6) | S | Base da diferença entre alíquotas. |
| 147 | PERCDIFALIQUOTAS | NUMBER(8,4) | S | Percentual de diferença de tributação. |
| 148 | GERAGNRE_CNPJCLIENTE | VARCHAR2(1) | S | Campo para definir se a GNRE será paga pelo cliente. |
| 149 | PRODDESCRICAOCONTRATO | VARCHAR2(300) | S | Indica a descrição do produto no contrato. |
| 150 | NUMOS | NUMBER(15) | S | Número da O.S. |
| 151 | DTINICIALSEP | DATE | S | Inicio de separação de O.S. |
| 152 | DTFINALSEP | DATE | S | Fim de separação de O.S. |
| 153 | DATACONFFIM | DATE | S | Fim de conferência de O.S. |
| 154 | SITUACAOOS | NUMBER(2) | S | Situação da O.S. |
| 155 | NUMVIASOS | NUMBER(3) | S | Vias emitidas da O.S. |
| 156 | DTINICIALPEND | DATE | S | Inicio de pendência de O.S. |
| 157 | DTFINALPEND | DATE | S | Fim de pendência de O.S. |
| 158 | CODFUNCPEND | NUMBER(8) | S | Conferente de pendência da O.S. |
| 159 | DTLIBOS | DATE | S | Liberação da O.S. |
| 160 | CODFUNCLIBOS | NUMBER(8) | S | Responsável por liberar O.S. |
| 161 | LOCALIZACAOOS | VARCHAR2(60) | S | Localização de carrinhos da O.S. |
| 162 | NUMOSORIGEM | NUMBER(15) | S | Número de OS Original. |
| 163 | PERDESCPOLITICA | NUMBER(8,2) | S | Percentual de desconto aplicado em política de desconto por valor de pedido.  |
| 164 | PVENDAANTERIOR | NUMBER(18,6) | S | Preço anterior a aplicação de política de desconto por valor de pedido.  |
| 165 | TIPOENTREGA | VARCHAR2(2) | S | Tipo de entrega do produto.  |
| 166 | TVBONIF | NUMBER(2) | S | Indica que tipo de venda será o brinde. |
| 167 | CODIGOBRINDE | NUMBER(6) | S | Indica o código do brinde. |
| 168 | CODFUNCAJUSTEOS | NUMBER(8) | S | Indica o código do funcionário que efetuou ajuste de O.S. |
| 169 | DTAJUSTEOS | DATE | S | Indica a data de ajuste de O.S. |
| 170 | POLITICAPRIORITARIA | VARCHAR2(1) | S | Indica se a política prioritária foi utilizada por item. |
| 171 | TIPOCALCULOST | VARCHAR2(1) | S | Indica a forma que o ST foi calculado. |
| 172 | QTUNITEMB | NUMBER(18,6) | S | Indica a quantidade uitaria da embalagem. |
| 173 | VLFRETE_RATEIO | NUMBER(18,6) | S | Indica o valor do frete no raterio. |
| 174 | VLOUTRAS_RATEIO | NUMBER(18,6) | S | Indica o valor de outras despesas no rateio. |
| 175 | BASEICST_ANT_RATEIO | NUMBER(18,6) | S | Indica o valor da base de ST antes do rateio de frete e outras despesas. |
| 176 | ST_ANT_RATEIO | NUMBER(18,6) | S | Indica o valor do ST antes do rateio de frete e outras despesas. |
| 177 | ST_DIF_RATEIO | NUMBER(18,6) | S | Indica o valor da diferença de ST antes do rateio de frete e outras despesas. |
| 178 | VLVERBACMV | NUMBER(18,6) | S | Valor verba fornecedor para redução CMV. |
| 179 | NUMVOLUMESCONFERENCIA | NUMBER(4) | S | Indica a quantidade de volumes confereido de itens. |
| 180 | ROTINA | VARCHAR2(40) | S | Indica o nome da rotina. |
| 181 | PERCDESCPIS | NUMBER(12,4) | S | Indica o % de desconto PIS aplicado na venda. |
| 182 | PERCDESCCOFINS | NUMBER(12,4) | S | Indica o % de desconto COFINS aplicado na venda. |
| 183 | VLDESCREDUCAOPIS | NUMBER(24,6) | S | Indica o valor de desconto PIS aplicado na venda.. |
| 184 | VLDESCREDUCAOCOFINS | NUMBER(24,6) | S | Indica valor de desconto COFINS aplicado na venda. |
| 185 | PERCOM4 | NUMBER(8,4) | S | Indica a comissão do terceiro profissional. |
| 186 | NUMETIQUETA | NUMBER(6) | S | Indica o número da etiqueta do produto quando definido como frios. |
| 187 | CODFUNCCONF2 | NUMBER(8) | S | Indica o código funcionário conferente 2. |
| 188 | TIPOSEPARACAO | VARCHAR2(1) | S | Indica o tipo de separação. |
| 189 | PRECOFVBRUTO | NUMBER(18,6) | S | Preço bruto/Negociado FV. |
| 190 | CODFUNCALTLOTE | NUMBER(8) | S | Indica o funcionário responsável por alterar o lote do item do pedido |
| 191 | DTGERACAOOS | DATE | S | Indica a data de geração da OS. |
| 192 | ALTERNATIVO | VARCHAR2(10) | S | Alternativa da formula de tinta. |
| 193 | SIGLAQUALIDADE | VARCHAR2(10) | S | Sigla da qualidade da tinta. |
| 194 | VOLUMEDESEJADO | NUMBER(12,4) | S | Volume de tinta vendido. |
| 195 | CODBASE | VARCHAR2(40) | S | Código da base vendida. |
| 196 | CODFORMULA | VARCHAR2(20) | S | Código da formula vendida. |
| 197 | USADEBCREDRCABRIND | VARCHAR2(1) | S | Gerar débito ou crédito para o RCA no pedido brinde. |
| 198 | CODDESCONTO | NUMBER(8) | S | Código da política de desconto. |
| 199 | CODCOMBO | NUMBER(10) | S | Código da política de combo. |
| 200 | DTENTREGA | DATE | S | Data Entrega Futura - Encomenda. |
| 201 | MOVIMENTACONTACORRENTERCA | VARCHAR2(1) | S | Movimenta conta corrente RCA. |
| 202 | VLREDPVENDASIMPLESNA | NUMBER(18,6) | S | Valor de redução no preço de venda para cliente Simples Nacional. |
| 203 | VLREDCMVSIMPLESNAC | NUMBER(18,6) | S | Valor red. cmv simples nacional. |
| 204 | IDPATRIMONIO | VARCHAR2(75) | S | Identificação patrimonial. |
| 205 | PERDESCFOB | NUMBER(5,2) | S | Percentual de Desconto FOB. |
| 206 | PRODDESCRICAODANFE | VARCHAR2(500) | S | Descrição produto Danfe |
| 207 | CODCONTRATO | NUMBER(6) | S | Descrição produto Danfe |
| 208 | PERCICM | NUMBER(12,4) | S |  |
| 209 | NUMCAIXA | VARCHAR2(10) | S | Numero da caixa |
| 210 | ORIGEMDEPRECO | CLOB | S |  |
| 211 | ROTINALANCULTALT | VARCHAR2(48) | S | Rotina que gravou o registro |
| 212 | PRECOMAXCONSUM | NUMBER(18,6) | S | Preço máximo consumidor |
| 213 | DESCPRECOFAB | NUMBER(18,6) | S | Desconto Preço de Fábrica |
| 214 | BONIFIC | VARCHAR2(1) | S | INFORMAR SE O PRODUTO É BONIFICAÇÃO OU NÃO |
| 215 | PERCDESCINDUSTRIA | NUMBER(18,6) | S | % desconto industria |
| 216 | CODLINHAPRAZO | NUMBER(6) | S | Código Linha Prazo |
| 217 | STPBASERCA | NUMBER(18,6) | S | Valor do ST na base RCA |
| 218 | STPTABELA | NUMBER(18,6) | S | Valor do ST no preço de  Tabela |
| 219 | QTLITRAGEM | NUMBER(18,6) | S | Quantidade em litragem do produto |
| 220 | IDVENDA | NUMBER(6) | S | Identificação de Venda |
| 221 | PERDESCBOLETO | NUMBER(18,6) | S | Percentual do Desconto do Produto que será concedido no Boleto |
| 222 | NUMSEQITEMCONTRATO | NUMBER(6) | S | Número da Sequência do Item no Contrato |
| 223 | NUMLISTA | NUMBER(6) | S | Número da Lista de Presente |
| 224 | VLDESCBOLETO | NUMBER(18,6) | S | Valor do Desconto do Boleto |
| 225 | VLDESCABATIMENTO | NUMBER(18,6) | S | Valor do desconto referente ao abatimento. |
| 226 | PERCDESCABATIMENTO | NUMBER(18,6) | S | Percentual de desconto referente ao abatimento |
| 227 | QTRESERVANT | NUMBER(20,6) | S | Quantidade reservada do item |
| 228 | VLIPIOUTRAS | NUMBER(18,6) | S | Esse atributo serve para armazenar o valor do ipi para devolução simbolica. |
| 229 | PERCIPIOUTRAS | NUMBER(18,6) | S | Esse atributo serve para armazenar o valor do ipi para devolução simbolica. |
| 230 | GRUPOFATURAMENTO | VARCHAR2(2) | S | GRUPO FATURAMENTO |
| 231 | PARTICIPAGIRO | VARCHAR2(1) | S | Participa Giro |
| 232 | RP_IMEDIATA | VARCHAR2(1) | S | Retirada Posterior Imediata |
| 233 | QTUN | NUMBER(16) | S | Quantidade em unidades de venda |
| 234 | ROTINALANCBRINDE | NUMBER(6) | S | Rotina Lançamento de Brinde |
| 235 | SUGESTAO | VARCHAR2(1) | S | Informa se o item foi sugestão de venda |
| 236 | CODEMITENTEITEMPEDIDO | NUMBER(8) | S | Código emite item pedido |
| 237 | SUGERIDO | VARCHAR2(1) | S | Informa se o item foi sugestão de venda |
| 238 | PERDESCINICOMISS | NUMBER(12,6) | S | Perc. Desc. Inicial Faixa Comissão. |
| 239 | PERDESCFIMCOMISS | NUMBER(12,6) | S | Perc. Desc. Final Faixa Comissão. |
| 240 | CONCEDERMAIORCOMISSREG | VARCHAR2(1) | S | Conceder maior Comissão das Faixas. |
| 241 | VLSUBTOTITEM | NUMBER(18,2) | S | Valor total do item do pedido. |
| 242 | PERDESCNEGOCIADO | NUMBER(18,6) | S | Percentudal de desconto negociado |
| 243 | FORMANEGOCIACAO | VARCHAR2(1) | S | Forma de negociação |
| 244 | PERDESCAVISTA | NUMBER(18,6) | S | Percentual de desconta à vista |
| 245 | NEGOCIACAOPOSTERIOR | VARCHAR2(1) | S | Negociação Posterior |
| 246 | CODPRECOFIXO | NUMBER(18,6) | S | Código preço fixo |
| 247 | VLACRESFRETEKG | NUMBER(12,6) | S | Valor de acréscimo Frete (kG) |
| 248 | STATUSSUCATA | NUMBER(1) | S | Status Sucata |
| 249 | PTABELAAUTPECAS | NUMBER(18,6) | S | Preço de tabela para auto peças. |
| 250 | GRPREGRABRINDE | NUMBER(6) | S | Grade brinde. |
| 251 | NUMITEMPED | NUMBER(8) | S | NÚMERO DO ITEM NO XML DA NOTA FISCAL. |
| 252 | VLITEMTRIBUTOS | NUMBER(18,6) | S | Valor dos impostos do item. |
| 253 | PERCTRIBUTOS | NUMBER(4,2) | S | Percentual a ser calculado sobre o item. |
| 254 | TOTALIZADORALIQUOTA | VARCHAR2(10) | S | Totalizador alíquotas. |
| 255 | PERDESCPAUTA | NUMBER(18,6) | S | PERCENTUAL DE DESCONTO UTILIZADO NA PAUTA |
| 256 | ORIGEMST | VARCHAR2(1) | S | Origem do cálculo de ST |
| 257 | VLDESCSOCIOTORCEDOR | NUMBER(18,2) | S | Desconto Sóscio Torcedor |
| 258 | NUMRECOPI | NUMBER(20) | S | Número sujeito ao controle de papel imune |
| 259 | NUMERORECOPI | VARCHAR2(20) | S | Número sujeito ao controle de papel imune |
| 260 | VENDAPBM | VARCHAR2(1) | S |  |
| 261 | UNIDADE | VARCHAR2(2) | S | Unidade da fórmula de tinta |
| 262 | AMBIENTE | VARCHAR2(50) | S | Nome do Ambiente |
| 263 | TAXACASOMOEDAREAL | NUMBER(12,6) | S | Taxa caso a moeda escolhida seja o real. |
| 264 | CODMOEDAESTRAGEIRA | NUMBER(6) | S | Guarda código da moeda estrageira no momento da venda |
| 265 | VLRMOEDAESTRAGEIRA | NUMBER(18,6) | S | Guarda valor da conversão do real para moeda extrangeira |
| 266 | QTDIASENTREGAITEM | NUMBER(4) | S | Qtde de dias para entregar o produto sem estoque |
| 267 | PERCIPIECF | NUMBER(12,4) | S | Percentual de IPI |
| 268 | VLIPIECF | NUMBER(18,6) | S | Valor do IPI |
| 269 | BASEIPIECF | NUMBER(18,6) | S | Base de cálculo do IPI |
| 270 | USAUNIDADEMASTER | VARCHAR2(1) | S | Indica se o produto foi gravado para respeitar a quantidade master ou não. |
| 271 | DTINICIOPROMOLOTE | DATE | S | Data Inicio da promoção |
| 272 | DTFIMPROMOLOTE | DATE | S | Data Fim da promoção |
| 273 | QT_SEPARADAMANIF | NUMBER(20,6) | S | Quantidade separada |
| 274 | CODCONTROLEVASILHAME | NUMBER(8) | S | Código de controle do vasilhame |
| 275 | CODVASILHAMEECF | NUMBER(8) | S | Código do vasilhame |
| 276 | QTSAIDAVASILHAME | NUMBER(8) | S | Quantidade de saída de vasilhames |
| 277 | QTVENDIDAVASILHAME | NUMBER(8) | S | Quantidade de vasilhames vendidos |
| 278 | VLACRESCVASILHAME | NUMBER(18,6) | S | Valor total de acrescímo de vasilhames |
| 279 | PVENDAVASILHAME | NUMBER(18,6) | S | Preço unitário de acrescíemo por vasilhame |
| 280 | PERCREDALIQIPI | NUMBER(18,6) | S | Percentual de Redução da Alíquota de IPI |
| 281 | CORTE | VARCHAR2(1) | S | Indica se o item sofreu corte |
| 282 | NUMSERIESAT | VARCHAR2(13) | S | Número de série equipamento SAT |
| 283 | CODDESCONTOBASERCA | NUMBER(8) | S | Código do desconto que movimentou o preço base rca. |
| 284 | CODPRODCESTA | NUMBER(6) | S | Código do produto cesta Kit Aberto |
| 285 | NUMSEQCESTABASICA | NUMBER(6) | S |  |
| 286 | CODINDICEMULTIPLICADOR | NUMBER(6) | S | código indice multiplicador de serviço |
| 287 | MARGEMMIN | NUMBER(18,2) | S | Margem de lucratividade mínima |
| 288 | PVENDALIQ | NUMBER(18,6) | S | preço líquido |
| 289 | NUMPEDRAS | NUMBER(10) | S | Número do Pedido de Rastreabilidade |
| 290 | NUMSEQRAS | NUMBER(10) | S | Número de Sequência de Rastreabilidade |
| 291 | VLFCPPART | NUMBER(18,6) | S | Valor do FUNCEP |
| 292 | VLICMSPART | NUMBER(18,6) | S | Valor do ICMS Interestadual para a UF Destino |
| 293 | PERCPROVPART | NUMBER(5,2) | S | Percentual provisório de partilha de ICMS |
| 294 | VLICMSDIFALIQPART | NUMBER(22,6) | S | Valor de ICMS do diferencial de aliquota da partilha de ICMS |
| 295 | PERCBASEREDPART | NUMBER(5,2) | S | Redução aplicada na base de Partilha no pedido |
| 296 | VLICMSPARTDEST | NUMBER(18,6) | S | Valor de icms partilha na UF destino |
| 297 | VLBASEPARTDEST | NUMBER(18,6) | S | Valor da base de calculo na Uf de destino |
| 298 | ALIQFCP | NUMBER(18,6) | S | Aliquota de FCP (fundo de combate a pobreza) |
| 299 | ALIQINTERNADEST | NUMBER(18,6) | S | Aliquota de ICMS interna na UF de destino |
| 300 | VLICMSPARTREM | NUMBER(18,6) | S | Valor do ICMS de partilha para UF remetente |
| 301 | ALIQINTERORIGPART | NUMBER(18,6) | S | Valor de ICMS da UF de origem |
| 302 | VLBASEPART_DESATIVADO | NUMBER(18,6) | S | Valor da base de calculo na Uf de destino |
| 303 | ALIQFCPPART_DESATIVADO | NUMBER(18,6) | S | Percentual do FUNCEP |
| 304 | ALIQICMSPART_DESATIVADO | NUMBER(18,6) | S | Percentual de ICMS na UF Destino |
| 305 | VLACRESPART_DESATIVADO | NUMBER(18,6) | S | Valor a ser acrescido ao Preço s/imposto(Valor ICMS + ValorFuncep) |
| 306 | VLIPIPTABELA | NUMBER(18,6) | S | Valor de IPI relativo ao preço de tabela |
| 307 | VLIPIPBASERCA | NUMBER(18,6) | S | Valor de IPI relativo ao preço base RCA |
| 308 | VLICMSPARTPTABELA | NUMBER(18,6) | S | Valor ICMS Partilha relativo ao preço de tabela |
| 309 | VLICMSPARTPBASERCA | NUMBER(18,6) | S | Valor ICMS Partilha relativo ao preço base RCA |
| 310 | INICIOINTERVALODESCQUANT | NUMBER(10,4) | S |  |
| 311 | CODCEST | VARCHAR2(7) | S | Código Cest |
| 312 | QTDIFPESO | NUMBER(18,6) | S | Diferença entre o peso pedido e o peso real do item. |
| 313 | ORIGMERCTRIB | VARCHAR2(1) | S | Código de origem da mercadoria para geração da situação tributária |
| 314 | VLTOTSERVICO | NUMBER(12,6) | S | VALOR TOTAL DO SERVIÇO DO PRODUTO NO PEDIDO |
| 315 | PRODUZIR_TINTA | VARCHAR2(1) | S | Produzir Tinta Manipulada |
| 316 | VLBASEPART | NUMBER(18,6) | S | Valor da base de calculo na Uf de destino |
| 317 | ALIQFCPPART | NUMBER(18,6) | S | Percentual do FUNCEP |
| 318 | ALIQICMSPART | NUMBER(18,6) | S | Percentual de ICMS na UF Destino |
| 319 | VLACRESPART | NUMBER(18,6) | S | Valor a ser acrescido ao Preço s/imposto(Valor ICMS + ValorFuncep) |
| 320 | CODDESCONTOSIMULADOR | NUMBER(8) | S | Indica o código do desconto de simulador no item |
| 321 | CODFIGVENDATRIANGULAR | NUMBER(4) | S | Define o código da figura tributaria para venda triangular |
| 322 | NUMOSSERVICO | NUMBER(6) | S | numero da ordem de serviço |
| 323 | CODBARRABALANCA | VARCHAR2(100) | S | Código de barra enviado pela balança |
| 324 | CODFISCAL | NUMBER(8) | S | Define o CFOP do item no pedido |
| 325 | SITTRIBUT | VARCHAR2(3) | S | Define o CST do item no pedido |
| 326 | VERSAOSERVICOPARTILHA | VARCHAR2(10) | S | Versão do servido de ICMS Partilha |
| 327 | PRODIMPORTADOPEPS | VARCHAR2(1) | S | Produto Importado PEPS |
| 328 | NUMTRANSENTPEPS | NUMBER(10) | S | Transação de Entrada PEPS |
| 329 | PTABELAFABRICAZFM | NUMBER(18,6) | S |  |
| 330 | CODSUPERVISOR | NUMBER(4) | S | Indica o código do supervisor por item |
| 331 | QTMINATACVENDA | NUMBER(18,6) | S | Qt. Min de Atacado do produto durante a venda |
| 332 | TIPODESCATACVENDA | VARCHAR2(1) | S | Tipo de desconto de atacado utilizado na venda |
| 333 | CODOFERTA | NUMBER(6) | S | Código da oferta programada utilizado na venda. |
| 334 | CNPJFABRICANTE | VARCHAR2(18) | S | CNPJ do fabricante da mercadoria. |
| 335 | FABRICANTE | VARCHAR2(60) | S | Nome do fabricante da mercadoria |
| 336 | INDESCALARELEVANTE | VARCHAR2(1) | S | Defini se a mercadoria é de escala relevante. |
| 337 | CODAGREGACAO | VARCHAR2(20) | S | Código de agregação do produto |
| 338 | VLDESCCARCACA | NUMBER(18,6) | S | Desconto concedido na venda com a promessa de entrega de carcaça |
| 339 | DEVOLUCAOCARCACA | VARCHAR2(1) | S | Identifica se existe promessa de devolução de carcaça ou não |
| 340 | NUMCHAVEEXP | VARCHAR2(44) | S | Número da Chave de acesso da Nfe recebida para exportação |
| 341 | NUMDRAWBACK | NUMBER(11) | S | Número do ato concessório do Drawback |
| 342 | NUMREGEXP | NUMBER(12) | S | Número do registro de Exportação |
| 343 | CODBENEFICIOFISCAL | VARCHAR2(10) | S | Codigo beneficio fiscal fonecido para o item |
| 344 | VLBASEFCPICMS | NUMBER(18,6) | S | Valor da base de calculo do Fundo de Combate a Pobreza |
| 345 | VLBASEFCPST | NUMBER(18,6) | S | Valor de base de calculo do Fundo de Combate a Pobreza ST |
| 346 | VLBCFCPSTRET | NUMBER(18,6) | S | Valor da base de calculo do FCP retido anteriormente por ST |
| 347 | PERFCPSTRET | NUMBER(12,4) | S | Percentual do FCP retido anteriormente por Substituição Tributaria |
| 348 | VLFCPSTRET | NUMBER(18,6) | S | Valor do FCP retido por Substituição Tributária |
| 349 | PERFCPSN | NUMBER(12,4) | S | Aliquota aplicável de cálculo do crédito(SIMPLES NACIONAL) |
| 350 | VLCREDFCPICMSSN | NUMBER(18,6) | S | Valor crédito do ICMS que pode ser aproveitado nos termos do art. 23 da LC 123 (SIMPLES NACIONAL) |
| 351 | VLFECP | NUMBER(18,6) | S | Valor Fecp. |
| 352 | VLACRESCIMOFUNCEP | NUMBER(18,6) | S | Valor Acrescimo FUNCEP |
| 353 | PERACRESCIMOFUNCEP | NUMBER(12,4) | S | Percentual Acrescimo FUNCEP |
| 354 | ALIQICMSFECP | NUMBER(12,4) | S | Alíquota de ICMS Fecp. |
| 355 | FATCONVCOMBO | NUMBER(18,6) | S | Indica o fator de conversão informado para inclusão do combo |
| 356 | TIPOCOMBO | VARCHAR2(1) | S | Identifica se o compo 'A - Ativando campanha ' ou 'C - Combo Contínuo'  |
| 357 | PGLP | NUMBER(7,4) | S | Percentual do GLP derivado de petroleo. |
| 358 | PGNN | NUMBER(7,4) | S | Percentual do Gás Natural Nacional. |
| 359 | PGNI | NUMBER(7,4) | S | Percentual do Gás Natural importado. |
| 360 | VPART | NUMBER(23,10) | S | Deve ser informado neste campo o valor por quilograma sem ICMS. |
| 361 | BRINDEVAREJO | VARCHAR2(1) | S | BRINDE DA ROTINA 2046 |
| 362 | UTILIZOUMOTORCALCULO | VARCHAR2(1) | S | Indica se o produto teve o seu cálculo de impostos feitos pelo motor de cálculo |
| 363 | ANP | NUMBER(9) | S | Código do produto na Agência Nacional de Petróleo |
| 364 | DESCANP | VARCHAR2(95) | S | Descrição do Produto na Agência Nacional de Petróleo |
| 365 | NUMSEQITEMBRINDE | NUMBER(20) | S | NÚMERO SEQUENCIAL DO ITEM QUE ORIGINOU O BRINDE |
| 366 | QTORIGTV8 | NUMBER(22,6) | S | Quantidade orignal do pedido de venda TV8 |
| 367 | CUPOMDESCONTO | VARCHAR2(50) | S | Cupom de desconto motor de oferta varejo |
| 368 | CODCUPOMDESCONTO | VARCHAR2(15) | S | Cod. Cupom de desconto motor de oferta varejo |
| 369 | BAIXAQTFRENTELOJA | VARCHAR2(1) | S | Identifica baixa de estoque na loja ou no deposito |
| 370 | CODMAQUINA | NUMBER(4) | S | CÓDIGO DA MÁQUINA |
| 371 | DTCONSOLIDACAOAUXPRO | DATE | S | Data de consolidação do estoque na PCAUXPRO |
| 372 | IDREMESSAWEB | NUMBER(22) | S | id da remessa web quando pedido for multi remessa |
| 373 | CODFORMULACMV | VARCHAR2(200) | S | Código da fórmula do CMV |
| 374 | NUMVERBACAMPANHA | NUMBER(8) | S | Nr. da verba atribuída a campanha. |
| 375 | PERCCUSTFORNEC | NUMBER(12,4) | S | Percentual custeado pelo fornecedor. |
| 376 | VLCUSTFORNEC | NUMBER(18,6) | S | Valor unitário custeado pelo fornecedor. |
| 377 | CODPRECOCESTA | NUMBER(22) | S | Código do preço fixo da cesta básica. |
| 378 | POSICAOCFAT | VARCHAR2(3) | S | Posição do Item na Trarnsferência CFAT (AF ou em nulo = Aguarda Iniciar processo Faturamento CFAT, SR = Gerou Saída da Filial Retira, EVE = Entrou na Filial Venda, SVE = Gerou Saída da Filial Venda, EVI = Entrou na Filial Virtual, F = Pedido foi faturado) |
| 379 | REGIMEESPISENSTFONTE | VARCHAR2(1) | S | Regime Especial Isenção ST Fonte |
| 380 | OBSERVACAOSTFONTE | VARCHAR2(2000) | S | Observação do ST Fonte |
| 381 | CODCONFIGFUNCEPMED | NUMBER(4) | S | Código Config. FECP Farma |
| 382 | CODDEPOSITO | NUMBER(10) | S | Código do Depósito da Filial de Venda |
| 383 | FATORGRAMATURALICIT | NUMBER(30,16) | S | Fator da Gramatura |
| 384 | NUMPEDCLI | VARCHAR2(15) | S | Número Pedido Cliente |
| 385 | NUMLOTEPROMOCAOMED | VARCHAR2(15) | S | Lote da Promoção para indução do lote |
| 386 | CODIGOINTEGRACAOWMS | VARCHAR2(20) | S | código da integração do WMS |
| 387 | CODPROMOCAOMED | NUMBER(9) | S | Código da Promoção Medicamento |
| 388 | NUMVIASMAPASEPRI | NUMBER(3) | S | Número de vias para o mapa de separação dos itens com retira imediata. |
| 389 | NUMEMPENHO | VARCHAR2(20) | S | Número do Empenho |
| 390 | CODEDITAL | NUMBER(9) | S | Código do Edital |
| 391 | LOTECONTRATO | VARCHAR2(10) | S | Lote do Item do edital |
| 392 | PRECOSEMIMPOSTO | NUMBER(18,6) | S | Campo que define o preço sem imposto do item |
| 393 | VLDESCCMVPROMOCAOMED | NUMBER(18,6) | S | Vl. Verba Fornec. Promocao |
| 394 | BCSTRETANTERIOR | NUMBER(18,6) | S | Base de calculo do ST recolhido anteriormente |
| 395 | VLICMSSUBSTITUTOANTERIOR | NUMBER(18,6) | S | Valor do ICMS substituto recolhido anteriormente |
| 396 | VLICMSSTRETANTERIOR | NUMBER(18,6) | S | Valor do ST recolhido anteriormente |
| 397 | QBCMONORET | NUMBER(21,10) | S | Base mono ret |
| 398 | ADREMICMSRET | NUMBER(7,4) | S | Icms retido |
| 399 | VICMSMONORET | NUMBER(23,10) | S | Valor ICMS mono ret |
| 400 | NUMSEQORIG | NUMBER(20) | S | Sequência original para auxiliar no desdobramento na atribuição de lotes pela 930 |
| 401 | USACASHBACK | VARCHAR2(1) | S |  Indica que o % informado PERDESCBOLETO  é relacionado a CASHBACK. |
| 402 | VLIPISUPENSO | NUMBER(18,6) | S | Valor de IPI Suspenso |
| 403 | VLIISUSPENSO | NUMBER(18,6) | S | Valor de II Suspenso |
| 404 | CODBENEFICIOFISCALCOMPLE | VARCHAR2(100) | S | Código do benefício fiscal complementar |
| 405 | INDDEDUZDESONERACAO | NUMBER(1) | S | Identifica se houve ou não desoneração ( 0 -sem desoneração / 1 - com desoneração). |
| 406 | QTCOMBOVIRTUAL | NUMBER(12,4) | S | Quantidade de Combo Utilizado da Campanha de Desconto |
| 407 | CUSTOULTENTFISCAL | NUMBER(18,3) | S | Custo Patrimonial |
| 408 | PERDESCMAXCAMPANHA | NUMBER(18,6) | S | Percentual máximo da política utilizada |
| 409 | PERDESCCAMPANHA | NUMBER(18,6) | S | Percentual informado pelo usuário |
| 410 | PBASECAMPANHA | NUMBER(18,6) | S | Preço de venda aplicado caso não houvesse nenhuma política |
| 411 | PRECOFIXOCAMPANHA | NUMBER(18,6) | S | Preço de venda partindo da política da 357, sem outra política de desconto |
| 412 | ALIQCBS | NUMBER(18,6) | S | Alíquota CBS |
| 413 | ALIQIBS | NUMBER(23,10) | S | Aliquota para calculo do valor IBS |
| 414 | ALIQIS | NUMBER(18,6) | S | Alíquota do Imposto Seletivo |
| 415 | BASECBS | NUMBER(18,6) | S | Valor base para calculo do valor CBS |
| 416 | BASEIBS | NUMBER(18,6) | S | Valor base para calculo do valor IBS |
| 417 | BASEIS | NUMBER(18,6) | S | Valor base para calculo do valor IS |
| 418 | VLCBS | NUMBER(23,10) | S | Valor da CBS |
| 419 | VLIBS | NUMBER(18,6) | S | Valor do imposto IBS |
| 420 | VLIS | NUMBER(18,6) | S | Valor do Imposto Seletivo |
| 421 | CODCBS | NUMBER(10) | S | Código da tributação do imposto CBS na tabela PCTRIBUTACAO |
| 422 | CODIBS | NUMBER(10) | S | Código da tributação do imposto IBS na tabela PCTRIBUTACAO |
| 423 | CODIS | NUMBER(10) | S | Código da tributação do imposto IS na tabela PCTRIBUTACAO |
| 424 | VLCBSPTABELA | NUMBER(18,6) | S | Valor do imposto CBS aplicado no preço de tabela |
| 425 | VLIBSPTABELA | NUMBER(18,6) | S | Valor do imposto IBS aplicado no preço de tabela |
| 426 | VLISPTABELA | NUMBER(18,6) | S | Valor do imposto IS aplicado no preço de tabela |
| 427 | VLCBSPBASERCA | NUMBER(18,6) | S | Valor do imposto CBS aplicado no preço base do RCA |
| 428 | VLIBSPBASERCA | NUMBER(18,6) | S | Valor do imposto IBS aplicado no preço base do RCA |
| 429 | VLISPBASERCA | NUMBER(18,6) | S | Valor do imposto IS aplicado no preço base do RCA |
| 430 | PRECOUTILIZADONFE | VARCHAR2(1) | S | Preço a ser utilizado |
| 431 | CODIGOTRIBUTACAOCBSIBS | NUMBER(10) | S | Código da figura tributária na Rotina 4000 |
| 432 | CSTIBSCBS | VARCHAR2(3) | S | CST de IBS e CBS |
| 433 | CCLASSTRIBIBSCBS | VARCHAR2(6) | S | CClassTrib IBS e CBS |
| 434 | VLBASEIBSCBS | NUMBER(23,10) | S | Base CBS   |
| 435 | IBSUF | NUMBER(7,4) | S | Alíquota do IBS de competência das UF |
| 436 | PDIFIBSUF | NUMBER(7,4) | S | Percentual do diferimento |
| 437 | VDIFIBSUF | NUMBER(23,10) | S | Valor do Diferimento |
| 438 | PREDALIQIBSUF | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 439 | PALIQEFETIBSUF | NUMBER(7,4) | S | Alíquota Efetiva do IBS de competência das UF que será aplicada a Base de Cálculo |
| 440 | VIBSUF | NUMBER(23,10) | S | Valor do IBS de competência da UF |
| 441 | PIBSMUN | NUMBER(7,4) | S | Alíquota do IBS de competência do município |
| 442 | PDIFIBSMUN | NUMBER(7,4) | S | Percentual do diferimento |
| 443 | VDIFIBSMUN | NUMBER(23,10) | S | Valor do Diferimento |
| 444 | PREDALIQIBSMUN | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 445 | PALIQEFETIBSMUN | NUMBER(7,4) | S | Alíquota Efetiva do IBS de competência do Município que será aplicada a Base de Cálculo |
| 446 | VIBSMUN | NUMBER(23,10) | S | Valor do IBS de competência do Município |
| 447 | PDIFCBS | NUMBER(7,4) | S | Percentual do diferimento |
| 448 | VDIFCBS | NUMBER(23,10) | S | Valor do Diferimento |
| 449 | PREDALIQCBS | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 450 | PALIQEFETCBS | NUMBER(7,4) | S | Alíquota Efetiva da CBS que será aplicada a Base de Cálculo |
| 451 | CSTTRIBREG | VARCHAR2(3) | S | Código de Situação Tributária do IBS e CBS |
| 452 | CCLASSTRIBREG | VARCHAR2(6) | S | Código de Classificação Tributária do IBS e CBS |
| 453 | PALIQEFETREGIBSUF | NUMBER(7,4) | S | Valor da alíquota do IBS da UF |
| 454 | VTRIBREGIBSUF | NUMBER(23,10) | S | Valor do Tributo do IBS da UF |
| 455 | ALIQEFETREGIBSMUN | NUMBER(7,4) | S | Valor da alíquota do IBS do Município |
| 456 | VTRIBREGIBSMUN | NUMBER(23,10) | S | Valor do Tributo do IBS do Município |
| 457 | PALIQEFETREGCBS | NUMBER(7,2) | S | Valor da alíquota da CBS |
| 458 | VTRIBREGCBS | NUMBER(23,10) | S | Valor do Tributo da CBS |
| 459 | PIBSUFCOMPRAGOV | NUMBER(7,4) | S | Alíquota do IBS de competência do Estado para compra governamental |
| 460 | VIBSUFCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo do IBS da UF calculado  para compra governamental |
| 461 | PIBSMUNCOMPRAGOV | NUMBER(7,4) | S | Alíquota do IBS de competência do Município para compra governamental |
| 462 | VIBSMUNCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo do IBS do Município calculado para compra governamental |
| 463 | PCBGCOMPRAGOV | NUMBER(7,4) | S | Alíquota da CBS para compra governamental |
| 464 | VCBSCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo da CBS calculado para compra governamental |
| 465 | CODIGOTRIBUTACAOIS | NUMBER(10) | S | Código da tributação IS cadastrada na rotina 4000 PCTRIBUTACAO. |
| 466 | CSTIS | VARCHAR2(3) | S | Código de Situação Tributária do Imposto Seletivo E |
| 467 | CCLASSTRIBIS | VARCHAR2(6) | S | Código de Classificação Tributária do Imposto Seletivo |
| 468 | VLBASEIS | NUMBER(23,10) | S | Valor da Base de Cálculo do Imposto Seletivo |
| 469 | ALIQESPECIFICAIS | NUMBER(7,4) | S | Alíquota específica por unidade de medida apropriada |
| 470 | PERMITIRCUMULATIVIDADEVERBACMV | VARCHAR2(1) | S | Indica que se item, ao ser processado, estava com o parâmetro S ou N para uso de Cumulatividade de Verbas. |
| 471 | CPFRESPTECNICOAGRICOLA | VARCHAR2(11) | S | CPF Responsável Técnico Agrícola |
| 472 | NUMRECEITUARIOAGRICOLA | VARCHAR2(30) | S | Número Receituário Agrícola |
| 473 | RETORNO_SERVICO_CBS | CLOB | S | Retorno CBS |
| 474 | RETORNO_SERVICO_IBS | CLOB | S | Retorno IBS |
| 475 | RETORNO_SERVICO_IS | CLOB | S | Retorno IS |
| 476 | VLBASEIPI | NUMBER(18,6) | S | Valor base de IPI |

## Chaves e índices

- Constraint PCPEDI_PK (P): NUMPED, CODPROD, NUMSEQ
- Índice PCPEDI_IDX12: CODPROD, CODFILIALRETIRA
- Índice PCPEDI_IDX13: NUMLOTE, CODPROD
- Índice PCPEDI_IDX14: CODPROMOCAO, CODPROD
- Índice PCPEDI_IDX15: CODPROD, CODIGOBRINDE, POSICAO
- Índice PCPEDI_IDX16: CODCLI, DATA, CODPROD
- Índice PCPEDI_IDX17: CODIGOBRINDE, SYS_NC00390$, DATA
- Índice PCPEDI_IDX18: CODPROD, NUMPED
- Índice PCPEDI_IDX19: DATA, SYS_NC00310$
- Índice PCPEDI_IDX20: NUMPED, NUMVERBAREBCMV, CODPROD
- Índice PCPEDI_IDX21: DATA, CODUSUR, CODPROD, BONIFIC
- Índice PCPEDI_IDX22: DATA
- Índice PCPEDI_IDX23: DATA, CODPROD
- Índice PCPEDI_IDX24: DATA, CONDVENDA
- Índice PCPEDI_IDX25: CODCLI, SYS_NC00310$
- Índice PCPEDI_IDX26: DATA, NUMPED, CODAUXILIAR
- Índice PCPEDI_IDX27: DATA, NUMPED
- Índice PCPEDI_IDX28: DTCONSOLIDACAOAUXPRO
- Índice PCPEDI_IDX29: NUMPED, CODUSUR
- Índice PCPEDI_IDX30: DATA, CODPROMOCAOMED, NUMVERBACAMPANHA
- Índice PCPEDI_INDEX1: NUMPED, CODPROD
- Índice PCPEDI_INDEX10: CODCLI, CODPROD
- Índice PCPEDI_INDEX11: NUMCARAUX, NUMPED
- Índice PCPEDI_INDEX3: NUMCAR, NUMPED, CODPROD
- Índice PCPEDI_INDEX4: NUMCAR, NUMSEQ, TIPOPESO
- Índice PCPEDI_INDEX5: NUMCAR, CODPROD
- Índice PCPEDI_INDEX6: DATA, CODUSUR, CODCLI, POSICAO, CODPROD
- Índice PCPEDI_INDEX9: DATACONF
- Índice PCPEDI_PK UNIQUE: NUMPED, CODPROD, NUMSEQ
