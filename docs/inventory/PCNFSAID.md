# RIBERDOCES.PCNFSAID — dicionário

Gerado por `gateway inventory` em 2026-10-25. 670 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | ESPECIE | VARCHAR2(2) | S |  |
| 2 | SERIE | VARCHAR2(5) | S |  |
| 3 | NUMNOTA | NUMBER(10) | S |  |
| 4 | DTSAIDA | DATE | N |  |
| 5 | VLTOTAL | NUMBER(12,2) | S | Valor total. |
| 6 | CODCONT | VARCHAR2(8) | N |  |
| 7 | CODFISCAL | NUMBER(8) | N |  |
| 8 | CODCLI | NUMBER(9) | N |  |
| 9 | CAIXA | NUMBER(8) | N | Caixa |
| 10 | CODUSUR | NUMBER(4) | N |  |
| 11 | DTENTREGA | DATE | S |  |
| 12 | COMISSAO | NUMBER(12,2) | S |  |
| 13 | VLTOTGER | NUMBER(12,2) | S | Valor total. |
| 14 | VLTABELA | NUMBER(12,2) | S |  |
| 15 | CODFILIAL | VARCHAR2(2) | N |  |
| 16 | VLDESCONTO | NUMBER(12,2) | S | Valor total do desconto. |
| 17 | TIPOVENDA | VARCHAR2(2) | N |  |
| 18 | OBS | VARCHAR2(120) | S |  |
| 19 | VLCUSTOREAL | NUMBER(14,2) | S |  |
| 20 | VLCUSTOFIN | NUMBER(14,2) | S |  |
| 21 | VLFRETE | NUMBER(14,2) | S |  |
| 22 | VLOUTRASDESP | NUMBER(14,2) | S |  |
| 23 | TOTPESO | NUMBER(18,6) | S |  |
| 24 | TOTVOLUME | NUMBER(18,6) | S |  |
| 25 | CODPRACA | NUMBER(6) | N |  |
| 26 | NUMITENS | NUMBER(4) | S |  |
| 27 | CODEMITENTE | NUMBER(8) | S |  |
| 28 | DTCANCEL | DATE | S |  |
| 29 | NUMCAR | NUMBER(8) | S |  |
| 30 | CODCOB | VARCHAR2(4) | S |  |
| 31 | NUMPED | NUMBER(10) | S |  |
| 32 | STRIBUT | NUMBER(14,2) | S |  |
| 33 | ICMSRETIDO | NUMBER(14,2) | S |  |
| 34 | BCST | NUMBER(14,2) | S |  |
| 35 | CODDEVOL | NUMBER(4) | S |  |
| 36 | NFFORNEC | NUMBER(10) | S |  |
| 37 | CODFORNEC | NUMBER(6) | S |  |
| 38 | CODDOC | VARCHAR2(4) | S |  |
| 39 | NUMTRANSVENDA | NUMBER(10) | N |  |
| 40 | CODPLPAG | NUMBER(4) | S |  |
| 41 | CODFUNCCANCEL | NUMBER(8) | S |  |
| 42 | NUMTAB | NUMBER(4) | S |  |
| 43 | NUMSEQ | NUMBER(20) | S |  |
| 44 | CODSUPERVISOR | NUMBER(4) | S |  |
| 45 | CONDVENDA | NUMBER(5) | S | Indica a condição da venda. |
| 46 | VLDEVOLUCAO | NUMBER(14,2) | S |  |
| 47 | VLBASEIPI | NUMBER(14,2) | S |  |
| 48 | VLIPI | NUMBER(12,2) | S |  |
| 49 | VLICMSANTECIPADO | NUMBER(12,2) | S |  |
| 50 | CODEPTO | NUMBER(6) | S |  |
| 51 | VLTOTALNF | NUMBER(12,2) | S |  |
| 52 | CODFILIALNF | VARCHAR2(2) | S |  |
| 53 | UF | VARCHAR2(2) | S | UF cliente. |
| 54 | VLBASE | NUMBER(12,2) | S |  |
| 55 | VLICMS | NUMBER(12,2) | S |  |
| 56 | VLISENTAS | NUMBER(12,2) | S |  |
| 57 | VLOUTRAS | NUMBER(12,2) | S |  |
| 58 | NUMIDF | VARCHAR2(30) | S |  |
| 59 | NUMFORMULARIO | NUMBER(10) | S |  |
| 60 | VLCUSTOREP | NUMBER(14,2) | S |  |
| 61 | VLCUSTOCONT | NUMBER(14,2) | S |  |
| 62 | DTDEVOL | DATE | S |  |
| 63 | NUMSELONF | NUMBER(10) | S |  |
| 64 | NUMSEQMONTAGEM | NUMBER(4) | S |  |
| 65 | NUMVOLUME | NUMBER(8) | S |  |
| 66 | NUMCUPOM | NUMBER(10) | S |  |
| 67 | VLRECEBIDODIN | NUMBER(10,2) | S |  |
| 68 | VLRECEBIDOCH | NUMBER(10,2) | S |  |
| 69 | VLRECEBIDOBOLETO | NUMBER(10,2) | S |  |
| 70 | VLTROCO | NUMBER(10,2) | S |  |
| 71 | CODFUNCLIBENT | NUMBER(8) | S |  |
| 72 | DTLIBENT | DATE | S |  |
| 73 | VLDESCPROM | NUMBER(12,2) | S |  |
| 74 | NUMVIAS | NUMBER(4) | S |  |
| 75 | CODCONTRATO | NUMBER(6) | S |  |
| 76 | FRETEDESPACHO | VARCHAR2(1) | S |  |
| 77 | FRETEREDESPACHO | VARCHAR2(1) | S |  |
| 78 | CODFORNECFRETE | NUMBER(6) | S |  |
| 79 | VLRECEBIDOCARTAO | NUMBER(12,2) | S |  |
| 80 | NUMTRANSVENDAORIGEM | NUMBER(10) | S |  |
| 81 | CODFUNCLANC | NUMBER(8) | S |  |
| 82 | DTCANHOTO | DATE | S |  |
| 83 | NUMTRANSVENDADESTINO | NUMBER(10) | S |  |
| 84 | CODUSUR3 | NUMBER(4) | S |  |
| 85 | CODUSUR2 | NUMBER(4) | S |  |
| 86 | HORALANC | VARCHAR2(2) | S |  |
| 87 | MINUTOLANC | VARCHAR2(2) | S |  |
| 88 | ROTINALANC | VARCHAR2(48) | S |  |
| 89 | EQUIPLANC | VARCHAR2(64) | S |  |
| 90 | FUNCLANC | VARCHAR2(30) | S |  |
| 91 | DTWMS | DATE | S |  |
| 92 | PRAZOMEDIO | NUMBER(4) | S |  |
| 93 | PRAZOADICIONAL | NUMBER(4) | S |  |
| 94 | VLFRETENF | NUMBER(12,2) | S |  |
| 95 | CODSUPERVISOR2 | NUMBER(4) | S |  |
| 96 | CODSUPERVISOR3 | NUMBER(4) | S |  |
| 97 | COMISSAO2 | NUMBER(12,2) | S |  |
| 98 | COMISSAO3 | NUMBER(12,2) | S |  |
| 99 | VLISS | NUMBER(18,6) | S |  |
| 100 | NUMLANCCP | NUMBER(8) | S |  |
| 101 | PRAZOPONDERADO | VARCHAR2(1) | S |  |
| 102 | NUMNFTRANSF | NUMBER(10) | S |  |
| 103 | SERIEECF | VARCHAR2(2) | S |  |
| 104 | NUMSERIEEQUIP | VARCHAR2(30) | S |  |
| 105 | VLBONIFIC | NUMBER(16,3) | S |  |
| 106 | BCSTGNRE | NUMBER(14,2) | S |  |
| 107 | ICMSRETIDOGNRE | NUMBER(14,2) | S |  |
| 108 | TIPODEVOL | VARCHAR2(2) | S |  |
| 109 | CODCLINF | NUMBER(9) | S | Código do cliente para o endereço de entraga da NF. |
| 110 | VLIR | NUMBER(18,6) | S |  |
| 111 | PERCFRETE | NUMBER(8,4) | S |  |
| 112 | NUMREGIAOFRETE | NUMBER(4) | S |  |
| 113 | CODCONDICAOVENDA | NUMBER(6) | S | Indica o código da condição de venda. |
| 114 | TOTVLDESCFIN | NUMBER(18,6) | S |  |
| 115 | TOTVLDESCCOM | NUMBER(18,6) | S |  |
| 116 | TOTVLBONIFIC | NUMBER(18,6) | S |  |
| 117 | TOTVLDESCFLEX | NUMBER(18,6) | S |  |
| 118 | TOTVLREDCOMISS | NUMBER(18,6) | S |  |
| 119 | NUMVIASPU | NUMBER(4) | S |  |
| 120 | NUMORCA | NUMBER(10) | S |  |
| 121 | COMISSAOSUP | NUMBER(18,6) | S |  |
| 122 | TOTVLREDCOMISSSUP | NUMBER(18,6) | S |  |
| 123 | DTEMISSAONFFORNEC | DATE | S |  |
| 124 | TIPOFRETE | VARCHAR2(1) | S |  |
| 125 | BROKER | VARCHAR2(1) | S |  |
| 126 | PERBASEREDOUTRASDESP | NUMBER(8,2) | S | Contém o percentual de redução para despesas acessórias (se houver). |
| 127 | NUMCAIXAFISCAL | NUMBER(4) | S |  |
| 128 | DTHORACHEGADA | DATE | S |  |
| 129 | DTHORASAIDA | DATE | S |  |
| 130 | KMFINALENTREGA | NUMBER(12,2) | S |  |
| 131 | VLPIS | NUMBER(18,6) | S | Valor do PIS a ser lançado na Geração de Nota Fiscal de Prestação de Serviço (rotina 1446). |
| 132 | PERDESCFIN | NUMBER(12,6) | S | Mesmo campo gravado no pedido, que é replicado durante o faturamento. |
| 133 | CODCLIRECEBEDOR | NUMBER(9) | S | Indica o cliente recebedor da mercadoria em pedidos feitos para redes de clientes. |
| 134 | NUMREGIAO | NUMBER(4) | S | Região de preço usada na nota fiscal. |
| 135 | NUMORCAFILIAL | NUMBER(10) | S | Número Orçamento da Filial. |
| 136 | APRESENTOUCUPOMPROMOCAO | VARCHAR2(1) | S |  |
| 137 | NSU | NUMBER(15) | S | NSU (número sequencial único) que deve ser gerado e impresso em todas as notas fiscais (entrada e saída). |
| 138 | DTNSU | DATE | S | Data e hora de geração do NSU (número sequencial único), que deve ser impresso em todas as notas fiscais (entrada e saída). |
| 139 | RESERVARITENSTV7 | VARCHAR2(1) | S |  |
| 140 | NUMPREVENDA | NUMBER(10) | S |  |
| 141 | GERACAORAZAOAUXILIAR | VARCHAR2(1) | S | Indica que já foi importado na geração do Razão Auxiliar (rotina 1050). |
| 142 | DTPGTOCAMPANHA | DATE | S |  |
| 143 | NUMECF | NUMBER(10) | S |  |
| 144 | GERACP | VARCHAR2(1) | S | Indica se pedido TV10 irá gerar Contas a Pagar. \|Campo do tipo caracter, de tamanho 1, com valor padrão "N |
| 145 | VLVERBACMVCLI | NUMBER(18,6) | S | Indica o valor total da verba CMV registrado no pedido. |
| 146 | DTEXPORTACAOSERVINT | DATE | S |  |
| 147 | EXPORTADOSERVINT | VARCHAR2(1) | S |  |
| 148 | IMPORTADOSERVPRINC | VARCHAR2(1) | S |  |
| 149 | DTIMPORTACAOSERVPRINC | DATE | S |  |
| 150 | NUMTRANSVENDAECF | NUMBER(10) | S |  |
| 151 | CHAVENFE | VARCHAR2(45) | S |  |
| 152 | SITUACAONFE | NUMBER(10) | S |  |
| 153 | PROTOCOLONFE | VARCHAR2(20) | S | Descricao coluna PROTOCOLONFE |
| 154 | DTHORANFE | DATE | S |  |
| 155 | IEENT | VARCHAR2(15) | S |  |
| 156 | NUMLOTENFE | VARCHAR2(15) | S | Número de Lote da Nota Fiscal Eletrônica. |
| 157 | PERCCOMPROFISSIONAL | NUMBER(8,4) | S | Indica o percentual de comissao do profissional cadastrado na rotina 340. |
| 158 | TIPOEMISSAO | VARCHAR2(1) | S | Tipo de emissão da NFE (Normal, Contingência, SCAN). Campo do tipo caracter, de tamanho 1, com valor default 1 |
| 159 | ENVIADA | VARCHAR2(1) | S | Indica se em caso de contingência, já enviou para o SEFAZ. |
| 160 | USAINTEGRACAOWMS | VARCHAR2(1) | S | Usa integração WMS |
| 161 | ENVIADANFE | VARCHAR2(1) | S | NF-e já enviada ao cliente. |
| 162 | AUTENTPGTOGNRE | VARCHAR2(200) | S | Autenticação bancaria. |
| 163 | BANCOPGTOGNRE | VARCHAR2(50) | S | Banco de pagto. |
| 164 | AGENCIAPGTOGNRE | VARCHAR2(30) | S | Agencia de pagto. |
| 165 | VENDAASSISTIDA | VARCHAR2(1) | S | Indica se a NF é de venda assistida. |
| 166 | CODEMITENTEPEDIDO | NUMBER(8) | S | Código do emitente do pedido. |
| 167 | NUMOP | NUMBER(8) | S | Indica o numero da ordem de produção. |
| 168 | ALIQUOTA | NUMBER(7,2) | S | Indica a alíquota de ICMS.\|Campo do tipo numérico, de tamanho 7, com 2 casas decimais. |
| 169 | CODCONSIGNATARIOFRETE | NUMBER(8) | S | Indica o código do consignatário do conhecimento de frete. |
| 170 | CODEXPORTADOR | NUMBER(6) | S | Indica o código de exportador. |
| 171 | UFDESEMBARACO | VARCHAR2(2) | S | Indica a UF desembaraço. |
| 172 | LOCALDESEMBARACO | VARCHAR2(30) | S | Indica o local desembaraço. |
| 173 | CODREMETENTEFRETE | NUMBER(9) | S | Indica o código do remetente do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 174 | CODDESTINATARIOFRETE | NUMBER(9) | S | Indica o código do destinatário do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 175 | CODVEICULO | NUMBER(4) | S | Indica o código do veículo do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 176 | CODMOTORISTA | NUMBER(8) | S | Indica o código do motorista do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 177 | CODPRACADESTFRETE | NUMBER(6) | S | Indica o código do frete (rota de entrega) do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 178 | OBSCONHECFRETE | VARCHAR2(400) | S | Indica a observação do conhecimento de frete.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 179 | VLSECCAT | NUMBER(18,6) | S | Indica o valor de despesa com escolta.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 180 | VLDESPACHO | NUMBER(18,6) | S | Indica o valor de despesa de despacho. (Utilizado apenas para o projeto 50507 - Transportadora) |
| 181 | VLLTR | NUMBER(18,6) | S | Indica o valor de despesa de LTR. (Utilizado apenas para o projeto 50507 - Transportadora) |
| 182 | VLPEDAGIO | NUMBER(18,6) | S | Indica o valor de despesa do pedágio.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 183 | VLSEGURO | NUMBER(18,6) | S | Indica o valor de despesa de seguro da carga.(Utilizado apenas para o projeto 50507 - Transportadora) |
| 184 | TENTATIVAENVIONFE | NUMBER(4) | S | Indica a quantidade de tentativas de envio da nota fiscal a SEFAZ. |
| 185 | CLIENTE | VARCHAR2(60) | S | Razão social do cliente. |
| 186 | CGC | VARCHAR2(18) | S | CNPJ do cliente. |
| 187 | IE | VARCHAR2(15) | S | IE do cliente. |
| 188 | TIPOFJ | VARCHAR2(1) | S | Tipo pessoa (em casos de dev. de clientes). |
| 189 | CONSUMIDORFINAL | VARCHAR2(1) | S | Indica se cliente é consumidor final. |
| 190 | CODATV1 | NUMBER(6) | S | Cód.Atividade do cliente. |
| 191 | CODPAIS | NUMBER(6) | S | Código do País. |
| 192 | DESCPAIS | VARCHAR2(60) | S | Descrição do País. |
| 193 | CGCFILIAL | VARCHAR2(14) | S | CNPJ da filial. |
| 194 | IEFILIAL | VARCHAR2(20) | S | IE da filial. |
| 195 | UFFILIAL | VARCHAR2(2) | S | UF da filial. |
| 196 | AGREGARSTPRODSINTEGRA | VARCHAR2(1) | S | Agregar ou não o ST ao valor dos itens (Sintegra). |
| 197 | TIPOALIQOUTRASDESP | VARCHAR2(1) | S | Tipo de tributação do frete e despesas acessórias. |
| 198 | IESUBSTTRIBUT | VARCHAR2(20) | S | IE como substituto tributário. |
| 199 | UFCODIGO | VARCHAR2(2) | S | Código da UF. |
| 200 | HISTORICO | VARCHAR2(1) | S | Gerou histórico para o registro, sim ou não. |
| 201 | SUBSERIE | VARCHAR2(2) | S | Indica a subserie da nota fiscal. |
| 202 | ESPECIEVOLUME | VARCHAR2(100) | S | Indica a especie dos volumes da nota fiscal. |
| 203 | MARCARVOLUME | VARCHAR2(100) | S | Indica a marca dos produtos da nota fiscal. |
| 204 | PROTOCOLONFECAN | VARCHAR2(20) | S | Indica o protocolo de cancelamento da NF-e. |
| 205 | PROTOCOLONFEAUTO | VARCHAR2(20) | S | Indica o protocolo de autorizaçãoda NF-e. |
| 206 | COMISSAOMOT | NUMBER(12,2) | S | Indica o valor de comissão para o motorista. |
| 207 | DTPAGCOMISSAOPROF | DATE | S | Indica quando a comissão foi paga para o profissional. |
| 208 | CODUSUR4 | NUMBER(6) | S | Indica o código do terceiro profissional. |
| 209 | COMISSAO4 | NUMBER(12,2) | S | Indica a comissão do terceiro profissional. |
| 210 | CLIENTEFONTEST | VARCHAR2(1) | S | Indica o cliente fonte ST. |
| 211 | ORGAOPUB | VARCHAR2(1) | S | Indica orgão público. |
| 212 | TRANSPORTADORA | VARCHAR2(60) | S | Indica a transportadora. |
| 213 | CODCLIFILIAL | NUMBER(9) | S | Indica o código do cliente para a filial. |
| 214 | CGCFRETE | VARCHAR2(18) | S | CNPJ do transportador. |
| 215 | IEFRETE | VARCHAR2(15) | S | IE do transportador. |
| 216 | UFFRETE | VARCHAR2(2) | S | UF do transportador. |
| 217 | VLPAUTAFRETE | NUMBER(12,2) | S | Vl.Pauta do frete na nota. |
| 218 | COBRANCA | VARCHAR2(30) | S | Cobrança utilizada para a nota. |
| 219 | ENDERECO | VARCHAR2(40) | S | Endereço do cliente. |
| 220 | BAIRRO | VARCHAR2(40) | S | Bairro do cliente. |
| 221 | MUNICIPIO | VARCHAR2(30) | S | Municpio do cliente. |
| 222 | CEP | VARCHAR2(10) | S | CEP do cliente. |
| 223 | SULFRAMA | VARCHAR2(15) | S | Incrição Suframa do cliente. |
| 224 | CODMUNICIPIO | NUMBER(10) | S | Cód.Municipio do cliente. |
| 225 | PLACAVEIC | VARCHAR2(10) | S | Placa do veículo utilizado para transporte. |
| 226 | CODCONTCLI | NUMBER(10) | S | Código da conta de cliente. |
| 227 | CODIBGE | NUMBER(10) | S | Código da cidade no IBGE. |
| 228 | TIPOEMPRESA | VARCHAR2(4) | S | Tipo da empresa do cliente. |
| 229 | CODFISCALFRETE | NUMBER(4) | S | CFOP do frete na nota fiscal. |
| 230 | PERCICMFRETE | NUMBER(4,2) | S | Aliq.ICMS do frete na nota fiscal. |
| 231 | CODFISCALOUTRASDESP | NUMBER(4) | S | CFOP da despesa na nota fiscal. |
| 232 | ALIQICMOUTRASDESP | NUMBER(4,2) | S | Aliq.ICMS da despesa na nota fiscal. |
| 233 | ATACADISTA | VARCHAR2(1) | S | Indica se o cliente é atacadista. |
| 234 | VLINSS | NUMBER(18,6) | S | Indica o valor do INSS. |
| 235 | NUMCARANTERIOR | NUMBER(8) | S | Indica o numero do carregamento anterior. |
| 236 | DTFECHACCRCA | DATE | S | Indica se foi lançado saldo no contacorrente do RCA para a NF. |
| 237 | CONTRIBUINTE | VARCHAR2(1) | S | Indica se o e contribuinte ou não. |
| 238 | DTCANCELWMS | DATE | S | Indica se o processo do WMS foi cancelado. |
| 239 | VLCOFINS | NUMBER(18,6) | S | Indica o valor total de crédito de COFINS na NF. |
| 240 | CODFISCALNF | NUMBER(8) | S | Indica o código fiscal da NF. |
| 241 | DTLANCTO | DATE | S | Indica a data da inclusão do lançamento. |
| 242 | RECARGA | VARCHAR2(1) | S | Indica pedido de recarga. |
| 243 | NUMCONTRATO | NUMBER(8) | S | Indica o número do contrato. |
| 244 | DTHORAAUTORIZACAOSEFAZ | DATE | S | Indica a hora que foi autorizado no sefaz. |
| 245 | RECIBONFE | VARCHAR2(20) | S | Indica o recibo de envio de lote. |
| 246 | LOTEDPEC | VARCHAR2(20) | S | Indica o número do lote dpec. |
| 247 | RECIBODPEC | VARCHAR2(20) | S | Indica o recibo de envio de lote dpec. |
| 248 | AMBIENTE | VARCHAR2(1) | S | Indica o ambiente em que foi enviada a nota. |
| 249 | PROTOCOLOCANCELAMENTO | VARCHAR2(20) | S | Indica o número do protocolo de cancelamento. |
| 250 | DTHORACANCELAMENTOSEFAZ | DATE | S | Indica a data e hora em que foi cancelada no sefaz. |
| 251 | DTFECHA | DATE | S | Data de fechamento da carga. |
| 252 | VLDESPDENTRONF | NUMBER(14,2) | S | Indica o Valor outras despesas na NF. |
| 253 | DTHORARECIBODPEC | DATE | S | Data e hora do recibo DPEC. |
| 254 | NUMTRANSVENDACONHEC | NUMBER(14,2) | S | Número de transação do conhecimento de frete. |
| 255 | OBSNFE | CLOB | S | Observações da NF-e. |
| 256 | TOTDIFALIQUOTAS | NUMBER(14,2) | S | Total do diferencial de alíquotas. |
| 257 | TOTPESOLIQ | NUMBER(18,6) | S | Peso liquido dos produtos. |
| 258 | TOTPESOAGRUPADO | NUMBER(18,6) | S | Peso total agrupado. |
| 259 | TOTPESOLIQAGRUPADO | NUMBER(18,6) | S | Peso total liquido agrupado. |
| 260 | TOTVOLUMEAGRUPADO | NUMBER(18,6) | S | Volume total agrupado. |
| 261 | DTLIBERACAOENTREGA | DATE | S | Data de liberação de entrega do pedido na expedição. |
| 262 | CODFUNCLIBERACAOENTREGA | NUMBER(8) | S | Código do funcionário que efetuou liberação de entrega do pedido na expedição. |
| 263 | HORAEMISSAO | VARCHAR2(8) | S | Hora da emissão da venda. |
| 264 | CONTADORCUPOMFISCAL | NUMBER(6) | S | Contador cupom fiscal. |
| 265 | VENDATRIANGULAR | VARCHAR2(1) | S | Venda Triangular. |
| 266 | SITUACAOCTE | NUMBER(10) | S | Situaçao do cte |
| 267 | RECIBOCTE | VARCHAR2(20) | S | Recibo cte |
| 268 | PROTOCOLOCTE | VARCHAR2(20) | S | Protocolo do cte |
| 269 | CHAVECTE | VARCHAR2(44) | S | Chave do cte |
| 270 | NUMLOTECTE | VARCHAR2(15) | S | Numero do lote do cte |
| 271 | DTA_HORAENVIOSEFAZ | DATE | S | Data e hora do envio do cte ao sefaz |
| 272 | AMBIENTECTE | VARCHAR2(1) | S | Ambiente do cte |
| 273 | AMBIENTENFE | VARCHAR2(1) | S | Ambiente da nfe |
| 274 | NUMSEQENTREGA | NUMBER(20) | S | Número de sequência de entrega |
| 275 | CODPRODPREDOMINANTE | NUMBER(6) | S | Código produto predominante conhecimento de frete eletrônico. |
| 276 | VLFRETERETIDO | NUMBER(14,2) | S | Valor de frete retido. |
| 277 | VLSTFRETERETIDO | NUMBER(14,2) | S | Valor ST de frete retido. |
| 278 | PERCFRETERETIDO | NUMBER(8,4) | S | Percentual de frete retido. |
| 279 | PERCSTFRETERETIDO | NUMBER(8,4) | S | Percentual de ST do frete retido. |
| 280 | PRAZO1 | NUMBER(4) | S | Prazo de pagamento em dias da primeira parcela. |
| 281 | PRAZO2 | NUMBER(4) | S | Prazo de pagamento em dias da segunda parcela. |
| 282 | PRAZO3 | NUMBER(4) | S | Prazo de pagamento em dias da terceira parcela. |
| 283 | PRAZO4 | NUMBER(4) | S | Prazo de pagamento em dias da quarta parcela. |
| 284 | PRAZO5 | NUMBER(4) | S | Prazo de pagamento em dias da quinta parcela. |
| 285 | PRAZO6 | NUMBER(4) | S | Prazo de pagamento em dias da sexta parcela. |
| 286 | PRAZO7 | NUMBER(4) | S | Prazo de pagamento em dias da sétima parcela. |
| 287 | PRAZO8 | NUMBER(4) | S | Prazo de pagamento em dias da oitava parcela. |
| 288 | PRAZO9 | NUMBER(4) | S | Prazo de pagamento em dias da nona parcela. |
| 289 | PRAZO10 | NUMBER(4) | S | Prazo de pagamento em dias da décima parcela. |
| 290 | PRAZO11 | NUMBER(4) | S | Prazo de pagamento em dias da décima primeira parcela. |
| 291 | PRAZO12 | NUMBER(4) | S | Prazo de pagamento em dias da décima segunda parcela. |
| 292 | TOTPESOBRUTO | NUMBER(18,6) | S | Total do peso bruto das mercadorias. |
| 293 | DTSAIDANF | DATE | S | Data de emissão da Nota Fiscal. |
| 294 | CUSTOBONIFICACAO | VARCHAR2(1) | S | Custo da bonificação TV5. |
| 295 | CODFORNECBONIFIC | NUMBER(6) | S | Código do Fornecedor da Bonificação TV5. |
| 296 | CODBNF | NUMBER(4) | S | Código do tipo de bonificação TV5. |
| 297 | JUSTIFICATIVACONTIGENCIA | VARCHAR2(256) | S | Motivo da entrada de contigência. |
| 298 | DTAHORAENTRADACONTIGENCIA | DATE | S | Data e Hora da entrada de contigência. |
| 299 | CONFERIDO | VARCHAR2(1) | S | Nota Fiscal Conferida. |
| 300 | CNPJCISP | VARCHAR2(8) | S | CNPJ para CISP. |
| 301 | CONSUMIUNUMNFE | VARCHAR2(1) | S | Se já tem um número de Nfe para o registro. |
| 302 | DTHORACTE | DATE | S | Data e Hora do CTE. |
| 303 | TENTATIVAENVIOCTE | NUMBER(10) | S | Nro de Tentativas de Envio do CTE. |
| 304 | PROTOCOLOCANCELAMENTOCTE | VARCHAR2(20) | S | Protocolo de cancelamento do CTE. |
| 305 | CODFORNECREDESPACHO | NUMBER(6) | S | Código fornecedore frete redespacho. |
| 306 | EMAILFRETE | VARCHAR2(100) | S | E-mail fornecedor frete. |
| 307 | EMAILFRETEREDESPACHO | VARCHAR2(100) | S | E-mail fornecedor frete redespacho. |
| 308 | DATAEMPENHO | DATE | S | Data de Empenho. |
| 309 | CODUNIDADEEXECUTORA | VARCHAR2(20) | S | Código da unidade executora. |
| 310 | VLDESCICMS | NUMBER(14,2) | S | Valor de desconto do ICMS de isenção. |
| 311 | VLTOTALSEMDESCICMS | NUMBER(12,2) | S | Valor total da NF sem desconto de Isenção de ICMS. |
| 312 | NUMEMPENHO | VARCHAR2(40) | S | Número de empenho do Órgão Púbico. |
| 313 | OBSEXP | VARCHAR2(200) | S | Obs. De ocorrência de entrega de nota fiscal. |
| 314 | CONTAORDEM | VARCHAR2(1) | S | Indica que a venda TV7, TV8 é do tipo conta e ordem. |
| 315 | VALORST | NUMBER(12,2) | S | Valor de ST da nota fiscal |
| 316 | MENSAGEMPISCOFINS | VARCHAR2(4000) | S | Mensagens da tributação PIS/COFINS |
| 317 | OBSPISCOFINS | VARCHAR2(500) | S | Observação da tributação de PIS/COFINS |
| 318 | VLBASEPISCOFINS | NUMBER(12,2) | S | Valor Base PIS/COFINS |
| 319 | PERPIS | NUMBER(8,4) | S | Percentual PIS |
| 320 | PERCOFINS | NUMBER(8,4) | S | Percentual COFINS |
| 321 | CODSITTRIBPISCOFINS | NUMBER(3) | S | Código Situação Tributário PIS COFINS |
| 322 | SITDOC | VARCHAR2(2) | S | Situação Documento |
| 323 | FINALIDADENFE | VARCHAR2(2) | S | Finalidade NF-e |
| 324 | SERIEORIGINAL | VARCHAR2(3) | S | SERIE ORIGINAL DA NOTA FISCAL ELETRONICA |
| 325 | DTEXPORTACAOWMS | DATE | S | Data  e hora de exportação |
| 326 | DTIMPORTACAOWMS | DATE | S | Data  e hora de importação |
| 327 | NFATUALWMS | VARCHAR2(1) | S | Atualização Nota Fiscal no WMAS Store. |
| 328 | EMAILDEST | VARCHAR2(100) | S | E-mail do destinatário |
| 329 | TELEFONE | VARCHAR2(20) | S | Telefone do destinatário |
| 330 | NUMENDERECO | VARCHAR2(6) | S | Número do Logradouro do destinatário |
| 331 | UFPLACAVEIC | VARCHAR2(2) | S | Gravar a UF da Placa do Veiculo |
| 332 | DESPESASRATEADA | VARCHAR2(1) | S | Informar se o frete e outras despesas foram rateados nos itens. |
| 333 | NUMTRANSENTORIGEM | NUMBER(10) | S | Campo que grava o número da transação de entrada associadas a nota fiscal de saída. |
| 334 | FORNECENTREGA | VARCHAR2(1) | S | Campo que grava a venda do cliente como venda com entrega pelo fornecedor. |
| 335 | NUMTRANSENTDEST | NUMBER(10) | S | Gravar o NUMTRANSENT da nota de destino |
| 336 | WMSCODFUNCIGNORA | NUMBER(8) | S | Funcionário que ignorou WMS |
| 337 | WMSIGNORADO | VARCHAR2(1) | S | WMS Ignorado |
| 338 | PISRETIDO | VARCHAR2(1) | S | Identificar se foi retido PIS na nota fiscal de saída |
| 339 | INSSRETIDO | VARCHAR2(1) | S | Identificar se foi retido INSS na nota fiscal de saída |
| 340 | IRRRETIDO | VARCHAR2(1) | S | Identificar se foi retido IRR na nota fiscal de saída |
| 341 | MOTORISTAVEICULO | VARCHAR2(40) | S | Motorista Veículo |
| 342 | CONSUMIUNUMCTE | VARCHAR2(1) | S | Consumiu número de cte |
| 343 | UTILIZAIPICALCICMS | VARCHAR2(1) | S | Utiliza IPI no calculo de ICMS |
| 344 | NUMOS | NUMBER(8) | S | Campo para armazenar o número da ordem de serviço que gerou a NF. |
| 345 | DTFAT | DATE | S | Data de faturamento |
| 346 | CODGERENTE | NUMBER(4) | S | Código do Gerente do RCA |
| 347 | NUMPEDCLI | VARCHAR2(15) | S | Número do pedido do cliente |
| 348 | DTDENEGADA | DATE | S | Data denegada |
| 349 | ENTREGA | VARCHAR2(1) | S | Indica se o item será entregue ao invés de ser retirando em venda balcão. |
| 350 | TAXAENTREGA | NUMBER(14,2) | S | Valor cobrado para realizar a entrega da compra. |
| 351 | UFPLACAVEI | VARCHAR2(2) | S | Uf da Placa do Veiculo |
| 352 | NORMAREGESPECIAL | VARCHAR2(250) | S | Norma Especifica (Regime Especial) |
| 353 | TIPOEMISSAOCTE | VARCHAR2(2) | S | Tipo Emissão Ct-e |
| 354 | DTEMISSAODECLARACAO | DATE | S | Data de emissão da declaração emitida pelo tomador do Ct-e, quando este não for contribuinte |
| 355 | PROTOCOLO | VARCHAR2(50) | S | Número do protocolo |
| 356 | CODVEICULO1 | NUMBER(4) | S | Veículo 1 |
| 357 | CODVEICULO2 | NUMBER(4) | S | Veículo 2 |
| 358 | COFINSRETIDO | VARCHAR2(1) | S | Identificar se foi retido cofins na nota fiscal de saída |
| 359 | CSLLRETIDO | VARCHAR2(1) | S | Identificar se foi retido csll na nota fiscal de saída |
| 360 | GERARBCRNFE | VARCHAR2(1) | S | Inserir os valores Base e ST BCR no XML da NF-e |
| 361 | NUMCARGAEDI | NUMBER(8) | S | Número Carga EDI |
| 362 | DEPOSITOFECHADO | VARCHAR2(1) | S | Identificação de deposito fechado |
| 363 | OBSNFCARREG | VARCHAR2(255) | S | OBSERVAÇÕES NA NOTA FISCAL REFERENTE A ENTREGA DA CARGA. |
| 364 | DEVSIMBOLICA | VARCHAR2(1) | S | Esse atributo serve para identificar se a operação, é uma devolução simbólica ou não. |
| 365 | VLIPIOUTRAS | NUMBER(18,6) | S | Esse atributo serve para armazenar o valor do ipi para devolução simbolica. |
| 366 | VLDESCABATIMENTO | NUMBER(18,6) | S | GRAVAR O VALOR DO DESCONTO DO ABATIMENTO |
| 367 | GERANFVENDA | VARCHAR2(1) | S | Informa se o número da nota foi gerado automatico. |
| 368 | ENVIADOEMAILCANCELADO | VARCHAR2(1) | S | ENVIADO EMAIL CANCELADO |
| 369 | NUMCONTRATOCARRETEIRO | VARCHAR2(20) | S | Número do Contrato Carreteiro |
| 370 | NUMLISTA | NUMBER(6) | S | Número da Lista de Presente |
| 371 | MOTESTORNONFE | VARCHAR2(250) | S | Motivo de estorno Nfe |
| 372 | NATOPERNFE | VARCHAR2(60) | S | Nat. Operação Nfe |
| 373 | VENDAEXPORTACAO | VARCHAR2(1) | S | Venda de exportação |
| 374 | IDINTEGRACAOMYFROTA | RAW(16) | S | Identifica a integração com My Frota. |
| 375 | NUMTRANSVENDATV13 | NUMBER(10) | S | Número de transação da nota fiscal tipo 13. |
| 376 | SITUACAONFEORIG | NUMBER(10) | S | Situação NFe Original. |
| 377 | CODMEDICOPRESCRIT | NUMBER(6) | S | Código do Médico Prescritor. |
| 378 | NUMPEDVANXML | VARCHAR2(20) | S | Num. Pedido VAN para gravar no XML da NF-e. |
| 379 | CNPJFORNEC | VARCHAR2(18) | S | CNPF do Fornecedor. |
| 380 | NUMCOMPROVANTE | NUMBER(20) | S | Número do comprovante do pagamento do pedagio. |
| 381 | CNPJRESPONSAVEL | VARCHAR2(18) | S | CNPF do Responsavel. |
| 382 | INUTILIZADA | VARCHAR2(1) | S | INFORMA SE FOI INUTILIZAÇÃO DE NUMERACAO. |
| 383 | AGREGASTVLMERC | VARCHAR2(1) | S | Agregar valor de ST ao valor da mercadoria. |
| 384 | RESPSEGURO | NUMBER(1) | S | RESPONSÁVEL PELO SEGURO |
| 385 | NOMESEGURADORA | VARCHAR2(30) | S | NOME DA SEGURADORA |
| 386 | NUMEROAPOLICE | VARCHAR2(20) | S | NUMERO DA APOLICE DO SEGURO |
| 387 | NUMEROAVERBACAO | VARCHAR2(40) | S | NUMERO DA AVERBAÇÃO DO SEGURO |
| 388 | SIMPLESNACIONAL | VARCHAR2(1) | S | Informar se o cliente é optante do simples |
| 389 | VLRAVERBACAO | NUMBER(13,2) | S | Valor da Averbação. |
| 390 | VLTRIBUTOS | NUMBER(14,2) | S | Guarda o Valor do Conhec. Frete X Perc. |
| 391 | VERSAOROTINA | VARCHAR2(20) | S | Versão da rotina gerarou os dados |
| 392 | CARTAODOTZ | VARCHAR2(18) | S | Número do CPF do cliente ou Número do cartão de fidelidade DOTZ. |
| 393 | CNPJCOLETA | VARCHAR2(18) | S | CNPJ do local de Coleta do CTE |
| 394 | RAZAOSOCIALCOLETA | VARCHAR2(60) | S | Razão Social do local de Coleta do CTE |
| 395 | LOGRADOUROCOLETA | VARCHAR2(60) | S | Logradouro do local de Coleta do CTE |
| 396 | NUMEROCOLETA | VARCHAR2(60) | S | Número do endereço do local de Coleta do CTE |
| 397 | COMPLEMENTOCOLETA | VARCHAR2(60) | S | Complemento do local de Coleta do CTE |
| 398 | BAIRROCOLETA | VARCHAR2(60) | S | Bairro do local de Coleta do CTE |
| 399 | CODCIDADECOLETA | NUMBER(6) | S | Código da Cidade do local de Coleta do CTE |
| 400 | CNPJFORNECPEDAGIO | VARCHAR2(18) | S | CPNJ da empresa fornecedora do vale pedagio |
| 401 | NUMVALEPEDAGIO | VARCHAR2(20) | S | Número do comprovante de compra do vale pedagio |
| 402 | CPNJRESPONSAVELPEDAGIO | VARCHAR2(18) | S | CNPJ do responsável pelo pagamento do vale pedágio |
| 403 | VALORPEDAGIO | NUMBER(13,2) | S | Valor do Vale Pedágio |
| 404 | NUMTRANSCTEANUL | NUMBER(10) | S | Para que o Ct-e de substituição(Saída) seja vinculado ao Ct-e que foi excluído fiscalmente |
| 405 | NUMTRANSCTESUBS | NUMBER(10) | S | Para que Ct-e de substituição(Saída) seja vinculado ao Ct-e de entrada |
| 406 | CHAVECTEREF | VARCHAR2(45) | S | Chave do cte referenciado |
| 407 | CNPJCPFCTEREF | VARCHAR2(18) | S | CNPJ/CPF do cte referenciado |
| 408 | IECTEREF | VARCHAR2(20) | S | Inscrição estadual do cte referenciado |
| 409 | RAZAOSOCIALCTEREF | VARCHAR2(60) | S | Razão social do cte referenciado |
| 410 | TIPODOCCTEREF | VARCHAR2(2) | S | Tipo do documento do cte referenciado |
| 411 | SERIEDOCCTEREF | VARCHAR2(3) | S | Série do documento do cte referenciado |
| 412 | SUBSERIEDOCCTEREF | VARCHAR2(2) | S | Sub-série do documento do cte referenciado |
| 413 | NUMDOCCTEREF | VARCHAR2(20) | S | Número do documento do cte referenciado |
| 414 | DTEMISSAOCTEREF | DATE | S | Data de emissão do cte referenciado |
| 415 | ORGAOPUBFEDERAL | VARCHAR2(1) | S | Orgão Pub. Federal |
| 416 | ORGAOPUBMUNICIPAL | VARCHAR2(1) | S | Orgão Pub. Municipal |
| 417 | NSUSOCIOTORCEDOR | VARCHAR2(15) | S | Número de Transação Sitef |
| 418 | CNPJCPFEMITENTESUB | VARCHAR2(14) | S | CNPJ ou CPF do emitente |
| 419 | MODELODOCSUB | VARCHAR2(2) | S | Modelo do documento fiscal |
| 420 | SERIEDOCSUB | VARCHAR2(3) | S | Serie do documento fiscal |
| 421 | SUBSERIEDOCSUB | VARCHAR2(3) | S | Subserie do documento fiscal |
| 422 | NUMERODOCSUB | NUMBER(10) | S | Número do documento fiscal |
| 423 | VALORDOCSUB | NUMBER(13,2) | S | Valor do documento fiscal |
| 424 | DTEMISSAODOCSUB | DATE | S | Data da emissão do documento fiscal |
| 425 | CHAVECTETOMADORSUB | VARCHAR2(45) | S | Chave de acesso do CT-e emitido pelo Tomador |
| 426 | CHAVECTEANULACAO | VARCHAR2(45) | S | Chave de acesso do CT-e de Anulação |
| 427 | CHAVECTESUBSTITUICAO | VARCHAR2(45) | S | Chave de acesso do CT-e a ser substituido (original) |
| 428 | CHAVENFETOMADORSUB | VARCHAR2(45) | S | Chave de acesso da NF-e emitida pelo Tomador |
| 429 | TIPODOCUMENTOSUB | VARCHAR2(3) | S | Flag para identificação do tipo do documento (NF/ NFE/ CT/ CTE) |
| 430 | UFCTEREF | VARCHAR2(2) | S | Sigla da UF do Emissor do documento anterior |
| 431 | VLDESCSOCIOTORCEDOR | NUMBER(18,2) | S | Desconto Sócio Torcedor |
| 432 | IDTIPOPRESENCA | VARCHAR2(2) | S | Indica o tipo de presença do comprador no estabelecimento no momento da operação |
| 433 | ROTINACADCTE | VARCHAR2(48) | S | Rotina que efetuou a geração do CTE |
| 434 | NUMREGEXP | NUMBER(12) | S | Número do Registro de Exportação  |
| 435 | NUMCHAVEEXP | VARCHAR2(44) | S | Número da Chave de Acesso da Nfe recebida para exportação |
| 436 | NUMDRAWBACK | NUMBER(11) | S | Número do ato concessório do Drawback |
| 437 | VLSUBTOTAL | NUMBER(18,2) | S | Valor do subtotal dos itens |
| 438 | IDLOCALDESTMERC | VARCHAR2(1) | S | Identificação do local de destino da mercadoria |
| 439 | CODPRODNFAJUSTE | VARCHAR2(60) | S | Código de produto para NF-e de ajuste |
| 440 | DESCRICAOPRODNFAJUSTE | VARCHAR2(120) | S | Descrição do produto paa NF ajuste |
| 441 | NCMNFAJUSTE | VARCHAR2(8) | S | NCM para item de NF de ajuste |
| 442 | UNIDADENFAJUSTE | VARCHAR2(6) | S | Unidade de medida para NF de ajuste |
| 443 | VLTOTBRUTOPRODAJUSTE | NUMBER(18,2) | S | Valor total bruto para NF de ajuste |
| 444 | OBSNFAJUSTE | VARCHAR2(1000) | S | Observação complementar para NF de ajuste |
| 445 | CODFORNECFRETECTEREF | NUMBER(6) | S | Código do fornecedor de frete referenciado (redespacho) informado na 1450 |
| 446 | DTPAGDOTZ | DATE | S | Data do pagamento para Dotz. |
| 447 | VLBONIFICACERTO | NUMBER(14,2) | S | Valor da bonificação acerto |
| 448 | CTECODTRANSPSUBCONTRATANTE | NUMBER(6) | S | Código da transportadora cte subcontratante |
| 449 | CTECODVEICULOSUBCONTRATANTE | NUMBER(6) | S | Código do veiculo cte subcontratante |
| 450 | CTECODTOMADORSERVICO | NUMBER(6) | S | Código do tomador de serviço cte subcontratado |
| 451 | CTECODTRANSPSUBCONTRATADO | NUMBER(6) | S | Código do veiculo cte subcontratado |
| 452 | CTESUBCONTRATADOCHAVECTEREF | VARCHAR2(45) | S | Chave do cte referenciado pelo tomador serviço |
| 453 | CODRECEBFRETECTEREF | NUMBER(6) | S | Código do Recebedor de frete referenciado(redespacho) informado na 1450 |
| 454 | TIPOSERVICOCTE | VARCHAR2(1) | S | Código do tipo de serviço do CT-e |
| 455 | TIPOSUBCONTRATACAOCTE | VARCHAR2(1) | S | Tipo do CT-e |
| 456 | TIPOPRESENCAADQUIRINTE | VARCHAR2(1) | S | Tipo presença do adquirinte |
| 457 | DOCEMISSAO | VARCHAR2(2) | S | Tipo de documento de emissão da venda. |
| 458 | VLGRIS | NUMBER(12,2) | S | Valor do GRIS no conhecimento |
| 459 | VLDESPPEDAGIO | NUMBER(12,2) | S | Valor de Outras despesas de pedágio |
| 460 | VLTAS | NUMBER(12,2) | S | Valor do TAS no conheciemento |
| 461 | NUMFECHAMENTOMOVCX | NUMBER(10) | S | Numero do fechamento da movimentacao |
| 462 | DTMOVIMENTOCX | DATE | S | Data da movimentacao do caixa. |
| 463 | IDENTIFICARCLIENTENFCE | VARCHAR2(1) | S | Indentificar cliente NFC-e |
| 464 | MD5LISTAARQ | VARCHAR2(300) | S | MD5 da lista de arquivos PAF-ECF |
| 465 | VENDANFSERIED | VARCHAR2(1) | S | Venda NF Serie D |
| 466 | SELO | VARCHAR2(20) | S | Selo da sefaz |
| 467 | NUMREMESSA | NUMBER(10) | S | Número remessa garantia |
| 468 | TIPOMOVGARANTIA | NUMBER(3) | S | Define qual tipo de movimentação de garantia a nota representa. |
| 469 | FORMULARIO | NUMBER(10) | S | Formulário da nota fiscal |
| 470 | NUMNOTAMANIF | NUMBER(10) | S | Número da nota de manifesto |
| 471 | ORIGEMFRETE | VARCHAR2(1) | S | Origem do percentual de Frete |
| 472 | NUMCIRURGIA | NUMBER(10) | S | Núm. Cirurgia |
| 473 | UTILIZAEXPDIVERSOS | VARCHAR2(1) | S | Valor do parâmetro na hora que foi gravada a NF |
| 474 | SITUACAONFEEPEC | NUMBER(3) | S | Situação Nfe EPEC |
| 475 | PROTOCOLONFEEPEC | VARCHAR2(20) | S | Protocolo de autorização EPEC |
| 476 | DATAHORAREGISTROEPEC | DATE | S | Data e hora de registro do EPEC |
| 477 | NUMTRANSENTORIGCONSIG | NUMBER(10) | S |  |
| 478 | NUMSERIESAT | VARCHAR2(13) | S | Número de serie do SAT |
| 479 | CHAVESAT | VARCHAR2(400) | S | Chave NFCE gerada pelo SAT |
| 480 | SITUACAOSAT | VARCHAR2(2) | S | Situação da Venda |
| 481 | EMISSNUMAUTOMATICO | VARCHAR2(1) | S | Emissão de nota própia |
| 482 | DATAHORAEMISSAOSAT | DATE | S | Data Hora Aprovação da venda |
| 483 | QRCODESAT | VARCHAR2(4000) | S | Código Consulta venda |
| 484 | NUMSESSAOSAT | NUMBER(10) | S |  |
| 485 | CODSEFAZSAT | NUMBER(10) | S |  |
| 486 | CODSTATUSSAT | NUMBER(10) | S | Status da venda |
| 487 | VLBOLETO | NUMBER(14,2) | S | Valor total de boleto na NF |
| 488 | NUMTRANSENTNFESTORNADA | NUMBER(10) | S | Número da nota de devolução que gerou o estorno |
| 489 | NOTADUPLIQUESVC | VARCHAR2(1) | S | Validar as notas quando elas forem enviadas em ambiente de SVC. Quando nota for duplicada o campo deverá ser preenchido |
| 490 | TIPOINSCPRESTSERV | NUMBER(1) | S | Tipo da Inscrição do Prestador de Serviço |
| 491 | INDICPRESTSERV | NUMBER(1) | S | Indicativo de Prestação de Serviços |
| 492 | NUMCOLETA | VARCHAR2(20) | S |  |
| 493 | NUMCOLETA_AUX | VARCHAR2(20) | S |  |
| 494 | NUMTRANSENTDEPENDENTE | NUMBER(10) | S | Transação de nota de entrada dependente |
| 495 | NUMTRANSVENDADEPENDENTE | NUMBER(10) | S | Transação de nota de saida dependente |
| 496 | ROTINACAD | VARCHAR2(48) | S | Gravar a rotina que inseriu o registro na PCNFSAID |
| 497 | QRCODENFCE | VARCHAR2(500) | S | link para geração do qrcode da nfce |
| 498 | NOMEARQUIVODOTZ | VARCHAR2(100) | S | Campo para armazenar o nome do arquivo dotz no qual foi gerado |
| 499 | NOMEARQUIVODOTZESTORNO | VARCHAR2(100) | S | Campo para armazenar o nome do arquivo dotz no qual foi gerado para estorno |
| 500 | PONTOSDOTZ | NUMBER(6) | S | Pontuação do dotz na venda |
| 501 | VLBASERETENCAO | NUMBER(18,6) | S | Valor de base para rentenção do servico |
| 502 | CODRETENCAOCPRB | NUMBER(1) | S | Código de retenção do CPRB |
| 503 | NUMVIASGNRE | NUMBER(4) | S | Número de impressões da nota na GNRE |
| 504 | QTDEREPROCNFE | NUMBER(6) | S | Quantidade de máxima de tentativas de aprovação da nota. |
| 505 | HORASAIDA | NUMBER(2) | S | Hora de saída |
| 506 | MINUTOSAIDA | NUMBER(2) | S | Minuto de saída |
| 507 | CTEREGIMEESPECIAL | VARCHAR2(1) | S | Define se o CT-e emitido pela rotina 1450 faz parte do regime especial PTA nº 45.000008189-09 |
| 508 | NFDEVENTFUTURA | VARCHAR2(1) | S | Devolução NF entrega Futura |
| 509 | UIDREGISTRO | NUMBER(10) | S | Id único do registro de assinatura |
| 510 | IDPARCEIRO | VARCHAR2(15) | S | Id do parceiro da assinatura do registro |
| 511 | ASSINATURA | VARCHAR2(100) | S | Assinatura do registro |
| 512 | IMPRIMIUDACTEJUNTODANFE | VARCHAR2(1) | S | Imprimiu o DACTe junto ao DANFe vinculado (Se S, significa que o Cte foi impresso junto a Nfe vinculada) |
| 513 | VLDESCFIN | NUMBER(18,6) | S | Valor do desconto financeiro |
| 514 | PERDESCFINTOT | NUMBER(18,6) | S | Percentual total do desconto financeiro na NF |
| 515 | HORAENTREGA | NUMBER(2) | S | Hora de entrega |
| 516 | MINUTOENTREGA | NUMBER(2) | S | Minuto de entrega |
| 517 | DTPAGCOMISSAO | DATE | S | Data de pagamento da comissão |
| 518 | TRANSFDEP | VARCHAR2(1) | S | Define se a NF é uma nota de transferência para o depósito fechado |
| 519 | TIPOTRANSFENTREDEPOSITO | VARCHAR2(1) | S | Tipo de transferência entre depósitos (E = Envio; R = Retorno) |
| 520 | DTPAGCOMISSAOOP | DATE | S | Fechamento de comissão do Operador |
| 521 | NUMNOTAVINCULADA | NUMBER(10) | S | Número de nota vinculada |
| 522 | TIPODOCARRECGNRE | VARCHAR2(1) | S | Tipo de documento - GNRE ou Documento Arrecadação |
| 523 | UFBENEFICIARIA | VARCHAR2(2) | S | Unidade Federação beneficiária |
| 524 | NUMDOCARREC | VARCHAR2(255) | S | Número do documento de arrecadação |
| 525 | CODAUTENTICACAO | VARCHAR2(255) | S | Código de autenticação |
| 526 | VLDOCARRECADACAO | NUMBER(20,6) | S | Valor do documento de arrecadação |
| 527 | DTVENCARREC | DATE | S | Data do Vencimento |
| 528 | DTPAGARREC | DATE | S | Data do Pagamento |
| 529 | VLOPERACAOFRETE | NUMBER(14,2) | S | Valor da Operação |
| 530 | VLBASESTFRETE | NUMBER(14,2) | S | Base de cálculo ST frete |
| 531 | VLICMSSTFRETE | NUMBER(14,2) | S | Valor do ICMS/ST Frete |
| 532 | CTEGLOBALIZADO | VARCHAR2(1) | S | CT-e Globalizado (Sim ou Não) |
| 533 | TOMADORCTE | NUMBER(1) | S | Define o tomador do Cte (0-Remetente/1-Expedidor/2-Recebedor/3-Destinatário) |
| 534 | VLCSLL | NUMBER(18,6) | S | Valor do CSLL RETIDO NA NF |
| 535 | CNPJSEGURADORA | VARCHAR2(14) | S | CNPJ da seguradora |
| 536 | INFGLOBALIZADO | VARCHAR2(256) | S | Informações adicionais do Cte Globalizado |
| 537 | QTDEREPROCCTE | NUMBER(6) | S | Quantidade de reprocessamento CT-e |
| 538 | PERCENTUALSTFRETE | NUMBER(14,2) | S |  |
| 539 | CODEDITAL | NUMBER(9) | S | Código Edital |
| 540 | FIL_MONTARKITAUTOMATICAMENTE | VARCHAR2(1) | S | grava o valor do parâmetro no momento da geração da nota fiscal |
| 541 | NUMNOTAORIGEM | NUMBER(10) | S | Numero de nota original |
| 542 | PAGCHEQUEMORADIA | VARCHAR2(1) | S | Define se o cliente faz pagamento usando cheque moradia |
| 543 | CODCOBSEFAZ | VARCHAR2(4) | S | Código de cobrança da SEFAZ |
| 544 | DIGITOVERIFICADORNFSE | VARCHAR2(100) | S | Digito Verificardor NFSE |
| 545 | LINKNFSE | VARCHAR2(100) | S | Link de consulta NFSE |
| 546 | PROTOCOLONFSE | VARCHAR2(16) | S | Protocolo NFSE |
| 547 | SERIERPS | VARCHAR2(3) | S | Serie RPS |
| 548 | NUMERORPS | NUMBER(10) | S | Numero RPS |
| 549 | SITUACAONFSE | NUMBER(10) | S | Situacao NFSE |
| 550 | NUMNOTANFSE | NUMBER(18) | S | Numero nota NFSE |
| 551 | MD5PAF | VARCHAR2(200) | S | Assinatura MD5 do registro do PAF |
| 552 | USAMULTUNIDITEM | VARCHAR2(1) | S | Usa multimplas unidades (Master ou Venda) por item na entrada mercadoria |
| 553 | CONDFINANC | VARCHAR2(9) | S | Condição de financiamento suppliercard. |
| 554 | PLANOSUPPLI | NUMBER(2) | S | Plano de quantidade de parcelas suppliercard. |
| 555 | TIPOREMESSA | VARCHAR2(1) | S | Indica o tipo de simples remessa M-Manutenção B-Baixa de Item |
| 556 | CODTRANSPAUTONOMO | NUMBER(10) | S | Indica o código do transportador autônomo |
| 557 | VLTRANSPAUTONOMO | NUMBER(18,6) | S | indica o valor à pagar ao transportador autônomo |
| 558 | TRANSPAUTONOMO | VARCHAR2(1) | S | Indica se o transporte foi realizar por um transportador autônomo |
| 559 | GERARCPTRANSPAUNOTOMO | VARCHAR2(1) | S | Indica se irá gerar contas à pagar para o transportador autônomo |
| 560 | CODFORNECTRANSPAUTONOMO | NUMBER(10) | S | Indica o código do transportador autônomo do frete |
| 561 | VLTOTALFCPST | NUMBER(18,6) | S | Valor Total do Fundo de Combate a Pobreza ST |
| 562 | VLTOTALFCPICMS | NUMBER(18,6) | S | Valor Total do Fundo de Combate a Pobreza ICMS |
| 563 | VLFRETEOUTRASDESP | NUMBER(14,2) | S | Valor fixo da Taxa de Frete que será impresso na NF. |
| 564 | VLTOTALNFCOMPOEMCTE | NUMBER(22,6) | S | Valor total das notas fiscais que compõem o conhecimento de frete |
| 565 | DEDUZIRDESONERORGAOPUB | VARCHAR2(1) | S | Deduzir ICMS desonerado NF Licitação Orgãos Públicos |
| 566 | TIPOCONTACORRENTE | VARCHAR2(1) | S | Tipo de conta corrente movimentada: RCA, Supervisor ou Gerente |
| 567 | NUMEROPROCESSO | VARCHAR2(14) | S | Número do processo |
| 568 | TIPOPROCESSO | VARCHAR2(1) | S | Tipo do processo |
| 569 | CODIGOINDICATIVOSUSPENSAO | VARCHAR2(14) | S | Código Indicativo de Suspensão |
| 570 | CODPROMOCAOMED | NUMBER(9) | S | Código da Promoção Medicamento |
| 571 | CODPLPAGETICO | NUMBER(4) | S | Código Plano Pagamento Etico |
| 572 | CODPLPAGGENERICO | NUMBER(4) | S | Código Plano Pagamento Genérico |
| 573 | VLPUNITMED | NUMBER(12,2) | S | Valor Total a Preço Unitário Medicamento |
| 574 | MATERIALDECONSTRUCAO | VARCHAR2(1) | S | Identifica se o pedido foi feito pelo modulo 041 |
| 575 | DEDUZIRICMSBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do ICMS da base de cálculo do PIS/COFINS |
| 576 | CODMUNINICTE | NUMBER(7) | S | Código do município de origem da prestação do serviço(Cte) |
| 577 | CODMUNFIMCTE | NUMBER(7) | S | Código do município de destino da prestação do serviço(Cte) |
| 578 | VLCREDICMSPRESUMIDO | NUMBER(12,2) | S | Valor de credito icms retido para ct-e |
| 579 | DATACONSOLIDACAOPREFAT | DATE | S | Data consolidação Pré-Faturamento |
| 580 | VERSAOLAYOUTNFE | VARCHAR2(5) | S | Versão do layout do arquivo na Sefaz |
| 581 | OBSNFEFISCO | CLOB | S | Informações Adicionais de Interesse do Fisco |
| 582 | PCT_ESTOQUE | VARCHAR2(1) | S | Gravar a efetivação da package de estoque |
| 583 | PCT_CUSTO | VARCHAR2(1) | S | Gravar a efetivação do serviço de custo |
| 584 | CODDEVEDORCTE | NUMBER(6) | S | Código do Tomador/Devedor de serviço informado na geração de CTE |
| 585 | VLSERVICOFRETEAUTON | NUMBER(10,2) | S | Valor do serviço de frete autônomo |
| 586 | VLBASEFRETEAUTON | NUMBER(10,2) | S | Valor da base de cálculo do frete autônomo |
| 587 | ALIQICMFRETEAUTON | NUMBER(10,2) | S | Aliquota de ICMS do frete autônomo |
| 588 | VLICMFRETEAUTON | NUMBER(10,2) | S | Valor de ICMS do frete autônomo |
| 589 | CODFISCALFRETEAUTON | NUMBER(8) | S | Código fiscal utilizado no processo de frete autônomo |
| 590 | CODMUNICIPFATGERADOR | NUMBER(10) | S | Código IBGE municipio fato gerador no processo de frete autônomo |
| 591 | NUMNOTAPREF | NUMBER(18) | S | Número da NF da Prefeitura |
| 592 | STBCR | NUMBER(18,6) | S | Total do ST BCR na nota fiscal |
| 593 | SOMAREPASSEOUTRASDESPNF | VARCHAR2(1) | S | Se somou vlrepasse nas outras despesas |
| 594 | NFIPIEMITIDA | VARCHAR2(1) | S | Define se NF-e de IPI já foi emitida |
| 595 | APROVEITASTDEVFORNEC | VARCHAR2(1) | S | Aproveita ST NF na devolução a fornecedor |
| 596 | INDICADORDESCONTO | VARCHAR2(1) | S | Indicador de desconto do PAFECF |
| 597 | INDICADORACRESCIMO | VARCHAR2(1) | S | Indicador de acréscimo do PAFECF |
| 598 | CODCOORDENADOR | NUMBER(10) | S | Código do Coordenador de Venda |
| 599 | USASRVTERCEIROS | VARCHAR2(1) | S | Identifica se durante o faturamento o valor do parametro 4182 |
| 600 | ENVIADASEMRESPOSTA | VARCHAR2(1) | S | Define se a nota foi enviada a Sefaz mas não houve resposta, utilizado para reprocessamento. |
| 601 | IDSOFITVIEW | VARCHAR2(10) | S | Indica o código da receita na SofitView |
| 602 | DTEXCLUSAOSOFITVIEW | TIMESTAMP(6) | S | Indica a data que a receita foi excluída na SofitView |
| 603 | QRCODE | VARCHAR2(1000) | S | QRCODE PARA NF-e/CT-e |
| 604 | CODGERENTENACIONAL | NUMBER(4) | S | Gerente Nacional |
| 605 | CODDIRETOR | NUMBER(4) | S | Diretor de Venda |
| 606 | CODGERENTEREGIONAL | NUMBER(4) | S | Gerente Regional |
| 607 | PEDIDOAVARIA | VARCHAR2(1) | S | Identificador de Pedido de Avaria |
| 608 | NUMTRANSENTCROSSDOCK | NUMBER(10) | S | Número Trans. Entrada Pedido Crossdocking |
| 609 | CODPROFISSIONAL | NUMBER(6) | S | Código do profissional utilizado no pedido. |
| 610 | NUMTRANSENTORIGTV10 | NUMBER(10) | S | Número da Transação de Entrada de Origem da Transferencia |
| 611 | REDUZICMSDOCTE | VARCHAR2(1) | S | Campo indica se o valor de ICMS vou reduzido do valor total do CT-e |
| 612 | VLJUROSPARCELAMENTO | NUMBER(18,2) | S | Valor de juros referente ao parcelamento do pagamento |
| 613 | PCT_ESTOQUE_CANCELAMENTO | VARCHAR2(1) | S | Utilizou PKG Estoque para Cancelar |
| 614 | PCT_CUSTO_CANCELAMENTO | VARCHAR2(1) | S | Utilizou PKG Custo para Cancelar |
| 615 | NUMCIOT | VARCHAR2(20) | S | Número do controto do CIOT |
| 616 | NFBRINDEZERADA | VARCHAR2(1) | S | NF Brinde com preço e impostos zerados |
| 617 | NFBRINDE | VARCHAR2(1) | S | Nota de Brinde |
| 618 | QTPALETE | NUMBER(10) | S | Campo utilizado para gravar quantidade de Palete do pedido |
| 619 | TOTPESOPALETE | NUMBER(12,6) | S | Campo utilizado para gravar o peso do Palete do pedido |
| 620 | CODIGONUMERICOCHAVE | VARCHAR2(8) | S | Código númerico que compoem a chave |
| 621 | TIPOIMPRESSAO | VARCHAR2(1) | S | Tipo de impressão (retrato ou paisagem) |
| 622 | DEDUZFUNCEPBASEPISCOFINS | VARCHAR2(1) | S | Deduzir valor do FUNCEP da Base de Cálculo do Pis/Cofins |
| 623 | CNPJINTERMEDIADOR | VARCHAR2(20) | S | CNPJ do Intermediador da Transação |
| 624 | DESCINTERMEDIADOR | VARCHAR2(60) | S | Nome do usuário ou identificação do perfil do vendedor no site do intermediador |
| 625 | DESFEZKITCANC | VARCHAR2(1) | S | Flag que indica se foi desfeito o kit/cesta/tinta. |
| 626 | DTEXPWMSSAAS | DATE | S | Data de integração da NF no WMS SaaS |
| 627 | NUMTRANSPAGDIGITAL | VARCHAR2(50) | S | Número identificador da transação |
| 628 | CALCDESONERACAOFATMEDICAM | VARCHAR2(1) | S | Calcular a Desoneração no Faturamento Farma/Hospitalar |
| 629 | GERARCPCT | VARCHAR2(1) | S | Gerar Contas a Pagar Conhecimento de Transporte |
| 630 | GERARCRCT | VARCHAR2(1) | S | Gerar Contas a Receber Conhecimento de Transporte |
| 631 | GEROUGRNE | VARCHAR2(50) | S | Identificador para saber se gerou GRNE por códigos |
| 632 | VOLTARESTOQUEFILIALVIRTUAL | VARCHAR2(1) | S | Indica se o estoque irá retornar para filial de venda |
| 633 | VOLTARESTOQUEFILIALRETIRA | VARCHAR2(1) | S | Indica se o estoque irá retornar para filial retira |
| 634 | TIPOPRESTACAO | VARCHAR2(1) | S | Tipo de prestação de serviço para o CTe Simplificado |
| 635 | CTESIMPLIFICADO | VARCHAR2(1) | S | Define se o CTe é Simplificado |
| 636 | VLFRETETRANSPORTARNF | NUMBER(14,2) | S | Valor do frete para transportar as mercadoria da Nota Fiscal |
| 637 | IDEXTERNOCOLETAENTREGA | VARCHAR2(40) | S | Index de integração |
| 638 | SITUACAOCOLETAENTREGA | VARCHAR2(40) | S | SITUACAO COLETA ENTREGA |
| 639 | DTENTGACOLETAENTREGA | DATE | S | Data Entrada GA Coleta Entrega  |
| 640 | DTULTALTOCOLETAENTREGA | DATE | S | Data Ultima Alto Coleta Eentrega |
| 641 | REJEITADOSEFAZ | VARCHAR2(1) | S | Indica se o rejeição da NFe é Local ou da SEFAZ |
| 642 | INDVAREJO | NUMBER(1) | S | Indicação de nota do varejo |
| 643 | TIPOORGAOPUB | VARCHAR2(2) | S | Tipo do Órgão Público. Preencher com: 0=Nenhum; 1=União; 2=Estado; 3=Distrito Federal ou 4=Município (Rotinas 202 e 302) |
| 644 | PERCREDORGAOPUB | NUMBER(12,4) | S | Percentual de redução da alíquota IBS, CBS e IS para Órgão Público (Rotinas 202 e 302) |
| 645 | VLTOTALIS | NUMBER(18,6) | S | Total do imposto seletivo |
| 646 | TPENTEGOV | VARCHAR2(1) | S | Tipo de ente governamental |
| 647 | PREDUTOR | NUMBER(7,4) | S | Percentual de redução de alíquota em compra governamental |
| 648 | VLIBSMUN | NUMBER(15,2) | S | Valor total do IBS do Município |
| 649 | VLIBSUF | NUMBER(15,2) | S | Valor total do IBS da UF |
| 650 | VLCBS | NUMBER(15,2) | S | Valor total da CBS |
| 651 | TIPONOTACREDDEB | VARCHAR2(2) | S | Destaque para tipo de Nota Débito e Crédito |
| 652 | VLTOTALIBS | NUMBER(15,2) | S | Valor total de IBS |
| 653 | VLTOTALCBS | NUMBER(15,2) | S | Valor total de CBS |
| 654 | VBCICBSCBS | NUMBER(15,2) | S | Valor total da BC do IBS e da CBS |
| 655 | IMPORTADOXML | VARCHAR2(1) | S | Identificador se o lançamento foi inserido por importação de xml |
| 656 | VBCIBSCBS | NUMBER(15,2) | S | Valor total da BC do IBS e da CBS |
| 657 | TPOPERGOV | NUMBER(1) | S | Tipo de operação com o ente governamental |
| 658 | NFCREDITO | VARCHAR2(1) | S | Nota de Crédito |
| 659 | TIPONFCREDITO | VARCHAR2(2) | S | Tipo NF Crédito |
| 660 | SAIDAREMESSAFULFILLMENT | VARCHAR2(1) | S | Identifica se a nota pertence a uma remessa fulfillment |
| 661 | CODOBSVENDA | NUMBER(6) | S | Codigo da Observacao da venda - PCOBSVENDA |
| 662 | INFADFISCO | VARCHAR2(2000) | S | Informação do XML conteúdo da tag InfAdFisco |
| 663 | ESTORNODECREDITO | VARCHAR2(1) | S | Estorna Créditos IBS/CBS |
| 664 | NUMNOTASERVICO | VARCHAR2(60) | S | Nº Nota Fiscal de Serviço |
| 665 | CHAVENFSE | VARCHAR2(60) | S | Chave Nota Fiscal de Serviço |
| 666 | ISUFEMIT | VARCHAR2(15) | S | Inscrição SUFRAMA do Emitente |
| 667 | IMPXML | VARCHAR2(1) | S | Indica se NFe foi gerada a partir de importação direta de um arquivo XML. |
| 668 | NUMTRANSVENDANFDEBITO | NUMBER(10) | S | Transação de venda da nota de débito |
| 669 | OPERCOMPGOV | VARCHAR2(1) | S | Operação compra governamental |
| 670 | VERSAODOCFISCAL | VARCHAR2(30) | S | Versão do DocFiscal |

