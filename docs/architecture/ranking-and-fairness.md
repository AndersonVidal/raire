# Ranking e Fairness — RAIRE

**Status:** proposta para revisão humana  
**Versão:** 1.0  
**Dono técnico:** Victor  
**Aprovadores:** produto + pesquisa + segurança  
**Relacionado:** `docs/product/prd.md` (seção 9), `docs/architecture/architecture-plan.md`, ADR 0005

---

## 1. Objetivo

Definir como a descoberta do RAIRE ordena profissionais de forma **relevante, explicável, auditável e justa**, sem posição paga e sem atributos sensíveis.

Fairness aqui é operacional:

- reduzir invisibilidade extrema;
- reduzir concentração excessiva de exposição;
- preservar relevância para o cliente;
- permitir que profissionais novos sejam descobertos;
- manter tudo mensurável, versionado e reversível.

---

## 2. Pipeline de descoberta

```txt
1. Entrada e normalização
2. Elegibilidade (hard filters)
3. Retrieval de candidatos
4. Score baseline
5. Ranking por bandas de relevância
6. Camada de fairness (caps, exploração, boost, diversidade)
7. Explicação
8. Paginação e resposta
9. Telemetria de exposição
```

### 2.1 Entrada e normalização

- termo livre normalizado (trim, lowercase, limites);
- categoria/subcategoria do catálogo;
- localização aproximada (cidade/bairro ou coordenada opcional);
- raio (presencial/híbrido); remoto ignora raio;
- filtros: tipo de atendimento, disponibilidade, verificação;
- sessão pseudonimizada para determinismo e ledger.

### 2.2 Elegibilidade (hard filters)

Um perfil só entra se:

- está `published`;
- não está suspenso;
- tem categoria compatível ativa;
- tem tipo de serviço compatível;
- atende a área/raio quando presencial/híbrido;
- não foi bloqueado pelo usuário;
- não está excluído/anonimizado.

**Elegibilidade não é ranking.** Nenhum perfil elegível é excluído por falta de avaliações.

### 2.3 Retrieval

- filtra por categoria, tipo, status e área;
- limita candidatos para latência previsível;
- usa índice geoespacial (PostGIS em produção; Haversine no adaptador local);
- paginação por cursor.

### 2.4 Score baseline

Score proposto (v1), com pesos versionados:

```txt
score = 0.40 * category_match
      + 0.20 * text_match
      + 0.15 * proximity
      + 0.10 * reputation
      + 0.05 * availability
      + 0.05 * verification
      + 0.05 * profile_completeness
```

| Sinal | Definição | Peso | Observação |
| --- | --- | --- | --- |
| `category_match` | 1 se categoria exata; 0,5 se subcategoria/sinônimo | 0,40 | principal sinal de relevância |
| `text_match` | similaridade simples de termo/bio/serviço | 0,20 | sem embeddings no MVP |
| `proximity` | 1 - distância/raio, limitado a [0,1] | 0,15 | zero para remoto |
| `reputation` | média bayesiana contextual | 0,10 | prior neutro para novo |
| `availability` | 1 se disponível hoje; 0,5 se na semana; 0 caso contrário | 0,05 | declarado pelo profissional |
| `verification` | 1 se identidade; 0,7 se contato; 0 se não | 0,05 | honesto e explicável |
| `profile_completeness` | campos preenchidos / campos esperados | 0,05 | não punir excessivamente |

**Regras:**

- pesos são configuração versionada;
- nenhum peso muda em produção sem registro;
- empates são desfeitos por hash determinístico de `session_seed + profile_id`;
- reputação usa média bayesiana com prior por categoria; perfil novo não recebe zero;
- `text_match` não usa atributo pessoal;
- `profile_completeness` tem peso baixo para não excluir P3.

### 2.5 Bandas de relevância

Antes de fairness, os candidatos são agrupados em bandas:

- banda A: score ≥ 0,75;
- banda B: 0,55–0,75;
- banda C: 0,35–0,55;
- banda D: < 0,35 (elegível, mas fora da primeira página).

