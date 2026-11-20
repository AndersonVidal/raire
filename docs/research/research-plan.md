# Plano de Pesquisa — RAIRE

**Status:** proposta  
**Versão:** 1.0  
**Dono:** Anderson  
**Apoio:** Felipe (ética/privacidade), Victor (dados/ranking)

---

## 1. Objetivo

Reduzir incerteza sobre o problema, a proposta de valor e os mecanismos de descoberta antes de investir em produto. Pesquisa não é marketing; é evidência para decisão.

---

## 2. Perguntas prioritárias

1. A dor de descoberta é intensa e frequente?
2. Profissionais com pouca presença digital adotariam a plataforma?
3. Clientes confiariam em profissionais fora do círculo social?
4. Fairness importa para o cliente? E para o profissional?
5. Qual densidade local mínima gera conexões?
6. O mural de necessidade resolve o cold start?
7. Reputação contextual é compreensível?
8. Qual o custo operacional de moderação?
9. Quais barreiras de acessibilidade são críticas?
10. Qual modelo de sustentabilidade não conflita com fairness?

---

## 3. Métodos

| Método | Uso | Amostra inicial |
| --- | --- | --- |
| Entrevista semiestruturada | dor, canais, confiança | 5 clientes + 5 profissionais por categoria |
| Teste de usabilidade | onboarding, busca, conexão | 3–5 por persona |
| Auditoria de plataformas | visibilidade e fairness | 3–5 concorrentes |
| Análise de funil | hipóteses de conversão | dados do protótipo/piloto |
| Experimento controlado | fairness, explicação, reputação | ver protocolo |
| Pesquisa quantitativa | satisfação e confiança | pós-jornada |

---

## 4. Ética e privacidade

- Consentimento informado.
- Nenhum dado real identificável em nota pública.
- Entrevistas gravadas só com autorização.
- Direito de retirar consentimento.
- Dados agregados e anonimizados.
- Não usar atributo sensível para segmentação.
- Revisão de Felipe antes de experimento.

---

## 5. Mapa hipótese → estudo

| Hipótese | Estudo | Evidência |
| --- | --- | --- |
| H1/H2 | entrevistas + auditoria | profissionais invisíveis, concentração |
| H3 | teste A/B | clique, confiança |
| H4/H5 | entrevistas + teste de perfil | decisão, compreensão |
| H6 | teste de selo | confiança |
| H7 | funil por distância | conversão |
| H8 | entrevistas + dados | frustração |
| H9 | comparação presencial/remoto | diferença de funil |
| H10 | piloto de pedidos | interesses/conexões |
| H11 | uso do feed | adoção |
| H12 | teste de usabilidade | abandono |
| H13/H14 | golden set + simulação | relevância/fairness |
| H15 | experimento | diversidade |
| H16/H17 | comparação com baseline | necessidade de IA |

---

## 6. Critérios de evidência

- **Forte:** triangulação de métodos + amostra adequada + replicação.
- **Média:** um método robusto + amostra razoável.
- **Fraca:** opinião isolada, amostra pequena, viés claro.
- **Sem evidência:** hipótese.

Nenhuma decisão estrutural se baseia em evidência fraca sem registro de risco.

---

## 7. Cronograma de pesquisa

| Fase | Atividade | Saída |
| --- | --- | --- |
| Pré-produto | entrevistas e auditoria | `docs/research/` |
| Protótipo | teste de usabilidade | `docs/experiments/` |
| Piloto | funil, satisfação, fairness | relatório |
| Pós-piloto | decisão de continuidade | `docs/decisions/` |

---

## 8. Registro

- Nota por estudo em `docs/research/`.
- Protocolo em `docs/experiments/`.
- Resultado com limitações.
- Decisão e impacto no PRD.
- Nenhum dado pessoal em nota pública.

## 9. Template

- [`_template.md`](_template.md)
