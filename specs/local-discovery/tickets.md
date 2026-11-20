# local-discovery — tickets

**Spec:** `spec.md`  
**Status:** proposto  
**Dono:** Victor  
**Revisores:** Anderson (copy/UX), Felipe (fairness/privacidade)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| LCL-01 | Contrato de busca e read model | Victor | FND-02, FND-06 | M | proposto |
| LCL-02 | Retrieval por categoria, área e tipo | Victor | LCL-01 | G | proposto |
| LCL-03 | Ranking baseline e pesos versionados | Victor | LCL-02 | G | proposto |
| LCL-04 | Explicação e página de transparência | Anderson | LCL-03 | M | proposto |
| LCL-05 | Busca HTTP/HTML e filtros | Victor + Anderson | LCL-03 | G | proposto |
| LCL-06 | Perfil público sem PII | Victor | LCL-02 | M | proposto |
| LCL-07 | Eventos de funil e métricas | Victor | LCL-05 | M | proposto |
| LCL-08 | Guardrails básicos de fairness | Victor | LCL-03 | M | proposto |
| LCL-09 | Fallback e performance | Victor | LCL-03 | M | proposto |
| LCL-10 | Golden set e avaliação offline | Victor | LCL-03 | M | proposto |
| LCL-11 | E2E busca e acessibilidade | Felipe | LCL-05, WEB-05 | M | proposto |

---

## Detalhes

### LCL-01 — Contrato de busca e read model

**Objetivo:** definir contrato de entrada/saída e read model de busca.  
**Requisitos:** L-01..L-06, RNF-02.  
**Aceite:** AC-01..AC-08.  
**Escopo:**

- query params;
- resposta;
- read model derivado;
- índices;
- paginação por cursor.

**Fora de escopo:** ranking.  
**Testes:** contrato, validação, paginação.

---

### LCL-02 — Retrieval por categoria, área e tipo

**Objetivo:** filtrar apenas perfis elegíveis.  
**Requisitos:** L-07..L-10.  
**Aceite:** AC-09..AC-13.  
**Escopo:**

- categoria ativa;
- tipo de serviço;
- área/raio;
- status;
- limite de candidatos.

**Testes:** integração, bordas, performance.

---

### LCL-03 — Ranking baseline e pesos versionados

**Objetivo:** ordenar por relevância de forma determinística.  
**Requisitos:** L-09..L-15.  
**Aceite:** AC-12..AC-17.  
**Escopo:**

- score;
- pesos versionados;
- bayesiana;
- tie-break;
- determinismo;
- fallback.

**Testes:** unitários, propriedades, fixtures.

---

### LCL-04 — Explicação e página de transparência

**Objetivo:** explicar por que o profissional aparece.  
**Requisitos:** L-16..L-20.  
**Aceite:** AC-18..AC-20.  
**Escopo:**

- até 3 motivos;
- copy permitida;
- página pública;
- revisão de produto.

**Depende de:** LCL-03.  
**Testes:** unitário, copy, contrato.

---

### LCL-05 — Busca HTTP/HTML e filtros

**Objetivo:** expor busca na API e web.  
**Requisitos:** L-01..L-06, L-21..L-25.  
**Aceite:** AC-01..AC-08, AC-21..AC-25.  
**Escopo:**

- endpoint `/v1/search`;
- página `/buscar`;
- filtros;
- estado vazio/erro;
- paginação.

**Depende de:** LCL-03, WEB-02.  
**Testes:** HTTP/E2E, a11y, performance.

---

### LCL-06 — Perfil público sem PII

**Objetivo:** exibir perfil público completo e seguro.  
**Requisitos:** L-22, L-23.  
**Aceite:** AC-22, AC-23.  
**Escopo:**

- bio, categorias, área;
- disponibilidade;
- portfólio;
- sinais;
- nenhuma PII.

**Depende de:** LCL-02.  
**Testes:** contrato, segurança, a11y.

---

### LCL-07 — Eventos de funil e métricas

**Objetivo:** medir descoberta sem PII.  
**Requisitos:** L-26..L-30.  
**Aceite:** AC-26..AC-29.  
**Escopo:**

- eventos;
- pseudonimização;
- policy_version;
- consentimento;
- dashboard.

**Depende de:** LCL-05.  
**Testes:** integração, privacidade, agregação.

---

### LCL-08 — Guardrails básicos de fairness

**Objetivo:** evitar concentração na primeira página.  
**Requisitos:** L-13, L-15.  
**Aceite:** AC-16.  
**Escopo:**

- diversidade de bairro/categoria;
- limite de repetição;
- telemetria.

**Depende de:** LCL-03.  
**Testes:** propriedade, fair.

---

### LCL-09 — Fallback e performance

**Objetivo:** busca resiliente e rápida.  
**Requisitos:** L-14, L-31, L-32.  
**Aceite:** AC-17, AC-30.  
**Escopo:**

- fallback;
- índices;
- benchmark;
- alerta.

**Testes:** caos leve, benchmark.

---

### LCL-10 — Golden set e avaliação offline

**Objetivo:** medir relevância antes de mudar ranking.  
**Requisitos:** RNF-16, L-10.  
**Aceite:** AC-18.  
**Escopo:**

- consultas;
- julgamento;
- nDCG/MRR;
- relatório.

**Testes:** determinismo, fixture.

---

### LCL-11 — E2E busca e acessibilidade

**Objetivo:** provar a jornada de busca.  
**Requisitos:** L-33..L-35, RNF-10.  
**Aceite:** AC-31, AC-32.  
**Escopo:**

- E2E busca;
- teclado;
- 360 px;
- axe;
- evidência.

**Depende de:** LCL-05, WEB-05.  
**Testes:** E2E, a11y.
