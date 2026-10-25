# RIBERDOCES.PCNFENT — dicionário

Gerado por `gateway inventory` em 2026-10-25. 532 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | ESPECIE | VARCHAR2(2) | N |  |
| 2 | SERIE | VARCHAR2(3) | S | Indica a serie.\|Campo do tipo caracter, de tamanho 3.  |
| 3 | NUMNOTA | NUMBER(10) | N |  |
| 4 | DTEMISSAO | DATE | N |  |
| 5 | DTENT | DATE | N |  |
| 6 | CODFORNEC | NUMBER(9) | N |  |
| 7 | VLTOTAL | NUMBER(12,2) | S |  |
| 8 | CODCONT | NUMBER(10) | N |  |
| 9 | CODFISCAL | NUMBER(8) | N |  |
| 10 | CODFILIAL | VARCHAR2(2) | N |  |
| 11 | OBS | VARCHAR2(255) | S |  |
| 12 | NUMVOL | NUMBER(12) | S |  |
| 13 | TOTPESO | NUMBER(18,6) | S |  |
| 14 | DTSAIDA | DATE | S |  |
| 15 | DTDESCARGA | DATE | S |  |
| 16 | TIPODESCARGA | VARCHAR2(1) | S |  |
| 17 | CODDEVOL | NUMBER(4) | S |  |
| 18 | VLDESCFIN | NUMBER(14,2) | S |  |
| 19 | NUMBONUS | NUMBER(10) | S | Descricao coluna NUMBONUS |
| 20 | CODMOTORISTADEVOL | NUMBER(8) | S |  |
| 21 | NUMTRANSENT | NUMBER(10) | N |  |
| 22 | VLIPI | NUMBER(12,2) | S |  |
| 23 | VLFRETE | NUMBER(12,2) | S |  |
| 24 | VLST | NUMBER(12,2) | S |  |
| 25 | VLDESCONTO | NUMBER(12,2) | S |  |
| 26 | NUMTAB | NUMBER(4) | S |  |
| 27 | VLBASEIPI | NUMBER(14,2) | S |  |
| 28 | CONCILIACAO | VARCHAR2(2) | S |  |
| 29 | DTCONCIL | DATE | S |  |
| 30 | CODFILIALNF | VARCHAR2(2) | S |  |
| 31 | CALCDESCARGA | VARCHAR2(1) | S |  |
| 32 | TEMPOAPRO | NUMBER(6,2) | S |  |
| 33 | TEMPODESCARGA | NUMBER(6,2) | S |  |
| 34 | ESTCOMISSAORCA | VARCHAR2(1) | S |  |
| 35 | NUMAPLIC | NUMBER(8) | S |  |
| 36 | UF | VARCHAR2(2) | S | UF do fornecedor. |
| 37 | VLBASE | NUMBER(12,2) | S |  |
| 38 | VLICMS | NUMBER(12,2) | S |  |
| 39 | VLISENTAS | NUMBER(12,2) | S |  |
| 40 | VLOUTRAS | NUMBER(12,2) | S |  |
| 41 | ALIQUOTA | NUMBER(8,2) | S |  |
| 42 | BASEICST | NUMBER(18,6) | S |  |
| 43 | CODUSURDEVOL | NUMBER(4) | S |  |
| 44 | NUMIDF | VARCHAR2(30) | S |  |
| 45 | NUMFORMULARIO | NUMBER(10) | S |  |
| 46 | DATAEXPARQOL | DATE | S |  |
| 47 | CODFUNCEXPARQOL | NUMBER(8) | S |  |
| 48 | NUMSELONF | NUMBER(10) | S |  |
| 49 | DTSELOFISCAL | DATE | S |  |
| 50 | CODFUNCLANC | NUMBER(8) | S |  |
| 51 | VLTOTGER | NUMBER(12,2) | S |  |
| 52 | NUMNOTAENTFUTURA | NUMBER(10) | S |  |
| 53 | NUMNEGOCIACAO | NUMBER(12) | S |  |
| 54 | LOTEDIGITADO | VARCHAR2(1) | S |  |
| 55 | GERANFVENDA | VARCHAR2(1) | S |  |
| 56 | GERANFDEVCLI | VARCHAR2(1) | S |  |
| 57 | HORALANC | VARCHAR2(2) | S |  |
| 58 | MINUTOLANC | VARCHAR2(2) | S |  |
| 59 | ROTINALANC | VARCHAR2(48) | S |  |
| 60 | EQUIPLANC | VARCHAR2(64) | S |  |
| 61 | FUNCLANC | VARCHAR2(30) | S |  |
| 62 | DTWMS | DATE | S |  |
| 63 | NFENTREGAFUTURA | VARCHAR2(1) | S |  |
| 64 | VLBONIFIC | NUMBER(12,2) | S |  |
| 65 | PERBASEREDOUTRASDESP | NUMBER(8,2) | S | Contém o percentual de redução para despesas acessórias (se houver). |
| 66 | PERDESCFIN | NUMBER(12,6) | S | Mesmo campo gravado no pedido, que é replicado durante a devolução do cliente. |
| 67 | VLSISCOMEX | NUMBER(18,6) | S |  |
| 68 | VLADUANEIRA | NUMBER(18,6) | S |  |
| 69 | VLOUTRASDESPIMP | NUMBER(18,6) | S |  |
| 70 | CODEXPORTADOR | NUMBER(6) | S |  |
| 71 | NUMTRANSBENEFIC | NUMBER(10) | S | Indica número da transação de entrada de beneficiamento no registro da nota de devolução de remessa para beneficiamento. |
| 72 | ENTRADATIPOI | VARCHAR2(1) | S |  |
| 73 | NUMTRANSORIGEM | NUMBER(10) | S |  |
| 74 | NUMENTMEX | NUMBER(10) | S | Campo de sequencia fiscal voltada para entradas de mercadoria no México. |
| 75 | NSU | NUMBER(15) | S | NSU (número sequencial único) que deve ser gerado e impresso em todas as notas fiscais (entrada e saída). Esta informação é referente a legislação de Santa Catarina.  |
| 76 | DTNSU | DATE | S | Data e hora de geração do NSU, que deve ser impresso em todas as notas fiscais (entrada e saída). |
| 77 | VLVERBACMV | NUMBER(18,6) | S | Indica o valor total da verba CMV registrado na nota fiscal de saída. \|Campo do tipo numérico, de tamanho 18, com 6 casas decimais. |
| 78 | CODROTINADESBLOQUEIO | NUMBER(6) | S |  |
| 79 | VLVERBACMVCLI | NUMBER(18,6) | S | Indica o valor total da verba CMV registrado na nota fiscal de saída. |
| 80 | NUMDIIMPORTACAO | VARCHAR2(40) | S | Campo utilizado para armazenar o número do documento de importação. |
| 81 | CHAVENFE | VARCHAR2(45) | S |  |
| 82 | SITUACAONFE | NUMBER(10) | S |  |
| 83 | PROTOCOLONFE | VARCHAR2(20) | S | Descricao coluna PROTOCOLONFE |
| 84 | DTHORANFE | DATE | S |  |
| 85 | PERCALTERCUSTOENT | NUMBER(8,4) | S | Percentual para alterar o custo da mercadoria na entrada. |
| 86 | NUMLOTENFE | VARCHAR2(15) | S | Número de Lote da Nota Fiscal Eletrônica. |
| 87 | EMISSAOPROPRIA | VARCHAR2(1) | S | Indica se a emissão da nota de entrada é propria ou de terceiros. |
| 88 | TIPOEMISSAO | VARCHAR2(1) | S | Tipo de emissão da NFE (Normal, Contingência, SCAN). |
| 89 | ENVIADA | VARCHAR2(1) | S | Indica se em caso de contingência, já enviou para o SEFAZ. |
| 90 | USAINTEGRACAOWMS | VARCHAR2(1) | S | Usa integração WMS |
| 91 | ENVIADANFE | VARCHAR2(1) | S | NF-e já enviada ao cliente. |
| 92 | DTVENCSTFORANF | DATE | S | Indica a data vencimento. |
| 93 | VLTOTALFRETEFOB | NUMBER(12,2) | S | Indica o valor total do frete FOB. |
| 94 | PESOFRETEFOB | NUMBER(18,6) | S | Indica o peso do frete FOB. |
| 95 | BLOQUEIOEDICAO | VARCHAR2(1) | S | Bloquear pedido para edição. |
| 96 | CODENTVEICULO | NUMBER(10) | S | Código de entrada do veículo. |
| 97 | DTTERMINOCIAP | DATE | S | Data do termino do CIAP. |
| 98 | CODRECEITARECOLHIM | VARCHAR2(10) | S | Código da receita. |
| 99 | DTENVIOWMS | DATE | S | Campo para gravar a data, hora e minutos que a nota fiscal foi enviada ao WMS. |
| 100 | NUMSEQENVIO | NUMBER(4) | S | Campo para gravar o número de seqüência dos arquivos enviado. |
| 101 | CODFUNCEXPINTWMS | NUMBER(8) | S | Campo para indicar código do funcionário que exportou arquivos para a Integração MWS. \|Campo do tipo numérico, de tamanho 8, sem casas decimais.  |
| 102 | CODFUNCENVIO | NUMBER(8) | S | Campo para indicar código do funcionário que exportou arquivos para a Integração MWS. \|Campo do tipo numérico, de tamanho 8, sem casas decimais.  |
| 103 | OBS1 | VARCHAR2(120) | S | Indica observação com relação ao recebimento de mercadoria. |
| 104 | OBS2 | VARCHAR2(120) | S | Indica observação com relação ao recebimento de mercadoria. |
| 105 | NUMOP | NUMBER(8) | S | Indica o numero da ordem de produção. |
| 106 | TIPONOTAFISCALBRINDE | VARCHAR2(1) | S | Campo que indica se os itens da nota é brinde. |
| 107 | UFDESEMBARACO | VARCHAR2(2) | S | Indica o local desembaraço. |
| 108 | LOCALDESEMBARACO | VARCHAR2(30) | S |  |
| 109 | DATADIIMPORTACAO | DATE | S | Indica a data de importação. |
| 110 | NUMPARCDEVOL | NUMBER(4) | S | Indica a quantidade de parcelas da devolução. |
| 111 | TENTATIVAENVIONFE | NUMBER(4) | S | Indica a quantidade de tentativas de envio da nota fiscal a SEFAZ. |
| 112 | FORNECEDOR | VARCHAR2(60) | S | Razão social do fornecedor. |
| 113 | CGC | VARCHAR2(18) | S | CNPJ do fornecedor. |
| 114 | IE | VARCHAR2(15) | S | IE do fornecedor. |
| 115 | TIPOFJ | VARCHAR2(1) | S | Tipo pessoa do cliente (quando devolucao). |
| 116 | TIPOFORNEC | VARCHAR2(1) | S | Tipo do fornecedor (industria / comercio). |
| 117 | PERCPISRED | NUMBER(5,2) | S | % Redução PIS. |
| 118 | CODPAIS | NUMBER(6) | S | Código do País. |
| 119 | DESCPAIS | VARCHAR2(60) | S | Descrição do País. |
| 120 | CGCFILIAL | VARCHAR2(14) | S | CNPJ da filial. |
| 121 | IEFILIAL | VARCHAR2(20) | S | IE da filial. |
| 122 | UFFILIAL | VARCHAR2(2) | S | UF da filial. |
| 123 | CODFORFILIAL | NUMBER(6) | S | Cód. Fornecedor para a filial. |
| 124 | AGREGARSTPRODSINTEGRA | VARCHAR2(1) | S | Indica se agrega o ST no produto (Sintegra). |
| 125 | TIPOALIQOUTRASDESP | VARCHAR2(1) | S | Tipo de tributação do frete e despesas. |
| 126 | CODCONTFOR | NUMBER(10) | S | Cod.Conta para fornecedor. |
| 127 | CODCONTFRE | NUMBER(10) | S | Cod.Conta para frete. |
| 128 | TIPOFRETECIFFOB | VARCHAR2(1) | S | Tipo de frete (CIF/FOB). |
| 129 | REVENDA | VARCHAR2(1) | S | Se fornecedor revende, sim ou não. |
| 130 | INDUSTRIALOCAL | VARCHAR2(1) | S | Se fornecedor tem industria local. |
| 131 | IESUBSTTRIBUT | VARCHAR2(20) | S | IE como substituto tributário na UF. |
| 132 | UFCODIGO | VARCHAR2(2) | S | Código da UF. |
| 133 | HISTORICO | VARCHAR2(1) | S | Se gerou ou não histórico. |
| 134 | CODPROFISSIONAL | NUMBER(6) | S | Indica o código do profissional que a mencadoria esta sendo entregue. |
| 135 | CODFORNECFRETE | NUMBER(8) | S | Indica o código da transportadora. |
| 136 | VLTOTALIPI | NUMBER(12,2) | S | Indica o valor do IPI na nota fiscal. |
| 137 | MODELO | VARCHAR2(2) | S | Identifica o modelo do documento fiscal. |
| 138 | CONSIDERAIPINOFRETEFOB | VARCHAR2(1) | S | Verifica se o valor do IPI entrar ou não na base do Frete FOB para distribuição entre os itens |
| 139 | PERCICMSFRETEFOBST | NUMBER(12,4) | S | Idinca o percentual de icms frete fob a ser utilizado como crédito no calculo do ST fora da nf ou BCR. |
| 140 | CODFUNCENVIOWMS | NUMBER(8) | S | Campo para indicar código do funcionário que exportou arquivos para a Integração MWS. |
| 141 | PROTOCOLONFECAN | VARCHAR2(20) | S | Indica o protocolo de cancelamento da NF-e. |
| 142 | PROTOCOLONFEAUTO | VARCHAR2(20) | S | Indica o protocolo de autorizaçãoda NF-e. |
| 143 | NUMVIAS | NUMBER(4) | S | Indica o numero de vias impressas. |
| 144 | ENDERECO | VARCHAR2(40) | S | Indica o endereço do fornecedor. |
| 145 | BAIRRO | VARCHAR2(40) | S | Indica o bairro do fornecedor. |
| 146 | MUNICIPIO | VARCHAR2(30) | S | Indica a cidade do fornecedor. |
| 147 | CEP | VARCHAR2(9) | S | Indica o CEP do fornecedor. |
| 148 | NUMTRANSVENDATROCA | NUMBER(10) | S | Indica o numero da transação a qual ocorreu a troca. |
| 149 | CONSUMIDORFINAL | VARCHAR2(1) | S | Indica consumidor final. |
| 150 | CODMUNICIPIO | NUMBER(10) | S | Indica o código municipio. |
| 151 | CODCONTABFORNEC | NUMBER(12) | S | Indica o código contabil. |
| 152 | ATACADISTA | VARCHAR2(1) | S | Indica se e atacadista. |
| 153 | CODIBGE | NUMBER(10) | S | Indica o código IBGE. |
| 154 | CODFISCALFRETE | NUMBER(4) | S | Indica o CFOP do frete. |
| 155 | PERCICMFRETE | NUMBER(4,2) | S | Indica a aliquota do ICMS frete. |
| 156 | CODFISCALOUTRASDESP | NUMBER(4) | S | Indica o CFOP das despesas. |
| 157 | ALIQICMOUTRASDESP | NUMBER(4,2) | S | Indica a aliquota do ICMS das despesas. |
| 158 | ESPECIEVOLUME | VARCHAR2(100) | S | Indica a espécie do volume. |
| 159 | SIMPLESNACIONAL | VARCHAR2(1) | S | Indica optante pelo simples nacional. |
| 160 | VLCREDPIS | NUMBER(18,6) | S | Indica o valor total de crédito de PIS na NF. |
| 161 | VLCREDCOFINS | NUMBER(18,6) | S | Indica o valor total de crédito de COFINS na NF. |
| 162 | NUMTRANSVENDAORIG | NUMBER(10) | S | Indica a transação original de entrada consignação. |
| 163 | DTLANCTO | DATE | S | Indica a data da inclusão do lançamento. |
| 164 | NUMNOTAVENDA | NUMBER(10) | S | Indica o número da NF de venda. |
| 165 | DTEMISSAONFVENDA | DATE | S | Indica a data da emissão da NF. |
| 166 | SERIENFVENDA | VARCHAR2(3) | S | Indica a serie da NF de venda. |
| 167 | CODFORNECNF | NUMBER(9) | S | Indica o código do fornecedor da NF. |
| 168 | DTHORAAUTORIZACAOSEFAZ | DATE | S | Indica a hora que foi autorizado no sefaz. |
| 169 | RECIBONFE | VARCHAR2(20) | S | Indica o recibo de envio de lote. |
| 170 | LOTEDPEC | VARCHAR2(20) | S | Indica o número do lote dpec. |
| 171 | RECIBODPEC | VARCHAR2(20) | S | Indica o recibo de envio de lote dpec. |
| 172 | AMBIENTE | VARCHAR2(1) | S | Indica o ambiente em que foi enviada a nota. |
| 173 | PROTOCOLOCANCELAMENTO | VARCHAR2(20) | S | Indica o número do protocolo de cancelamento. |
| 174 | DTHORACANCELAMENTOSEFAZ | DATE | S | Indica a data e hora em que foi cancelada no sefaz. |
| 175 | VLSEGURO | NUMBER(12,2) | S | Indica o valor do seguro. |
| 176 | VLDESPDENTRONF | NUMBER(12,2) | S | Indica o Valor outras despesas na NF. |
| 177 | NUMLIVRO_ENTRADA | VARCHAR2(10) | S | Número do livro de entrada do registro da nota. |
| 178 | FOLHALIVRO_ENTRADA | NUMBER(5) | S | Folha do livro de entrada do registro da nota. |
| 179 | DTHORARECIBODPEC | DATE | S | Data e hora do recibo DPEC. |
| 180 | TOTDIFALIQUOTAS | NUMBER(14,2) | S | Total do diferencial de alíquotas. |
| 181 | OBSNFE | CLOB | S | Observações da NF-e. |
| 182 | CONSSTNFPISCOFINS | VARCHAR2(1) | S | Considerar o ST Nota fiscal no PIS/COFINS. |
| 183 | UTILIZAOUTDESPCALCSUFRAMA | VARCHAR2(1) | S | Utiliza outras despesa para calculo do desconto SUFRAMA. |
| 184 | DTA_HORAENVIOSEFAZ | DATE | S | Data e hora do envio do cte ao sefaz. |
| 185 | AMBIENTENFE | VARCHAR2(1) | S | Ambiente nfe. |
| 186 | TOTPESOLIQ | NUMBER(18,6) | S | Total do peso líquido dos proudtos  informados na nota fiscal. |
| 187 | REMESSAUNICA | VARCHAR2(1) | S | Remessa única |
| 188 | JUSTIFICATIVACONTIGENCIA | VARCHAR2(256) | S | Motivo da entrada de contigência. |
| 189 | DTAHORAENTRADACONTIGENCIA | DATE | S | Data e Hora da entrada de contigência. |
| 190 | POSSECODFORNEC | NUMBER(6) | S | Código Fornecedor da empresa que está de posse da mercadoria. |
| 191 | CONFERIDO | VARCHAR2(1) | S | Nota Fiscal Conferida. |
| 192 | CONSUMIUNUMNFE | VARCHAR2(1) | S | Se já tem um número de Nfe para o registro. |
| 193 | VLPEDAGIO | NUMBER(12,2) | S | Valor total do pedágio. |
| 194 | DEDFRETECIFCREDPRESICMS | VARCHAR2(1) | S | Considera frete FOB no calculo da base PIS/COFINS |
| 195 | CODIGOCLASSECONSUMOENERGIA | VARCHAR2(2) | S | Código da Classe de Consumo de Energia. |
| 196 | TIPOLIGACAOENERGIA | VARCHAR2(1) | S | Tipo de Ligação de Energia. |
| 197 | CODIGOGRUPOTENSAOENERGIA | VARCHAR2(2) | S | Código do Grupo de Tensão de Energia. |
| 198 | CODTRIBPISCOFINS | NUMBER(4) | S | Código da Situação Tributária do PIS/COFINS. |
| 199 | VLPIS | NUMBER(18,6) | S | Valor do PIS na Nota Fiscal. |
| 200 | VLCOFINS | NUMBER(18,6) | S | Valor do COFINS na Nota Fiscal. |
| 201 | NUMEROEND | VARCHAR2(6) | S | Número do Logradouro do Fornecedor. |
| 202 | SITDOCPISCOFINS | VARCHAR2(2) | S | Situação do Documento do PIS/COFINS |
| 203 | VLBASEPIS | NUMBER(18,6) | S | Valor da Base do PIS |
| 204 | PERPIS | NUMBER(12,4) | S | Alíquota do PIS |
| 205 | VLBASECOFINS | NUMBER(18,6) | S | Valor da Base do COFINS |
| 206 | PERCOFINS | NUMBER(12,4) | S | Alíquota do COFINS |
| 207 | USACREDPISCOFINSENTBONIFIC | VARCHAR2(1) | S | Utiliza crédito PIS/COFINS na entrada bonificada |
| 208 | OBSPISCOFINS | VARCHAR2(500) | S | Observação da tributação de PIS/COFINS |
| 209 | FINALIDADENFE | VARCHAR2(1) | S | Finalidade NF-e |
| 210 | SERIEORIGINAL | VARCHAR2(3) | S | SERIE ORIGINAL DA NOTA FISCAL ELETRONICA |
| 211 | DTEXPORTACAOWMS | DATE | S | Data  e hora de exportação |
| 212 | DTIMPORTACAOWMS | DATE | S | Data  e hora de importação |
| 213 | EMAILDEST | VARCHAR2(100) | S | E-mail do destinatário |
| 214 | TELEFONE | VARCHAR2(20) | S | Telefone do destinatário |
| 215 | USAOUTRASDESPSEGUROPISCOFINS | VARCHAR2(1) | S | Somar outras despesas e seguros na base PIS/COFINS |
| 216 | CALCDESPFINPRAZOPGTO | VARCHAR2(1) | S | Calcular vendor por prazo de pagamento |
| 217 | RATEIOPESOLIQ | VARCHAR2(1) | S | Informa que o rateio vai ser por peso líquido |
| 218 | TOTALPESOLIQDI | NUMBER(12,4) | S | Peso líquido total da DI |
| 219 | VLIRRF | NUMBER(12,2) | S | "Valor de IRRF	Valor de IRRF" |
| 220 | VLSESTSENAT | NUMBER(12,2) | S | "Valor de VLSESTSENAT	Valor de VLSESTSENAT" |
| 221 | VLISS | NUMBER(12,2) | S | "INC	Valor de VLISS	Valor de VLISS" |
| 222 | VLINSS | NUMBER(12,2) | S | "INC	Valor de VLINSS	Valor de VLINSS" |
| 223 | VLCSRF | NUMBER(12,2) | S | "INC	Valor de VLCSRF	Valor de VLCSRF" |
| 224 | OBSLIVROFISCAL | VARCHAR2(500) | S | Observação |
| 225 | MOVESTOQUEIMPORT | VARCHAR2(1) | S | Utiliza estoque em Trânsito na entrada Importação |
| 226 | FECHAMENTOCEAUTOMATICO | VARCHAR2(1) | S | Fechamento automático do controle de embarque |
| 227 | TIPOTRIBENTIMP | VARCHAR2(1) | S | Tipo de tributação na entrada de importação |
| 228 | CALCBASEICMSALIQREDENTIMP | VARCHAR2(1) | S | Calcular base de ICMS com alíquota reduzida na entrada de importação |
| 229 | INCLUIRIMPOSTOSNFTRANSPIMP | VARCHAR2(1) | S | Incluir impostos no preço da mercadoria NF de transporte importação |
| 230 | CONSIDERARIIPUNIT | VARCHAR2(1) | S | Considerar o imposto de importação no preço unitário do produto DANFE |
| 231 | DESPESASRATEADA | VARCHAR2(1) | S | Informar se o frete e outras despesas foram rateados nos itens |
| 232 | EXPORTADO | VARCHAR2(1) | S | Marcação se o registro foi exportado |
| 233 | WMSCODFUNCIGNORA | NUMBER(8) | S | Funcionário que ignorou WMS |
| 234 | WMSIGNORADO | VARCHAR2(1) | S | WMS Ignorado |
| 235 | MARCAVOLUME | VARCHAR2(40) | S | Marca do container |
| 236 | NUMVOLUME | VARCHAR2(20) | S | Numeração do container |
| 237 | SOMARVLOUTRASDESPADUANEIRAS | VARCHAR2(1) | S | Somar valor outras despesas e despesas aduaneiras no DANFE. |
| 238 | NFEXTEMPORANEA | VARCHAR2(1) | S | Informar se a nota fiscal é extemporanea ou não |
| 239 | UTILIZAIPICALCICMS | VARCHAR2(1) | S | Utiliza IPI no calculo de ICMS |
| 240 | NUMCONTRATO | NUMBER(8) | S | NÚmero do contrato de comodato |
| 241 | NUMOS | NUMBER(8) | S | Campo para armazenar o número da ordem de serviço que gerou a NF. |
| 242 | CHAVECTE | VARCHAR2(44) | S | Chave CTE |
| 243 | NORMAREGESPECIAL | VARCHAR2(250) | S | Norma Especifica (Regime Especial) |
| 244 | SOMARPISCOFINSVLOUTRASDESPIMP | VARCHAR2(1) | S | Incluir PIS/COFINS no valor de Outras Despesas DANFe Importação |
| 245 | TIPOEMISSAOCTE | VARCHAR2(2) | S | Tipo Emissão Ct-e |
| 246 | DTEMISSAODECLARACAO | DATE | S | Data de emissão da declaração emitida pelo tomador do Ct-e, quando este não for contribuinte |
| 247 | DTEMISSAOPRODRURAL | DATE | S | Data de emissão NF produtor rural |
| 248 | MODELOPRODRURAL | VARCHAR2(2) | S | Modelo da NF produtor rural |
| 249 | SERIEPRODRURAL | VARCHAR2(3) | S | Serie da NF produtor rural |
| 250 | NUMNOTAPRODRURAL | NUMBER(10) | S | Numero da NF produtor rural |
| 251 | CHAVENFEPRODRURAL | VARCHAR2(45) | S | Chave da Nfe produtor rural |
| 252 | CALCCREDICMS | VARCHAR2(1) | S | Calcular custo com crédito de ICMS |
| 253 | CALCCREDPISCOFINS | VARCHAR2(1) | S | Calcular custo com crédito de PIS/COFINS |
| 254 | CALCCREDICMCUSTOCONT | VARCHAR2(1) | S | Calcular crédito de ICMS no custo contábil |
| 255 | CALCCREDPISCUSTOCONT | VARCHAR2(1) | S | Calcular crédito de PIS/COFINS no custo contábil |
| 256 | DTPROXATU | DATE | S | Data para a proxima atualização do custo financeiro |
| 257 | PRAZO | NUMBER(6) | S | Prazo medio |
| 258 | REBAIXACUSTOCOMDESCFIN | VARCHAR2(1) | S | Calcular custo com descontofinanceiro |
| 259 | DEVSIMBOLICA | VARCHAR2(1) | S | Devolução simbólica |
| 260 | ENVIADOEMAILCANCELADO | VARCHAR2(1) | S | ENVIADO EMAIL CANCELADO |
| 261 | PLACAVEICULO | VARCHAR2(10) | S | Placa do veículo |
| 262 | UFVEICULO | VARCHAR2(2) | S | UF do veículo |
| 263 | NATOPERNFE | VARCHAR2(60) | S | Natureza de operação da NF-e do estorno. |
| 264 | MOTESTORNONFE | VARCHAR2(250) | S | Observação do motivo da geração da NF-e de estorno. |
| 265 | USADRAWBACK | VARCHAR2(1) | S | Usa regime DrawBack Suspensão |
| 266 | SITUACAONFEORIG | NUMBER(10) | S | Situação NFe Original. |
| 267 | CONSSTGUIAPISCOFINS | VARCHAR2(1) | S | Considera ST Guia Base do PISCOFINS. |
| 268 | DTCANCEL | DATE | S | Data de Cancelamento |
| 269 | GERARBCRNFE | VARCHAR2(1) | S | Gerar ICMS ST BCR |
| 270 | OBSAUTORIDEVOL | VARCHAR2(200) | S | Observação de Autorização de Devolução. |
| 271 | AGREGASTVLMERC | VARCHAR2(1) | S | Agregar valor de ST ao valor da mercadoria. |
| 272 | NOMESEGURADORA | VARCHAR2(30) | S | Nome da Seguradora. |
| 273 | NUMEROAPOLICE | VARCHAR2(20) | S | Número da Apolice do Seguro |
| 274 | NUMEROAVERBACAO | VARCHAR2(40) | S | Número da Averbação do Seguro. |
| 275 | RESPSEGURO | NUMBER(1) | S | Responsável pelo seguro. |
| 276 | IMPORTADOXML | VARCHAR2(1) | S | Informa se a entrada é importada de xml |
| 277 | VERSAOROTINA | VARCHAR2(20) | S | Versão da rotina gerarou os dados |
| 278 | TIPOEMBALAGEMPEDIDO | VARCHAR2(1) | S | Tipo de embalagem V = Venda, M = Master |
| 279 | SITUACAOCTE | NUMBER(10) | S | Situação Cte |
| 280 | PROTOCOLOCTE | VARCHAR2(20) | S | Protocolo Cte |
| 281 | RECIBOCTE | VARCHAR2(20) | S | Recibo Cte |
| 282 | NUMLOTECTE | VARCHAR2(15) | S | Número do lote do Cte |
| 283 | EMITENFSIMBOLICACONSIG | VARCHAR2(1) | S | Emite Nota Fiscal Devolução Simbólica de Consignação |
| 284 | ALTCUSTOENTREMESSA | VARCHAR2(1) | S | Calcula Custo na entrada tipo R (Remessa) |
| 285 | CALCPISCOFINSFRETEFOBCSTPROD | VARCHAR2(1) | S | Define Tributação de PIS/COFINS do Frete FOB pelo CST do produto. |
| 286 | CALCICMSFRETEFOBCSTPROD | VARCHAR2(1) | S | Define Tributação de ICMS do Frete FOB pelo CST do produto. |
| 287 | CODVIA | NUMBER(10) | S | Código da via de transporte |
| 288 | CODPROJETO | NUMBER(6) | S | Código do Projeto. Cadastrado na 607 e usado na 631,749 e 3402. |
| 289 | NFEEMIAVUSEF | VARCHAR2(1) | S | Nfe de Emissão Avulsa SEFAZ |
| 290 | CODPRODNFAJUSTE | VARCHAR2(60) | S | Código de produto para NF-e de ajuste |
| 291 | DESCRICAOPRODNFAJUSTE | VARCHAR2(120) | S | Descrição do produto paa NF ajuste |
| 292 | NCMNFAJUSTE | VARCHAR2(8) | S | NCM para item de NF de ajuste |
| 293 | UNIDADENFAJUSTE | VARCHAR2(6) | S | Unidade de medida para NF de ajuste |
| 294 | OBSNFAJUSTE | VARCHAR2(1000) | S | Observação complementar para NF de ajuste |
| 295 | VLTOTBRUTOPRODAJUSTE | NUMBER(18,2) | S | Valor total bruto para NF de ajuste |
| 296 | NUMTRANSINTEGRACAO | VARCHAR2(40) | S | Número de transação de integração |
| 297 | CONSUMIUNUMCTE | VARCHAR2(1) | S | Consumiu numeração de Cte |
| 298 | TENTATIVAENVIOCTE | NUMBER(10) | S | Tentativa de envio do Cte |
| 299 | NUMPEDPRINC | NUMBER(8) | S | Número do pedido de compra principal |
| 300 | TIPOMOVGARANTIA | NUMBER(3) | S | Define qual tipo de movimentação de garantia a nota representa. |
| 301 | RESULTADOLAUDO | VARCHAR2(1) | S | Resultado Laudo |
| 302 | IDTIPOPRESENCA | VARCHAR2(2) | S | Identificação do tipo de presença |
| 303 | UTILIZAOUTDESPNFBASEST | VARCHAR2(1) | S | Considerar o valor Outras Despesas NF + Seguro na base de calculo ST |
| 304 | DTVENCSTGUIA | DATE | S | Data vencimento ST GUIA |
| 305 | ORIGEMFRETE | VARCHAR2(1) | S | Origem do percentual de Frete |
| 306 | TIPOOPERCIAP | VARCHAR2(1) | S | Tipo de operação referente documento fiscal do CIAP |
| 307 | SITUACAONFEEPEC | NUMBER(3) | S | Situação Nfe EPEC |
| 308 | PROTOCOLONFEEPEC | VARCHAR2(20) | S | Protocolo de autorização EPEC |
| 309 | DATAHORAREGISTROEPEC | DATE | S | Data e hora de registro do EPEC |
| 310 | TIPOPRESENCAADQUIRINTE | VARCHAR2(1) | S | Indica o tipo de presença do comprador no estabelecimento no momento da operação |
| 311 | ALIQISSRET | NUMBER(12,4) | S | Alíquota do ISSQN Retido |
| 312 | CODFISCALISS | NUMBER(4) | S | CFPS - Código Fiscal de Prestação de Serviço |
| 313 | CONSFRETEFOBBASESTNF | VARCHAR2(1) | S | Considera Frete FOB base ST NF |
| 314 | GERACPICMSIMPORTACAO | VARCHAR2(1) | S | Gerar contas a pagar icms importação |
| 315 | GERARPCLANCICMSANTECIP | VARCHAR2(1) | S | Gerar contas a pagar do ICMS Antecipado |
| 316 | GERARPCLANCSTGUIAENT | VARCHAR2(1) | S | Gerar contas a pagar do ST Guia |
| 317 | VLBASECALCRETENCAOISS | NUMBER(18,6) | S | Valor da Base de Cálculo de retenção ISS |
| 318 | VLCONTABILISS | NUMBER(18,6) | S | Valor Contábil ISS |
| 319 | VLISENTASISS | NUMBER(18,6) | S | Valor de isentas ISS |
| 320 | VLISSRETIDOTOMADOR | NUMBER(18,6) | S | Valor do ISS Retido Tomador |
| 321 | NOTADUPLIQUESVC | VARCHAR2(1) | S | Validar as notas quando elas forem enviadas em ambiente de SVC. Quando a nota for duplicada o campo deverá ser preenchido |
| 322 | NUMTRANSENTDEPENDENTE | NUMBER(10) | S | Transação de nota de entrada dependente |
| 323 | NUMTRANSVENDADEPENDENTE | NUMBER(10) | S | Transação de nota de saida dependente |
| 324 | ROTINACAD | VARCHAR2(48) | S | Gravar a rotina que inseriu o registro na PCNFENT |
| 325 | CONTRIBUINTE | VARCHAR2(1) | S | Identifica quando o fornecedor é contribuinte |
| 326 | NUMVIASGNRE | NUMBER(4) | S | Número de impressões da nota na GNRE |
| 327 | QTDEREPROCNFE | NUMBER(4) | S | Número de impressões da nota na GNRE |
| 328 | GERARDIFALIDECNOVO | VARCHAR2(1) | S | Gerar Diferencial de Alíquota com nova sistemática de cálculo Decreto 8.915/15 art. 65 inciso III |
| 329 | VINCULARDOCUMENTOCT | VARCHAR2(1) | S | Vincular Documento Nfe com o Cte para cálculo de frete no CIAP |
| 330 | NUMTRANSORIGEMCT | NUMBER(10) | S | Número de Transação de origem do Cte vinculado com a Nfe para cálculo do CIAP |
| 331 | CODFORNECFUNRURAL | NUMBER(6) | S | Código fornecedor funrural |
| 332 | CALCFUNRURAL | VARCHAR2(1) | S | Calcula contas a pagar separado funrural |
| 333 | CONSCAPATAZIAICMS | VARCHAR2(1) | S | Considera Capatazia na base de ICMS |
| 334 | TRANSFDEP | VARCHAR2(1) | S | Define se a NF é uma nota de transferência para o depósito fechado |
| 335 | RESTRINGIRCREDICMSTV10 | VARCHAR2(1) | S | Restringir credito de ICMS nas entradas TV9 e TV10 |
| 336 | DEDUZIRSUFRAMAALIQEXT | VARCHAR2(1) | S | Deduzir Suframa Créd. ICMS Aliq. Ext. ST |
| 337 | CONSIPICALCICMSANT | VARCHAR2(1) | S | Considera IPI no cálculo do ICMS antecipado |
| 338 | GERARVLDEPRECCREDIM | VARCHAR2(1) | S | Gerar valor a depreciar deduzido valor do Crédito Imobilizado |
| 339 | MOTIVO | VARCHAR2(30) | S | Motivo da Devolução |
| 340 | DTVENCARREC | DATE | S | Data do Vencimento |
| 341 | TIPODOCARRECGNRE | VARCHAR2(1) | S | Tipo de documento - GNRE ou Documento Arrecadação |
| 342 | UFBENEFICIARIA | VARCHAR2(2) | S | Unidade Federação beneficiária |
| 343 | NUMDOCARREC | VARCHAR2(255) | S | Número do documento de arrecadação |
| 344 | CODAUTENTICACAO | VARCHAR2(255) | S | Código de autenticação |
| 345 | VLDOCARRECADACAO | NUMBER(20,6) | S | Valor do documento de arrecadação |
| 346 | DTPAGARREC | DATE | S | Data do Pagamento |
| 347 | TIPOVIATRANSPORTE | NUMBER(10) | S | Tipo de via de transporte da mercadoria |
| 348 | QTDEREPROCCTE | NUMBER(6) | S | Quantidade de reprocessamento CT-e |
| 349 | OBSDESACORDOCTE | VARCHAR2(256) | S | Observação do evento de prestação de desacordo do Cte |
| 350 | PROTOCOLODESACORDOCTE | VARCHAR2(20) | S | Protocolo do evento de prestação de desacordo do Cte |
| 351 | DTHORADESACORDOCTE | DATE | S | Data hora do evento de prestação de desacordo do Cte |
| 352 | ATUALIZOUCUSTOCONT | VARCHAR2(1) | S | Indica se o custo contábil foi atualizado na Nota N |
| 353 | PERCIRRF | NUMBER(12,4) | S | Alíquota de IRRF |
| 354 | PERCISS | NUMBER(12,4) | S | Alíquota de ISS |
| 355 | PERCCSRF | NUMBER(12,4) | S | Alíquota de CSRF |
| 356 | PERCINSS | NUMBER(12,4) | S | Alíquota de INSS |
| 357 | CODCONTSERV | NUMBER(10) | S | Código conta de NF Serviço |
| 358 | GERARNFCOMPLEOUTROSCUSTOS | VARCHAR2(1) | S | Gerar NF complementar de outros custos |
| 359 | CALCREDPISSERVICO | VARCHAR2(1) | S | Calcula crédito de PIS/COFINS no custo de NF serviço |
| 360 | CALCREDPISSERVICOCONT | VARCHAR2(1) | S | Calcula crédito de PIS/COFINS no custo contábil de NF serviço |
| 361 | VALORTV1 | NUMBER(18,6) | S | Valor da devolução de produtos não bonificados |
| 362 | VALORBNFTV1 | NUMBER(18,6) | S | Valor da devolução de produtos bonificados |
| 363 | VALORCREDITODEV | NUMBER(18,6) | S | Valor do crédito criado na devolução |
| 364 | USAMULTUNIDITEM | VARCHAR2(1) | S | Usa multimplas unidades (Master ou Venda) por item no pedido de compra |
| 365 | DEDUZIRCREDITOOUTORGADOCUSTO | VARCHAR2(1) | S | Deduzir crédito ST Outorgado do custo |
| 366 | UTICREDICMSFRETEFOBCUSTOCONT | VARCHAR2(1) | S | Calcula custo contabil com crédito ICMS Frete FOF |
| 367 | USADEBCREDOPERLOG | VARCHAR2(1) | S | Utilizar déb./créd. de conta corrente de Operador Logístico |
| 368 | CONSIDEREFILIALDEPCALCCUSTOENT | VARCHAR2(1) | S | Calcular custo na entrada de mercadoria utilizando estoque deposito terceiro |
| 369 | VLTOTALFCPST | NUMBER(18,6) | S | Valor Total do Fundo de Combate a Pobreza ST |
| 370 | VLTOTALFCPICMS | NUMBER(18,6) | S | Valor Total do Fundo de Combate a Pobreza ICMS |
| 371 | CSTICMSNFAJUSTE | VARCHAR2(3) | S | Código de situação tributária de ICMS para nota do tipo ajuste |
| 372 | CSTPISCOFINSNFAJUSTE | NUMBER(3) | S | Código de situação tributária de PIS/COFINS para nota do tipo ajuste |
| 373 | DEDUZIRSUFRAMADIFALIQ | VARCHAR2(1) | S | Deduzir Suframa Diferença de Alíquotas |
| 374 | DERRUBARCARGA | VARCHAR2(1) | S | Define se derrubou carga(Sim ou Não) |
| 375 | ENTRADAINSERVIVEL | VARCHAR2(1) | S | Indica que a entrada é para produtos inserviveis |
| 376 | RETORNOTV13 | VARCHAR2(1) | S | Gravar SIM neste campo se o parametro "Devolver Simples Remessa TV13/Totalidade", da filial que gerou a nota de devolução estiver igual a SIM. |
| 377 | DEDUZIRDESONERORGAOPUB | VARCHAR2(1) | S | Deduzir ICMS desonerado NF Licitação Orgãos Públicos |
| 378 | CONSIDERARSTGUIADIFPRECO | VARCHAR2(1) | S | Considerar ST Guia para calculo da diferença de preço entrada de mercadoria |
| 379 | DEDUZIRICMSBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do ICMS da base de cálculo do PIS/COFINS |
| 380 | CODMUNINICTE | NUMBER(7) | S | Código do município de origem da prestação do serviço(Cte) |
| 381 | CODMUNFIMCTE | NUMBER(7) | S | Código do município de destino da prestação do serviço(Cte) |
| 382 | VERSAOLAYOUTNFE | VARCHAR2(5) | S | Versão do layout do arquivo na Sefaz |
| 383 | CALCFRETEPORITEM | VARCHAR2(1) | S | Considera Frete FOB por Item |
| 384 | OBSNFEFISCO | CLOB | S | Informações Adicionais de Interesse do Fisco |
| 385 | PCT_ESTOQUE | VARCHAR2(1) | S | Gravar a efetivação da package de estoque |
| 386 | PCT_CUSTO | VARCHAR2(1) | S | Gravar a efetivação do serviço de custo |
| 387 | DETENTORREGESPECIAL | CHAR(1) | S | Detentor de Regime Especial |
| 388 | PERCACRESCIMOCUSTOREAL | NUMBER(12,4) | S | Acréscimo no custo real entrada de mercadoria |
| 389 | PERCACRESCIMOCUSTOFIN | NUMBER(12,4) | S | Acréscimo no custo financeiro entrada de mercadoria |
| 390 | CUSTOZEROCUSCONTBONIFIC | VARCHAR2(1) | S | Considera custo contabil zero mercadoria bonificada |
| 391 | CANCELARPRODUCAO | VARCHAR2(1) | S | CANCELAR PRODUÇÃO TINTOMÉTRICO |
| 392 | DEDUZIRFCPICMSDOFCPST | VARCHAR2(1) | S | Deduzir o FCP ICMS do valor do FCP ST |
| 393 | CONSIDERARPAUTAZEROSTGUIA | VARCHAR2(1) | S | Considerar pauta zero para cálculo do ST Guia |
| 394 | CALCCREDICMSBASEREDUZIDA | VARCHAR2(1) | S | Utiliza crédito de ICMS sobre base reduzida |
| 395 | CODSERVISS | VARCHAR2(4) | S | Código Serviço ISS |
| 396 | QRCODE | VARCHAR2(1000) | S | QRCODE PARA NF-e/CT-e |
| 397 | REMESSACOMBENEFIC | VARCHAR2(1) | S | Retorno de remessa com Entrada de Beneficiamento |
| 398 | CONSPOLITICATRIBUTNOTAH | VARCHAR2(1) | S | Considera o politica comercial e tributação da nota H na nota J |
| 399 | CONTAEORDEMBONIFICADA | VARCHAR2(1) | S | Entrada de conta e ordem bonificada |
| 400 | ENTRADACONSERTOGARANTIA | VARCHAR2(1) | S | Indica se a entrada é para conserto de garantia |
| 401 | UTILIZAPRECOPERSNFIMP | VARCHAR2(1) | S | Define se útiliza o calculo personalizado para o preço do produto |
| 402 | PRECOPERSNFIMP_AFRMM | VARCHAR2(1) | S | Define se o AFRMM vai ser considerado no preço do produto |
| 403 | CROSSDOCKING | VARCHAR2(1) | S | Indica que é uma entada para crossdocking |
| 404 | PRECOPERSNFIMP_FRETE | VARCHAR2(1) | S | Define se útiliza o calculo personalizado para o preço do produto |
| 405 | CODENDENTCLI | NUMBER(6) | S | Campo que informa o endereço de entrega utilizado na nota de saida |
| 406 | VLCARGACTE | NUMBER(22,6) | S | Valor Carga Cte |
| 407 | DTMANIFDEST | DATE | S | Data de manifestação do documento |
| 408 | PERCREDBASEPISCOFINSFRETE | NUMBER(12,4) | S | Percentual de redução da base para o calculo pis/cofins Frete |
| 409 | PCT_ESTOQUE_CANCELAMENTO | VARCHAR2(1) | S | Utilizou PKG Estoque para Cancelar |
| 410 | PCT_CUSTO_CANCELAMENTO | VARCHAR2(1) | S | Utilizou PKG Custo para Cancelar |
| 411 | PCT_RECALC_TRIBUT | VARCHAR2(1) | S | Gravar a efetivação do Recálculo de Tributação |
| 412 | PCT_RECALC_PISCOFINS | VARCHAR2(1) | S | Gravar a efetivação do Recálculo de PisCofins |
| 413 | ORIGEMCUSTOTRANSF | VARCHAR2(1) | S | Define a origem do custo tranferência |
| 414 | CALCCREDFUNCEPCONT | VARCHAR2(1) | S | Calcular credito FUNCEP (FCP ICMS) Contabil |
| 415 | CALCCREDFUNCEP | VARCHAR2(1) | S | Calcular credito FUNCEP (FCP ICMS) Gerencial |
| 416 | NFBRINDE | VARCHAR2(1) | S | Nota de Brinde |
| 417 | CODIGONUMERICOCHAVE | VARCHAR2(8) | S | Código númerico que compoem a chave |
| 418 | TIPOIMPRESSAO | VARCHAR2(1) | S | Tipo de impressão (retrato ou paisagem) |
| 419 | ZERARALIQICMSPRESUMIDO | VARCHAR2(1) | S | Zerar Alíquota ICMS PRESUMIDO para produto Bonificado |
| 420 | DEDUZIRIPIBASEDIFALIQ | VARCHAR2(1) | S | Deduzir IPI da base Diferencial Alíquota |
| 421 | CONSFRETEFOBBASEDIFALIQ | VARCHAR2(1) | S | Deduzir IPI da base Diferencial Alíquota |
| 422 | DEDUZFUNCEPBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do FUNCEP da Base de Cálculo do Pis/Cofins |
| 423 | NUMEMPENHO | VARCHAR2(40) | S | Numero do Empenho |
| 424 | CNPJINTERMEDIADOR | VARCHAR2(20) | S | CNPJ do Intermediador da Transação |
| 425 | DESCINTERMEDIADOR | VARCHAR2(60) | S | Nome do usuário ou identificação do perfil do vendedor no site do intermediador |
| 426 | ALIMENTAFUNDOMULTIFILIAL | VARCHAR2(1) | S | Campo que ira definir o momento de que o fundo multifilial sera alimentado |
| 427 | APLICAVERBAREBCUSTO | VARCHAR2(1) | S | Indica se utiliza aplicação de verba para rebaixa de custo |
| 428 | APLICVERBAREBCUSTO | VARCHAR2(1) | S |  |
| 429 | NFDEVOLMAISNEG | VARCHAR2(1) | S | Flag que marca se foi devolução oriunda do processo MAIS NEGOCIO |
| 430 | SULFRAMA | VARCHAR2(15) | S | Inscrição estadual sulframa |
| 431 | USAICMSENTPROD | VARCHAR2(1) | S | UTILIZA A TRIBUTACAO CADASTRADA NA 212 |
| 432 | SOMAICMSDIFTOTIMP | VARCHAR2(1) | S | SOMA O ICMS DIFERIDO AO TOTAL NA IMPORTAÇÃO |
| 433 | SOMAVLDENTRONFCUSTOCONTTV10 | VARCHAR2(1) | S | Define se será somado os valores dentro da nota, ao custocont ao realizar um TV10 ou TV 9 |
| 434 | ORIGEMCUSTOCONTTRANSF | VARCHAR2(1) | S | Define a origem do custocont na tranferência |
| 435 | SOMAVLFORANFCUSTOCONTTV10 | VARCHAR2(1) | S | Define se será somado os valores fora da nota, ao custocont ao realizar um TV10 ou TV 9 |
| 436 | CODCONTASTGUIAENT | NUMBER(10) | S | Cod.Conta para gerar despesa do ST Guia |
| 437 | CODCONTATFCPGUIAENT | NUMBER(10) | S | Cod.Conta para gerar despesa do FCP ST Guia |
| 438 | CODPRACA | NUMBER(6) | S | Código da praça que é cadastrado na rotina 572 |
| 439 | DEDUZICMSSTBCRSEMST | VARCHAR2(1) | S | Considera ICMS e ST BCR para calculo dos custos sem ST |
| 440 | GERARC500ENERGIA | VARCHAR2(1) | S | Gerar C500 com modelo 55  Sped Contribuições rotina 1097 |
| 441 | VOLTARESTOQUEFILIALVIRTUAL | VARCHAR2(1) | S | Indica se o  estoque irá retornar para filial de venda |
| 442 | VOLTARESTOQUEFILIALRETIRA | VARCHAR2(1) | S | Indica se o estoque irá retornar para filial retira |
| 443 | ZERARPISCOFINSIMPXML | VARCHAR2(1) | S | Zerar os valores do PIS e COFINS na importação de XML de produtos que são para consumo |
| 444 | IDEXTERNOCOLETAENTREGA | VARCHAR2(40) | S | Index de integração |
| 445 | DTULTALTOCOLETAENTREGA | DATE | S | Data Ultima Alto Coleta Eentrega |
| 446 | DTENTGACOLETAENTREGA | DATE | S | Data Entrega Coleta Entrega  |
| 447 | SITUACAOCOLETAENTREGA | VARCHAR2(40) | S | SITUACAO COLETA ENTREGA |
| 448 | REJEITADOSEFAZ | VARCHAR2(1) | S | Indica se o rejeição da NFe é Local ou da SEFAZ |
| 449 | DEDUZIRCAPATAZIABASEII | VARCHAR2(1) | S | DEDUZIR CAPATAZIA DA BASE II |
| 450 | DEDUZIRCAPATAZIABASEPISCOFINS | VARCHAR2(1) | S | DEDUZIR CAPATAZIA BASE PIS/COFINS |
| 451 | DEDUZIRICMSDIFBASEICMSATECIP | VARCHAR2(1) | S | Deduzir ICMS Diferido da Base do ICMS Antecipado |
| 452 | UTILIZATOTALPERSNFIMP | VARCHAR2(1) | S | Define se utiliza o cálculo personalizado para o valor total da nota de importação (tipo N). O valor total será gravado na movimentação conforme configuração destes parâmetros.  |
| 453 | VLTOTALPERSNFIMPDEDUCOES | VARCHAR2(1) | S | Define se o Vl. Deduções vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 454 | VLTOTALPERSNFIMPFRETE | VARCHAR2(1) | S | Define se o Vl. Frete vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 455 | VLTOTALPERSNFIMPSEGURO | VARCHAR2(1) | S | Define se o Vl. Seguro vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 456 | VLTOTALPERSNFIMPCAPATAZIA | VARCHAR2(1) | S | Define se o Vl. Capatazia vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 457 | VLTOTALPERSNFIMPSISCOMEX | VARCHAR2(1) | S | Define se o Vl. Siscomex vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 458 | VLTOTALPERSNFIMPDESPADUAN | VARCHAR2(1) | S | Define se o Vl. Despesas Aduaneiras vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 459 | VLTOTALPERSNFIMPAFRMM | VARCHAR2(1) | S | Define se o Vl. AFRMM vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 460 | VLTOTALPERSNFIMPANTIDAMPING | VARCHAR2(1) | S | Define se o Vl. Antidamping vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 461 | VLTOTALPERSNFIMPACRESCIMOS | VARCHAR2(1) | S | Define se o Vl. Acréscimos vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 462 | VLTOTALPERSNFIMPII | VARCHAR2(1) | S | Define se o Vl. II vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 463 | VLTOTALPERSNFIMPIPI | VARCHAR2(1) | S | Define se o Vl. IPI vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 464 | VLTOTALPERSNFIMPICMS | VARCHAR2(1) | S | Define se o Vl. ICMS vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 465 | VLTOTALPERSNFIMPST | VARCHAR2(1) | S | Define se o Vl. ST vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 466 | VLTOTALPERSNFIMPCREDPIS | VARCHAR2(1) | S | Define se o Vl. Crédito de PIS vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 467 | VLTOTALPERSNFIMPCREDCOFINS | VARCHAR2(1) | S | Define se o Vl. Crédito de COFINS vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 468 | VLTOTALPERSNFIMPFECOPICMS | VARCHAR2(1) | S | Define se o Vl. Fecop ICMS vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 469 | VLTOTALPERSNFIMPFECOPST | VARCHAR2(1) | S | Define se o Vl. Fecop ST vai ser considerado no cálculo do valor total em notas de importação (tipo N) |
| 470 | VLFRETENACIONAL | NUMBER(12,2) | S | Valor do Frente em Território Nacional |
| 471 | VLSEGURONACIONAL | NUMBER(12,2) | S | Valor do Seguro em Território Nacional |
| 472 | INFFRETESEGURONOTAN | VARCHAR2(1) | S | Informar Frete ou Seguro Nacional na Nota de Nacionalização |
| 473 | DEDUZIRCBSCUSTOREAL | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOREAL |
| 474 | DEDUZIRCBSCUSTOCONT | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOCONT |
| 475 | DEDUZIRCBSCUSTOREP | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOREP |
| 476 | DEDUZIRCBSCUSTOFIN | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOFIN |
| 477 | DEDUZIRCBSCUSTOREALSEMST | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOREALSEMST |
| 478 | DEDUZIRCBSCUSTOFINSEMST | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOFINSEMST |
| 479 | DEDUZIRCBSCUSTOREALLIQ | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOREALLIQ |
| 480 | DEDUZIRCBSCUSTOFISCAL | VARCHAR2(1) | S | DEDUZIR CBS PARA CALCULO DO CUSTOFISCAL |
| 481 | DEDUZIRIBSCUSTOREAL | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOREAL |
| 482 | DEDUZIRIBSCUSTOCONT | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOCONT |
| 483 | DEDUZIRIBSCUSTOREP | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOREP |
| 484 | DEDUZIRIBSCUSTOFIN | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOFIN |
| 485 | DEDUZIRIBSCUSTOREALSEMST | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOREALSEMST |
| 486 | DEDUZIRIBSCUSTOFINSEMST | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOFINSEMST |
| 487 | DEDUZIRIBSCUSTOREALLIQ | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOREALLIQ |
| 488 | DEDUZIRIBSCUSTOFISCAL | VARCHAR2(1) | S | DEDUZIR IBS PARA CALCULO DO CUSTOFISCAL |
| 489 | DEDUZIRISCUSTOREAL | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOREAL |
| 490 | DEDUZIRISCUSTOCONT | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOCONT |
| 491 | DEDUZIRISCUSTOREP | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOREP |
| 492 | DEDUZIRISCUSTOFIN | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOFIN |
| 493 | DEDUZIRISCUSTOREALSEMST | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOREALSEMST |
| 494 | DEDUZIRISCUSTOFINSEMST | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOFINSEMST |
| 495 | DEDUZIRISCUSTOREALLIQ | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOREALLIQ |
| 496 | DEDUZIRISCUSTOFISCAL | VARCHAR2(1) | S | DEDUZIR IS PARA CALCULO DO CUSTOFISCAL |
| 497 | VLTOTALBASEIS | NUMBER(18,6) | S | Total da base de cálculo do imposto seletivo |
| 498 | VLTOTALIS | NUMBER(18,6) | S | Total do imposto seletivo |
| 499 | SOMATOTALNF_IS | VARCHAR2(1) | S | Define se irá somar o valor do IS no valor total da nota S(Sim) ou N(Não) |
| 500 | TPENTEGOV | VARCHAR2(1) | S | Tipo de ente governamental |
| 501 | PREDUTOR | NUMBER(7,4) | S | Percentual de redução de alíquota em compra governamental |
| 502 | VLIBSMUN | NUMBER(15,2) | S | Valor total do IBS do Município |
| 503 | VLIBSUF | NUMBER(15,2) | S | Valor total do IBS da UF |
| 504 | VLCBS | NUMBER(15,2) | S | Valor total de CBS |
| 505 | VLIBS | NUMBER(15,2) | S | Valor imposto ibs geral |
| 506 | TIPONOTACREDDEB | VARCHAR2(2) | S | Destaque para tipo de Nota Débito e Crédito |
| 507 | VLTOTALIBS | NUMBER(15,2) | S | Valor total de IBS |
| 508 | VLTOTALCBS | NUMBER(15,2) | S | Valor total de CBS |
| 509 | VBCICBSCBS | NUMBER(15,2) | S | Valor total da BC do IBS e da CBS |
| 510 | VBCIBSCBS | NUMBER(13,2) | S | Valor total da BC do IBS e da CBS |
| 511 | TIPOORGAOPUB | VARCHAR2(2) | S | Tipo do Órgão Público. Preencher com: 0=Nenhum; 1=União; 2=Estado; 3=Distrito Federal ou 4=Município (Rotinas 202 e 302) |
| 512 | PERCREDORGAOPUB | NUMBER(12,4) | S | Percentual de redução da alíquota IBS, CBS e IS para Órgão Público (Rotinas 202 e 302) |
| 513 | TPOPERGOV | NUMBER(1) | S | Tipo de operação com o ente governamental |
| 514 | VLPISRETIDO | NUMBER(12,2) | S | Valor PIS Retido |
| 515 | VLCOFINSRETIDO | NUMBER(12,2) | S | Valor COFINS Retido |
| 516 | TIPOFATURAMENTONFSERVICO | VARCHAR2(1) | S | Tipo do faturamento NFs de Serviço |
| 517 | FINALIDADENFSERVICO | VARCHAR2(1) | S | Finalidade NFs de Serviço |
| 518 | CODMUNNFSERVICO | NUMBER(7) | S | Código do IBGE NFs de Serviço |
| 519 | NFCREDITO | VARCHAR2(1) | S | Nota de crédito |
| 520 | TIPONFCREDITO | VARCHAR2(2) | S | Tipo NF Crédito |
| 521 | SOMAICMSTV10 | VARCHAR2(1) | S | Somar o valor do ICMS no custo do produto na entrada TV10 |
| 522 | TRANSFTV10 | VARCHAR2(1) | S | EXIBE SE A ENTRADA É UMA TRANSFERÊNCIA TV10 |
| 523 | ENTRADAREMESSAFULFILLMENT | VARCHAR2(1) | S | Identifica se a nota pertence a uma remessa fulfillment |
| 524 | ESTORNODECREDITO | VARCHAR2(1) | S | Indica estorno Créditos IBS/CBS |
| 525 | PAGCHEQUEMORADIA | VARCHAR2(1) | S | Define se a devolução foi oriunda de NFe usando pagamento usando cheque moradia |
| 526 | CONSACRCUSTOCONTIMP | VARCHAR2(1) | S | Considerar Acréscimo Custo no custo médio contábil na importação |
| 527 | NUMNOTASERVICO | VARCHAR2(60) | S | Nº Nota Fiscal de Serviço |
| 528 | CHAVENFSE | VARCHAR2(60) | S | Chave Nota Fiscal de Serviço |
| 529 | NUMREGEXP | NUMBER(12) | S | Número do Registro de Exportação |
| 530 | ISUFEMIT | VARCHAR2(15) | S | Inscrição SUFRAMA do Emitente |
| 531 | IMPXML | VARCHAR2(1) | S | Indica se NFe foi gerada a partir de importação direta de um arquivo XML. |
| 532 | VERSAODOCFISCAL | VARCHAR2(30) | S | Versão do DocFiscal |

