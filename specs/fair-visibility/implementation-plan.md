# fair-visibility — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Victor  
**Depende de:** local-discovery aprovada; ADR 0005 aceito

---

## 1. Estratégia

Entregar fairness em camadas: política/ledger, reranking, métricas, simulação, experimentos, fallback e governança. Nenhuma política é ativada sem simulação, guardrail e aprovação humana.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Política | schema, versão, status | unit, auditoria | discovery |
| 2 | Ledger | impressão, pseudonimização, caps | integração, privacidade | 1 |
| 3 | Reranking | caps, exploração, diversidade | unit, propriedades | 1, 2 |
| 4 | Boost | novo, decaimento | unit, fair | 3 |
| 5 | Métricas | Gini, top, novos | unit, integração | 2 |
| 6 | Dashboard | tela admin | E2E, a11y | 5 |
| 7 | Simulação | golden set, offline | determinismo | 1, discovery |
| 8 | Experimentos | assignment, guardrail | unit, fair | 1, 5 |
| 9 | Fallback | flag, versão anterior | caos | 3 |
| 10 | Governança | relatório, transparência | revisão | 5, 8 |

---

## 3. Paralelismo

- Política (1) e ledger (2) em sequência.
- Reranking (3) e métricas (5) em paralelo após ledger.
- Boost (4) e dashboard (6) em paralelo.
- Simulação (7) e experimentos (8) em paralelo.
- Fallback (9) e governança (10) no fim.

Caminho crítico: política → ledger → reranking → métricas → simulação → governança.

---

## 4. Contratos e dados

- política versionada;
- ledger append-only;
- métricas agregadas;
- experiment assignment;
- auditoria.

---

## 5. Feature flags

| Flag | Default | Uso |
| --- | --- | --- |
| `discovery.fairness.enabled` | off | ativa camada |
| `discovery.fairness.policy_version` | v1 | política ativa |
| `discovery.experiments.enabled` | off | experimentos |

---

## 6. Migrações e rollback

- política nova sem alterar histórico;
- ledger retido;
- rollback por flag e versão anterior;
- simulação não altera produção.

---

## 7. Observabilidade

- Gini, top, novos, posição;
- guardrails;
- fallback;
- experimentos;
- alertas.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| relevância cai | piso + guardrail |
| gaming | caps + anomalia |
| política opaca | transparência |
| feedback loop | exploração |
| complexidade | política simples |
| atributo sensível | schema proibido + review |

---

## 9. Critérios de saída

- [ ] política versionada;
- [ ] ledger sem PII;
- [ ] caps/exploração/boost/diversidade;
- [ ] métricas;
- [ ] simulação;
- [ ] experimento com guardrail;
- [ ] fallback;
- [ ] governança;
- [ ] evidência;
- [ ] gates.
