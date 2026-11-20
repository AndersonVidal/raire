# service-request — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Anderson  
**Depende de:** foundation-platform aprovada; ADR 0004

---

## 1. Estratégia

Entregar em fatias verticais: modelo, ciclo do pedido, feed, interesse, conexão/contato, notificações, abuso e aceite. Contato é a parte mais sensível; deve entrar com teste de segurança antes de liberar.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Modelo | tabelas, estados, invariantes | migration, transições | foundation |
| 2 | Pedido | criar/editar/fechar/expirar | unit, integração | 1 |
| 3 | Feed | elegibilidade e filtros | integração, contrato | 2, discovery |
| 4 | Interesse | enviar/retirar/limites | unit, integração | 3 |
| 5 | Conexão | aceite/recusa/estados | integração, authz | 4 |
| 6 | Contato | consentimento/mascaramento | segurança, contrato | 5 |
| 7 | Notificação | outbox/retry/histórico | integração, falha | 2, 4 |
| 8 | Abuso | bloqueio/denúncia | integração, moderação | 5 |
| 9 | E2E | jornada completa | E2E, a11y | 5, 6 |

---

## 3. Paralelismo

- Modelo (1) e notificação (7) podem começar após foundation.
- Pedido (2) e feed (3) em sequência.
- Interesse (4) e abuso (8) em paralelo após conexão.
- Contato (6) e notificação (7) em paralelo.
- E2E (9) no fim.

Caminho crítico: modelo → pedido → feed → interesse → conexão → contato → E2E.

---

## 4. Contratos e dados

- migrations;
- outbox;
- idempotência;
- read model público sem PII;
- estados auditáveis.

---

## 5. Feature flags

| Flag | Default | Uso |
| --- | --- | --- |
| `requests.enabled` | off | libera mural |
| `requests.feed` | off | libera feed |
| `requests.contact_release` | off | libera contato |
| `notifications.in_app` | on | histórico in-app |

---

## 6. Migrações e rollback

- forward-only;
- rollback por flag;
- dados de pedido não apagados em rollback;
- expiração por job idempotente.

---

## 7. Observabilidade

- funil pedido → interesse → aceite → contato;
- tempo até primeiro interesse;
- falha de notificação;
- abuso/bloqueio;
- sem PII.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| spam | limites + triagem |
| contato exposto | teste de segurança |
| baixa liquidez | notificação e curadoria |
| PII na descrição | validação |
| aceite falso | confirmação e denúncia |
| notificação falha | outbox/retry |

---

## 9. Critérios de saída

- [ ] fluxo completo;
- [ ] contato protegido;
- [ ] limites;
- [ ] notificações;
- [ ] PII fora;
- [ ] acessibilidade;
- [ ] auditoria;
- [ ] evidência;
- [ ] gates.