## Chaves e índices

- Constraint PCNFENT_PK (P): NUMTRANSENT, CODCONT
- Índice IDX$$_3ABD0002: ESPECIE, NUMTRANSENT, NUMNOTA
- Índice IDX$$_3ABD0004: ESPECIE, DTCANCEL, NUMTRANSENT, NUMNOTA
- Índice IDX$$_3ABD0007: NUMTRANSENT, NUMNOTA
- Índice PCNFENT_IDX1: NUMNOTA
- Índice PCNFENT_IDX10: SYS_NC00326$, DTENT, ESPECIE
- Índice PCNFENT_IDX1000: NUMNOTA, DTENT, SYS_NC00326$, ESPECIE, TIPODESCARGA, CHAVENFE, SYS_NC00328$
- Índice PCNFENT_IDX1000_1: DTENT, NUMNOTA
- Índice PCNFENT_IDX11: SITUACAONFE, ESPECIE
- Índice PCNFENT_IDX12: DTEMISSAO, SITUACAONFE, ESPECIE
- Índice PCNFENT_IDX13: CHAVENFE
- Índice PCNFENT_IDX14: RECIBONFE
- Índice PCNFENT_IDX15: RECIBOCTE
- Índice PCNFENT_IDX16: RECIBODPEC
- Índice PCNFENT_IDX17: SITUACAOCTE, SYS_NC00326$
- Índice PCNFENT_IDX18: NUMVIAS, SITUACAONFE, ESPECIE, DTCANCEL, DTSAIDA
- Índice PCNFENT_IDX19: NUMLOTENFE
- Índice PCNFENT_IDX2: DTENT
- Índice PCNFENT_IDX20: ESPECIE, SYS_NC00327$
- Índice PCNFENT_IDX21: CHAVECTE
- Índice PCNFENT_IDX22: DTENT, TEMPODESCARGA, CODFISCAL, CODCONT, OBS
- Índice PCNFENT_IDX24: NUMNOTAENTFUTURA, CODFORNEC
- Índice PCNFENT_IDX25: CODFILIAL, DTENT, CODFISCAL, SYS_NC00343$
- Índice PCNFENT_IDX26: DTENT, DTCANCEL, NUMTRANSENT, CHAVENFE
- Índice PCNFENT_IDX3: CODFORNEC
- Índice PCNFENT_IDX6: NUMBONUS
- Índice PCNFENT_IDX7: NUMTRANSENT, TIPODESCARGA
- Índice PCNFENT_IDX8: NUMTRANSVENDAORIG, TIPODESCARGA
- Índice PCNFENT_IDX9: NUMTRANSENT, CODFILIAL, DTWMS
- Índice PCNFENT_PK UNIQUE: NUMTRANSENT, CODCONT
