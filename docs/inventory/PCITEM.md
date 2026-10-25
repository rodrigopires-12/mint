# RIBERDOCES.PCITEM — dicionário

Gerado por `gateway inventory` em 2026-10-25. 359 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODPROD | NUMBER(6) | N | Codigo do produto. |
| 2 | NUMPED | NUMBER(10) | N | Numero do pedido. |
| 3 | PCOMPRA | NUMBER(18,6) | S | Preco de compra. |
| 4 | QTPEDIDA | NUMBER(20,6) | S | Quantidade pedida. |
| 5 | QTENTREGUE | NUMBER(20,6) | S | Quantidade entregue. |
| 6 | DTULTENT | DATE | S | Data da ultima entrada. |
| 7 | PERIPI | NUMBER(12,4) | S | Percentual do IPI. |
| 8 | PLIQUIDO | NUMBER(18,6) | S | Preco liquido. |
| 9 | PERICM | NUMBER(12,4) | S | Percentual do ICMS. |
| 10 | PERCDESC | NUMBER(12,4) | S | Percentual de desconto comercial. |
| 11 | VLIPI | NUMBER(18,6) | S | Valor do IPI. |
| 12 | PERCST | NUMBER(12,4) | S | Percentual ST. |
| 13 | PERCBON | NUMBER(12,4) | S | Percentual de bonificacao. |
| 14 | PERCFRETE | NUMBER(18,6) | S | Percentual de frete CIF. |
| 15 | PERCDESPFIN | NUMBER(12,4) | S | Percentual de despesa financeira(Vendor). |
| 16 | PERCDESCFIN | NUMBER(12,4) | S | Percentual de desconto financeiro. |
| 17 | QTESTDISP | NUMBER(16,3) | S | Quantidade do estoque disponivel. |
| 18 | QTESTIDEAL | NUMBER(16,3) | S | Quantidade do estoque ideal. |
| 19 | QTSUGESTAO | NUMBER(16,3) | S | Quantidade de sugestao de compra. |
| 20 | CUSTOULTENT | NUMBER(18,6) | S | Custo da ultima entrada. |
| 21 | DTULTCOMPRA | DATE | S | Data da ultima compra. |
| 22 | MOEDA | VARCHAR2(1) | S | Moeda (rotina 3001). |
| 23 | DTDOLAR | VARCHAR2(1) | S | Data de cotacao da moeda (rotina 3001). |
| 24 | PERCDESC1 | NUMBER(12,4) | S | Percentual de desconto comercial 1. |
| 25 | PERCDESC2 | NUMBER(12,4) | S | Percentual de desconto comercial 2. |
| 26 | PERCDESC3 | NUMBER(12,4) | S | Percentual de desconto comercial 3. |
| 27 | PERCDESC4 | NUMBER(12,4) | S | Percentual de desconto comercial 4. |
| 28 | PERICMSANTECIPADO | NUMBER(12,4) | S | Percentual do ICMS antecipado. |
| 29 | PERCOUTRASDESP | NUMBER(12,4) | S | Percentual de outras despesas dentro da NF. |
| 30 | VLOUTRASDESP | NUMBER(18,6) | S | Valor de outras despesas dentro da NF. |
| 31 | VLFRETE | NUMBER(18,6) | S | Valor do frete CIF. |
| 32 | VLST | NUMBER(18,6) | S | Valor do ST NF. |
| 33 | PERCFRETEFOB | NUMBER(18,6) | S | Percentual do frete FOB. |
| 34 | PERCDESPADICIONAL | NUMBER(12,4) | S | Percentual ST GUIA. |
| 35 | VLBONIFIC | NUMBER(18,6) | S | Valor da bonificacao dinheiro. |
| 36 | VLFRETECONHEC | NUMBER(18,6) | S | Valor frete FOB. |
| 37 | VLDESPFIN | NUMBER(18,6) | S | Valor da despesa financeira(Vendor) |
| 38 | VLDESPADICIONAL | NUMBER(18,6) | S | Valor do ST GUIA. |
| 39 | VLICMSANTECIPADO | NUMBER(18,6) | S | Valor do ICMS antecipado. |
| 40 | PERCVLBONIFIC | NUMBER(12,4) | S | Percentua bonificação em dinheiro. |
| 41 | VLPERCBON | NUMBER(18,6) | S | Valor bonificação em mercadoria. |
| 42 | PCOMPRAANT | NUMBER(18,6) | S | Preco de compra anterior. |
| 43 | PERCINTEIRO | NUMBER(6,3) | S | Classificação arroz rotina 1330. |
| 44 | PERCQUEBRADO | NUMBER(6,3) | S | Classificação arroz rotina 1330. |
| 45 | PERCIMPUREZA | NUMBER(6,3) | S | Classificação arroz rotina 1330. |
| 46 | PERCVERMELHO | NUMBER(6,3) | S | Classificação arroz rotina 1330. |
| 47 | PERCUMIDADE | NUMBER(6,3) | S | Classificação arroz rotina 1330. |
| 48 | PERCICMRED | NUMBER(12,4) | S | Percentual de reducao do ICMS ST. |
| 49 | PERCIVA | NUMBER(12,4) | S | Percentual do IVA ST. |
| 50 | PERCALIQINT | NUMBER(12,4) | S | Percentual da aliquota interna ST. |
| 51 | PERCALIQEXT | NUMBER(12,4) | S | Percentual da aliquota externa ST. |
| 52 | PERPIS | NUMBER(12,4) | S | Percentual de credito do PIS. |
| 53 | VLCREDPIS | NUMBER(18,6) | S | Valor do credito do PIS. |
| 54 | VLCREDICMS | NUMBER(18,6) | S | Valor do credito do ICMS. |
| 55 | PERICMFRETE | NUMBER(8,4) | S | Percentual ICMS frete CIF. |
| 56 | VLCREDICMSFRETE | NUMBER(18,6) | S | Valor do ICMS frete CIF. |
| 57 | PERCBONNF | NUMBER(12,4) | S | Não utilizado. |
| 58 | VLPERCBONNF | NUMBER(18,6) | S | Não utilizado. |
| 59 | PCOMPRAMED | NUMBER(18,6) | S | Não utilizado. |
| 60 | QTPEDIDAMED | NUMBER(16,3) | S | Não utilizado. |
| 61 | PLIQUIDOMED | NUMBER(18,6) | S | Não utilizado. |
| 62 | TIPOCALCST | VARCHAR2(1) | S | Tipo de ST: NF, Guia, Ambos, BCR. |
| 63 | VLSTBONIFIC | NUMBER(18,6) | S | Não utilizado. |
| 64 | VLPAUTA | NUMBER(18,6) | S | Valor da pauta ST. |
| 65 | PERCSUFRAMA | NUMBER(12,4) | S | Percentual do SUFRAMA. |
| 66 | VLSUFRAMA | NUMBER(18,6) | S | Valor do SUFRAMA. |
| 67 | PERCDESC5 | NUMBER(12,4) | S | Percentual de desconto comercial 5. |
| 68 | PERCCREDICMPRESUMIDO | NUMBER(8,4) | S | Percentual crédito ICMS presumido. |
| 69 | PERCDESC6 | NUMBER(12,4) | S | Percentual de desconto comercial 6. |
| 70 | PERCDESC7 | NUMBER(12,4) | S | Percentual de desconto comercial 7. |
| 71 | PERCDESC8 | NUMBER(12,4) | S | Percentual de desconto comercial 8. |
| 72 | PERCDESC9 | NUMBER(12,4) | S | Percentual de desconto comercial 9. |
| 73 | PERCDESC10 | NUMBER(12,4) | S | Percentual de desconto comercial 10. |
| 74 | VLCREDPRESUMIDO | NUMBER(18,6) | S | Valor do credito presumido. |
| 75 | PERCIVA2 | NUMBER(12,4) | S | Alíquota de ST Outorgado |
| 76 | PERCALIQINT2 | NUMBER(12,4) | S | Não utilizado. |
| 77 | VLPAUTA2 | NUMBER(12,4) | S | Não utilizado. |
| 78 | PERCALIQEXT2 | NUMBER(12,4) | S | Não utilizado. |
| 79 | VLFRETEPORKG | NUMBER(18,6) | S | Valor do frete por quilo. |
| 80 | NUMSEQ | NUMBER(6) | N | Número de sequencia digitação do mesmo item. |
| 81 | PERCMEXIVA | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 82 | PERCMEXIEPS | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 83 | PERCMEXFPGC | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 84 | VLMEXIVA | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 85 | VLMEXIEPS | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 86 | VLMEXFPGC | NUMBER(10,6) | S | Campo usado no Winthor Mexico. |
| 87 | PRODBONIFICADO | VARCHAR2(1) | S | Identificador produto bonificado. |
| 88 | VLADICIONALBCST | NUMBER(18,6) | S | Valor adicional base de calculo ST. |
| 89 | VLIPIPORKG | NUMBER(18,6) | S | Valor do IPI por quilo. |
| 90 | DATAENTREGA | DATE | S | Previsão da Data de Entrega do item do Pedido de Compra.  |
| 91 | PRODUTOGRADE | VARCHAR2(1) | S | Produto de grade. |
| 92 | PTABELA | NUMBER(18,6) | S | Campo para indicar se o produto sera aproveitado credito de icms na entrada |
| 93 | PERCBASEREDST | NUMBER(18,6) | S | Será utilizado para o cálculo de redução de base de ST (rotina 209). |
| 94 | APLICPERCIVAPAUTA | VARCHAR2(1) | S | Determina se o valor de IVA será agregado no valor da paula (rotina 209). |
| 95 | REDBASEIVA | NUMBER(18,6) | S | Será utilizado para Redução da Base de Cálculo de Substituição Tributária na Entrada. |
| 96 | REDBASEALIQEXT | NUMBER(18,6) | S | Será utilizado para Redução da Base de Cálculo da Aliquota Externa ref a substituição Tributária na Entrada. |
| 97 | PERCBONOUTRAS | NUMBER(18,6) | S | Percentual de outras bonificacoes. |
| 98 | VLBONIFICOUTRAS | NUMBER(18,6) | S | Valor de outras bonificacoes. |
| 99 | PERCOFINS | NUMBER(12,4) | S | Indica o valor do percentual de cofins. |
| 100 | APROVEITACREDICMS | VARCHAR2(1) | S | Aproveita credito ICMS na impotação rotina 30023. |
| 101 | BASEICMS | NUMBER(18,6) | S | Base de icms item a item do ped. compra.  |
| 102 | BASEICST | NUMBER(18,6) | S | Base de subst. tributaria item a item do ped. compra.  |
| 103 | VLICMS | NUMBER(18,6) | S | Valor de icms item a item do ped. compra.  |
| 104 | PERCIMPORTACAO | NUMBER(12,4) | S | Percentual de Importação |
| 105 | APROVEITACREDPISCOFINS | VARCHAR2(1) | S | Aproveita credito de PIS/COFINS na importação rotina 3002. |
| 106 | VLPAUTAICMSANTEC | NUMBER(18,6) | S | Valor pauta ICMS antecipado. |
| 107 | PERCDIREITOSADUANEIROS | NUMBER(8,4) | S | Direitos aduaneiros. |
| 108 | PERCIMPOSTOCONSUMO | NUMBER(8,4) | S | Imposto de consumo. |
| 109 | PERCIMPOSTOSELO | NUMBER(8,4) | S | Imposto selo. |
| 110 | VLCREDCOFINS | NUMBER(18,6) | S | Indica valor crédito COFINS. |
| 111 | PERCREDICMS | NUMBER(12,4) | S | Indica o percentual de crédito ICMS. |
| 112 | CALCCREDIPI | VARCHAR2(1) | S | Calcular Custo com Crédito de IPI (Compra). |
| 113 | PERCIMPPRODUTORURAL | NUMBER(12,4) | S | Indica o percentual do imposto do produtor rural. |
| 114 | ROTINALANC | NUMBER(6) | S | Indica o número da rotina que alterou o registro. |
| 115 | PERCICMSFRETEFOBST | NUMBER(12,4) | S | Percentual de crédito de ICMS sobre o valor Frete FOB, utilizado no calculo do ST |
| 116 | QTENTREGUEMED | NUMBER(16,3) | S | Não utilizado. |
| 117 | PERCICMSDIFERIDO | NUMBER(8,4) | S | Indica o percentual do ICMS diferido. |
| 118 | VLPAUTAICMS | NUMBER(18,6) | S | Indica o valor de pauta ICMS. |
| 119 | VLSTBCR | NUMBER(18,6) | S | Indica o valor ST BCR. |
| 120 | VLBASESTBCR | NUMBER(18,6) | S | Indica o valor Base ST BCR. |
| 121 | PERCSTBCR | NUMBER(12,4) | S | Indica o percentual ST BCR. |
| 122 | VLBASESTGUIA | NUMBER(18,6) | S | Indica o valor base ST guia. |
| 123 | PERCDESPFORANF | NUMBER(12,4) | S | Indica o percentual de despesa fora da nota. |
| 124 | VLDESPFORANF | NUMBER(18,6) | S | Indica o valor da despesa fora da nota. |
| 125 | PERCDIFALIQUOTAS | NUMBER(8,4) | S | Indica o percentual de diferença de alíquota. |
| 126 | VLDIFALIQUOTAS | NUMBER(18,6) | S | Indica o valor da diferença de alíquota. |
| 127 | BASEDIFALIQUOTAS | NUMBER(18,6) | S | Indica o valor base da diferença de alíquota. |
| 128 | PERCSEGURO | NUMBER(18,6) | S | Indica o percentual de seguro. |
| 129 | VLSEGURO | NUMBER(18,6) | S | Indica o valor do seguro. |
| 130 | PERCDESPDENTRONF | NUMBER(18,6) | S | Indica o percentual outras despesa na NF. |
| 131 | VLDESPDENTRONF | NUMBER(18,6) | S | Indica o Valor outras despesas na NF. |
| 132 | QTBONIFMERC | NUMBER(16,3) | S | Indica a quantidade de bonificação em mercadoria. |
| 133 | PERCALIQEXTGUIA | NUMBER(12,4) | S | Percentual da alíquota externa para calculo do ST Guia. |
| 134 | NUMADICAO | NUMBER(3) | S | Número de adição |
| 135 | NUMSEQADICAO | NUMBER(3) | S | Número sequencial de controle de adição |
| 136 | PERCDESCICMSDIF | NUMBER(8,4) | S | Percentual do desconto de ICMS Difererido. |
| 137 | PERCIISUSPENSO | NUMBER(8,4) | S | Percenctual de imposto de importação suspenso para cáculo do PIS, COFINS, ICMS e IPI. |
| 138 | PERPISCALCDI | NUMBER(8,4) | S | Percentual de PIS calculo da D.I. |
| 139 | PERCOFINSCALCDI | NUMBER(8,4) | S | Percentual de COFINS calculo da D.I. |
| 140 | UTILIZACREDREDPISCOFINS | VARCHAR2(1) | S | Alíquota crédito reduzido PIS/COFINS |
| 141 | VLPAUTAIPI | NUMBER(18,6) | S | Informar valor da pauta do IPI na tabela PCMOVCOMPLE. |
| 142 | QTBONIFICMERCOUTRAS | NUMBER(20,6) | S | Quantidade de mercadoria bonificação outras |
| 143 | CUSTOULTPEDCOMPRA | NUMBER(18,6) | S | Custo da última entrada pedido de compra. |
| 144 | IPIPORVALOR | VARCHAR2(1) | S | IPI fixo por valor. |
| 145 | PERCMVAORIG | NUMBER(12,4) | S | Percentual de MVA Original para cálculo do MVA ajustado |
| 146 | APLICREDBASEIVAPLIQ | VARCHAR2(1) | S | Aplicar redução base IVA preço liquido |
| 147 | APLICREDBASEIVAPLIQBCR | VARCHAR2(1) | S | Aplicar redução base IVA preço liquido BCR |
| 148 | PISCOFINSRETIDO | VARCHAR2(1) | S | Considerar a informação de PIS/COFINS retido na figura tributária. |
| 149 | NUMLICENCAIMP | NUMBER(20) | S | Número da licença de importação. |
| 150 | DTEMISSAOLICENCAIMP | DATE | S | Data de emissão da licença de importação |
| 151 | DTVALIDADELICENCAIMP | DATE | S | Data de validade da licença de importação |
| 152 | CODFUNCAUTORIZA | NUMBER(8) | S | Funcionário que autorizou incluir item no pedido |
| 153 | DTAUTORIZACAO | DATE | S | Data da autorização para incluir item no pedido |
| 154 | DTLEMISSAOLICENCAIMPORT | DATE | S | Data de emissão da licença de importação |
| 155 | NUMLICENCAIMPORT | VARCHAR2(20) | S | Número da licença de importação |
| 156 | DTVENCLICENCAIMPORT | DATE | S | Data de vencimento da licença de importação |
| 157 | VLPAUTAPISCOFINS | NUMBER(18,6) | S | Valor pauta PIS/COFINS |
| 158 | USAPISCOFINSLIT | VARCHAR2(1) | S | Usa PIS/COFINS por litragem |
| 159 | BASEPISCOFINSLIT | NUMBER(18,6) | S | Base PIS/COFINS por litragem |
| 160 | CODSITTRIBPISCOFINS | NUMBER(3) | S | Indica o Código da Situação Tributária Referente ao PIS/COFINS |
| 161 | PERCALTERCUSTOENT | NUMBER(6,2) | S | Percentual de acréscimo no custo do produto por não reembolso da avaria. |
| 162 | VLADUANEIRA | NUMBER(18,6) | S | Valor referente a despesas aduaneiras |
| 163 | VLSISCOMEX | NUMBER(18,6) | S | Valor referente a despesas com SISCOMEX |
| 164 | VLOUTRASDESPIMP | NUMBER(18,6) | S | Valor referente a outras despesas de importação |
| 165 | USAPMCBASEST | VARCHAR2(1) | S | Usa PMC base ST |
| 166 | PERCREDPMC | NUMBER(12,4) | S | Perc.redução PMC |
| 167 | PERCOUTROSCUSTOS | NUMBER(18,6) | S | Percentual de outros custos do item de importação. |
| 168 | VLOUTROSCUSTOS | NUMBER(18,6) | S | Valor de outros custos do item de importação. |
| 169 | PERCCAPATAZIA | NUMBER(18,6) | S | Percentual da capatazia rateado entre os itens de importação. |
| 170 | VLCAPATAZIA | NUMBER(18,6) | S | Valor da capatazia rateado entre os itens de importação. |
| 171 | PERCSISCOMEXCUSTO | NUMBER(18,6) | S | Percentual Siscomex Custo |
| 172 | VLSISCOMEXCUSTO | NUMBER(18,6) | S | Valor Siscomex Custo |
| 173 | PERCADUANEIRACUSTO | NUMBER(18,6) | S | Percentual Despesa Aduaneira Custo |
| 174 | VLADUANEIRACUSTO | NUMBER(18,6) | S | Valor Despesa Aduaneira Custo |
| 175 | PERCOUTRASDESPIMPCUSTO | NUMBER(18,6) | S | Percentual Outras Despesas Importação Custo |
| 176 | VLOUTRASDESPIMPCUSTO | NUMBER(18,6) | S | Valor Outras Despesas Importação Custo |
| 177 | PERCOUTROSCUSTOSCUSTO | NUMBER(18,6) | S | Percentual Outros Custos Importação Custo |
| 178 | VLOUTROSCUSTOSCUSTO | NUMBER(18,6) | S | Valor Outros Custos Importação Custo |
| 179 | NUMDRAWBACK | VARCHAR2(20) | S | Número do ato concessório do regime DRAWBACK. |
| 180 | CODFUNCALTER | NUMBER(8) | S | Código do funcionário que fez a alteração no pedido de compra, cadastro pela rotina 528 |
| 181 | PERCFRETENFCUSTO | NUMBER(18,6) | S | Percentual de frete NF custo |
| 182 | VLFRETENFCUSTO | NUMBER(18,6) | S | Valor do frete NF custo |
| 183 | PERCSEGUROCUSTO | NUMBER(18,6) | S | Percentual de seguro NF custo |
| 184 | VLSEGUROCUSTO | NUMBER(18,6) | S | Valor do seguro NF custo |
| 185 | PERCCAPATAZIACUSTO | NUMBER(18,6) | S | Percentual de capatazia NF custo |
| 186 | VLCAPATAZIACUSTO | NUMBER(18,6) | S | Valor da capatazia NF custo |
| 187 | PERCDESPDENTRONFCUSTO | NUMBER(18,6) | S | INC	Percentual de acréscimo NF custo |
| 188 | VLDESPDENTRONFCUSTO | NUMBER(18,6) | S | INC	Valor do acréscimo NF custo |
| 189 | VLIMPORTACAO | NUMBER(18,6) | S | Percentual imposto importação |
| 190 | PERCSISCOMEX | NUMBER(18,6) | S | Percentual despesa SISCOMEX |
| 191 | PERCADUANEIRA | NUMBER(18,6) | S | Percentual despesa aduaneira |
| 192 | PERCOUTRASDESPIMP | NUMBER(18,6) | S | Percentual outras depesas importação |
| 193 | PERCCARGATRIBMEDIA | NUMBER(18,6) | S | Percentual de carga tributária média, utilizado no SEFAZ MT para calculo da substituição tributária |
| 194 | BASEICMSANTECIPADO | NUMBER(18,6) | S | Base do ICMS Antecipado |
| 195 | BASECREDPRESUMIDO | NUMBER(18,6) | S | Base do Crédito Presumido |
| 196 | STATUS | VARCHAR2(2) | S | Status da movimentação do item: Gerencial ou Fiscal |
| 197 | CALCCREDIPICONT | VARCHAR2(1) | S | Calcula crédito IPI custo contabil na importação |
| 198 | APROVEITACREDICMSCONT | VARCHAR2(1) | S | Aproveita crédito ICMS custo contabil na importação |
| 199 | APROVEITACREDPISCOFINSCONT | VARCHAR2(1) | S | Aproveita crédito PIS/COFINS custo contabil na importação |
| 200 | PERCICMSBASEICMSANTECIPADO | NUMBER(12,4) | S | Aliq. Para base do ICMS antecipado |
| 201 | GERABASEPISCOFINSSEMALIQ | VARCHAR2(1) | S | Defina se deve gerar base de PIS/CONFINS mesmo quando não for informado aliquotas de PIS/CONFINS |
| 202 | CODSITTRIBST | VARCHAR2(2) | S | Gravar o código da situação tributária st |
| 203 | OBS | VARCHAR2(500) | S | Observação do Item |
| 204 | QTMINSUGCOMPRA | NUMBER(22,8) | S | Quantidade mínima de compra calculo da sugestão de compra |
| 205 | MULTIPLOCOMPRAS | NUMBER(18,6) | S | Quantidade multipla de compra calculo da sugestão de compra |
| 206 | OBSITEM | VARCHAR2(500) | S | Quantidade multipla de compra calculo da sugestão de compra |
| 207 | VLBASEIPISUSPENSO | NUMBER(18,6) | S | Base do IPI suspenso |
| 208 | VLIPISUSPENSO | NUMBER(18,6) | S | Valor do IPI suspenso |
| 209 | PERCIPISUSPENSO | NUMBER(12,4) | S | Percentual do IPI suspenso |
| 210 | CONSIISUSPENSOBASEICMS | VARCHAR2(1) | S | Considera II Suspenso base ICMS. |
| 211 | CONSIPISUSPENSOBASEICMS | VARCHAR2(1) | S | Considera IPI Suspenso base ICMS. |
| 212 | VLIISUSPENSO | NUMBER(18,6) | S | Valor do II Suspenso. |
| 213 | VLALTERCUSTOENT | NUMBER(18,6) | S | Acréscmo no Custo da Última Entrada. |
| 214 | CONSIDERASTNFCUSTO | VARCHAR2(1) | S | Considera ST NF  no Custo. |
| 215 | CONSIDERASTGUIACUSTO | VARCHAR2(1) | S | Considera ST Guia no Custo. |
| 216 | CONSIDERASTNFCUSTOCONT | VARCHAR2(1) | S | Considera ST NF no Custo Contábil. |
| 217 | CONSIDERASTGUIACUSTOCONT | VARCHAR2(1) | S | Considera ST Guia no Custo Contábil. |
| 218 | UTILIZABASEPISCOFINSSUSP | VARCHAR2(1) | S | Utiliza Base PIS/COFINS Suspenso. |
| 219 | NUMFCI | VARCHAR2(36) | S | Ficha de Conteúdo de Importação. |
| 220 | VLPARCELAIMPFCI | NUMBER(18,6) | S | Valor da Parcela Importada. |
| 221 | PERCCONTEUDOIMPFCI | NUMBER(5,2) | S | Conteúdo de Importação. |
| 222 | VLIMPORTACAOFCI | NUMBER(18,6) | S | Valor de Importação. |
| 223 | CONTROLARFCI | VARCHAR2(1) | S | Controla Ficha de Conteúdo de Importação. |
| 224 | QTDIGITADA | NUMBER(18,6) | S | Quantidade digitada pelo usúario. |
| 225 | LOTELICIT | VARCHAR2(10) | S | Lote da licitação |
| 226 | NUMEROITEMLICIT | NUMBER(9) | S | Número item licitação |
| 227 | CUSTOULTPEDCOMPRASEMST | NUMBER(18,6) | S | Custo Ult.Pedido de Compra sem ST. |
| 228 | PERCAFRMM | NUMBER(18,6) | S | Percentual de Adicional de Frete para a Renovação da Marinha Mercante |
| 229 | VLAFRMM | NUMBER(18,6) | S | Valor do Adicional de Frete para a Renovação da Marinha Mercante |
| 230 | PERCAFRMMCUSTO | NUMBER(18,6) | S | Percentual de Adicional de Frete para a Renovação da Marinha Mercante |
| 231 | VLAFRMMCUSTO | NUMBER(18,6) | S | Valor do Adicional de Frete para a Renovação da Marinha Mercante |
| 232 | PERCIVAICMANTECIP | NUMBER(12,4) | S | Percentual IVA ICMS Antecipado |
| 233 | PERCALIQINTICMANTECIP | NUMBER(12,4) | S | Percentual Aliquota Interna ICMS Antecipado |
| 234 | PERCALIQEXTICMANTECIP | NUMBER(12,4) | S | Percentual Aliquota Externa  ICMS Antecipado |
| 235 | VLPISCALCDI | NUMBER(18,6) | S | Valor do PIS suspenso |
| 236 | VLCOFINSCALCDI | NUMBER(18,6) | S | Valor do COFINS suspenso |
| 237 | MOTIVO | VARCHAR2(60) | S | Motivo cancelamento |
| 238 | CODPRODORIGEM | NUMBER(6) | S | Código produto origem na Importação. |
| 239 | PERCBONIFIC | NUMBER(18,6) | S | % Bonificação Dinheiro |
| 240 | VERBADINPORPERC | VARCHAR2(1) | S | Calcular verba dinheiro por percentual |
| 241 | QTFATORCONVERSAO | NUMBER(6,2) | S | Permite alterar fator de conversão da importação xml. |
| 242 | VLADICIONALBCICMSANTECIP | NUMBER(18,6) | S | Valor adicional base Icms Antecipado |
| 243 | APLICPERCIVAPAUTAICMSANTECIP | VARCHAR2(1) | S | Aplica IVA sobre valor de pauta ICMS antecipado |
| 244 | PERCICMSFRETEFOBICMSANTECIP | NUMBER(12,4) | S | % ICMS Frete FOB p/ calc.Icms Antecipado |
| 245 | PERCMVAORIGICMSANTECIP | NUMBER(12,4) | S | % IVA Original ICMS Antecipado |
| 246 | PERCCARGATRIBMEDIAICMSANTECIP | NUMBER(18,6) | S | % Carga Média tributária Icms Antecipado |
| 247 | REDBASEIVAICMSANTECIP | NUMBER(18,6) | S | % Redução IVA Icms Antecipado |
| 248 | REDBASEALIQEXTICMSANTECIP | NUMBER(18,6) | S | % Redução Aliquota externa ICMS Antecipado. |
| 249 | PERCANTIDUMPING | NUMBER(18,6) | S | Percentual Antidumping DI |
| 250 | VLANTIDUMPING | NUMBER(18,6) | S | Valor Antidumping DI |
| 251 | PERCANTIDUMPINGCUSTO | NUMBER(18,6) | S | Percentual Antidumping Custo |
| 252 | VLANTIDUMPINGCUSTO | NUMBER(18,6) | S | Valor Antidumping Custo |
| 253 | TIPOEMBALAGEMPEDIDO | VARCHAR2(1) | S | Representa a unidade em que o item foi digitado (Venda ou Master) |
| 254 | MULTIMPLOUNIDADEMASTER | VARCHAR2(1) | S | Representa se a quantidade digitada é ou não multiplo da unidade master |
| 255 | PESOLIQDI | NUMBER(12,6) | S | Peso Líquido calculo rateiro Importação |
| 256 | VLDESCICMSDIF | NUMBER(18,6) | S | Valor de ICMS diferido |
| 257 | OBSBLOQUEIO | VARCHAR2(500) | S | Observação com o motivo do bloqueio |
| 258 | CODFAB | VARCHAR2(30) | S |  |
| 259 | NUMSEQCONTROLE | NUMBER(6) | S | Número sequencial utilizado no vinculo do controle de embarque e pedido master |
| 260 | NUMPEDPRINC | NUMBER(10) | S | Número do pedido de compra principal. |
| 261 | NUMPEDMASTERORIGEM | NUMBER(10) | S | Número do pedido master origem. |
| 262 | NUMSEQORIGEM | NUMBER(6) | S | Número sequêncial do item de origem. |
| 263 | VLSTOUTORGADO | NUMBER(18,6) | S | Valor de ST Outorgado |
| 264 | PERCICMSDESONERACAO | NUMBER(12,4) | S | Percentual de ICMS Desoneração |
| 265 | VLICMSDESONERACAO | NUMBER(18,6) | S | Valor de ICMS Desoneração |
| 266 | VLBASEFCPICMS | NUMBER(18,6) | S | Valor da base de calculo do Fundo de Combate a Pobreza |
| 267 | PERACRESCIMOFUNCEP | NUMBER(12,4) | S | Percentual do Fundo de Combate a Pobreza ICMS |
| 268 | VLACRESCIMOFUNCEP | NUMBER(18,6) | S | Valor do Fundo de Combate a Pobreza ICMS |
| 269 | VLBASEFCPST | NUMBER(18,6) | S | Valor de base de calculo do Fundo de Combate a Pobreza ST |
| 270 | ALIQICMSFECP | NUMBER(12,4) | S | Percentual do Fundo de Combate a Pobreza ST |
| 271 | VLFECP | NUMBER(18,6) | S | Valor do Fundo de Combate a Pobreza ST |
| 272 | USAMAIORVALORPARACALCULOICMS | VARCHAR2(1) | S | Usa maior valor para calculo do ICMS (Pauta de ICMS e preço de compra) |
| 273 | USABASESTNOFCP | VARCHAR2(1) | S | Base FCP ST - Considera base do ST |
| 274 | CODMOTIVOICMSDESONERADO | VARCHAR2(1) | S | Motivo desoneração suframa |
| 275 | QTRESIDUO | NUMBER(20,6) | S | Quantidade de residuo do item |
| 276 | USAICMSDESONERACAO | VARCHAR2(1) | S | Indica se o ICMS desonerado será aproveitado na nota |
| 277 | ALIQICMSFECPSTGUIA | NUMBER(12,4) | S | Alíquota de FCP ST Guia |
| 278 | VLFECPSTGUIA | NUMBER(18,6) | S | Valor de FCP ST Guia |
| 279 | VLBASEFCPSTGUIA | NUMBER(18,6) | S | Base de FCP ST Guia |
| 280 | BASEICMSBCR | NUMBER(18,6) | S | Base do ICMS substituto |
| 281 | VLICMSBCR | NUMBER(18,6) | S | Valor do ICMS substituto |
| 282 | VLBCFCPSTRET | NUMBER(18,6) | S | Base do FCP ST retido |
| 283 | VLFCPSTRET | NUMBER(18,6) | S | Valor do FCP ST retido |
| 284 | PERFCPSTRET | NUMBER(12,4) | S | Percentual do FCP ST retido |
| 285 | USABASEREDICMSPRESUMIDO | VARCHAR2(1) | S | Utilizar Base Reduzida para Crédito Presumido |
| 286 | CODPRAZOENT | VARCHAR2(4) | S | Código do prazo de entrega |
| 287 | QTTRANSITOTV10 | NUMBER(22,8) | S | Quantidade em transito com o tipo de venda igual a 10 - Transferência |
| 288 | QTBLOQUEADA | NUMBER(20,6) | S | Quantidade bloqueada do item |
| 289 | QTINDENIZ | NUMBER(20,6) | S | Quantidade indenizavel do item |
| 290 | USAPMPFBASEST | VARCHAR2(1) | S | Usa Preço Máximo Pessoa Física como Base de ST |
| 291 | USAPFABBASEST | VARCHAR2(1) | S | Usa Preço de Fábrica como Base de ST |
| 292 | PERCREDPMPF | NUMBER(12,4) | S | Alíquota de redução do Preço Máximo Pessoa Física |
| 293 | PERCREDPFAB | NUMBER(12,4) | S | Alíquota de redução do Preço de Fábrica |
| 294 | CALCULAPISCOFINSCOMIPI | VARCHAR2(1) | S | Considera IPI para cálculo PIS/COFINS |
| 295 | PMERC | NUMBER(18,6) | S | PREÇO DE MERCADORIA |
| 296 | USAMAIORVALORPARACALCULOIPI | VARCHAR2(1) | S | Considera maior entre Preço Líquido e Pauta IPI/Valor IPI |
| 297 | REGIMEESPECIALBA | VARCHAR2(1) | S | Regime Especial BA |
| 298 | REGIMEESPECIALPBMED | VARCHAR2(1) | S | REGIME ESPECIAL PB MED. |
| 299 | USAPMPFBASESTMG | VARCHAR2(1) | S | Usa PMPF Base ST MG |
| 300 | PLIQULTENT | NUMBER(18,6) | S | Preço de compra da última entrada. |
| 301 | CODFIGURA | NUMBER(8) | S | Código figura tributária rotina 212. |
| 302 | ALIQEFETIVAFEEF | NUMBER(12,4) | S | ALIQEFETIVAFEEF |
| 303 | IVAFEEF | NUMBER(12,4) | S | PERC IVA FEEF |
| 304 | ALIQINTERNAFEEF | NUMBER(12,4) | S | ALIQINTERNAFEEF |
| 305 | ALIQEXTERNAFEEF | NUMBER(12,4) | S | ALIQEXTERNAFEEF |
| 306 | PERCCREDALIQEXTERNAFEEF | NUMBER(12,4) | S | PERCCREDALIQEXTERNAFEEF |
| 307 | PERCSTSAIDAFEEF | NUMBER(12,4) | S | PERCSTSAIDAFEEF |
| 308 | PERCFEEF | NUMBER(12,4) | S | PERCENTUAL FEEF |
| 309 | BASECALCSTFEEF | NUMBER(22,6) | S | BASECALCSTFEEF |
| 310 | VLCALCFEEF | NUMBER(22,6) | S | VLCALCFEEF |
| 311 | CODROTA | NUMBER(10) | S | Código da rota(PCROTACOMPRAC) vinculado ao produto |
| 312 | CODFILIALPEDIDOCOMPRA | VARCHAR2(2) | S | Código da filial do primeiro pedido de compra (ORDEM 1) relacionado a (PCROTACOMPRAI) |
| 313 | CODFILIALDESTINO | VARCHAR2(2) | S | Código da filial de destino da compra |
| 314 | ORDEM | NUMBER(10) | S | Ordem sequencial relacionada ao CODROTA |
| 315 | UTILIZAPMPFSTMG | VARCHAR2(1) | S | Usa PMPF ST MG |
| 316 | PREDALIQCBS | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 317 | SOMATOTALNF_CBSIBS | VARCHAR2(1) | S | Define se o valor do CBSIBS irá somar no valor total da nota S(Sim) ou N(Não) |
| 318 | SOMATOTALNF_IS | VARCHAR2(1) | S | Define se irá somar o valor do IS no valor total da nota S(Sim) ou N(Não) |
| 319 | VLIS | NUMBER(23,10) | S | Valor do Imposto Seletivo |
| 320 | ALIQCBS | NUMBER(10,4) | S | Alíquota CBS |
| 321 | ALIQEFETREGIBSMUN | NUMBER(7,4) | S | Valor da alíquota do IBS do Município |
| 322 | ALIQIS | NUMBER(10,4) | S | Alíquota do Imposto Seletivo |
| 323 | CCLASSTRIBIBSCBS | VARCHAR2(6) | S | CClassTrib IBS e CBS |
| 324 | CCLASSTRIBIS | VARCHAR2(6) | S | Código de Classificação Tributária do Imposto Seletivo |
| 325 | CCLASSTRIBREG | VARCHAR2(6) | S | Código de Classificação Tributária do IBS e CBS |
| 326 | CSTIBSCBS | VARCHAR2(3) | S | CST de IBS e CBS |
| 327 | CSTIS | VARCHAR2(3) | S | Código de Situação Tributária do Imposto Seletivo |
| 328 | CSTTRIBREG | VARCHAR2(3) | S | Código de Situação Tributária do IBS e CBS |
| 329 | IBSUF | NUMBER(7,4) | S | Alíquota do IBS de competência das UF |
| 330 | PALIQEFETCBS | NUMBER(7,4) | S | Alíquota Efetiva da CBS que será aplicada a Base de Cálculo |
| 331 | PALIQEFETIBSMUN | NUMBER(7,4) | S | Alíquota Efetiva do IBS de competência do Município que será aplicada a Base de Cálculo |
| 332 | PALIQEFETIBSUF | NUMBER(7,4) | S | Alíquota Efetiva do IBS de competência das UF que será aplicada a Base de Cálculo |
| 333 | PALIQEFETREGCBS | NUMBER(7,4) | S | Valor da alíquota da CBS |
| 334 | PALIQEFETREGIBSUF | NUMBER(7,4) | S | Valor da alíquota do IBS da UF |
| 335 | PCBGCOMPRAGOV | NUMBER(7,4) | S | Alíquota da CBS para compra governamental |
| 336 | PIBSMUN | NUMBER(7,4) | S | Alíquota do IBS de competência dos municípios |
| 337 | PIBSMUNCOMPRAGOV | NUMBER(7,4) | S | Alíquota do IBS de competência do Município para compra governamental |
| 338 | PIBSUFCOMPRAGOV | NUMBER(7,4) | S | Alíquota do IBS de competência do Estado para compra governamental |
| 339 | PREDALIQIBSMUN | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 340 | PREDALIQIBSUF | NUMBER(7,4) | S | Percentual da redução de alíquota |
| 341 | VCBSCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo da CBS calculado para compra governamental |
| 342 | VIBSMUN | NUMBER(23,10) | S | Valor do IBS de competência do Município |
| 343 | VIBSMUNCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo do IBS do Município calculado para compra governamental |
| 344 | VIBSUF | NUMBER(23,10) | S | Valor do IBS de competência da UF |
| 345 | VIBSUFCOMPRAGOV | NUMBER(23,10) | S | Valor do Tributo do IBS da UF calculado para compra governamental |
| 346 | VLBASEIBSCBS | NUMBER(23,10) | S | Base de cálculo para IBS CBS |
| 347 | VLBASEIS | NUMBER(23,10) | S | Valor da Base de Cálculo do Imposto Seletivo |
| 348 | VLCBS | NUMBER(23,10) | S | Valor da CBS |
| 349 | VTRIBREGCBS | NUMBER(23,10) | S | Valor do Tributo da CBS |
| 350 | VTRIBREGIBSMUN | NUMBER(23,10) | S | Valor do Tributo do IBS do Município |
| 351 | VTRIBREGIBSUF | NUMBER(23,10) | S | Valor do Tributo do IBS da UF |
| 352 | VLIBS | NUMBER(23,10) | S | SOMA DO VALOR IBSUF E IBSMUN |
| 353 | VDIFCBS | NUMBER(23,10) | S | Valor do Diferimento |
| 354 | PDIFCBS | NUMBER(7,4) | S | Percentual do diferimento |
| 355 | VDIFIBSMUN | NUMBER(23,10) | S | Valor do Diferimento |
| 356 | PDIFIBSMUN | NUMBER(7,4) | S | Percentual do diferimento |
| 357 | ALIQESPECIFICAIS | NUMBER(10,4) | S | Alíquota específica por unidade de medida apropriada |
| 358 | VDIFIBSUF | NUMBER(23,10) | S | Valor do Diferimento |
| 359 | PDIFIBSUF | NUMBER(7,4) | S | Percentual do diferimento |

## Chaves e índices

- Constraint PCITEM_PK (P): NUMPED, CODPROD, NUMSEQ
- Índice PCITEM_IDX1: NUMPED, CODPROD, SYS_NC00316$
- Índice PCITEM_IDX2: CODPROD
- Índice PCITEM_IDX3: DTULTENT
- Índice PCITEM_IDX4: CODFILIALDESTINO
- Índice PCITEM_PK UNIQUE: NUMPED, CODPROD, NUMSEQ
