# RIBERDOCES.PCEST — dicionário

Gerado por `gateway inventory` em 2026-10-25. 193 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODFILIAL | VARCHAR2(2) | N |  |
| 2 | CODPROD | NUMBER(6) | N |  |
| 3 | QTEST | NUMBER(22,8) | S |  |
| 4 | ESTMIN | NUMBER(16,3) | S |  |
| 5 | DTULTENT | DATE | S |  |
| 6 | DTULTSAIDA | DATE | S |  |
| 7 | QTVENDMES | NUMBER(16,3) | S |  |
| 8 | QTVENDMES1 | NUMBER(16,3) | S |  |
| 9 | QTVENDMES2 | NUMBER(16,3) | S |  |
| 10 | QTVENDMES3 | NUMBER(16,3) | S |  |
| 11 | QTPERDADIA | NUMBER(16,3) | S |  |
| 12 | QTPERDAMES | NUMBER(16,3) | S |  |
| 13 | VLVENDMES | NUMBER(12,2) | S |  |
| 14 | QTVENDDIA | NUMBER(16,3) | S |  |
| 15 | VLVENDDIA | NUMBER(12,2) | S |  |
| 16 | QTPEDIDA | NUMBER(16,3) | S |  |
| 17 | VLCUSTODIAREAL | NUMBER(12,2) | S |  |
| 18 | VLCUSTOMESREAL | NUMBER(12,2) | S |  |
| 19 | VLCUSTODIAFIN | NUMBER(12,2) | S |  |
| 20 | VLCUSTOMESFIN | NUMBER(12,2) | S |  |
| 21 | QTESTANT | NUMBER(16,3) | S |  |
| 22 | PCUSTOANT | NUMBER(18,6) | S |  |
| 23 | QTRESERV | NUMBER(22,8) | S |  |
| 24 | QTULTENT | NUMBER(16,3) | S |  |
| 25 | CUSTOULTENT | NUMBER(18,6) | S |  |
| 26 | QTESTGER | NUMBER(22,8) | S |  |
| 27 | QTENTDIA | NUMBER(16,3) | S |  |
| 28 | QTENTMES | NUMBER(16,3) | S |  |
| 29 | VLVENDDIAREAL | NUMBER(12,2) | S |  |
| 30 | VLVENDMESREAL | NUMBER(12,2) | S |  |
| 31 | CUSTOCONT | NUMBER(18,6) | S |  |
| 32 | QTBLOQUEADA | NUMBER(20,6) | S |  |
| 33 | FIMESTOQUE | VARCHAR2(1) | S |  |
| 34 | QTINDENIZ | NUMBER(20,6) | S |  |
| 35 | QTULTINVENT | NUMBER(16,3) | S |  |
| 36 | QTVENDAPERDIDA | NUMBER(16,3) | S |  |
| 37 | CUSTOREAL | NUMBER(18,6) | S |  |
| 38 | CUSTOFIN | NUMBER(18,6) | S |  |
| 39 | CUSTOREP | NUMBER(18,6) | S |  |
| 40 | VALORULTENT | NUMBER(18,6) | S |  |
| 41 | QTBLOQTRANSF | NUMBER(16,3) | S |  |
| 42 | QTPENDENTE | NUMBER(16,3) | S |  |
| 43 | QTVENDSEMANA | NUMBER(16,3) | S |  |
| 44 | QTVENDSEMANA1 | NUMBER(16,3) | S |  |
| 45 | QTVENDSEMANA2 | NUMBER(16,3) | S |  |
| 46 | QTVENDSEMANA3 | NUMBER(16,3) | S |  |
| 47 | QTGIRODIA | NUMBER(16,3) | S |  |
| 48 | PERCEVOLUCAOVENDA | NUMBER(12,4) | S |  |
| 49 | VLVENDMESANT | NUMBER(14,2) | S |  |
| 50 | VLCUSTOMESFINANT | NUMBER(14,2) | S |  |
| 51 | VLCUSTOMESREALANT | NUMBER(14,2) | S |  |
| 52 | NUMNOTAULTENT | NUMBER(10) | S |  |
| 53 | PCOMPRAULTENT | NUMBER(18,6) | S |  |
| 54 | VLSTULTENT | NUMBER(18,6) | S |  |
| 55 | MODULO | NUMBER(2) | S |  |
| 56 | RUA | NUMBER(4) | S |  |
| 57 | APTO | NUMBER(4) | S |  |
| 58 | DTULTINVENT | DATE | S |  |
| 59 | CUSTODOLAR | NUMBER(18,6) | S |  |
| 60 | CUSTOULTENTMED | NUMBER(18,6) | S |  |
| 61 | VALORULTENTMED | NUMBER(18,6) | S |  |
| 62 | QTRESERVMED | NUMBER(16,3) | S |  |
| 63 | QTREQUISICAO | NUMBER(16,3) | S |  |
| 64 | CUSTOREALSEMST | NUMBER(18,6) | S |  |
| 65 | DTULTPEDCOMPRA | DATE | S |  |
| 66 | CUSTOULTPEDCOMPRA | NUMBER(18,6) | S |  |
| 67 | CUSTOULTENTFIN | NUMBER(18,6) | S |  |
| 68 | DTULTDESDOBRA | DATE | S |  |
| 69 | CORREDOR | NUMBER(4) | S |  |
| 70 | QTLOJA | NUMBER(20,8) | S |  |
| 71 | BASEBCR | NUMBER(18,6) | S |  |
| 72 | STBCR | NUMBER(18,6) | S |  |
| 73 | CUSTOPROXIMACOMPRA | NUMBER(18,6) | S |  |
| 74 | NUMTRANSENTULTENT | NUMBER(10) | S |  |
| 75 | CUSTOULTENTANT | NUMBER(18,6) | S |  |
| 76 | MOTIVOBLOQESTOQUE | VARCHAR2(80) | S |  |
| 77 | CODDEVOL | NUMBER(4) | S | Código do motivo de devolução que é cadastrada na rotina 518 onde o tipo é BO. |
| 78 | QTDEVOLMES | NUMBER(16,3) | S |  |
| 79 | QTDEVOLMES1 | NUMBER(16,3) | S |  |
| 80 | QTDEVOLMES2 | NUMBER(16,3) | S |  |
| 81 | QTDEVOLMES3 | NUMBER(16,3) | S |  |
| 82 | ESTMAX | NUMBER(16,3) | S |  |
| 83 | VLULTENTCONTSEMST | NUMBER(18,6) | S | Valor última entrada contábil.  |
| 84 | CUSTOULTENTLIQ | NUMBER(18,6) | S | Custo da última entrada sempre deduzindo o crédito de icms e pis/cofins, independente dos parâmetros da filial.  |
| 85 | CUSTOREALLIQ | NUMBER(18,6) | S | Custo real sempre deduzindo o crédito de icms e pis/cofins, independente dos parâmetros da filial.  |
| 86 | MODULOCX | NUMBER(2) | S | Indica o modulo do endereço do produto. |
| 87 | RUACX | NUMBER(4) | S | Indica rua do endereço do produto. |
| 88 | APTOCX | NUMBER(4) | S | Indica o apto do endereço do produto. |
| 89 | DTULTALTERSRVPRC | DATE | S | Indica a data da ultima alteração nesta tabela. |
| 90 | DTPRIMCOMPRA | DATE | S | Indica a data da primeira compra. |
| 91 | ESTIDEAL | NUMBER(16,3) | S | Indica o estoque ideal por filial. |
| 92 | DTULTFAT | DATE | S | Indica a data do último faturamento. |
| 93 | VLULTPCOMPRA | NUMBER(18,6) | S | Indica o valor do último preço de compra. |
| 94 | CUSTONFSEMST | NUMBER(18,6) | S | Custo da NF sem ST |
| 95 | BASEICMSULTENT | NUMBER(18,6) | S | Base de ICMS para última entrada |
| 96 | CUSTOULTENTCONT | NUMBER(18,6) | S | Custo da última entrada utilizado para cálculo do custo contábil |
| 97 | QTTRANSITO | NUMBER(22,8) | S | Quantidade em Trânsito |
| 98 | VLSTGUIAULTENT | NUMBER(18,6) | S | Valor do ST GUIA da última entrada. |
| 99 | CUSTONFSEMSTGUIAULTENT | NUMBER(18,6) | S | Base de cálculo ST GUIA. |
| 100 | IVAULTENT | NUMBER(8,4) | S | IVA. |
| 101 | ALIQICMS1ULTENT | NUMBER(12,4) | S | Alíquota ICMS interna. |
| 102 | ALIQICMS2ULTENT | NUMBER(12,4) | S | Alíquota ICMS externa. |
| 103 | REDBASEIVAULTENT | NUMBER(18,6) | S | Redução base de cálculo do ST. |
| 104 | PERCICMSFRETEFOBSTULTENT | NUMBER(12,4) | S | % ICMS frete FOB ST. |
| 105 | VLFRETECONHECULTENT | NUMBER(18,6) | S | Valor frete FOB. |
| 106 | PERCALIQEXTGUIAULTENT | NUMBER(12,4) | S | % aliq. externa guia. |
| 107 | BASEICMSULTENTTAB | NUMBER(18,6) | S | Base do ICMS Futuro. |
| 108 | VLSTULTENTTAB | NUMBER(18,6) | S | Valor do ST futuro. |
| 109 | VLSTGUIAULTENTTAB | NUMBER(18,6) | S | Valor do ST GUIA futuro. |
| 110 | CUSTONFSEMSTTAB | NUMBER(18,6) | S | Custo da NF sem ST futuro. |
| 111 | CUSTONFSEMSTGUIAULTENTTAB | NUMBER(18,6) | S | Base de cálculo ST GUIA futuro. |
| 112 | IVAULTENTTAB | NUMBER(8,4) | S | Valor do IVA futuro. |
| 113 | ALIQICMS1ULTENTTAB | NUMBER(12,4) | S | Alíquota ICMS interna futuro. |
| 114 | ALIQICMS2ULTENTTAB | NUMBER(12,4) | S | Alíquota ICMS externa futuro. |
| 115 | REDBASEIVAULTENTTAB | NUMBER(18,6) | S | Redução da base cálculo ST futuro. |
| 116 | PERCICMSFRETEFOBSTULTENTTAB | NUMBER(12,4) | S | Percentual ICMS frete FOB ST futuro. |
| 117 | VLFRETECONHECULTENTTAB | NUMBER(18,6) | S | Valor frete FOB futuro. |
| 118 | PERCALIQEXTGUIAULTENTTAB | NUMBER(12,4) | S | Percentual alíquota externa GUIA futuro. |
| 119 | PERCMVAORIGULTENT | NUMBER(12,4) | S | Percentual de MVA Original para cálculo do MVA ajustado - Última entrada |
| 120 | PERCMVAORIGTAB | NUMBER(12,4) | S | Percentual de MVA Original para cálculo do MVA ajustado - Futuro |
| 121 | ASSINATURA | VARCHAR2(255) | S | Código MD-5. |
| 122 | QTESTORNOPRODUZDIA | NUMBER(22,8) | S | QUANTIDADE ESTORNADA NO DIA. |
| 123 | QTESTORNOPRODUZMES | NUMBER(22,8) | S | QUANTIDADE ESTORNADA NO MES. |
| 124 | QTPRODUZIDADIA | NUMBER(22,8) | S | QUANTIDADE PRODUZIDA NO DIA. |
| 125 | QTPRODUZIDAMES | NUMBER(22,8) | S | QUANTIDADE PRODUZIDA NO MES. |
| 126 | ROTINALANC | VARCHAR2(48) | S | ROTINA QUE GRAVOU A INFORMACAO |
| 127 | DTEMISSAOULTENT | DATE | S | Data da última NF de Entrada |
| 128 | DTALTERACAO1107 | DATE | S | Dt. Alteração pela rotina 1107 |
| 129 | CUSTOFORNEC | NUMBER(12,6) | S | Custo fornecedor |
| 130 | QTGIROCURVATURA | NUMBER(16,3) | S | Qtde de giro curvatura |
| 131 | QTGIROCURVATURAANT | NUMBER(16,3) | S | Qtde de giro curvatura anterior  |
| 132 | QTFRENTELOJA | NUMBER(22,6) | S | Quantidade no estoque da loja |
| 133 | NUMFCI | VARCHAR2(36) | S | Ficha de Conteúdo de Importação. |
| 134 | VLPARCELAIMPFCI | NUMBER(18,6) | S | Valor da Parcela Importada. |
| 135 | PERCCONTEUDOIMPFCI | NUMBER(5,2) | S | Conteúdo de Importação. |
| 136 | VLIMPORTACAOFCI | NUMBER(18,6) | S | Valor de Importação. |
| 137 | CUSTOFORNECSEMST | NUMBER(18,6) | S | Valor do Custo Fornecedor sem ST. |
| 138 | CUSTOPROXIMACOMPRASEMST | NUMBER(18,6) | S | Valor do Custo da próxima compra sem ST. |
| 139 | CUSTOFINSEMST | NUMBER(18,6) | S | Valor do Custo médio Financeiro sem ST NF e GUIA. |
| 140 | CUSTOULTENTFINSEMST | NUMBER(18,6) | S | Custo Ult.Entrada Financeira sem ST. |
| 141 | CUSTOULTENTSEMST | NUMBER(18,6) | S | Custo Ult.Entrada sem ST. |
| 142 | CUSTOULTPEDCOMPRASEMST | NUMBER(18,6) | S | Custo Ult.Pedido de Compra sem ST. |
| 143 | BASEIPIULTENT | NUMBER(18,6) | S | Base de IPI da ultima entrada |
| 144 | VLIPIULTENT | NUMBER(18,6) | S | Valor de IPI da ultima entrada |
| 145 | PERCIPIULTENT | NUMBER(18,6) | S | Percentual de IPI da ultima entrada |
| 146 | QTTRANSITOTV13 | NUMBER(22,8) | S | Quantidade de estoque em trânsito |
| 147 | CODCEST | VARCHAR2(7) | S |  |
| 148 | IDENTIFICADOR | NUMBER(16) | S | Identificar o número da transação que realizou a ultima alteração |
| 149 | TEMESTOQUEECOMMERCE | VARCHAR2(1) | S | Identifica se tem ou não estoque no e-commerce (InfraCommerce/Unilever) |
| 150 | QTINDUSTRIA | NUMBER(20,6) | S | Estoque físico da industria no estoque geral |
| 151 | VLFCPICMSULTENTTAB | NUMBER(18,6) | S | Valor do Fundo de Combate a Pobreza ICMS |
| 152 | VLFCPSTULTENTTAB | NUMBER(18,6) | S | Valor do Fundo de Combate a Pobreza ST |
| 153 | VLBASEFCPSTULTENTTAB | NUMBER(18,6) | S | Valor de base de calculo do Fundo de Combate a Pobreza ST |
| 154 | VLBCFCPSTRET | NUMBER(18,6) | S | Valor da base de calculo do FCP retido anteriormente por ST |
| 155 | PERFCPSTRET | NUMBER(12,4) | S | Percentual do FCP retido anteriormente por Substituição Tributaria |
| 156 | VLFCPSTRET | NUMBER(18,6) | S | Valor do FCP retido por Substituição Tributária |
| 157 | PERCREDBASEEFET | NUMBER(8,4) | S | Percentual de redução de base de calculo efetiva |
| 158 | VLBASEEFET | NUMBER(18,6) | S | Valor da base de calculo efetiva |
| 159 | PERCICMSEFET | NUMBER(8,4) | S | Percentual do ICMS efetivo |
| 160 | VLICMSEFET | NUMBER(18,6) | S | Valor do ICMS efetivo |
| 161 | NUMTRANSULTENTET | NUMBER(10) | S | Transação de entrada de transferência |
| 162 | NUMTRANSULTENTED | NUMBER(10) | S | Transação de entrada de devolução |
| 163 | ALIQICMS1 | NUMBER(12,4) | S | Alíquota interna de ICMS para cálculo de ST. |
| 164 | NUMERO | NUMBER(9,2) | S |  |
| 165 | NUMEROCX | NUMBER(9,2) | S | Indica o numero do endereço do produto. |
| 166 | NUMTRANSULTENTST | NUMBER(16) | S | Transação da última entrada com ST Retido |
| 167 | NUMERO_AUX | NUMBER(7) | S |  |
| 168 | NUMEROCX_AUX | NUMBER(7) | S |  |
| 169 | ESTBONIFIC | VARCHAR2(1) | S | ESTOQUE BONIFICADO |
| 170 | VLICMSBCR | NUMBER(18,6) | S | Valor de ICMS Substituto. |
| 171 | BASEICMSBCR | NUMBER(18,6) | S | Base de ICMS Substituto. |
| 172 | TIPOCURVA | VARCHAR2(1) | S | Tipo de Curva para giro customizado |
| 173 | CURVA | VARCHAR2(250) | S | Chave da Curva para giro customizado |
| 174 | SUBCURVA | VARCHAR2(250) | S | Chave da Sub-Curva para giro customizado |
| 175 | FREQUENCIA | VARCHAR2(250) | S | Chave da Frequencia para giro customizado |
| 176 | NUMTRANSENTULTENTBNF | NUMBER(12) | S | Última transação de entrada bonificada |
| 177 | QTCROSSDOCK | NUMBER(22,8) | S | Saldo Estoque em Crossdocking |
| 178 | QTESTOQUEEMTERCEIRO | NUMBER(22,8) | S | Quantidade de estoque em terceiros. |
| 179 | QTESTOQUEDETERCEIRO | NUMBER(22,8) | S | Quantidade de estoque de terceiros. |
| 180 | QTTRANSITOTV10 | NUMBER(22,8) | S | Quantidade em transito com o tipo de venda igual a 10 - Transferência. |
| 181 | DTHORAULTSAIDA | DATE | S | Data e hora da última movimentação de saida do estoque |
| 182 | DTHORAULTALTDISP | DATE | S | Data e hora da ultima alteração do disponível 	 |
| 183 | QBCMONORET | NUMBER(21,10) | S | Informar a base de cálculo do ICMS retida anteriormente em quantidade (tributada) conforme unidade de medida estabelecida na legislação. |
| 184 | ADREMICMSRET | NUMBER(7,4) | S | Alíquota ad rem do imposto retido anteriormente estabelecida em legislação para o produto. |
| 185 | VICMSMONORET | NUMBER(23,10) | S | Valor do ICMS retido anteriormente obtido pela multiplicação da alíquota ad rem pela quantidade do produto conforme unidade de medida estabelecida em legislação. |
| 186 | VLIPISUSPENSOULTENT | NUMBER(18,6) | S | Valor de IPI Suspenso da Ultima Entrada |
| 187 | VLIISUSPENSOULTENT | NUMBER(18,6) | S | Valor de II Suspenso da Ultima Entrada |
| 188 | PLIQULTENT | NUMBER(18,6) | S | Preço de compra da última entrada. |
| 189 | CUSTOULTENTFISCAL | NUMBER(18,6) | S | Custo Ultima Entrada Fiscal |
| 190 | CUSTOFISCAL | NUMBER(18,6) | S | Custo Medio Fiscal |
| 191 | PERICMSANTECIPADOULTENT | NUMBER(12,4) | S | PERCENTUAL DE ICMS ANTECIPADO DA ULTIMA ENTRADA |
| 192 | PERICMULTENT | NUMBER(12,4) | S | PERCENTUAL DE ICMS DA ULTIMA ENTRADA |
| 193 | QTTRANSITOBENEFIC | NUMBER(22,8) | S | Quantidade de estoque em trânsito específica para beneficiamento |

## Chaves e índices

- Constraint PCEST_PK (P): CODFILIAL, CODPROD
- Índice PCEST_AVARIA: CODPROD, CODFILIAL
- Índice PCEST_IDX2: CODPROD
- Índice PCEST_IDX3: CODPROD, CODFILIAL, NUMTRANSENTULTENT
- Índice PCEST_PK UNIQUE: CODFILIAL, CODPROD
