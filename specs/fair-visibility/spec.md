# fair-visibility

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** RF-07, RF-20  
**RNFs:** RNF-16, RNF-13, RNF-09  
**Módulos:** fairness, discovery, analytics, moderation (governança), web (admin/transparência)  
**Depende de:** `local-discovery`; ADR 0005  
**Dono:** Victor  
**Revisores:** Anderson (produto), Felipe (segurança/fairness)  
**Aprovação humana:** pendente (inclui revisão de pesquisa e segurança)

---

## 1. Contexto

A busca baseline já existe. Agora o RAIRE precisa tornar a distribuição de visibilidade um mecanismo explícito, mensurável e governável. Esta spec entrega a camada de fairness, o ledger de exposição, os experimentos e o dashboard de métricas.

**Hipótese central:** é possível reduzir concentração de visibilidade e aumentar a descoberta de profissionais novos sem queda inaceitável de relevância percebida.

---

## 2. Objetivo e não objetivo

**Objetivo:** implementar política de fairness versionada, ledger de exposição, reranking, métricas, experimentos e governança.

**Não objetivo:**

- ML/embeddings no ranking;
- leilão de posição;
- personalização individual profunda;
- fairness de preço (não há preço);
- múltiplas cidades;
- explicabilidade por atribuição complexa.

---

## 3. Princípios e restrições

- nenhum atributo sensível, direto ou indireto;
- coortes observáveis e agregadas;
- fairness não pode violar o piso de relevância;
- nenhuma posição paga;
- toda política é versionada, auditável e reversível;
- mudança de política exige simulação offline, guardrails e revisão humana;
- experimentos têm hipótese, métrica, guardrail, duração e critério de parada;
- o sistema funciona com a camada desligada (fallback baseline).

---

## 4. Requisitos

### A. Política de fairness

- **F-01** Política é objeto versionado com pesos, caps, exploração, boost de novo, diversidade e piso de relevância.
- **F-02** Política publicada é imutável; alteração cria nova versão.
- **F-03** Política tem status `draft`, `simulated`, `approved`, `active`, `retired`.
- **F-04** Ativação exige aprovação registrada de produto, pesquisa e segurança.
- **F-05** Política ativa é aplicada de forma determinística por `policy_version`.
- **F-06** É possível desligar a camada por feature flag sem deploy.

### B. Ledger de exposição

- **F-07** Cada impressão registra perfil, sessão pseudonimizada, posição, política, experimento e timestamp.
- **F-08** Ledger é append-only e não contém PII.
- **F-09** Caps são calculados por sessão e por dia, com janela configurável.
- **F-10** Ledger suporta replay e auditoria de decisões.
- **F-11** Retenção do ledger é definida; dados agregados sobrevivem ao detalhe.

### C. Reranking

- **F-12** Camada aplica caps de exposição por perfil.
- **F-13** Camada reserva slots de exploração para candidatos elegíveis fora do topo baseline.
- **F-14** Camada aplica boost de novo profissional com decaimento por tempo/conexões.
- **F-15** Camada aplica diversidade de bairro/categoria na primeira página.
- **F-16** Nenhum candidato abaixo do piso de relevância entra por fairness.
- **F-17** Camada não altera a ordem quando não há candidatos elegíveis suficientes; registra o motivo.
- **F-18** Resultado é determinístico para a mesma entrada, sessão e política.

### D. Experimentos

- **F-19** Experimentos têm ID, hipótese, variantes, métrica primária, guardrails, tamanho e duração.
- **F-20** Assignment é determinístico por usuário/sessão e resistente a mudança de dispositivo.
- **F-21** Experimento pode ser interrompido por guardrail sem deploy.
- **F-22** Nenhum experimento usa atributo sensível para segmentação.
- **F-23** Resultados são registrados em `docs/experiments` com decisão.

### E. Métricas e dashboard

- **F-24** Métricas: Gini de impressões, share top 10/20, share de novos, posição média por coorte, exposição elegível e conversão.
- **F-25** Dashboard mostra métricas por política e por experimento.
- **F-26** Métricas são calculadas a partir de dados agregados; não expõem usuário.
- **F-27** Alertas disparam quando guardrail é violado ou distribuição desvia do esperado.
- **F-28** Relatório periódico de fairness é gerado e revisado.

### F. Governança e transparência

- **F-29** Admin vê políticas, simulações e métricas; ativação exige aprovação.
- **F-30** Toda mudança gera `audit_event` com ator, política, motivo e correlação.
- **F-31** Página pública de transparência é atualizada quando a política muda de forma relevante.
- **F-32** Profissional pode contestar posição por canal claro.
- **F-33** Explicação do resultado reflete a política ativa e os sinais usados.

### G. Avaliação offline

- **F-34** Simulador reproduz consultas históricas com política candidata.
- **F-35** Simulação calcula relevância (nDCG/MRR), fairness e cobertura.
- **F-36** Simulação é determinística e versionada.
- **F-37** Nenhuma política é ativada sem relatório de simulação.

---

## 5. Regras de negócio e invariantes

- política ativa única por ambiente;
- versão publicada não é editada;
- nenhum atributo sensível no schema;
- caps não podem ser negativos;
- exploração + baseline = 100% dos slots;
- piso de relevância obrigatório;
- assignment estável;
- audit append-only;
- fair não promove abaixo do piso.

