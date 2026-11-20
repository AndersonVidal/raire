# Harness de Desenvolvimento — RAIRE

**Status:** operacional para planejamento  
**Versão:** 2.0  
**Data:** 2026-11-20  
**Base:** adaptação do harness do repositório `raire_old`, com gates adicionais de fairness, privacidade, pesquisa e divisão explícita de trabalho.

---

## 1. Para que serve

O harness organiza o desenvolvimento por agentes especializados e pessoas. Ele não substitui julgamento humano: cria trilhos, artefatos, critérios e evidências para que o trabalho seja previsível.

Princípios:

- **spec antes de código**;
- **testes antes de implementação**;
- **revisão independente antes de merge**;
- **evidência antes de "pronto"**;
- **humano decide escopo, hipótese e risco aceitável**;
- **nenhuma decisão de fairness, privacidade ou segurança é automática**.

> **Estado atual:** o repositório está na fase de organização. Nenhum código de produto foi implementado. O harness existe para orientar a próxima fase, não para justificar implementação sem aprovação.

---

## 2. Mapa do harness

```txt
AGENTS.md                      # contrato de entrada para qualquer agente
CLAUDE.md                      # contexto do projeto (lei do repositório)
.claude/agents/                # especialistas
.claude/commands/              # comandos de fluxo
specs/
  README.md                    # quadro de specs
  _template/                   # modelos de spec, aceite, tickets, plano, review
  <feature>/                   # uma pasta por feature
docs/
  product/prd.md               # fonte de requisitos
  product/                     # personas, métricas, riscos, roadmap, questões
  architecture/                # plano e desenho técnico
  process/                     # fluxo, pipeline, review, QA, segurança, CI/CD
  security/                    # ameaças, privacidade e LGPD
  team/                        # RACI, task board, handoffs
  decisions/                   # ADRs
  research/                    # evidências e perguntas
  experiments/                 # protocolos e resultados
```

---

## 3. Agentes

| Agente | Papel | Pode editar? |
| --- | --- | --- |
| `orchestrator` | coordena o fluxo, mantém quadro, cobra gates | sim (specs/board) |
| `product-manager` | PRD, escopo, prioridade, critérios de produto | sim (docs de produto) |
| `researcher` | pesquisa, hipóteses, experimentos, evidências | sim (research/experiments) |
| `spec-writer` | spec, critérios de aceite, contratos conceituais | sim (specs) |
| `architect` | arquitetura, ADRs, dependências, aprovação técnica | sim (docs/architecture, ADRs) |
| `backend-go` | API, domínio, persistência, jobs | sim (código backend) |
| `data-ranking` | ranking, matching, fairness, avaliação offline | sim (código de ranking) |
| `frontend` | web acessível, design system, estados | sim (código web) |
| `qa` | plano e execução de testes, evidências | sim (testes/evidências) |
| `reviewer` | revisão independente de código e spec | não (comenta) |
| `security-privacy` | segurança, LGPD, abuso, ameaças | não (comenta) |
| `trust-safety` | moderação, fraude, reputação, políticas | não (comenta) |
| `devops` | CI/CD, infra, observabilidade, custos | sim (infra) |

**Regra de ouro:** quem implementa não aprova. `reviewer`, `security-privacy` e `trust-safety` nunca editam o código que revisam.

---

## 4. Comandos

| Comando | Uso | Saída |
| --- | --- | --- |
| `/discovery <pergunta>` | investigar hipótese ou dúvida | nota de pesquisa + atualização de hipótese |
| `/prd <tema>` | criar/atualizar requisito no PRD | RF/RNF + critérios + perguntas |
| `/spec <feature>` | criar spec completa | `spec.md`, `acceptance.md`, `tickets.md`, `implementation-plan.md` |
| `/tickets <feature>` | refinar/atualizar tickets | quadro e detalhes |
| `/implement <feature> <ticket>` | implementar ticket aprovado | código + testes + evidência |
| `/review <feature>` | review em sequência | parecer consolidado |
| `/test <escopo>` | rodar testes do escopo | relatório de falhas |
| `/adr <decisão>` | registrar decisão | ADR + índice |
| `/release <feature>` | preparar entrega | checklist + rollback + evidências |
| `/status` | ver próximo passo | resumo do quadro e bloqueios |

