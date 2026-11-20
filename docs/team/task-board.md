# Task Board — RAIRE

**Status:** planejamento  
**Versão:** 1.0  
**Data:** 2026-11-20  
**Donos:** Victor, Anderson, Felipe  
**Regra:** nenhum ticket entra em implementação antes da spec aprovada e do DoR.

---

## 1. Legenda

| Status | Significado |
| --- | --- |
| `proposto` | ideia de trabalho; ainda não refinado |
| `refinado` | escopo, aceite e dependências claros |
| `pronto` | DoR atendido; pode entrar em implementação |
| `em implementação` | trabalho ativo |
| `em review` | PR aberto |
| `em QA` | testes/aceite |
| `bloqueado` | depende de decisão/externo |
| `concluído` | DoD atendido |
| `cancelado` | não será feito |

**Tamanhos:** `P` (até 1 dia), `M` (2–3), `G` (4–5), `GG` (quebrar).

---

## 2. Resumo por pessoa

### Victor — arquitetura, backend e ranking

| Fase | Tickets | Total |
| --- | --- | --- |
| F1 | FND-01, FND-02, FND-03, FND-04, FND-05, FND-06, FND-14 | 7 |
| F2 | LCL-01, LCL-02, LCL-03, LCL-06, LCL-07, LCL-08, LCL-09, LCL-10 | 8 |
| F5 | FAV-01, FAV-02, FAV-03, FAV-04, FAV-05, FAV-07, FAV-08, FAV-09, FAV-10 | 9 |
| Review | arquitetura, migrations, contratos, performance | contínuo |

### Anderson — produto, web e acessibilidade

| Fase | Tickets | Total |
| --- | --- | --- |
| F1 | WEB-01, WEB-02, WEB-03, WEB-04, FND-07, FND-11, FND-12 | 7 |
| F2 | LCL-04, LCL-05 (apoio), WEB-05, WEB-06 | 4 |
| F3 | SRV-01, SRV-02, SRV-03, SRV-04, SRV-05, SRV-07, SRV-10, WEB-07, WEB-08 | 9 |
| F4 | WEB-09 (apoio), WEB-10, WEB-12 | 3 |
| F5 | FAV-06, FAV-10 (apoio) | 2 |
| Review | produto, copy, acessibilidade | contínuo |

### Felipe — QA, segurança, privacidade e DevOps

| Fase | Tickets | Total |
| --- | --- | --- |
| F1 | FND-08, FND-09, FND-10, FND-13, WEB-11 | 5 |
| F3 | SRV-06, SRV-08, SRV-09 | 3 |
| F4 | TRS-01, TRS-02, TRS-03, TRS-04, TRS-05, TRS-06, TRS-07, TRS-08, TRS-09, TRS-10, WEB-09, WEB-11 | 12 |
| F6 | evidências, segurança, release, runbook | contínuo |
| Review | segurança, privacidade, QA, trust & safety | contínuo |

---

## 3. Board por fase

### F0 — Planejamento e revisão

| ID | Tarefa | Dono | Depende de | Status |
| --- | --- | --- | --- | --- |
| PLAN-01 | Revisar PRD v1.0 e responder Q-07, Q-44–Q-50 | Victor | — | proposto |
| PLAN-02 | Revisar personas, jornadas e copy | Anderson | — | proposto |
| PLAN-03 | Revisar riscos, segurança e LGPD | Felipe | — | proposto |
| PLAN-04 | Revisar spec foundation-platform | Victor + Felipe | PLAN-01 | proposto |
| PLAN-05 | Revisar spec web-experience | Anderson + Felipe | PLAN-02 | proposto |
| PLAN-06 | Revisar spec local-discovery | Victor + Anderson | PLAN-01 | proposto |
| PLAN-07 | Revisar spec service-request | Anderson + Felipe | PLAN-02 | proposto |
| PLAN-08 | Revisar spec trust-reputation | Felipe + Victor | PLAN-03 | proposto |
| PLAN-09 | Revisar spec fair-visibility | Victor + Felipe | PLAN-06 | proposto |
| PLAN-10 | Aprovar specs (humano) | Fundador | PLAN-04..PLAN-09 | bloqueado |
| PLAN-11 | Refinar tickets e dependências | Todos | PLAN-10 | proposto |
| PLAN-12 | Preparar ambiente e CI | Felipe | PLAN-10 | proposto |

---

### F1 — Fundação

