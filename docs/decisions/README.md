# Architecture Decision Records — RAIRE

**Status:** índice em revisão  
**Uso:** registrar decisões arquiteturais relevantes, alternativas e consequências.

---

## ADRs propostos

| ADR | Decisão | Status |
| --- | --- | --- |
| [0001](0001-go-modular-monolith.md) | Backend em Go com monólito modular | proposto |
| [0002](0002-postgres-postgis.md) | PostgreSQL + PostGIS em produção; SQLite local/CI | proposto |
| [0003](0003-outbox-events.md) | Outbox transacional e eventos internos | proposto |
| [0004](0004-session-rbac.md) | Sessão web, CSRF, RBAC e MFA | proposto |
| [0005](0005-explainable-fair-ranking.md) | Ranking explicável com camada de fairness | proposto |
| [0006](0006-web-stack.md) | Web server-rendered mobile-first | proposto |

## Regras

1. Toda decisão estrutural tem ADR.
2. ADR começa como `proposto`.
3. `aceito` só com aprovação humana.
4. Decisão substituída recebe `substituído por ADR-xxxx`.
5. ADR não é ata de reunião; é registro de decisão e trade-off.
6. Nenhum ADR autoriza código sem spec aprovada.

## Template

- [`_template.md`](_template.md)
