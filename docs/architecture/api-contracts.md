# Contratos de API — RAIRE MVP

**Status:** proposta para revisão  
**Versão:** 1.0  
**Dono:** Victor  
**Revisores:** Anderson (web), Felipe (segurança)  
**Relacionado:** `docs/architecture/architecture-plan.md`, `docs/architecture/data-model.md`

> Contrato conceitual. O `api/openapi.yaml` só será criado na implementação, a partir deste documento.

---

## 1. Convenções

| Item | Regra |
| --- | --- |
| Base | `/v1` |
| Formato | JSON (`application/json`) |
| Auth | cookie de sessão `HttpOnly` + CSRF em métodos mutáveis |
| Erros | `{ "error": { "code", "message", "details", "correlation_id" } }` |
| Paginação | cursor opaco (`next_cursor`) |
| Idempotência | header `Idempotency-Key` em criação sensível |
| Rate limit | headers `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `Retry-After` |
| Datas | ISO 8601 UTC |
| IDs | opacos, prefixados (`acc_`, `prf_`, `req_`, etc.) |
| PII | nunca em resposta pública, log ou evento |

### 1.1 Códigos de erro

| Código | HTTP | Uso |
| --- | --- | --- |
| `VALIDATION_ERROR` | 400 | entrada inválida |
| `UNAUTHENTICATED` | 401 | sem sessão válida |
| `FORBIDDEN` | 403 | sem permissão |
| `NOT_FOUND` | 404 | recurso inexistente ou invisível |
| `CONFLICT` | 409 | estado incompatível |
| `RATE_LIMITED` | 429 | limite excedido |
| `PROFILE_INCOMPLETE` | 422 | perfil não publicável |
| `INVALID_SERVICE_AREA` | 422 | área/raio inválido |
| `EMAIL_NOT_CONFIRMED` | 403 | e-mail ainda não confirmado |
| `TOKEN_EXPIRED` | 422 | token expirado |
| `TOKEN_INVALID` | 422 | token inválido/reutilizado |
| `INVALID_AVAILABILITY` | 422 | janela de disponibilidade inválida |
| `INVALID_IMAGE` | 422 | imagem inválida/grande |
| `INVALID_PASSWORD` | 422 | senha atual incorreta |
| `CSRF_INVALID` | 403 | CSRF ausente/inválido |
| `MFA_REQUIRED` | 403 | MFA necessário |
| `REQUEST_EXPIRED` | 422 | pedido expirado |
| `INTEREST_LIMIT_REACHED` | 422 | limite de interesses |
| `CONTACT_PROTECTED` | 403 | contato antes do aceite |
| `INVALID_STATE_TRANSITION` | 409 | transição inválida |
| `PII_NOT_ALLOWED` | 422 | PII em campo público |
| `REVIEW_NOT_ELIGIBLE` | 422 | avaliação sem conexão elegível |
| `INTERNAL_ERROR` | 500 | erro interno sem detalhe |

### 1.2 Papéis

- `client`
- `professional`
- `operator`
- `admin`

Um usuário pode ser cliente e profissional; o perfil profissional é recurso separado.

---

## 2. Autenticação

### POST `/v1/auth/signup`

```json
{
  "email": "ana@example.test",
  "password": "senha-longa",
  "display_name": "Ana",
  "accept_terms": true,
  "accept_privacy": true
}
```

Resposta `201`:

```json
{
  "account": { "id": "acc_...", "email": "ana@example.test", "email_verified": false },
  "dev_confirm_token": "..." 
}
```

`dev_confirm_token` só existe em modo dev/teste. Em produção, o token vai por e-mail.

### POST `/v1/auth/login`

```json
{ "email": "ana@example.test", "password": "senha-longa", "mfa_code": "123456" }
```

Resposta `200`:

```json
{
  "account": { "id": "acc_...", "display_name": "Ana", "roles": ["client", "professional"] },
  "csrf_token": "..."
}
```

`mfa_code` é obrigatório para `operator`/`admin`.

### POST `/v1/auth/logout`

Revoga a sessão. Resposta `204`.

### POST `/v1/auth/confirm-email`

```json
{ "token": "..." }
```

### POST `/v1/auth/resend-confirmation`

```json
{ "email": "ana@example.test" }
```

Resposta sempre `202`, sem revelar existência.

### POST `/v1/auth/password-reset/request`

```json
{ "email": "ana@example.test" }
```

Resposta sempre `202`, sem revelar existência.

### POST `/v1/auth/password-reset/confirm`

```json
{ "token": "...", "new_password": "..." }
```

---

## 3. Conta e privacidade

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/me` | dados da conta autenticada |
| PATCH | `/v1/me` | nome, preferência de contato, telefone |
| POST | `/v1/me/password` | trocar senha; revoga as outras sessões |
| GET | `/v1/me/consents` | consentimentos e versões |
| POST | `/v1/me/consents/{kind}/revoke` | revogar consentimento |
| GET | `/v1/me/export` | exportação JSON dos dados |
| POST | `/v1/me/delete` | solicitar exclusão/anonimização |
| GET | `/v1/me/requests` | solicitações de direitos |

