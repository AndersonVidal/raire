# trust-reputation — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Felipe  
**Depende de:** foundation + service-request aprovadas

---

## 1. Estratégia

Entregar confiança em fatias: modelo de denúncia/auditoria, desfecho, elegibilidade, avaliação, reputação, moderação, recurso, anti-manipulação e verificação. Segurança e moderação entram cedo.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Modelo | reports/actions/appeals/audit | migration, invariantes | foundation |
| 2 | Desfecho | outcome de conexão | unit, integração | service-request |
| 3 | Elegibilidade | regra de avaliação | unit, propriedades | 2 |
| 4 | Avaliação | criar/editar/responder | unit, integração, segurança | 3 |
| 5 | Reputação | cálculo contextual | unit, fixture | 4 |
| 6 | Moderação | fila/ações/comunicação | integração, authz | 1, admin |
| 7 | Recurso | appeal/revisão | integração | 6 |
| 8 | Anti-manipulação | detecção/sinalização | unit, integração | 4 |
| 9 | Verificação | selo de contato | integração, segurança | foundation |
| 10 | E2E | avaliação/denúncia | E2E, a11y | 4, 6 |

---

## 3. Paralelismo

- Modelo (1) e verificação (9) podem começar após foundation.
- Desfecho (2) e avaliação (4) em sequência.
- Reputação (5) e moderação (6) em paralelo após avaliação.
- Recurso (7) e anti-manipulação (8) em paralelo.
- E2E (10) no fim.

Caminho crítico: modelo → desfecho → elegibilidade → avaliação → moderação → E2E.

---

## 4. Contratos e dados

- auditoria append-only;
- denúncia sem expor denunciante;
- review pública sem PII;
- reputação recalculável;
- retenção.

---

## 5. Feature flags

| Flag | Default | Uso |
| --- | --- | --- |
| `reviews.enabled` | off | libera avaliações |
| `moderation.manual_queue` | on | fila manual |
| `verification.contact` | off | selo de contato |
| `reputation.contextual` | on | reputação por categoria |

---

## 6. Migrações e rollback

- forward-only;
- rollback por flag;
- remoção de conteúdo mantém auditoria;
- expurgo por job.

---

## 7. Observabilidade

- cobertura de avaliação;
- distribuição de notas;
- denúncias por 100 conexões;
- tempo de triagem/resolução;
- recursos;
- alertas de manipulação.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| avaliação falsa | elegibilidade + detecção |
| retaliação | anonimato + janela |
| remoção de crítica | critério + auditoria |
| documento exposto | cifra + retenção |
| moderação lenta | prioridade + SLA |
| viés | contexto + prior neutro |

---

## 9. Critérios de saída

- [ ] elegibilidade;
- [ ] avaliação bilateral;
- [ ] reputação contextual;
- [ ] denúncia/moderação/recurso;
- [ ] anti-manipulação;
- [ ] PII protegida;
- [ ] acessibilidade;
- [ ] evidência;
- [ ] gates.
