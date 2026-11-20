# Specs

Uma pasta por feature. **Nenhum código nasce sem spec com status `aprovada`** (ou waiver registrado).

## Status

| Feature | Status | Dono | Requisitos | Depende de |
| --- | --- | --- | --- | --- |
| [foundation-platform](foundation-platform/spec.md) | aprovada | Victor | RF-01, RF-02, RF-03, RF-04, RF-05, RF-19; RNF-01, RNF-05, RNF-08, RNF-09, RNF-14 | — |
| [web-experience](web-experience/spec.md) | em revisão | Anderson | RNF-10, RNF-11, RNF-12; jornadas 5.1–5.10 | foundation |
| [local-discovery](local-discovery/spec.md) | em revisão | Victor | RF-06, RF-07, RF-20; RNF-02, RNF-13, RNF-16 | foundation |
| [service-request](service-request/spec.md) | em revisão | Anderson | RF-08, RF-09, RF-10, RF-11, RF-12, RF-16; RNF-05, RNF-06 | foundation |
| [trust-reputation](trust-reputation/spec.md) | em revisão | Felipe | RF-12, RF-13, RF-14, RF-15, RF-17, RF-18; RNF-08, RNF-14 | foundation, service-request |
| [fair-visibility](fair-visibility/spec.md) | em revisão | Victor | RF-07, RF-20; RNF-16 | local-discovery |

## Regras

1. Spec nasce em `specs/<feature>/spec.md` com status `rascunho`.
2. `aprovada` só com revisão humana registrada.
3. Toda spec tem `acceptance.md`, `tickets.md` e `implementation-plan.md`.
4. Todo ticket referencia um requisito da spec.
5. Toda entrega anexa evidência em `evidence/`.
6. Mudança de escopo atualiza a spec antes do código.
7. Spec não substitui ADR: decisão arquitetural relevante vai para `docs/decisions`.
8. Dono e revisores explícitos.

## Templates

- [`_template/spec.md`](_template/spec.md)
- [`_template/acceptance.md`](_template/acceptance.md)
- [`_template/tickets.md`](_template/tickets.md)
- [`_template/implementation-plan.md`](_template/implementation-plan.md)
- [`_template/review.md`](_template/review.md)
- [`_template/qa-report.md`](_template/qa-report.md)
- [`_template/security-review.md`](_template/security-review.md)

## Ordem sugerida

```txt
foundation-platform
  ├── web-experience
  ├── local-discovery
  │     └── fair-visibility
  └── service-request
        └── trust-reputation
```

Fairness básica (caps, explicação, telemetria) entra junto com `local-discovery`; a camada completa e os experimentos ficam em `fair-visibility`.

## Aprovação

Ver `docs/process/approvals.md`. A `foundation-platform` foi aprovada em 2026-11-20; as demais specs seguem em revisão.

## Estado da implementação

Nenhum código de produto implementado. A `foundation-platform` está aprovada e os tickets FND estão prontos; a implementação pode começar pelo `FND-01`.