| ID | Tarefa | Dono | Tam. | Depende de | Aceite | Status |
| --- | --- | --- | --- | --- | --- | --- |
| FND-01 | Estrutura modular, contratos e dependências | Victor | M | PLAN-10 | AC-37 | proposto |
| FND-02 | Modelo de dados, migrations e seed | Victor | G | FND-01 | AC-01, AC-09..AC-18, AC-21, AC-24..AC-30 | proposto |
| FND-03 | Identidade, sessão, RBAC e rate limit | Victor | G | FND-01, FND-02 | AC-01..AC-08 | proposto |
| FND-04 | Confirmação de e-mail, reset e MFA | Victor | M | FND-03 | AC-03, AC-06, AC-08, AC-33 | proposto |
| FND-05 | Catálogo, sinônimos e admin | Victor | M | FND-02 | AC-17, AC-28..AC-30 | proposto |
| FND-06 | Perfil, área e disponibilidade | Victor | G | FND-02, FND-03 | AC-09..AC-16 | proposto |
| FND-07 | Portfólio profissional | Anderson | M | FND-06 | AC-24..AC-27 | proposto |
| FND-08 | Privacidade e direitos | Felipe | G | FND-03, FND-06 | AC-18..AC-23 | proposto |
| FND-09 | Admin, suspensão e auditoria | Felipe | M | FND-03, FND-06 | AC-31..AC-33 | proposto |
| FND-10 | Observabilidade, health e flags | Felipe | M | FND-01 | AC-37..AC-39 | proposto |
| FND-11 | Web: conta | Anderson | G | FND-03, FND-04 | AC-01..AC-08, AC-34..AC-36 | proposto |
| FND-12 | Web: onboarding e perfil | Anderson | G | FND-06, FND-11 | AC-09..AC-16, AC-34..AC-36 | proposto |
| FND-13 | E2E, acessibilidade e evidência | Felipe | M | FND-11, FND-12 | AC-34..AC-37 | proposto |
| FND-14 | OpenAPI, README e handoff | Victor | P | todos | todos | proposto |
| WEB-01 | Design tokens e componentes | Anderson | M | — | AC-01..AC-05 | proposto |
| WEB-02 | Layout, navegação e estados | Anderson | M | WEB-01 | AC-06..AC-10 | proposto |
| WEB-03 | Conta acessível | Anderson | M | WEB-02, FND-11 | AC-11..AC-13 | proposto |
| WEB-04 | Perfil e portfólio | Anderson | M | WEB-02, FND-12, FND-07 | AC-14..AC-16, AC-21 | proposto |

---

### F2 — Descoberta

| ID | Tarefa | Dono | Tam. | Depende de | Aceite | Status |
| --- | --- | --- | --- | --- | --- | --- |
| LCL-01 | Contrato de busca e read model | Victor | M | FND-02, FND-06 | AC-01..AC-08 | proposto |
| LCL-02 | Retrieval | Victor | G | LCL-01 | AC-09..AC-13 | proposto |
| LCL-03 | Ranking baseline | Victor | G | LCL-02 | AC-12..AC-17 | proposto |
| LCL-04 | Explicação e transparência | Anderson | M | LCL-03 | AC-18..AC-20 | proposto |
| LCL-05 | Busca HTTP/HTML e filtros | Victor + Anderson | G | LCL-03 | AC-01..AC-08, AC-21..AC-25 | proposto |
| LCL-06 | Perfil público sem PII | Victor | M | LCL-02 | AC-22, AC-23 | proposto |
| LCL-07 | Eventos de funil | Victor | M | LCL-05 | AC-26..AC-29 | proposto |
| LCL-08 | Guardrails básicos | Victor | M | LCL-03 | AC-16 | proposto |
| LCL-09 | Fallback e performance | Victor | M | LCL-03 | AC-17, AC-30 | proposto |
| LCL-10 | Golden set | Victor | M | LCL-03 | AC-18 | proposto |
| LCL-11 | E2E busca e a11y | Felipe | M | LCL-05, WEB-05 | AC-31, AC-32 | proposto |
| WEB-05 | Busca e resultados | Anderson | G | WEB-02, LCL-05 | AC-17..AC-19, AC-22 | proposto |
| WEB-06 | Perfil público | Anderson | M | WEB-05, FND-07 | AC-20, AC-21 | proposto |

---

### F3 — Demanda e conexão

