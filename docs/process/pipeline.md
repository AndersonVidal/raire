# Pipeline de Trabalho — RAIRE

**Status:** operacional para planejamento  
**Versão:** 1.0  
**Data:** 2026-11-20  
**Donos:** Victor (arquitetura/backend), Anderson (web/produto), Felipe (QA/segurança/DevOps)  
**Relacionado:** `docs/process/harness.md`, `docs/team/raci.md`, `docs/team/task-board.md`, `specs/README.md`

---

## 1. Objetivo

Definir como o trabalho sai do PRD e chega a uma entrega verificável, com donos claros, artefatos obrigatórios, gates e evidências. O pipeline vale para pessoas e agentes.

---

## 2. Visão geral

```txt
[1] Descoberta / requisito
        |
[2] Spec + aceite + tickets + plano
        |
[3] Revisão humana da spec
        |
[4] Refinamento de ticket e dependências
        |
[5] Implementação com testes primeiro
        |
[6] Review de código independente
        |
[7] Gate de segurança e privacidade
        |
[8] Gate de trust & safety
        |
[9] Gate de QA
        |
[10] Gate de produto
        |
[11] Evidência + release candidate
        |
[12] Aprovação humana para merge/piloto
```

Nenhuma etapa pode ser pulada por urgência. Um bloqueio em qualquer gate impede o avanço.

---

## 3. Estágios, entradas, saídas e donos

| # | Estágio | Entrada | Saída | Dono | Gate |
| --- | --- | --- | --- | --- | --- |
| 1 | Descoberta | pergunta, hipótese, evidência | nota de pesquisa, decisão de investigar | Anderson (produto) + pesquisa | pergunta clara e ética |
| 2 | Definição | RF/RNF do PRD | `spec.md`, `acceptance.md`, `tickets.md`, `implementation-plan.md` | Victor/Anderson/Felipe por feature | spec completa |
| 3 | Aprovação | spec em revisão | spec `aprovada` + registro humano | Fundador + produto + arquitetura + segurança | aprovação registrada |
| 4 | Refinamento | spec aprovada | tickets priorizados, dependências, donos | Dono da feature | DoR |
| 5 | Implementação | ticket + aceite | código + testes + evidência | Dono técnico | testes verdes |
| 6 | Review | PR + diff | `review.md` | Reviewer (não autor) | sem bloqueios |
| 7 | Segurança/privacidade | PR + modelo de ameaças | `security-review.md` | Felipe | sem bloqueios críticos |
| 8 | Trust & safety | PR + fluxos de confiança | parecer de trust & safety | Felipe (com apoio de Anderson) | sem bloqueios |
| 9 | QA | PR + aceite | `qa-report.md` + evidência | Felipe | matriz completa |
| 10 | Produto | PR + métricas/copy | aceite de produto | Anderson | aderente ao PRD |
| 11 | Evidência | relatórios | `evidence/` + checklist | Dono + QA | evidência completa |
| 12 | Release | release candidate | aprovação humana + merge/piloto | Fundador + Felipe | checklist de release |

---

## 4. Papéis no pipeline

### 4.1 Victor — arquitetura, backend e ranking

**Responsabilidades:**

- manter arquitetura, ADRs e contratos;
- implementar foundation, discovery e fairness;
- revisar migrations, performance, segurança estrutural;
- garantir determinismo e explicabilidade do ranking;
- revisar código de backend de Anderson/Felipe quando necessário;
- aprovar mudanças de arquitetura.

**Não faz:**

- aprovar o próprio código;
- aprovar política de fairness sozinho;
- liberar piloto sem segurança e produto.

### 4.2 Anderson — web, produto e acessibilidade

**Responsabilidades:**

- manter PRD, personas, jornadas e copy PT-BR;
- implementar web-experience e fluxos de demanda;
- garantir WCAG 2.2 AA, 360 px e teclado;
- validar critérios de produto e aceite;
- conduzir pesquisa qualitativa e teste de usabilidade;
- revisar copy, estados e acessibilidade.

**Não faz:**

- aprovar o próprio código;
- prometer verificação inexistente;
- usar dark pattern.

### 4.3 Felipe — QA, segurança, privacidade e DevOps

**Responsabilidades:**

- manter pipeline, CI/CD, observabilidade e runbooks;
- implementar trust-reputation e moderação;
- revisar segurança, LGPD, PII e abuso;
- executar QA, acessibilidade, performance e evidências;
- bloquear release com risco crítico;
- manter o registro de riscos e o plano de incidente.

**Não faz:**

- aprovar o próprio código;
- esconder falha;
- liberar exceção de segurança sem humano.

### 4.4 Regra de independência

- Quem implementa não aprova.
- O revisor não edita o código revisado.
- O gate de segurança/privacidade tem poder de bloqueio.
- O gate de produto valida aderência ao PRD, não apenas funcionamento técnico.
- Conflito entre gates escala para o fundador com registro.