**Regra:** `GET /v1/me` nunca retorna senha, token, documento ou localização precisa.

---

## 4. Catálogo

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/categories` | categorias ativas |
| GET | `/v1/categories/{slug}` | detalhe + sinônimos |
| POST | `/v1/admin/categories` | criar (operator/admin) |
| PATCH | `/v1/admin/categories/{id}` | editar |
| POST | `/v1/admin/categories/{id}/deactivate` | desativar |

Categoria inativa não aparece em onboarding/busca.

---

## 5. Perfil profissional e portfólio

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/professionals/me` | perfil do dono |
| POST | `/v1/professionals/me` | criar rascunho |
| PATCH | `/v1/professionals/me` | editar |
| POST | `/v1/professionals/me/publish` | publicar |
| POST | `/v1/professionals/me/pause` | pausar |
| POST | `/v1/professionals/me/resume` | reativar |
| GET | `/v1/professionals/{id}` | perfil público (sem PII) |
| GET | `/v1/professionals/{id}/preview` | prévia do dono |
| GET | `/v1/professionals/{id}/portfolio` | portfólio público |
| POST | `/v1/professionals/me/portfolio` | criar item |
| PATCH | `/v1/portfolio/{id}` | editar item |
| DELETE | `/v1/portfolio/{id}` | remover/despublicar |

### 5.1 Payload de perfil

```json
{
  "display_name": "Ana Souza",
  "bio": "Eletricista com 10 anos de experiência",
  "service_type": "in_person",
  "category_ids": ["cat_eletricista"],
  "city": "Cidade Piloto",
  "neighborhood": "Centro",
  "radius_km": 10,
  "availability_text": "Seg–Sex, 8h–18h",
  "contact_preference": "platform"
}
```

**Validação:**

- `display_name` 2–80;
- `bio` ≤ 500;
- ao menos 1 categoria ativa;
- `in_person`/`hybrid` exigem cidade/bairro/raio;
- `remote` não exige;
- raio 1–50 km;
- contato nunca público.

---

## 6. Descoberta

### GET `/v1/search`

Query:

```txt
q, category, city, neighborhood, lat, lng, radius_km, service_type,
availability, verification, cursor, limit
```

Resposta:

```json
{
  "items": [
    {
      "profile_id": "prf_...",
      "display_name": "Ana Souza",
      "categories": ["eletricista"],
      "service_type": "in_person",
      "distance_label": "até 2 km",
      "availability_label": "Hoje",
      "verification": "contact_verified",
      "rating": { "average": 4.6, "count": 18 },
      "reasons": ["Perto de você", "Disponível hoje"],
      "position": 1
    }
  ],
  "next_cursor": "...",
  "policy_version": "ranking-v1",
  "total_estimate": 42
}
```

