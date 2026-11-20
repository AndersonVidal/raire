# Runbook — Fundação local

**Status:** implementação local disponível; não é runbook de produção  
**Dono:** Victor  
**Última atualização:** 2026-11-20  
**Relacionado:** `api/openapi.yaml`, `specs/foundation-platform/`, `docs/status/2026-11-20-foundation-implementation.md`

> Esta fundação roda com SQLite no desenvolvimento e no CI. PostgreSQL/PostGIS
> de produção ainda **não** tem adaptador implementado; é um follow-up bloqueante
> para produção (ver `docs/status/2026-11-20-foundation-follow-ups.md`).

---

## 1. Pré-requisitos

- Go 1.24 ou superior (testado com Go 1.27.1 no Windows).
- Nenhum banco externo para o fluxo local.
- `curl` opcional para smoke test.

---

## 2. Variáveis de ambiente

| Variável | Default | Uso |
| --- | --- | --- |
| `RAIRE_ENV` | `prod` | `dev`, `test` ou `prod`. Sem valor, assume `prod` e exige chave. |
| `RAIRE_ADDR` | `:8080` | Endereço HTTP. |
| `RAIRE_DB` | `raire.db` | Arquivo SQLite local. |
| `RAIRE_SECURE_COOKIES` | `true` fora de dev/test | Cookies `Secure`. |
| `RAIRE_TRUST_PROXY` | `false` | Confiar em `X-Forwarded-For` (só atrás de proxy conhecido). |
| `RAIRE_ENCRYPTION_KEY` | chave de dev | Base64 de 32 bytes para AES-256-GCM. Obrigatória fora de dev/test. |
| `RAIRE_TERMS_VERSION` | `2026-11-20` | Versão registrada no consentimento. |
| `RAIRE_PRIVACY_VERSION` | `2026-11-20` | Versão registrada no consentimento. |
| `RAIRE_RATE_LIMIT_ENABLED` | `true` | Liga/desliga limites em memória. |
| `RAIRE_ADMIN_EMAIL` | — | Usado só no comando `seed-admin`. |
| `RAIRE_ADMIN_PASSWORD` | — | Usado só no comando `seed-admin`. |

**Nunca** use a chave de dev em produção. Gere uma chave nova:

```bash
openssl rand -base64 32
```

---

## 3. Subir localmente

```bash
RAIRE_ENV=dev go run ./cmd/raire
```

Saída esperada: log JSON `server_started` em `RAIRE_ADDR`. O binário:

- aplica migrations pendentes;
- roda seed idempotente de catálogo e flags;
- inicia o job diário de expurgo de localização precisa;
- expõe `/v1/healthz` e `/v1/readyz`.

Verificação:

```bash
curl -s http://localhost:8080/v1/healthz
curl -s http://localhost:8080/v1/readyz
```

---

## 4. Criar operador de desenvolvimento

```bash
RAIRE_ENV=dev \
RAIRE_ADMIN_EMAIL=ops@example.test \
RAIRE_ADMIN_PASSWORD='senha-segura-123' \
go run ./cmd/raire seed-admin
```

O comando só funciona em `dev`/`test`, grava a senha com Argon2id, criptografa o
segredo TOTP e imprime o segredo e os códigos de recuperação **uma vez**.
Guarde-os apenas no gerenciador de segredos local; não cole em issue, log ou chat.

Login de operador/admin exige TOTP de 6 dígitos (ou um código de recuperação de
uso único). Sem MFA, a API responde `MFA_REQUIRED`.

---

## 5. Fluxo mínimo por API

```bash
# 1. cadastro
curl -i -X POST http://localhost:8080/v1/auth/signup \
  -H 'Content-Type: application/json' \
  -H 'Idempotency-Key: demo-1' \
  -d '{"email":"ana@example.test","password":"senha-segura-123","display_name":"Ana","accept_terms":true,"accept_privacy":true}'

# Em dev, o token de confirmação vem no header X-Dev-Confirm-Token.
# 2. confirmação
curl -i -X POST http://localhost:8080/v1/auth/confirm-email \
  -H 'Content-Type: application/json' \
  -d '{"token":"<token>"}'

# 3. login (guarde o cookie e o csrf_token)
curl -i -c cookies.txt -X POST http://localhost:8080/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"ana@example.test","password":"senha-segura-123"}'

# 4. perfil (POST/PATCH exigem X-CSRF-Token)
curl -b cookies.txt -X POST http://localhost:8080/v1/professionals/me \
  -H 'Content-Type: application/json' \
  -H 'X-CSRF-Token: <csrf_token>' \
  -d '{"display_name":"Ana Souza","service_type":"in_person","category_ids":["cat_eletricista"],"city":"Cidade Piloto","neighborhood":"Centro","radius_km":10}'

# 5. publicação
curl -b cookies.txt -X POST http://localhost:8080/v1/professionals/me/publish \
  -H 'X-CSRF-Token: <csrf_token>'
```

