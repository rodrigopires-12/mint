# RIBERDOCES.PCMOVCOMPLE — dicionário

Gerado por `gateway inventory` em 2026-10-25. 587 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | NUMTRANSITEM | NUMBER(18) | N | Número do item relacionado a tabela PCMOV |
| 2 | DTREGISTRO | DATE | S | Data de criação do registro |
| 3 | PERPISCALCDI | NUMBER(8,4) | S | Percentual de PIS calculo da D.I. |
| 4 | PERCOFINSCALCDI | NUMBER(8,4) | S | Percentual de COFINS calculo da D.I. |
| 5 | EXTIPI | VARCHAR2(3) | S | Exceção da TIPI para geração do XML NFe |
| 6 | PISRETIDO | NUMBER(5,2) | S | Percentual PIS retido venda. |
| 7 | COFINSRETIDO | NUMBER(5,2) | S | Percentual COFINS retido venda. |
| 8 | IRRETIDO | NUMBER(5,2) | S | Percentual IR retido venda. |
| 9 | CSLLRETIDO | NUMBER(5,2) | S | Percentual CSLL retido venda. |
| 10 | VLPISRETIDO | NUMBER(18,6) | S | Valor PIS retido venda. |
| 11 | VLCOFINSRETIDO | NUMBER(18,6) | S | Valor COFINS retido venda. |
| 12 | VLIRRETIDO | NUMBER(18,6) | S | Valor IR retido venda. |
| 13 | VLCSLLRETIDO | NUMBER(18,6) | S | Valor CSLL retido venda. |
| 14 | VLREDCMVSIMPLESNAC | NUMBER(18,6) | S | Valor red. cmv simples nacional. |
| 15 | VLREDPVENDASIMPLESNA | NUMBER(18,6) | S | Valor de redução no preço de venda para cliente Simples Nacional. |
| 16 | DTESTORNO | DATE | S | Data de Estorno da Movimentação. |
| 17 | CUSTONFSEMSTGUIAULTENT | NUMBER(18,6) | S | Custo da NF sem ST Guia atual. |
| 18 | VLSTULTENTANT | NUMBER(18,6) | S | Valor do ST anterior. |
| 19 | VLSTGUIAULTENTANT | NUMBER(18,6) | S | Valor do ST GUIA anterior. |
| 20 | CUSTONFSEMSTGUIAULTENTANT | NUMBER(18,6) | S | Custo da NF sem ST Guia anterior. |
| 21 | IVAULTENTANT | NUMBER(12,4) | S | Valor do IVA anterior. |
| 22 | ALIQICMS1ULTENTANT | NUMBER(12,4) | S | Alíquota ICMS interna anterior. |
| 23 | ALIQICMS2ULTENTANT | NUMBER(12,4) | S | Alíquota ICMS externa anterior. |
| 24 | REDBASEIVAULTENTANT | NUMBER(18,6) | S | Redução da base cálculo ST anterior. |
| 25 | PERCICMSFRETEFOBSTULTENTANT | NUMBER(12,4) | S | Percentual ICMS frete FOB ST anterior. |
| 26 | VLFRETECONHECULTENTANT | NUMBER(18,6) | S | Valor frete FOB anterior. |
| 27 | PERCALIQEXTGUIAULTENTANT | NUMBER(12,4) | S | Percentual alíquota externa GUIA anterior. |
| 28 | VLICMS | NUMBER(18,6) | S | Valor do ICMS. |
| 29 | PERDIFEREIMENTOICMS | NUMBER(8,4) | S | Percentual diferimento ICMS. |
| 30 | VLICMSDIFERIDO | NUMBER(18,6) | S | Valor ICMS diferido. |
| 31 | VLPAUTAIPI | NUMBER(18,6) | S | Informar valor da pauta do IPI na tabela PCMOVCOMPLE. |
| 32 | STBCRANT | NUMBER(18,6) | S | ST Bcr anterior. |
| 33 | BASEBCRANT | NUMBER(18,6) | S | Base ST BCR anterior. |
| 34 | VLPEDAGIO | NUMBER(18,6) | S | Valor do pedágio sobre o item. |
| 35 | PERCPEDAGIO | NUMBER(12,4) | S | Percentual do pedágio sobre o item. |
| 36 | CODFILIALDEST | VARCHAR2(2) | S | Código da filial/depósito de destino para armazenagem |
| 37 | DEDUZIRICMSIMPORTACAO | VARCHAR2(1) | S | Somar ICMS ao custo unitário na primeira NF importação "N". |
| 38 | CODSITTRIBIPI | NUMBER(3) | S | Situação Tributária do IPI |
| 39 | APLICREDBASEIVAPLIQ | VARCHAR2(1) | S | Aplicar redução base IVA preço liquido |
| 40 | APLICREDBASEIVAPLIQBCR | VARCHAR2(1) | S | Aplicar redução base IVA preço liquido BCR |
| 41 | CALCTRIBMANUAL | VARCHAR2(1) | S | Parâmetro para informar se calculo da tributação vai ser manual ou automática |
| 42 | PRODDESCRICAODANFE | VARCHAR2(500) | S | Descrição produto Danfe |
| 43 | PERCMVAORIG | NUMBER(12,4) | S | Percentual de MVA Original para cálculo do MVA ajustado |
| 44 | VLPAUTAPISCOFINS | NUMBER(18,6) | S | Valor pauta PIS/COFINS |
| 45 | USAPISCOFINSLIT | VARCHAR2(1) | S | Usa PIS/COFINS por litragem |
| 46 | PISCOFINSDEDUZIDO | VARCHAR2(1) | S | PIS/COFINS deduzido do valor do produto |
| 47 | VLRICMSSIMPLESNAC | NUMBER(18,6) | S | Valor de ICMS Simples Nacional |
| 48 | PERCICMSSIMPLESNAC | NUMBER(8,4) | S | Percentual Simples Nacional |
| 49 | CODTRIBPISCOFINS | NUMBER(4) | S | Código da tributação de PIS/COFINS |
| 50 | PERCALTERCUSTOENT | NUMBER(6,2) | S | Percentual de acréscimo no custo do produto por não reembolso da avaria. |
| 51 | VLDESCORGAOPUB | NUMBER(8,2) | S | Desconto para orgão público. |
| 52 | CODFORNEC | NUMBER(6) | S |  |
| 53 | IBANC | VARCHAR2(20) | S |  |
| 54 | PRACA | VARCHAR2(20) | S |  |
| 55 | QTPERDA | NUMBER(14,8) | S | Campo para armazenar a quantidade de insumos que foram perdidas durante o processo de beneficiamento. |
| 56 | SUBSTANCIA | VARCHAR2(10) | S | Gravar a substancia do produto para obter um controle especial |
| 57 | VLBASEFRETE | NUMBER(18,6) | S | Valor da base do frete do item |
| 58 | VLBASEOUTROS | NUMBER(18,6) | S | Valor da base da outras despesas do item |
| 59 | PRECOMAXCONSUM | NUMBER(18,6) | S | Preço máximo consumidor |
| 60 | DESCPRECOFAB | NUMBER(18,6) | S | Desconto Preço de Fábrica |
| 61 | PORIGINAL | NUMBER(18,6) | S | Preço original |
| 62 | PERACRESCIMOFUNCEP | NUMBER(8,4) | S | Percentual de Acréscimo do Fundo de Combate e Erradicação da Pobreza |
| 63 | VLACRESCIMOFUNCEP | NUMBER(20,8) | S | Valor de Acréscimo do Fundo de Combate e Erradicação da Pobreza |
| 64 | VLOUTROSCUSTOS | NUMBER(18,6) | S | Valor de outros custos do item de importação. |
| 65 | BONIFIC | VARCHAR2(1) | S | INFORMAR SE O PRODUTO É BONIFICAÇÃO OU NÃO |
| 66 | PERCCAPATAZIA | NUMBER(18,6) | S | Percentual da capatazia rateado entre os itens de importação. |
| 67 | VLCAPATAZIA | NUMBER(18,6) | S | Valor da capatazia rateado entre os itens de importação. |
| 68 | USAPMCBASEST | VARCHAR2(1) | S | Usa PMC base ST |
| 69 | PERCREDPMC | NUMBER(12,4) | S | Perc.redução PMC |
| 70 | CODPRODACABCESTA | NUMBER(6) | S | Código de produto acabado da cesta |
| 71 | EANCODPROD | NUMBER(14) | S | EAN Cód. Prod. |
| 72 | PERCFRETENFCUSTO | NUMBER(12,4) | S | Percentual de frete NF custo |
| 73 | VLFRETENFCUSTO | NUMBER(18,6) | S | Valor do frete NF custo |
| 74 | PERCSEGUROCUSTO | NUMBER(12,4) | S | Percentual de seguro NF custo |
| 75 | VLSEGUROCUSTO | NUMBER(18,6) | S | Valor do seguro NF custo |
| 76 | PERCCAPATAZIACUSTO | NUMBER(18,6) | S | Percentual de capatazia NF custo |
| 77 | VLCAPATAZIACUSTO | NUMBER(18,6) | S | Valor da capatazia NF custo |
| 78 | PERCDESPDENTRONFCUSTO | NUMBER(12,4) | S | Percentual de acréscimo NF custo |
| 79 | VLDESPDENTRONFCUSTO | NUMBER(18,6) | S | Valor do acréscimo NF custo |
| 80 | PERCOUTRASDESPIMP | NUMBER(12,4) | S | Percentual de outras despesas importação |
| 81 | TIPOCALCST | VARCHAR2(1) | S | Tipo do calculo do ST |
| 82 | IPIPORVALOR | VARCHAR2(1) | S | IPI com valor fixo |
| 83 | PERCADUANEIRA | NUMBER(12,4) | S | Percentual de despesas aduaneira |
| 84 | PERCSISCOMEX | NUMBER(12,4) | S | Percentual de taxas com siscomex |
| 85 | UTILIZACREDREDPISCOFINS | VARCHAR2(1) | S | Alíquota crédito reduzido PIS/COFINS |
| 86 | QTLITRAGEM | NUMBER(18,6) | S | Quantidade litragem |
| 87 | NUMDRAWBACK | VARCHAR2(20) | S | Número do Ato Concessório do regime |
| 88 | ALIQICMSFECP | NUMBER(8,4) | S | Alíquota de ICMS Fecp. |
| 89 | VLFECP | NUMBER(18,6) | S | Valor Fecp. |
| 90 | VLICMSANTECIPADO | NUMBER(18,6) | S | Valor do ICMS Antecipado |
| 91 | PERCDESCINDUSTRIA | NUMBER(18,6) | S | % desconto industria |
| 92 | CODLINHAPRAZO | NUMBER(6) | S | Código Linha Prazo |
| 93 | VLDESCONTOCUSTO | NUMBER(18,6) | S | Valor das Deduções Custo |
| 94 | PTABELACUSTO | NUMBER(18,6) | S | Preço de Compra Custo |
| 95 | PERCSISCOMEXCUSTO | NUMBER(12,4) | S | Percentual Siscomex Custo |
| 96 | VLSISCOMEXCUSTO | NUMBER(18,6) | S | Valor Siscomex Custo |
| 97 | PERCADUANEIRACUSTO | NUMBER(12,4) | S | Percentual Despesa Aduaneira Custo |
| 98 | VLADUANEIRACUSTO | NUMBER(18,6) | S | Valor Despesa Aduaneira Custo |
| 99 | PERCOUTRASDESPIMPCUSTO | NUMBER(12,4) | S | Percentual Outras Despesas Importação Custo |
| 100 | VLOUTRASDESPIMPCUSTO | NUMBER(18,6) | S | Valor Outras Despesas Importação Custo |
| 101 | PERCOUTROSCUSTOSCUSTO | NUMBER(12,4) | S | Percentual Outros Custos Importação Custo |
| 102 | VLOUTROSCUSTOSCUSTO | NUMBER(18,6) | S | Valor Outros Custos Importação Custo |
| 103 | CUSTOULTENTCONT | NUMBER(22,6) | S | Custo da última entrada contábil - entrada |
| 104 | CUSTOULTENTCONTANT | NUMBER(22,6) | S | Custo da última entrada contábil - anterior |
| 105 | BASEICMSANTECIPADO | NUMBER(18,6) | S | Base do ICMS Antecipado |
| 106 | BASECREDPRESUMIDO | NUMBER(18,6) | S | Base do Crédito Presumido |
| 107 | PERCCARGATRIBMEDIA | NUMBER(18,6) | S | Percentual de carga tributária média, utilizado no SEFAZ MT para calculo da substituição tributária |
| 108 | IDVENDA | NUMBER(6) | S | Identificação de Venda |
| 109 | PARTICIPAGIRO | VARCHAR2(1) | S | Parcicipa Giro |
| 110 | NUMLISTA | NUMBER(6) | S | Número da Lista de Presente |
| 111 | NUMSEQITEMCONTRATO | NUMBER(6) | S | Número da Sequência do Item no Contrato |
| 112 | PERCIPIOUTRAS | NUMBER(12,4) | S | Percentual de IPI |
| 113 | VLIPIOUTRAS | NUMBER(18,6) | S | Valor do IPI |
| 114 | BASEIPIOUTRAS | NUMBER(18,6) | S | Base do IPI |
| 115 | PERCSTOUTRAS | NUMBER(12,4) | S | Percentual de ST |
| 116 | VLSTOUTRAS | NUMBER(18,6) | S | Valor do ST |
| 117 | BASESTOUTRAS | NUMBER(18,6) | S | Base do ST |
| 118 | VLDESCABATIMENTO | NUMBER(18,6) | S | GRAVAR O VALOR DO DESCONTO DO ABATIMENTO |
| 119 | CALCCREDIPICONT | VARCHAR2(1) | S | Calcula crédito IPI custo contabil na importação |
| 120 | APROVEITACREDICMSCONT | VARCHAR2(1) | S | Aproveita crédito ICMS custo contabil na importação |
| 121 | APROVEITACREDPISCOFINSCONT | VARCHAR2(1) | S | Aproveita crédito PIS/COFINS custo contabil na importação |
| 122 | VLBASEPISCOFINSBKP | NUMBER(18,6) | S | Backup da base de PIS/COFINS |
| 123 | PERPISBKP | NUMBER(12,4) | S | Backup da alíquota de PIS |
| 124 | VLCREDPISBKP | NUMBER(18,6) | S | Backup do valor do PIS |
| 125 | PERCOFINSBKP | NUMBER(12,4) | S | Backup da alíquota de COFINS |
| 126 | VLCREDCOFINSBKP | NUMBER(18,6) | S | Backup do valor do COFINS |
| 127 | GERABASEPISCOFINSSEMALIQ | VARCHAR2(1) | S | Defina se deve gerar base de PIS/CONFINS mesmo quando não for informado aliquotas de PIS/CONFINS |
| 128 | PERCICMSBASEICMSANTECIPADO | NUMBER(12,4) | S | Aliq. Para base do ICMS antecipado |
| 129 | NUMTRANSITEMDEV | NUMBER(10) | S | Número de transação do item de devolução. Permite rastrear o item da venda apartir do item devolvido. |
| 130 | CODSITTRIBST | VARCHAR2(2) | S | GRAVAR O CÓDIGO DA SITUAÇÃO TRIBUTÁRIA ST |
| 131 | PERCDESCABATIMENTO | NUMBER(18,6) | S | Percentual de abatimento |
| 132 | QTUN | NUMBER(16) | S | Quantidade em unidades de venda |
| 133 | NUMINVENT | NUMBER(8) | S | Número do Inventário |
| 134 | ICMSPAUTA | VARCHAR2(1) | S | Indica se a base de ICMS foi pauta. |
| 135 | CODEMITENTEITEMPEDIDO | NUMBER(8) | S | Codigo emitente que incluiu o item no pedido |
| 136 | VLBASEIPISUSPENSO | NUMBER(18,6) | S | Base do IPI suspenso |
| 137 | VLIPISUSPENSO | NUMBER(18,6) | S | Valor do IPI suspenso |
| 138 | PERCIPISUSPENSO | NUMBER(12,4) | S | Percentual do IPI suspenso |
| 139 | PERDESCINICOMISS | NUMBER(12,6) | S | Perc. Desc. Inicial Faixa Comissão. |
| 140 | PERDESCFIMCOMISS | NUMBER(12,6) | S | Perc. Desc. Final Faixa Comissão. |
| 141 | CONCEDERMAIORCOMISSREG | VARCHAR2(1) | S | Conceder maior Comissão das Faixas. |
| 142 | ANP | NUMBER(9) | S | Código ANP combustível. |
| 143 | UNIDADETRIB | VARCHAR2(6) | S | Unidade tributável. |
| 144 | FATORCONVTRIB | NUMBER(12,6) | S | Fator conversão tributável. |
| 145 | CONSIISUSPENSOBASEICMS | VARCHAR2(1) | S | Considera II Suspenso base ICMS. |
| 146 | CONSIPISUSPENSOBASEICMS | VARCHAR2(1) | S | Considera IPI Suspenso base ICMS. |
| 147 | VLIISUSPENSO | NUMBER(18,6) | S | Valor do II Suspenso. |
| 148 | VLALTERCUSTOENT | NUMBER(18,6) | S | Acréscmo no Custo da Última Entrada. |
| 149 | VLSUBTOTITEM | NUMBER(18,6) | S | Valor SubTotal do Item - ECF |
| 150 | REGIMEESPECIAL | VARCHAR2(1) | S | Regime Especial. |
| 151 | CONSIDERASTNFCUSTO | VARCHAR2(1) | S | Considera ST NF  no Custo. |
| 152 | CONSIDERASTGUIACUSTO | VARCHAR2(1) | S | Considera ST Guia no Custo. |
| 153 | CONSIDERASTNFCUSTOCONT | VARCHAR2(1) | S | Considera ST NF no Custo Contábil. |
| 154 | CONSIDERASTGUIACUSTOCONT | VARCHAR2(1) | S | Considera ST Guia no Custo Contábil. |
| 155 | UTILIZABASEPISCOFINSSUSP | VARCHAR2(1) | S | Utiliza Base PIS/COFINS Suspenso. |
| 156 | VLSUBTOTDESCONTO | NUMBER(12,2) | S | Valor do subtotal de desconto do item. |
| 157 | NUMFCI | VARCHAR2(36) | S | Ficha de Conteúdo de Importação. |
| 158 | VLPARCELAIMPFCI | NUMBER(18,6) | S | Valor da Parcela Importada. |
| 159 | PERCCONTEUDOIMPFCI | NUMBER(5,2) | S | Conteúdo de Importação. |
| 160 | VLIMPORTACAOFCI | NUMBER(18,6) | S | Valor de Importação. |
| 161 | CONTROLARFCI | VARCHAR2(1) | S | Controla Ficha de Conteúdo de Importação. |
| 162 | CODPRODGNRE | NUMBER(6) | S | CÓDIGO DO PRODUTO PARA GNRE. |
| 163 | VLITEMTRIBUTOS | NUMBER(18,6) | S | VALOR DOS TRIBUTOS POR ITEM. |
| 164 | ENDERECADO | VARCHAR2(1) | S | Informa se o produto já foi endereçado. |
| 165 | CUSTOFINSEMST | NUMBER(18,6) | S | Valor do Custo médio Financeiro sem ST NF e GUIA. |
| 166 | CUSTOFINSEMSTANT | NUMBER(18,6) | S | Valor do Custo médio Financeiro sem ST NF e GUIA Anterior. |
| 167 | QTDIGITADA | NUMBER(18,6) | S | Quantidade digitada pelo usúario. |
| 168 | VLDESCSOCIOTORCEDOR | NUMBER(18,2) | S | Desconto Sóscio Torcedor |
| 169 | SUJDESONERACAO | VARCHAR2(1) | S | Este parâmetro ira dizer se o produto esta ou não sujeito a desoneração |
| 170 | CNAEDESONERACAO | VARCHAR2(60) | S | Código CNAE sera utilizado no recálculo desoneração na rotina 1000 |
| 171 | PERCDESONERACAO | NUMBER(10,2) | S | Percentual de desoneração será utilizado no recálculo de desoneração na rotina 1000 |
| 172 | NUMRECOPI | NUMBER(20) | S | Número sujeito ao controle de papel imune |
| 173 | NUMERORECOPI | VARCHAR2(20) | S | Número sujeito ao controle de papel imune |
| 174 | CUSTOULTENTFINSEMST | NUMBER(18,6) | S | Custo Ult.Entrada Financeira sem ST. |
| 175 | CUSTOULTENTSEMST | NUMBER(18,6) | S | Custo Ult.Entrada sem ST. |
| 176 | CUSTOULTENTFINSEMSTANT | NUMBER(18,6) | S | Custo Ult.Entrada Financeira sem ST Anterior |
| 177 | CUSTOULTENTSEMSTANT | NUMBER(18,6) | S | Custo Ult.Entrada sem ST Anterior |
| 178 | VLBASEPISCOFINSFRETEFOB | NUMBER(18,6) | S | Valor base PIS/COFINS Frete FOB |
| 179 | PERPISFRETEFOB | NUMBER(18,6) | S | Percentual PIS Frete FOB |
| 180 | VLCREDPISFRETEFOB | NUMBER(18,6) | S | Valor PIS Frete FOB |
| 181 | PERCOFINSFRETEFOB | NUMBER(18,6) | S | Percentual COFINS Frete FOB |
| 182 | VLCREDCOFINSFRETEFOB | NUMBER(18,6) | S | Valor COFINS Frete FOB |
| 183 | CODSITTRIBPISCOFINSFRETEFOB | NUMBER(3) | S | Código Situação Tributária PIS/COFINS Frete FOB |
| 184 | VLBASEICMSFRETEFOB | NUMBER(18,6) | S | Valor base ICMS Frete FOB |
| 185 | PERCICMSFRETEFOB | NUMBER(18,6) | S | Percentual ICMS Frete FOB |
| 186 | PERCICMSCUSTOFRETEFOB | NUMBER(18,6) | S | Percentual ICMS p/ calculo do custo Frete FOB |
| 187 | VLICMSFRETEFOB | NUMBER(18,6) | S | Valor do ICMS Frete FOB |
| 188 | VLICMSCUSTOFRETEFOB | NUMBER(18,6) | S | Valor do ICMS p/ calculo do custo Frete FOB |
| 189 | CODSITTRIBICMSFRETEFOB | VARCHAR2(3) | S | Código Situação Tributária ICMS Frete FOB |
| 190 | PERCAFRMM | NUMBER(12,4) | S | Percentual de Adicional de Frete para a Renovação da Marinha Mercante |
| 191 | VLAFRMM | NUMBER(18,6) | S | Valor do Adicional de Frete para a Renovação da Marinha Mercante |
| 192 | PERCAFRMMCUSTO | NUMBER(12,4) | S | Percentual de Adicional de Frete para a Renovação da Marinha Mercante |
| 193 | VLAFRMMCUSTO | NUMBER(18,6) | S | Valor do Adicional de Frete para a Renovação da Marinha Mercante |
| 194 | ORIGMERCTRIB | VARCHAR2(1) | S | Código de origem da mercadoria na situação tributária |
| 195 | VLPISCALCDI | NUMBER(18,6) | S | Valor do PIS suspenso |
| 196 | VLCOFINSCALCDI | NUMBER(18,6) | S | Valor do COFINS suspenso |
| 197 | USAUNIDADEMASTER | VARCHAR2(1) | S | Indica se o produto foi gravado para respeitar a quantidade master ou não. |
| 198 | NUMTRANSTRANSFORIG | NUMBER(10) | S | Transação da PCMOV das transferencias virtuais. |
| 199 | VERSAOREGISTRO | VARCHAR2(10) | S | Define a versão do sistema que o registro foi gravado. |
| 200 | NUMTRANSNUMEROSERIE | NUMBER(10) | S | Número da transação de movimentação |
| 201 | PERCBONIFIC | NUMBER(18,6) | S | % Bonificação Dinheiro |
| 202 | VLADICIONALBCICMSANTECIP | NUMBER(18,6) | S | Valor adicional base Icms Antecipado |
| 203 | APLICPERCIVAPAUTAICMSANTECIP | VARCHAR2(1) | S | Aplica IVA sobre valor de pauta ICMS antecipado |
| 204 | PERCICMSFRETEFOBICMSANTECIP | NUMBER(12,4) | S | % ICMS Frete FOB p/ calc.Icms Antecipado |
| 205 | PERCMVAORIGICMSANTECIP | NUMBER(12,4) | S | % IVA Original ICMS Antecipado |
| 206 | PERCCARGATRIBMEDIAICMSANTECIP | NUMBER(18,6) | S | % Carga Média tributária Icms Antecipado |
| 207 | REDBASEIVAICMSANTECIP | NUMBER(18,6) | S | % Redução IVA Icms Antecipado |
| 208 | REDBASEALIQEXTICMSANTECIP | NUMBER(18,6) | S | % Redução Aliquota externa ICMS Antecipado. |
| 209 | PERCANTIDUMPING | NUMBER(12,4) | S | Percentual Antidumping DI |
| 210 | VLANTIDUMPING | NUMBER(18,6) | S | Valor Antidumping DI |
| 211 | PERCANTIDUMPINGCUSTO | NUMBER(12,4) | S | Percentual Antidumping Custo |
| 212 | VLANTIDUMPINGCUSTO | NUMBER(18,6) | S | Valor Antidumping Custo |
| 213 | VLITEMTRIBUTOSEST | NUMBER(18,6) | S | Valor dos tributos estaduais do item |
| 214 | VLITEMTRIBUTOSMUNIC | NUMBER(18,6) | S | Valor dos tributos municipais do item |
| 215 | UNIDADELICIT | VARCHAR2(2) | S | Unidade Licitação |
| 216 | QTUNITLICIT | NUMBER(22,8) | S | Quantidade unidade licitação |
| 217 | CODMOTIVOICMSDESONERADO | VARCHAR2(2) | S | Código do Motivo da Desoneração do ICMS |
| 218 | VLICMSDESONERACAO | NUMBER(18,6) | S | Valor do ICMS Desonerado |
| 219 | CODFIGURADESONICMS | NUMBER(8) | S | Código de figura desoneração de ICMS |
| 220 | PERCREDALIQIPI | NUMBER(18,6) | S | Percentual de Redução da Alíquota de IPI |
| 221 | FATORCONVCOM | NUMBER(12,6) | S | Fator conversão comercial |
| 222 | PRECOUTILIZADONFE | VARCHAR2(2) | S | Tipo de preço utilizado na NFE |
| 223 | ALIQREDUCAOPIS | NUMBER(12,4) | S | % de Redução de Alíquota de PIS |
| 224 | ALIQREDUCAOCOFINS | NUMBER(12,4) | S | % de Redução Alíquota da COFINS |
| 225 | ENVIARALIQREDUCAOPISCOFINS | VARCHAR2(1) | S | Enviar informações de % redução de aliq pis e % de redução da confins em infAdProd |
| 226 | NUMSERIESAT | VARCHAR2(13) | S | Número de série equipamento SAT |
| 227 | VLICMSCHEIO | NUMBER(18,6) | S | Valor do icms sem diferimento |
| 228 | NUMPEDCLI | VARCHAR2(15) | S | Número do pedido do cliente |
| 229 | PERCICMSREDFRETEFOB | NUMBER(18,6) | S | Percentual de Redução de ICMS Frete FOB |
| 230 | PERCCARGEFETIV | NUMBER(6,2) | S | % de carga efetiva do produto |
| 231 | CODCEST | VARCHAR2(7) | S |  |
| 232 | CODENQIPI | VARCHAR2(3) | S |  |
| 233 | DESCRICAONFCOMP | VARCHAR2(120) | S | Descrição para notas complementares |
| 234 | VLSTDEVFORNEC | NUMBER(18,6) | S | Valor de ST para devoluções fornecedor ressarcimento |
| 235 | VLBASESTDEVFORNEC | NUMBER(18,6) | S | Valor da base ST para devoluções fornecedor ressarcimento |
| 236 | PERCICMS | NUMBER(8,4) | S | Percentual de ICMS |
| 237 | CODCNAESERV | NUMBER(20) | S | Código CNAE do Serviço |
| 238 | CODTIPOSERVICO | VARCHAR2(9) | S | Código do serviço conforme tabela 06 do layout do SPED |
| 239 | PERCRETSOBREVLTOTALCPRBSERV | NUMBER(18,6) | S | Percentual sobre valor total CPRB servico |
| 240 | PERCRETCPRBSERV | NUMBER(18,6) | S | Percentual retenção do CPRB sobre o serviço |
| 241 | PERCRETSOBREVLTOTALINCISERV | NUMBER(18,6) | S | Percentual sobre valor total de incidência do serviço |
| 242 | TERMPOSERVICO | NUMBER(6) | S | Tempo adicional do serviço |
| 243 | RETERINSSCPRBSERV | VARCHAR2(2) | S | Reter INSS-CPRB serviço |
| 244 | RETERVALORINCSER | VARCHAR2(2) | S | Reter valor de incidência do serviço |
| 245 | TIPOPROCESSOSERV | VARCHAR2(2) | S | Tipo de processo do serviço |
| 246 | NUMEROPROCESSOSERV | VARCHAR2(100) | S | Número do processo do serviço |
| 247 | PERCRETINCIDENCIASERV | NUMBER(18,6) | S | Percentual de retenção de incidência do serviço |
| 248 | VLFCPPART | NUMBER(22,10) | S | Valor do FUNCEP |
| 249 | VLICMSPART | NUMBER(22,10) | S | Valor de ICMS de partilha (valor que deverá ser utilizado para acrescer o valor do produto) |
| 250 | VLBASEPARTDEST | NUMBER(22,10) | S | Valor da base de calculo na Uf de destino |
| 251 | ALIQFCP | NUMBER(18,6) | S | Aliquota de FCP (fundo de combate a pobreza) |
| 252 | ALIQINTERNADEST | NUMBER(18,6) | S | Aliquota de ICMS interna na UF de destino |
| 253 | VLICMSPARTDEST | NUMBER(22,10) | S | Valor do ICMS Interestadual para a UF Destino |
| 254 | PERCPROVPART | NUMBER(5,2) | S | Percentual provisório de partilha de ICMS |
| 255 | VLICMSDIFALIQPART | NUMBER(26,10) | S | Valor de ICMS do diferencial de aliquota da partilha de ICMS |
| 256 | PERCBASEREDPART | NUMBER(5,2) | S | Redução aplicada na base de Partilha na nota |
| 257 | VLPISRETORGPUB | NUMBER(18,6) | S | Valor do PIS que será retido pelo orgão publico |
| 258 | VLCOFINSRETORGPUB | NUMBER(18,6) | S | Valor do Cofins que será retido pelo orgão publico |
| 259 | VLIRPJRETORGPUB | NUMBER(18,6) | S | Valor do IRPJ que será retido pelo orgão publico |
| 260 | VLCSLLRETORGPUB | NUMBER(18,6) | S | Valor do CSLL que será retido pelo orgão publico |
| 261 | VLICMSPARTREM | NUMBER(22,10) | S | Valor do ICMS de partilha para UF remetente |
| 262 | ALIQINTERORIGPART | NUMBER(18,6) | S | Valor de ICMS da UF de origem |
| 263 | BNFNAOCALCULAICMS | VARCHAR2(1) | S | Flag da rotina 514 |
| 264 | PESOLIQDI | NUMBER(12,6) | S | Peso Líquido calculo rateiro Importação |
| 265 | QT_ENTRADAMANIF | NUMBER(20,8) | S | Quantidade de entrada que ficou no caminhão |
| 266 | VLBASEPART_DESATIVADO | NUMBER(18,6) | S | Valor da base de calculo na Uf de destino |
| 267 | ALIQFCPPART_DESATIVADO | NUMBER(18,6) | S | Percentual do FUNCEP |
| 268 | ALIQICMSPART_DESATIVADO | NUMBER(18,6) | S | Percentual de ICMS na UF Destino |
| 269 | VLACRESPART_DESATIVADO | NUMBER(18,6) | S | Valor a ser acrescido ao Preço s/imposto(Valor ICMS + ValorFuncep) |
| 270 | VLRETENCAO | NUMBER(18,6) | S | Valor total da retenção para os impostos do sped Reinf |
| 271 | VLRETENCAOSUBCONTRATADA | NUMBER(18,6) | S | Valor da retenção da subcontratada |
| 272 | CNPJSUBCONTRATADA | NUMBER(18) | S | CNPJ da subcontratada para o servico |
| 273 | NOTASUBCONTRATADA | NUMBER(18) | S | Nota para o servico da empresa subcontratada |
| 274 | SERIENOTASUBCONTRATADA | NUMBER(18) | S | Serie da nota da subcontratada |
| 275 | TEMPOSERVICO | NUMBER(18) | S | Tempo adicional do servico cadastrado |
| 276 | NUMITEMPED | NUMBER(18,6) | S | Número de sequencia de item do pedido do cliente |
| 277 | VLIPIPTABELA | NUMBER(18,6) | S | Valor de IPI relativo ao preço de tabela |
| 278 | VLIPIPBASERCA | NUMBER(18,6) | S | Valor de IPI relativo ao preço base RCA |
| 279 | STPTABELA | NUMBER(18,6) | S | Valor de ST relativo ao preço de tabela |
| 280 | STPBASERCA | NUMBER(18,6) | S | Valor de ST relativo ao preço base RCA |
| 281 | VLICMSPARTPTABELA | NUMBER(18,6) | S | Valor ICMS Partilha relativo ao preço de tabela |
| 282 | VLICMSPARTPBASERCA | NUMBER(18,6) | S | Valor ICMS Partilha relativo ao preço base RCA |
| 283 | PERDESCCUSTO | NUMBER(8,4) | S | Percentual de desconto do custo |
| 284 | CFOPORIG | NUMBER(8) | S | Cód. Fiscal original da nota |
| 285 | MODBCST | NUMBER(1) | S | Modalidade da base st nota fiscal |
| 286 | MODBCSTGUIA | NUMBER(1) | S | Modalidade da base st guia |
| 287 | PRECOFAB | NUMBER(18,6) | S | Preço de Fábrica do produto |
| 288 | VLIPIDEVFORNEC | NUMBER(18,6) | S | Valor IPI Devolução fornecedor |
| 289 | NUMSEQORIG | NUMBER(20) | S | Numero de sequencia do pedido original para vendas do autosserviço. |
| 290 | FATORENT | NUMBER(27,15) | S | Fator usado na entrada de mercadoria |
| 291 | MULTIMPLOUNIDADEMASTER | VARCHAR2(1) | S | Representa se a quantidade digitada é ou não multiplo da unidade master |
| 292 | VLBASEPART | NUMBER(18,6) | S | Valor da base de calculo na Uf de destino |
| 293 | ALIQFCPPART | NUMBER(18,6) | S | Percentual do FUNCEP |
| 294 | ALIQICMSPART | NUMBER(18,6) | S | Percentual de ICMS na UF Destino |
| 295 | VLACRESPART | NUMBER(18,6) | S | Valor a ser acrescido ao Preço s/imposto(Valor ICMS + ValorFuncep) |
| 296 | PERCIPIDEVFORNEC | NUMBER(12,4) | S | Percentual de IPI devolução fornecedor |
| 297 | VLBASEIPIDEVFORNEC | NUMBER(18,6) | S | Valor base IPI Devolução Fornecedor |
| 298 | MOSTRARITENSNAEMBMASTER | VARCHAR2(2) | S | Resultado do parametro 2494 - Mostrar Item Apenas na Embalagem Master |
| 299 | FATORCONVEMB | NUMBER(18,6) | S | Fator de conversão da embalagem |
| 300 | UNIDADEEMB | VARCHAR2(2) | S | Unidade da embalagem |
| 301 | BASEIMPORTACAO | NUMBER(18,6) | S | Base do imposto de importação |
| 302 | DTPAGCOMISSAO | DATE | S | Data do Pagamento da Comissão por Item do RCA |
| 303 | DTPAGCOMISSAOOP | DATE | S | Data do Pagamento da Comissão por Item do Operador |
| 304 | CODUSURPGCOMISSAO | NUMBER(8) | S | Usuário que gerou o Pagamento da Comissão por Item do RCA |
| 305 | CODUSURPGCOMISSAOOP | NUMBER(8) | S | Usuário que gerou o Pagamento da Comissão por Item do Operador |
| 306 | VLBASEST | NUMBER(18,6) | S | Valor da Base de ST, que será informada quando o ST for agregado ao valor de outras despesas |
| 307 | IVA | NUMBER(18,6) | S | Valor do IVA, que será informada quando o ST for agregado ao valor de outras despesas |
| 308 | ALIQ1 | NUMBER(18,6) | S | Valor da Alíquota1, que será informada quando o ST for agregado ao valor de outras despesas |
| 309 | VLST | NUMBER(18,6) | S | Valor do ST, para quando o ST for agregado ao valor de outras despesas. |
| 310 | BASEICMS_MANIF | NUMBER(18,6) | S | Base de ICMS que será destacada na NF Manifesto |
| 311 | PERCICM_MANIF | NUMBER(18,6) | S | Percentual de ICMS da NF de devolução de manifesto |
| 312 | VLICMS_MANIF | NUMBER(18,6) | S | Valor de ICMS na devolução de NF Manifesto |
| 313 | VLICMS_MANUAL | NUMBER(18,6) | S | Base de ICMS informado manualmente em uma NF |
| 314 | VLDESCONTONF | NUMBER(18,6) | S | Valor do desconto que será apresentado no DANFe |
| 315 | VERSAOSERVICOPARTILHA | VARCHAR2(10) | S | Versão do serviço de cálculo do ICMS Partilha |
| 316 | PATCHSERVICOPARTILHA | NUMBER(4) | S | Patch do serviço de cálculo do ICMS Partilha |
| 317 | ESFERAPUBLICA | VARCHAR2(1) | S | ESFERA PUBLICA DO CLIENTE (FEDERAL, ESTADUAL, MUNICIPAL) |
| 318 | PRODIMPORTADOPEPS | VARCHAR2(1) | S | Produto Importado PEPS |
| 319 | NUMTRANSENTPEPS | NUMBER(10) | S | Transação de Entrada PEPS |
| 320 | NUMSEQENT | NUMBER(5) | S | Sequência de digitação do item no XML ou entrada. |
| 321 | BASEICMSIMPORIG | NUMBER(18,6) | S | Base ICMS Original da Importação |
| 322 | PERCICMSIMPORIG | NUMBER(12,4) | S | Percentual original da Importação |
| 323 | VLICMSMPORIG | NUMBER(18,6) | S | Valor original da Importação |
| 324 | VLIMPPRODUTORURAL | NUMBER(18,6) | S | Valor imposto Produtor Rural |
| 325 | PERCOUTROSCUSTOSCONT | NUMBER(12,4) | S | Percentual de outros custos contabil |
| 326 | VLOUTROSCUSTOSCONT | NUMBER(18,6) | S | Valor de outros custos contabil |
| 327 | PERCOUTROSCUSTOSCUSTOCONT | NUMBER(12,4) | S | Percentual de outros custos custo contabil |
| 328 | VLOUTROSCUSTOSCUSTOCONT | NUMBER(18,6) | S | Valor de outros custos custo contabil |
| 329 | CODSUPERVISOR | NUMBER(4) | S | Indica o código do SUPERVISOR |
| 330 | MOVEST | VARCHAR2(1) | S | Define se a movimentação irá movimentar ou não o estoque. |
| 331 | QTMINATACVENDA | NUMBER(18,6) | S | Qt. Min de Atacado do produto durante a venda |
| 332 | TIPODESCATACVENDA | VARCHAR2(1) | S | Tipo de desconto de atacado utilizado na venda |
| 333 | VLBASEPISCOFINSNFSERVICO | NUMBER(18,6) | S | Base PIS/COFINS NF Serviço |
| 334 | PERPISNFSERVICO | NUMBER(12,4) | S | Alíquota de PIS NF Serviço |
| 335 | VLCREDPISNFSERVICO | NUMBER(18,6) | S | Valor crédito PIS NF Serviço |
| 336 | PERCOFINSNFSERVICO | NUMBER(12,4) | S | Alíquota de COFINS NF Serviço |
| 337 | VLCREDCOFINSNFSERVICO | NUMBER(18,6) | S | Valor crédito COFINS NF Serviço |
| 338 | CODSITTRIBPISCOFINSNFSERVICO | NUMBER(3) | S | CST de PIS/COFINS NF Serviço |
| 339 | PERCNFSERVICO | NUMBER(12,4) | S | Percentual de NF Serviço por item |
| 340 | VLNFSERVICO | NUMBER(18,6) | S | Valor de NF Serviço por item |
| 341 | VLNFSERVICOPORKG | NUMBER(18,6) | S | Valor em kilos da NF Serviço por item |
| 342 | APLICAMVABASEICMS | VARCHAR2(1) | S | Define se irá aplicou o MVA do ST na base de ICMS |
| 343 | VLBASEFCPICMS | NUMBER(18,6) | S | Valor da base de calculo do Fundo de Combate a Pobreza |
| 344 | VLBASEFCPST | NUMBER(18,6) | S | Valor de base de calculo do Fundo de Combate a Pobreza ST |
| 345 | VLBCFCPSTRET | NUMBER(18,6) | S | Valor da base de calculo do FCP retido anteriormente por ST |
| 346 | PERFCPSTRET | NUMBER(12,4) | S | Percentual do FCP retido anteriormente por Substituição Tributaria |
| 347 | VLFCPSTRET | NUMBER(18,6) | S | Valor do FCP retido por Substituição Tributária |
| 348 | PERFCPSN | NUMBER(12,4) | S | Aliquota aplicável de cálculo do crédito(SIMPLES NACIONAL) |
| 349 | VLCREDFCPICMSSN | NUMBER(18,6) | S | Valor crédito do ICMS que pode ser aproveitado nos termos do art. 23 da LC 123 (SIMPLES NACIONAL) |
| 350 | NUMCHAVEEXP | VARCHAR2(44) | S | Número da Chave de acesso da Nfe recebida para exportação |
| 351 | NUMREGEXP | NUMBER(12) | S | Número do registro de Exportação |
| 352 | INDESCALARELEVANTE | VARCHAR2(1) | S | Indicador de escala relevante - 'S' ou 'N'  |
| 353 | CNPJFABRICANTE | VARCHAR2(18) | S | CNPJ do fabricante da mercadoria |
| 354 | FABRICANTE | VARCHAR2(60) | S | Razão social do fabricante |
| 355 | UNIDADETRIBEX | VARCHAR2(6) | S | Unidade tributável exterior |
| 356 | FATORCONVTRIBEX | NUMBER(12,6) | S | Fator de conversão do exterior a ser utilizado na NF-e |
| 357 | DESCANP | VARCHAR2(95) | S | Descrição do produto conforme ANP |
| 358 | CODBENEFICIONOTA | VARCHAR2(10) | S | Armazena o código de benefício fiscal que veio na nota (xml) |
| 359 | CODAGREGACAO | VARCHAR2(20) | S | Código de agregação do produto |
| 360 | CODBENEFICIOFISCAL | VARCHAR2(10) | S | Código do beneficio fiscal utilizado na UF |
| 361 | CODOFERTA | NUMBER(6) | S | Código da oferta programada utilizado na venda. |
| 362 | MD5PAF | VARCHAR2(200) | S | Assinatura MD5 do registro do PAF |
| 363 | PESOBRUTOFRETE | NUMBER(18,6) | S | Peso bruto do frete FOB |
| 364 | PERCIVA2 | NUMBER(12,4) | S | Alíquota de ST Outorgado |
| 365 | VLSTOUTORGADO | NUMBER(18,6) | S | Valor de ST Outorgado |
| 366 | CODCOMPRADOR | NUMBER(8) | S | Campo referente ao código do comprador |
| 367 | PERCICMSDESONERACAO | NUMBER(12,4) | S | Percentual de ICMS Desoneração |
| 368 | ESTCOMISSAORCA | VARCHAR2(1) | S | Indica se o item tem estorno de comissão ou não. |
| 369 | CAGREG | VARCHAR2(20) | S | Chave para rastreio de lote |
| 370 | QTDERRUBADAPARCIAL | NUMBER(20,6) | S | Define a quantidade parcial da derrubada de carga |
| 371 | CODMOVINSERVIVEL | NUMBER(8) | S | Codigo da movimentação do inservivel |
| 372 | MODALIDADEBASESTFONTE | NUMBER(4) | S | Modalidade da Base de Cálculo do ST Fonte |
| 373 | TIPOENTREGA | VARCHAR2(2) | S | TIPO DE ENTREGA DO ITEM |
| 374 | VLSTOUTRASDESP | NUMBER(18,6) | S | Valor do ST calculado na venda triangular |
| 375 | NUMSEQPEDAGRUPADO | VARCHAR2(50) | S | Sequência de itens agrupados |
| 376 | QTSOBRA | NUMBER(18,6) | S | Quantidade sobra de matéria prima |
| 377 | VLMATERIAPRIMA | NUMBER(18,6) | S | Valor da composição da matéria prima |
| 378 | QTRETORNOTV13 | NUMBER(20,6) | S | Gravar neste campo a quantidade do saldo que não teve geração de TV14 (pcest.qttransitotv13), para a pkg de estoque saber quantos deverá movimentar o estoque. |
| 379 | CODPOLITICADESC | NUMBER(10) | S | Código da Política de Desconto Financeiro |
| 380 | PGLP | NUMBER(7,4) | S | Percentual do GLP derivado de petroleo. |
| 381 | PGNN | NUMBER(7,4) | S | Percentual do Gás Natural Nacional. |
| 382 | PGNI | NUMBER(7,4) | S | Percentual do Gás Natural importado. |
| 383 | VPART | NUMBER(23,10) | S | Deve ser informado neste campo o valor por quilograma sem ICMS. |
| 384 | UTILIZOUMOTORCALCULO | VARCHAR2(1) | S | Indica se o produto teve o seu cálculo de impostos feitos pelo motor de cálculo |
| 385 | STATUSPREDEVOLUCAO | VARCHAR2(2) | S | Status da pré-devolução |
| 386 | VLTAXAPRAZOMEDCMV | NUMBER(18,6) | S | Valor da Taxa de Prazo no CMV |
| 387 | PARTICIPACOMISSGARANTIDA | VARCHAR2(1) | S | Participa da Comissão Garantida |
| 388 | PERCOMREGIAOMED | NUMBER(8,4) | S | Perc. Comissão Região Medicamento |
| 389 | PERCOMREGIAOMED2 | NUMBER(8,4) | S | Perc. Comissão Região Med. RCA 2 |
| 390 | PERCTAXAPRAZOMEDCMV | NUMBER(18,6) | S | Perc. Taxa Prazo CMV Medicamento |
| 391 | PRECOMAXCONSUMLICITTAB | NUMBER(18,6) | S | Preço Máx. Consumidor Licitação |
| 392 | QTDEDOACAOPEDLICIT | NUMBER(20,6) | S | Quantidade Doação Licitação |
| 393 | QTPEDLICIT | NUMBER(20,6) | S | Quantidade Pedido Licitação |
| 394 | REGIMEESPISENSTFONTE | VARCHAR2(1) | S | Regime Especial Isenção ST Fonte |
| 395 | STREPASSEBCRMED | NUMBER(18,6) | S | Valor do ST Destacado como BCR |
| 396 | TIPOCONVERSAOPEDLICIT | NUMBER(2) | S | Tipo de Conversão de Licitação |
| 397 | TIPOMEDICAMENTO | VARCHAR2(1) | S | Tipo de Medicamento |
| 398 | UNIDADECONVERSAOPEDLICIT | VARCHAR2(2) | S | Unidade Produto para Conversão Licitação |
| 399 | VLDESCBOLETOCMV | NUMBER(18,6) | S | Valor Desconto Financeiro no CMV |
| 400 | ANVISANUMREGMED | VARCHAR2(20) | S | Número Registro ANVISA |
| 401 | CLASSEVENDA | VARCHAR2(1) | S | Classe de Venda do Produto |
| 402 | CODCONFIGFUNCEPMED | NUMBER(4) | S | Cód. Configuração FUNCEP Medicamento |
| 403 | CODEDITAL | NUMBER(9) | S | Código do Edital |
| 404 | CODPROMOCAOMED | NUMBER(9) | S | Código da Promoção Medicamento |
| 405 | CODUSUR2 | NUMBER(4) | S | Código do RCA 2 |
| 406 | FATORCONVERSAOPEDLICIT | NUMBER(30,16) | S | Fator Conversão Licitação |
| 407 | INICIOINTERVALODESCQUANT | NUMBER(10,4) | S | Inicio Intervalo Desconto por Qtde |
| 408 | LOTECONTRATO | VARCHAR2(10) | S | Lote do Contrato do Edital |
| 409 | MIXACORDOPARCERIA | VARCHAR2(1) | S | Mix Participante Acordo Parceria |
| 410 | NUMEMPENHO | VARCHAR2(20) | S | Número Empenho Edital |
| 411 | NUMTRANSPEPS | NUMBER(10) | S | Número Transação Entrada PEPS |
| 412 | OBSERVACAOCOMERCIALMED | VARCHAR2(20) | S | Observação Comercial NF Medicamento |
| 413 | USAMAIORVALORPARACALCULOICMS | VARCHAR2(1) | S | Usa maior valor para calculo do ICMS (Pauta de ICMS e preço de compra) |
| 414 | QTINDUSTRIA | NUMBER(20,6) | S | Estoque indústira na movimentação de entrada, saída e cancelamento |
| 415 | OBSDEVOLUCAO | VARCHAR2(100) | S | Observação sobre o item que esta sendo devolvido |
| 416 | CUSTOREALLIQ | NUMBER(18,6) | S | Custo real médio sem acréscimo |
| 417 | CUSTOREALLIQANT | NUMBER(18,6) | S | Custo real médio sem acréscimo anterior |
| 418 | CUSTOULTENTLIQ | NUMBER(18,6) | S | Custo da última entrada real sem acréscimo |
| 419 | CUSTOULTENTLIQANT | NUMBER(18,6) | S | Custo da última entrada real sem acréscimo anterior |
| 420 | CODCONTACONTSPED | VARCHAR2(255) | S | Código da conta contábil que será utilizada no sped |
| 421 | DATACONSOLIDACAOPREFAT | DATE | S | Data consolidação Pré-Faturamento |
| 422 | USABASESTNOFCP | VARCHAR2(1) | S | Base FCP ST - Considera base do ST |
| 423 | PERCREDBASEEFET | NUMBER(8,4) | S | Percentual de redução de base de calculo efetiva |
| 424 | VLBASEEFET | NUMBER(18,6) | S | Valor da base de calculo efetiva |
| 425 | PERCICMSEFET | NUMBER(8,4) | S | Percentual do ICMS efetivo |
| 426 | VLICMSEFET | NUMBER(18,6) | S | Valor do ICMS efetivo |
| 427 | GUID_JSON_DEV | VARCHAR2(37) | S | GUID do JSON do serviço de custo para devolução |
| 428 | PAUTAICMSFRETEFOB | NUMBER(18,6) | S | Valor de pauta do ICMS |
| 429 | CUSTOREALEST | NUMBER(22,6) | S | Correção da rotina que não estava estornando as quantidades quando era estorno da avaria.Custo real estoque |
| 430 | NUMTRANSULTENTETANT | NUMBER(10) | S | Transação de entrada de transferência |
| 431 | NUMTRANSULTENTEDANT | NUMBER(10) | S | Transação de entrada de devolução |
| 432 | VLSERVPRODUCAO | NUMBER(18,6) | S | Valor do serviço para produzir um item |
| 433 | VLSERVPRODUCAOICMS | NUMBER(18,6) | S | Valor do icms para produzir um item |
| 434 | VLSERVPRODUCAOCREDPIS | NUMBER(18,6) | S | Valor do pis para produzir um item |
| 435 | VLSERVPRODUCAOCREDCOFINS | NUMBER(18,6) | S | Valor do cofins para produzir um item |
| 436 | BASESTTRANSFCD | NUMBER(18,6) | S | Base do ST utilizado na transferência do CD Virtual |
| 437 | VLSTTRANSFCD | NUMBER(18,6) | S | Valor do ST utilizado na transferência do CD Virtual |
| 438 | BASEICMSTRANSFCD | NUMBER(18,6) | S | Base do ICMS utilizado na transferência do CD Virtual |
| 439 | VLICMSTRANSFCD | NUMBER(18,6) | S | Valor do ICMS utilizado na transferência do CD Virtual |
| 440 | CODSERVISS | VARCHAR2(4) | S | Código Serviço ISS |
| 441 | CUSTOREPCANC | NUMBER(18,6) | S | Custo médio de reposição de cancelamento |
| 442 | CUSTOREALCANC | NUMBER(18,6) | S | Custo médio real de cancelamento |
| 443 | CUSTOREALSEMSTCANC | NUMBER(18,6) | S | Custo médio real sem st de cancelamento |
| 444 | CUSTOFINCANC | NUMBER(18,6) | S | Custo médio financeiro de cancelamento |
| 445 | CUSTOFINSEMSTCANC | NUMBER(18,6) | S | Custo médio financeiro sem st de cancelamento |
| 446 | CUSTOCONTCANC | NUMBER(18,6) | S | Custo médio contabil de cancelamento |
| 447 | CUSTOREPANTCANC | NUMBER(18,6) | S | Custo médio de resposição anterior ao cancelamento |
| 448 | CUSTOREALANTCANC | NUMBER(18,6) | S | Custo médio real anterior ao cancelamento |
| 449 | CUSTOREALSEMSTANTCANC | NUMBER(18,6) | S | Custo médio real sem st anterior ao cancelamento |
| 450 | CUSTOFINANTCANC | NUMBER(18,6) | S | Custo médio financeiro anterior ao cancelamento |
| 451 | CUSTOFINSEMSTANTCANC | NUMBER(18,6) | S | Custo médio financeiro sem st anterior ao cancelamento |
| 452 | CUSTOCONTANTCANC | NUMBER(18,6) | S | Custo médio contabil anterior ao cancelamento |
| 453 | VLDESPFIN | NUMBER(18,6) | S | Valor da despesa financeira |
| 454 | GUID_JSON | VARCHAR2(37) | S | Identificar único Json para o registro |
| 455 | AGREGARICMSNOVLTRANSF | VARCHAR2(1) | S | Agregar o ICMS ao valor da transferência |
| 456 | VLPAUTABCR | NUMBER(18,6) | S | Valor pauta do ST BCR |
| 457 | VLADICIONALBCSTBCR | NUMBER(18,6) | S | Valor adicional base de cálc. ST BCR |
| 458 | REDBASEIVABCR | NUMBER(18,6) | S | Redução da Base de IVA BCR |
| 459 | REDBASEALIQEXTBCR | NUMBER(18,6) | S | Redução alíq. Externa BCR |
| 460 | PERCIVABCR | NUMBER(10,2) | S | Percentual IVA BCR |
| 461 | PERCICMSFRETEFOBSTBCR | NUMBER(12,4) | S | Percentual ICMS Frete FOB ST BCR |
| 462 | PERCALIQINTBCR | NUMBER(12,4) | S | Percentual alíq. Interna BCR |
| 463 | PERCALIQEXTBCR | NUMBER(12,4) | S | Percentual alíq. Externa BCR |
| 464 | PERICMBCR | NUMBER(12,4) | S | Percentual ICMS BCR |
| 465 | VLBASEIPIECF | NUMBER(18,6) | S | Valor Base IPI |
| 466 | PERCIPIECF | NUMBER(18,6) | S | Percentual IPI |
| 467 | VLIPIECF | NUMBER(18,6) | S | Valor IPI |
| 468 | VLCREDICMSBKP | NUMBER(18,6) | S | Valor do crédito de ICMS |
| 469 | PERCREDICMSBKP | NUMBER(12,4) | S | Percentual do crédito de ICMS |
| 470 | NUMTRANSULTENTSTANT | NUMBER(16) | S | Transação da última entrada com ST Retido |
| 471 | ALIQICMS1RET | NUMBER(12,4) | S | Aliquota de ICMS 1 retido |
| 472 | ESTBONIFIC | VARCHAR2(1) | S | ESTOQUE BONIFICADO |
| 473 | PERCREDBASCALPC | NUMBER(12,4) | S | Perc. de red. base de cálculo PIS/COFINS |
| 474 | DESCRICAONFE | VARCHAR2(120) | S | Descrição do produto na NFE |
| 475 | CODMOTISENCAOANVISA | VARCHAR2(255) | S | Informação de código de isenção para registro da Anvisa |
| 476 | DESCPRODNFE | VARCHAR2(120) | S | Descrição do Produto conforme NFe |
| 477 | USAICMSDESONERACAO | VARCHAR2(1) | S | Indica se o ICMS desonerado será aproveitado na nota |
| 478 | CONSIDERACREDSTFCPBCRCUSTOCONT | VARCHAR2(1) | S | Considera Crédito ST e FCP BCR custo contábil |
| 479 | CONSIDERACREDSTFCPBCRCUSTOGER | VARCHAR2(1) | S | Considera Crédito ST e FCP BCR custo gerencial |
| 480 | CONSIDERACREDICMSBCRCUSTOCONT | VARCHAR2(1) | S | Considera Crédito ICMS BCR custo contábil |
| 481 | CONSIDERACREDICMSBCRCUSTOGER | VARCHAR2(1) | S | Considera Crédito ICMS BCR custo gerencial |
| 482 | VLICMSBCRANT | NUMBER(18,6) | S | Valor de ICMS Substituto anterior. |
| 483 | BASEICMSBCRANT | NUMBER(18,6) | S | Base de ICMS Substituto anterior. |
| 484 | PERMITECREDITOPRESUMIDO | CHAR(1) | S | Identifica os produtos que podem calcular o crédito presumido. |
| 485 | CODFABRICA | VARCHAR2(30) | S | Código de fábrica do produto |
| 486 | UNIDADECOMERCIAL | VARCHAR2(6) | S | Unidade comercial do produto |
| 487 | NUMVERBACAMPANHA | NUMBER(8) | S | Nr. da verba atribuída a campanha. |
| 488 | PERCCUSTFORNEC | NUMBER(12,4) | S | Percentual custeado pelo fornecedor. |
| 489 | VLCUSTFORNEC | NUMBER(18,6) | S | Valor unitário custeado pelo fornecedor. |
| 490 | CODPRECOCESTA | NUMBER(22) | S | Código do preço fixo da cesta básica. |
| 491 | ALIQICMSFECPSTGUIA | NUMBER(12,4) | S | Alíquota de FECP ST Guia |
| 492 | VLFECPSTGUIA | NUMBER(18,6) | S | Valor do FECP ST Guia |
| 493 | VLBASEFCPSTGUIA | NUMBER(18,6) | S | Base do FECP ST Guia |
| 494 | NUMTRANSENTULTENTBNFANT | NUMBER(12) | S | Transação bonificada anterior a entrada bonificada |
| 495 | EXCLUIRICMSBASEPISCOFINS | VARCHAR2(1) | S | Exclusão do icms da base do Pis/Cofins |
| 496 | VLPISEXIGIBSUSPENSA | NUMBER(18,6) | S | Valor PIS exigibilidade suspensa (Valor Inteiro sem deduzir ICMS) |
| 497 | VLCOFINSEXIGIBSUSPENSA | NUMBER(18,6) | S | Valor da COFINS exigibilidade suspensa  (Valor Inteiro sem deduzir ICMS) |
| 498 | VLBASEPISCOFINSEXIGIBSUSP | NUMBER(18,6) | S | Valor Base PIS/COFINS Exigibilidade suspensa (Valor Inteiro sem deduzir ICMS) |
| 499 | NATUREZARECEITA | VARCHAR2(30) | S | Determina o código de natureza da receita |
| 500 | FATORCALCULADOXML | NUMBER(18,9) | S | Fator calculado no XML |
| 501 | GERABONUS | VARCHAR2(1) | S | Indica se irá gerar bônus para produto tintométrico |
| 502 | CODEXCECAOPISCOFINS | NUMBER(6) | S | CODIGO DA EXCEÇÃO DE PIS/COFINS |
| 503 | PRODUSAEXCECAOPISCOFINS | VARCHAR2(2) | S | SE O PRODUTO USOU EXCEÇÃO DE DE PIS/COFINS NA ENTRADA |
| 504 | NITEMXML | NUMBER(38) | S |  |
| 505 | NAOPERSISTEIVAPCMOV | VARCHAR2(1) | S | Indica se sistema irá gravar o campo iva na pocmov se for ST FONTE |
| 506 | USABASEREDICMSPRESUMIDO | VARCHAR2(1) | S | Utilizar Base Reduzida para Crédito Presumido |
| 507 | VLDESCCOMERCIALICMISENCAO | NUMBER(18,6) | S | Valor Desc. Isenção ICM destaque Comercial |
| 508 | PERCBASEREDSTFONTEORIGINAL | NUMBER(8,4) | S | Perc. Base Red. St. Fonte Original |
| 509 | VLFECPTRANSFCD | NUMBER(18,6) | S | Valor do FECP utilizado na transferência do CD Virtual |
| 510 | PERCFECPTRANSFCD | NUMBER(18,6) | S | Percentual do FECP utilizado na transferência do CD Virtual |
| 511 | BASEFECPTRANSFCD | NUMBER(18,6) | S | Base do FECP utilizado na transferência do CD Virtual |
| 512 | VLFUNCEPTRANSFCD | NUMBER(20,8) | S | Valor do FECP ICMS utilizado na transferência do CD Virtual |
| 513 | PERCFUNCEPTRANSFCD | NUMBER(8,4) | S | Percentual do FECP ICMS utilizado na transferência do CD Virtual |
| 514 | BASEFUNCEPTRANSFCD | NUMBER(18,6) | S | Base do FECP ICMS utilizado na transferência do CD Virtual |
| 515 | XML_QCOM | NUMBER(18,6) | S | Quantidade comercial no XML |
| 516 | XML_QTRIB | NUMBER(18,6) | S | Quantidade tributada do XML |
| 517 | RATEAR_FRETE | VARCHAR2(1) | S | Indicar se houve o rateio de frete na inclusão do registro; |
| 518 | VLDESCCMVPROMOCAOMED | NUMBER(18,6) | S | Vl. Verba Fornec. Promocao |
| 519 | NUMLOTEFORNEC | VARCHAR2(20) | S | Numero do lote no fornecedor |
| 520 | BCSTRETANTERIOR | NUMBER(18,6) | S | Base de calculo do ST recolhido anteriormente |
| 521 | VLICMSSUBSTITUTOANTERIOR | NUMBER(18,6) | S | Valor do ICMS substituto recolhido anteriormente |
| 522 | VLICMSSTRETANTERIOR | NUMBER(18,6) | S | Valor do ST recolhido anteriormente |
| 523 | DESTACICMSSTANTERIOR | VARCHAR2(1) | S | Destacar o ICMS ST recolhido anteriormente em tags proprias |
| 524 | USAPMPFBASEST | VARCHAR2(1) | S | Usa Preço Máximo Pessoa Física como Base de ST |
| 525 | USAPFABBASEST | VARCHAR2(1) | S | Usa Preço de Fábrica como Base de ST |
| 526 | PERCREDPMPF | NUMBER(12,4) | S | Alíquota de redução do Preço Máximo Pessoa Física |
| 527 | PERCREDPFAB | NUMBER(12,4) | S | Alíquota de redução do Preço de Fábrica |
| 528 | PMPFMEDICAMENTO | NUMBER(18,6) | S | Preço Máximo Pessoa Física |
| 529 | PRECOFABRICA | NUMBER(18,6) | S | Preço de Fábrica |
| 530 | SOMARVALORDIFBCICMS | VARCHAR2(1) | S | Somar Valor Diferido a BC ICMS |
| 531 | CALCULAPISCOFINSCOMIPI | VARCHAR2(1) | S | Considera IPI para cálculo PIS/COFINS |
| 532 | EXCLUIRDIFALBASEPISCOFINS | VARCHAR2(1) | S | Campo responsável por definir a exclusão do DIFAL da base do pis/cofins conforme tributação |
| 533 | VLFCPDIF | NUMBER(20,8) | S | Valor FCP diferido |
| 534 | VICMSSTDESON | NUMBER(18,6) | S | Valor do icms st desonerado |
| 535 | QBCMONO | NUMBER(22,10) | S | BC do icms próprio |
| 536 | ADREMICMS | NUMBER(12,4) | S | Alíquota ad rem do ICMS |
| 537 | VICMSMONO | NUMBER(22,10) | S | Valor do ICMS próprio |
| 538 | QBCMONORETEN | NUMBER(22,10) | S | BC do ICMS sujeito a retenção |
| 539 | ADREMICMSRETEN | NUMBER(12,4) | S | Alíquota ad rem do ICMS |
| 540 | VICMSMONORETEN | NUMBER(22,10) | S | Valor do ICMS com retenção |
| 541 | PREDADREM | NUMBER(12,4) | S | Percentual de redução da alíquota adremICMS |
| 542 | MOTREDADREM | NUMBER(1) | S | Motivo da redução do adrem |
| 543 | QBCMONODIF | NUMBER(22,10) | S | BC do ICMS diferido |
| 544 | ADREMICMSDIF | NUMBER(12,4) | S | Alíquota ad rem do ICMS diferido |
| 545 | VICMSMONODIF | NUMBER(22,10) | S | Valor do ICMS diferido |
| 546 | QBCMONORET | NUMBER(22,10) | S | BC do ICMS retido anteriormente |
| 547 | ADREMICMSRET | NUMBER(12,4) | S | Alíquota ad rem do imposto retido anteriormente |
| 548 | VICMSMONORET | NUMBER(23,10) | S | Valor do ICMS retido anteriormente |
| 549 | CODBENEFICIOFISCALCOMPLE | VARCHAR2(100) | S | Código de beneficio fiscal complementar |
| 550 | CUSTOULTENTFISCAL | NUMBER(18,6) | S | Custo Ultima Entrada Fiscal |
| 551 | CUSTOFISCAL | NUMBER(18,6) | S | Custo Medio Fiscal |
| 552 | USAMAIORVALORPARACALCULOIPI | VARCHAR2(1) | S | Considera maior entre Preço Líquido e Pauta IPI/Valor IPI |
| 553 | CUSTOFISCALANT | NUMBER(18,6) | S | Custo Fiscal Anterior |
| 554 | CUSTOFISCALANTCANC | NUMBER(18,6) | S | Custo Fiscal Anterior ao Cancelamento |
| 555 | CUSTOFISCALCANC | NUMBER(18,6) | S | Custo Fiscal do Cancelamento |
| 556 | VLIISUSPENSOULTENT_ANT | NUMBER(18,6) | S | Valor II Suspenso de Ultima Entrada Anterior |
| 557 | VLIPISUSPENSOULTENT_ANT | NUMBER(18,6) | S | Valor de IPI Suspenso de Ultima Entrada Anterior |
| 558 | PLIQULTENT | NUMBER(18,6) | S | Preço de compra da última entrada. |
| 559 | PLIQULTENTANT | NUMBER(18,6) | S | Preço de compra da entrada anterior. |
| 560 | SOMACREDCUSTOICMSVLUTENT | VARCHAR2(1) | S | Somar CredCusto ICMS ao vlultent em nota de importação 3010 |
| 561 | INDDEDUZDESONERACAO | VARCHAR2(1) | S | Indica se deduz desoneração do total da nota 0 não 1 sim.  |
| 562 | CCREDPRESUMIDO | VARCHAR2(10) | S | Código de Benefício Fiscal de Crédito Presumido na UF aplicado ao item |
| 563 | REGIMEESPECIALBA | VARCHAR2(1) | S | Regime Especial BA |
| 564 | REGIMEESPECIALPBMED | VARCHAR2(1) | S | Regime Especial PB MED. |
| 565 | QTCONSUMIDA | NUMBER(20,6) | S | Quantidade consumida do produto acabado  |
| 566 | USAPMPFBASESTMG | VARCHAR2(1) | S | Usa PMPF Base ST MG |
| 567 | REPROCESSO | VARCHAR2(1) | S | Reprocessamento de movimentação interna de produtos |
| 568 | PERCICMBKPDIFPR | NUMBER(22,4) | S | Alíquota de ICMS para destaque no XML Diferimento Paraná |
| 569 | ALIQEFETIVAFEEF | NUMBER(12,4) | S | ALIQEFETIVAFEEF |
| 570 | IVAFEEF | NUMBER(12,4) | S | PERCIVAFEEF |
| 571 | ALIQINTERNAFEEF | NUMBER(12,4) | S | ALIQINTERNAFEEF |
| 572 | ALIQEXTERNAFEEF | NUMBER(12,4) | S | ALIQEXTERNAFEEF |
| 573 | PERCCREDALIQEXTERNAFEEF | NUMBER(12,4) | S | PERCCREDALIQEXTERNAFEEF |
| 574 | PERCSTSAIDAFEEF | NUMBER(12,4) | S | PERCSTSAIDAFEEF |
| 575 | PERCFEEF | NUMBER(12,4) | S | PERCVLFEEF |
| 576 | BASECALCSTFEEF | NUMBER(22,6) | S | BASECALCSTFEEF |
| 577 | VLCALCFEEF | NUMBER(22,6) | S | VLCALCFEEF |
| 578 | VLFRETEINTERNACIONAL | NUMBER(18,6) | S | Valor do Frete Internacional |
| 579 | VLSEGUROINTERNACIONAL | NUMBER(18,6) | S | Valor do Seguro Internacional |
| 580 | VLFRETECUSTOINTERNACIONAL | NUMBER(18,6) | S | Valor do Frete Internacional Custo |
| 581 | VLSEGUROCUSTOINTERNACIONAL | NUMBER(18,6) | S | Valor do Seguro Internacional Custo |
| 582 | PRECOBASE | NUMBER(22,10) | S | PREÇO BASE PARA CALCULO CBS-IBS-IS |
| 583 | FATORAUTOMATICO | VARCHAR2(1) | S | Fator automático (rotina 253)  |
| 584 | SITTRIBUTORIG | VARCHAR2(3) | S | Armazena o CST (Código de Situação Tributária) informado no XML do fornecedor, exatamente como recebido. |
| 585 | CUSTOULTENTFISCALANT | NUMBER(18,6) | S | Custo Ultima Entrada Fiscal Anterior |
| 586 | UTILIZAPMPFSTMG | VARCHAR2(1) | S | Usa PMPF ST MG |
| 587 | CBENEFRBC | VARCHAR2(10) | S | Código de Benefício Fiscal RBC |

## Chaves e índices

- Constraint PCMOVCOMPLE_PK (P): NUMTRANSITEM
- Constraint PCMOVCOMPLE_UK1 (U): NUMTRANSNUMEROSERIE, NUMTRANSITEM
- Índice PCMOVCOMPLE_IDX1: NUMTRANSITEM, NUMINVENT
- Índice PCMOVCOMPLE_IDX2: CODPROMOCAOMED, NUMVERBACAMPANHA
- Índice PCMOVCOMPLE_IDX3: NUMTRANSITEM, BONIFIC
- Índice PCMOVCOMPLE_PK UNIQUE: NUMTRANSITEM
- Índice PCMOVCOMPLE_UK1 UNIQUE: NUMTRANSNUMEROSERIE, NUMTRANSITEM
