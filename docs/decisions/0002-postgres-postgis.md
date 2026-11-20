# ADR 0002 — PostgreSQL + PostGIS em produção; SQLite local/CI

**Status:** proposto  
**Data:** 2026-11-20  
**Decisores:** arquitetura, backend, segurança  
**Relacionado:** `docs/architecture/architecture-plan.md`, `docs/architecture/data-model.md`

## Contexto

O MVP precisa de transações confiáveis, consultas geoespaciais por raio/área, busca textual simples e baixo custo operacional. O harness precisa rodar testes sem infraestrutura externa na máquina de cada pessoa. O volume inicial é pequeno e a relevância ainda será validada.

## Opções consideradas

1. **PostgreSQL + PostGIS em tudo** — fidelidade de produção, mas exige Docker/instalação local e aumenta o atrito do harness.
2. **SQLite em tudo** — zero infraestrutura, mas sem PostGIS e com diferenças de concorrência/tipos.
3. **PostgreSQL + PostGIS em produção e SQLite local/CI** — fidelidade onde importa, velocidade no desenvolvimento; diferença confinada ao adaptador.
4. **PostgreSQL + search engine dedicado** — complexidade prematura.

## Decisão

Usar **PostgreSQL 16+ com PostGIS** como fonte de verdade em produção/staging e **SQLite** no ambiente local/CI, atrás de interfaces de repositório. A query geoespacial é isolada em um adaptador; o domínio não conhece o banco. A busca textual inicial usa `pg_trgm` em produção e uma implementação simples no adaptador local.

## Consequências

**Positivas:**

- testes rodam sem Docker;
- produção usa ACID + PostGIS;
- menos atrito para o time;
- domínio e casos de uso ficam independentes do banco;
- migração para search engine dedicado só com gatilho.

**Negativas:**

- duas implementações de adaptador;
- diferenças de SQL/dialeto;
- risco de teste passar em SQLite e falhar em Postgres.

**Mitigações:**

- testes de contrato do repositório;
- migrations revisadas para os dois bancos quando possível;
- staging com Postgres desde cedo;
- query geoespacial isolada e testada;
- CI com job Postgres quando o custo justificar.

## Critérios de reversão

- diferença de comportamento causar bug em produção;
- custo de manter dois adaptadores superar o benefício;
- Postgres local se tornar trivial no harness.

## Status

Proposto. Requer aprovação humana antes da implementação.
