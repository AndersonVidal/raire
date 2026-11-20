# Handoff — Victor

**Papel:** Arquiteto + backend + ranking  
**Ownership:** `foundation-platform`, `local-discovery`, `fair-visibility`  
**Revisa:** arquitetura, backend, migrations, contratos, performance, fairness técnica  
**Não aprova:** o próprio código; política de fairness sozinho  
**Status (2026-11-20):** `foundation-platform` implementada localmente e em review; decisões Q-07/Q-44–Q-50 registradas; próximos passos são Postgres, outbox worker e revisão independente.
**Implementação:** `docs/status/2026-11-20-foundation-implementation.md`  
**Follow-ups:** `docs/status/2026-11-20-foundation-follow-ups.md`

---

## 1. Leitura obrigatória

1. `AGENTS.md` e `CLAUDE.md`.
2. `docs/product/prd.md` — visão, escopo, RFs, RNFs e decisões.
3. `docs/product/requirements-catalog.md` — RF-01 a RF-20.
4. `docs/architecture/architecture-plan.md` — contextos e dependências.
5. `docs/architecture/ranking-and-fairness.md` — pipeline, score e governança.
6. `docs/process/pipeline.md` — fluxo, gates e DoR/DoD.
7. `docs/team/raci.md` — decisões e independência.
8. `docs/status/2026-11-20-decision-draft.md` — recomendação para Q-07/Q-44–Q-50.
9. `docs/status/2026-11-20-foundation-review-checklist.md` — o que revisar na spec foundation.
10. `specs/foundation-platform/`, `specs/local-discovery/`, `specs/fair-visibility/`.
11. `docs/product/open-questions.md` — Q-07, Q-44–Q-50.

---

## 2. Foco da primeira semana

### Dia 1 — Contexto

- [ ] Ler PRD, arquitetura e specs da sua frente.
- [ ] Listar dúvidas de requisito, dados e contrato.
- [ ] Verificar se as decisões Q-07, Q-44–Q-50 estão respondidas.

### Dia 2 — Revisão de arquitetura

- [ ] Validar contextos: identity, catalog, directory, discovery, fairness, demand, reputation, moderation, platform.
- [ ] Validar regras de dependência e composição.
- [ ] Revisar ADRs 0001–0005.
- [ ] Apontar o que é complexidade prematura.

### Dia 3 — Revisão de dados

- [ ] Validar modelo de dados proposto nas specs.
- [ ] Confirmar estratégia SQLite local/CI + Postgres alvo.
- [ ] Revisar migrations, seed, auditoria e outbox.
- [ ] Definir invariantes que precisam de teste.

### Dia 4 — Revisão de ranking/fairness

- [ ] Validar score baseline e pesos.
- [ ] Validar caps, exploração, boost e piso de relevância.
- [ ] Definir golden set e métricas.
- [ ] Definir fallback e rollback.
- [ ] Verificar que nenhum atributo sensível entrou.

### Dia 5 — Review cruzado e handoff

- [ ] Revisar specs de Anderson e Felipe sob ótica de contrato/dados.
- [ ] Atualizar task board com dependências.
- [ ] Preparar parecer para aprovação humana.
- [ ] Registrar riscos técnicos e decisões pendentes.

---

## 3. Entregáveis da fase de planejamento

- [ ] Parecer de arquitetura das 6 specs.
- [ ] ADRs 0001–0005 revisados.
- [ ] Modelo de dados e invariantes validados.
- [ ] Estratégia de migrations e rollback.
- [ ] Contrato conceitual de busca e ranking.
- [ ] Golden set inicial e plano de avaliação offline.
- [ ] Dependências e ordem de implementação.
- [ ] Riscos técnicos registrados.

---

## 4. Primeiros tickets após aprovação

**Status:** implementados localmente e em review; próximos são FUP-01 (Postgres), FUP-02 (outbox worker) e FUP-03 (login CSRF).

| Ordem | Ticket | Entrega | Status |
| --- | --- | --- | --- |
| 1 | FND-01 | estrutura modular, contratos e regras de dependência | em review |
| 2 | FND-02 | modelo de dados, migrations e seed | em review |
| 3 | FND-03 | identidade, sessão, RBAC, rate limit, CSRF | em review |
| 4 | FND-04 | confirmação, reset, MFA e recovery codes | em review |
| 5 | FND-05 | catálogo e sinônimos | em review |
| 6 | FND-06 | perfil, área, disponibilidade e contato | em review |
| 7 | FND-15 | segurança de conta e abuso | em review |
| 8 | FND-16 | dados sensíveis e localização | em review |
| 9 | FUP-01 | adaptador PostgreSQL/PostGIS | pendente |
| 10 | FUP-02 | worker de outbox | pendente |
| 11 | LCL-01..LCL-03 | read model, retrieval e ranking baseline | pendente |

Detalhes no [`task-board.md`](task-board.md) e em `specs/*/tickets.md`.

---

## 5. Responsabilidades de review

- Revisar PRs de backend de Anderson/Felipe.
- Revisar migrations e mudanças de schema.
- Revisar performance de busca e ranking.
- Revisar fairness técnica (sem atributo sensível, determinismo, fallback).
- Bloquear merge se houver:
  - acoplamento entre contextos;
  - migration sem rollback;
  - ranking não determinístico;
  - PII em read model público;
  - queda de performance sem justificativa.

---

## 6. Decisões que você aplica

**Status:** decididas em 2026-11-20; registro em `docs/product/decisions-2026-11-20.md`.

- Q-07: um perfil por conta.
- Q-44: SQLite local/CI + Postgres produção.
- Q-45: web server-rendered.
- Q-46: sessão por cookie + CSRF.
- Q-47: MFA TOTP + recovery codes.
- Q-48: outbox interno.
- Q-49: busca geoespacial isolada por interface.
- Q-50: feature flags em banco.

Implemente conforme a spec `foundation-platform` v1.1 e registre qualquer desvio em ADR.

---

## 7. Critérios de sucesso

- arquitetura simples, testável e reversível;
- nenhuma dependência circular;
- migrations seguras;
- ranking explicável, determinístico e justo;
- performance dentro do orçamento;
- contratos claros para Anderson e Felipe;
- nenhum código antes da aprovação humana.
