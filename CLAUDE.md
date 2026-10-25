# Plataforma de Dados Winthor — contexto do projeto

## O que estamos construindo

Plataforma multi-tenant que extrai dados do ERP **Winthor v30 (Oracle)** de ~100 clientes no primeiro ano, consolida em um DW (**ClickHouse**) e entrega:

1. **Dashboards e relatórios padrão** em front próprio, com envio agendado por e-mail e WhatsApp. Público leigo em dados.
2. **Acesso via Claude** (Claude Dashboards / chat) por um **servidor MCP remoto** que expõe a camada semântica. Público avançado, opcional.

Todos os clientes têm o mesmo schema Winthor v30 e usam o mesmo modelo semântico. **Não há customização por cliente na v1.** A configuração por cliente se limita a: conexão, janelas de carga, lista de filiais, parâmetros de regra (tabela `meta.tenant_param`) e destinatários de relatórios.

## Documentos de referência

- `docs/escopo-v1.md`: **contrato do modelo do DW** (dimensões, fatos, colunas, origens, regras), as análises por onda e os pontos a confirmar. A lista de tabelas Oracle de cada onda está na seção "Carga: janela e tabelas de origem".
- `PROMPT_FASE1.md`: fase em andamento. As demais fases serão detalhadas quando chegarmos nelas.

## Arquitetura (ELT)

```
Oracle/Winthor do cliente (tabelas PC*, somente leitura)
   │  Gateway (Windows Server, serviço, só conexões de saída)
   │  ou worker "modo direto" na nossa infra (VPN / IP liberado)
   ▼
Object storage: landing em Parquet, um prefixo por tenant (fonte de reprocessamento)
   │  Loader
   ▼
ClickHouse   raw (1:1 com a origem) → staging (tipado, sem duplicados) → marts (dimensões e fatos de docs/escopo-v1.md)
                                     └──────────────── dbt ────────────────┘
   ▼
Cube (camada semântica) — única porta de acesso aos dados
   ├── Front próprio + entrega agendada (e-mail, WhatsApp)
   └── Servidor MCP remoto ──► Claude

Control plane: tenants, jobs, watermarks, agendamentos, execuções do dbt, frescor dos dados, alertas
```

Camadas:

- **landing (Parquet):** permite recarregar o ClickHouse sem voltar ao Oracle do cliente. Retenção configurável (padrão proposto: 90 dias).
- **raw:** espelho 1:1 das tabelas de origem, só com as colunas usadas. Colunas de controle: `tenant_id`, `_batch_id`, `_loaded_at`, `_version`, `_is_deleted`. **Sem regra de negócio e sem renomear colunas.**
- **staging:** renomeia, tipa e deduplica (fica a última versão não deletada de cada linha). Trata colunas opcionais que variam entre clientes.
- **marts:** dimensões e fatos de `docs/escopo-v1.md`. Nenhum consumidor (Cube, MCP, front) lê `raw` ou `staging`.

## Decisões já tomadas (mudar só com ADR em `docs/adr/`)

1. **Linguagem:** Go para gateway, extractor, control plane e loader. dbt (SQL) para modelagem. Cube para a camada semântica. TypeScript para o servidor MCP (SDK oficial) e para o front (Next.js). Motivo do Go: binário único e boa concorrência. O gargalo esperado é Oracle e rede, não CPU. Revisitar com benchmark.
2. **Gateway: Windows Server 2016 ou mais novo, apenas na v1, instalado como serviço (background).** Go 1.21+ não roda em 2012 R2; esses clientes usam o modo direto. Linux fica fora de escopo por enquanto, mas a lógica de extração não pode depender de código específico de Windows. Só o wrapper de serviço, a criptografia de segredos (DPAPI) e o log de eventos são específicos de plataforma.
3. **Um extrator, dois modos de deploy:** gateway no cliente e worker na nossa infra usam o mesmo código.
4. **Gateway só faz conexões de saída e funciona por pull.** Ele busca jobs e configuração no control plane. Nunca abre porta de entrada.
5. **ELT a partir das tabelas do Winthor.** Extraímos as tabelas `PC*` 1:1 (só as colunas usadas), sem views e sem consultas prontas. Cada tabela tem um `sources/winthor/<TABELA>.yaml` (colunas, chave primária, watermark, partição, estratégia de carga), proposto por você e aprovado por mim. Consulta SQL de extração só como exceção, com ADR. Nunca crie nem altere objetos no Oracle.
6. **Sem CDC na v1.** Tabelas pequenas usam full refresh. As grandes usam incremental por watermark mais janela móvel de reprocessamento. Títulos em aberto (receber e pagar) entram sempre, fora da janela. Deletes são detectados por reconciliação de chaves.
7. **Multi-tenant em tabelas compartilhadas:** `tenant_id` é o primeiro campo do `ORDER BY`, com row policies no ClickHouse. O `tenant_id` vem do prefixo autenticado do objeto no storage, **nunca do conteúdo do arquivo**. A camada semântica injeta o filtro de tenant de forma obrigatória. Deve existir **teste automatizado de isolamento** (tenant A nunca enxerga dados do B).
8. **A camada semântica é a única porta de acesso aos dados** para front e MCP. O MCP **não** expõe SQL livre, só ferramentas de alto nível (listar métricas e dimensões, consultar métrica por dimensão, filtro e período). Somente leitura, com limite de linhas e timeout.
9. **Dados em repouso em região Brasil** (LGPD e latência): S3 e ClickHouse Cloud em AWS `sa-east-1`. Segredos em secret manager ou DPAPI, conforme o componente. Nenhuma credencial no repositório.
10. **Transformação:** dbt Core, executado em lote para os tenants que receberam dados novos (filtro por parâmetro), disparado pelo control plane depois do commit do loader. Parâmetros por tenant ficam na tabela `meta.tenant_param`, nunca em código por tenant. Plano inicial para os marts: dimensões reconstruídas por completo; fatos atualizados por apagar e reinserir uma janela móvel por tenant, sem `FINAL` nas consultas. **Isso será validado em um spike na fase 2 com o maior tenant.**
11. **Camada semântica: Cube Core autohospedado** (API, refresh worker e Cube Store). Sem Airflow ou Dagster na v1: o control plane dispara um worker que roda o dbt.