## Chaves e índices

- Constraint PCNFSAID_PK (P): NUMTRANSVENDA
- Constraint PCNFSAID_UK (U): UIDREGISTRO, IDPARCEIRO
- Índice IDX$$_152F0001: NUMCAR, CODVEICULO
- Índice IDX$$_3ABD0005: ORGAOPUB, ESPECIE, DTSAIDA
- Índice IDX$$_D36F0001: CODFISCAL, DTCANCEL
- Índice IDX_003300: NUMSERIESAT, SERIE, DTSAIDA
- Índice PCNFSAID_IDX1: NUMNOTA
- Índice PCNFSAID_IDX10: CODSUPERVISOR, CODUSUR, DTSAIDA
- Índice PCNFSAID_IDX11: NUMCAR, NUMTRANSVENDA
- Índice PCNFSAID_IDX12: NUMCAR, NUMNOTA
- Índice PCNFSAID_IDX13: NUMCAR, NUMPED
- Índice PCNFSAID_IDX14: NUMPED
- Índice PCNFSAID_IDX15: NUMCAR, CODCLI, NUMSEQMONTAGEM
- Índice PCNFSAID_IDX16: SITUACAONFE, ESPECIE, CONDVENDA, SYS_NC00503$, DTCANCEL
- Índice PCNFSAID_IDX2: DTSAIDA
- Índice PCNFSAID_IDX22: NUMTRANSVENDA, NUMNOTA
- Índice PCNFSAID_IDX23: NUMTRANSVENDAORIGEM, DTCANCEL
- Índice PCNFSAID_IDX24: SITUACAONFE, ESPECIE
- Índice PCNFSAID_IDX25: NUMTRANSVENDA, CODFISCAL, CONDVENDA, DTCANCEL
- Índice PCNFSAID_IDX29: NUMPED, NUMTRANSVENDA
- Índice PCNFSAID_IDX31: CODFILIAL, NUMTRANSVENDA
- Índice PCNFSAID_IDX32: CODFILIAL, DTSAIDA, CODFISCAL, DTCANCEL
- Índice PCNFSAID_IDX33: CHAVENFE
- Índice PCNFSAID_IDX34: TIPOEMISSAO, ESPECIE, SITUACAONFE
- Índice PCNFSAID_IDX35: CODFORNEC, TIPODEVOL
- Índice PCNFSAID_IDX36: NUMORCA, DTCANCEL
- Índice PCNFSAID_IDX37: SYS_NC00504$
- Índice PCNFSAID_IDX38: CODEMITENTE, CAIXA, NUMTRANSVENDA
- Índice PCNFSAID_IDX39: NUMTRANSVENDACONHEC
- Índice PCNFSAID_IDX4: CODCLI
- Índice PCNFSAID_IDX40: NUMLANCCP
- Índice PCNFSAID_IDX41: SITUACAOCTE, SYS_NC00503$
- Índice PCNFSAID_IDX42: TIPOVENDA, DTSAIDA
- Índice PCNFSAID_IDX43: NUMTRANSVENDAORIGEM
- Índice PCNFSAID_IDX44: CHAVECTE
- Índice PCNFSAID_IDX45: NUMTRANSENTORIGEM
- Índice PCNFSAID_IDX46: NUMTRANSVENDATV13
- Índice PCNFSAID_IDX47: NUMTRANSCTEANUL
- Índice PCNFSAID_IDX48: NUMTRANSCTESUBS
- Índice PCNFSAID_IDX49: DTSAIDA, NUMNOTA, SYS_NC00503$
- Índice PCNFSAID_IDX50: DTCANCEL, SYS_NC00503$
- Índice PCNFSAID_IDX51: NUMTRANSVENDADESTINO
- Índice PCNFSAID_IDX52: RECIBONFE
- Índice PCNFSAID_IDX53: RECIBOCTE
- Índice PCNFSAID_IDX54: RECIBODPEC
- Índice PCNFSAID_IDX55: DTSAIDA, CONDVENDA, CODFISCAL
- Índice PCNFSAID_IDX56: NUMVIAS, SITUACAONFE, ESPECIE, DTCANCEL, DTSAIDA
- Índice PCNFSAID_IDX57: DTCANCEL, CONDVENDA
- Índice PCNFSAID_IDX58: NUMLOTENFE
- Índice PCNFSAID_IDX59: ESPECIE, SYS_NC00587$, SYS_NC00588$
- Índice PCNFSAID_IDX60: CONDVENDA, CODCOB
- Índice PCNFSAID_IDX61: CODFILIAL, DTSAIDA, CODFISCAL, CONDVENDA, DTCANCEL
- Índice PCNFSAID_IDX62: DTSAIDA, SYS_NC00515$, CODFILIAL
- Índice PCNFSAID_IDX63: DTSAIDA, CODFILIAL, DTCANCEL
- Índice PCNFSAID_IDX64: DTFECHA
- Índice PCNFSAID_IDX65: CODCLI, DTSAIDA
- Índice PCNFSAID_IDX66: NUMTRANSENTNFESTORNADA
- Índice PCNFSAID_IDX67: NUMTRANSVENDADEPENDENTE
- Índice PCNFSAID_IDX68: DTSAIDA, NUMPED, NUMTRANSVENDA
- Índice PCNFSAID_PK UNIQUE: NUMTRANSVENDA
- Índice PCNFSAID_UK UNIQUE: UIDREGISTRO, IDPARCEIRO
