# Evidência — testes da fundação — 2026-11-20

**Status:** execução local concluída; revisão independente executada, aprovação pendente  
**Implementação:** fatia de Victor (`FND-01..06`, `FND-14..16`; dependências de `FND-09/10/17`)  
**Commit de código:** `2602cba`  
**Ambiente:** Windows, Go 1.27.1, SQLite via `modernc.org/sqlite`, sem CGO  
**Banco:** arquivo SQLite temporário por teste; nenhuma fixture real

---

## 1. Comandos e resultados

| Comando | Resultado | Arquivo |
| --- | --- | --- |
| `gofmt -l cmd internal db` | sem saída | — |
| `go vet ./...` | sem achados | `go-vet-2026-11-20.txt` |
| `go build ./...` | verde | `go-build-2026-11-20.txt` |
| `go test ./... -count=1 -v` | todos os pacotes verdes | `go-test-2026-11-20.txt` |
| `go test ./... -coverpkg=./internal/...` | 64,6% total | — |
| `go test ./internal/api -cover` | 84,9% | — |
| smoke do binário + curl | health/ready 200; signup/confirm/login OK | — |

`go test -race ./...` não foi executado neste ambiente porque CGO está
desabilitado. O CI deve rodar em Linux com `CGO_ENABLED=1`.

---

## 2. Cobertura por área

| Área | Evidência |
| --- | --- |
| identidade/sessão | signup, confirmação, login, rotação, logout, reset, troca de senha, MFA, recovery code, rate limit, idempotência |
| catálogo | sinônimos, categoria inativa, CRUD admin, seed idempotente, auditoria |
| perfil | presencial/remoto/híbrido, área, raio, disponibilidade, draft, publish, pause, resume, categoria inativa |
| PII | perfil público e preview sem e-mail/telefone/coordenada; logs sem PII; contato mascarado |
| dados sensíveis | AES-256-GCM, expiração de 90 dias, expurgo e auditoria |
| admin | suspensão, revogação de sessão, reativação, auditoria, paginação/filtros, flags |
| plataforma | migrations, seed, audit append-only, outbox, flags, health/ready, OpenAPI, arquitetura |

---

## 3. Testes de aceite cobertos

`AC-01..AC-19`, `AC-23`, `AC-28..AC-32`, `AC-37..AC-46`, `AC-48`, `AC-50..AC-55`, `AC-59`.

Fora da fatia: `AC-20..AC-22`, `AC-24..AC-27`, `AC-34..AC-36`, `AC-47`,
`AC-49`, `AC-56..AC-58`, `AC-60` — portfólio, direitos, web, a11y e copy.

---

## 4. Smoke test

Fluxo executado com o binário real:

1. `GET /v1/healthz` → `200 {"status":"ok"}`
2. `GET /v1/readyz` → `200 {"status":"ready"}`
3. `POST /v1/auth/signup` → `202 {"status":"ok"}` (token de dev no header quando a conta é nova)
4. `POST /v1/auth/resend-confirmation` → token de dev
5. `POST /v1/auth/confirm-email` → `204`
6. `POST /v1/auth/login` → `200`, cookies `raire_session`/`raire_csrf`, CSRF no corpo

O banco usado foi `workspace/smoke.db` (temporário, ignorado pelo Git). Nenhum
dado real foi usado.

---

## 5. Limitações

- Sem Postgres/PostGIS, E2E, axe, carga, race ou pentest.
- Sem deploy, staging ou dados reais.
- Cobertura não substitui revisão independente de segurança/QA.

## 6. Correções após revisão independente

A revisão em contexto separado apontou enumeração no signup, lost update de
suspensão, consumo não atômico de recovery code, token de e-mail sem
`RowsAffected`, perda de verificação de telefone ao re-salvar, PII em texto
público, acesso direto da API ao repositório e default inseguro de ambiente.
Todas foram corrigidas e têm teste dedicado; a suíte final tem 72 funções.
