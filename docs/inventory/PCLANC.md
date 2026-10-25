# RIBERDOCES.PCLANC — dicionário

Gerado por `gateway inventory` em 2026-10-25. 203 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | RECNUM | NUMBER(8) | N |  |
| 2 | DTLANC | DATE | S |  |
| 3 | CODCONTA | NUMBER(10) | N |  |
| 4 | CODFORNEC | NUMBER(9) | N | Descricao coluna CODFORNEC |
| 5 | HISTORICO | VARCHAR2(200) | N |  |
| 6 | NUMNOTA | NUMBER(10) | N |  |
| 7 | DUPLIC | VARCHAR2(1) | S |  |
| 8 | VALOR | NUMBER(12,2) | N |  |
| 9 | DTVENC | DATE | S |  |
| 10 | VPAGO | NUMBER(12,2) | S |  |
| 11 | DTPAGTO | DATE | S |  |
| 12 | LOCALIZACAO | VARCHAR2(200) | S | Descricao coluna LOCALIZACAO |
| 13 | NUMNOTADEV | NUMBER(10) | S |  |
| 14 | VALORDEV | NUMBER(12,2) | S |  |
| 15 | CODFILIAL | VARCHAR2(2) | N |  |
| 16 | INDICE | VARCHAR2(1) | N | Indice de risco do título. |
| 17 | PRORROG | NUMBER(4) | S |  |
| 18 | NUMCHEQUE | VARCHAR2(20) | S |  |
| 19 | NUMBORDERO | NUMBER(6) | S |  |
| 20 | DTBORDER | DATE | S |  |
| 21 | VPAGOBORDERO | NUMBER(14,2) | S |  |
| 22 | NUMSEQBORDERO | NUMBER(6) | S |  |
| 23 | NUMCHEQUE2 | NUMBER(8) | S |  |
| 24 | NUMBANCO | NUMBER(4) | S |  |
| 25 | DESCONTOFIN | NUMBER(14,2) | S |  |
| 26 | TXPERM | NUMBER(14,2) | S |  |
| 27 | TIPOLANC | VARCHAR2(1) | S |  |
| 28 | DTEMISSAO | DATE | S |  |
| 29 | DTCHEQ | DATE | S |  |
| 30 | NUMVIA | NUMBER(2) | S |  |
| 31 | TIPOPARCEIRO | VARCHAR2(1) | S |  |
| 32 | NOMEFUNC | VARCHAR2(40) | S |  |
| 33 | ASSINATURA | VARCHAR2(40) | S |  |
| 34 | DTASSINATURA | DATE | S |  |
| 35 | TIPOPAGTO | VARCHAR2(3) | S |  |
| 36 | NUMBCODESTDOC | NUMBER(4) | S |  |
| 37 | NUMAGDESTDOC | NUMBER(6) | S |  |
| 38 | NUMCCDESTDOC | NUMBER(12) | S |  |
| 39 | NUMDVDESTDOC | VARCHAR2(2) | S |  |
| 40 | NUMCODBARRA | VARCHAR2(44) | S |  |
| 41 | NUMTRANSENT | NUMBER(10) | S |  |
| 42 | NUMTAB | NUMBER(8) | S |  |
| 43 | CODFORNECPRINC | NUMBER(9) | S | Descricao coluna CODFORNECPRINC |
| 44 | HISTORICO2 | VARCHAR2(200) | S |  |
| 45 | MOEDA | VARCHAR2(4) | S | Descricao coluna MOEDA |
| 46 | DTMOEDA | VARCHAR2(1) | S |  |
| 47 | VLIRRF | NUMBER(8,2) | S |  |
| 48 | VLSESTSENAT | NUMBER(18,6) | S | Grava o valor do SEST/SENAT. |
| 49 | NUMTRANSADIANTFOR | NUMBER(10) | S |  |
| 50 | CODFUNCBAIXA | NUMBER(8) | S |  |
| 51 | NUMAGDESTDOC2 | NUMBER(6) | S |  |
| 52 | NUMVIASFICHACP | NUMBER(4) | S |  |
| 53 | DVAG | VARCHAR2(2) | S |  |
| 54 | CODFUNCAUTOR | NUMBER(8) | S |  |
| 55 | DTAUTOR | DATE | S |  |
| 56 | CODROTINABAIXA | NUMBER(6) | S |  |
| 57 | NUMTRANS | NUMBER(10) | S |  |
| 58 | PORTADORCHEQUE | VARCHAR2(40) | S |  |
| 59 | CODFUNCRECLASSIFIC | NUMBER(8) | S |  |
| 60 | DTRECLASSIFIC | DATE | S |  |
| 61 | CODFUNCESTORNOBAIXA | NUMBER(8) | S |  |
| 62 | DTESTORNOBAIXA | DATE | S |  |
| 63 | DTPRORROG | DATE | S |  |
| 64 | CODCOMPRADOR | NUMBER(8) | S | Comprador do produto. Na rotina de 1301 - Entrada de Mercadoria será pesquisado o código comprador do pedido de compra e gravar nesse campo. |
| 65 | MOEDAESTRANGEIRA | NUMBER(6) | S |  |
| 66 | COTACAO | NUMBER(18,6) | S |  |
| 67 | RECNUMPRINC | NUMBER(8) | S | Armazena em cada lançamento de rateio, o número do primeiro lançamento gerado. Agrupara lançamentos gerados que foram rateados em diversas contas diferentes. |
| 68 | COTACAOBAIXA | NUMBER(18,6) | S |  |
| 69 | FROTA_QTLITROS | NUMBER(20,6) | S | Dados de Frota: quantidade de litros que foi abastecido (rotinas 631/749).  |
| 70 | FROTA_CODPRACA | NUMBER(4) | S | Dados de Frota: praça que foi feito o abastecimento, conserto, troca de pneu, travessia de balsa, etc. (rotinas 631/749).  |
| 71 | FROTA_NUMCAR | NUMBER(8) | S | Dados de Frota: carregamento (viagem) em que as despesas foram feitas (rotina 749). |
| 72 | FROTA_CODVEICULO | NUMBER(4) | S | Dados de Frota: código do veículo. Necessário para infomar despesas que não são feitas em viagem, como pintura de baú (rotina 749). |
| 73 | FROTA_COMISSAO | NUMBER(14,2) | S | Dados de Frota: valor da comissão paga ao motorista (rotina 1206).  |
| 74 | NUMTRANSVENDA | NUMBER(10) | S |  |
| 75 | NUMDIIMPORTACAO | VARCHAR2(40) | S | Campo utilizado para armazenar o número do documento de importação  |
| 76 | ADIANTAMENTO | VARCHAR2(1) | S | Campo para identificar se o lançamento de contas a pagar é proveniente de um adiantamento - (Rotina 745). |
| 77 | FROTA_DTABASTECE | DATE | S | Dados de Frota: data e horário do abastecimento.  |
| 78 | FROTA_KMABASTECE | NUMBER(12,2) | S | Dados de Frota: quilometragem do abastecimento.  |
| 79 | FROTA_PESO | NUMBER(5) | S | Dados de Frota: quantidade de peso do frete. |
| 80 | DTCOTACAO | DATE | S | Campo para armazenar a data da cotação do pedido de compra. |
| 81 | VALORMOEDAESTRANGEIRA | NUMBER(18,6) | S | Campo para armazenar o valor do lançamento em moeda estrangeira. |
| 82 | PERCDESCFINREAL | NUMBER(14,2) | S | Indica o percentual de desconto financeiro. |
| 83 | PREST | VARCHAR2(2) | S | Ligação com PCPREST.PREST. |
| 84 | NUMNEGOCIACAO | NUMBER(12) | S | Numero da negociação de importação. |
| 85 | NFSERVICO | VARCHAR2(1) | S | Campo para indicar se o título gerado é referente a uma nota fiscal de serviço. |
| 86 | DTPAGADIANTFOR | DATE | S | Data que o lancto. adiant. fornecedor foi pago (conciliado) pela 746. |
| 87 | DTULTALTER | DATE | S | Data da última alteração do lançamento pela rotina 750. |
| 88 | CODFUNCULTALTER | NUMBER(8) | S | Código do funcionário da última alteração do lançamento pela rotina 750. |
| 89 | BOLETO | VARCHAR2(1) | S | . |
| 90 | CODPROJETO | NUMBER(6) | S | Código de Projeto. Cadastrado na 607 e informado na 749 e 631. |
| 91 | DTDESD | DATE | S | Indica a data de desdobramento dos títulos gerados. |
| 92 | PAGTONOMEDOCLIENTE | VARCHAR2(1) | S | Indica se o responsável pela pagamento da GNRE é o cliente. |
| 93 | DTCOMPETENCIA | DATE | S | Indica a data da competencia. |
| 94 | VLISS | NUMBER(18,6) | S | Grava o valor do ISS da NF de serviço. |
| 95 | VLINSS | NUMBER(18,6) | S | Grava o valor do INSS da NF de serviço. |
| 96 | VLCSRF | NUMBER(18,6) | S | Grava o valor do CSRF da NF de serviço. |
| 97 | FUNCBLOQ | NUMBER(10) | S | Funcionário que realizou o bloqueio. |
| 98 | DTBLOQ | DATE | S | Data do bloqueio. |
| 99 | OBSBLOQ | VARCHAR2(200) | S | Observação de bloqueio. |
| 100 | VLPIS | NUMBER(18,6) | S | Valor do imposto PIS. |
| 101 | VLCOFINS | NUMBER(18,6) | S | Valor do imposto COFINS. |
| 102 | IDCONTROLEEMBARQUE | VARCHAR2(20) | S | Identificação do controle de embarque. |
| 103 | NUMCONTRATOCAMBIO | VARCHAR2(20) | S | Número do contrato de câmbio. |
| 104 | VLVARIACAOCAMBIAL | NUMBER(18,6) | S | Valor da variação cambial. |
| 105 | VLOUTRAS | NUMBER(12) | S | Outros valores. |
| 106 | DTAGENDAMENTO | DATE | S | Data do agendamento. |
| 107 | AGENDAMENTO | VARCHAR2(1) | S | Indica se o lançamento foi enviado para agendamento S ou N. |
| 108 | NUMTRANSENTNF | NUMBER(10) | S | Número de transação da NF de entrada |
| 109 | FORMAPGTO | VARCHAR2(4) | S | Forma de pagamento |
| 110 | CODROTINACAD | VARCHAR2(40) | S | Rotina de cadastro do contas a pagar |
| 111 | CODROTINAALT | VARCHAR2(40) | S | Rotina de alteração do contas a pagar |
| 112 | CODRECEITA | NUMBER(4) | S | Código da Receita para Lançamentos de Tributos |
| 113 | VRECEITABRUTA | NUMBER(18,2) | S | Valor da Receita bruta utilizado no darf simples |
| 114 | PERCRECEITABRUTA | NUMBER(4,2) | S | Percentual da receita bruta utilizado no darf simples |
| 115 | DTCANCEL | DATE | S | Data cancelamento contas a pagar de entrada de NF. |
| 116 | NUMPARCELAMENTO | NUMBER(6) | S | Número chave do parcelamento |
| 117 | NUMSEQPARCELAMENTO | NUMBER(10) | S | Número sequencial de parcelamento utilizado para ordenação |
| 118 | PARCELA | VARCHAR2(6) | S | Número de prestação |
| 119 | DTIMPRESSAO | DATE | S | Data de impressão do Cheque/Boleto. |
| 120 | FORNECEDOR | VARCHAR2(60) | S | Parceiro do lançamento. |
| 121 | VLRPARIDADEMOEDA | NUMBER(14,2) | S | Valor da Paridade da Moeda. |
| 122 | MOEDAPARIDADE | NUMBER(6) | S | Moeda da Paridade. |
| 123 | NUMSEQCONTRATOEMPRESTIMO | NUMBER(10) | S | Número sequencial do contrato. |
| 124 | CODLANCAMENTORM | VARCHAR2(30) | S | Código de lançamento RM |
| 125 | RECNUMBAIXA | NUMBER(8) | S | RecNum do lançamento principal da baixa. |
| 126 | TIPOAUTORIZACAO | VARCHAR2(2) | S | Tipo de autorização do pagamento do borderô. |
| 127 | VLRUTILIZADOADIANTFORNEC | NUMBER(12,2) | S | Valor que já foi utilizado pelo adiantamento a Fornecedor |
| 128 | NUMBANCOCONTRATO | NUMBER(10) | S | Número do banco do contrato de cambio |
| 129 | LACREDIGCONECSOCIAL | NUMBER(11) | S | Lacre + Digito de Conectividade Social |
| 130 | OBSRETORNO | VARCHAR2(200) | S | Observação do processamento do arquivo retorno. |
| 131 | VLDESPFIN | NUMBER(18,6) | S | Valor Desp. financ. (Vendor) |
| 132 | VLINSSTOMADOR | NUMBER(18,6) | S | Grava o valor do INSS do tomador de serviço. |
| 133 | DTPROCESSAMENTO | DATE | S | Data de processamento do arquivo retorno. |
| 134 | DTREMESSA | DATE | S | Data de geração da remessa de pagamento. |
| 135 | CODSERVICOSPEDEFD | NUMBER(8) | S | Código do serviço da NF de serviço. |
| 136 | IDENTIFICADORFGTS | VARCHAR2(16) | S | Identificador do FGTS |
| 137 | SERIE | VARCHAR2(5) | S | Serie |
| 138 | CHAPA_RM | VARCHAR2(16) | S | Chapa de identificação da RM |
| 139 | GNRE | VARCHAR2(1) | S | Informar se o titulo foi gerado pela GNRE |
| 140 | DTVENCORIGEM | DATE | S | Data de vencimento origem |
| 141 | CODROTINAVERSAOCAD | VARCHAR2(100) | S | Código e versão da rotina de cadastro |
| 142 | CODROTINAVERSAOALT | VARCHAR2(100) | S | Código e versão da rotina de alteração |
| 143 | CODROTINAVERSAOBAIXA | VARCHAR2(100) | S | Código e versão da rotina de baixa |
| 144 | CODROTINAVERSAODESD | VARCHAR2(100) | S | Código e versão da rotina de desdobramento |
| 145 | CODCOBSEFAZ | VARCHAR2(4) | S | Código de cobrança da SEFAZ |
| 146 | BANDEIRACARTAO | NUMBER(3) | S | Bandeira do Cartão de Credito ou Debito |
| 147 | CNPJCREDENCCARTAO | VARCHAR2(18) | S | CNPJ da Credenciadora do Cartão |
| 148 | RAZAOCREDENCCARTAO | VARCHAR2(60) | S | Razão Social da Credenciadora do Cartão |
| 149 | NSUTEF | VARCHAR2(15) | S | Número de autorização do Cartão |
| 150 | OPCAOPAGAMENTOIPVA | NUMBER(1) | S | Campo referente a opção de pagamento de IPVA |
| 151 | TIPOSERVICO | VARCHAR2(2) | S | Tipo de Serviço (20 - Pag. Fornecedor, 30 - Pag. Salários, 22 - Pag. Tributos, 99 - Outros Pag.) |
| 152 | DTENVIOFLUIG | DATE | S | Campo indicador de envio de documento para o Fluig |
| 153 | NUMSOLICITACAOFLUIG | NUMBER(10) | S | Campo com numero de solicitacao gerado pelo Fluig para indicacao de recebimento do documento pelo mesmo. |
| 154 | QTDENVIOSFLUIG | NUMBER(4) | S | Campo com a quantidade de vezes de envio do documento ao Fluig. |
| 155 | CODFUNCAUTSUPLORC | NUMBER(8) | S | CODIGO FUNCIONARIO AUTORIZACAO SUPLEMENTACAO ORCAMENTARIA |
| 156 | DATAAUTSUPLORC | DATE | S | DATA AUTORIZACAO SUPLEMENTACAO ORCAMENTARIA |
| 157 | CODMOEDABAIXA | VARCHAR2(4) | S | Código da moeda de baixa de movimentação bancária |
| 158 | DTREJEICAOSUPLORC | DATE | S | Data da rejeição da suplementação orçamentária |
| 159 | CODFUNCREJEICAOSUPLORC | NUMBER(8) | S | Funcionário que rejeitou a suplementqação orçamentária |
| 160 | HISTORICOREJEICAOSUPLORC | VARCHAR2(200) | S | Histórico da rejeição da suplementação orçamentária |
| 161 | DTAUTOR2 | DATE | S | Data da segunda autorização |
| 162 | CODFUNCAUTOR2 | NUMBER(8) | S | Código do segundo funcionário que atutorização |
| 163 | DATACONSOLIDACAOPREFAT | DATE | S | Data consolidação Pré-Faturamento |
| 164 | UTILIZOURATEIOCONTA | VARCHAR2(1) | S | Definte se o lançamento usou o rateio de conta |
| 165 | PRCRATEIOUTILIZADO | NUMBER(18,6) | S | Percentual do rateio que foi utilizado da conta gerencial |
| 166 | DTAUTOR1 | DATE | S | Data da primeira autorização |
| 167 | CODFUNCAUTOR1 | NUMBER(8) | S | Código do primeiro funcionário que atutorização |
| 168 | AUTENTICACAO | VARCHAR2(64) | S | Autenticação bancária Segmento Z |
| 169 | NUMPRESTACAOCONTA | NUMBER(10) | S | Número da prestação de contas da rot. 782 |
| 170 | NUMADIANTAMENTO | NUMBER(10) | S | Número do adiantamento da rot. 782 |
| 171 | DTCOTACAOBAIXA | DATE | S | Data de cotação na quitação do título negociado em moeda estrangeira. |
| 172 | CODSERVICOSPEDIRPF | NUMBER(10) | S | Código do serviço de rendimento de pessoa física |
| 173 | CODSERVICOSPEDIRPJ | NUMBER(10) | S | Código do serviço de rendimento de pessoa jurídica |
| 174 | REINFEVENTOR4040 | VARCHAR2(1) | S | Definir se o lançamento se refere ao REINF Evento 4040 |
| 175 | CODSERVICOSPEDBNI | NUMBER(10) | S | Código do serviço de rendimento de beneficiário não identificado |
| 176 | DTCOMPETENCIAIRRETIDO | DATE | S | Informar o mês competência do IRRF para EFD-REINF |
| 177 | CHAVEPIX | VARCHAR2(100) | S | Chave pix para transferência |
| 178 | TIPOCHAVEPIX | VARCHAR2(20) | S | Descrição do tipo da chave pix |
| 179 | CODTIPOCHAVEPIX | VARCHAR2(2) | S | Código do tipo da chave pix |
| 180 | QRCODEPIX | VARCHAR2(500) | S | Qr-code para pagamento com pix |
| 181 | RESCISAO | VARCHAR2(1) | S | Campo utilizado para informar se é contas a pagar de rescisão de contrato de RCA |
| 182 | CODAGRUPAMENTO | NUMBER(10) | S | Código do agrupamento - Vínculo com PCLANCAGRUPADO |
| 183 | DTAGRUPAMENTO | DATE | S | Data de agrupamento |
| 184 | DTCANCELAGRUPAMENTO | DATE | S | Data do cancelamento do agrupamento |
| 185 | CODFUNCAGRUPAMENTO | NUMBER(8) | S | Código funcionário que agrupou o título |
| 186 | CODFUNCCANCELAGRUPAMENTO | NUMBER(8) | S | Código do funcionário que cancelou o agrupamento |
| 187 | AGRUPADO | VARCHAR2(1) | S | Título agrupado |
| 188 | PROVPARCEMP | VARCHAR2(1) | S | Parcelamento proveniente de empréstimo de contrato FINIMP |
| 189 | LANCEXCLUIDO | VARCHAR2(1) | S | Campo utilizado pelo módulo contábil para identificar lançamentos que foram desdobrados e cancelados posteriormente |
| 190 | VLBASEIR | NUMBER(18,6) | S | Valor Base IR |
| 191 | VLBASEPCC | NUMBER(18,6) | S | Valor Base PIS COFINS CSRF |
| 192 | LOGENVIOCP | VARCHAR2(200) | S | LOG DE ENVIO DE CONTAS A PAGAR |
| 193 | DTENVIOCP | DATE | S | Data de envio do contas a pagar |
| 194 | STATUSENVIOCP | VARCHAR2(30) | S | Status do envio de contas a pagar |
| 195 | LOGRETORNOCP | VARCHAR2(200) | S | Log de retorno de contas a pagar |
| 196 | STATUSRETORNOCP | VARCHAR2(30) | S | Status do retorno do contas a pagar |
| 197 | DTRETORNOCP | DATE | S | Data de retorno do contas a pagar |
| 198 | JSONENVIOCP | CLOB | S | Json de envio do contas a pagar |
| 199 | JSONRETORNOCP | CLOB | S | Json de retorno do contas a pagar |
| 200 | MOTIVONAOEFETIVACAOCP | VARCHAR2(200) | S | Descrição do motivo da não efetivação de pagamento do contas a pagar |
| 201 | DTGERACAOTXTCP | DATE | S | Data de geração do arquivo remessa do contas a pagar |
| 202 | NUMNOTASERVICO | VARCHAR2(60) | S | Número de Nota de Serviço |
| 203 | BORDEROESCRITURAL | VARCHAR2(1) | S | Determina se o borderô é do tipo escritural |

