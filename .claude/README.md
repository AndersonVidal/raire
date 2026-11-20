# .claude — Agentes e Comandos

Este diretório contém o harness de agentes do RAIRE.

## Agentes

- `orchestrator` — fluxo e status.
- `product-manager` — PRD e prioridade.
- `researcher` — hipóteses e experimentos.
- `spec-writer` — specs, aceite e tickets.
- `architect` — arquitetura e ADRs.
- `backend-go` — backend.
- `data-ranking` — busca, ranking e fairness.
- `frontend` — web acessível.
- `qa` — testes e evidências.
- `reviewer` — review independente.
- `security-privacy` — segurança e LGPD.
- `trust-safety` — moderação e reputação.
- `devops` — CI/CD e operação.

## Comandos

`/discovery`, `/prd`, `/spec`, `/tickets`, `/implement`, `/review`, `/test`, `/adr`, `/release`, `/status`.

## Regra

Quem implementa não aprova. Agente de review não edita o código revisado.

## Estado

Fase de planejamento. Nenhum código implementado. Implementação bloqueada até aprovação humana das specs.