| ID | Tarefa | Dono | Tam. | Depende de | Aceite | Status |
| --- | --- | --- | --- | --- | --- | --- |
| SRV-01 | Modelo de pedido/interesse/conexão | Anderson | M | FND-02 | AC-01..AC-03 | proposto |
| SRV-02 | Ciclo do pedido | Anderson | G | SRV-01, FND-03 | AC-01..AC-10 | proposto |
| SRV-03 | Feed de oportunidades | Anderson | G | SRV-02, LCL-02 | AC-11..AC-15 | proposto |
| SRV-04 | Interesses e limites | Anderson | G | SRV-03, FND-03 | AC-16..AC-22, AC-34..AC-37 | proposto |
| SRV-05 | Aceite e conexão | Anderson | G | SRV-04 | AC-23..AC-29 | proposto |
| SRV-06 | Contato consentido | Felipe | M | SRV-05, FND-08 | AC-24, AC-25, AC-38, AC-39 | proposto |
| SRV-07 | Notificações e outbox | Anderson | M | SRV-02, SRV-04 | AC-30..AC-33 | proposto |
| SRV-08 | Bloqueio e denúncia | Felipe | M | SRV-05, TRS-01 | AC-28, AC-37 | proposto |
| SRV-09 | Acessibilidade e E2E | Felipe | M | SRV-05, WEB-08 | AC-40 | proposto |
| SRV-10 | Contrato e documentação | Anderson | P | SRV-05 | todos | proposto |
| WEB-07 | Mural, pedido e interesses | Anderson | G | WEB-02, SRV-03, SRV-05 | AC-23..AC-26 | proposto |
| WEB-08 | Conexão, contato e avaliação | Anderson | G | WEB-07, TRS-02, TRS-04 | AC-27..AC-30 | proposto |

---

### F4 — Confiança e reputação

| ID | Tarefa | Dono | Tam. | Depende de | Aceite | Status |
| --- | --- | --- | --- | --- | --- | --- |
| TRS-01 | Modelo de denúncia/ação/auditoria | Felipe | M | FND-02 | AC-21..AC-29 | proposto |
| TRS-02 | Desfecho de conexão | Felipe | M | SRV-05 | AC-01 | proposto |
| TRS-03 | Elegibilidade de avaliação | Felipe | M | TRS-02 | AC-01..AC-03 | proposto |
| TRS-04 | Avaliação bilateral e resposta | Felipe | G | TRS-03 | AC-04..AC-09 | proposto |
| TRS-05 | Reputação contextual | Felipe | G | TRS-04 | AC-10..AC-15 | proposto |
| TRS-06 | Fila de moderação e ações | Felipe | G | TRS-01, FND-09 | AC-21..AC-27, AC-29 | proposto |
| TRS-07 | Recurso e auditoria | Felipe | M | TRS-06 | AC-28 | proposto |
| TRS-08 | Anti-manipulação | Felipe | M | TRS-04 | AC-30..AC-33 | proposto |
| TRS-09 | Verificação de contato | Felipe | M | FND-04 | AC-16..AC-20 | proposto |
| TRS-10 | E2E avaliação/denúncia | Felipe | M | TRS-04, TRS-06, WEB-09 | AC-34..AC-37 | proposto |
| WEB-09 | Denúncia, moderação e admin | Felipe | G | WEB-02, TRS-06, FND-09 | AC-31..AC-34, AC-37 | proposto |
| WEB-10 | Privacidade, consentimento e transparência | Anderson | M | WEB-02, FND-08 | AC-35..AC-37, AC-46 | proposto |
| WEB-11 | Acessibilidade, 360 px e evidência | Felipe | M | WEB-03..WEB-10 | AC-38..AC-44, AC-47..AC-49 | proposto |
| WEB-12 | Copy PT-BR e revisão | Anderson | M | WEB-03..WEB-10 | AC-45, AC-46 | proposto |

---

### F5 — Fairness e experimentação

| ID | Tarefa | Dono | Tam. | Depende de | Aceite | Status |
| --- | --- | --- | --- | --- | --- | --- |
| FAV-01 | Política versionada | Victor | G | LCL-03 | AC-01..AC-06 | proposto |
| FAV-02 | Ledger de exposição | Victor | M | LCL-07 | AC-07..AC-11 | proposto |
| FAV-03 | Caps, exploração, diversidade | Victor | G | FAV-01, FAV-02 | AC-12, AC-13, AC-15..AC-17 | proposto |
| FAV-04 | Boost de novo | Victor | M | FAV-03 | AC-14 | proposto |
| FAV-05 | Métricas de fairness | Victor | G | FAV-02 | AC-24..AC-28 | proposto |
| FAV-06 | Dashboard admin | Anderson | M | FAV-05, WEB-09 | AC-27, AC-29 | proposto |
| FAV-07 | Simulação offline | Victor | G | FAV-01, LCL-10 | AC-34..AC-37 | proposto |
| FAV-08 | Experimentos e guardrails | Victor | G | FAV-01, FAV-05 | AC-19..AC-23 | proposto |
| FAV-09 | Fallback e rollback | Victor | M | FAV-03 | AC-05, AC-17, AC-18 | proposto |
| FAV-10 | Relatório e governança | Victor + Anderson | M | FAV-05, FAV-08 | AC-29..AC-33 | proposto |

---

### F6 — Piloto

