# Modelo de Dados — RAIRE MVP

**Status:** proposta para revisão  
**Versão:** 1.0  
**Dono:** Victor  
**Revisores:** Felipe (segurança/privacidade), Anderson (produto)  
**Relacionado:** `docs/architecture/architecture-plan.md`, ADR 0002

> Modelo lógico, não implementação. Nomes de tabelas e campos são proposta; migrations só após aprovação da spec.

---

## 1. Princípios

- Banco relacional como fonte de verdade.
- SQLite local/CI; PostgreSQL + PostGIS em produção.
- PII separada de dados públicos.
- IDs opacos, sem dado pessoal.
- Timestamps em UTC.
- Auditoria append-only.
- Eventos via outbox na mesma transação.
- Migrations forward-only com rollback lógico.

---

## 2. Convenções

| Convenção | Regra |
| --- | --- |
| Chave primária | `id TEXT` opaco (prefixo + random) |
| Timestamps | `created_at`, `updated_at`, `deleted_at` em UTC |
| Soft delete | apenas onde a lei/auditoria exigir; preferir anonimização |
| Enums | `TEXT` com check constraint |
| Dinheiro | inteiro em centavos; fora do MVP |
| Localização | lat/lng aproximada; precisa separada e cifrada |
| JSON | apenas para payload de evento/configuração versionada |
| Índices | toda FK e todo filtro de busca |

---

## 3. Entidades

### 3.1 Identity

```txt
accounts
  id, email_normalized (unique), password_hash, display_name,
  role (client|professional|operator|admin),
  email_verified_at, contact_email, contact_phone,
  contact_preference (platform|email|phone),
  mfa_secret, suspended_at, deleted_at, created_at, updated_at

sessions
  id, account_id, token_hash (unique), csrf_token_hash,
  created_at, expires_at, revoked_at, last_seen_at, ip_hash, user_agent_hash

email_tokens
  id, account_id, kind (confirm|reset), token_hash, expires_at, used_at, created_at

consents
  id, account_id, kind (terms|privacy|research|marketing),
  version, granted_at, revoked_at

account_requests
  id, account_id, kind (export|delete|correct), status,
  requested_at, resolved_at, resolution_note
```

**Invariantes:**

- e-mail único normalizado;
- senha nunca em texto;
- token de uso único;
- sessão revogável;
- consentimento versionado;
- exclusão anonimiza dados públicos e preserva auditoria.

### 3.2 Catalog

```txt
categories
  id, slug (unique), name, description, active, sort_order, created_at, updated_at

category_synonyms
  id, category_id, term, normalized_term, created_at

service_offerings
  id, profile_id, category_id, title, description, service_type,
  active, created_at, updated_at
```

**Invariantes:**

- categoria inativa não aparece em onboarding/busca;
- sinônimo normalizado único por categoria;
- desativação preserva histórico;
- perfil com categoria inativa sai da busca até ajuste.

### 3.3 Directory

```txt
professional_profiles
  id, account_id (unique), display_name, bio, service_type (in_person|remote|hybrid),
  status (draft|published|paused|suspended),
  city, neighborhood, public_area_label,
  precise_lat_enc, precise_lng_enc, radius_km,
  availability_text, availability_json,
  contact_preference, verification (unverified|contact_verified|identity_verified),
  published_at, created_at, updated_at, suspended_at, suspended_reason

profile_categories
  profile_id, category_id, primary (bool)

portfolio_items
  id, profile_id, title, description, category_id, image_ref,
  position, status (draft|published|removed), created_at, updated_at, removed_at
```

**Invariantes:**

- `account_id` único no MVP (um perfil por conta);
- `published` exige display_name, categoria e área/remoto;
- presencial/híbrido exige área válida;
- remoto não exige área;
- raio entre mínimo e máximo;
- localização precisa cifrada e nunca no read model público;
- portfólio não é experiência verificada.