A camada de fairness **não promove** candidato abaixo do piso de relevância (default: percentil 40). Ela escolhe dentro das bandas elegíveis.

---

## 3. Camada de fairness

### 3.1 Objetivos

- evitar que o mesmo perfil domine a primeira página;
- dar chance de descoberta a profissionais novos ou com pouco histórico;
- manter diversidade de bairro/categoria;
- preservar relevância e confiança;
- ser determinística, auditável e reversível.

### 3.2 Mecanismos

| Mecanismo | Default | Regra |
| --- | --- | --- |
| Cap por sessão | 2 impressões por perfil | excedente vai para depois ou é omitido |
| Cap por dia | 10 impressões por perfil | janela diária |
| Exploração | 20% dos slots da primeira página | candidatos elegíveis fora do topo baseline |
| Boost de novo | 30 dias ou 5 conexões | decaimento linear; nunca acima do piso de relevância |
| Diversidade | mínimo de 2 bairros e 2 categorias na primeira página | quando houver candidatos |
| Piso de relevância | percentil 40 | fairness não promove abaixo do piso |
| Fallback | baseline sem fairness | se a camada falhar ou flag off |

### 3.3 Algoritmo proposto

```txt
1. baseline = ordenar(candidatos, score)
2. elegiveis_fairness = candidatos com score >= piso
3. topo = primeiros N (ex.: 10)
4. reservar 20% dos slots para exploracao:
   - escolher de coortes sub-representadas (novo, poucas avaliacoes)
   - respeitar relevancia e diversidade
5. preencher o restante com baseline, aplicando caps
6. aplicar diversidade de bairro/categoria
7. se faltar candidato elegivel, preencher com baseline e registrar motivo
8. reordenar por banda + score dentro dos slots
9. gerar explicacao por resultado
```

### 3.4 Coortes observáveis

Permitidas:

- profissional novo (≤ 30 dias ou ≤ 5 conexões);
- poucas avaliações (< 3);
- baixa atividade recente (sem conexão em 30 dias);
- região/bairro;
- categoria;
- tipo de serviço.

**Proibidas:** raça, gênero, idade, deficiência, religião, orientação sexual, inferência ou proxy.

### 3.5 Explicação

Cada resultado pode exibir até 3 motivos derivados dos sinais usados:

- "Perto de você";
- "Disponível hoje";
- "Contato verificado";
- "Bem avaliado em [categoria]";
- "Novo na plataforma";
- "Combina com sua busca".

**Proibido:** "melhor", "mais confiável", "garantido", "recomendado pela IA", urgência falsa.

A página "Como o RAIRE ordena os resultados" explica os critérios em linguagem simples, sem revelar pesos internos que facilitem gaming.

---

## 4. Política de fairness

### 4.1 Estrutura

```json
{
  "version": "fairness-v1",
  "status": "draft|simulated|approved|active|retired",
  "weights": { "category_match": 0.4, "text_match": 0.2, "proximity": 0.15, "reputation": 0.1, "availability": 0.05, "verification": 0.05, "completeness": 0.05 },
  "caps": { "per_session": 2, "per_day": 10 },
  "exploration_pct": 0.2,
  "new_boost_days": 30,
  "new_boost_connections": 5,
  "relevance_floor_percentile": 40,
  "diversity": { "min_neighborhoods": 2, "min_categories": 2 },
  "guardrails": { "ctr_drop_pp": 2, "ndcg_drop_pct": 5 }
}
```

### 4.2 Regras

- política publicada é imutável; alteração cria nova versão;
- ativação exige aprovação registrada de produto, pesquisa e segurança;
- política ativa é única por ambiente;
- aplicação determinística por `policy_version`;
- flag `discovery.fairness.enabled` desliga sem deploy;
- rollback para versão anterior;
- toda mudança gera `audit_event`.

### 4.3 Ciclo de vida

```txt
draft -> simulated -> approved -> active -> retired
```

- `draft`: em edição;
- `simulated`: relatório offline gerado;
- `approved`: aprovada por humano;
- `active`: em produção;
- `retired`: substituída.

---

## 5. Ledger de exposição

