# local-discovery

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** RF-06, RF-07, RF-20  
**RNFs:** RNF-02, RNF-13, RNF-16  
**Módulos:** discovery, directory (read), catalog (read), fairness (guardrails), platform, web  
**Depende de:** `foundation-platform`; ADR 0002, ADR 0005  
**Dono:** Victor  
**Revisores:** Anderson (copy/UX), Felipe (fairness/privacidade)  
**Aprovação humana:** pendente

---

## 1. Contexto

Com perfis e catálogo publicados, o cliente precisa encontrar profissionais relevantes por serviço e localização. Esta spec entrega a busca, o ranking baseline explicável, o perfil público e a instrumentação de funil.

A camada completa de fairness e experimentos fica em `fair-visibility`; aqui entram guardrails básicos e a estrutura para medição.

**Problema do usuário:** "preciso de um profissional perto de mim e não sei em quem confiar nem por onde começar."

---

## 2. Objetivo e não objetivo

**Objetivo:** permitir busca por termo/categoria/localização, com filtros, ranking determinístico, explicação e perfil público sem PII.

**Não objetivo:**

- camada completa de fairness e experimentos;
- busca semântica/embeddings/LLM;
- personalização individual profunda;
- mapa interativo avançado;
- ordenação paga;
- múltiplas cidades/idiomas;
- recomendação por histórico comportamental.

---

## 3. Princípios e restrições

- busca presencial considera distância e área;
- busca remota não usa distância como sinal;
- nenhum atributo sensível no ranking;
- nenhuma posição paga;
- explicação deriva dos sinais realmente usados;
- localização pública é aproximada;
- resultado vazio é estado de produto, não erro;
- ranking é determinístico por seed e versão de política;
- performance p95 < 500 ms.

---

## 4. Requisitos

### A. Entrada e filtros

- **L-01** Busca aceita termo livre, categoria/subcategoria e localização (cidade/bairro ou coordenada opcional).
- **L-02** Para presencial/híbrido, aceita raio em km com limites configuráveis; remoto ignora raio.
- **L-03** Filtros: tipo de atendimento, disponibilidade, verificação. Filtro inválido retorna erro claro.
- **L-04** Localização é normalizada para área aproximada; coordenada precisa não é persistida em evento analítico.
- **L-05** Busca sem termo e sem categoria retorna sugestões de categorias, não lista aleatória.
- **L-06** Termos vazios, muito longos ou com caracteres de controle são rejeitados ou normalizados.

### B. Retrieval e ranking

- **L-07** Retrieval filtra apenas perfis `published`, ativos, com categoria compatível, tipo compatível e área que contém o ponto (quando presencial/híbrido).
- **L-08** Retrieval limita candidatos para latência previsível.
- **L-09** Score baseline combina correspondência de categoria/texto, distância (presencial), disponibilidade, verificação, reputação contextual, completude e atividade.
- **L-10** Pesos são versionados em configuração; nenhum peso muda em produção sem registro.
- **L-11** Empates são desfeitos por hash determinístico de sessão + perfil.
- **L-12** Reputação usa média bayesiana com prior por categoria; perfil novo não recebe zero.
- **L-13** Busca aplica guardrails básicos: limite de repetição do mesmo perfil na página e diversidade mínima de bairro/categoria.
- **L-14** Se a camada de ranking/fairness falhar, a busca usa fallback baseline e registra alerta; não retorna erro 500.
- **L-15** Ordenação é determinística para a mesma entrada, sessão e versão de política.

### C. Explicação

- **L-16** Cada resultado exibe até 3 motivos curtos derivados dos sinais usados.
- **L-17** Motivos permitidos: "Perto de você", "Disponível hoje", "Verificado", "Bem avaliado em [categoria]", "Novo na plataforma", "Combina com sua busca".
- **L-18** Proibido usar "melhor", "mais confiável", "garantido", "recomendado pela IA" ou urgência falsa.
- **L-19** Página pública "Como o RAIRE ordena os resultados" explica os critérios em linguagem simples.
- **L-20** Explicação não revela pesos internos, dados de outros usuários ou lógica que facilite gaming.

