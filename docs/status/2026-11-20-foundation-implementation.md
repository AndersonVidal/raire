# Implementação da Fundação — 2026-11-20

**Status:** fatia de Victor implementada e verificada localmente; revisão independente pendente  
**Autor da implementação:** Victor (agente no harness)  
**Aprovação:** nenhuma aprovação de código foi registrada por IA  
**Escopo:** `FND-01`..`FND-06`, `FND-14`..`FND-16`; dependências antecipadas de `FND-09`, `FND-10` e `FND-17`  
**Fora de escopo:** portfólio, privacidade/direitos, web, E2E, Postgres de produção, piloto

---

## 1. Resumo

A fundação saiu do papel. O repositório agora tem um monólito modular em Go com:

- SQLite local/CI, migrations forward-only e seed idempotente;
- contas, Argon2id, sessão por cookie, CSRF, rate limit e MFA TOTP;
- confirmação de e-mail, reset de senha e códigos de recuperação;
- catálogo curado com sinônimos e administração por operador;
- perfil profissional com área, disponibilidade, contato e localização precisa cifrada;
- suspensão/reativação com auditoria append-only e revogação de sessão;
- feature flags, health/readiness, logs JSON sem PII e expurgo de localização;
- OpenAPI, runbook e testes automatizados.

**Nada foi implantado.** Não há deploy, Postgres, web, dados reais ou aprovação de piloto.

---

## 2. Entregas por ticket

| Ticket | Entrega | Estado |
| --- | --- | --- |
| FND-01 | `cmd/raire`, contextos `identity/catalog/directory/platform/api/app`, teste de arquitetura, health | em review |
| FND-02 | migrations `0001`, seed `0002`, tabelas de conta/sessão/perfil/catálogo/auditoria/outbox/flags/idempotência | em review |
| FND-03 | signup, login, logout, sessão, RBAC, CSRF, rate limit, Argon2id | em review |
| FND-04 | confirmação, reenvio, reset, TOTP, recovery codes, MFA cifrado | em review |
| FND-05 | catálogo, sinônimos, CRUD admin, desativação, auditoria | em review |
| FND-06 | perfil, área, disponibilidade, contato, publicação/pausa/retomada, suspensão | em review |
| FND-09 (antecipado) | listagem admin com cursor/filtros, suspensão/reativação, auditoria | em review por Felipe |
| FND-10 (parcial) | health/readiness, logs JSON, flags, métricas em memória | em andamento por Felipe |
| FND-14 | OpenAPI 3.1, README, runbook, handoff | em review |
| FND-15 | troca de senha, CSRF, rate limits, idempotência, recovery codes | em review |
| FND-16 | precedência de contato, campo de verificação de telefone (fluxo OTP pendente), janelas, AES-256-GCM, expurgo | em review |
| FND-17 (parcial) | categorias, flags, paginação/filtros de auditoria | em andamento por Felipe |

---

## 3. Estrutura criada

```txt
cmd/raire/                     binário API + job de expurgo
db/migrations/                 schema forward-only
db/seeds/                      catálogo e flags idempotentes
internal/identity/             conta, sessão, consentimento, MFA
internal/catalog/              categorias, sinônimos
internal/directory/            perfil, área, disponibilidade, contato, localização
internal/platform/             config, db, http, audit, outbox, flags, segurança, rate limit
internal/contracts/            erros estáveis
internal/api/                  handlers JSON /v1
internal/app/                  composição de dependências
api/openapi.yaml               contrato da fundação
docs/runbooks/foundation.md    operação local
```

Regras de dependência são verificadas por `internal/architecture`:

- `platform` não importa contexto de negócio;
- `domain` não importa HTTP, banco ou adaptadores;
- `contracts` não importa projeto;
- `api` não importa banco nem adaptadores;
- só o composition root (`internal/app`) importa adaptadores.

---

## 4. Verificação executada