**Regras:**

- sem termo/categoria → sugestões, não lista aleatória;
- termo vazio/longo/controle → validação;
- sem localização em presencial → pedir cidade/bairro;
- distância pública em faixa, não coordenada;
- explicação derivada de sinais reais;
- fallback se fairness falhar.

### GET `/v1/search/suggestions`

Sugestões de categorias e termos.

### GET `/v1/public/ranking-transparency`

Conteúdo público sobre critérios de ordenação.

---

## 7. Pedidos e interesses

| Método | Rota | Uso |
| --- | --- | --- |
| POST | `/v1/requests` | criar pedido |
| GET | `/v1/requests/{id}` | detalhe (dono ou elegível) |
| PATCH | `/v1/requests/{id}` | editar enquanto aberto |
| POST | `/v1/requests/{id}/close` | fechar |
| POST | `/v1/requests/{id}/cancel` | cancelar |
| GET | `/v1/requests/me` | meus pedidos |
| GET | `/v1/requests/feed` | oportunidades elegíveis |
| POST | `/v1/requests/{id}/interests` | enviar interesse |
| GET | `/v1/requests/{id}/interests` | listar interessados (dono) |
| POST | `/v1/interests/{id}/accept` | aceitar |
| POST | `/v1/interests/{id}/decline` | recusar |
| POST | `/v1/interests/{id}/withdraw` | retirar |

### 7.1 Payload de pedido

```json
{
  "category_id": "cat_eletricista",
  "title": "Troca de tomada",
  "description": "Preciso trocar duas tomadas na sala",
  "city": "Cidade Piloto",
  "neighborhood": "Centro",
  "urgency": "week",
  "budget_min_cents": 10000,
  "budget_max_cents": 20000,
  "contact_preference": "platform"
}
```

**Validação:**

- descrição 20–1000;
- PII explícita (telefone/e-mail/endereço) bloqueada com `PII_NOT_ALLOWED`;
- limite de 5 pedidos ativos por cliente;
- expiração padrão 7 dias;
- status válido.

### 7.2 Interesse

```json
{ "message": "Consigo atender amanhã à tarde." }
```

**Regras:**

- um interesse por profissional/pedido;
- limite 10 por pedido, 20/dia por profissional;
- pedido expirado/fechado → `REQUEST_EXPIRED`/`INVALID_STATE_TRANSITION`;
- aceite cria conexão e registra consentimento;
- contato só aparece após aceite.

---

