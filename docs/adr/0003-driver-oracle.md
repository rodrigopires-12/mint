# ADR 0003 — Driver Oracle do extrator

- Status: **proposto** (2026-10-25); confirmar com o benchmark do E1.5
- Decisão provisória: **go-ora** como driver padrão; **godror** como build alternativa

## Contexto

O gateway roda em Windows Server 2016+ de ~100 clientes, com Oracle Database **11gR2, 12c/18c e 19c** na base instalada. Performance e simplicidade de instalação são requisitos críticos. Dois drivers Go foram considerados:

| | go-ora (`github.com/sijms/go-ora/v2`) | godror (`github.com/godror/godror`) |
|---|---|---|
| Implementação | Go puro (protocolo TNS) | cgo sobre ODPI-C |
| Dependências no cliente | Nenhuma: binário único | Oracle Instant Client + VC++ Redistributable |
| Compilação | `CGO_ENABLED=0`, cross-compile trivial | Compilador C; cross-compile difícil |
| Oracle suportado | 10g em diante (protocolo) | Depende do Instant Client: o 23 conecta só em 19c+; para 11g/12c é preciso o 19 |
| Maturidade | Menor; conversões de charset e `NUMBER` precisam de validação | Alta (mesma base do cx_Oracle/node-oracledb) |
| Licença | MIT | Apache-2.0 (ODPI-C: Apache/UPL) |

## Decisão provisória

`Source` é uma interface; as duas implementações existem no repositório (`internal/source/goora` padrão; `internal/source/godror` atrás da build tag `godror`, a partir do E1.5). O E1.5 mede ambas sobre `PCMOV` com array size de 1k a 10k linhas (rows/s, MB/s, CPU, RSS). go-ora fica como padrão **salvo** deficit de throughput acima de ~30% ou falha de fidelidade (charset, `NUMBER` sem escala, `DATE`) detectada pelo `gateway check` do E1.4 e pelos scripts de amostra.

## Consequências

- Pacote de instalação sem dependências nativas e sem Instant Client (se go-ora prevalecer).
- Se godror prevalecer, o pacote inclui o Instant Client 19 Basic Light (~80 MB) e exige VC++ 2017; a versão do Instant Client passa a ser diagnosticada pelo `doctor`.
- O servidor de dev não tem compilador C; o bench do godror depende de instalar mingw-w64 (choco, elevado) ou zig portátil.
- SQL do gateway compatível com 11g: `ROWNUM` em vez de `FETCH FIRST`.
