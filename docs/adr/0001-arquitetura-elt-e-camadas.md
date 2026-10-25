# ADR 0001 — Arquitetura ELT e camadas do DW

- Status: aceito (2026-10-25)
- Origem: `CLAUDE.md`, decisões 3 a 8, 10 e 11

## Contexto

Cerca de 100 clientes rodam o mesmo ERP Winthor v30 sobre Oracle. Precisamos consolidar os dados em um DW multi-tenant e entregar dashboards, relatórios agendados e acesso via Claude (MCP), sem customização por cliente na v1 e sem jamais escrever no Oracle do cliente.

## Decisão

1. **ELT a partir das tabelas `PC*`**, extraídas 1:1 (só as colunas usadas), sem views nem consultas prontas no Oracle. Um `sources/winthor/<TABELA>.yaml` por tabela descreve colunas, chave, watermark, partição e estratégia. Nenhum objeto é criado ou alterado no Oracle.
2. **Um extrator, dois modos de deploy**: gateway Windows no cliente (serviço, só conexões de saída, pull de jobs) e worker em modo direto na nossa infra. Mesmo código.
3. **Landing em Parquet** no object storage, um prefixo por tenant, com retenção configurável (padrão 90 dias). É a fonte de reprocessamento: recarregar o ClickHouse nunca exige voltar ao Oracle.
4. **ClickHouse em três camadas**: `raw` (espelho 1:1, colunas de controle `tenant_id`, `_batch_id`, `_loaded_at`, `_version`, `_is_deleted`; sem regra de negócio), `staging` (tipagem, renomeação, deduplicação) e `marts` (modelo de `docs/escopo-v1.md`). `staging` e `marts` são dbt Core.
5. **Sem CDC na v1**: full refresh para tabelas pequenas; watermark + janela móvel para as grandes; títulos em aberto sempre; deletes por reconciliação de chaves.
6. **Multi-tenant em tabelas compartilhadas**: `tenant_id` é o primeiro campo do `ORDER BY`, com row policies. O `tenant_id` vem do prefixo autenticado do objeto, nunca do conteúdo do arquivo. Teste automatizado de isolamento é obrigatório.
7. **Camada semântica (Cube Core) como única porta de acesso** para front e MCP. O MCP não expõe SQL livre.
8. **Orquestração mínima**: o control plane dispara o loader e o worker do dbt; sem Airflow ou Dagster na v1.

## Consequências

- O gateway precisa de apenas `SELECT` no Oracle, com guarda no código que recusa qualquer outra instrução.
- Toda regra de negócio fica no dbt (catálogo versionado de regras), o que facilita paridade com o TARGIT e auditoria.
- Reprocessar é barato (Parquet) e não impacta o ERP do cliente.
- Mudanças nesta arquitetura exigem novo ADR.
