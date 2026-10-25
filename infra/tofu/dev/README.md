# Infra de desenvolvimento (AWS `sa-east-1`)

Recursos planejados para o ambiente de dev (ADR 0002). **Nada é criado sem aprovação explícita**: o código HCL (OpenTofu) é proposto no entregável que precisa do recurso, o `plan` é revisado e só então o `apply` roda.

| Recurso | Entregável | Observações |
|---|---|---|
| Bucket S3 de landing (Parquet), versionamento desligado, ciclo de vida de 90 dias, bloqueio de acesso público, SSE-S3 | E1.6 | Prefixo por tenant: `tenant=<id>/table=<TABELA>/batch=<id>/` |
| Usuário ou papel IAM do loader e do gateway em modo direto (mínimo privilégio) | E1.6 | URLs pré-assinadas chegam no E1.9 |
| EC2 com ClickHouse open-source (`m6i.large`, gp3 200 GB, security group restrito ao IP do servidor de dev), parada quando ociosa | E1.7 | Substitui o ClickHouse Cloud, que só oferece `sa-east-1` no plano Enterprise |
| EC2 Windows Server 2019/2022 limpa para testar o serviço do gateway, parada quando ociosa | E1.3 | Acesso RDP só do IP do servidor de dev |

Estado do OpenTofu em dev: local, em `D:\mint\tfstate` (gitignored). Binário do OpenTofu: portátil em `D:\tools` (sem elevação).
