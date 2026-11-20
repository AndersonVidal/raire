# foundation-platform — plano de implementação

**Status:** aprovado  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Victor  
**Depende de:** spec aprovada; ADRs 0001, 0002, 0004, 0006 aceitos; decisões Q-07/Q-44–Q-50 registradas

---

## 1. Estratégia

Entregar a fundação em fatias verticais, começando por contratos e banco, depois identidade, catálogo e perfil, depois admin e privacidade, fechando com web/E2E e hardening. Tudo atrás de feature flags quando possível.

Critérios de início (atendidos em 2026-11-20):

- spec `aprovada`;
- ADRs 0001, 0002, 0004 e 0006 `aceitos`;
- questões Q-07 e Q-44–Q-50 decididas;
- critérios de aceite revisados;
- ambiente local definido (Go + SQLite; Postgres em staging).

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Contratos e esqueleto | módulos, OpenAPI inicial, regras de dependência | build, health, arquitetura | — |
| 2 | Dados | migrations, seed, auditoria, outbox | migration, seed, constraints | 1 |
| 3 | Identidade | cadastro, login, logout, sessão, RBAC | unit, integração, rate limit | 1, 2 |
| 4 | Confirmação/MFA | e-mail, reset, TOTP admin, recovery codes | integração, segurança | 3 |
| 5 | Catálogo | categorias, sinônimos, admin | unit, integração | 2 |
| 6 | Perfil | CRUD, publicação, área, disponibilidade, contato | unit, integração, contrato | 2, 3 |
| 7 | Portfólio | itens, imagem, status | unit, integração | 6 |
| 8 | Privacidade | consentimento, PII, exportação/exclusão | integração, redaction | 3, 6 |
| 9 | Admin | suspensão, MFA, auditoria | integração, authz | 3, 6 |
| 10 | Web conta | cadastro/login/confirmação | render, E2E, a11y | 3, 4 |
| 11 | Web perfil | onboarding, prévia, edição | render, E2E, a11y | 6, 10 |
| 12 | Operação | logs, métricas, flags, health | health, logs | 1 |
| 13 | Aceite | E2E, axe, evidência | E2E, a11y | 10–12 |
| 14 | Handoff | OpenAPI, README, ADRs, runbook | validação | todos |
| 15 | Segurança de conta | password change, CSRF, rate limit, idempotência | unit, integração, segurança | 3, 4 |
| 16 | Dados sensíveis | contato, disponibilidade, criptografia, expurgo | unit, integração, segurança | 6 |
| 17 | Admin completo | categorias, flags, direitos, paginação | integração, authz, métricas | 5, 9 |
| 18 | Web estados | sucesso, sessão expirada, termos, copy | E2E, a11y, copy | 10, 11 |
| 19 | QA hardening | CSRF, rate limit, PII, criptografia, contrato | unit, integração, contrato, segurança | 13, 15–18 |

---

## 3. Paralelismo

Podem rodar em paralelo após a fatia 2:

- catálogo (5) e identidade (3);
- web conta (10) e perfil backend (6);
- privacidade (8) e admin (9), com revisão de segurança;
- segurança de conta (15) e dados sensíveis (16), após 3/4 e 6;
- admin completo (17) e web estados (18), após 5/9 e 10/11;
- QA hardening (19) no fim, com evidência.

Caminho crítico: contratos → dados → identidade → perfil → web perfil → hardening → aceite.

---

## 4. Contratos e dados

- migrations forward-only;
- seed idempotente;
- `audit_events` append-only;
- `outbox_events` transacional;
- read model público sem PII;
- OpenAPI versionado em `/v1`;
- erros padronizados.

---

## 5. Feature flags

| Flag | Default | Uso | Dono |
| --- | --- | --- | --- |
| `foundation.enabled` | off | libera cadastro/perfil | Victor |
| `admin.enabled` | off | libera área admin | Felipe |
| `privacy.export_enabled` | off | libera exportação/exclusão | Felipe |
| `catalog.use_synonyms` | on | busca no catálogo | Victor |
| `auth.password_change` | off | libera troca de senha | Victor |
| `auth.mfa_recovery` | off | libera recovery codes | Felipe |
| `admin.account_requests` | off | libera direitos no admin | Felipe |
| `profile.precise_location` | off | libera localização precisa | Felipe |

---

## 6. Migrações e rollback

- forward-only com rollback lógico;
- expand/contract para colunas novas;
- backup antes de migração destrutiva;
- rollback: flag off + imagem anterior + migration compatível;
- teste em banco vazio e com seed;
- teste de contrato SQLite/Postgres;
- job de expurgo de localização precisa com auditoria e rollback.

---

## 7. Observabilidade

- métricas: cadastro iniciado/concluído, confirmação, login, publicação, suspensão, catálogo, flags, rate limit, outbox.
- logs: `correlation_id`, ator pseudonimizado, sem PII.
- alertas: taxa de erro de cadastro, falha de e-mail, falha de migration, backlog de outbox, pico de rate limit, falha de expurgo.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| onboarding longo | fatiar formulário, teste com profissionais reais |
| PII em log | redaction, testes, revisão de segurança |
| admin sem MFA | bloqueio de release |
| catálogo ruim | seed revisado, sinônimos, feedback |
| acoplamento de contextos | regras de dependência no CI |
| SQLite ≠ Postgres | repositórios isolados, testes de contrato |
| localização precisa vaza | cifra, acesso restrito, retenção, expurgo, testes |
| e-mail duplicado enumera | resposta genérica + teste |
| CSRF/rate limit ausente | middleware + testes de segurança |
| admin sem paginação | cursor + filtros + teste de performance |
| disponibilidade ambígua | janelas + validação + teste |

---

## 9. Critérios de saída

- [ ] spec aprovada;
- [ ] ADRs aceitos;
- [ ] aceite completo (AC-01..AC-60);
- [ ] testes unitários, integração e E2E verdes;
- [ ] acessibilidade validada;
- [ ] segurança/privacidade aprovadas;
- [ ] CSRF, rate limit e idempotência testados;
- [ ] localização precisa cifrada e expurgada;
- [ ] migração e rollback testados;
- [ ] observabilidade e flags;
- [ ] evidência anexada;
- [ ] handoff revisado.
