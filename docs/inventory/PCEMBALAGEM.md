# RIBERDOCES.PCEMBALAGEM — dicionário

Gerado por `gateway inventory` em 2026-10-25. 149 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODAUXILIAR | NUMBER(20) | N |  |
| 2 | CODPROD | NUMBER(6) | N |  |
| 3 | EMBALAGEM | VARCHAR2(12) | S |  |
| 4 | UNIDADE | VARCHAR2(2) | S |  |
| 5 | QTUNIT | NUMBER(18,6) | S |  |
| 6 | PTABELA | NUMBER(12,3) | S |  |
| 7 | PVENDA | NUMBER(12,3) | S |  |
| 8 | DTULTALTPTABELA | DATE | S |  |
| 9 | DTULTALTPVENDA | DATE | S |  |
| 10 | CODFUNCALTPTABELA | NUMBER(8) | S |  |
| 11 | CODFUNCALTPVENDA | NUMBER(8) | S |  |
| 12 | QTMAXVENDA | NUMBER(18,6) | S |  |
| 13 | MARGEM_ESP | NUMBER(6,2) | S |  |
| 14 | MARGEM | NUMBER(6,2) | S |  |
| 15 | DTULTALTCOM | DATE | S |  |
| 16 | EXPORTACAMPO | VARCHAR2(1) | S |  |
| 17 | PRAZOVAL | NUMBER(4) | S |  |
| 18 | CODFILIAL | VARCHAR2(2) | N |  |
| 19 | POFERTA | NUMBER(12,2) | S |  |
| 20 | DTOFERTAINI | DATE | S |  |
| 21 | DTOFERTAFIM | DATE | S |  |
| 22 | PRECOANTERIOR | NUMBER(12,3) | S |  |
| 23 | EXCLUIDO | VARCHAR2(1) | S |  |
| 24 | PERMITEMULTIPLICACAO | VARCHAR2(1) | S |  |
| 25 | USABALANCATOLEDO | VARCHAR2(1) | S |  |
| 26 | TIPOEMBALAGEM | VARCHAR2(1) | S |  |
| 27 | DTEMISSAOETIQ | DATE | S |  |
| 28 | QTMINGONDOLA | NUMBER(6) | S |  |
| 29 | QTMAXGONDOLA | NUMBER(6) | S |  |
| 30 | NUMSEQATU | NUMBER | S |  |
| 31 | FATORPRECO | NUMBER(20,8) | S |  |
| 32 | QTMINATACADOF | NUMBER(6) | S |  |
| 33 | PTABELAATAC | NUMBER(12,3) | S |  |
| 34 | PVENDAATAC | NUMBER(12,3) | S |  |
| 35 | PRECOANTERIORATAC | NUMBER(12,3) | S |  |
| 36 | QTMINIMAATACADO | NUMBER(18,6) | S |  |
| 37 | QTMINIMAATACADOF | NUMBER(18,6) | S |  |
| 38 | ENVIAFV | VARCHAR2(1) | S |  |
| 39 | MARGEMATAC_ESP | NUMBER(6,2) | S |  |
| 40 | DTULTALTPTABELAATAC | DATE | S |  |
| 41 | DTULTALTPVENDAATAC | DATE | S |  |
| 42 | MARGEMATAC | NUMBER(6,2) | S |  |
| 43 | OBS | VARCHAR2(2) | S |  |
| 44 | MARGEMIDEALATAC | NUMBER(6,2) | S |  |
| 45 | DTOFERTAATACINI | DATE | S |  |
| 46 | DTOFERTAATACFIM | DATE | S |  |
| 47 | POFERTAATAC | NUMBER(18,6) | S |  |
| 48 | PESOBRUTO | NUMBER(20,8) | S | Indica o Peso Bruto da embalagem.  |
| 49 | PESOLIQ | NUMBER(20,8) | S | Indica o Peso Líquido da embalagem.  |
| 50 | ENVIABALANCA | VARCHAR2(1) | S |  |
| 51 | QTMULTIPLA | NUMBER(6) | S | Indica a quantidade múltipla embalagem.  |
| 52 | PERMITEVENDAATACADO | VARCHAR2(1) | S | Indica restrição de venda por embalagem para o Atacado.  |
| 53 | DESCRICAOECF | VARCHAR2(40) | S | Indica a descrição do ECF.  |
| 54 | ACEITAPRECOREPLICADO | VARCHAR2(1) | S | Indica se a embalagem deve ou não replicação de preço.  |
| 55 | LAYOUTETIQUETA | VARCHAR2(40) | S | Definição de layout para embalagem.  |
| 56 | PERVARIACAOPTABELA | NUMBER(20,8) | S | Indica o percentual de variação de preço. |
| 57 | DTEMISSAOETIQPOFERTAS | DATE | S | Indica a data de emissão da etiqueta de oferta. |
| 58 | DTULTALTERSRVPRC | DATE | S | Indica a data da ultima alteração nesta tabela. |
| 59 | FATORCONVERSAO | NUMBER(20,8) | S | Indica o fator de conversão para o da unidade da embalagem. |
| 60 | UNMEDIDA | VARCHAR2(4) | S | Indica a unidade de medida utilizada para a embalagem. |
| 61 | CODFUNCALTPTABELAATAC | NUMBER(8) | S | Código Funcionário que alterou o preço de tabela atacado. |
| 62 | CODFUNCALTPVENDAATAC | NUMBER(8) | S | Código Funcionário que alterou o preço de venda atacado. |
| 63 | CODFUNCALTPOFERTA | NUMBER(8) | S | Código Funcionário que alterou o preço de oferta. |
| 64 | CODFUNCALTPOFERTAATAC | NUMBER(8) | S | Código Funcionário que alterou o preço de oferta atacado. |
| 65 | PERDESCCARTAOFIDELIDADE | NUMBER(5,2) | S |  |
| 66 | ASSINATURA | VARCHAR2(255) | S | Código MD-5  |
| 67 | CONTROLEVASILHAME | VARCHAR2(1) | S | Controle Vasilhame |
| 68 | CODFUNCINATIVO | NUMBER(8) | S | Funcionário desativação embalagem |
| 69 | DTINATIVO | DATE | S | Data desativação embalagem |
| 70 | MOTIVOOFERTA | VARCHAR2(60) | S | Motivo da oferta |
| 71 | DTULTALT | DATE | S | Data  e hora de última alteração no registro |
| 72 | USARESTAURANTE | VARCHAR2(1) | S | Usa embalagem em restaurante |
| 73 | IMPDATAEMBALAGEMBALANCA | VARCHAR2(1) | S | Exporta data da embalagem para balanças |
| 74 | ENVIAINFNUTRI | VARCHAR2(1) | S | Exportar as informações nutricionais da embalagem para balança |
| 75 | PTABELAWEB | NUMBER(8,2) | S | Preço futuro que será aplicado ao produto disponibilizado na web |
| 76 | PVENDAWEB | NUMBER(8,2) | S | Preço de venda que será aplicado ao produto disponibilizado na web |
| 77 | POFERTAWEB | NUMBER(8,2) | S | Preço de oferta que será aplicado ao produto disponibilizado na web |
| 78 | DTOFERTAWEBINI | DATE | S | Data do inicio do período de vigência do preço de oferta do produto disponibilizado na web |
| 79 | DTOFERTAWEBFIM | DATE | S | Data do fim do período de vigência da oferta do produto disponibilizado na web |
| 80 | SITUACAOWEB | VARCHAR2(1) | S | Situação do Produto |
| 81 | DESCRICAOWEB | VARCHAR2(100) | S | Descrição do Produto na WEB |
| 82 | QTDIAS | NUMBER(3) | S | Qt. Dias para estoque segurança |
| 83 | QTFIXAMULTIPLICCHECKOUT | NUMBER(9,3) | S | Quantidade fixa para multiplicação no checkout |
| 84 | DTULTALTPTABELAWEB | DATE | S | Data da última alteração do preço de tabela para web |
| 85 | CODCOB | VARCHAR2(4) | S | Código da cobrança |
| 86 | CODCLI | NUMBER(9) | S | Código do cliente |
| 87 | QTEMISSAOETIQ | NUMBER(8) | S | Quantidade de etiquetas a emitir desta embalagem |
| 88 | CODINFEXTRABAL | NUMBER(6) | S | Informação extra usada para vincular uma receita ao exportar a embalagem para a balança. |
| 89 | TARAF | NUMBER(7,3) | S | Tara da balança |
| 90 | PERMITEVENDACOMINTERNO | VARCHAR2(1) | S | Permite venda com código interno - PDV. |
| 91 | INDICPESAGEMOBRIGATORIA | VARCHAR2(1) | S | Indica se a pesagem e obrigatória - PDV |
| 92 | INDICPERMITEPORPREVENDA | VARCHAR2(1) | S | Indica se o produto só pode ser vendido dentro de uma pré venda. |
| 93 | INDICPERMITEDIGITACAOPRECO | VARCHAR2(1) | S | Permite digitação de preço. |
| 94 | INDICPERMITEDIGITACAODESCONTO | VARCHAR2(1) | S | Permite digitação de desconto. |
| 95 | INDICGERARCODPESOVAR | VARCHAR2(1) | S | Permite gerar código peso variável. |
| 96 | INDICALERTAQUANTIDADE | VARCHAR2(1) | S | Indica se o sistema deve apresentar o alerta de quantidade para operador. |
| 97 | INDICPRODUCAOPROPRIA | VARCHAR2(1) | S | Indica a produção própria. |
| 98 | CODDIGITAQTDE | VARCHAR2(1) | S | Permite digitação da quantidade. |
| 99 | PTABELAANT | NUMBER(12,3) | S | Preço futuro anterior |
| 100 | PTABELAATACANT | NUMBER(12,3) | S | Preço futuro atacado anterior |
| 101 | PVENDAWEBANT | NUMBER(18,6) | S | Preço venda web anterior da embalagem. |
| 102 | PTABELAWEBANT | NUMBER(18,6) | S | Preço futuro web anterior. |
| 103 | PCOMINT1 | NUMBER(6,2) | S | Taxa de comissão para vendedor interno. |
| 104 | PCOMEXT1 | NUMBER(6,2) | S | Taxa de comissão para vendedor externo. |
| 105 | PCOMREP1 | NUMBER(6,2) | S | Taxa de comissão para representante. |
| 106 | INDCVENDECODINTERNO | VARCHAR2(1) | S | Permite vender com código interno. |
| 107 | BEBALCOOLICA | VARCHAR2(1) | S | Definição de Item como bebida alcoolica ou não. |
| 108 | DTAPLICPRECOVAREJO | DATE | S | Data de agendamento de preço futuro varejo. |
| 109 | DTAPLICPRECOATAC | DATE | S | Data de agendamento de preço futuro atacado. |
| 110 | SOCIOTORCEDOR | VARCHAR2(1) | S | Embalagem sócio torcedor |
| 111 | DESTINOOFERTAVAREJO | VARCHAR2(2) | S | Define para onde será aplicado o preço de oferta de varejo |
| 112 | DESTINOOFERTAATAC | VARCHAR2(2) | S | Define para onde será aplicado o preço de oferta de atacado |
| 113 | GIROMES | NUMBER(16,3) | S | Giro mês da embalagem |
| 114 | GIROMEDIODIA | NUMBER(16,3) | S | Giro médio dia da embalagem |
| 115 | DTULTALTGIRO | DATE | S | Data da última alteração de giro |
| 116 | ENVIATELEMARKETING | VARCHAR2(1) | S | Envia telemarketing |
| 117 | ENVIAFRENTECAIXA | VARCHAR2(1) | S | Envia frente de caixa |
| 118 | IMPRIMERESTAURANTE | VARCHAR2(1) | S | Imprime restaurante |
| 119 | ALTURA | NUMBER(20,8) | S | Altura da embalagem |
| 120 | LARGURA | NUMBER(20,8) | S | Largura da embalagem |
| 121 | COMPRIMENTO | NUMBER(20,8) | S | Comprimento da embalagem |
| 122 | VOLUME | NUMBER(20,8) | S | Volume da embalagem |
| 123 | OBRIGALEITURACODBARRA | VARCHAR2(1) | S | Obriga leitura de código de barras da embalagem |
| 124 | IMPRESSAORESTAURANTE | VARCHAR2(1) | S | Permite impressão restaurante. |
| 125 | PRODUTOCOZINHA | VARCHAR2(1) | S | Informa se é producao de cozinha. |
| 126 | TEMPOPREPARO | NUMBER(5) | S | Tempo de Preparo no caso de produto de cozinha. |
| 127 | TEMPOALERTA | NUMBER(5) | S | Tempo de alerta para o preparo no caso de produto de cozinha. |
| 128 | PERMITEIMPRESSAOETIQUETA | VARCHAR2(1) | S | Flag para indicar se embalagem permite impressão de Etiqueta |
| 129 | IONSYNC | CHAR(1) | S |  |
| 130 | USAECOMMERCEUNILEVER | VARCHAR2(1) | S | Valida se a embalagem será enviada ou não para o e-commerce da Unilever |
| 131 | USAETIQUETAANTIFURTO | VARCHAR2(1) | S | INDICA SE PRODUTO POSSUI ETIQUETA ANTI-FURTO |
| 132 | MD5PAF | VARCHAR2(200) | S | Assinatura MD5 do registro do PAF |
| 133 | ACEITAMARGEMREPLICADA | CHAR(1) | S | Define se aceita ou não replicação de margem |
| 134 | ENVIAPENTREGA | CHAR(1) | S | Valor para definir se o produto é ou não a pronta entrega |
| 135 | INFOPRODWEB | CLOB | S | Informações do Produto WEB |
| 136 | IDCIASHOP | NUMBER(10) | S | Indentificador do embalagem vindo do ciashop |
| 137 | PRODSEMCODBARRAS | CHAR(1) | S | Define se o produto tem ou não código de barras para listagem de embalagens no Caixa |
| 138 | PERMITEVENDACRTALIMENTACAO | VARCHAR2(1) | S | Define se permite ou não vender o produto utilizando cartão alimentação |
| 139 | QTDVIASIMPRESSAOCOZINHA | NUMBER(8) | S | Quantidade de vias de impressão para embalagens que utilizam a opção "Produto Cozinha" |
| 140 | JUSTIFICATIVAPRECO | VARCHAR2(50) | S | JUSTIFICATIVA DE ALTERACAO DE PRECO |
| 141 | CODFRACIONADOR | VARCHAR2(4) | S | Código do Fracionador |
| 142 | CODAUXILIARANTERIOR | NUMBER(20) | S | Código auxiliar que possuía antes de ser editado. |
| 143 | ENVIAECOMMERCE | VARCHAR2(1) | S | E enviado para o e commerce |
| 144 | ATIVOECOMMERCE | VARCHAR2(1) | S | Está ativo para o e-commerce |
| 145 | CODFILIALINTEGRACAO | NUMBER(3) | S | Código da filial de integração |
| 146 | DTALTERC5 | TIMESTAMP(6) | S | Data de alteração |
| 147 | DTCADASTRO | TIMESTAMP(6) | S | Data de cadastro |
| 148 | DTULALTERINTEGRA | TIMESTAMP(6) | S | Data alteração |
| 149 | CHECAPESOETIQUETA | VARCHAR2(1) | S | Indica se a embalagem deve realizar checagem de peso e etiqueta durante os processos operacionais |

