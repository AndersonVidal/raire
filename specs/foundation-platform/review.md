# foundation-platform — review

**Status:** revisão técnica e independente executada; aprovação formal pendente  
**Spec:** `spec.md`  
**Implementação:** commit local de 2026-11-20  
**Dono:** Victor  
**Revisores pendentes:** Anderson (produto/PII), Felipe (segurança/QA)  
**Checklist de spec:** `docs/status/2026-11-20-foundation-review-checklist.md`

## Veredito

`pronto para aprovação independente` — não aprovado. A implementação cobre a fatia de Victor e as dependências de `FND-09/10/17` usadas para testar a fundação. O código compila, passa `go vet` e a suíte automatizada. Uma revisão em contexto separado encontrou e ajudou a corrigir enumeração no signup, lost update de suspensão e atomicidade de recovery codes. Ainda faltam Postgres, worker de outbox, E2E, acessibilidade e aprovação humana.

## Gates

| Gate | Veredito | Revisor | Data |
| --- | --- | --- | --- |
| código | em review | Anderson/Felipe | 2026-11-20 |
| arquitetura | automatizado verde | teste `internal/architecture` | 2026-11-20 |
| segurança/privacidade | revisão interna feita; independente pendente | Felipe | 2026-11-20 |
| trust & safety | pendente | Felipe | — |
| QA | backend verde; E2E/a11y pendentes | Felipe | 2026-11-20 |
| produto | pendente | Anderson | — |

## 1. Escopo revisado

- `cmd/raire`, `internal/**`, `db/**`, `api/openapi.yaml`;
- migrations, seed, auth, catálogo, perfil, admin, auditoria, outbox, flags;
- testes unitários, integração HTTP com SQLite real, contrato OpenAPI e arquitetura;
- documentação de operação e evidências.

## 2. Achados corrigidos durante a implementação

| # | Achado | Correção |
| --- | --- | --- |
| 1 | DSN SQLite com caminho absoluto do Windows falhava (`out of memory`) | `file:///C:/...` normalizado em `platform/db` |
| 2 | Segredo TOTP era gravado em texto | AES-256-GCM com prefixo `enc:v1:`; teste garante que o segredo não aparece no banco |
| 3 | Criação de perfil promovia papel depois do commit e podia ficar inconsistente | promoção feita antes do insert; retry do POST é recuperável |
| 4 | Suspensão revogava sessão em transação separada | suspensão, auditoria, outbox e revogação na mesma transação |
| 5 | API não impedia cache de resposta autenticada | `Cache-Control: no-store` em `/v1/` |
| 6 | Recovery codes podiam ser reutilizados em corrida | tabela `mfa_recovery_codes` com consumo condicional no login |
| 7 | E-mail duplicado precisava de resposta genérica | status, corpo e headers de produção idênticos; token só em dev para conta criada; teste compara as duas respostas |
| 8 | Teste de arquitetura não existia | regra automatizada para `platform`, `domain`, `contracts`, adapters e `api` |
| 9 | Perfil podia sobrescrever suspensão concorrente | coluna `version` + update condicional; `ErrConflict` vira 409; teste de lost update |
| 10 | Token de e-mail não verificava `RowsAffected` | confirmação/reset retornam `TOKEN_INVALID` se o token já foi consumido |
| 11 | Telefone re-salvo perdia verificação | só limpa `contact_phone_verified_at` quando o número muda |
| 12 | Texto público aceitava PII | validação de bio, disponibilidade, cidade, bairro e nome com `PII_NOT_ALLOWED` |
| 13 | API acessava repositório direto | consultas de auditoria/flags passaram pelo `App`; API não usa mais `UoW` |
| 14 | Default de ambiente era dev | sem `RAIRE_ENV`, assume produção e exige chave; token de dev só com opt-in |

## 3. Achados abertos

| # | Severidade | Achado | Follow-up |
| --- | --- | --- | --- |
| 1 | Alta | Sem adaptador PostgreSQL/PostGIS; SQLite é o único banco | FUP-01 |
| 2 | Alta | Outbox sem worker; eventos ficam pendentes | FUP-02 |
| 3 | Alta | Login não tem CSRF/origin check | FUP-03 |
| 4 | Média | Rate limit e métricas são por instância | FUP-04 |
| 5 | Média | Sem enrollment/rotação de MFA pela API | FUP-05 |
| 6 | Média | Sem fluxo de verificação de telefone; selo não é emitido | FUP-09 |
| 7 | Média | Flags ainda não gateiam features | FUP-06 |
| 8 | Média | `evidence` administrativo pode receber PII | FUP-08 |
| 9 | Baixa | Limpeza de chaves de idempotência/tokens expirados | FUP-07 |
| 10 | Baixa | `LastSeenAt` atualiza no máximo a cada 5 min, mas ainda gera escrita | monitorar no CI/staging |

## 4. Aderência à spec

- Conta, sessão, CSRF, MFA, rate limit, confirmação, reset e preferência de contato: implementados.
- Catálogo, sinônimos, categoria inativa e auditoria: implementados.
- Perfil, área, disponibilidade, publicação, pausa, retomada, suspensão e categoria inativa: implementados.
- Localização precisa cifrada, retenção e expurgo auditável: implementados.
- API pública sem PII: testada.
- Portfólio, exportação/exclusão, direitos, web, a11y e E2E: fora desta fatia.

## 5. Evidências

- `specs/foundation-platform/evidence/2026-11-20-foundation-tests.md`
- `docs/status/2026-11-20-foundation-implementation.md`
- `docs/status/2026-11-20-foundation-follow-ups.md`

## 6. O que o humano deve olhar

- O veredito de segurança e QA independentes.
- O risco de manter SQLite até o adaptador Postgres existir.
- O login CSRF e o rate limit compartilhado antes de qualquer staging.
- A decisão de aceitar flags sem gating até FND-10.
- Nenhuma aprovação de piloto ou produção está implícita nesta revisão.
