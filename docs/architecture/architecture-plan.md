# Plano de Arquitetura — RAIRE MVP

**Status:** proposta para revisão humana  
**Versão:** 1.0  
**Data:** 2026-11-20  
**Dono:** Victor  
**Relacionado:** `docs/product/prd.md`, ADRs 0001–0006, `docs/architecture/ranking-and-fairness.md`, `docs/architecture/data-model.md`

> As decisões deste documento são propostas. Elas só se tornam vinculantes quando registradas em ADR com status `aceito` e aprovadas pelo humano responsável. O objetivo é evitar complexidade prematura e manter o domínio compreensível.

---

## 1. Premissas

1. O MVP atende uma região piloto, poucas categorias e dezenas de milhares de usuários.
2. A prioridade é aprender com segurança, não escalar para milhões.
3. O backend será em Go, modular, com forte separação de responsabilidades.
4. Um único deployable no início; workers podem rodar no mesmo binário.
5. Banco relacional é a fonte de verdade. SQLite para local/CI; PostgreSQL + PostGIS como alvo de produção.
6. Eventos de domínio via outbox interno, sem broker externo no primeiro release.
7. A experiência web é mobile-first, server-rendered e acessível; sem app nativo no MVP.
8. Fairness, explicabilidade e privacidade são requisitos arquiteturais.

---

## 2. Estilo arquitetural

### 2.1 Decisão proposta: monólito modular

Adotar um **monólito modular em Go** com contextos delimitados, contratos explícitos e dependências unidirecionais. Microserviços só entram com gatilho medido (seção 12).

**Por quê:**

- o domínio ainda está em validação;
- transações locais simplificam consistência de conexão, avaliação e moderação;
- um único deploy reduz custo e complexidade operacional;
- contextos bem separados permitem extrair serviços depois;
- o time pequeno ganha velocidade com menos infraestrutura.

**Evitar:**

- monólito em camadas único, com ranking, perfil e moderação misturados;
- microserviços prematuros;
- eventos como única forma de comunicação interna;
- dependências circulares;
- "shared kernel" grande e genérico.

### 2.2 Estrutura de diretórios proposta

```txt
raire/
├── cmd/
│   └── raire/                 # API + web no mesmo binário
├── internal/
│   ├── identity/              # contas, sessões, papéis, consentimento
│   ├── catalog/               # categorias, serviços, sinônimos
│   ├── directory/             # perfis, área, disponibilidade, portfólio
│   ├── discovery/             # retrieval, score, explicação, paginação
│   ├── fairness/              # políticas, exposição, reranking
│   ├── demand/                # pedidos, interesses, conexões
│   ├── reputation/            # avaliações, desfechos, sinais
│   ├── moderation/            # denúncias, ações, recursos, auditoria
│   ├── platform/              # config, db, http, auth, eventos, logs
│   ├── contracts/             # DTOs e eventos compartilhados mínimos
│   ├── store/sqlite/          # adaptador de persistência local/CI
│   ├── api/                   # handlers JSON /v1
│   ├── web/                   # handlers HTML e templates
│   └── app/                   # composição
├── api/openapi.yaml
├── db/migrations/
├── db/seeds/
├── web/templates/
├── web/static/
├── specs/
├── docs/
└── deployments/
```

Cada contexto segue uma divisão interna leve:

```txt
internal/<contexto>/
├── domain/          # entidades, VOs, invariantes, erros
├── app/             # casos de uso, portas, transações
└── adapters/        # http, sqlite/postgres, mailer
```

Não é obrigatório criar todas as pastas vazias. A regra é: **domínio não importa HTTP, banco ou framework**.

### 2.3 Regras de dependência

1. `domain` importa apenas biblioteca padrão e tipos do próprio domínio.
2. `app` importa `domain`, `contracts` e portas de `platform`; nunca `adapters`.
3. `adapters` importam `app`, `domain` e `platform`.
4. Um contexto não importa `adapters` de outro contexto.
5. Integração entre contextos acontece por:
   - interfaces de aplicação expostas pelo contexto provedor;
   - eventos de domínio via outbox;
   - views de leitura explícitas.
6. `platform` não importa nenhum contexto de negócio.
7. `contracts` não importa contextos; só tipos estáveis e versionados.
8. `web` e `api` consomem casos de uso; nunca acessam banco diretamente.

**Ferramentas propostas:** `go vet`, `golangci-lint`, `depguard` e teste de arquitetura.

---

## 3. Contextos delimitados