### D. Resultado e perfil público

- **L-21** Resultado mostra nome, categorias, distância aproximada, sinais de confiança, disponibilidade e explicação.
- **L-22** Perfil público mostra bio, serviços, área aproximada, disponibilidade, sinais e avaliações públicas (quando existirem).
- **L-23** Perfil público nunca retorna e-mail, telefone, documento, coordenada precisa ou endereço.
- **L-24** Paginação por cursor; limite configurável; sem offset profundo.
- **L-25** Estado vazio sugere ajustar filtros ou ampliar raio; estado de erro permite tentar novamente.

### E. Instrumentação

- **L-26** Eventos: `search_started`, `search_results_returned`, `profile_impression`, `profile_clicked`, `contact_intent`.
- **L-27** Eventos usam IDs pseudonimizados, `policy_version`, posição e coorte observável; nunca PII.
- **L-28** Consentimento de analytics é respeitado; sem consentimento, eventos são agregados ou descartados.
- **L-29** Logs de busca não contêm termo livre associado a usuário identificável.
- **L-30** Dashboard de funil: busca, resultado, clique, intenção de contato.

### F. Performance e acessibilidade

- **L-31** Busca p95 < 500 ms no piloto; APIs p95 < 300 ms.
- **L-32** Consultas usam índices e paginação por cursor.
- **L-33** Página de resultados funciona em 360 px, teclado e leitor de tela.
- **L-34** Filtros têm rótulos, foco visível e estado aplicado compreensível.
- **L-35** Imagens e componentes respeitam orçamento de performance e `prefers-reduced-motion`.

---

## 5. Regras de negócio e invariantes

- nenhum perfil inelegível aparece;
- nenhuma PII na resposta;
- mesma entrada + sessão + política = mesma ordem;
- fallback não retorna erro;
- evento de impressão só depois de resultado renderizado;
- remoto não usa distância;
- categoria inativa não aparece;
- perfil suspenso sai no próximo request;
- explicação só cita sinais usados.

---

## 6. Estados e transições

A busca não tem estado persistente; o perfil tem:

```txt
draft -> published -> paused -> published
published/paused -> suspended
```

A resposta de busca pode ser:

```txt
results | empty | error | degraded (fallback)
```

---

## 7. Contratos e dados

### 7.1 Endpoints

```txt
GET /v1/search
GET /v1/search/suggestions
GET /v1/professionals/{id}
GET /v1/professionals/{id}/reviews
GET /v1/public/ranking-transparency
```

### 7.2 Resposta de busca

```json
{
  "items": [
    {
      "profile_id": "prf_...",
      "display_name": "Ana Souza",
      "categories": ["eletricista"],
      "distance_label": "até 2 km",
      "service_type": "in_person",
      "availability_label": "Hoje",
      "verification": "contact_verified",
      "rating": { "average": 4.6, "count": 18 },
      "reasons": ["Perto de você", "Disponível hoje"],
      "position": 1
    }
  ],
  "next_cursor": "...",
  "policy_version": "ranking-v1",
  "total_estimate": 42,
  "degraded": false
}
```

### 7.3 Entidades/read models

- `professional_search_document` (derivado de directory/catalog);
- `visibility_events`;
- `ranking_config`;
- `search_audit`.

### 7.4 Invariantes

- nenhum perfil inelegível aparece;
- nenhuma PII na resposta;
- mesma entrada + sessão + política = mesma ordem;
- fallback não retorna erro;
- evento de impressão só depois de resultado renderizado.

---

## 8. Casos de borda