## Chaves e índices

- Constraint PCEMBALAGEM_PK (P): CODFILIAL, CODAUXILIAR
- Índice IDX_E_SINC_C5_COVERING: DTALTERC5, CODFILIAL, CODPROD, PVENDA, DTULTALTPVENDA, PVENDAATAC, DTULTALTPVENDAATAC, CODAUXILIAR
- Índice PCEMBALAGEM_C5_IDX01: CODPROD, CODFILIAL, SYS_NC00149$, DTINATIVO, CODAUXILIAR, DTALTERC5
- Índice PCEMBALAGEM_IDX1: CODPROD, CODAUXILIAR
- Índice PCEMBALAGEM_IDX3: DTOFERTAFIM
- Índice PCEMBALAGEM_IDX4: DTOFERTAINI
- Índice PCEMBALAGEM_IDX5: ENVIAFRENTECAIXA
- Índice PCEMBALAGEM_IDX6: QTUNIT, QTMINIMAATACADO
- Índice PCEMBALAGEM_IDX7: CODPROD, CODFILIAL, QTMINGONDOLA, QTMAXGONDOLA
- Índice PCEMBALAGEM_IDX_C5_TESTE: DTALTERC5, CODFILIAL, CODPROD
- Índice PCEMBALAGEM_PK UNIQUE: CODFILIAL, CODAUXILIAR