| Contexto | Responsabilidade | Não faz |
| --- | --- | --- |
| Identity | contas, autenticação, sessão, papéis, consentimento | perfil profissional, ranking |
| Catalog | categorias, subcategorias, sinônimos, tipo de serviço | perfil, ranking |
| Directory | perfil público, serviços, área, disponibilidade, portfólio | busca, reputação, moderação |
| Demand | pedidos, interesses, conexões, estados | ranking, avaliação |
| Discovery | retrieval, score, filtros, explicação, paginação | regras de domínio de perfil |
| Fairness | política de distribuição, exposição, reranking, experimentos | busca HTTP, dados pessoais |
| Reputation | avaliações, respostas, verificações, sinais | moderação de conteúdo |
| Moderation | denúncias, análise, ações, recursos, auditoria | ranking, reputação |
| Platform | config, banco, HTTP, auth, eventos, logs, métricas | regras de negócio |

### 3.1 Mapa de dependências permitido

```txt
identity   ──> platform
catalog    ──> platform
directory  ──> catalog, identity (ids), platform
demand     ──> directory, catalog, identity (ids), platform
discovery  ──> directory (read), catalog (read), fairness, platform
fairness   ──> platform (config, clock)
reputation ──> demand (connection status), directory (ids), identity (ids), platform
moderation ──> identity, directory, demand, reputation, platform
contracts  ──> (nada)
```

Setas indicam "pode usar a interface de aplicação/read model". Nenhum contexto importa `adapters` de outro.

---

## 4. Modelo de domínio

Detalhes em `docs/architecture/data-model.md`. Resumo:

| Agregado | Entidades/VOs | Invariantes principais |
| --- | --- | --- |
| Account | Account, Session, Role, Consent | e-mail único; senha com hash forte; sessão revogável; consentimento versionado |
| ProfessionalProfile | Profile, DisplayName, Bio, VerificationStatus | publica com categoria e área/remoto; nome não vazio |
| ServiceOffering | Offering, ServiceType, CategoryRef | presencial exige área; remoto não exige; categoria ativa |
| ServiceArea | City, Neighborhood, Radius | raio mínimo/máximo; localização precisa privada |
| Availability | Window, Weekday, TimeRange | janelas válidas; fuso consistente |
| PortfolioItem | Item, Title, Description, Image | conteúdo do dono; sem PII de terceiros |
| ServiceRequest | Request, Description, Urgency, BudgetRange, Status | cliente dono; status válido; expiração; localização aproximada |
| Interest | Interest, Message, Status | um por profissional/pedido; limite por período |
| Connection | Connection, Consent, ContactChannel, Status | só após aceite; contato liberado com consentimento |
| Review | Review, Rating, Tags, Comment, Response | uma por lado/conexão; elegibilidade; sem PII |
| Verification | Verification, Type, EvidenceRef, Status | evidência com retenção curta; decisão auditável |
| Report | Report, Target, Reason, Priority, Status | alvo válido; prioridade por risco; não expõe denunciante |
| ModerationAction | Action, Reason, Actor, Appeal | ação auditável; recurso com prazo; menor privilégio |
| FairnessPolicy | Policy, Weights, Caps, Exploration, Version | versão imutável; aprovação humana; rollback |
| VisibilityEvent | Impression, Position, Cohort, PolicyVersion | pseudonimizado; append-only; sem PII |
| Experiment | Experiment, Assignment, Guardrail | assignment determinístico; critério de parada |

### 4.1 Máquinas de estado

**ServiceRequest:** `draft → open → receiving_interests → matched/closed → expired/cancelled`.

**Interest:** `sent → viewed → accepted/declined → withdrawn/expired`.

**Connection:** `pending → accepted → contacted → completed/cancelled/disputed`.

**Review:** `eligible → submitted → published/under_review → removed`.

**Verification:** `not_started → pending → verified/rejected → expired`.

**Report:** `open → triaged → investigating → resolved/dismissed → appealed`.

### 4.2 Eventos de domínio propostos

| Evento | Produtor | Consumidores iniciais |
| --- | --- | --- |
| `identity.account_created` | Identity | Directory, analytics |
| `identity.consent_changed` | Identity | Analytics, Moderation |
| `directory.profile_published` | Directory | Discovery index, analytics |
| `directory.profile_updated` | Directory | Discovery index |
| `demand.request_created` | Demand | Notifications, analytics |
| `demand.interest_expressed` | Demand | Notifications, analytics |
| `demand.connection_accepted` | Demand | Reputation, notifications |
| `demand.connection_completed` | Demand | Reputation, analytics |
| `reputation.review_published` | Reputation | Directory trust signals, analytics |
| `moderation.action_taken` | Moderation | Identity, Directory, audit |
| `fairness.exposure_recorded` | Fairness/Discovery | Analytics, fairness dashboard |
| `fairness.policy_activated` | Fairness | Discovery, audit |

