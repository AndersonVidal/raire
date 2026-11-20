# Métricas, Validação e Experimentação — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0, seção 11  
**Uso:** orientar instrumentação, dashboards, experimentos e critérios de continuidade do piloto.

---

## 1. Objetivos da medição

- Saber se o RAIRE gera **conexões úteis**, não apenas cliques.
- Avaliar se a distribuição de visibilidade é mais justa sem perder relevância.
- Detectar riscos de segurança, privacidade e manipulação.
- Aprender com o piloto e decidir continuidade com evidência.
- Nunca usar métrica para justificar violação de privacidade ou fairness.

---

## 2. Hierarquia de métricas

```txt
Métrica central
  └── Conexões consentidas de qualidade por semana
      ├── Oferta e ativação
      ├── Descoberta e relevância
      ├── Demanda e interesse
      ├── Conexão e desfecho
      ├── Confiança e reputação
      ├── Fairness
      ├── Segurança e privacidade
      └── Acessibilidade e usabilidade
```

---

## 3. Métrica central

**Conexões consentidas de qualidade por semana na região piloto.**

Definição operacional:

- conexão aceita por ambas as partes;
- pelo menos um desfecho registrado como "contato útil" ou "serviço realizado";
- sem denúncia procedente na conexão;
- cliente e profissional não bloqueados.

Métricas de apoio: relevância percebida, distribuição de exposição, incidentes.

---

## 4. Métricas por dimensão

### 4.1 Oferta e participação profissional

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Perfis completos publicados | perfis com categoria, área/tipo e disponibilidade | 150–300 |
| Tempo até primeiro perfil publicado | mediana do cadastro à publicação | ≤ 1 dia |
| Taxa de conclusão do onboarding | perfis publicados / cadastros de profissional | ≥ 60% |
| Profissionais ativos | com login ou oportunidade vista em 30 dias | ≥ 70% dos publicados |
| Cobertura de categorias | categorias com ≥ 5 profissionais | 100% das categorias piloto |

### 4.2 Descoberta

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Buscas com resultado relevante | busca com ≥ 1 resultado elegível | ≥ 70% |
| Busca → perfil | sessões de busca com clique em perfil | ≥ 25% |
| Tempo até primeiro resultado | p95 da busca | ≤ 500 ms |
| Relevância percebida | pesquisa pós-busca (1–5) | ≥ 4,0 |
| Busca vazia | buscas sem resultado | ≤ 30% |
| Uso de filtros | buscas com ao menos um filtro | monitorar |

### 4.3 Demanda

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Pedidos publicados | pedidos ativos por semana | monitorar |
| Pedidos com ≥ 3 interesses | pedidos que recebem 3+ interesses | ≥ 50% |
| Tempo até primeiro interesse | mediana | < 24 h |
| Pedidos expirados sem interesse | pedidos que expiram sem interesse | ≤ 30% |
| Taxa de fechamento | pedidos fechados / publicados | ≥ 60% |

### 4.4 Conexão e desfecho

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Interesse → conexão aceita | interesses aceitos / enviados | ≥ 30% |
| Conexão → contato | conexões com contato registrado | ≥ 70% |
| Conexão → serviço realizado | desfecho "serviço realizado" | ≥ 40% |
| Conexão → avaliação | conexões com ao menos uma avaliação | ≥ 20% |
| Tempo até conexão | mediana do interesse ao aceite | ≤ 48 h |

### 4.5 Confiança e reputação

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Cobertura de avaliações | conexões concluídas com avaliação | ≥ 20% no primeiro ciclo |
| Distribuição de notas | histograma por categoria | monitorar |
| Taxa de resposta do avaliado | respostas / avaliações elegíveis | ≥ 20% |
| Denúncias procedentes | denúncias confirmadas / conexões | < 2% |
| Contas com e-mail confirmado | confirmadas / cadastradas | ≥ 80% |
| Perfis com selo de contato | verificados / publicados | ≥ 40% (quando disponível) |

### 4.6 Fairness

| Métrica | Definição | Meta inicial |
| --- | --- | --- |
| Gini de impressões | concentração de exposição entre elegíveis | redução vs. baseline |
| Share top 10/20 | exposição dos 10/20 mais exibidos por sessão/dia | monitorar; evitar concentração |
| Novos no top 20 | perfis novos no top 20 | ≥ 20% |
| Posição média por coorte | média de posição por coorte observável | monitorar |
| Conversão por coorte | impressão → clique → interesse → conexão | sem queda injustificada |
| Relevância vs. fairness | guardrail de CTR/nDCG | queda ≤ 2 p.p. / 5% |

**Coortes permitidas:** novo, poucas avaliações, baixa atividade, região/bairro, categoria, tipo de serviço. **Proibido:** atributos sensíveis ou proxies.

### 4.7 Segurança, privacidade e moderação

| Métrica | Definição | Meta |
| --- | --- | --- |
| Incidentes de privacidade | exposição indevida confirmada | 0 |
| Vazamento de contato antes do aceite | casos confirmados | 0 |
| Denúncias por 100 conexões | volume | monitorar |
| Tempo de triagem | mediana até primeira análise | ≤ 48 h |
| Tempo de resolução | mediana até decisão | ≤ 5 dias |
| Recursos procedentes | reversões / recursos | monitorar |
| Contas suspensas | por fraude/abuso | monitorar |