Contrato completo: `api/openapi.yaml`.

---

## 6. Testes e verificações

```bash
gofmt -l cmd internal db
go vet ./...
go test ./... -count=1
go build ./...
```

O teste de arquitetura roda `go list -json ./...` e falha se um contexto importar
adaptadores de outro ou se o domínio importar HTTP/banco.

`go test -race` exige CGO; o ambiente local sem gcc não roda. O CI deve rodar
`CGO_ENABLED=1 go test -race ./...` em Linux.

---

## 7. Banco e migrations

- Migrations ficam em `db/migrations/` e são embutidas no binário.
- `schema_migrations` registra a versão aplicada.
- Rollback é lógico: flag off + binário anterior + migration compatível.
- Seed fica em `db/seeds/` e usa chave natural (`slug`, `key`) com `INSERT OR IGNORE`.
- Para resetar o ambiente local, pare o processo e apague `raire.db*`.

**Não** rode migration destrutiva sem backup e revisão. Produção usa Postgres,
que ainda não tem adaptador (follow-up bloqueante).

---

## 8. Expurgo de localização precisa

O job roda na inicialização e a cada 24 h:

1. seleciona perfis com `precise_lat_enc` e expiração vencida (ou ausente);
2. limpa `precise_lat_enc`, `precise_lng_enc` e `precise_location_expires_at`;
3. grava `audit_events` com ação `directory.location_purged` por perfil;
4. publica `directory.locations_purged` no outbox.

Para testar sem esperar 90 dias, use o teste
`TestPreciseLocationEncryptedAndPurged` ou chame o serviço com um relógio futuro.

---

## 9. Observabilidade

- Logs JSON com `correlation_id`; o access log não inclui query string nem corpo.
- `X-Correlation-ID` aceito quando tem 8–64 caracteres seguros; caso contrário, é gerado.
- Respostas `/v1/` têm `Cache-Control: no-store`.
- Métricas em memória: cadastro, confirmação, login, perfil, catálogo, flags, rate limit.
- Exportação de métricas para backend externo é follow-up (FND-10).

---

## 10. Problemas comuns

| Sintoma | Causa provável | Ação |
| --- | --- | --- |
| `readyz` 503 | banco fechado ou migrations ausentes | verificar `RAIRE_DB` e logs de startup |
| `MFA_REQUIRED` no login | operador sem código TOTP | informar `mfa_code` ou recovery code |
| `CSRF_INVALID` | POST/PATCH sem `X-CSRF-Token` | ler cookie `raire_csrf` ou refazer login |
| `TOKEN_INVALID` | token já usado | solicitar novo reenvio/reset |
| `EMAIL_NOT_CONFIRMED` | publicação antes da confirmação | confirmar e-mail |
| `INVALID_SERVICE_AREA` | híbrido/presencial sem cidade/bairro/raio | completar área |
| `RATE_LIMITED` | limite por IP/conta | aguardar `Retry-After` |

---

## 11. Limites conhecidos

- Sem adaptador Postgres/PostGIS.
- Rate limit em memória (não compartilhado entre instâncias).
- Sem worker de outbox; eventos ficam pendentes para inspeção.
- Sem enrollment/rotação de MFA pela API; contas criadas antes da migration
  `0004` precisam de recovery codes novos.
- Sem fluxo de verificação de telefone por OTP; o campo e o mascaramento existem,
  mas o selo não é emitido.
- Sem web, portfólio, exportação/exclusão e direitos do titular.
- Sem teste de carga, pentest ou DPIA.
- Não aprovado para piloto ou produção.
