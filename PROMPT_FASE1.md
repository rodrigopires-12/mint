# Fase 1 — Ingestão: Oracle/Winthor → Parquet → ClickHouse (raw)

Leia o `CLAUDE.md` e o `docs/escopo-v1.md` antes de qualquer coisa. O primeiro traz o contexto, a arquitetura, as decisões já tomadas e as regras de segurança com o Oracle. O segundo traz o modelo do DW e, na seção "Carga: janela e tabelas de origem", as tabelas Oracle de cada onda.

## Objetivo

Construir o caminho de ingestão de ponta a ponta: um **gateway Windows**, rodando como serviço em background, que extrai as **tabelas de origem do Winthor** (1:1, só as colunas usadas), grava Parquet em object storage e alimenta tabelas `raw` no ClickHouse. O mesmo código também precisa rodar como worker no modo direto, na nossa infra.

Nesta fase **não há transformação de negócio**. A camada `raw` é um espelho da origem. Staging e marts são da fase 2.

**Performance do gateway é o requisito crítico.** Isso inclui throughput de extração e também impacto mínimo no Oracle de produção do cliente, que roda o ERP.

## Premissas desta fase

- Você roda no **servidor de desenvolvimento**, que tem acesso de rede ao Oracle de teste da empresa. Sem Oracle em Docker.
- As tabelas a extrair são as listadas no `docs/escopo-v1.md`, por onda (cerca de 45). **As colunas extraídas são as citadas no escopo, mais chaves e coluna de watermark.** Qualquer coluna adicional precisa da minha aprovação. Onde o escopo citar a coluna de origem, use-a. Onde não citar, proponha e registre em `docs/open-questions.md`. Coluna do escopo que não existir na base também vai para lá.
- Cada tabela tem um `sources/winthor/<TABELA>.yaml` (`status: proposed` ou `approved`). Você propõe, eu aprovo. Nunca crie nem altere objetos no Oracle.
- Credenciais só em `config.local.yaml` ou variáveis de ambiente (gitignored). Nunca peça a senha no chat, e não leia o arquivo `TARGIT Enterprise - Winthor.improj` (ver `CLAUDE.md`).
- O usuário Oracle de desenvolvimento tem só `SELECT`. O gateway é somente leitura (só `SELECT` e `WITH`), com guarda no código.
- **Dados pessoais não passam pelo seu contexto.** Use contagens e agregações. Para conferência de amostras com acentos e nulos, escreva um script que eu executo.
- Dependências de dev (ClickHouse, MinIO ou S3, Postgres): ver "Hospedagem e ambientes" no `CLAUDE.md`. Não crie recursos de nuvem sem o meu OK.
- O gateway é só Windows Server por enquanto. O serviço é testado em um Windows Server limpo (pode ser uma VM), e não no servidor de desenvolvimento.
- **Ordem por onda:** o inventário cobre todas as tabelas. Do E1.4 ao E1.9 você prova o mecanismo em **3 tabelas que eu escolho**. Minha sugestão: `PCFILIAL` (pequena, full refresh), `PCMOV` (grande, incremental) e `PCPREST` (regra de títulos em aberto). Só no E1.10 as demais entram, onda por onda.

## Como trabalhar: portões de verificação

Cada entregável abaixo é um portão. Ao terminar um entregável:

1. Escreva `docs/verification/E1.x.md` com o passo a passo exato para eu verificar e o resultado esperado.
2. Crie `make verify E=1.x` automatizando o que der (saída PASS/FAIL).
3. Resuma em poucas linhas: o que foi feito, como verificar e pendências.
4. **Pare e aguarde meu OK.** Se uma verificação falhar, corrija antes de seguir.

## Antes de codar (plan mode)

Apresente um plano com riscos e perguntas. Eu quero ver respondido:

- **Ambiente do servidor de desenvolvimento:** sistema operacional (Windows ou Linux), se há Docker, se há `make`, e se há saída para a internet (módulos Go, GitHub, ClickHouse e S3 de dev). Mostre o impacto em Makefile ou scripts alternativos, e onde sobem ClickHouse, MinIO e Postgres de dev.
- **Versão do Go vs. Windows Server:** Go 1.21 e posteriores não rodam em Windows Server 2012/2012 R2 (o Go 1.20 foi o último). Pergunte qual é o Windows Server mais antigo que precisamos suportar e mostre o impacto na escolha da versão do Go.
- **Driver Oracle:** implemente `Source` como interface e avalie **godror** (ODPI-C, exige cgo, Oracle Instant Client e VC++ Redistributable no servidor do cliente) e **go-ora** (Go puro, sem dependências nativas). Meça os dois no `bench` e decida por throughput e simplicidade de instalação. Registre em ADR. Considere a compatibilidade com as versões do Oracle Database da base instalada (confirmar comigo).
- **Extração de tabelas grandes:** como particionar em blocos paralelos (por coluna de data, por faixa numérica da chave, ou por faixas de `ROWID`, já que agora lemos tabelas e não consultas com joins). Diga quais privilégios cada opção exige (por exemplo, acesso às views de extents para `ROWID`).
- **Consistência entre blocos:** ler todos os blocos no mesmo SCN (`AS OF SCN`, que exige privilégio de flashback) ou aceitar leituras em momentos diferentes e confiar na janela móvel de reprocessamento. Mostre o risco de `ORA-01555` em cada caso. Registre em ADR.
- **Watermark em tabelas do Winthor:** muitas tabelas não têm timestamp de última alteração confiável. Proponha, por tabela, a coluna candidata (data de negócio) e como validar se ela muda quando a linha é alterada. **Não invente.** O que não der para validar vira pergunta.
- **Instalação:** como será o pacote (zip com `install.ps1` na v1; MSI pode ficar para depois).

## Entregáveis

### E1.1 — Base do projeto e `gateway doctor`

**Entrega:**
- Estrutura do monorepo, Makefile (ou a alternativa aprovada), build para Windows (`make build-windows` gera `gateway.exe`), CI que compila em runner Windows, carregamento de configuração.
- `sources/winthor/_waves.yaml` com as tabelas de cada onda, extraídas do `docs/escopo-v1.md`.
- `.claude/settings.json` com regras de negação de leitura para arquivos de credenciais (proposto por você, aprovado por mim).
- ADR 0001 (arquitetura ELT e camadas) e ADR 0002 (hospedagem e ambientes), a partir do `CLAUDE.md`.
- `gateway doctor`: conecta ao Oracle de teste e mostra versão do banco, usuário, `NLS_CHARACTERSET` e relógio. Verifica os **privilégios** do usuário (por `SESSION_PRIVS`, `USER_ROLE_PRIVS` e `USER_TAB_PRIVS`) e avisa se houver qualquer coisa além de leitura. Para cada tabela de `_waves.yaml`, confirma que existe e que dá para ler (`SELECT` com zero linhas), e lista as que faltam.
- A guarda somente-leitura já fica ativa aqui.

**Como eu verifico:**
- `doctor` no ambiente de desenvolvimento mostra OK ou FAIL por verificação, e o número de tabelas checadas é igual ao de `_waves.yaml`.
- Com a senha errada, a mensagem é clara, sem stack trace e sem a senha na saída.
- Um arquivo de teste com `UPDATE` ou `DELETE` é recusado pela guarda.
- Se eu usar um usuário com privilégio extra no ambiente de teste, o `doctor` avisa.
- `make build-windows` gera o `gateway.exe` e o CI compila em runner Windows.

### E1.2 — Inventário das tabelas