### 4.8 Acessibilidade e usabilidade

| Métrica | Definição | Meta |
| --- | --- | --- |
| Jornadas críticas WCAG 2.2 AA | auditoria | 100% |
| Conclusão em 360 px | teste de usabilidade | ≥ 90% |
| Conclusão por teclado | teste manual | 100% das jornadas críticas |
| Erros de formulário | taxa de erro por campo | monitorar |
| Abandono no onboarding | etapas perdidas | reduzir vs. baseline |
| Tempo de tarefa | busca, pedido, interesse | dentro do alvo de UX |

### 4.9 Sustentabilidade operacional

| Métrica | Definição | Meta |
| --- | --- | --- |
| Custo de infraestrutura | mensal | dentro do orçamento |
| Custo por conexão | infra / conexões | monitorar |
| Custo de moderação | horas / casos | monitorar |
| Disponibilidade | capacidades críticas | 99,5% |
| Erros 5xx | requisições | < 0,5% |

---

## 5. Eventos de instrumentação

Nomes propostos (sem PII):

```txt
account_signup_started
account_signup_completed
email_confirmed
profile_created
profile_published
profile_updated
search_started
search_results_returned
profile_impression
profile_clicked
request_created
request_closed
request_expired
interest_sent
interest_accepted
interest_declined
connection_created
connection_contacted
connection_completed
outcome_recorded
review_submitted
review_published
report_created
moderation_action_taken
```

**Regras:**

- IDs pseudonimizados; nunca e-mail, telefone, documento ou coordenada precisa.
- `policy_version`, posição, coorte observável e `correlation_id` quando aplicável.
- Consentimento de analytics respeitado; sem consentimento, agregar ou descartar.
- Evento de impressão só depois de resultado renderizado.
- Logs de busca não associam termo livre a usuário identificável.

---

## 6. Experimentos prioritários

| ID | Hipótese | Desenho | Métrica primária | Guardrails | Critério de sucesso |
| --- | --- | --- | --- | --- | --- |
| EXP-01 | boost de novos aumenta diversidade sem perder relevância | A/B | novos no top 20 | CTR, satisfação | +20% novos; queda ≤ 2 p.p. |
| EXP-02 | explicação do ranking aumenta confiança e clique | A/B | clique no perfil + confiança | relevância | +10% clique; confiança maior |
| EXP-03 | distância é decisiva em serviços presenciais | análise de funil | conversão por faixa | relevância | queda > 30% acima de X km |
| EXP-04 | mural gera conexões de qualidade | piloto 2 categorias | pedidos com 3+ interesses; aceite | spam, denúncias | ≥ 50% e ≥ 30% |
| EXP-05 | reputação contextual melhora decisão | A/B | confiança e arrependimento | privacidade | melhora sem expor PII |

### 6.1 Protocolo de experimento

Todo experimento deve ter:

1. hipótese falsificável;
2. população e amostra;
3. variantes e assignment determinístico;
4. métrica primária e secundárias;
5. guardrails e critério de parada;
6. duração e tamanho mínimo;
7. revisão de privacidade e fairness;
8. registro em `docs/experiments/`;
9. decisão final documentada.

**Nenhum experimento começa sem guardrail e revisão humana.**

---

## 7. Metas do piloto e critérios de continuidade

### 7.1 Metas iniciais

| Dimensão | Meta |
| --- | --- |
| Oferta | 150–300 perfis completos |
| Descoberta | ≥ 70% das buscas com resultado relevante |
| Demanda | ≥ 50% dos pedidos com 3+ interesses |
| Conexão | ≥ 30% de interesses aceitos |
| Confiança | ≥ 20% das conexões com avaliação |
| Fairness | Gini menor que baseline; ≥ 20% de novos no top 20 |
| Segurança | 0 incidentes críticos de privacidade |
| Acessibilidade | 100% das jornadas críticas WCAG 2.2 AA |

### 7.2 Critérios de continuidade

Continuar se:

- conexões de qualidade crescerem por 3 semanas consecutivas;
- relevância percebida se mantiver ≥ 4,0;
- fairness dentro dos guardrails;
- nenhum incidente crítico sem correção;
- operação de moderação sustentável.

Revisar/reduzir se:

- liquidez não atingir o mínimo;
- concentração de visibilidade piorar;
- custo operacional inviável;
- risco de segurança/privacidade inaceitável.

### 7.3 Critérios de parada

- vazamento de localização precisa ou contato;
- ranking com atributo sensível;
- moderação incapaz de responder a risco físico;
- fraude generalizada;
- impossibilidade de cumprir LGPD.

---

## 8. Governança das métricas

- Definição versionada em `docs/product/metrics-and-experiments.md`.
- Toda métrica tem dono, fonte, fórmula e limitação.
- Dashboard agregado, sem PII.
- Revisão mensal de métricas e guardrails.
- Nenhuma decisão de produto baseada em métrica isolada.
- Resultados de experimento registrados com decisão e aprendizado.