### 3.4 Demand

```txt
service_requests
  id, client_id, category_id, title, description, city, neighborhood,
  urgency (today|week|flexible), budget_min_cents, budget_max_cents,
  status (draft|open|receiving_interests|matched|closed|expired|cancelled),
  contact_preference, expires_at, created_at, updated_at, closed_at

request_interests
  id, request_id, professional_id, message, status
  (sent|viewed|accepted|declined|withdrawn|expired),
  created_at, updated_at

connections
  id, request_id, interest_id, client_id, professional_id,
  status (pending|accepted|contacted|completed|cancelled|disputed),
  client_contact_consent_at, professional_contact_consent_at,
  accepted_at, contacted_at, completed_at, cancelled_at, created_at, updated_at
```

**Invariantes:**

- pedido pertence a um cliente;
- interesse único por profissional/pedido;
- pedido fechado/expirado não recebe interesse;
- conexão só existe a partir de interesse aceito ou solicitação direta aceita;
- contato só após aceite e consentimento registrado;
- endereço exato nunca no read model;
- transições válidas apenas.

### 3.5 Reputation

```txt
connection_outcomes
  id, connection_id, author_id, kind (contacted|negotiated|completed|not_completed|useful),
  note, created_at

reviews
  id, connection_id, author_id, subject_id, subject_role (client|professional),
  rating (1..5), tags_json, comment, status
  (submitted|published|under_review|removed),
  created_at, updated_at, published_at, removed_at, removed_reason

review_responses
  id, review_id (unique), author_id, comment, created_at, updated_at

verifications
  id, account_id, kind (contact|identity), status
  (not_started|pending|verified|rejected|expired),
  evidence_ref_enc, decided_by, decided_at, expires_at, created_at
```

**Invariantes:**

- uma avaliação por autor/conexão;
- autor ≠ avaliado;
- avaliação exige conexão elegível;
- resposta única;
- remoção só com motivo e auditoria;
- documento restrito e com retenção;
- reputação contextual, não score único.

### 3.6 Moderation

```txt
reports
  id, reporter_id, target_type (profile|request|review|connection|portfolio),
  target_id, reason, details, priority (low|normal|high|critical),
  status (open|triaged|investigating|resolved|dismissed|appealed),
  created_at, updated_at, resolved_at, resolution_note

moderation_actions
  id, report_id, actor_id, action
  (none|warning|content_removed|suspended|blocked|reactivated),
  reason, target_type, target_id, created_at

appeals
  id, moderation_action_id, appellant_id, reason, status
  (open|accepted|rejected), created_at, resolved_at, resolution_note

audit_events
  id, actor_id, actor_role, action, target_type, target_id,
  reason, correlation_id, metadata_json, created_at
```

**Invariantes:**

- denúncia tem alvo válido;
- denunciante não é exposto;
- toda ação tem ator, motivo e timestamp;
- auditoria append-only;
- recurso tem prazo e revisor diferente quando possível.

### 3.7 Fairness e medição

```txt
fairness_policies
  id, version (unique), status (draft|simulated|approved|active|retired),
  config_json, created_by, created_at, approved_by, approved_at, activated_at

visibility_events
  id, profile_id, session_hash, position, policy_version, experiment_id,
  cohort, query_hash, event_type (impression|click), created_at

experiments
  id, name, hypothesis, status (draft|running|stopped|concluded),
  config_json, guardrails_json, created_at, started_at, stopped_at

experiment_assignments
  id, experiment_id, subject_hash, variant, assigned_at
```

**Invariantes:**

- política ativa única por ambiente;
- versão publicada não é editada;
- nenhum atributo sensível no schema;
- ledger append-only;
- assignment estável;
- retenção definida.

### 3.8 Platform

