# ADR 0003 — Outbox transacional e eventos internos

**Status:** aceito  
**Data:** 2026-11-20  
**Decisores:** arquitetura, backend  
**Relacionado:** `docs/architecture/architecture-plan.md`

## Contexto

O RAIRE precisa desacoplar efeitos colaterais (notificação, indexação de busca, métricas, moderação) sem introduzir um broker externo no primeiro release. Perder um evento ou publicar antes do commit quebraria consistência.

## Opções consideradas

1. **Outbox transacional + worker interno** — evento gravado na mesma transação do agregado; worker lê e processa.
2. **Broker externo desde o início** — mais escalável, mais complexo e caro.
3. **Chamadas síncronas diretas** — simples, mas acopla contextos e arrisca inconsistência.

## Decisão

Adotar outbox transacional em tabela `outbox_events`, com worker interno lendo eventos pendentes, processamento idempotente por `event_id` e retry com backoff. Sem broker externo no MVP.

## Consequências

**Positivas:**

- consistência entre agregado e evento;
- baixo custo operacional;
- base para extrair broker depois;
- eventos auditáveis.

**Negativas:**

- latência de processamento;
- worker precisa de monitoramento;
- tabela pode crescer.

**Mitigações:**

- índice em `processed_at`;
- expurgo/arquivamento;
- alerta de backlog;
- idempotência testada.

## Critérios de reversão

- backlog > capacidade;
- necessidade de múltiplos consumidores externos;
- latência de evento crítica.

## Status

Aceito em 2026-11-20. Implementação condicionada à spec aprovada.
