# Review de Planejamento — 2026-11-20

**Status:** concluído (planejamento)  
**Escopo:** PRD, docs, specs, tickets, planos e pipeline  
**Fora de escopo:** código, migrations, testes de produto, deploy  
**Veredito:** **aprovado com ressalvas para planejamento**; implementação **bloqueada** até aprovação humana

---

## 1. Participantes

| Gate | Revisor | Foco |
| --- | --- | --- |
| Produto | Anderson | escopo, personas, jornadas, copy, métricas |
| Arquitetura | Victor | contextos, dados, contratos, ranking, performance |
| Segurança/privacidade | Felipe | auth, PII, LGPD, abuso, moderação |
| QA | Felipe | aceite, testabilidade, evidência |
| Fairness | Victor + Felipe | política, coortes, guardrails |

> Revisão simulada no harness. A aprovação final é humana.

---

## 2. Resumo

A organização está consistente e navegável:

- PRD v1.0 revisado e rastreável;
- 6 specs com requisitos numerados, critérios de aceite, tickets e planos;
- 67 tickets de feature + planejamento/piloto no CSV;
- arquitetura, dados, API, segurança, privacidade e fairness documentados;
- pipeline, RACI, handoffs e sprint plan;
- QA documental sem links quebrados, sem código e sem segredo.

O que impede implementação:

1. specs não aprovadas por humano;
2. questões bloqueantes sem decisão;
3. ADRs ainda propostos;
4. DPIA/jurídico pendente para piloto;
5. região/categorias do piloto não definidas.

---

## 3. Bloqueios

| ID | Bloqueio | Gate | Dono | Correção esperada |
| --- | --- | --- | --- | --- |
| B-01 | Specs em revisão, nenhuma aprovada | produto/arquitetura | Fundador | registrar aprovação |
| B-02 | Q-07, Q-44–Q-50 sem decisão | arquitetura | Victor | decidir antes de F1 |
| B-03 | Q-08, Q-14–Q-18, Q-37 sem decisão | produto/arquitetura | Victor + Anderson | decidir antes de F2 |
| B-04 | Q-25–Q-31 sem decisão | produto/segurança | Anderson + Felipe | decidir antes de F3 |
| B-05 | Q-06, Q-09, Q-10, Q-32–Q-36 sem decisão | segurança/trust | Felipe | decidir antes de F4 |
| B-06 | Q-19–Q-24 sem decisão | fairness | Victor + Anderson | decidir antes de F5 |
| B-07 | DPIA e retenção não validadas | jurídico/privacidade | Felipe + jurídico | concluir antes do piloto |
| B-08 | Região/categorias do piloto indefinidas | produto | Anderson | decidir antes de F6 |
| B-09 | ADRs 0001–0006 propostos | arquitetura | Victor | aceitar antes de F1 |
| B-10 | SQLite local ≠ Postgres produção | arquitetura/QA | Victor + Felipe | testes de contrato + staging |

---

## 4. Achados por gate

### 4.1 Produto

**Aprovado com ressalvas.**

- O PRD separa MVP de validação e fatia técnica; isso evita escopo infinito.
- As jornadas cobrem os fluxos críticos.
- A copy de transparência está honesta.
- Ressalvas:
  - região e categorias ainda não definidas;
  - alguns limites (interesses, pedidos, aceite) são defaults e precisam de decisão;
  - o piloto precisa de critério de continuidade claro.

### 4.2 Arquitetura

**Aprovado com ressalvas.**

- Monólito modular, contextos e dependências estão claros.
- Modelo de dados cobre as jornadas.
- Ranking/fairness está isolado e testável.
- Ressalvas:
  - SQLite/Postgres exige disciplina de contrato;
  - outbox e idempotência precisam de teste;
  - performance de busca precisa de benchmark;
  - migrations e rollback precisam de teste em banco vazio e com seed.

### 4.3 Segurança/privacidade

**Aprovado com ressalvas para planejamento; bloqueado para piloto até DPIA.**

- Auth, sessão, CSRF, RBAC e MFA estão definidos.
- PII separada; contato mascarado; localização aproximada.
- Rate limits e auditoria estão previstos.
- Ressalvas:
  - bases legais e retenção precisam de jurídico;
  - DPIA não iniciada;
  - moderação manual precisa de SLA e operação;
  - plano de incidente precisa de simulação antes do piloto.

### 4.4 QA

**Aprovado com ressalvas.**

- Toda spec tem matriz de aceite e tipos de teste.
- Acessibilidade, performance, segurança e fairness têm cobertura planejada.
- Ressalvas:
  - nenhum teste executado (correto nesta fase);
  - caminhos de teste ainda são propostas;
  - evidências ainda vazias;
  - é preciso definir dataset sintético e golden set.

### 4.5 Fairness

**Aprovado com ressalvas.**

- Política versionada, coortes observáveis, caps, exploração, boost, diversidade e piso.
- Guardrails e fallback definidos.
- Ressalvas:
  - defaults (20%, 2/10, 30 dias/5 conexões, percentil 40) precisam de validação;
  - golden set não existe;
  - simulação offline precisa ser implementada antes de ativar política;
  - aprovação humana obrigatória.

---

## 5. Sugestões

1. Aprovar primeiro `foundation-platform` e `web-experience`; depois discovery e service-request.
2. Resolver Q-07 e Q-44–Q-50 antes de qualquer código.
3. Manter fairness básica em F2 e camada completa em F5.
4. Criar o golden set cedo, junto com a busca baseline.
5. Rodar E2E e acessibilidade a cada jornada, não no fim.
6. Validar SQLite/Postgres com testes de contrato desde F1.
7. Fazer simulação de incidente antes do piloto.
8. Reduzir o escopo de F1 se o tempo apertar: conta + perfil + catálogo + admin mínimo.

---

## 6. O que o humano deve decidir

1. Aprovar ou rejeitar as 6 specs.
2. Responder as questões bloqueantes.
3. Aceitar ou revisar os defaults de fairness.
4. Confirmar a estratégia SQLite/Postgres.
5. Confirmar o escopo da fatia técnica 1.
6. Autorizar (ou não) o início de F1.
7. Definir região/categorias do piloto antes de F6.
8. Nomear jurídico/DPO para LGPD.

---

## 7. Conclusão

A documentação está pronta para revisão humana. Não há código, teste de produto, migration ou deploy. A implementação deve começar somente após aprovação das specs e decisão das questões bloqueantes.
