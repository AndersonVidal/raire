# local-discovery — mapa de aceite

**Status:** em revisão (testes propostos)  
**Spec:** `spec.md`  
**Dono:** Victor  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `A11y` = acessibilidade, `Perf` = performance, `Fair` = fairness, `Sec` = segurança.
- Dados de busca são sintéticos.
- Golden set é versionado.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | busca por termo e categoria | L-01 | U/I | `discovery` › "busca por termo e categoria retorna elegíveis" | pendente |
| AC-02 | busca por cidade/bairro | L-01, L-04 | I | `discovery` › "localização é normalizada para área aproximada" | pendente |
| AC-03 | raio em presencial | L-02 | U | `discovery` › "raio limita resultados presenciais" | pendente |
| AC-04 | remoto ignora raio | L-02 | U | `discovery` › "remoto não usa distância" | pendente |
| AC-05 | filtros válidos | L-03 | I | `discovery` › "filtros de tipo, disponibilidade e verificação" | pendente |
| AC-06 | filtro inválido | L-03 | U | `discovery` › "filtro inválido retorna VALIDATION_ERROR" | pendente |
| AC-07 | busca vazia sugere categorias | L-05 | U | `discovery` › "sem termo/categoria retorna sugestões" | pendente |
| AC-08 | termo inválido | L-06 | U | `discovery` › "termo vazio/controle é rejeitado" | pendente |
| AC-09 | retrieval só published | L-07 | I | `discovery` › "draft/paused/suspended não aparecem" | pendente |
| AC-10 | área contém ponto | L-07 | U | `discovery` › "perfil fora do raio não aparece" | pendente |
| AC-11 | limite de candidatos | L-08 | I | `discovery` › "retrieval limita candidatos" | pendente |
| AC-12 | score baseline | L-09 | U | `discovery` › "score combina sinais com pesos versionados" | pendente |
| AC-13 | pesos versionados | L-10 | U | `discovery` › "mudança de peso exige nova versão" | pendente |
| AC-14 | determinismo | L-11, L-15 | U | `discovery` › "mesma entrada/sessão/política = mesma ordem" | pendente |
| AC-15 | reputação bayesiana | L-12 | U | `discovery` › "novo profissional recebe prior, não zero" | pendente |
| AC-16 | guardrails de diversidade | L-13 | U/Fair | `discovery` › "primeira página tem diversidade mínima" | pendente |
| AC-17 | fallback | L-14 | I | `discovery` › "falha de fairness usa baseline e marca degraded" | pendente |
| AC-18 | explicação | L-16, L-17 | U | `discovery` › "até 3 motivos derivados dos sinais" | pendente |
| AC-19 | copy proibida | L-18 | U | `discovery` › "motivos não usam 'melhor' ou 'garantido'" | pendente |
| AC-20 | transparência | L-19, L-20 | C | `discovery` › "página pública explica critérios sem pesos" | pendente |
| AC-21 | resultado completo | L-21 | C | `discovery` › "card tem nome, categoria, distância, sinais e explicação" | pendente |
| AC-22 | perfil público | L-22 | C | `discovery` › "perfil mostra bio, área aproximada, portfólio" | pendente |
| AC-23 | PII | L-23 | C/Sec | `discovery` › "resposta pública não contém PII" | pendente |
| AC-24 | paginação | L-24 | I | `discovery` › "cursor avança sem duplicar" | pendente |
| AC-25 | estado vazio | L-25 | E2E | `discovery` › "vazio sugere ajuste" | pendente |
| AC-26 | eventos de funil | L-26, L-27 | I | `discovery` › "eventos usam IDs pseudonimizados e policy_version" | pendente |
| AC-27 | consentimento | L-28 | I | `discovery` › "sem consentimento, evento é descartado/agregado" | pendente |
| AC-28 | logs sem termo | L-29 | Sec | `discovery` › "log de busca não associa termo a usuário" | pendente |
| AC-29 | dashboard | L-30 | I | `discovery` › "funil busca→clique→intenção" | pendente |
| AC-30 | performance | L-31, L-32 | Perf | `discovery` › "busca p95 < 500 ms no dataset" | pendente |
| AC-31 | acessibilidade | L-33, L-34 | A11y | `web` › "resultados e filtros por teclado" | pendente |
| AC-32 | reduced motion | L-35 | A11y | `web` › "respeita prefers-reduced-motion" | pendente |

## Cenários Gherkin

```gherkin
Cenário: cliente busca eletricista no Centro
  Dado que estou na busca
  Quando informo "eletricista" e "Centro"
  Então vejo profissionais publicados que atendem na região
  E cada resultado mostra uma explicação curta
  E não vejo e-mail, telefone ou coordenada precisa

Cenário: busca remota não usa distância
  Dado que filtro por atendimento remoto
  Quando busco "consultoria"
  Então resultados remotos aparecem independentemente da distância
  E a explicação não cita proximidade

Cenário: ranking é determinístico
  Dado a mesma busca, sessão e política
  Quando executo duas vezes
  Então a ordem dos resultados é idêntica

Cenário: fairness falha e há fallback
  Dado que a camada de fairness está indisponível
  Quando executo a busca
  Então recebo resultados baseline
  E a resposta indica modo degradado
  E um alerta é registrado

Cenário: perfil suspenso sai da busca
  Dado um perfil publicado
  Quando o admin suspende o perfil
  Então a próxima busca não o retorna
```

## Evidências

- Relatório de testes: pendente
- Golden set: pendente
- Performance: pendente
- Acessibilidade: pendente
- Fairness: pendente
