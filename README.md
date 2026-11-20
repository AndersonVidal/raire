# RAIRE

RAIRE é uma plataforma de descoberta e conexão entre profissionais autônomos locais e pessoas que buscam serviços. O diferencial do projeto é tratar **fairness, explicabilidade, privacidade e confiança** como requisitos de produto desde o primeiro release.

## Estado atual

**Fundação implementada localmente, em review. Ainda não há deploy, web, Postgres de produção ou piloto.**

O que está pronto:

- PRD revisado (v1.0), decisões Q-01..Q-50 e `foundation-platform` aprovada;
- harness de agentes, comandos e gates;
- monólito modular em Go com SQLite local/CI;
- conta, sessão, CSRF, rate limit, MFA, confirmação e reset;
- catálogo com sinônimos e administração;
- perfil, área, disponibilidade, contato, localização cifrada e suspensão;
- auditoria append-only, outbox, feature flags, health/readiness e logs JSON;
- OpenAPI, runbook, testes automatizados e evidência;
- specs de web, discovery, pedidos, confiança e fairness em revisão.

Status de implementação: [docs/status/2026-11-20-foundation-implementation.md](docs/status/2026-11-20-foundation-implementation.md).  
Follow-ups: [docs/status/2026-11-20-foundation-follow-ups.md](docs/status/2026-11-20-foundation-follow-ups.md).  
Runbook: [docs/runbooks/foundation.md](docs/runbooks/foundation.md).  
QA documental: [docs/status/2026-11-20-qa-documental.md](docs/status/2026-11-20-qa-documental.md).

O que vem depois:

1. revisar a fundação (`FND-01..FND-19`) com Anderson e Felipe;
2. implementar portfólio, privacidade/direitos e web (`FND-07/08/11/12/13/18`);
3. adaptar Postgres/PostGIS, worker de outbox e rate limit compartilhado;
4. gates de review, QA, segurança e trust & safety;
5. evidências finais e preparação de piloto (bloqueada até DPIA/jurídico).

## Como rodar a fundação

```bash
go test ./... -count=1
RAIRE_ENV=dev go run ./cmd/raire
```

No PowerShell: `$env:RAIRE_ENV='dev'; go run ./cmd/raire`.

O servidor sobe em `http://localhost:8080`, aplica migrations/seed e expõe
`/v1/healthz` e `/v1/readyz`. O contrato está em [`api/openapi.yaml`](api/openapi.yaml)
e o passo a passo em [`docs/runbooks/foundation.md`](docs/runbooks/foundation.md).

Em dev, o token de confirmação aparece no header `X-Dev-Confirm-Token`. Nunca
rode com `RAIRE_ENV=dev` em produção; sem `RAIRE_ENV`, o servidor assume o modo
seguro de produção e exige `RAIRE_ENCRYPTION_KEY`.



## Comece por aqui

- [PRD revisado](docs/product/prd.md) — escopo, requisitos, critérios de aceite e roadmap.
- [PRD v0.1 (arquivo)](docs/product/prd-v0.1.md) — versão anterior preservada para rastreabilidade.
- [Catálogo detalhado de requisitos](docs/product/requirements-catalog.md) — RF-01 a RF-20.
- [Decisões aprovadas](docs/product/decisions-2026-11-20.md) — Q-01 a Q-50.
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
