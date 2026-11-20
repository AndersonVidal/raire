# service-request — tickets

**Spec:** `spec.md`  
**Status:** proposto  
**Dono:** Anderson  
**Revisores:** Victor (contratos), Felipe (abuso/PII)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| SRV-01 | Modelo de pedido, interesse e conexão | Anderson | FND-02 | M | proposto |
| SRV-02 | Publicação, edição, fechamento e expiração | Anderson | SRV-01, FND-03 | G | proposto |
| SRV-03 | Feed de oportunidades elegíveis | Anderson | SRV-02, LCL-02 | G | proposto |
| SRV-04 | Interesses, limites e anti-spam | Anderson | SRV-03, FND-03 | G | proposto |
| SRV-05 | Aceite/recusa e conexão | Anderson | SRV-04 | G | proposto |
| SRV-06 | Consentimento e liberação de contato | Felipe | SRV-05, FND-08 | M | proposto |
| SRV-07 | Notificações e outbox | Anderson | SRV-02, SRV-04 | M | proposto |
| SRV-08 | Bloqueio e denúncia na conexão | Felipe | SRV-05, TRS-01 | M | proposto |
| SRV-09 | Acessibilidade e E2E do fluxo | Felipe | SRV-05, WEB-08 | M | proposto |
| SRV-10 | Contrato e documentação | Anderson | SRV-05 | P | proposto |

---

## Detalhes

### SRV-01 — Modelo de pedido, interesse e conexão

**Objetivo:** modelar pedido, interesses e conexões com invariantes.  
**Requisitos:** S-01..S-03, RNF-05.  
**Aceite:** AC-01..AC-03.  
**Escopo:**

- tabelas e constraints;
- estados;
- índices;
- expiração;
- auditoria.

**Testes:** migration, invariantes, transições.

---

### SRV-02 — Publicação, edição, fechamento e expiração

**Objetivo:** ciclo do pedido.  
**Requisitos:** S-01..S-10.  
**Aceite:** AC-01..AC-10.  
**Escopo:**

- criar/editar/fechar/cancelar;
- PII na descrição;
- limite de pedidos;
- expiração;
- moderação de conteúdo.

**Fora de escopo:** feed (SRV-03).  
**Testes:** unit, integração, segurança, rate limit.

---

### SRV-03 — Feed de oportunidades elegíveis

**Objetivo:** profissional vê pedidos compatíveis.  
**Requisitos:** S-11..S-15.  
**Aceite:** AC-11..AC-15.  
**Escopo:**

- elegibilidade por categoria/área;
- paginação;
- filtros;
- identidade resumida do cliente;
- estado vazio.

**Depende de:** LCL-02.  
**Testes:** integração, contrato, a11y.

---

### SRV-04 — Interesses, limites e anti-spam

**Objetivo:** manifestação de interesse com limites.  
**Requisitos:** S-16..S-22, S-33..S-36.  
**Aceite:** AC-16..AC-22, AC-34..AC-37.  
**Escopo:**

- enviar/retirar/recusar;
- um por pedido;
- limites;
- duplicidade;
- abuso;
- bloqueio.

**Testes:** unit, integração, segurança.

---

### SRV-05 — Aceite/recusa e conexão

**Objetivo:** aceite cria conexão e transições.  
**Requisitos:** S-23..S-28.  
**Aceite:** AC-23..AC-29.  
**Escopo:**

- aceitar/recusar;
- conexão `accepted`;
- estados;
- conclusão/contestação;
- auditoria.

**Testes:** integração, máquina de estados, authz.

---

### SRV-06 — Consentimento e liberação de contato

**Objetivo:** contato só após aceite, conforme preferência.  
**Requisitos:** S-24, S-37, S-38.  
**Aceite:** AC-24, AC-25, AC-38, AC-39.  
**Escopo:**

- registrar consentimento;
- derivar contato visível;
- mascarar antes;
- respeitar canal.

**Testes:** segurança, contrato, integração.

---

### SRV-07 — Notificações e outbox

**Objetivo:** notificar eventos transacionais.  
**Requisitos:** S-29..S-32.  
**Aceite:** AC-30..AC-33.  
**Escopo:**

- outbox;
- retry;
- preferências;
- histórico in-app;
- sem PII.

**Testes:** integração, falha/retry, privacidade.

---

### SRV-08 — Bloqueio e denúncia na conexão

**Objetivo:** proteger usuários em qualquer etapa.  
**Requisitos:** S-27, S-36.  
**Aceite:** AC-28, AC-37.  
**Escopo:**

- bloquear usuário;
- denunciar conexão;
- impedir novas interações.

**Depende de:** TRS-01.  
**Testes:** integração, authz, moderação.

---

### SRV-09 — Acessibilidade e E2E do fluxo

**Objetivo:** provar a jornada completa.  
**Requisitos:** S-39..S-41, RNF-10.  
**Aceite:** AC-40.  
**Escopo:**

- E2E pedido → interesse → aceite;
- teclado/360 px;
- axe;
- evidência.

**Depende de:** SRV-05, WEB-08.  
**Testes:** E2E, a11y.

---

### SRV-10 — Contrato e documentação

**Objetivo:** documentar endpoints e fluxos.  
**Requisitos:** RNF-01, RNF-18.  
**Aceite:** todos.  
**Escopo:**

- OpenAPI;
- diagramas de estado;
- runbook;
- handoff.

**Testes:** validação de contrato.