---

## 5. Fluxo obrigatório

```txt
1. Descoberta
   /discovery -> evidência -> atualização de hipótese/PRD
        |
        v
2. Definição
   /prd -> RF/RNF -> /spec -> spec rascunho
        |
        v
3. Aprovação humana da spec
   status: aprovada (somente humano)
        |
        v
4. Planejamento
   /tickets -> tickets priorizados -> /adr se necessário
        |
        v
5. Implementação
   /implement -> testes primeiro -> código -> evidência
        |
        v
6. Review
   /review -> reviewer -> security-privacy -> trust-safety -> qa -> product-manager
        |
        v
7. Entrega
   /release -> branch -> PR -> gates verdes -> merge
        |
        v
8. Aprendizado
   métricas -> /discovery -> decisão -> próxima iteração
```

### 5.1 Gate de descoberta

Antes de escrever requisito:

- a pergunta está clara?
- há evidência ou é hipótese explícita?
- qual decisão muda se a resposta for sim/não?
- o experimento é ético e privacy-safe?

### 5.2 Gate de spec

Uma spec só é `aprovada` quando:

- contexto e problema estão claros;
- RFs/RNFs do PRD estão citados;
- requisitos são numerados e verificáveis;
- critérios de aceite cobrem feliz, erro e borda;
- contratos e dados estão descritos;
- fora de escopo está explícito;
- riscos de fairness/privacidade/segurança foram avaliados;
- perguntas abertas têm dono e prazo;
- um humano aprovou.

### 5.3 Gate de implementação

Antes de abrir PR:

- testes escritos antes e falhando;
- testes passando;
- lint/typecheck/build verdes;
- migrations testadas;
- contratos atualizados;
- evidência anexada;
- nada de `.skip`, `TODO` órfão ou log de debug;
- feature flag e rollback definidos.

### 5.4 Gate de review

Ordem obrigatória:

1. `reviewer` — spec, código, testes, dependências.
2. `security-privacy` — auth, PII, abuso, LGPD.
3. `trust-safety` — moderação, reputação, fraude, fairness operacional.
4. `qa` — aceite, regressão, acessibilidade, performance.
5. `product-manager` — aderência ao PRD e critério de produto.

O parecer final é consolidado. Bloqueio de qualquer gate impede merge.

### 5.5 Gate de release

- pipeline verde;
- smoke em staging;
- migração aplicada e rollback testado;
- feature flag com default seguro;
- dashboard e alertas;
- comunicação de mudança;
- aprovação humana para piloto/produção.

---

## 6. Papéis humanos

O harness automatiza execução, não responsabilidade.

| Decisão | Humano responsável |
| --- | --- |
| Visão, escopo e prioridade | fundador/produto |
| Aprovação de spec | produto + arquitetura |
| Política de fairness | produto + pesquisa + segurança |
| Base legal, retenção, DPIA | jurídico/privacidade |
| Aceite de risco de segurança | segurança/privacidade |
| Liberação para piloto | fundador + segurança |
| Mudança de arquitetura | arquitetura |
| Contratação/parceria | fundador |

---

## 7. Artefatos e evidências

Toda feature deve ter:

```txt
specs/<feature>/
├── spec.md
├── acceptance.md
├── tickets.md
├── implementation-plan.md
├── review.md              # parecer consolidado
├── qa-report.md           # execução de testes
├── security-review.md     # parecer de segurança/privacidade
└── evidence/              # prints, relatórios, métricas, fixtures
```

Evidência mínima:

- saída de testes com data e commit;
- screenshots/vídeo das jornadas críticas;
- relatório de acessibilidade;
- resultado de performance;
- simulação de fairness quando aplicável;
- migração e plano de rollback;
- riscos residuais.

