# web-experience — plano de implementação

**Status:** proposta  
**Spec:** `spec.md`  
**Tickets:** `tickets.md`  
**Dono:** Anderson  
**Depende de:** foundation, local-discovery, service-request, trust-reputation

---

## 1. Estratégia

Construir a experiência por jornada vertical, começando por tokens e layout, depois conta, perfil, busca, pedidos, conexão e confiança. Acessibilidade é testada por jornada, não no fim.

---

## 2. Fatias

| # | Fatia | Entrega | Testes | Depende de |
| --- | --- | --- | --- | --- |
| 1 | Base visual | tokens, componentes, layout | contraste, foco, axe | — |
| 2 | Conta | cadastro, login, confirmação, reset | E2E, a11y | foundation |
| 3 | Perfil | onboarding, prévia, portfólio | E2E, a11y | foundation |
| 4 | Busca | formulário, resultados, perfil público | E2E, a11y, perf | discovery |
| 5 | Pedidos | mural, feed, interesse | E2E, a11y | service-request |
| 6 | Conexão | contato, avaliação | E2E, segurança | service-request + trust |
| 7 | Confiança | denúncia, moderação, admin | E2E, authz | trust |
| 8 | Privacidade | consentimento, exportação, transparência | E2E, a11y | foundation |
| 9 | Evidência | axe, teclado, 360 px, performance | a11y, perf | todas |

---

## 3. Paralelismo

- Base visual (1) e conta (2) podem começar juntas após tokens.
- Perfil (3) e busca (4) em paralelo, com contrato acordado.
- Pedidos (5) e conexão (6) em sequência.
- Confiança (7) e privacidade (8) em paralelo.
- Evidência (9) no fim de cada jornada.

Caminho crítico: base → conta → perfil → busca → pedidos → conexão → evidência.

---

## 4. Contratos e dados

- consome API `/v1` e endpoints HTML;
- sem lógica de negócio no cliente;
- formulários com CSRF;
- estados derivados do servidor;
- sem PII no HTML público.

---

## 5. Feature flags

| Flag | Default | Uso |
| --- | --- | --- |
| `web.new_search` | off | nova busca |
| `web.portfolio` | off | portfólio |
| `web.admin` | off | admin |
| `web.transparency` | on | página de transparência |

---

## 6. Migrações e rollback

- UI não altera schema;
- rollback por flag e versão anterior;
- compatibilidade com API versionada.

---

## 7. Observabilidade

- eventos de funil;
- erros de formulário;
- performance;
- acessibilidade;
- sem PII.

---

## 8. Riscos e mitigação

| Risco | Mitigação |
| --- | --- |
| onboarding longo | etapas curtas + salvamento |
| acessibilidade no fim | teste por jornada |
| contato exposto | teste de contrato |
| performance ruim | orçamento e JS mínimo |
| copy enganosa | revisão de produto |
| responsividade | teste 360 px |

---

## 9. Critérios de saída

- [ ] WCAG 2.2 AA nas jornadas críticas;
- [ ] 360 px sem rolagem horizontal;
- [ ] teclado e leitor de tela;
- [ ] contraste e foco;
- [ ] estados completos;
- [ ] copy revisada;
- [ ] performance dentro do orçamento;
- [ ] evidência anexada;
- [ ] review de produto e QA.
