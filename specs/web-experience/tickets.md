# web-experience — tickets

**Spec:** `spec.md`  
**Status:** proposto  
**Dono:** Anderson  
**Revisores:** Victor (contratos), Felipe (a11y/QA)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| WEB-01 | Design tokens e componentes acessíveis | Anderson | — | M | proposto |
| WEB-02 | Layout base, navegação e estados globais | Anderson | WEB-01 | M | proposto |
| WEB-03 | Jornada de conta acessível | Anderson | WEB-02, FND-11 | M | proposto |
| WEB-04 | Jornada de perfil e portfólio | Anderson | WEB-02, FND-12, FND-07 | M | proposto |
| WEB-05 | Busca e resultados acessíveis | Anderson | WEB-02, LCL-05 | G | proposto |
| WEB-06 | Perfil público e portfólio | Anderson | WEB-05, FND-07 | M | proposto |
| WEB-07 | Mural, pedido e interesses | Anderson | WEB-02, SRV-03, SRV-05 | G | proposto |
| WEB-08 | Conexão, contato e avaliação | Anderson | WEB-07, TRS-02, TRS-04 | G | proposto |
| WEB-09 | Denúncia, moderação e admin | Felipe | WEB-02, TRS-06, FND-09 | G | proposto |
| WEB-10 | Privacidade, consentimento e transparência | Anderson | WEB-02, FND-08 | M | proposto |
| WEB-11 | Acessibilidade, 360 px e evidência | Felipe | WEB-03..WEB-10 | M | proposto |
| WEB-12 | Copy PT-BR e revisão de produto | Anderson | WEB-03..WEB-10 | M | proposto |

---

## Detalhes

### WEB-01 — Design tokens e componentes acessíveis

**Objetivo:** base visual consistente e acessível.  
**Requisitos:** W-01..W-07.  
**Aceite:** AC-01..AC-05.  
**Escopo:**

- tokens de cor, tipografia, espaçamento, foco, borda e movimento;
- componentes: botão, link, input, select, textarea, checkbox, radio, alert, badge, card, paginação, dialog, skeleton;
- estados;
- documentação curta.

**Fora de escopo:** tema escuro completo.  
**Testes:** contraste, foco, teclado, axe.  
**DoD:** componentes documentados e testados.

---

### WEB-02 — Layout base, navegação e estados globais

**Objetivo:** estrutura de página acessível e consistente.  
**Requisitos:** W-08..W-14.  
**Aceite:** AC-06..AC-10.  
**Escopo:**

- skip link;
- header/footer/main;
- nav;
- título e idioma;
- alertas globais;
- estados de sessão/erro/offline.

**Testes:** HTML semântico, landmarks, teclado, axe.

---

### WEB-03 — Jornada de conta acessível

**Objetivo:** cadastro, login, confirmação e reset.  
**Requisitos:** W-15, W-16.  
**Aceite:** AC-11..AC-13.  
**Escopo:**

- formulários com validação inline;
- mensagens específicas;
- CSRF;
- estados;
- link de confirmação em dev.

**Depende de:** FND-11.  
**Testes:** E2E, a11y, 360 px.

---

### WEB-04 — Jornada de perfil e portfólio

**Objetivo:** onboarding e edição do perfil.  
**Requisitos:** W-17, W-18.  
**Aceite:** AC-14..AC-16, AC-21.  
**Escopo:**

- etapas curtas;
- progresso;
- salvamento;
- prévia;
- edição/pausa;
- portfólio.

**Depende de:** FND-12, FND-07.  
**Testes:** E2E, a11y, 360 px.

---

### WEB-05 — Busca e resultados acessíveis

**Objetivo:** busca e resultados compreensíveis.  
**Requisitos:** W-19..W-21.  
**Aceite:** AC-17..AC-19, AC-22.  
**Escopo:**

- formulário de busca;
- filtros;
- cards;
- explicação;
- estado vazio/erro;
- paginação.

**Depende de:** LCL-05.  
**Testes:** E2E, a11y, performance.

---

### WEB-06 — Perfil público e portfólio

**Objetivo:** perfil público informativo e sem PII.  
**Requisitos:** W-22.  
**Aceite:** AC-20, AC-21.  
**Escopo:**

- bio, categorias, área aproximada;
- disponibilidade;
- sinais;
- portfólio;
- ação de interesse.

**Depende de:** FND-07, LCL-06.  
**Testes:** contrato, a11y, PII.

---

### WEB-07 — Mural, pedido e interesses

**Objetivo:** cliente publica pedido e vê interessados; profissional vê feed.  
**Requisitos:** W-23..W-26.  
**Aceite:** AC-23..AC-26.  
**Escopo:**

- formulário de pedido;
- lista/detalhe;
- feed;
- interesse;
- estados.

**Depende de:** SRV-03, SRV-05.  
**Testes:** E2E, a11y, 360 px.

---

### WEB-08 — Conexão, contato e avaliação

**Objetivo:** conexão consentida, contato e avaliação.  
**Requisitos:** W-27, W-28.  
**Aceite:** AC-27..AC-30.  
**Escopo:**

- estado da conexão;
- contato mascarado/liberado;
- próxima ação;
- avaliação/resposta.

**Depende de:** TRS-02, TRS-04.  
**Testes:** E2E, segurança (contato), a11y.

---

### WEB-09 — Denúncia, moderação e admin

**Objetivo:** denúncia e operação.  
**Requisitos:** W-29, W-30.  
**Aceite:** AC-31..AC-34, AC-37.  
**Escopo:**

- formulário de denúncia;
- fila de denúncias;
- auditoria;
- fairness;
- notificações.

**Depende de:** TRS-06, FND-09.  
**Testes:** E2E, authz, a11y.

---

### WEB-10 — Privacidade, consentimento e transparência

**Objetivo:** usuário entende e controla seus dados.  
**Requisitos:** W-31, W-32.  
**Aceite:** AC-35..AC-37, AC-46.  
**Escopo:**

- página de privacidade;
- consentimentos;
- exportação;
- exclusão;
- transparência de ranking.

**Depende de:** FND-08.  
**Testes:** E2E, a11y, copy.

---

### WEB-11 — Acessibilidade, 360 px e evidência

**Objetivo:** provar acessibilidade e responsividade.  
**Requisitos:** W-33..W-41, W-48..W-52.  
**Aceite:** AC-38..AC-44, AC-47..AC-49.  
**Escopo:**

- axe em páginas críticas;
- teclado;
- leitor de tela;
- zoom;
- reduced motion;
- performance;
- evidência.

**Depende de:** WEB-03..WEB-10.  
**Testes:** axe, manual, Lighthouse, screenshots.

---

### WEB-12 — Copy PT-BR e revisão de produto

**Objetivo:** linguagem clara e honesta.  
**Requisitos:** W-42..W-47.  
**Aceite:** AC-45, AC-46.  
**Escopo:**

- revisão de copy;
- estados de erro;
- transparência;
- remoção de jargão;
- revisão de promessas.

**Testes:** revisão de produto, checklist de copy.
