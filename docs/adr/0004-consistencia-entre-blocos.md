# ADR 0004 — Consistência entre blocos paralelos de extração

- Status: **proposto** (2026-10-25); validar no E1.5 e E1.8

## Contexto

Tabelas grandes são extraídas em N blocos paralelos (faixas de data, de chave numérica ou de `ROWID`). Há duas formas de tratar a consistência entre blocos: fixar um SCN único (`AS OF SCN`) ou aceitar que cada bloco seja consistente apenas no seu próprio instante.

| | Sem SCN fixo (padrão) | `AS OF SCN` único |
|---|---|---|
| Privilégios | Só `SELECT` | `FLASHBACK` na tabela (ou `FLASHBACK ANY TABLE`) e `EXECUTE` em `DBMS_FLASHBACK` (ou leitura de `V$DATABASE`) |
| Undo necessário | Pela duração de cada bloco (minutos) | Desde o SCN inicial até o fim da extração (horas em carga inicial) |
| Risco de ORA-01555 | Baixo | Alto em ERP com `UNDO_RETENTION` curto |
| Linhas inseridas/alteradas durante a extração | Podem escapar nesta execução; a janela móvel da próxima execução as captura | Visão única |
| Deletes | Reconciliação de chaves (E1.8) | Idem |

## Decisão

1. **Padrão: sem SCN fixo.** Cada bloco é uma consulta independente, com faixa fechada à direita (`>= inicio AND < fim`), o que evita duplicidade nas fronteiras. A janela móvel de reprocessamento e a reconciliação de chaves cobrem alterações, inserções tardias e deletes.
2. **Opcional por tabela** (`consistency: scn` no YAML): todos os blocos com `AS OF SCN`, só quando o cliente conceder os privilégios. O `doctor` informa se estão disponíveis.
3. Em `ORA-01555`, o bloco é dividido ao meio e repetido (E1.5), em ambos os modos.
4. O manifest de cada bloco registra o instante de início e, quando disponível, o SCN.

## Consequências

- Nenhum privilégio além de `SELECT` é exigido do cliente para o caminho padrão.
- Os marts (fase 2) devem tolerar pequena defasagem entre tabelas lidas em momentos diferentes dentro de uma mesma execução; o frescor por tabela é exposto no `/status`.