---

## 6. Estados e transições

**Policy:**

```txt
draft -> simulated -> approved -> active -> retired
```

**Experiment:**

```txt
draft -> running -> stopped/concluded
```

---

## 7. Contratos e dados

### 7.1 Endpoints

```txt
GET  /v1/admin/fairness/policies
POST /v1/admin/fairness/policies
POST /v1/admin/fairness/policies/{id}/simulate
POST /v1/admin/fairness/policies/{id}/approve
POST /v1/admin/fairness/policies/{id}/activate
GET  /v1/admin/fairness/metrics
GET  /v1/admin/fairness/experiments
POST /v1/admin/fairness/experiments
POST /v1/admin/fairness/experiments/{id}/stop
GET  /v1/public/ranking-transparency
```

### 7.2 Entidades

- `fairness_policies`
- `visibility_events`
- `experiments`, `experiment_assignments`
- `fairness_metrics_daily`
- `audit_events`

### 7.3 Invariantes

- política ativa única por ambiente;
- versão publicada não editada;
- nenhum atributo sensível no schema;
- caps não negativos;
- exploração + baseline = 100%;
- piso de relevância obrigatório;
- assignment estável;
- audit append-only.

---

## 8. Casos de borda

- poucos candidatos elegíveis → preencher com baseline e registrar;
- todos os candidatos acima do cap → não relaxar; retornar menos e registrar;
- novo profissional sem avaliações → prior neutro, boost elegível;
- política falha → fallback baseline + alerta;
- experimento com amostra insuficiente → não concluir;
- guardrail violado → parar experimento e revisar;
- empate de score → hash determinístico;
- categoria/bairro sem diversidade → registrar motivo.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `POLICY_NOT_SIMULATED` | ativar sem simulação | "Simule a política antes de ativar." |
| `POLICY_NOT_APPROVED` | ativar sem aprovação | "A política precisa de aprovação." |
| `POLICY_INVALID` | config inválida | "Confira pesos, caps e piso." |
| `EXPERIMENT_GUARDRAIL` | guardrail violado | "Experimento interrompido por guardrail." |
| `FORBIDDEN` | sem permissão | "Ação restrita à governança." |
| `FAIRNESS_FALLBACK` | camada falhou | "Resultados em modo baseline." |

---

## 10. Privacidade, segurança e fairness

- **PII:** ledger pseudonimizado; sem PII.
- **Authz:** admin/operator com MFA; ativação exige papel.
- **Rate limit:** admin por operador.
- **Fairness:** núcleo da spec.
- **Acessibilidade:** dashboard/transparência acessíveis.

---

## 11. Observabilidade e métricas

- Gini, top 10/20, novos no top 20, posição por coorte;
- conversão por coorte;
- guardrails;
- fallback;
- backlog do ledger;
- alertas.

---

## 12. Fora de escopo

- ML/embeddings;
- leilão de posição;
- personalização profunda;
- fairness de preço;
- múltiplas cidades;
- explicabilidade complexa.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| fairness reduz relevância | abandono | piso, guardrails, experimento |
| gaming de exposição | injustiça | caps, anomalia, logs |
| política opaca | perda de confiança | transparência e explicação |
| feedback loop | concentração | exploração, holdout |
| complexidade prematura | atraso | política simples e isolada |
| uso indevido de coortes | discriminação | proibir atributos sensíveis, revisão |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- política versionada e imutável;
- caps e exploração;
- boost de novo com decaimento;
- piso de relevância;
- determinismo;
- experimento e guardrail;
- métricas e dashboard;
- auditoria e rollback;
- fallback.

---

## 15. Gates de review específicos

- [ ] nenhum atributo sensível;
- [ ] política versionada e imutável;
- [ ] simulação offline anexada;
- [ ] guardrails e critério de parada;
- [ ] determinismo e replay;
- [ ] caps/exploração testados por propriedade;
- [ ] fallback testado;
- [ ] auditoria e rollback;
- [ ] revisão de produto, pesquisa e segurança;
- [ ] transparência atualizada.

---

## 16. Perguntas para o humano

- [ ] Q-19: percentual de exploração (default 20%).
- [ ] Q-20: cap por sessão/dia (default 2/10).
- [ ] Q-21: boost de novo (default 30 dias/5 conexões).
- [ ] Q-22: piso de relevância (default percentil 40).
- [ ] Q-23: guardrail de relevância (default 2 p.p./5%).
- [ ] Q-24: quem aprova política.

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| F-01..F-06 | RF-07 | FAV-01 | AC-01..AC-06 |
| F-07..F-11 | RF-20 | FAV-02 | AC-07..AC-11 |
| F-12..F-18 | RF-07 | FAV-03, FAV-04, FAV-09 | AC-11..AC-18 |
| F-19..F-23 | RF-20 | FAV-08 | AC-19..AC-23 |
| F-24..F-28 | RF-20 | FAV-05, FAV-06 | AC-24..AC-28 |
| F-29..F-33 | RF-07 | FAV-10 | AC-29..AC-33 |
| F-34..F-37 | RNF-16 | FAV-07 | AC-34..AC-37 |
