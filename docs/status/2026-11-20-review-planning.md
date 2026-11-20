# Review de Planejamento — 2026-11-20

**Status:** concluído (planejamento)  
**Escopo:** PRD, docs, specs, tickets, planos e pipeline  
**Fora de escopo:** código, migrations, testes de produto, deploy  
**Veredito:** **foundation aprovada; demais specs em revisão**; implementação da foundation liberada; piloto bloqueado até DPIA/jurídico

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

- PRD v1.0 revisado, aprovado e rastreável;
- 6 specs com requisitos numerados, critérios de aceite, tickets e planos;
- 72 tickets de feature + planejamento/piloto no CSV (96 linhas);
- arquitetura, dados, API, segurança, privacidade e fairness documentados;
- pipeline, RACI, handoffs e sprint plan;
- QA documental sem links quebrados, sem código e sem segredo.

O que continua pendente:

1. revisão das specs que não são a fundação;
2. DPIA/jurídico antes do piloto;
3. região/categorias reais do piloto antes de F6;
4. implementação, testes e evidências da fundação.

---

## 3. Bloqueios e decisões

| ID | Item | Gate | Dono | Status |
| --- | --- | --- | --- | --- |
| B-01 | Aprovação da foundation | produto/arquitetura | Humano | resolvido em 2026-11-20 |
| B-02 | Q-07, Q-44–Q-50 | arquitetura | Victor | decididas |
| B-03 | Q-08, Q-14–Q-18, Q-37 | produto/arquitetura | Victor + Anderson | decididas |
| B-04 | Q-25–Q-31 | produto/segurança | Anderson + Felipe | decididas |
| B-05 | Q-06, Q-09, Q-10, Q-32–Q-36 | segurança/trust | Felipe | decididas |
| B-06 | Q-19–Q-24 | fairness | Victor + Anderson | decididas; ativação depende de simulação |
| B-07 | DPIA e retenção | jurídico/privacidade | Felipe + jurídico | pendente antes do piloto |
| B-08 | Região/categorias reais do piloto | produto | Anderson | pendente antes de F6 |
| B-09 | ADRs 0001–0006 | arquitetura | Humano | aceitos em 2026-11-20 |
| B-10 | SQLite local ≠ Postgres produção | arquitetura/QA | Victor + Felipe | mitigação: testes de contrato + staging |

---

## 4. Achados por gate

> **Atualização 2026-11-20:** as ressalvas de produto, arquitetura, QA e fairness foram resolvidas com as decisões Q-01..Q-50 e a aprovação da foundation. Permanecem apenas região real do piloto e DPIA/jurídico.

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

## 6. O que o humano já decidiu (2026-11-20)

1. Aprovou as recomendações de Q-01..Q-50.
2. Aprovou a `foundation-platform` v1.1.
3. Aceitou os ADRs 0001–0006.
4. Aprovou os parâmetros da política de fairness (ativação condicionada a simulação).
5. Autorizou o início da F1.

## 7. O que continua pendente

1. Revisar e aprovar as demais specs.
2. Definir região/categorias reais do piloto antes de F6.
3. Nomear jurídico/DPO e concluir DPIA antes do piloto.
4. Implementar, testar e revisar a fundação.

## 8. Conclusão

A documentação está pronta e a fundação foi aprovada. Não há código, teste de produto, migration ou deploy. A implementação começa pelo `FND-01`; o piloto continua bloqueado até DPIA/jurídico.
