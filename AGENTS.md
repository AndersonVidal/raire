# AGENTS.md — Contrato de entrada para agentes

Este arquivo é o ponto de entrada para qualquer agente de IA que trabalhe no RAIRE. Leia antes de agir.

## 1. O que é o RAIRE

Projeto orientado por pesquisa que investiga a conexão entre profissionais autônomos locais e pessoas que buscam serviços, com foco em:

- descoberta justa;
- distribuição de visibilidade;
- reputação e confiança;
- marketplaces locais;
- matching e relevância;
- acessibilidade digital.

O repositório é uma base de conhecimento viva. **Hipóteses não são fatos.** Nenhuma arquitetura, tecnologia ou requisito é definitivo sem decisão registrada.

## 2. Leia nesta ordem

1. `CLAUDE.md` — contexto e regras do repositório.
2. `docs/product/prd.md` — requisitos e escopo revisados.
3. `docs/architecture/architecture-plan.md` — arquitetura da implementação.
4. `docs/process/harness.md` — fluxo de trabalho e gates.
5. `docs/process/pipeline.md` — pipeline operacional e donos.
6. `docs/process/review-criteria.md` — critérios de review.
7. `specs/README.md` — estado das specs.
8. A spec/ticket que você vai trabalhar.

## 3. Estado atual

**Fase: organização e planejamento. Nenhum código de produto foi implementado ainda.**

O repositório contém:

- PRD revisado (v1.0) em `docs/product/prd.md`;
- PRD v0.1 preservado em `docs/product/prd-v0.1.md`;
- harness de agentes, comandos, specs, gates e templates;
- specs, critérios de aceite, tickets e planos de implementação;
- pipeline de trabalho e divisão de tarefas entre Victor, Anderson e Felipe;
- documentos de arquitetura, segurança, privacidade, métricas e roadmap.

A preparação da implementação foi autorizada pelo waiver de planejamento em `docs/process/waivers/2026-11-20-mvp-planning.md`. A `foundation-platform` foi aprovada em 2026-11-20 e as decisões Q-01..Q-50 foram registradas; as demais specs seguem em revisão. **Ainda não há código de produto implementado.** O waiver não autoriza piloto ou produção.

## 4. Fluxo obrigatório

```txt
/discovery -> /prd -> /spec -> aprovação humana -> /tickets -> /implement -> /review -> /release
```

Regras:

- Nenhum código nasce sem spec `aprovada` (ou waiver registrado).
- Nenhum ticket avança sem critério de aceite.
- Testes primeiro, código depois.
- Quem implementa não aprova.
- Review em sequência: `reviewer` → `security-privacy` → `trust-safety` → `qa` → `product-manager`.
- Merge só com gates verdes e sem bloqueios.
- Evolução incremental; nada de overengineering.

## 5. Agentes

| Agente | Use para |
| --- | --- |
| `orchestrator` | próximo passo, status, destravar |
| `product-manager` | PRD, escopo, prioridade, aceite de produto |
| `researcher` | hipótese, entrevista, experimento, evidência |
| `spec-writer` | spec, aceite, tickets, plano |
| `architect` | arquitetura, ADR, dependências |
| `backend-go` | API, domínio, banco, jobs |
| `data-ranking` | busca, ranking, fairness, simulação |
| `frontend` | web acessível, design system |
| `qa` | testes, E2E, acessibilidade, performance |
| `reviewer` | review independente |
| `security-privacy` | segurança, LGPD, abuso |
| `trust-safety` | moderação, reputação, fraude |
| `devops` | CI/CD, infra, observabilidade |

## 6. Comandos

`/discovery`, `/prd`, `/spec`, `/tickets`, `/implement`, `/review`, `/test`, `/adr`, `/release`, `/status`.

## 7. Equipe e divisão de trabalho

| Pessoa | Papel principal | Ownership |
| --- | --- | --- |
| **Victor** | Arquitetura + backend + ranking | `foundation-platform`, `local-discovery`, `fair-visibility`; ADRs; revisão de arquitetura |
| **Anderson** | Web, produto e acessibilidade | `web-experience`, jornadas, design system, copy PT-BR; apoio a `service-request` |
| **Felipe** | QA, segurança, privacidade e DevOps | `trust-reputation`, moderação, CI/CD, evidências, gates de segurança |

Detalhes em `docs/team/raci.md` e `docs/process/pipeline.md`.

## 8. Não negociáveis

- Não invente requisito.
- Não trate hipótese como fato.
- Não use atributo sensível em ranking.
- Não venda posição.
- Não exponha localização precisa, contato ou documento.
- Não logue PII.
- Não remova crítica legítima.
- Não faça ação administrativa sem auditoria.
- Não use dado real em teste.
- Não esconda teste que falha.
- Não faça merge do próprio código.
- Não faça deploy sem staging e aprovação humana.
- Não marque spec/ADR como aprovado sem humano (waiver de exercício é a exceção registrada).

## 9. Quando travar

Pare e pergunte ao humano com:

1. o que tentou;
2. o que esperava;
3. o que aconteceu;
4. a spec/requisito relacionado;
5. até 2 alternativas.

Nunca contorne um teste, esconda erro ou invente regra.

## 10. Convenções

- Documentação/UI em PT-BR; código e identificadores em inglês.
- Branch: `feat/<feature>`, `fix/<issue>`, `docs/<tema>`, `chore/<tema>`.
- Commit: Conventional Commits.
- Datas em UTC no banco.
- Segredos fora do repositório.
- Toda entrega tem evidência em `specs/<feature>/evidence/`.
