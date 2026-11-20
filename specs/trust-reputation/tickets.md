# trust-reputation — tickets

**Spec:** `spec.md`  
**Status:** proposto  
**Dono:** Felipe  
**Revisores:** Victor (dados), Anderson (copy/jornada)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| TRS-01 | Modelo de denúncia, ação e auditoria | Felipe | FND-02 | M | proposto |
| TRS-02 | Desfecho de conexão | Felipe | SRV-05 | M | proposto |
| TRS-03 | Elegibilidade de avaliação | Felipe | TRS-02 | M | proposto |
| TRS-04 | Avaliação bilateral e resposta | Felipe | TRS-03 | G | proposto |
| TRS-05 | Reputação contextual e sinais | Felipe | TRS-04 | G | proposto |
| TRS-06 | Fila de moderação e ações | Felipe | TRS-01, FND-09 | G | proposto |
| TRS-07 | Recurso e auditoria | Felipe | TRS-06 | M | proposto |
| TRS-08 | Anti-manipulação de avaliações | Felipe | TRS-04 | M | proposto |
| TRS-09 | Verificação de contato | Felipe | FND-04 | M | proposto |
| TRS-10 | E2E de avaliação, denúncia e moderação | Felipe | TRS-04, TRS-06, WEB-09 | M | proposto |

---

## Detalhes

### TRS-01 — Modelo de denúncia, ação e auditoria

**Objetivo:** modelar reports, actions, appeals e audit.  
**Requisitos:** T-23..T-31, RNF-14.  
**Aceite:** AC-21..AC-29.  
**Escopo:**

- tabelas e constraints;
- prioridade;
- status;
- auditoria append-only;
- recurso.

**Testes:** migration, invariantes, authz.

---

### TRS-02 — Desfecho de conexão

**Objetivo:** registrar desfecho declarado.  
**Requisitos:** T-01, RF-13.  
**Aceite:** AC-01.  
**Escopo:**

- `connection_outcomes`;
- tipos;
- autoria;
- divergência;
- elegibilidade.

**Depende de:** SRV-05.  
**Testes:** unit, integração, authz.

---

### TRS-03 — Elegibilidade de avaliação

**Objetivo:** determinar quando avaliar.  
**Requisitos:** T-01, T-02.  
**Aceite:** AC-01..AC-03.  
**Escopo:**

- conexão completed;
- contestação;
- uma por lado;
- autoavaliação.

**Testes:** unit, propriedades, integração.

---

### TRS-04 — Avaliação bilateral e resposta

**Objetivo:** criar, editar e responder avaliações.  
**Requisitos:** T-03..T-08.  
**Aceite:** AC-04..AC-09.  
**Escopo:**

- nota/tags/comentário;
- PII;
- janela de edição;
- resposta;
- publicação.

**Testes:** unit, integração, segurança, a11y.

---

### TRS-05 — Reputação contextual e sinais

**Objetivo:** calcular e exibir reputação por contexto.  
**Requisitos:** T-09..T-15.  
**Aceite:** AC-10..AC-15.  
**Escopo:**

- média bayesiana;
- distribuição;
- recência;
- sinais;
- snapshot recalculável.

**Testes:** unit, fixture, revisão de fairness.

---

### TRS-06 — Fila de moderação e ações

**Objetivo:** operar denúncias com prioridade e auditoria.  
**Requisitos:** T-23..T-28, T-31.  
**Aceite:** AC-21..AC-27, AC-29.  
**Escopo:**

- fila;
- análise;
- ações;
- comunicação;
- auditoria.

**Depende de:** TRS-01, FND-09.  
**Testes:** integração, authz, auditoria.

---

### TRS-07 — Recurso e auditoria

**Objetivo:** permitir contestação e revisão independente.  
**Requisitos:** T-29, T-30.  
**Aceite:** AC-28.  
**Escopo:**

- appeal;
- prazo;
- revisor diferente;
- decisão.

**Testes:** integração, authz, auditoria.

---

### TRS-08 — Anti-manipulação de avaliações

**Objetivo:** detectar padrões abusivos.  
**Requisitos:** T-32..T-36.  
**Aceite:** AC-30..AC-33.  
**Escopo:**

- autoavaliação;
- repetição;
- conta nova;
- denúncia falsa;
- sinalização para revisão.

**Testes:** unit, integração, fair.

---

### TRS-09 — Verificação de contato

**Objetivo:** selo honesto de contato.  
**Requisitos:** T-16..T-22.  
**Aceite:** AC-16..AC-20.  
**Escopo:**

- confirmação de e-mail/telefone;
- status;
- selo;
- expiração/revogação;
- documento fora do primeiro ciclo.

**Depende de:** FND-04.  
**Testes:** integração, segurança, contrato.

---

### TRS-10 — E2E de avaliação, denúncia e moderação

**Objetivo:** provar a jornada de confiança.  
**Requisitos:** T-39..T-42, RNF-10.  
**Aceite:** AC-34..AC-37.  
**Escopo:**

- E2E conexão → avaliação;
- denúncia → moderação → recurso;
- teclado/360 px/axe;
- evidência.

**Depende de:** TRS-04, TRS-06, WEB-09.  
**Testes:** E2E, a11y, segurança.
