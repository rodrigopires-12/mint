# RIBERDOCES.PCDEPTO — dicionário

Gerado por `gateway inventory` em 2026-10-25. 24 colunas.

| # | Coluna | Tipo Oracle | Nulo | Comentário |
|---|---|---|---|---|
| 1 | CODEPTO | NUMBER(6) | N |  |
| 2 | DESCRICAO | VARCHAR2(25) | N |  |
| 3 | PERCPARTVENDAPREV | NUMBER(8,4) | S |  |
| 4 | PERCMARGEMPREV | NUMBER(8,4) | S |  |
| 5 | TIPOMERC | VARCHAR2(2) | S |  |
| 6 | EMITEQTUNIT | VARCHAR2(1) | S |  |
| 7 | ATUALIZAINVGERAL | VARCHAR2(1) | S |  |
| 8 | MARGEMPREVISTA | NUMBER(6,2) | S | % de margem prevista padrão de lucratividade |
| 9 | REFERENCIA | VARCHAR2(20) | S | Referência |
| 10 | PERDESCMAXIDEAL | NUMBER(10,2) | S | % Desconto máximo ideal |
| 11 | PERDESCMAXPOSSIVEL | NUMBER(10,2) | S | % Desconto máximo possível |
| 12 | PERDESCMAXAVISTA | NUMBER(10,2) | S | % Desconto máximo venda a vista |
| 13 | PERCCOMGARANTIDA | NUMBER(10,2) | S | % Comissão mínima garantida |
| 14 | IDINTEGRACAOCIASHOP | VARCHAR2(250) | S | Referência do departamento no E-commerce CiaShop |
| 15 | ENVIAECOMMERCE | VARCHAR2(1) | S | Identifica se o produto será enviado ao e-commerce |
| 16 | DTULTALTER | DATE | S | Data da ultima alteração |
| 17 | CODCAMPLOMADEE | VARCHAR2(200) | S |  Código da Campanha |
| 18 | CODADWORDS | VARCHAR2(200) | S | Código de Remarketing |
| 19 | ATIVO | VARCHAR2(1) | S | Ativo |
| 20 | DESCRICAOECOMMERCE | VARCHAR2(200) | S | Descrição do e-commerce |
| 21 | TITULO | VARCHAR2(200) | S | Titulo |
| 22 | CODDEPTOPRINC | NUMBER(6) | S | Código do departamento principal |
| 23 | DTCADASTRO | DATE | S | Data de cadastro |
| 24 | DTALTERC5 | TIMESTAMP(6) | S | DATA ALTERACAO |

## Chaves e índices

- Constraint PCDEPTO_PK (P): CODEPTO
- Índice PCDEPTO_PK UNIQUE: CODEPTO
