# fair-visibility — mapa de aceite

**Status:** em revisão (testes propostos)  
**Spec:** `spec.md`  
**Dono:** Victor  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `Fair` = fairness, `Perf` = performance, `Sec` = segurança.
- Dados de exposição são sintéticos.
- Coortes são agregadas.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | política versionada | F-01, F-02 | U | `fairness` › "nova versão é imutável" | pendente |
| AC-02 | status da política | F-03 | U | `fairness` › "transições draft→simulated→approved→active" | pendente |
| AC-03 | aprovação registrada | F-04 | I | `fairness` › "ativação sem aprovação falha" | pendente |
| AC-04 | determinismo | F-05, F-18 | U/Fair | `fairness` › "mesma entrada/sessão/política = mesma ordem" | pendente |
| AC-05 | flag desliga | F-06 | I | `fairness` › "flag off usa baseline" | pendente |
| AC-06 | config inválida | F-01 | U | `fairness` › "caps negativos/piso inválido são rejeitados" | pendente |
| AC-07 | ledger de impressão | F-07 | I | `fairness` › "impressão grava perfil/posição/política" | pendente |
| AC-08 | ledger sem PII | F-08 | Sec | `fairness` › "ledger não contém e-mail/telefone" | pendente |
| AC-09 | caps por sessão/dia | F-09 | U | `fairness` › "cap limita repetição" | pendente |
| AC-10 | replay | F-10 | I | `fairness` › "replay reconstrói decisão" | pendente |
| AC-11 | retenção | F-11 | I | `fairness` › "job agrega e expurga ledger" | pendente |
| AC-12 | cap aplicado | F-12 | U | `fairness` › "perfil não passa do cap" | pendente |
| AC-13 | exploração | F-13 | U | `fairness` › "20% dos slots são exploração elegível" | pendente |
| AC-14 | boost de novo | F-14 | U | `fairness` › "novo recebe boost com decaimento" | pendente |
| AC-15 | diversidade | F-15 | U | `fairness` › "primeira página tem 2 bairros/categorias" | pendente |
| AC-16 | piso de relevância | F-16 | U | `fairness` › "candidato abaixo do piso não entra" | pendente |
| AC-17 | poucos elegíveis | F-17 | U | `fairness` › "sem candidatos, preenche baseline e registra" | pendente |
| AC-18 | fallback | F-06, F-18 | I | `fairness` › "falha usa baseline + alerta" | pendente |
| AC-19 | experimento | F-19 | U | `fairness` › "experimento exige hipótese/guardrail" | pendente |
| AC-20 | assignment estável | F-20 | U | `fairness` › "mesmo usuário mantém variante" | pendente |
| AC-21 | parada por guardrail | F-21 | I | `fairness` › "guardrail violado para experimento" | pendente |
| AC-22 | sem atributo sensível | F-22 | Sec/Fair | `fairness` › "schema não tem atributo sensível" | pendente |
| AC-23 | registro de experimento | F-23 | I | `fairness` › "resultado e decisão registrados" | pendente |
| AC-24 | Gini | F-24 | U | `fairness` › "cálculo de Gini correto" | pendente |
| AC-25 | share top 10/20 | F-24 | U | `fairness` › "share por sessão/dia" | pendente |
| AC-26 | share de novos | F-24 | U | `fairness` › "novos no top 20" | pendente |
| AC-27 | dashboard | F-25, F-26 | I | `fairness` › "dashboard agrega sem PII" | pendente |
| AC-28 | alerta de guardrail | F-27 | I | `fairness` › "alerta quando guardrail viola" | pendente |
| AC-29 | admin vê políticas | F-29 | I | `fairness` › "admin lista políticas e simulações" | pendente |
| AC-30 | auditoria | F-30 | I | `fairness` › "ativação grava audit_event" | pendente |
| AC-31 | transparência | F-31, F-33 | C | `fairness` › "página pública reflete política ativa" | pendente |
| AC-32 | contestação | F-32 | I | `fairness` › "profissional pode contestar posição" | pendente |
| AC-33 | explicação coerente | F-33 | U | `fairness` › "motivos vêm dos sinais usados" | pendente |
| AC-34 | simulador | F-34 | U | `fairness` › "simula consultas históricas" | pendente |
| AC-35 | métricas offline | F-35 | U | `fairness` › "calcula nDCG/MRR/Gini/cobertura" | pendente |
| AC-36 | determinismo offline | F-36 | U | `fairness` › "mesma simulação = mesmo resultado" | pendente |
| AC-37 | ativação exige simulação | F-37 | I | `fairness` › "ativar sem relatório falha" | pendente |

## Cenários Gherkin

```gherkin
Cenário: política só ativa após simulação e aprovação
  Dado uma política em rascunho
  Quando tento ativar
  Então recebo POLICY_NOT_SIMULATED
  Quando simulo e aprovo
  Então posso ativar

Cenário: exploração não promove abaixo do piso
  Dado um candidato com score abaixo do percentil 40
  Quando a camada de fairness escolhe slots
  Então esse candidato não é promovido

Cenário: guardrail violado interrompe experimento
  Dado um experimento em execução
  Quando o CTR cai além do guardrail
  Então o experimento é interrompido
  E um alerta é registrado

Cenário: fallback
  Dado que a camada de fairness está indisponível
  Quando executo a busca
  Então o baseline é usado
  E a resposta indica degradação
```

## Evidências

- Relatório de simulação: pendente
- Métricas: pendente
- Segurança/fairness: pendente
- Acessibilidade do dashboard: pendente