## Hospedagem e ambientes

| Ambiente | Onde | O que roda |
|---|---|---|
| dev | Servidor Windows Server 2019 com acesso de rede ao Oracle de teste (onde roda o Claude Code) | Código Go, gateway e testes de integração contra o Oracle. Dependências de dev na nuvem (ADR 0002): bucket S3 de dev e ClickHouse open-source em EC2, ambos em `sa-east-1`; Postgres nativo no Windows, em `D:`. Caches do Go e dados que crescem ficam em `D:\mint` (o `C:` tem pouco espaço) |
| staging e prod, dados | AWS `sa-east-1` | S3 (landing Parquet) e ClickHouse na mesma região. **Atenção:** o ClickHouse Cloud só oferece `sa-east-1` como private region do plano Enterprise; como hospedar o ClickHouse em produção é decidido antes da Fase 2 (ADR 0002) |
| staging e prod, serviços | Contêineres. Piloto: Railway. Produção: decidir em ADR antes do primeiro cliente real, entre Railway e AWS `sa-east-1` (ECS Fargate) | control plane e Postgres, loader, worker do dbt, Cube Core, MCP, front, delivery |

Por que assim: o Railway oferece quatro regiões (US West, US East, EU West e Sudeste Asiático), nenhuma no Brasil, e os buckets do Railway ficam fora da rede privada e sem backup automático. Por isso o dado em repouso fica na AWS `sa-east-1`. Os serviços sem estado trocam pouco dado e podem começar no Railway.

Regras de portabilidade (para poder migrar sem reescrever):

- Um `Dockerfile` por serviço, configuração só por variáveis de ambiente.
- Nenhum estado no disco do contêiner: o estado fica em Postgres, ClickHouse e S3.
- Nada específico do Railway (nem volumes, nem buckets). IaC em `/infra`.
- **O loader move o mínimo de dados entre regiões.** Prefira delegar a leitura do Parquet ao ClickHouse (função `s3()`). Se o loader precisar ler os arquivos, ele roda na AWS `sa-east-1`. Decisão e benchmark no ADR do E1.7.
- Não crie nem altere recursos de nuvem (serviço ClickHouse Cloud, buckets, contas) sem o meu OK. Proponha o IaC ou os comandos e eu aprovo.

## Regras de segurança com o Oracle (valem sempre)

- O banco usado em desenvolvimento é o **Oracle de teste da empresa**, acessado direto do servidor onde você roda. Não use Oracle em Docker.
- O usuário Oracle de desenvolvimento deve ter **só `SELECT`** nas tabelas necessárias, garantido no banco. A guarda no código é a segunda camada. O `gateway doctor` avisa se o usuário tiver privilégios além de leitura.
- O gateway executa apenas `SELECT` e `WITH`. Qualquer outro comando é recusado por uma guarda no código. Nunca rode DDL ou DML no Oracle.
- As credenciais ficam em `config.local.yaml` ou variáveis de ambiente (gitignored). **Nunca** peça senha no chat, grave em arquivo versionado ou escreva em log.
- O arquivo `TARGIT Enterprise - Winthor.improj` contém credenciais Oracle em texto claro. **Não o leia, não o copie e não o versione.** Se precisar de algo dele, eu forneço um extrato sem credenciais. Mantenha `.claude/settings.json` com regras de negação de leitura para `config.local.yaml`, `.env*` e `*.improj` (proponha e eu aprovo).
- **Dados pessoais** (CPF, CNPJ, nome, endereço, telefone, e-mail de clientes finais) não devem aparecer no seu contexto nem em logs. Para inspecionar dados, use contagens, agregações e colunas não pessoais. Quando eu precisar conferir amostras com acentos e nulos, escreva um script que **eu** executo.
- Testes unitários usam mocks da interface `Source`. Testes de integração rodam contra o Oracle de teste só quando `ORACLE_DSN` estiver definido.

