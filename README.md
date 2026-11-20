# RAIRE

RAIRE é uma plataforma de descoberta e conexão entre profissionais autônomos locais e pessoas que buscam serviços. O diferencial do projeto é tratar **fairness, explicabilidade, privacidade e confiança** como requisitos de produto desde o primeiro release.

## Estado atual

**Fase de organização e planejamento. Nenhum código de produto foi implementado ainda.**

O que está pronto:

- PRD revisado (v1.0);
- harness de agentes, comandos e gates;
- specs com critérios de aceite, tickets e planos de implementação;
- arquitetura, ADRs, segurança, privacidade, métricas e roadmap;
- pipeline de trabalho e divisão de tarefas entre Victor, Anderson e Felipe.

Status detalhado: [docs/status/2026-11-20-planning.md](docs/status/2026-11-20-planning.md).  
QA documental: [docs/status/2026-11-20-qa-documental.md](docs/status/2026-11-20-qa-documental.md).  
Review de planejamento: [docs/status/2026-11-20-review-planning.md](docs/status/2026-11-20-review-planning.md).

O que vem depois:

1. revisão e aprovação humana das specs;
2. implementação por spec/ticket, com testes primeiro;
3. gates de review, QA, segurança e trust & safety;
4. evidências e preparação de piloto.

## Comece por aqui

- [PRD revisado](docs/product/prd.md) — escopo, requisitos, critérios de aceite e roadmap.
- [PRD v0.1 (arquivo)](docs/product/prd-v0.1.md) — versão anterior preservada para rastreabilidade.
- [Catálogo detalhado de requisitos](docs/product/requirements-catalog.md) — RF-01 a RF-20.
- [Personas e jornadas](docs/product/personas-and-journeys.md).
- [Métricas e experimentos](docs/product/metrics-and-experiments.md).
- [Riscos](docs/product/risks.md).
- [Roadmap](docs/product/roadmap.md).
- [Arquitetura](docs/architecture/architecture-plan.md).
- [Ranking e fairness](docs/architecture/ranking-and-fairness.md).
- [Harness](docs/process/harness.md).
- [Pipeline e divisão de trabalho](docs/process/pipeline.md).
- [Checklist de acessibilidade](docs/process/accessibility-checklist.md).
- [RACI](docs/team/raci.md).
- [Quadro de tarefas](docs/team/task-board.md).
- [Plano de sprints](docs/team/sprint-plan.md).
- [Tickets em CSV](docs/team/tickets.csv).
- [Specs](specs/README.md).
- [ADRs](docs/decisions/README.md).
- [Segurança e privacidade](docs/security/privacy-and-lgpd.md).

## Como este repositório funciona

1. Requisito nasce no PRD.
2. Spec é escrita em `specs/<feature>/`.
3. Humano aprova a spec.
4. Tickets são refinados e priorizados.
5. Implementação começa pelos testes.
6. Review, QA, segurança e trust & safety rodam em sequência.
7. Evidência é anexada antes de considerar pronto.

## Equipe

| Pessoa | Papel | Ownership principal |
| --- | --- | --- |
| **Victor** | Arquitetura + backend + ranking | foundation, discovery, fairness |
| **Anderson** | Web + produto + acessibilidade | web-experience, jornadas, design system |
| **Felipe** | QA + segurança + privacidade + DevOps | trust-reputation, moderação, evidências e gates |

Detalhes em [RACI](docs/team/raci.md), [pipeline](docs/process/pipeline.md) e [handoffs](docs/team/README.md).

## Licença

Projeto em estágio de pesquisa e validação. Nenhuma licença de produção foi definida.
