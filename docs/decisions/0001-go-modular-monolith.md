# ADR 0001 — Backend em Go com monólito modular

**Status:** proposto  
**Data:** 2026-11-20  
**Decisores:** arquitetura, backend, produto  
**Relacionado:** `docs/architecture/architecture-plan.md`

## Contexto

O RAIRE está em fase de validação. O backend precisa de clareza, desempenho previsível, baixo custo operacional e separação forte de responsabilidades. A preferência registrada é Golang, arquitetura modular e possível orientação a eventos. Não há evidência que justifique microserviços.

## Opções consideradas

1. **Monólito modular em Go** — um deployable, contextos delimitados, contratos internos, eventos via outbox.
2. **Microserviços desde o início** — serviços separados por contexto, comunicação por rede.
3. **Monólito em camadas tradicional** — um pacote por camada (handlers, services, repositories).

## Decisão

Adotar monólito modular em Go, com contextos delimitados (`identity`, `catalog`, `directory`, `demand`, `discovery`, `fairness`, `reputation`, `moderation`, `platform`), regras de dependência explícitas e possibilidade de extração futura.

Microserviços só entram com gatilho medido: escala, times independentes, deploy isolado ou tecnologia específica.

## Consequências

**Positivas:**

- transações locais e consistência simples;
- menor custo e complexidade operacional;
- refatoração mais fácil durante a validação;
- testes de integração mais baratos;
- fronteiras claras para extrair depois.

**Negativas:**

- exige disciplina para não virar monólito emaranhado;
- deploy único acopla o ciclo de release;
- escala independente limitada.

**Mitigações:**

- regras de dependência no CI;
- revisão de arquitetura;
- contratos e eventos versionados;
- gatilhos documentados para extração.

## Critérios de reversão

- time > 8 pessoas ou necessidade de deploy independente;
- contexto com carga desproporcional;
- tecnologia específica que justifique serviço separado.

## Status

Proposto. Requer aprovação humana antes da implementação.