## Chaves e índices

- Constraint PCLANC_PK (P): RECNUM
- Índice PCLANC_IDX1: DTLANC
- Índice PCLANC_IDX10: DTPAGTO, CODFILIAL
- Índice PCLANC_IDX11: CODFILIAL, CODCONTA
- Índice PCLANC_IDX13: CODCONTA, LOCALIZACAO, TIPOPARCEIRO, CODROTINABAIXA
- Índice PCLANC_IDX14: NUMBORDERO, NUMBANCO
- Índice PCLANC_IDX15: NUMCHEQUE, NUMBANCO
- Índice PCLANC_IDX16: NUMTRANS, NUMBANCO
- Índice PCLANC_IDX17: SYS_NC00136$, SYS_NC00137$
- Índice PCLANC_IDX18: RECNUMPRINC
- Índice PCLANC_IDX19: NUMTRANSADIANTFOR, SYS_NC00140$, DTESTORNOBAIXA
- Índice PCLANC_IDX2: CODCONTA, NUMTRANSENT
- Índice PCLANC_IDX20: DTPAGTO, CODFILIAL, CODCONTA
- Índice PCLANC_IDX21: DTPAGTO, NUMBORDERO, CODFILIAL
- Índice PCLANC_IDX22: CODFORNEC, DTPAGTO, NUMBORDERO
- Índice PCLANC_IDX23: CODSERVICOSPEDEFD, RECNUM, DTEMISSAO
- Índice PCLANC_IDX24: DTLANC, CODFORNEC, CODFILIAL
- Índice PCLANC_IDX25: NUMTRANSENT
- Índice PCLANC_IDX26: NUMTRANSVENDA
- Índice PCLANC_IDX27: CODSERVICOSPEDIRPF, RECNUM, DTCOMPETENCIA
- Índice PCLANC_IDX28: CODSERVICOSPEDIRPJ, RECNUM, DTEMISSAO, DTVENC
- Índice PCLANC_IDX29: DTEMISSAO
- Índice PCLANC_IDX3: CODFORNEC
- Índice PCLANC_IDX30: DTCOMPETENCIAIRRETIDO
- Índice PCLANC_IDX31: NUMNOTASERVICO
- Índice PCLANC_IDX4: NUMNOTA
- Índice PCLANC_IDX5: DTVENC
- Índice PCLANC_IDX6: CODFORNECPRINC, DTPAGTO
- Índice PCLANC_IDX7: NUMBORDERO, NUMSEQBORDERO
- Índice PCLANC_IDX8: NUMCHEQUE, RECNUM
- Índice PCLANC_IDX9: NUMTRANSENT, CODCONTA, RECNUM
- Índice PCLANC_IDXPCSIS500_1: SYS_NC00208$
- Índice PCLANC_IDXPCSIS500_2: SYS_NC00209$
- Índice PCLANC_IDXPCSIS500_3: SYS_NC00210$
- Índice PCLANC_IDXPCSIS500_4: SYS_NC00211$
- Índice PCLANC_IDXPCSIS500_5: SYS_NC00212$, CODCONTA, DTCANCEL
- Índice PCLANC_PK UNIQUE: RECNUM
