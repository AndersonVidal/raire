# RAIRE Context

> **Comece aqui:** leia `AGENTS.md`, `docs/product/prd.md`, `docs/architecture/architecture-plan.md` e `docs/process/harness.md`. Este arquivo é a lei do repositório; o PRD é a fonte de requisitos.

## Sobre o Projeto

RAIRE é um projeto orientado por pesquisa com foco na conexão entre profissionais autônomos locais e pessoas que buscam serviços.

O projeto investiga:
- descoberta justa de profissionais;
- distribuição de visibilidade;
- reputação e confiança;
- dinâmica de marketplaces locais;
- matching entre oferta e demanda;
- mecanismos de recomendação;
- acessibilidade digital para trabalhadores autônomos.

O repositório funciona como:
- base de conhecimento;
- memória arquitetural;
- ambiente de pesquisa;
- centralizador de documentação;
- registro de decisões;
- contexto persistente para IA assistida;
- implementação de referência do MVP.

---

# Estado Atual

**Fase: organização e planejamento. Nenhum código de produto foi implementado ainda.**

O PRD foi revisado para a versão 1.0 (`docs/product/prd.md`) e o harness foi montado a partir do repositório `raire_old`, com adaptações:

- PRD revisado, com escopo, critérios de aceite, pipeline e divisão de trabalho;
- PRD v0.1 arquivado em `docs/product/prd-v0.1.md`;
- visão, pesquisa, arquitetura, processo, ADRs, segurança e specs;
- specs com critérios de aceite, tickets e planos de implementação;
- pipeline de trabalho, RACI e handoffs por pessoa;
- templates de review, QA e segurança;
- nenhuma implementação, migration ou teste executado.

A preparação da implementação foi autorizada pelo waiver de planejamento (`docs/process/waivers/2026-11-20-mvp-planning.md`). **Esse waiver não autoriza código, merge, piloto ou produção.** A implementação começa somente após aprovação humana das specs.

---

# Objetivos Principais

- Estruturar o domínio do problema.
- Validar hipóteses.
- Organizar conhecimento.
- Investigar mecanismos de descoberta justa.
- Explorar arquitetura de longo prazo.
- Construir uma base sólida para futuras validações e MVPs.
- Utilizar IA como ferramenta auxiliar de pesquisa, organização e desenvolvimento.

---

# Direcionamento Técnico Atual

Decisões para o MVP técnico (ver ADRs):

- Backend em **Go** com **monólito modular**.
- Persistência relacional: **SQLite** para local/CI e testes; **PostgreSQL + PostGIS** como alvo de produção (ADR 0002 revisado).
- Migrations versionadas e seed idempotente.
- Frontend **server-rendered** com `html/template`, mobile-first e acessível; sem SPA no MVP (ADR 0006 revisado).
- API JSON versionada em `/v1` e páginas HTML no mesmo binário.
- Sessão em cookie `HttpOnly`, CSRF, RBAC e MFA TOTP para admin.
- Eventos de domínio via outbox interno, sem broker externo no MVP.
- Fairness, explicabilidade e privacidade como requisitos arquiteturais.

Tecnologias ainda não estão definidas de forma definitiva para produção. Toda mudança relevante exige ADR.

---

# Estrutura do Repositório

> Estrutura **planejada** para a implementação. As pastas de código (`cmd/`, `internal/`, `db/`, `api/`, `web/`) serão criadas apenas após a aprovação humana das specs.

```txt
raire/
├── AGENTS.md
├── CLAUDE.md
├── README.md
├── .claude/
│   ├── agents/
│   └── commands/
├── cmd/
│   └── raire/                 # binário da API + web
├── internal/
│   ├── identity/
│   ├── catalog/
│   ├── directory/
│   ├── discovery/
│   ├── fairness/
│   ├── demand/
│   ├── reputation/
│   ├── moderation/
│   ├── platform/
│   ├── store/sqlite/
│   ├── api/
│   ├── web/
│   └── app/
├── db/
│   ├── migrations/
│   └── seeds/
├── api/openapi.yaml
├── web/
│   ├── templates/
│   └── static/
├── specs/
├── docs/
│   ├── vision/
│   ├── research/
│   ├── product/
│   ├── architecture/
│   ├── process/
│   ├── security/
│   ├── team/
│   ├── status/
│   ├── experiments/
│   ├── decisions/
│   └── diary/
├── references/
├── assets/
└── workspace/
```

