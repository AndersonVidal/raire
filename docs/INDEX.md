# RAIRE Knowledge Index

Este documento é o índice central da base de conhecimento do projeto RAIRE. Ele facilita navegação, organização, rastreabilidade e contexto para agentes de IA.

**Estado:** organização e planejamento. Nenhum código implementado.

---

## Comece por aqui

1. [`AGENTS.md`](../AGENTS.md) — contrato de entrada.
2. [`CLAUDE.md`](../CLAUDE.md) — contexto e regras.
3. [`docs/product/prd.md`](product/prd.md) — PRD v1.0.
4. [`docs/process/pipeline.md`](process/pipeline.md) — pipeline e donos.
5. [`docs/team/task-board.md`](team/task-board.md) — tarefas.
6. [`docs/status/2026-11-20-planning.md`](status/2026-11-20-planning.md) — status detalhado.

---

## Product

| Documento | Uso |
| --- | --- |
| [prd.md](product/prd.md) | PRD revisado v1.0 |
| [prd-v0.1.md](product/prd-v0.1.md) | versão anterior arquivada |
| [requirements-catalog.md](product/requirements-catalog.md) | RF-01 a RF-20 detalhados |
| [personas-and-journeys.md](product/personas-and-journeys.md) | personas, atores, JTBD e jornadas |
| [metrics-and-experiments.md](product/metrics-and-experiments.md) | métricas, metas e experimentos |
| [risks.md](product/risks.md) | riscos e mitigação |
| [roadmap.md](product/roadmap.md) | fases F0–F7 |
| [open-questions.md](product/open-questions.md) | decisões pendentes e defaults |
| [catalog-seed.md](product/catalog-seed.md) | seed inicial proposto |
| [transparency-ranking.md](product/transparency-ranking.md) | copy pública do ranking |

## Vision

- [problem.md](vision/problem.md)
- [proposal.md](vision/proposal.md)
- [principles.md](vision/principles.md)
- [hypothesis.md](vision/hypothesis.md)
- [service-types.md](vision/service-types.md)
- [glossary.md](vision/glossary.md)

## Research

- [research-plan.md](research/research-plan.md)
- [fairness.md](research/fairness.md)
- [ranking-signals.md](research/ranking-signals.md)
- [_template.md](research/_template.md)

## Architecture

- [architecture-plan.md](architecture/architecture-plan.md)
- [ranking-and-fairness.md](architecture/ranking-and-fairness.md)
- [data-model.md](architecture/data-model.md)
- [api-contracts.md](architecture/api-contracts.md)

## Process

- [harness.md](process/harness.md) — agentes, comandos e gates
- [pipeline.md](process/pipeline.md) — pipeline, donos, SLAs
- [review-criteria.md](process/review-criteria.md) — gates de review
- [qa-strategy.md](process/qa-strategy.md) — testes e evidências
- [accessibility-checklist.md](process/accessibility-checklist.md) — checklist WCAG 2.2 AA
- [security-privacy.md](process/security-privacy.md) — segurança e LGPD
- [ci-cd.md](process/ci-cd.md) — CI/CD, ambientes e release
- [approvals.md](process/approvals.md) — aprovações e gates
- [waivers/2026-11-20-mvp-planning.md](process/waivers/2026-11-20-mvp-planning.md) — autorização de planejamento

## Security

- [threat-model.md](security/threat-model.md)
- [privacy-and-lgpd.md](security/privacy-and-lgpd.md)
- [incident-response.md](security/incident-response.md)

## Team

- [README.md](team/README.md) — índice e handoffs
- [raci.md](team/raci.md) — RACI e decisões
- [task-board.md](team/task-board.md) — todas as tarefas
- [sprint-plan.md](team/sprint-plan.md) — plano de sprints
- [tickets.csv](team/tickets.csv) — importável em ferramenta de gestão
- [handoff-victor.md](team/handoff-victor.md)
- [handoff-anderson.md](team/handoff-anderson.md)
- [handoff-felipe.md](team/handoff-felipe.md)

## Decisions

- [README.md](decisions/README.md)
- [0001 — Go modular](decisions/0001-go-modular-monolith.md)
- [0002 — Postgres/SQLite](decisions/0002-postgres-postgis.md)
- [0003 — Outbox](decisions/0003-outbox-events.md)
- [0004 — Sessão/RBAC/MFA](decisions/0004-session-rbac.md)
- [0005 — Ranking/fairness](decisions/0005-explainable-fair-ranking.md)
- [0006 — Web server-rendered](decisions/0006-web-stack.md)

## Specs

- [specs/README.md](../specs/README.md)
- [foundation-platform](../specs/foundation-platform/spec.md)
- [web-experience](../specs/web-experience/spec.md)
- [local-discovery](../specs/local-discovery/spec.md)
- [service-request](../specs/service-request/spec.md)
- [trust-reputation](../specs/trust-reputation/spec.md)
- [fair-visibility](../specs/fair-visibility/spec.md)

## Experiments

- [_template.md](experiments/_template.md)

## Outras pastas

- `references/` — materiais externos.
- `assets/` — logos, diagramas, mockups.
- `workspace/` — rascunhos descartáveis.
- `docs/diary/` — reflexões e mudanças de direção.

---

## Regra de ouro

Hipótese não é fato. Spec não é código. Código não é aprovado por quem escreveu.
