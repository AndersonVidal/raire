# Time — Índice e Handoffs

**Status:** ativo  
**Versão:** 1.0  
**Equipe:** Victor, Anderson, Felipe

---

## 1. Comece por aqui

1. Leia `AGENTS.md` e `CLAUDE.md`.
2. Leia o PRD revisado (`docs/product/prd.md`).
3. Leia o pipeline (`docs/process/pipeline.md`) e seu handoff.
4. Leia a spec da sua frente e o task board.
5. Verifique as questões abertas que bloqueiam seu trabalho.
6. Não implemente nada antes da aprovação humana da spec.

---

## 2. Handoffs

| Pessoa | Handoff | Foco |
| --- | --- | --- |
| Victor | [`handoff-victor.md`](handoff-victor.md) | arquitetura, foundation, discovery, fairness |
| Anderson | [`handoff-anderson.md`](handoff-anderson.md) | produto, web, acessibilidade, service-request |
| Felipe | [`handoff-felipe.md`](handoff-felipe.md) | QA, segurança, privacidade, trust-reputation, DevOps |

---

## 3. Documentos de trabalho

- [`raci.md`](raci.md) — responsabilidades e decisões.
- [`task-board.md`](task-board.md) — todas as tarefas, donos, dependências e status.
- [`sprint-plan.md`](sprint-plan.md) — plano de sprints e ordem de execução.
- [`tickets.csv`](tickets.csv) — lista importável para ferramenta de gestão.
- [`../process/pipeline.md`](../process/pipeline.md) — fluxo, gates e SLAs.
- [`../process/approvals.md`](../process/approvals.md) — o que está aprovado e o que falta.
- [`../process/waivers/2026-11-20-mvp-planning.md`](../process/waivers/2026-11-20-mvp-planning.md) — autorização de planejamento sem código.

---

## 4. Primeira semana (todos)

### Objetivo

Sair de "documentação pronta" para "specs revisadas e plano executável", sem código.

| Dia | Atividade | Saída |
| --- | --- | --- |
| 1 | Ler PRD, pipeline, RACI e specs da sua frente | dúvidas registradas |
| 2 | Revisar spec da sua frente e critérios de aceite | comentários no `spec.md` |
| 3 | Resolver questões abertas bloqueantes da sua frente | Q-xx atualizada |
| 4 | Revisar tickets, dependências e tamanho | task board atualizado |
| 5 | Reunião de review cruzado | parecer e pendências para aprovação |

### Critério de saída da semana

- [ ] specs da frente revisadas;
- [ ] questões bloqueantes com default ou decisão;
- [ ] tickets refinados e atribuídos;
- [ ] riscos e evidências esperadas claros;
- [ ] nenhum código implementado;
- [ ] pedido de aprovação humana preparado.

---

## 5. Reuniões

| Reunião | Quando | Participantes | Saída |
| --- | --- | --- | --- |
| Daily assíncrono | diário | todos | bloqueios e próximos passos |
| Refinamento | 2×/semana | dono da spec + revisores | tickets prontos |
| Review sync | por PR | autor + revisor + gate | parecer |
| Demo | semanal | todos | evidência e feedback |
| Riscos/fairness | mensal | todos | riscos atualizados |
| Retro | 2 semanas | todos | melhoria do harness |

---

## 6. Regras de convivência

- Critique o trabalho, não a pessoa.
- Bloqueio se comunica cedo.
- Decisão relevante se registra.
- Ninguém aprova o próprio código.
- Ninguém esconde teste que falha.
- Segurança, privacidade e fairness não são "departmento de ninguém": são de todos.
