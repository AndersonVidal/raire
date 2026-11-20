# ADR 0005 — Ranking explicável com camada de fairness

**Status:** proposto  
**Data:** 2026-11-20  
**Decisores:** produto, arquitetura, pesquisa, segurança  
**Relacionado:** `docs/architecture/ranking-and-fairness.md`, `docs/product/prd.md` (seção 9)

## Contexto

O diferencial do RAIRE é descoberta justa. Um ranking puramente por popularidade concentra visibilidade e pune profissionais novos. Um ranking puramente por fairness pode destruir relevância. É preciso um desenho explícito, mensurável e reversível.

## Opções consideradas

1. **Baseline + camada de fairness versionada** — score relevante, caps, exploração, boost de novo, diversidade, piso de relevância, explicação.
2. **Ranking cronológico/aleatório** — justo, mas irrelevante.
3. **Ranking por popularidade** — relevante, mas concentra e exclui novos.
4. **ML/embeddings desde o início** — opaco, caro e sem dados.

## Decisão

Adotar pipeline:

1. elegibilidade;
2. retrieval;
3. score baseline com pesos versionados;
4. bandas de relevância;
5. camada de fairness (caps, exploração, boost, diversidade, piso);
6. explicação;
7. telemetria;
8. fallback baseline.

Nenhum atributo sensível. Nenhuma posição paga. Política versionada, simulada offline, aprovada por humano e reversível por flag.

## Consequências

**Positivas:**

- diferencial de produto explícito;
- auditável e testável;
- protege profissionais novos;
- permite experimentos com guardrails.

**Negativas:**

- complexidade maior que ranking simples;
- exige métricas e governança;
- risco de reduzir relevância.

**Mitigações:**

- piso de relevância;
- guardrails de CTR/nDCG;
- fallback;
- simulação offline;
- revisão humana.

## Critérios de reversão

- queda de relevância além do guardrail;
- complexidade operacional inviável;
- evidência de que fairness não melhora descoberta.

## Status

Proposto. Requer aprovação humana antes da implementação.
