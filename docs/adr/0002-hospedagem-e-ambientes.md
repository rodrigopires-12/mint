# ADR 0002 — Hospedagem e ambientes

- Status: aceito para dev e piloto (2026-10-25); **produção pendente** (ver "Alerta")
- Origem: `CLAUDE.md`, seção "Hospedagem e ambientes", e plan mode do E1.1

## Contexto

Dados em repouso precisam ficar no Brasil (LGPD e latência). O Railway, escolhido para o piloto dos serviços sem estado, não tem região no Brasil nem buckets com backup. O servidor de desenvolvimento é um Windows Server 2019 físico sem Linux local (Docker só com contêineres Windows, sem WSL2, Hyper-V não instalado).

## Decisão

| Ambiente | Dados | Serviços |
|---|---|---|
| dev | Bucket S3 `sa-east-1` (landing). **ClickHouse open-source em uma EC2 `sa-east-1`** (ex.: `m6i.large`, gp3, desligada quando ociosa). Postgres nativo no Windows do servidor de dev, em `D:` | Gateway e worker rodam no servidor de dev (Windows). Windows Server limpo para testar o serviço: EC2 Windows em `sa-east-1`, desligada quando ociosa |
| staging / prod | S3 `sa-east-1`. ClickHouse: **decisão pendente** (abaixo) | Contêineres. Piloto no Railway; produção decidida em ADR antes do primeiro cliente real (Railway vs ECS Fargate `sa-east-1`) |

Regras de portabilidade (inalteradas): um `Dockerfile` por serviço; configuração só por variáveis de ambiente; nenhum estado em disco de contêiner; nada específico do Railway; IaC em `/infra` (OpenTofu, estado local em dev, gitignored). O loader move o mínimo de dados entre regiões: preferir `INSERT ... SELECT s3()` no ClickHouse (benchmark no E1.7).

Nenhum recurso de nuvem é criado sem aprovação explícita: o IaC é proposto e aplicado só depois do OK.

## Alerta: ClickHouse Cloud não oferece `sa-east-1` pública

Em 2026-10-25 a documentação do ClickHouse Cloud lista `sa-east-1` (São Paulo) apenas como **private region**, disponível só no plano **Enterprise**. A decisão 9 do `CLAUDE.md` (ClickHouse Cloud em `sa-east-1`) não é executável nos planos Basic e Scale. Opções para staging/prod, a decidir antes da Fase 2:

1. ClickHouse Cloud Enterprise com private region `sa-east-1` (custo e contrato a negociar).
2. ClickHouse open-source autogerido em `sa-east-1` (EC2 ou EKS), com backups para S3. É o que dev usa; operação por nossa conta.
3. ClickHouse Cloud em `us-east-1` com parecer jurídico LGPD (dados fora do Brasil): descartado por padrão.

## Por que dev na nuvem e não local

O MinIO foi arquivado em 2025 (sem binários desde outubro). O servidor de dev não executa contêineres Linux. Instalar Hyper-V exigiria reboot de um servidor compartilhado. Usar S3 real e ClickHouse na mesma região reproduz o caminho de produção (função `s3()` sem egress) e mantém os dados do Oracle de teste no Brasil.

## Consequências

- `make up`/`make down` ligam e desligam a EC2 do ClickHouse e o serviço Postgres local, para controlar custo.
- A verificação do E1.6 usa `aws s3 ls` no lugar do console do MinIO.
- O `CLAUDE.md` é atualizado para refletir este ADR.