Cada impressão registra:

- `profile_id`;
- `session_hash` (pseudonimizado);
- `position`;
- `policy_version`;
- `experiment_id` (quando aplicável);
- `cohort` observável;
- `query_hash`;
- timestamp UTC.

**Regras:**

- append-only;
- sem PII;
- retenção definida;
- dados agregados sobrevivem ao detalhe;
- suporta replay e auditoria.

---

## 6. Métricas

| Métrica | Definição | Meta |
| --- | --- | --- |
| Gini de impressões | concentração entre elegíveis | redução vs. baseline |
| Share top 10/20 | exposição dos mais exibidos | monitorar |
| Share de novos | novos no top 20 | ≥ 20% |
| Posição média por coorte | média de posição | monitorar |
| Conversão por coorte | impressão → clique → interesse → conexão | sem queda injustificada |
| Relevância | CTR, nDCG, satisfação | guardrail |
| Cobertura | perfis elegíveis exibidos | monitorar |
| Diversidade | bairros/categorias na primeira página | ≥ 2/2 quando possível |

### 6.1 Guardrails

- queda de CTR ≤ 2 p.p. vs. baseline;
- queda de nDCG ≤ 5%;
- satisfação do cliente não cai;
- Gini melhora ou não piora;
- novos no top 20 ≥ 20%;
- nenhum incidente de privacidade.

### 6.2 Critério de parada

Se guardrail violado:

1. desligar flag `discovery.fairness.enabled`;
2. voltar política anterior;
3. registrar incidente;
4. revisar antes de reativar.

---

## 7. Experimentos

Todo experimento tem:

- ID, hipótese e variantes;
- assignment determinístico por usuário/sessão;
- métrica primária e secundárias;
- guardrails;
- tamanho e duração;
- critério de parada;
- revisão de privacidade e fairness;
- registro em `docs/experiments`.

**Proibido:** segmentar por atributo sensível; usar clique isolado como sucesso; rodar sem guardrail.

---

## 8. Avaliação offline

### 8.1 Golden set

- consultas representativas por categoria e região;
- julgamento de relevância por produto/pesquisa;
- coortes balanceadas;
- casos de concentração e de poucos candidatos;
- dados sintéticos.

### 8.2 Simulador

- reproduz consultas históricas com política candidata;
- calcula nDCG/MRR, Gini, cobertura e share de novos;
- é determinístico e versionado;
- gera relatório antes de ativação.

### 8.3 Critério de ativação

Nenhuma política é ativada sem:

- relatório de simulação;
- guardrails dentro do limite;
- revisão humana;
- plano de monitoramento e rollback.

---

## 9. Governança

| Decisão | Aprovadores |
| --- | --- |
| Criar política | Victor |
| Simular | Victor |
| Aprovar | Anderson + pesquisa + Felipe |
| Ativar | Fundador (ou delegado) |
| Desativar por incidente | Felipe (segurança) |
| Aposentar | Victor + Anderson |

Toda decisão gera `audit_event` com ator, política, motivo e correlação.

---

## 10. Transparência

- página pública "Como o RAIRE ordena os resultados";
- explicação curta por resultado;
- canal de contestação;
- registro versionado de mudanças;
- relatório periódico de fairness;
- comunicação de mudanças relevantes.

---

## 11. Fora de escopo do MVP

- ML/embeddings no ranking;
- leilão de posição;
- personalização individual profunda;
- fairness de preço (não há preço);
- múltiplas cidades;
- explicabilidade por atribuição complexa;
- otimização automática sem revisão humana.

---

## 12. Critérios de aceite

- [ ] Nenhum atributo sensível no schema, feature ou proxy.
- [ ] Política versionada e imutável.
- [ ] Caps, exploração, boost, diversidade e piso testados.
- [ ] Determinismo e replay.
- [ ] Fallback testado.
- [ ] Ledger append-only sem PII.
- [ ] Métricas e guardrails definidos.
- [ ] Simulação offline anexada.
- [ ] Aprovação humana registrada.
- [ ] Transparência atualizada.
- [ ] Rollback testado.