## 8. Conexões e contato

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/connections/me` | minhas conexões |
| GET | `/v1/connections/{id}` | detalhe (partes) |
| POST | `/v1/connections/{id}/contacted` | registrar contato |
| POST | `/v1/connections/{id}/complete` | concluir |
| POST | `/v1/connections/{id}/cancel` | cancelar |
| POST | `/v1/connections/{id}/dispute` | contestar |

### 8.1 Resposta de conexão

```json
{
  "id": "conn_...",
  "status": "accepted",
  "context": { "request_id": "req_...", "category": "eletricista" },
  "contact": {
    "visible": true,
    "channel": "email",
    "value": "a***@example.test"
  },
  "counterpart": { "display_name": "Ana Souza", "profile_id": "prf_..." }
}
```

Antes do aceite, `contact.visible = false` e `value = null`. O contato completo só é retornado às partes autorizadas, conforme preferência.

---

## 9. Desfecho, avaliações e reputação

| Método | Rota | Uso |
| --- | --- | --- |
| POST | `/v1/connections/{id}/outcomes` | registrar desfecho |
| POST | `/v1/connections/{id}/reviews` | avaliar |
| GET | `/v1/professionals/{id}/reviews` | avaliações públicas |
| POST | `/v1/reviews/{id}/response` | resposta do avaliado |
| POST | `/v1/reviews/{id}/report` | denunciar avaliação |

### 9.1 Avaliação

```json
{
  "rating": 5,
  "tags": ["pontual", "organizado"],
  "comment": "Serviço bem feito."
}
```

**Regras:**

- conexão `completed`;
- uma por lado;
- autor ≠ avaliado;
- comentário ≤ 500, sem PII;
- janela de edição 24 h;
- resposta única;
- reputação contextual por categoria/papel.

---

## 10. Denúncias, notificações e admin

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/admin/professionals` | perfis; cursor, status, q |
| POST | `/v1/admin/professionals/{id}/suspend` | suspender com motivo/evidência |
| POST | `/v1/admin/professionals/{id}/reactivate` | reativar |
| GET | `/v1/admin/categories` | listar categorias |
| POST | `/v1/admin/categories` | criar categoria |
| PATCH | `/v1/admin/categories/{id}` | editar categoria |
| POST | `/v1/admin/categories/{id}/deactivate` | desativar categoria |
| GET | `/v1/admin/feature-flags` | listar flags |
| PATCH | `/v1/admin/feature-flags/{key}` | alterar flag com default seguro |
| GET | `/v1/admin/account-requests` | solicitações de direitos |
| POST | `/v1/admin/account-requests/{id}/resolve` | resolver solicitação |
| POST | `/v1/reports` | denunciar |
| GET | `/v1/notifications` | histórico |
| POST | `/v1/notifications/{id}/read` | marcar como lida |
| GET | `/v1/admin/reports` | fila de moderação |
| POST | `/v1/admin/reports/{id}/resolve` | resolver |
| POST | `/v1/admin/moderation-actions` | aplicar ação |
| POST | `/v1/admin/moderation-actions/{id}/appeal` | recurso |
| GET | `/v1/admin/audit-events` | auditoria; cursor, ator, alvo, período |
| GET | `/v1/admin/fairness/policies` | políticas |
| POST | `/v1/admin/fairness/policies` | criar política |
| POST | `/v1/admin/fairness/policies/{id}/simulate` | simular |
| POST | `/v1/admin/fairness/policies/{id}/approve` | aprovar |
| POST | `/v1/admin/fairness/policies/{id}/activate` | ativar |
| GET | `/v1/admin/fairness/metrics` | métricas |

**Admin exige MFA e autorização por papel.** Toda ação gera `audit_event` e métrica quando aplicável. Listagens usam cursor e limite.

---

## 11. Health

| Método | Rota | Uso |
| --- | --- | --- |
| GET | `/v1/healthz` | liveness |
| GET | `/v1/readyz` | readiness (banco/migrations) |

---

## 12. Rate limits

| Ação | Limite |
| --- | --- |
| login | 5/min por conta + IP |
| signup | 3/h por IP |
| busca | 60/min por sessão |
| pedido | 5/dia por cliente |
| interesse | 20/dia por profissional |
| denúncia | 10/dia por usuário |
| admin | limite por operador + auditoria |

---

## 13. Segurança do contrato

- CSRF obrigatório em POST/PATCH/DELETE com cookie.
- `Idempotency-Key` em criação de pedido/interesse/conexão.
- Nenhum erro retorna stack ou detalhe interno.
- Nenhuma resposta pública retorna PII.
- `404` em vez de `403` quando revelar existência for risco.
- Rate limit antes de processar payload caro.
- Logs com `correlation_id`, sem corpo sensível.

---

## 14. Critérios de aceite do contrato

- [ ] Todos os endpoints têm request/response documentados.
- [ ] Erros padronizados.
- [ ] Authz por recurso em cada rota.
- [ ] PII fora de resposta pública.
- [ ] Paginação por cursor.
- [ ] Rate limit definido.
- [ ] Idempotência nas criações sensíveis.
- [ ] OpenAPI validado contra fixtures.
- [ ] Revisão de segurança/privacidade registrada.