## Roadmap e portões de verificação

Cada fase se divide em entregáveis pequenos, e cada entregável tem uma verificação concreta que eu executo. **Não avance para o próximo entregável sem o meu OK.** Só a fase 1 está detalhada (`PROMPT_FASE1.md`).

- **Fase 1 — Ingestão** (Oracle → Parquet → ClickHouse `raw`, gateway Windows), por onda. Verificação: contagens e somas batem entre Oracle e ClickHouse em todas as tabelas, incremental e deletes funcionam, o serviço Windows sobrevive a falhas e reboot.
- **Fase 2 — Transformação e camada semântica** (dbt `staging` e `marts` + Cube), onda 1, depois 2 e 3. Inclui o spike de atualização dos fatos e de joins grandes, `meta.tenant_param` e a foto diária de estoque. Verificação: testes dbt passam, os marts batem com os números de referência do TARGIT que eu forneço, uma chamada à API do Cube retorna o mesmo valor que o SQL de referência, e o teste de isolamento entre tenants passa.
- **Fase 3 — Servidor MCP remoto.** Verificação: o MCP Inspector lista as ferramentas, a autenticação por tenant funciona, o conector adicionado ao Claude responde a um conjunto de perguntas de referência com os valores corretos, e um Claude Dashboard é gerado a partir dele. As análises de `docs/escopo-v1.md` servem de base para essas perguntas.
- **Fase 4 — Front dos dashboards padrão** (os marcados com ★ em `docs/escopo-v1.md`). Verificação: login multi-tenant, um dashboard por domínio, os valores na tela batem com a camada semântica (teste automatizado) e o "atualizado em" aparece.
- **Fase 5 — Entrega agendada** (e-mail e WhatsApp). Verificação: e-mail de teste com o relatório recebido, WhatsApp com template aprovado entregue a um número de teste, agendamento por tenant executando sozinho, opt-in registrado.
- **Fase 6 — Operação e onboarding.** Inclui o ADR de hospedagem de produção. Verificação: um tenant novo é provisionado de ponta a ponta por um único comando ou fluxo, alertas disparam em falha simulada e o runbook cobre os incidentes comuns.

## Notas sobre o Winthor

- Os clientes podem diferir em versão do Oracle Database, charset (`NLS_CHARACTERSET`) e colunas extras. O gateway deve diagnosticar essas diferenças no `doctor` e falhar de forma clara.
- Mantenha LOBs e dados pessoais desnecessários fora do `raw`.
- **Nunca invente regra de negócio.** Definições como "faturamento" (cancelamentos, bonificações, devoluções, condição de venda) vêm de especialista de negócio. Os pontos a confirmar estão em `docs/escopo-v1.md`. Quando faltar a definição, registre a dúvida em `docs/open-questions.md` e siga com o que está definido.

## Como trabalhar neste repositório

- Você (Claude Code) roda no servidor de desenvolvimento. Comece cada fase em **plan mode**: apresente plano, riscos e perguntas, e implemente só depois da aprovação.
- Termine cada entregável com `docs/verification/<id>.md` (passo a passo exato e resultado esperado) e, quando der, `make verify E=<id>` com PASS/FAIL. Depois pare e aguarde meu OK.
- Decisões relevantes vão em ADR curto (`docs/adr/NNNN-titulo.md`). Resultados de benchmark vão em `docs/benchmarks.md`, com números reais.
- Código, identificadores, comentários e commits em **inglês**. Documentação, rótulos de UI e nomes de métricas de negócio em **pt-BR**.
- Logs estruturados. **Nunca** logar credenciais ou dados pessoais de clientes finais.
- Todo componente expõe métricas e healthcheck. O frescor dos dados ("atualizado em") é um dado de primeira classe.
- Comandos comuns no `Makefile` (`make build-windows`, `make test`, `make verify E=1.x`, `make doctor`, `make bench`, `make up`/`make down` para ligar e desligar as dependências de dev na nuvem). O servidor tem GNU make 4.4 (choco) e os recipes rodam no bash do Git for Windows. Mantenha este arquivo atualizado quando a estrutura mudar.

## Estrutura do repositório (monorepo)

```
/sources        configuração por tabela de origem: sources/winthor/<TABELA>.yaml e a lista de tabelas por onda
/gateway        binário Windows do cliente + biblioteca de extração (Go)
/control-plane  API, agendador de cargas, execuções do dbt, estado (Go + Postgres)
/loader         object storage → ClickHouse raw (Go)
/dbt            modelos staging e marts, testes, seeds
/semantic       Cube: métricas, dimensões, segurança por tenant
/mcp-server     MCP remoto + OAuth (TypeScript)
/web            front dos dashboards padrão (Next.js)
/delivery       renderização e envio agendado (e-mail, WhatsApp)
/infra          IaC (OpenTofu) dos ambientes de dev e produção em AWS sa-east-1
/docs           escopo-v1.md, adr/, verification/, open-questions.md, benchmarks.md
```