| ID | Tarefa | Dono | Depende de | Status |
| --- | --- | --- | --- | --- |
| PIL-01 | Fechar região, categorias e metas | Anderson | specs aprovadas | proposto |
| PIL-02 | DPIA e revisão jurídica | Felipe | PIL-01 | proposto |
| PIL-03 | Política de privacidade e termos | Felipe + Anderson | PIL-02 | proposto |
| PIL-04 | Moderação operante e SLA | Felipe | TRS-06 | proposto |
| PIL-05 | Auditoria de acessibilidade | Felipe | WEB-11 | proposto |
| PIL-06 | Auditoria de segurança | Felipe | todas | proposto |
| PIL-07 | Observabilidade, alertas e runbook | Felipe | FND-10 | proposto |
| PIL-08 | Feature flags e rollback | Victor | todas | proposto |
| PIL-09 | Onboarding assistido de oferta | Anderson | PIL-01 | proposto |
| PIL-10 | Execução do piloto | Todos | PIL-01..PIL-09 | proposto |
| PIL-11 | Relatório de aprendizado | Anderson | PIL-10 | proposto |
| PIL-12 | Decisão de continuidade | Fundador | PIL-11 | proposto |

---

## 4. Caminho crítico

```txt
PLAN-10 (aprovação)
  -> FND-01 -> FND-02 -> FND-03 -> FND-04
  -> FND-06 -> LCL-01 -> LCL-02 -> LCL-03 -> LCL-05
  -> SRV-01 -> SRV-02 -> SRV-03 -> SRV-04 -> SRV-05 -> SRV-06
  -> TRS-01 -> TRS-02 -> TRS-03 -> TRS-04 -> TRS-06
  -> FAV-01 -> FAV-02 -> FAV-03 -> FAV-05 -> FAV-07
  -> PIL-01..PIL-10
```

Acessibilidade, segurança e privacidade correm em paralelo ao caminho crítico, mas bloqueiam o piloto.

---

## 5. Dependências bloqueantes

| Bloqueio | Afeta | Dono | Ação |
| --- | --- | --- | --- |
| Aprovação humana das specs | todos os tickets | Fundador | reunião de aprovação |
| Q-07, Q-44–Q-50 | FND-01..FND-03 | Victor | decidir antes de F1 |
| Q-08, Q-14–Q-18, Q-37 | LCL-01..LCL-05 | Victor + Anderson | decidir antes de F2 |
| Q-25–Q-31 | SRV-01..SRV-05 | Anderson + Felipe | decidir antes de F3 |
| Q-06, Q-09, Q-10, Q-32–Q-36 | TRS-01..TRS-09 | Felipe | decidir antes de F4 |
| Q-19–Q-24 | FAV-01..FAV-08 | Victor + Anderson | decidir antes de F5 |
| Q-01, Q-03, Q-38–Q-43 | PIL-01..PIL-10 | Todos | decidir antes de F6 |

---

## 6. Ordem sugerida por pessoa

### Victor

1. PLAN-01, PLAN-04, PLAN-06, PLAN-09.
2. FND-01, FND-02, FND-03, FND-05, FND-06, FND-04, FND-14.
3. LCL-01, LCL-02, LCL-03, LCL-06, LCL-07, LCL-08, LCL-09, LCL-10.
4. FAV-01, FAV-02, FAV-03, FAV-04, FAV-05, FAV-07, FAV-08, FAV-09, FAV-10.

### Anderson

1. PLAN-02, PLAN-05, PLAN-07.
2. WEB-01, WEB-02, WEB-03, WEB-04, FND-07, FND-11, FND-12.
3. LCL-04, WEB-05, WEB-06.
4. SRV-01..SRV-05, SRV-07, SRV-10, WEB-07, WEB-08.
5. WEB-12, FAV-06, FAV-10 (apoio).

### Felipe

1. PLAN-03, PLAN-08, PLAN-12.
2. FND-08, FND-09, FND-10, FND-13, WEB-11.
3. SRV-06, SRV-08, SRV-09.
4. TRS-01..TRS-10, WEB-09.
5. PIL-02..PIL-07, evidências e release.

---

## 7. Critérios de Ready

- [ ] spec aprovada;
- [ ] ticket com aceite;
- [ ] dependências resolvidas;
- [ ] contrato claro;
- [ ] PII avaliada;
- [ ] flag/rollback definidos;
- [ ] dono e revisor definidos.

## 8. Critérios de Done

- [ ] código revisado;
- [ ] testes verdes;
- [ ] aceite coberto;
- [ ] acessibilidade;
- [ ] segurança/privacidade;
- [ ] fairness quando aplicável;
- [ ] observabilidade;
- [ ] documentação/ADR/OpenAPI;
- [ ] evidência;
- [ ] PR mergeado;
- [ ] board atualizado.

---

## 9. Métricas do board

- tickets por status;
- idade média em review;
- bloqueios por categoria;
- taxa de retrabalho;
- cobertura de aceite;
- tempo de spec → primeira entrega.