**Entrega:** `docs/tables-inventory.md` e um `sources/winthor/<TABELA>.yaml` (`status: proposed`) para **cada tabela de todas as ondas**, com:
- Colunas extraídas, com o tipo Oracle e o tipo ClickHouse proposto.
- Chave primária real (pelas constraints do dicionário, confirmando unicidade) ou, se não houver, a chave candidata.
- Coluna de watermark candidata, coluna de partição candidata e a estratégia proposta (`full`, `watermark`, `window` ou `open_items`).
- Linhas e tamanho (do dicionário, se o privilégio permitir; senão `COUNT(*)`), e a distribuição por mês quando houver coluna de data.
- Riscos: LOB, `LONG`, tabela larga, sem chave, linhas alteradas retroativamente.
- Colunas do escopo que não existem na base, e perguntas.
- Respostas, com evidência, aos pontos do `docs/escopo-v1.md` que dependem só do dicionário do Oracle: **#1** (`PCDATAS`), **#2** (chaves reais), **#11** (rateio de centro de custo), **#12** (tipos reais), **#13** (`PCNFENT` ou `PCMOV` trazem o pedido de compra?) e **#14** (`DTVENC`). Os demais pontos dependem de especialista de negócio e ficam abertos em `docs/open-questions.md`.

O `doctor` passa a avisar sobre tabelas sem arquivo aprovado.

**Como eu verifico:**
- Há uma entrada por tabela (nenhuma faltando).
- Eu reviso chave, watermark, partição e estratégia de cada uma, corrijo o que for preciso e mudo o status para `approved`. É um portão humano.
- `doctor` mostra todas as tabelas como aprovadas depois da minha revisão.
- Os pontos #1, #2, #11, #12, #13 e #14 têm resposta com evidência.

### E1.3 — Gateway como serviço Windows (esqueleto)

**Entrega:** `gateway.exe service install|uninstall|start|stop|status` e `gateway.exe run` (modo console com a mesma lógica). O serviço ainda não extrai nada: ele grava um heartbeat em log a cada 10 segundos. Requisitos:
- Início automático (atrasado) e recuperação automática em caso de falha.
- Parada graciosa, com timeout.
- Instância única.
- Log em arquivo com rotação (em `%ProgramData%`) e eventos de início, parada e erro no Windows Event Log.
- Segredos criptografados com DPAPI (`gateway.exe secret set`), nunca em texto puro no YAML.
- Suporte a proxy HTTP(S) configurável.
- Pacote zip com `gateway.exe`, `config.example.yaml` e `install.ps1`.

**Como eu verifico (em um Windows Server limpo):**
- `install.ps1` instala. O serviço aparece em `services.msc` e `sc query` mostra RUNNING.
- O arquivo de log recebe os heartbeats e o Event Log registra início e parada.
- Encerrar o processo pelo Gerenciador de Tarefas faz o serviço voltar sozinho.
- Reiniciar o servidor faz o serviço subir sozinho.
- `uninstall` remove o serviço e deixa só os logs.
- O `config.yaml` não contém nenhuma senha legível.

### E1.4 — Extração para Parquet local (correção antes de velocidade)

**Entrega:** `gateway extract <TABELA> --out <dir>` em fluxo único, convertendo para Arrow e gravando Parquet (zstd) mais um manifest. Mapeamento de tipos documentado (`NUMBER` com e sem precisão, `DATE` com hora, tratamento de `CHAR` com espaços, conversão de charset para UTF-8), com `NULL` preservado. A `raw` mantém o valor como está na origem, exceto pela conversão de charset. `gateway check <TABELA>` compara contagem de linhas e somas de colunas numéricas entre Oracle e Parquet.

**Como eu verifico (nas 3 tabelas que eu escolhi):**
- `gateway check` retorna PASS.
- Com DuckDB, `SELECT count(*)` no Parquet é igual ao `SELECT COUNT(*)` da tabela no Oracle.
- Um script que eu executo imprime uma amostra de 20 linhas com acentos e valores nulos dos dois lados, e eu confiro que são iguais.

### E1.5 — Performance: blocos paralelos e pipeline em streaming

**Entrega:**
- **Particionamento em blocos** conforme a decisão do plan mode e o arquivo da tabela (data, faixa numérica da chave ou faixas de `ROWID`). Sem coluna de partição, usa fluxo único e avisa.
- **N configurável**, com teto global de sessões no Oracle.
- **Fetch em lote** com array size e prefetch ajustáveis, convertido direto para Arrow, sem passar por CSV ou JSON.
- **Pipeline** leitura → encoding → gravação em paralelo, com backpressure e memória limitada (padrão ≤ 1 GiB de RSS, independente do tamanho da tabela).
- **Proteção do ERP:** limite de sessões, limite de MB/s, janela de horário permitida e pausa adaptativa quando as leituras ficarem lentas.
- **Resiliência:** em `ORA-01555` (snapshot too old), reduzir o bloco e tentar de novo. Reconectar em queda de rede. Timeouts em tudo.
- `gateway bench <TABELA>` com um baseline "só fetch" (loop de leitura sem encoding nem upload).