---

## 8. Convenções

- **Branch:** `feat/<feature>`, `fix/<issue>`, `docs/<tema>`, `chore/<tema>`.
- **Commit:** Conventional Commits (`feat(discovery): ...`).
- **PR:** título Conventional Commits, link para spec/ticket, checklist de gates.
- **Merge:** squash ou merge commit; nunca direto em `main` sem review.
- **Idioma:** documentação, UI e copy em PT-BR; código e identificadores em inglês.
- **Datas:** UTC no banco; formatadas no cliente.
- **Dinheiro:** se existir, inteiro em centavos; no MVP, fora de escopo.
- **Segredos:** nunca no repositório.
- **Dados reais:** nunca em testes, seeds ou exemplos.

---

## 9. O que o agente deve fazer quando travar

1. Pare.
2. Descreva o que tentou, o que esperava e o que aconteceu.
3. Aponte a spec/requisito relacionado.
4. Proponha no máximo 2 alternativas.
5. Pergunte ao humano com contexto.

Nunca contorne um teste que falha, esconda um erro ou invente requisito.

---

## 10. Definição de pronto (DoD)

Uma entrega só está pronta quando:

- [ ] spec aprovada e vinculada;
- [ ] critérios de aceite testados;
- [ ] testes unitários/integração/E2E aplicáveis verdes;
- [ ] acessibilidade verificada;
- [ ] segurança e privacidade revisadas;
- [ ] fairness avaliada quando aplicável;
- [ ] observabilidade e alertas definidos;
- [ ] documentação e ADR atualizados;
- [ ] feature flag e rollback definidos;
- [ ] evidência anexada;
- [ ] parecer final sem bloqueios;
- [ ] PR aprovado e merge realizado.

---

## 11. Definição de pronto para começar (DoR)

Um ticket só entra em implementação quando:

- spec está `aprovada`;
- ticket tem critério de aceite;
- dependências estão resolvidas;
- contrato/entrada/saída está claro;
- dados e PII foram avaliados;
- feature flag e rollback foram pensados;
- dono e revisor estão definidos.

---

## 12. Métricas do harness

- tempo de spec → primeira entrega;
- taxa de retrabalho por gate;
- bloqueios por categoria;
- cobertura de critérios de aceite;
- incidentes pós-merge;
- tempo de review;
- dívida de documentação;
- aderência ao fluxo.

Revisar o harness a cada ciclo de 4 semanas.

---

## 13. O que foi adaptado do `raire_old`

| Dimensão | `raire_old` | `raire` atual |
| --- | --- | --- |
| Estado | PRD e specs em rascunho | PRD revisado + specs + tickets + pipeline + handoffs |
| Stack | Go, Postgres/PostGIS, React/Vite | Go modular, SQLite local/CI, Postgres alvo, web server-rendered |
| Spec | spec, acceptance, tickets, implementation-plan | mesmos artefatos + matriz de rastreabilidade e dono |
| Review | reviewer, security, qa | reviewer, security-privacy, trust-safety, qa, product-manager |
| Gates | fairness, LGPD, localização/contato | mesmos + gate de acessibilidade e gate de produto |
| Pesquisa | research-plan e experimentos | mantido, com protocolo de experimento e guardrails |
| Entrega | PR com gates | branch → PR → gates → revisão humana → merge |
| Governança | CLAUDE.md como lei | `AGENTS.md` + `CLAUDE.md` + PRD + ADRs + specs + RACI |
| Divisão | agentes | agentes + Victor, Anderson e Felipe com ownership explícito |

**Mantido:** spec antes de código, testes primeiro, review independente, Conventional Commits, ADRs, evidência, proibição de `.skip`/log de debug, regras de dependência.

**Adaptado:** stack do MVP, documentação enxuta, separação entre MVP de validação e fatia técnica, pipeline com handoffs, task board único.

**Deliberadamente não copiado:** regras específicas de outros projetos, deploy automatizado sem aprovação humana e fluxo "sem PR".