Regras:

- eventos são imutáveis, versionados e sem PII;
- payload carrega IDs e dados mínimos;
- consumo é idempotente por `event_id`;
- outbox na mesma transação do agregado;
- sem broker externo no MVP: worker lê outbox e executa handlers internos.

---

## 5. Dados

### 5.1 Banco

**Produção:** PostgreSQL 16+ com PostGIS.  
**Local/CI:** SQLite (mesmo modelo lógico, repositórios isolados).  
**Justificativa:** zero infraestrutura para o harness e testes; Postgres/PostGIS como alvo de produção. A diferença fica confinada ao adaptador e à query geoespacial.

**Por quê Postgres na produção:**

- ACID para conexão, avaliação e moderação;
- PostGIS resolve raio, polígono e distância;
- menos infraestrutura para operar;
- migração para search engine dedicado só com gatilho.

### 5.2 Tabelas principais (proposta)

```txt
accounts, sessions, consents, email_tokens
professional_profiles, service_offerings, service_areas, availability_windows, portfolio_items
categories, category_synonyms
service_requests, request_interests, connections
reviews, review_responses, verifications
reports, moderation_actions, appeals
fairness_policies, experiments, experiment_assignments
visibility_events
audit_events, outbox_events, notifications, idempotency_keys
```

### 5.3 Índices e consultas

- GiST em `service_areas.area` e `service_requests.location` (Postgres).
- B-tree em `(category_id, status)`, `(profile_id, status)`, `(request_id, status)`.
- Índice trigram (`pg_trgm`) em nome/bio/categoria para busca textual (Postgres).
- Índice composto para feed de pedidos por categoria/área/status.
- Paginação por cursor (`created_at`, `id`).

### 5.4 Separação de PII

- Dados públicos em tabelas de perfil.
- Dados pessoais em tabelas separadas, com criptografia em nível de aplicação para campos sensíveis.
- Chaves fora do banco (secret manager).
- Views/read models públicos não fazem join com PII.
- Eventos analíticos usam IDs pseudonimizados e nunca e-mail/telefone.

### 5.5 Migrations

- Forward-only, versionadas, revisadas.
- Expand/contract para mudanças destrutivas.
- Rollback lógico (flag ou compatibilidade).
- Seed de catálogo versionado e idempotente.
- Teste em banco vazio e com seed.

---

## 6. APIs

### 6.1 Estilo

- REST JSON versionado em `/v1`.
- OpenAPI 3.1 como contrato-fonte.
- Erros padronizados: `code`, `message`, `details`, `correlation_id`.
- Paginação por cursor.
- Idempotency-Key em criações sensíveis.
- Rate limit por IP, conta e ação.
- Sessão por cookie `HttpOnly` + CSRF.
- HTML server-rendered no mesmo binário.

### 6.2 Endpoints iniciais

Detalhes em `docs/architecture/api-contracts.md`.

