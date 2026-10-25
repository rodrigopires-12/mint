# Perguntas abertas

Dúvidas registradas durante a Fase 1. Cada item traz a decisão provisória adotada enquanto não há resposta.

## Infraestrutura e plataforma

| # | Pergunta | Provisório | Aberta em |
|---|---|---|---|
| Q1 | Nome definitivo do módulo Go e do repositório GitHub. | `github.com/the12company/mint` | 2026-10-25 |
| Q2 | Hospedagem do ClickHouse em staging/prod: o ClickHouse Cloud só oferece `sa-east-1` como private region (Enterprise). Ver ADR 0002. | Dev usa ClickHouse OSS em EC2 `sa-east-1`; prod decide antes da Fase 2 | 2026-10-25 |
| Q3 | Clientes com Windows Server 2012 R2: o gateway (Go 1.27) exige 2016+. Usar modo direto para eles? | Fora do gateway na v1 | 2026-10-25 |
| Q4 | Assinatura de código do `gateway.exe` (certificado) e empacotamento MSI. | Zip + `install.ps1` sem assinatura | 2026-10-25 |
| Q5 | Qual alias TNS é o Oracle de teste (`DISTRIBUIDORA` ou `INDUSTRIA`) e qual o owner das tabelas `PC*` (`WINT`?). | `doctor` resolve o owner por sinônimos/`ALL_TABLES` | 2026-10-25 |

## Modelo e tipos

| # | Pergunta | Provisório | Aberta em |
|---|---|---|---|
| Q6 | `NUMBER` sem precisão declarada na `raw`: `String` decimal ou `Decimal(38,10)`? | Decidir no mapeamento de tipos do E1.4 | 2026-10-25 |
| Q7 | Pontos #1, #2, #11, #12, #13 e #14 de `docs/escopo-v1.md` (dependem só do dicionário). | Respondidos no E1.2 | 2026-10-25 |
| Q8 | Pontos #3 a #10 de `docs/escopo-v1.md` (dependem de especialista de negócio). | Seguem como no escopo | 2026-10-25 |
| Q9 | Colunas de watermark por tabela: nenhuma tabela do Winthor tem timestamp de alteração confiável conhecido. | Data de negócio + janela móvel (45 dias); validação por tabela no E1.2 | 2026-10-25 |