**Como eu verifico (na tabela grande):**
- `docs/benchmarks.md` traz rows/s, MB/s, CPU, memória e gargalo identificado para N = 1, 2, 4 e 8.
- O pipeline completo atinge pelo menos **80% do baseline "só fetch"**. Se não atingir, vem uma justificativa e um plano de otimização. Se o encoding em Go for o limitante mesmo depois de otimizado, isso fica registrado para avaliarmos uma alternativa.
- `gateway check` retorna PASS com N = 1 e com N = 8, com as mesmas contagens e somas (sem linhas duplicadas ou perdidas nas fronteiras entre blocos).
- Com o teto em 2 sessões, `v$session` no Oracle nunca mostra mais de 2 sessões do usuário do gateway.
- O RSS fica dentro do limite durante a maior extração.

### E1.6 — Upload para object storage (idempotente e retomável)

**Entrega:** interface `ObjectStore`. Nesta fase vale o modo direto, com credenciais S3/MinIO; as URLs pré-assinadas entram no E1.9. Upload multipart com retry e backoff. Manifest por bloco (linhas, checksum, faixa, tempo). Chaves de objeto determinísticas, no formato `tenant=<id>/table=<TABELA>/batch=<id>/part-<n>.parquet`. Retenção configurável.

**Como eu verifico:**
- Após uma extração, o console do MinIO mostra os blocos e os manifests.
- Rodar de novo gera as mesmas chaves, sem duplicar objetos.
- Matar o processo no meio e reexecutar completa só o que faltava, e o log mostra quais blocos foram pulados.
- A soma das linhas dos manifests é igual ao `COUNT(*)` no Oracle.

### E1.7 — Loader: camada raw no ClickHouse

**Entrega:**
- DDL gerado a partir do arquivo aprovado da tabela: `raw.<tabela>` com `ReplacingMergeTree` (coluna de versão e flag de deleção), `tenant_id` como primeiro campo do `ORDER BY`, mais `_batch_id`, `_loaded_at`, `_version` e `_is_deleted`. A partição por mês, quando fizer sentido, é decidida por tabela com base no inventário.
- `loader run <TABELA>` que carrega os Parquets do storage e registra o commit. Permite reprocessar uma janela sem afetar outros dados.
- O `tenant_id` vem do **prefixo do objeto**, nunca do conteúdo do arquivo.
- **Duas abordagens, com benchmark:** (a) delegar a leitura ao ClickHouse (`INSERT ... SELECT` com a função `s3()`) e (b) o loader ler o Parquet e inserir. Registre a decisão em ADR, considerando que em produção o loader deve mover o mínimo de dados entre regiões (ver `CLAUDE.md`).

**Como eu verifico:**
- `SELECT count() FROM raw.<tabela> FINAL WHERE tenant_id = 'piloto'` é igual ao `COUNT(*)` no Oracle.
- `SHOW CREATE TABLE` mostra a chave e os tipos que eu aprovei.
- Um script que eu executo mostra uma amostra de 20 linhas, com acentos, correta.
- Rodar o loader duas vezes não duplica dados (com `FINAL`).
- Um arquivo Parquet com um `tenant_id` diferente no conteúdo é carregado com o `tenant_id` do prefixo.

### E1.8 — Incremental e deletes

**Entrega:**
- Estratégias por tabela: `full`, `watermark`, `window` e `open_items` (títulos em aberto sempre entram, fora da janela; a definição por tabela vem do inventário e é aprovada por mim).
- `StateStore` como interface (SQLite local por enquanto; o control plane assume no E1.9).
- O watermark **só avança depois que o loader confirma o commit**.
- Reconciliação de chaves para detectar deletes, marcando as linhas removidas na flag de deleção.
- Relatório de reconciliação por execução (contagem e somas, Oracle vs ClickHouse).
- `PCEST` (posição atual de estoque) entra como `full`. A necessidade da foto diária é registrada em `docs/open-questions.md` para a fase 2.