```txt
POST   /v1/auth/signup
POST   /v1/auth/login
POST   /v1/auth/logout
POST   /v1/auth/confirm-email
POST   /v1/auth/password-reset/request
POST   /v1/auth/password-reset/confirm

GET    /v1/me
PATCH  /v1/me
GET    /v1/me/consents
POST   /v1/me/consents/{kind}/revoke
GET    /v1/me/export
POST   /v1/me/delete

GET    /v1/categories
GET    /v1/categories/{slug}

GET    /v1/professionals
GET    /v1/professionals/{id}
GET    /v1/professionals/{id}/preview
GET    /v1/professionals/me
POST   /v1/professionals/me
PATCH  /v1/professionals/me
POST   /v1/professionals/me/publish
POST   /v1/professionals/me/pause
POST   /v1/professionals/me/resume
GET    /v1/professionals/{id}/portfolio
POST   /v1/professionals/me/portfolio
PATCH  /v1/portfolio/{id}
DELETE /v1/portfolio/{id}

GET    /v1/search
GET    /v1/search/suggestions

POST   /v1/requests
GET    /v1/requests/{id}
PATCH  /v1/requests/{id}
POST   /v1/requests/{id}/close
POST   /v1/requests/{id}/cancel
GET    /v1/requests/me
GET    /v1/requests/feed
POST   /v1/requests/{id}/interests
GET    /v1/requests/{id}/interests
POST   /v1/interests/{id}/accept
POST   /v1/interests/{id}/decline
POST   /v1/interests/{id}/withdraw

GET    /v1/connections/me
GET    /v1/connections/{id}
POST   /v1/connections/{id}/contacted
POST   /v1/connections/{id}/complete
POST   /v1/connections/{id}/cancel
POST   /v1/connections/{id}/reviews
GET    /v1/professionals/{id}/reviews
POST   /v1/reviews/{id}/response
POST   /v1/reviews/{id}/report

POST   /v1/reports
GET    /v1/notifications
POST   /v1/notifications/{id}/read

GET    /v1/admin/reports
POST   /v1/admin/reports/{id}/resolve
POST   /v1/admin/moderation-actions
GET    /v1/admin/audit-events
GET    /v1/admin/fairness/policies
POST   /v1/admin/fairness/policies
POST   /v1/admin/fairness/policies/{id}/simulate
POST   /v1/admin/fairness/policies/{id}/approve
POST   /v1/admin/fairness/policies/{id}/activate
GET    /v1/admin/fairness/metrics

GET    /v1/public/ranking-transparency
GET    /v1/healthz
GET    /v1/readyz
```

### 6.3 Contratos e compatibilidade

- Campos novos são opcionais; remoção só em versão maior.
- Enums têm fallback `unknown` no cliente.
- Contratos validados em CI contra fixtures.
- Frontend nunca depende de campo não documentado.
- Endpoints internos de worker não são públicos.

---

## 7. Descoberta e ranking

Detalhamento em `docs/architecture/ranking-and-fairness.md`. Resumo:

1. **Retrieval:** filtro por categoria, tipo de serviço, status, disponibilidade e área.
2. **Baseline score:** relevância textual/categoria, distância, disponibilidade, verificação, reputação contextual, completude e atividade.
3. **Fairness layer:** caps de exposição, exploração, boost de novo, diversidade.
4. **Explanation:** motivos legíveis derivados dos sinais usados.
5. **Telemetry:** impressão, clique, interesse e conexão com policy version.
6. **Fallback:** baseline sem fairness se a camada falhar, com alerta.
7. **Governança:** política versionada, simulação offline, canary e revisão humana.

**Fora do MVP:** aprendizado online, embeddings, LLM para ranking, leilão de posição.

---

## 8. Segurança e privacidade na arquitetura

### 8.1 Autenticação e autorização

- Senhas com Argon2id.
- Sessões revogáveis, rotação e expiração.
- RBAC: cliente, profissional, operador, administrador.
- Admin com MFA TOTP.
- Autorização por recurso em pedidos, perfis, avaliações e moderação.
- Tokens de reset de uso único e curta duração.

### 8.2 Superfícies de ataque prioritárias

- cadastro e login;
- busca e perfil público (scraping, enumeração);
- pedidos e interesses (spam, assédio);
- contato (vazamento);
- avaliações (manipulação);
- admin (abuso de poder);
- localização (exposição física);
- eventos/analytics (PII).

### 8.3 Controles mínimos

- rate limit por conta/IP/ação;
- validação e normalização de entrada;
- output encoding e CSP;
- CSRF em fluxos com cookie;
- criptografia de PII sensível;
- segredos em secret manager;
- logs estruturados sem PII;
- auditoria append-only;
- dependências verificadas;
- backup cifrado e teste de restore.

### 8.4 LGPD

- base legal documentada por finalidade;
- registro de operações de tratamento;
- DPIA antes do piloto;
- consentimento granular;
- direitos do titular;
- retenção e anonimização;
- DPO definido antes do piloto;
- plano de incidente.

---

## 9. Observabilidade e operação

### 9.1 Pilares

- **Logs:** JSON estruturado, correlation ID, sem PII.
- **Métricas:** latência, erro, saturação, funil, conexão, fairness.
- **Traces:** OpenTelemetry quando justificado.
- **Auditoria:** eventos de negócio e ações administrativas append-only.

### 9.2 SLOs iniciais

| SLO | Meta |
| --- | --- |
| Disponibilidade da API | 99,5% |
| Latência p95 de busca | < 500 ms |
| Latência p95 de APIs | < 300 ms |
| Erro 5xx | < 0,5% |
| Freshness de perfil na busca | < 60 s |
| Notificação in-app | > 99% em 5 min |