```txt
outbox_events
  id, event_type, payload_json, occurred_at, processed_at, attempts, last_error

notifications
  id, account_id, kind, title, body, read_at, created_at

feature_flags
  key (unique), enabled, description, owner, review_at, updated_at

idempotency_keys
  key (unique), account_id, operation, created_at, expires_at
```

---

## 4. Relacionamentos principais

```txt
account 1--1 professional_profile
professional_profile 1--N portfolio_item
professional_profile N--N category (profile_categories)
category 1--N category_synonym
account 1--N service_request (client)
service_request 1--N request_interest
request_interest 1--1 connection
connection 1--N review
review 1--1 review_response
account 1--N report (reporter)
report 1--N moderation_action
moderation_action 1--N appeal
professional_profile 1--N visibility_event
fairness_policy 1--N visibility_event
```

---

## 5. Separação de PII

| Dado | Onde fica | Exposição |
| --- | --- | --- |
| e-mail | `accounts.email_normalized` | nunca público |
| telefone | `accounts.contact_phone` | mascarado até aceite |
| localização precisa | `professional_profiles.precise_lat_enc/lng_enc` | nunca pública |
| localização aproximada | `city`, `neighborhood`, `radius_km` | pública |
| documento | `verifications.evidence_ref_enc` | restrito à operação |
| denúncia | `reports.details` | restrito à moderação |
| eventos | `visibility_events` | pseudonimizados |

---

## 6. Migrations e seed

- `0001_init.sql`: tabelas base, constraints e índices.
- `0002_catalog_seed.sql`: categorias, sinônimos e bairros sintéticos.
- `0003_admin_seed.sql`: admin de desenvolvimento (não produção).
- Toda migration tem `up` e plano de rollback lógico.
- Seed idempotente por chave natural (`slug`, `email`).
- Teste em banco vazio e com seed.

---

## 7. Invariantes que exigem teste

- e-mail único;
- senha nunca em texto;
- sessão revogada não autentica;
- token de uso único;
- perfil publicado exige categoria e área/remoto;
- raio válido;
- localização precisa não aparece em resposta pública;
- categoria inativa remove perfil da busca;
- pedido fechado não recebe interesse;
- interesse único por profissional/pedido;
- contato não aparece antes do aceite;
- conexão só após aceite;
- avaliação única por lado;
- autoavaliação bloqueada;
- remoção de avaliação exige motivo e auditoria;
- ação administrativa gera auditoria;
- política de fairness ativa é única;
- ledger de exposição é append-only;
- evento de outbox é idempotente.

---

## 8. Consultas críticas

| Consulta | Índice/filtro |
| --- | --- |
| busca por categoria + área + status | `(category_id, status)` + geo |
| perfil público por id | PK |
| feed de pedidos por categoria/área/status | `(category_id, status, expires_at)` |
| interesses por pedido | `(request_id, status)` |
| conexões por usuário | `(client_id)`, `(professional_id)` |
| avaliações por perfil | `(subject_id, status)` |
| fila de denúncias | `(status, priority, created_at)` |
| ledger por política/sessão | `(policy_version, session_hash)` |
| auditoria por alvo | `(target_type, target_id, created_at)` |

---

## 9. Retenção e exclusão

| Dado | Ação |
| --- | --- |
| conta excluída | anonimizar e-mail, nome, contato; preservar auditoria |
| localização precisa | expurgar após 90 dias |
| eventos analíticos | agregar e expurgar após 13 meses |
| denúncias/auditoria | 5 anos, acesso restrito |
| documentos | 90 dias após decisão |
| logs | 30–90 dias, sem PII |

---

## 10. Critérios de aceite do modelo

- [ ] Cobre todas as jornadas do PRD.
- [ ] Invariantes documentadas e testáveis.
- [ ] PII separada e protegida.
- [ ] Índices para consultas críticas.
- [ ] Migrations e rollback definidos.
- [ ] Seed idempotente.
- [ ] Nenhum atributo sensível no schema.
- [ ] Revisão de segurança/privacidade registrada.