---

# Propósito das Pastas

## docs/

Memória oficial do projeto.

- `docs/vision/` — problema, proposta, princípios, hipóteses, glossário.
- `docs/research/` — perguntas, referências, fairness, marketplaces, confiança.
- `docs/product/` — PRD, personas, jornadas, métricas, riscos, roadmap.
- `docs/architecture/` — plano, ranking/fairness, notas de implementação.
- `docs/process/` — harness, pipeline, review, QA, segurança, CI/CD, waivers.
- `docs/security/` — modelo de ameaças, privacidade e LGPD.
- `docs/team/` — RACI e divisão de trabalho.
- `docs/experiments/` — protocolos e resultados.
- `docs/decisions/` — ADRs.
- `docs/diary/` — reflexões e mudanças de direção.

## specs/

Uma pasta por feature, com `spec.md`, `acceptance.md`, `tickets.md`, `implementation-plan.md`, `review.md`, `qa-report.md`, `security-review.md` e `evidence/`.

## references/, assets/, workspace/

Materiais externos, recursos visuais e área descartável, respectivamente.

---

# Papel da IA no Projeto

Agentes de IA operam o repositório com autonomia controlada. O harness está descrito em `docs/process/harness.md`.

Especialistas em `.claude/agents/`:
- `orchestrator` — coordena fluxo e gates;
- `product-manager` — PRD, escopo e prioridade;
- `researcher` — hipóteses, evidências e experimentos;
- `spec-writer` — spec, aceite e tickets;
- `architect` — arquitetura, ADRs e dependências;
- `backend-go` — API, domínio e persistência;
- `data-ranking` — busca, ranking e fairness;
- `frontend` — web acessível;
- `qa` — testes, E2E, acessibilidade e performance;
- `reviewer` — review independente;
- `security-privacy` — segurança e LGPD;
- `trust-safety` — moderação, reputação e abuso;
- `devops` — CI/CD, infra e observabilidade.

Comandos: `/discovery`, `/prd`, `/spec`, `/tickets`, `/implement`, `/review`, `/test`, `/adr`, `/release`, `/status`.

A IA deve:
- separar hipótese de fato;
- citar requisito e spec;
- pedir decisão humana em escopo, risco, fairness e privacidade;
- nunca aprovar o próprio código;
- nunca inventar requisito.

---

# Fluxo Obrigatório

```txt
/discovery -> /prd -> /spec -> aprovação humana -> /tickets -> /implement -> /review -> /release
```

1. Feature nasce de requisito no PRD.
2. Spec em `specs/<feature>/spec.md` com status `rascunho`.
3. Status `aprovada` só com revisão humana (ou waiver registrado).
4. Implementação só com spec aprovada: testes primeiro, código depois.
5. Review em sequência: `reviewer` → `security-privacy` → `trust-safety` → `qa` → `product-manager`.
6. Merge só com gates verdes e sem bloqueios.
7. Toda entrega anexa evidência em `specs/<feature>/evidence/`.

**Nunca:**
- commitar direto na `main` sem review;
- abrir PR sem spec/ticket;
- usar dado real em teste;
- expor PII em log/evento;
- vender posição no ranking;
- usar atributo sensível;
- remover crítica legítima;
- marcar spec/ADR como aprovado sem humano.

---

# Regras Gerais

- Não assumir requisitos inexistentes.
- Não transformar hipóteses em fatos.
- Sempre priorizar clareza e rastreabilidade.
- Manter linguagem técnica e objetiva.
- Documentar decisões relevantes.
- Preferir evolução iterativa.
- O repositório deve funcionar como uma base de conhecimento viva.
