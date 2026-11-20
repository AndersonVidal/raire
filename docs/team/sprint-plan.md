# Plano de Sprints — RAIRE

**Status:** proposta  
**Versão:** 1.0  
**Equipe:** Victor, Anderson, Felipe  
**Premissa:** 2 semanas por sprint, dedicação parcial. Ajustar após aprovação das specs.

---

## Sprint 0 — Revisão e aprovação (1 semana)

**Objetivo:** sair de documentação para specs aprovadas.

| Pessoa | Trabalho | Saída |
| --- | --- | --- |
| Victor | revisar foundation, discovery, fairness; decidir Q-07/Q-44–Q-50 | parecer + decisões |
| Anderson | revisar personas, web, service-request; decidir Q-08/Q-11–Q-18/Q-25–Q-31 | parecer + copy |
| Felipe | revisar riscos, segurança, trust; decidir Q-06/Q-09/Q-10/Q-32–Q-43 | parecer + checklist |

**Saída:** specs revisadas, questões bloqueantes com default, pedido de aprovação humana.

---

## Sprint 1 — Fundação (F1)

**Objetivo:** conta + perfil publicável.

| Pessoa | Tickets |
| --- | --- |
| Victor | FND-01, FND-02, FND-03, FND-05, FND-06 |
| Anderson | WEB-01, WEB-02, WEB-03, FND-07, FND-11 |
| Felipe | FND-08, FND-09, FND-10 |

**Saída:** cadastro → confirmação → perfil → preview; admin suspende; PII protegida; testes verdes.

---

## Sprint 2 — Perfil e descoberta (F1/F2)

**Objetivo:** perfil completo + busca explicável.

| Pessoa | Tickets |
| --- | --- |
| Victor | FND-04, FND-14, LCL-01, LCL-02, LCL-03, LCL-06 |
| Anderson | FND-12, WEB-04, LCL-04, WEB-05, WEB-06 |
| Felipe | FND-13, WEB-11 (início), LCL-09, LCL-11 |

**Saída:** busca presencial/remota, ranking determinístico, perfil público sem PII, a11y básica.

---

## Sprint 3 — Demanda e conexão (F3)

**Objetivo:** mural → interesse → aceite → contato.

| Pessoa | Tickets |
| --- | --- |
| Victor | LCL-07, LCL-08, LCL-10 (apoio) |
| Anderson | SRV-01, SRV-02, SRV-03, SRV-04, SRV-05, SRV-07, SRV-10, WEB-07, WEB-08 |
| Felipe | SRV-06, SRV-08, SRV-09, WEB-11 |

**Saída:** fluxo completo, contato protegido, notificações, E2E.

---

## Sprint 4 — Confiança (F4)

**Objetivo:** desfecho → avaliação → reputação → denúncia → moderação.

| Pessoa | Tickets |
| --- | --- |
| Victor | revisão de dados/reputação, apoio a TRS-05 |
| Anderson | WEB-09 (apoio), WEB-12 |
| Felipe | TRS-01..TRS-10, WEB-09 |

**Saída:** avaliação bilateral, reputação contextual, denúncia, fila de moderação, recurso, E2E.

---

## Sprint 5 — Fairness (F5)

**Objetivo:** política + ledger + métricas + simulação.

| Pessoa | Tickets |
| --- | --- |
| Victor | FAV-01, FAV-02, FAV-03, FAV-04, FAV-05, FAV-07, FAV-08, FAV-09, FAV-10 |
| Anderson | FAV-06, FAV-10 (apoio), revisão de transparência |
| Felipe | revisão de fairness/segurança, evidências |

**Saída:** política versionada, métricas, simulação, experimento com guardrail, fallback.

---

## Sprint 6 — Piloto (F6)

**Objetivo:** gates verdes e preparação do piloto.

| Pessoa | Trabalho |
| --- | --- |
| Victor | performance, migrations, flags, rollback |
| Anderson | acessibilidade, copy, onboarding assistido |
| Felipe | segurança, DPIA, moderação, runbook, evidências |

**Saída:** checklist de release, evidências, aprovação humana.

---

## Regras de sprint

- Nenhum ticket entra sem DoR.
- Nenhum ticket sai sem DoD.
- Bloqueio é escalado no daily.
- Capacidade reservada para review e correções (20%).
- Acessibilidade e segurança não são “sobra” de sprint.
- Se um ticket G não couber, quebrar em P.
