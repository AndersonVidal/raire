# fair-visibility — tickets

**Spec:** `spec.md`  
**Status:** proposto  
**Dono:** Victor  
**Revisores:** Anderson (produto), Felipe (segurança/fairness)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| FAV-01 | Política de fairness versionada | Victor | LCL-03 | G | proposto |
| FAV-02 | Ledger de exposição | Victor | LCL-07 | M | proposto |
| FAV-03 | Caps, exploração e diversidade | Victor | FAV-01, FAV-02 | G | proposto |
| FAV-04 | Boost de novo e decaimento | Victor | FAV-03 | M | proposto |
| FAV-05 | Métricas de fairness | Victor | FAV-02 | G | proposto |
| FAV-06 | Dashboard admin | Anderson | FAV-05, WEB-09 | M | proposto |
| FAV-07 | Simulação offline | Victor | FAV-01, LCL-10 | G | proposto |
| FAV-08 | Experimentos e guardrails | Victor | FAV-01, FAV-05 | G | proposto |
| FAV-09 | Fallback e rollback | Victor | FAV-03 | M | proposto |
| FAV-10 | Relatório e governança | Victor + Anderson | FAV-05, FAV-08 | M | proposto |

---

## Detalhes

### FAV-01 — Política de fairness versionada

**Objetivo:** modelar política, versão e ciclo de vida.  
**Requisitos:** F-01..F-06.  
**Aceite:** AC-01..AC-06.  
**Escopo:**

- schema;
- status;
- imutabilidade;
- aprovação;
- ativação;
- flag.

**Testes:** unit, integração, auditoria.

---

### FAV-02 — Ledger de exposição

**Objetivo:** registrar impressões sem PII.  
**Requisitos:** F-07..F-11.  
**Aceite:** AC-07..AC-11.  
**Escopo:**

- evento;
- pseudonimização;
- caps;
- replay;
- retenção.

**Testes:** integração, privacidade, replay.

---

### FAV-03 — Caps, exploração e diversidade

**Objetivo:** aplicar reranking justo.  
**Requisitos:** F-12, F-13, F-15, F-16, F-17.  
**Aceite:** AC-12, AC-13, AC-15..AC-17.  
**Escopo:**

- caps;
- slots de exploração;
- diversidade;
- piso;
- registro de motivo.

**Testes:** unit, propriedades, fair.

---

### FAV-04 — Boost de novo e decaimento

**Objetivo:** dar chance a novos sem quebrar relevância.  
**Requisitos:** F-14.  
**Aceite:** AC-14.  
**Escopo:**

- coorte novo;
- boost;
- decaimento;
- limite pelo piso.

**Testes:** unit, fair, fixture.

---

### FAV-05 — Métricas de fairness

**Objetivo:** medir distribuição e conversão.  
**Requisitos:** F-24..F-28.  
**Aceite:** AC-24..AC-28.  
**Escopo:**

- Gini;
- top 10/20;
- novos;
- posição por coorte;
- dashboard;
- alertas.

**Testes:** unit, integração, agregação.

---

### FAV-06 — Dashboard admin

**Objetivo:** visualizar políticas, métricas e experimentos.  
**Requisitos:** F-25, F-29.  
**Aceite:** AC-27, AC-29.  
**Escopo:**

- tela;
- filtros;
- acessibilidade;
- sem PII.

**Depende de:** WEB-09.  
**Testes:** E2E, a11y.

---

### FAV-07 — Simulação offline

**Objetivo:** avaliar política antes de ativar.  
**Requisitos:** F-34..F-37.  
**Aceite:** AC-34..AC-37.  
**Escopo:**

- golden set;
- simulador;
- nDCG/MRR/Gini/cobertura;
- relatório.

**Depende de:** LCL-10.  
**Testes:** determinismo, fixture, fair.

---

### FAV-08 — Experimentos e guardrails

**Objetivo:** testar mudanças com segurança.  
**Requisitos:** F-19..F-23.  
**Aceite:** AC-19..AC-23.  
**Escopo:**

- experimento;
- assignment;
- guardrail;
- parada;
- registro.

**Testes:** unit, integração, fair.

---

### FAV-09 — Fallback e rollback

**Objetivo:** garantir disponibilidade e reversibilidade.  
**Requisitos:** F-06, F-17, F-18.  
**Aceite:** AC-05, AC-17, AC-18.  
**Escopo:**

- baseline;
- flag;
- versão anterior;
- alerta.

**Testes:** caos, integração, fair.

---

### FAV-10 — Relatório e governança

**Objetivo:** documentar decisão e transparência.  
**Requisitos:** F-29..F-33.  
**Aceite:** AC-29..AC-33.  
**Escopo:**

- relatório;
- auditoria;
- página pública;
- contestação;
- revisão.

**Testes:** revisão, contrato, a11y.