---

## 5. Fluxo de ticket

### 5.1 Estados

```txt
proposto -> refinado -> pronto -> em implementação -> em review
   -> em QA -> pronto para merge -> mergeado -> concluído
```

Estados de exceção: `bloqueado`, `cancelado`, `reaberto`, `waived` (com aprovação humana).

### 5.2 Campos obrigatórios do ticket

- ID e título;
- spec e requisito (RF/RNF);
- objetivo;
- escopo incluído;
- fora de escopo;
- critérios de aceite (IDs de `acceptance.md`);
- tipo de teste;
- dependências;
- dono;
- revisor;
- tamanho (`P`, `M`, `G`; `GG` deve ser quebrado);
- feature flag/rollback quando aplicável;
- definição de pronto;
- status.

### 5.3 Regras de tamanho

| Tamanho | Estimativa | Regra |
| --- | --- | --- |
| P | até 1 dia | pode ir direto para implementação após DoR |
| M | 2–3 dias | precisa de plano curto |
| G | 4–5 dias | precisa de fatia e evidência |
| GG | > 5 dias | **quebrar** antes de começar |

---

## 6. Fluxo de branches e commits

### 6.1 Branches

```txt
main
 └── feat/<feature>          # feature completa ou fatia
      └── feat/<feature>-<ticket>  # opcional para tickets grandes
fix/<issue>
docs/<tema>
chore/<tema>
```

- `main` é protegida: sem push direto.
- Toda mudança entra por PR.
- PR pequeno e vinculado a uma spec/ticket.
- Não misturar refactor amplo com feature.

### 6.2 Commits

Conventional Commits:

```txt
feat(foundation): add email confirmation token
fix(discovery): handle empty search results
docs(fairness): document exploration policy
test(service-request): cover interest limit
chore(ci): add govulncheck
```

### 6.3 PR

Título Conventional Commits. Corpo com:

```md
## Spec e escopo
- Spec: specs/<feature>/spec.md
- Ticket: <ID>
- RF/RNF: <IDs>

## O que mudou
...

## Como testar
...

## Evidência
...

## Riscos e rollback
...

## Checklist de gates
- [ ] review
- [ ] segurança/privacidade
- [ ] trust & safety
- [ ] QA
- [ ] produto
```

---

## 7. Gates

### 7.1 Gate de código (reviewer)

- aderência à spec e ao ticket;
- testes cobrindo aceite;
- testes verdes, sem `.skip`;
- lint/vet/build verdes;
- erros tipados, entrada validada, authz por recurso;
- transações/idempotência onde necessário;
- sem PII em log/evento;
- dependências justificadas;
- documentação e OpenAPI atualizados.

**Bloqueia:** teste falhando, ausência de teste para critério must, authz quebrada, PII, segredo, dependência vulnerável crítica/alta.

### 7.2 Gate de segurança/privacidade

Checklist em `docs/process/security-privacy.md` e `docs/security/privacy-and-lgpd.md`.

**Bloqueia:** acesso indevido, injeção, bypass de auth, exposição de contato/localização, PII em log, segredo commitado, falta de rate limit em auth.

### 7.3 Gate de trust & safety

- denúncia e moderação auditáveis;
- avaliação só após conexão elegível;
- crítica legítima preservada;
- sem atributo sensível no ranking;
- sem posição paga;
- contato/identidade protegidos.

**Bloqueia:** moderação sem auditoria, remoção silenciosa de crítica, autoavaliação, ranking com proxy sensível.

### 7.4 Gate de QA

- matriz de aceite → teste completa;
- unitários, integração, HTTP/E2E verdes;
- acessibilidade automatizada + manual;
- 360 px e teclado;
- performance dentro do orçamento;
- evidência anexada.

**Bloqueia:** critério sem teste/justificativa, falha reproduzível, regressão, acessibilidade quebrada em jornada crítica.

### 7.5 Gate de produto

- copy PT-BR clara e honesta;
- estados de carregando/vazio/erro/sucesso;
- jornada em 360 px;
- sem dark pattern;
- métrica e instrumentação definidas;
- aderência ao PRD e ao escopo.

**Bloqueia:** promessa enganosa, pay-to-win, fluxo inacessível, requisito inventado.

---

## 8. Pipeline de CI/CD

Detalhes em `docs/process/ci-cd.md`.

### 8.1 Pull request

```txt
1. gofmt -l
2. go vet ./...
3. go test ./... -count=1
4. go test ./... -coverprofile=coverage.out
5. go build ./...
6. govulncheck ./...
7. gosec ./...
8. secret scan
9. migration test (banco vazio + seed)
10. acessibilidade (quando web)
11. contrato/OpenAPI (quando API)
```

### 8.2 Staging

