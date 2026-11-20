# Waiver — Planejamento do MVP (sem código)

**ID:** W-2026-11-20-01  
**Status:** encerrado para a fundação em 2026-11-20  
**Data:** 2026-11-20  
**Solicitante:** fundador/produto  
**Aprovador:** instrução humana registrada no pedido de trabalho  
**Escopo:** organização, PRD, docs, specs, tickets, planos, pipeline e divisão de tarefas  
**Encerramento:** `foundation-platform` aprovada; demais specs seguem o fluxo normal de revisão

---

## 1. Motivo

O repositório estava reduzido ao PRD v0.1. Era necessário reconstruir o harness, revisar o PRD e preparar o trabalho antes de qualquer implementação. O pedido humano autorizou explicitamente **não implementar código ainda**.

## 2. O que este waiver autoriza

- criar e editar documentação;
- revisar o PRD;
- criar specs, critérios de aceite, tickets e planos;
- criar ADRs, processos, templates e task board;
- definir divisão de trabalho entre Victor, Anderson e Felipe;
- preparar evidências e checklists de gate (sem executá-los).

## 3. O que este waiver NÃO autoriza

- implementar código de produto;
- criar migrations, endpoints ou testes de código;
- executar testes, scans ou deploy;
- marcar spec/ADR como `aprovada`/`aceito`;
- iniciar piloto ou produção;
- tratar hipótese como fato.

## 4. Critérios de saída

Este waiver termina quando:

1. as specs estiverem revisadas e aprovadas por humano; ou
2. o fundador revogar o waiver.

**Status:** encerrado para a `foundation-platform` em 2026-11-20. As demais specs seguem o fluxo normal: revisão → aprovação → ticket → implementação.

Após o término, vale o fluxo normal: spec aprovada → ticket → implementação → review → QA → segurança → release.

## 5. Riscos aceitos

| Risco | Mitigação |
| --- | --- |
| documentação divergir do código futuro | toda mudança de escopo atualiza PRD/spec antes do código |
| excesso de documentação atrasar | foco em specs e tickets acionáveis; cortar o que não ajuda |
| default virar decisão sem revisão | defaults marcados como propostas e com dono/prazo |

## 6. Registro

- Pedido humano: "não implemente código ainda, apenas a organização de tudo e etc... principalmente tasks para quem vai trabalhar no repo".
- Data: 2026-11-20.
- Aprovação para implementação: **pendente**.
