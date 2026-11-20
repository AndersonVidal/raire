# Architecture Decision Records — RAIRE

**Status:** índice aprovado em 2026-11-20  
**Uso:** registrar decisões arquiteturais relevantes, alternativas e consequências.

---

## ADRs aceitos

| ADR | Decisão | Status |
| --- | --- | --- |
| [0001](0001-go-modular-monolith.md) | Backend em Go com monólito modular | aceito |
| [0002](0002-postgres-postgis.md) | PostgreSQL + PostGIS em produção; SQLite local/CI | aceito |
| [0003](0003-outbox-events.md) | Outbox transacional e eventos internos | aceito |
| [0004](0004-session-rbac.md) | Sessão web, CSRF, RBAC e MFA | aceito |
| [0005](0005-explainable-fair-ranking.md) | Ranking explicável com camada de fairness | aceito |
| [0006](0006-web-stack.md) | Web server-rendered mobile-first | aceito |

## Regras

1. Toda decisão estrutural tem ADR.
2. ADR começa como `proposto`.
3. `aceito` só com aprovação humana.
4. Decisão substituída recebe `substituído por ADR-xxxx`.
5. ADR não é ata de reunião; é registro de decisão e trade-off.
6. Nenhum ADR autoriza código sem spec aprovada.

## Template

- [`_template.md`](_template.md)
