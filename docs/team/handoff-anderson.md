# Handoff — Anderson

**Papel:** Produto + web + acessibilidade  
**Ownership:** `web-experience`, `service-request`; apoio em `foundation-platform` (portfólio) e `local-discovery` (copy/UX)  
**Revisa:** produto, copy PT-BR, jornadas, acessibilidade, estados de UI  
**Não aprova:** o próprio código; verificação que a plataforma não faz

---

## 1. Leitura obrigatória

1. `AGENTS.md` e `CLAUDE.md`.
2. `docs/product/prd.md` — visão, escopo, RFs, RNFs e decisões.
3. `docs/product/personas-and-journeys.md` — personas e jornadas.
4. `docs/product/metrics-and-experiments.md` — métricas e experimentos.
5. `docs/product/open-questions.md` — questões de produto.
6. `docs/process/pipeline.md` — fluxo, gates e DoR/DoD.
7. `docs/process/accessibility-checklist.md` — checklist WCAG 2.2 AA.
8. `docs/team/raci.md` — decisões e independência.
9. `specs/web-experience/`, `specs/service-request/`, `specs/foundation-platform/` (RF-04).
10. `docs/architecture/architecture-plan.md` — contratos e frontend.

---

## 2. Foco da primeira semana

### Dia 1 — Contexto

- [ ] Ler PRD, personas, jornadas e specs da sua frente.
- [ ] Listar dúvidas de produto, copy e UX.
- [ ] Verificar Q-01, Q-02, Q-08, Q-11, Q-12, Q-14–Q-18, Q-25–Q-31, Q-37.

### Dia 2 — Revisão de jornadas

- [ ] Validar onboarding do profissional em 360 px.
- [ ] Validar busca e comparação do cliente.
- [ ] Validar mural, interesses e conexão.
- [ ] Apontar fricções e estados faltantes.

### Dia 3 — Revisão de acessibilidade

- [ ] Definir requisitos de teclado, foco, rótulos e contraste.
- [ ] Definir estados de carregando, vazio, erro e sucesso.
- [ ] Definir orçamento de performance e `prefers-reduced-motion`.
- [ ] Definir matriz de dispositivos/navegadores.

### Dia 4 — Revisão de produto

- [ ] Validar copy e promessas.
- [ ] Validar critérios de aceite de produto.
- [ ] Validar métricas de sucesso por jornada.
- [ ] Verificar ausência de dark pattern e pay-to-win.

### Dia 5 — Review cruzado e handoff

- [ ] Revisar specs de Victor e Felipe sob ótica de produto/UX.
- [ ] Atualizar task board.
- [ ] Preparar parecer para aprovação humana.
- [ ] Registrar riscos de inclusão e adoção.

---

## 3. Entregáveis da fase de planejamento

- [ ] Parecer de produto/UX das 6 specs.
- [ ] Jornadas críticas revisadas.
- [ ] Design system mínimo definido (tokens, componentes, estados).
- [ ] Matriz de acessibilidade e dispositivos.
- [ ] Copy PT-BR revisada.
- [ ] Métricas de sucesso e instrumentação por jornada.
- [ ] Riscos de inclusão e adoção registrados.

---

## 4. Primeiros tickets após aprovação

| Ordem | Ticket | Entrega |
| --- | --- | --- |
| 1 | WEB-01 | design tokens e componentes acessíveis |
| 2 | WEB-02 | layout base, navegação e estados globais |
| 3 | WEB-03 | jornada de conta acessível |
| 4 | WEB-04 | jornada de perfil/portfólio |
| 5 | SRV-02 | publicação e gestão de pedidos |
| 6 | SRV-03 | feed de oportunidades |

Detalhes no [`task-board.md`](task-board.md) e em `specs/*/tickets.md`.

---

## 5. Responsabilidades de review

- Revisar copy, estados e acessibilidade de todos os PRs de UI.
- Revisar aderência ao PRD e ao escopo.
- Revisar se a jornada funciona em 360 px e por teclado.
- Bloquear merge se houver:
  - promessa enganosa;
  - dark pattern;
  - contato/PII exposto;
  - estado de erro ausente;
  - inacessibilidade em jornada crítica.

---

## 6. Perguntas que você precisa responder

- Q-01: região piloto.
- Q-02: categorias iniciais.
- Q-08: área por raio ou bairro/polígono.
- Q-11: catálogo inicial.
- Q-12: portfólio com imagem?
- Q-14: busca exige localização?
- Q-15: raio padrão/máximo.
- Q-16: resultados por página.
- Q-17: distância exata/faixa/bairro.
- Q-18: reputação na busca?
- Q-25..Q-31: regras de pedido, interesse, conexão e moderação.
- Q-37: reputação na busca ou só no perfil.

Registre a decisão em `docs/product/open-questions.md`.

---

## 7. Critérios de sucesso

- jornadas simples e compreensíveis;
- WCAG 2.2 AA nas jornadas críticas;
- 360 px sem rolagem horizontal;
- copy honesta e sem jargão;
- estados de UI completos;
- nenhum dark pattern;
- nenhum código antes da aprovação humana.
