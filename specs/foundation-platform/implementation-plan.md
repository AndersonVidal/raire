# foundation-platform — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Victor  
**Depende de:** aprovação humana das specs + ADRs 0001, 0002, 0004, 0006

---

## 1. Estratégia

Entregar a fundação em fatias verticais, começando por contratos e banco, depois identidade, catálogo e perfil, depois admin e privacidade, fechando com web/E2E. Tudo atrás de feature flags quando possível.

Nenhum código começa antes de:

- spec `aprovada`;
- ADRs 0001, 0002 e 0004 `aceitos`;
- questões Q-07, Q-44–Q-50 decididas;
- critérios de aceite revisados;
- ambiente local definido.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Contratos e esqueleto | módulos, OpenAPI inicial, regras de dependência | build, health, arquitetura | — |
| 2 | Dados | migrations, seed, auditoria, outbox | migration, seed, constraints | 1 |
| 3 | Identidade | cadastro, login, logout, sessão, RBAC | unit, integração, rate limit | 1, 2 |
| 4 | Confirmação/MFA | e-mail, reset, TOTP admin | integração, segurança | 3 |
| 5 | Catálogo | categorias, sinônimos, admin | unit, integração | 2 |
| 6 | Perfil | CRUD, publicação, área, disponibilidade | unit, integração, contrato | 2, 3 |
| 7 | Portfólio | itens, imagem, status | unit, integração | 6 |
| 8 | Privacidade | consentimento, PII, exportação/exclusão | integração, redaction | 3, 6 |
| 9 | Admin | suspensão, MFA, auditoria | integração, authz | 3, 6 |
| 10 | Web conta | cadastro/login/confirmação | render, E2E, a11y | 3, 4 |
| 11 | Web perfil | onboarding, prévia, edição | render, E2E, a11y | 6, 10 |
| 12 | Operação | logs, métricas, flags, health | health, logs | 1 |
| 13 | Aceite | E2E, axe, evidência | E2E, a11y | 10–12 |
| 14 | Handoff | OpenAPI, README, ADRs, runbook | validação | todos |

---

## 3. Paralelismo

Podem rodar em paralelo após a fatia 2:

- catálogo (5) e identidade (3);
- web conta (10) e perfil backend (6);
- privacidade (8) e admin (9), com revisão de segurança.

Caminho crítico: contratos → dados → identidade → perfil → web perfil → aceite.

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

---

## 6. Migrações e rollback

- forward-only com rollback lógico;
- expand/contract para colunas novas;
- backup antes de migração destrutiva;
- rollback: flag off + imagem anterior + migration compatível;
- teste em banco vazio e com seed.

---

## 7. Observabilidade

- métricas: cadastro iniciado/concluído, confirmação, login, publicação, suspensão.
- logs: `correlation_id`, ator pseudonimizado, sem PII.
- alertas: taxa de erro de cadastro, falha de e-mail, falha de migration.

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

---

## 9. Critérios de saída

- [ ] spec aprovada;
- [ ] ADRs aceitos;
- [ ] aceite completo;
- [ ] testes unitários, integração e E2E verdes;
- [ ] acessibilidade validada;
- [ ] segurança/privacidade aprovadas;
- [ ] migração e rollback testados;
- [ ] observabilidade e flags;
- [ ] evidência anexada;
- [ ] handoff revisado.