- deploy automático do merge;
- smoke test das jornadas críticas;
- migration aplicada;
- rollback testado;
- evidência anexada.

### 8.3 Produção/piloto

- aprovação humana registrada;
- feature flag com default seguro;
- canário ou rolling;
- monitoramento e alerta;
- rollback pronto.

---

## 9. Definição de Ready (DoR)

Um ticket está pronto quando:

- [ ] spec `aprovada`;
- [ ] requisito e critério de aceite claros;
- [ ] dependências resolvidas;
- [ ] contrato de entrada/saída definido;
- [ ] dados/PII avaliados;
- [ ] feature flag e rollback definidos;
- [ ] dono e revisor definidos;
- [ ] dúvidas bloqueantes resolvidas.

---

## 10. Definição de Done (DoD)

Um ticket está concluído quando:

- [ ] código revisado e aprovado;
- [ ] testes unitários/integração/E2E aplicáveis verdes;
- [ ] aceite coberto;
- [ ] acessibilidade validada;
- [ ] segurança/privacidade aprovadas;
- [ ] fairness avaliada quando aplicável;
- [ ] observabilidade e alertas definidos;
- [ ] documentação/ADR/OpenAPI atualizados;
- [ ] evidência anexada;
- [ ] PR mergeado;
- [ ] task board atualizado.

---

## 11. SLAs de trabalho

| Etapa | SLA alvo | Escalonamento |
| --- | --- | --- |
| Revisão de spec | 2 dias úteis | autor cobra no daily |
| Review de código | 1 dia útil | se > 2 dias, prioriza no daily |
| Gate de segurança | 1 dia útil | bloqueio crítico imediato |
| Gate de QA | 1 dia útil | falha crítica vira bloqueio |
| Correção de bloqueio | 1 dia útil | se > 2 dias, reavaliar escopo |
| Triagem de denúncia | 48 h | risco físico: 4 h |
| Resposta a incidente | imediato | fundador + segurança |

---

## 12. Cerimônias

| Cerimônia | Frequência | Duração | Saída |
| --- | --- | --- | --- |
| Daily assíncrono | diário | 10 min | bloqueios e próximos passos |
| Refinamento de spec/ticket | 2×/semana | 45 min | tickets prontos |
| Review sync | conforme PR | 30 min | parecer e bloqueios |
| Demo interna | semanal | 30 min | evidência e feedback |
| Revisão de riscos/fairness | mensal | 60 min | riscos atualizados |
| Retrospectiva | 2 semanas | 45 min | melhoria do harness |

---

## 13. Comunicação e registro

- Decisão relevante → ADR ou `docs/decisions`.
- Aprendizado → `docs/experiments` ou `docs/research`.
- Mudança de escopo → PRD + spec + task board.
- Bloqueio → task board + daily.
- Incidente → runbook + postmortem sem culpa.
- Evidência → `specs/<feature>/evidence/`.

---

## 14. Critérios de bloqueio de merge

Merge é proibido se:

- spec não está aprovada;
- ticket não tem aceite;
- teste falha ou foi escondido;
- cobertura de aceite está incompleta;
- há PII em log/evento;
- há vulnerabilidade crítica/alta alcançável;
- authz por recurso está ausente;
- acessibilidade de jornada crítica está quebrada;
- ranking usa atributo sensível ou posição paga;
- documentação/migration/rollback está ausente;
- revisor é o autor;
- gate de segurança/privacidade não aprovou.

---

## 15. Pipeline de release

```txt
1. Release candidate (branch + tag)
2. Checklist de release
3. Staging + smoke + E2E
4. Migração + rollback testados
5. Feature flags com default seguro
6. Aprovação humana
7. Deploy canário/rolling
8. Monitoramento + alertas
9. Comunicação de mudança
10. Postmortem se necessário
```

Checklist detalhado em `docs/process/ci-cd.md`.

---

## 16. Primeiro ciclo de trabalho

**Status:** fundação aprovada em 2026-11-20; implementação pode começar.

| Ordem | Trabalho | Dono | Status |
| --- | --- | --- | --- |
| 1 | Revisar PRD e decidir Q-01..Q-50 | Todos | concluído |
| 2 | Revisar e aprovar foundation-platform | Victor + humano | concluído |
| 3 | Implementar FND-01..FND-19 | Victor/Anderson/Felipe | próximo |
| 4 | Revisar e aprovar web-experience | Anderson | pendente |
| 5 | Revisar e aprovar local-discovery e service-request | Victor/Anderson | pendente |
| 6 | Revisar e aprovar trust-reputation | Felipe | pendente |
| 7 | Revisar e aprovar fair-visibility | Victor | pendente |
| 8 | Preparar piloto | Todos | pendente |

Detalhes em `docs/team/task-board.md` e `docs/team/sprint-plan.md`.