- sem resultados → sugestões + ampliar raio;
- localização não informada em busca presencial → pedir cidade/bairro, não assumir;
- perfil suspenso durante a busca → removido no próximo request;
- categoria desativada → não aparece;
- raio muito grande → limitar e avisar;
- termo ofensivo → busca retorna vazio e orienta denúncia, sem ecoar conteúdo;
- ranking lento → fallback e alerta;
- muitos resultados → paginação;
- coordenada inválida → erro de validação.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `VALIDATION_ERROR` | filtro/termo inválido | "Confira os filtros da busca." |
| `LOCATION_REQUIRED` | presencial sem localização | "Informe a cidade ou bairro." |
| `RADIUS_OUT_OF_RANGE` | raio inválido | "Escolha um raio entre 1 e 50 km." |
| `RATE_LIMITED` | muitas buscas | "Muitas buscas. Aguarde um instante." |
| `SEARCH_DEGRADED` | fallback | "Resultados em modo simplificado." |
| `INTERNAL_ERROR` | erro | "Não foi possível buscar agora. Tente novamente." |

---

## 10. Privacidade, segurança e fairness

- **PII:** nenhuma resposta pública com e-mail, telefone, documento ou coordenada.
- **Authz:** busca pública; perfil público sem dados privados.
- **Rate limit:** 60/min por sessão.
- **Fairness:** guardrails básicos e telemetria; camada completa em `fair-visibility`.
- **Acessibilidade:** L-33..L-35.

---

## 11. Observabilidade e métricas

- eventos de funil;
- latência e taxa de erro;
- taxa de resultado vazio;
- policy version;
- alerta de fallback;
- logs sem termo livre associado a usuário.

---

## 12. Fora de escopo

- camada completa de fairness e experimentos;
- busca semântica/embeddings/LLM;
- personalização profunda;
- mapa interativo avançado;
- ordenação paga;
- múltiplas cidades/idiomas;
- recomendação por histórico.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| ranking concentra visibilidade | injustiça | guardrails e métricas; fairness completa depois |
| busca lenta | abandono | índices, paginação, benchmark |
| explicação enganosa | perda de confiança | derivar de sinais reais, revisão de copy |
| PII em evento | vazamento | pseudonimização, redaction, testes |
| resultado irrelevante | baixa conversão | golden set, avaliação offline |
| gaming de clique | métrica distorcida | coortes, detecção de anomalia |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- busca por categoria + localização;
- busca remota sem distância;
- filtros;
- ranking determinístico;
- explicação;
- fallback;
- PII;
- vazio/erro;
- performance;
- acessibilidade.

---

## 15. Gates de review específicos

- [ ] nenhum atributo sensível no ranking;
- [ ] PII fora da resposta e dos eventos;
- [ ] determinismo testado;
- [ ] fallback testado;
- [ ] performance medida;
- [ ] acessibilidade dos resultados e filtros;
- [ ] explicação revisada por produto/pesquisa;
- [ ] golden set e métricas de relevância.

---

## 16. Perguntas para o humano

- [ ] Q-08: área por raio ou bairro/polígono.
- [ ] Q-14: busca exige localização?
- [ ] Q-15: raio padrão/máximo.
- [ ] Q-16: resultados por página.
- [ ] Q-17: distância exata/faixa/bairro.
- [ ] Q-18: reputação na busca?
- [ ] Q-37: reputação na busca ou só no perfil.
- [ ] Q-49: busca geoespacial isolada.

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| L-01..L-06 | RF-06 | LCL-01, LCL-05 | AC-01..AC-08 |
| L-07..L-15 | RF-06, RF-07 | LCL-02, LCL-03, LCL-08 | AC-09..AC-17 |
| L-16..L-20 | RF-07 | LCL-04 | AC-18..AC-20 |
| L-21..L-25 | RF-06 | LCL-05, LCL-06 | AC-21..AC-25 |
| L-26..L-30 | RF-20 | LCL-07 | AC-26..AC-29 |
| L-31..L-35 | RNF-02, RNF-13 | LCL-09, LCL-11 | AC-30..AC-32 |