**Como eu verifico (eu altero o Oracle de teste manualmente, seguindo `docs/verification/E1.8.md`):**
- Uma segunda execução sem mudanças no banco extrai 0 linhas novas (ou só a janela móvel).
- Alterar 1 linha: depois do ciclo, o ClickHouse reflete o novo valor.
- Inserir 1 linha: ela aparece.
- Apagar 1 linha: depois da reconciliação, ela fica marcada como deletada e some das consultas normais.
- Matar o processo no meio de uma carga e reiniciar: sem perda, sem duplicação e sem watermark adiantado.
- O relatório de reconciliação mostra 0 divergências.

### E1.9 — Control plane mínimo e modo gateway completo

**Entrega:**
- API em Go com Postgres: tenants, conexões, tabelas habilitadas, agendamentos, execuções, blocos e watermarks (substitui o SQLite).
- O serviço Windows faz polling de jobs (só saída) e envia heartbeat e progresso.
- Upload por **URLs pré-assinadas** com escopo por tenant e TTL curto. O gateway nunca recebe credenciais do storage.
- Autenticação gateway ↔ control plane por token com rotação e revogação.
- O loader é disparado pelo control plane quando o manifest chega.
- Endpoint `/status` com o frescor por tenant e por tabela, e alertas simples (carga falhou ou atrasou, heartbeat ausente).
- O mesmo código roda como worker no modo direto, selecionado por flag.

**Como eu verifico:**
- O serviço Windows registrado no control plane aparece no `/status` com heartbeat de menos de 60 segundos.
- Agendo uma tabela (cron) e ela roda sozinha, sem eu tocar em nada, e os dados sobem.
- O `/status` mostra o "atualizado em" de cada tabela.
- Desligar a rede do gateway dispara o alerta de heartbeat ausente. Ao religar, ele retoma sozinho.
- A configuração do gateway não contém nenhuma credencial do storage.
- Revogar o token faz o gateway ser recusado.

### E1.10 — Carga completa por onda (três portões: 10a, 10b e 10c)

**Entrega:** para cada onda, nessa ordem (**10a:** onda 1, **10b:** onda 2, **10c:** onda 3), carga inicial de todas as tabelas da onda com o histórico desde 01/jan do ano anterior (configurável por tenant), seguida de um ciclo incremental. Cada portão traz um relatório de reconciliação Oracle vs ClickHouse por tabela e o registro do tempo e do impacto no Oracle em `docs/benchmarks.md`.

**Como eu verifico (a cada onda):**
- O relatório cobre todas as tabelas da onda, sem divergência de contagem ou de soma.
- O teto de sessões e o limite de MB/s configurados foram respeitados durante a carga (conferido em `v$session`).
- Um ciclo incremental logo depois extrai só a janela móvel e continua sem divergência.
- Só passo para a onda seguinte depois do meu OK.

### E1.11 — Fechamento da fase

**Entrega:**
- Relatório final de reconciliação Oracle vs ClickHouse para todas as tabelas de todas as ondas.
- Teste automatizado de isolamento entre tenants (dois tenants fictícios, um não enxerga o outro).
- Guia de instalação e operação do gateway em Windows Server, runbook de falhas comuns, ADRs e `open-questions.md` atualizados.

**Como eu verifico:**
- `make verify-phase1` retorna PASS.
- Alguém que não participou do projeto instala o gateway em um Windows Server limpo seguindo só o guia, e a primeira carga funciona.
- O relatório de reconciliação cobre todas as tabelas, sem divergência.
- O teste de isolamento passa.

## Fora de escopo nesta fase

Staging e marts (dbt), camada semântica (Cube), MCP, front, entrega por e-mail e WhatsApp, foto diária de estoque, CDC em tempo real, gateway em Linux, auto-atualização do gateway, MSI.