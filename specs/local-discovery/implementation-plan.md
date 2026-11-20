# local-discovery — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Victor  
**Depende de:** foundation-platform aprovada; ADR 0002 e ADR 0005 aceitos

---

## 1. Estratégia

Entregar busca em fatias verticais: contrato/read model, retrieval, ranking, explicação, HTTP/web, telemetria, guardrails e performance. Fairness completa fica para `fair-visibility`.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Contrato/read model | query, resposta, índices | contrato, paginação | foundation |
| 2 | Retrieval | filtros de elegibilidade | integração, bordas | 1 |
| 3 | Ranking | score, pesos, tie-break | unit, propriedades | 2 |
| 4 | Explicação | motivos, transparência | copy, contrato | 3 |
| 5 | Busca web/API | endpoint e página | E2E, a11y | 3 |
| 6 | Perfil público | perfil sem PII | contrato, segurança | 2 |
| 7 | Telemetria | eventos, dashboard | integração, privacidade | 5 |
| 8 | Guardrails | diversidade, fallback | fair, caos | 3 |
| 9 | Golden set | avaliação offline | determinismo | 3 |
| 10 | Aceite | E2E, a11y, perf | E2E, perf | 5–9 |

---

## 3. Paralelismo

- Contrato/read model (1) e perfil público (6) podem começar após foundation.
- Ranking (3) e explicação (4) em sequência.
- Busca web (5) e telemetria (7) em paralelo após ranking.
- Guardrails (8) e golden set (9) em paralelo.
- Aceite (10) no fim.

Caminho crítico: contrato → retrieval → ranking → busca → aceite.

---

## 4. Contratos e dados

- read model sem PII;
- índices para categoria/área/status;
- paginação por cursor;
- policy version;
- eventos pseudonimizados.

---

## 5. Feature flags

| Flag | Default | Uso |
| --- | --- | --- |
| `discovery.enabled` | off | libera busca |
| `discovery.fairness.enabled` | off | ativa fairness |
| `discovery.explanations` | on | explicação por resultado |
| `discovery.remote_search` | off | busca remota |

---

## 6. Migrações e rollback

- read model derivado, sem migration destrutiva;
- índices adicionados em migration separada;
- rollback por flag e política anterior;
- seed de busca sintético.

---

## 7. Observabilidade

- latência p50/p95;
- taxa de erro;
- resultado vazio;
- fallback;
- funil;
- política ativa.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| concentração de visibilidade | guardrails + métricas |
| busca lenta | índices + benchmark |
| PII em evento | pseudonimização + review |
| ranking opaco | explicação + transparência |
| relevância baixa | golden set + experimento |
| fallback silencioso | alerta + degradação explícita |

---

## 9. Critérios de saída

- [ ] busca presencial e remota;
- [ ] ranking determinístico;
- [ ] explicação;
- [ ] fallback;
- [ ] PII fora;
- [ ] performance;
- [ ] acessibilidade;
- [ ] golden set;
- [ ] evidência;
- [ ] gates.
