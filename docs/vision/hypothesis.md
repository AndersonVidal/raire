# Hipóteses — RAIRE

**Status:** hipóteses em investigação  
**Versão:** 1.0  
**Regra:** hipótese não é fato; toda hipótese tem evidência esperada e critério de refutação.

---

## 1. Hipóteses de descoberta e visibilidade

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H1 | Profissionais com pouca presença digital são difíceis de descobrir | entrevistas + auditoria de canais | profissionais dizem que não é problema |
| H2 | Ranking por popularidade concentra oportunidades | análise de plataformas + simulação | concentração não se confirma |
| H3 | Explicação do ranking aumenta confiança | teste A/B + qualitativo | sem diferença ou confiança cai |

## 2. Hipóteses de confiança e reputação

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H4 | Clientes precisam de sinais além da nota | entrevistas + teste de perfil | nota isolada basta |
| H5 | Reputação contextual melhora decisão | teste com distribuição/tags | não muda decisão |
| H6 | Verificação honesta aumenta confiança | teste de selo | selo não é compreendido |

## 3. Hipóteses de contexto geográfico

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H7 | Distância é decisiva em serviços presenciais | funil por faixa | conversão não varia |
| H8 | Área de atendimento mal declarada gera frustração | entrevistas + dados | não há problema |
| H9 | Serviços remotos têm lógica diferente | comparação de funil | comportamento é igual |

## 4. Hipóteses de comportamento

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H10 | Mural de necessidade resolve cold start | piloto com pedidos | pedidos não recebem interesse |
| H11 | Profissionais querem busca ativa | uso do feed | pouco uso |
| H12 | Contato mascarado reduz medo | teste de usabilidade | usuários abandonam |

## 5. Hipóteses de matching e relevância

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H13 | Categoria + distância + disponibilidade bastam no MVP | golden set + conversão | relevância baixa |
| H14 | Fairness pode coexistir com relevância | simulação + experimento | queda inaceitável |
| H15 | Diversidade de bairro/categoria melhora descoberta | experimento | não melhora |

## 6. Hipóteses de IA e recomendação

| ID | Hipótese | Evidência esperada | Critério de refutação |
| --- | --- | --- | --- |
| H16 | IA generativa é necessária no MVP | comparação com baseline | baseline é suficiente |
| H17 | Personalização profunda melhora conexão | experimento | não melhora ou prejudica fairness |

---

## 7. Prioridade de investigação

1. H1, H2, H4, H7, H10.
2. H3, H5, H13, H14.
3. H6, H8, H9, H11, H12.
4. H15, H16, H17.

## 8. Como registrar evidência

- `docs/research/` para notas e referências;
- `docs/experiments/` para protocolo e resultado;
- `specs/<feature>/evidence/` para evidência de entrega;
- decisão em `docs/decisions/` quando mudar rumo.

## 9. Regra de linguagem

- "Acreditamos que..." → hipótese.
- "Observamos que..." → evidência.
- "Decidimos..." → decisão registrada.
- "Sabemos..." → somente com evidência forte e revisão.