### 9.3 Alertas

- queda de conversão de busca;
- aumento de 5xx ou latência;
- pico de denúncias/spam;
- falha de outbox/worker;
- drift de fairness;
- falha de backup/restore;
- exposição de PII em logs.

---

## 10. Frontend

### 10.1 Proposta

- Web responsivo mobile-first, server-rendered.
- `html/template` + CSS, sem SPA no MVP.
- Design system próprio com tokens de acessibilidade.
- Sem lógica de negócio no cliente.
- SEO básico para landing e perfis públicos.
- PWA pode ser avaliada depois.

### 10.2 Requisitos de interface

- fluxo principal em até 3 toques para busca;
- formulários curtos, com salvamento progressivo;
- estados de carregando, vazio, erro e sucesso;
- componentes acessíveis: foco, rótulos, contraste, teclado, leitor de tela;
- PT-BR, sentence case, linguagem simples;
- sem dark patterns, urgência falsa ou notificação agressiva.

---

## 11. Infraestrutura e entrega

### 11.1 Ambientes

- **Local:** SQLite, sem dependência externa; mail catcher/dev inbox.
- **CI:** banco efêmero, testes unitários, integração, contrato, build, scans.
- **Staging:** Postgres, dados sintéticos, smoke/E2E.
- **Piloto/produção:** região única, deploy rolling/canário, feature flags.

### 11.2 Pipeline

1. `gofmt`/`go vet`/`golangci-lint`.
2. Testes unitários.
3. Testes de integração com banco real.
4. Validação OpenAPI e contratos.
5. Build do binário.
6. `govulncheck`, secret scan, dependency scan.
7. E2E e acessibilidade.
8. Migrations em banco efêmero.
9. Deploy em staging + smoke.
10. Aprovação humana.
11. Deploy canário + monitoramento + rollback.

### 11.3 Feature flags

- `foundation.enabled`
- `discovery.enabled`
- `discovery.fairness.enabled`
- `requests.enabled`
- `reviews.enabled`
- `moderation.manual_queue`
- `analytics.research_consent`
- `admin.enabled`

Toda flag tem dono, descrição, default seguro e data de revisão.

---

## 12. Gatilhos para evoluir a arquitetura

| Gatilho | Evolução proposta |
| --- | --- |
| Busca p95 > 800 ms ou relevância limitada | avaliar OpenSearch/Meilisearch |
| Mais de 2 regiões com requisitos distintos | particionamento por região |
| Volume de eventos > capacidade do outbox | broker externo (NATS/Kafka) |
| Time > 8 pessoas ou deploy independente | extrair 1–2 serviços |
| Necessidade de ML/ranking sofisticado | feature store e serviço de scoring, com avaliação de fairness |
| Regulação/exigência de auditoria mais forte | trilha de auditoria dedicada |
| Moderação manual inviável | triagem assistida, sempre com revisão humana |

**Regra:** nenhuma evolução sem métrica, ADR e plano de rollback.

---

## 13. Riscos técnicos

| Risco | Mitigação |
| --- | --- |
| Ranking injusto ou inexplicável | política versionada, simulação, métricas, revisão humana |
| Vazamento de localização/contato | PII separada, criptografia, mascaramento, testes |
| Acoplamento entre contextos | regras de dependência, testes de arquitetura, review |
| Performance de busca | índices, paginação, benchmark, gatilho para search engine |
| Complexidade de eventos | outbox simples, contratos versionados, idempotência |
| Baixa qualidade do catálogo | curadoria, seed versionado, revisão de sinônimos |
| Custo de moderação | fila priorizada, critérios claros, automação só como triagem |
| Dependência de fornecedor | containers, formatos abertos, isolamento de integrações |
| SQLite ≠ Postgres | repositórios isolados, testes de contrato, migration review |

---

## 14. Critérios de aceite da arquitetura

- [ ] Contextos e dependências revisados por arquitetura.
- [ ] ADRs 0001–0006 decididos ou substituídos.
- [ ] Modelo de dados inicial validado com os fluxos do PRD.
- [ ] Estratégia de PII e localização aprovada por segurança/privacidade.
- [ ] Ranking e fairness com plano de avaliação offline.
- [ ] Pipeline CI e ambientes definidos.
- [ ] SLOs e alertas iniciais acordados.
- [ ] Gatilhos de evolução documentados.
- [ ] Nenhuma decisão tratada como definitiva sem ADR aceito.