| Verificação | Comando | Resultado |
| --- | --- | --- |
| Formatação | `gofmt -l cmd internal db` | sem saída |
| Vet | `go vet ./...` | sem achados |
| Testes | `go test ./... -count=1` | todos verdes |
| Cobertura | `go test ./... -coverpkg=./internal/...` | 64,6% total; API 84,9% |
| Build | `go build ./...` | verde |
| Smoke local | binário + SQLite + curl | health/ready 200, signup, confirmação e login OK |
| Race | `go test -race ./...` | não executado: CGO desabilitado no Windows local |

**Testes:** 72 funções em 23 arquivos, incluindo unitários, integração HTTP com SQLite real, contrato OpenAPI e arquitetura.

Evidência: [`specs/foundation-platform/evidence/2026-11-20-foundation-tests.md`](../../specs/foundation-platform/evidence/2026-11-20-foundation-tests.md).

### 4.1 Revisão independente

Um revisor em contexto separado encontrou dois problemas altos e lacunas de
atomicidade. Todos foram corrigidos antes deste fechamento:

- signup não enumera mais conta: sempre `202 {"status":"ok"}`;
- perfil tem `version` e update condicional, impedindo que um PATCH desfaça
  suspensão concorrente;
- suspensão, auditoria, outbox e revogação de sessão são atômicas;
- recovery codes usam tabela com consumo condicional no login;
- token de confirmação/reset verifica `RowsAffected`;
- telefone re-salvo não perde verificação;
- texto público rejeita e-mail/telefone/documento/coordenada;
- API não acessa mais `UoW`/repositório direto para auditoria/flags;
- `RAIRE_ENV` ausente assume produção e exige chave.

---

## 5. Cobertura de aceite

Cobertos por teste automatizado:

`AC-01..AC-19`, `AC-23`, `AC-28..AC-32`, `AC-37..AC-46`, `AC-48`, `AC-50..AC-55`, `AC-59`.

Parciais ou fora desta fatia:

| AC | Motivo | Dono |
| --- | --- | --- |
| AC-20..AC-22, AC-58 | exportação, exclusão e direitos | FND-08/FND-17 — Felipe |
| AC-24..AC-27, AC-47, AC-56 | portfólio e imagem | FND-07 — Anderson |
| AC-33 | login sem MFA é rejeitado; endpoint admin sem sessão MFA não existe por desenho | FND-04/FND-09 |
| AC-34..AC-36, AC-49, AC-57, AC-60 | web, termos, a11y, copy | FND-11/12/13/18 — Anderson/Felipe |
| AC-41 | não há fluxo de cliente/pedido na fundação | `service-request` |

---

## 6. Follow-ups registrados

Lista completa em [`docs/status/2026-11-20-foundation-follow-ups.md`](2026-11-20-foundation-follow-ups.md). Os bloqueantes para produção/piloto são:

1. adaptador PostgreSQL/PostGIS + testes de contrato;
2. worker de outbox e idempotência de consumo;
3. proteção CSRF de login e verificação de origem;
4. rate limit compartilhado entre instâncias;
5. enrollment/rotação de MFA pela operação;
6. DPIA, jurídico, DPO e região piloto (Q-01/Q-03/Q-40/Q-38–Q-43).

---

## 7. O que não foi feito

- Nenhum deploy, migration em Postgres, piloto ou dado real.
- Nenhuma tela web, acessibilidade executada, E2E ou axe.
- Nenhum portfólio, exportação, exclusão ou fila de direitos.
- Nenhum worker de outbox, tracing ou exportação de métricas.
- Nenhum pentest, teste de carga ou simulação de incidente.
- Nenhuma aprovação de review, segurança, QA ou produto por humano.

---

## 8. Próximo passo

1. Felipe revisa `FND-09`, `FND-10` e a segurança de `FND-03/04/15/16`.
2. Anderson revisa `FND-05/06` sob a ótica de produto/PII e inicia `FND-07`/`FND-11`/`FND-12`.
3. Victor ataca o adaptador Postgres e o worker de outbox.
4. QA independente roda `go test -race`, E2E e acessibilidade.
5. Somente depois: revisão de produto, segurança e aprovação humana.
