# RACI e Divisão de Trabalho — RAIRE

**Status:** aprovado para a fundação em 2026-11-20  
**Versão:** 1.1  
**Data:** 2026-11-20  
**Equipe:** Victor, Anderson, Felipe

**Legenda:** R = responsável pela execução; A = aprovador final; C = consultado; I = informado.

---

## 1. Papéis

| Pessoa | Papel | Foco | Não faz |
| --- | --- | --- | --- |
| **Victor** | Arquiteto + backend + ranking | arquitetura, foundation, discovery, fairness, migrations, performance | aprovar o próprio código; decidir sozinho política de fairness |
| **Anderson** | Produto + web + acessibilidade | PRD, jornadas, design system, telas, copy PT-BR, pesquisa qualitativa | aprovar o próprio código; prometer verificação inexistente |
| **Felipe** | QA + segurança + privacidade + DevOps | trust-reputation, moderação, CI/CD, testes, evidências, gates | aprovar o próprio código; liberar exceção de segurança sozinho |

---

## 2. RACI por área

| Área | Victor | Anderson | Felipe | Fundador |
| --- | --- | --- | --- | --- |
| PRD e escopo | C | R | C | A |
| Specs de foundation/discovery/fairness | R | C | C | A |
| Specs de web-experience/service-request | C | R | C | A |
| Specs de trust-reputation | C | C | R | A |
| Arquitetura e ADRs | R | C | C | A |
| Modelo de dados e migrations | R | C | C | I |
| API e contratos | R | C | C | I |
| Web e acessibilidade | C | R | C | I |
| Ranking e fairness | R | C | C | A |
| Segurança e privacidade | C | C | R | A |
| QA e evidências | C | C | R | I |
| CI/CD e operação | C | I | R | A |
| Moderação e trust & safety | C | C | R | A |
| Pesquisa e experimentos | C | R | C | A |
| Release/piloto | C | C | R | A |

---

## 3. RACI por spec

| Spec | Responsável | Revisores obrigatórios | Aprovador |
| --- | --- | --- | --- |
| `foundation-platform` | Victor | Anderson (produto/PII), Felipe (segurança) | Fundador |
| `web-experience` | Anderson | Victor (contratos), Felipe (a11y/QA) | Fundador |
| `local-discovery` | Victor | Anderson (copy/UX), Felipe (fairness/privacidade) | Fundador |
| `service-request` | Anderson | Victor (contratos), Felipe (abuso/PII) | Fundador |
| `trust-reputation` | Felipe | Victor (dados), Anderson (copy/jornada) | Fundador |
| `fair-visibility` | Victor | Anderson (produto), Felipe (segurança/fairness) | Fundador |

---

## 4. Direitos de decisão

| Decisão | Decide | Consulta | Informa |
| --- | --- | --- | --- |
| Escopo do MVP | Fundador | Anderson, Victor, Felipe | todos |
| Prioridade de ticket | Anderson | Victor, Felipe | todos |
| Arquitetura | Victor | Felipe | Anderson |
| Mudança de schema/migration | Victor | Felipe | Anderson |
| Política de ranking/fairness | Fundador | Victor, Anderson, Felipe | todos |
| Base legal/retenção | Jurídico | Felipe, Fundador | todos |
| Aceite de risco de segurança | Felipe | Fundador | todos |
| Liberação para piloto | Fundador | Felipe, Anderson, Victor | todos |
| Exceção de gate | Fundador + dono do gate | — | todos |

---

## 5. Regra de independência

- Quem implementa não aprova.
- O revisor não edita o código revisado.
- O gate de segurança/privacidade tem poder de bloqueio.
- O gate de produto valida aderência ao PRD e à experiência.
- Conflito entre gates escala para o fundador com registro em `docs/process/approvals.md`.

---

## 6. Matriz de backup

| Se ausente | Assume |
| --- | --- |
| Victor | Felipe (arquitetura/backend) com revisão posterior de Victor |
| Anderson | Felipe (produto/web básico) com revisão posterior de Anderson |
| Felipe | Victor (QA/segurança básica) com revisão posterior de Felipe |

Nenhum backup pode aprovar o próprio código.

---

## 7. Capacidade e foco

| Pessoa | Capacidade semanal | Foco principal | Foco secundário |
| --- | --- | --- | --- |
| Victor | ~30 h | foundation, discovery, fairness | revisão de arquitetura |
| Anderson | ~30 h | web, produto, acessibilidade | service-request |
| Felipe | ~30 h | QA, segurança, trust, DevOps | moderação |

> Ajustar conforme disponibilidade real. O plano assume dedicação parcial e prioriza o caminho crítico.

---

## 8. Escalonamento

1. Bloqueio técnico → dono da spec.
2. Conflito entre frentes → daily.
3. Risco de segurança/privacidade → Felipe + Fundador.
4. Mudança de escopo → Anderson + Fundador.
5. Mudança de arquitetura → Victor + Felipe.
6. Fairness → Victor + Anderson + Felipe + Fundador.
