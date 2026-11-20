# Aprovações e Gates — RAIRE

**Status:** operacional  
**Versão:** 1.0  
**Uso:** registrar quem aprovou o quê, quando e com quais ressalvas.

> Nenhum agente de IA pode aprovar spec, PRD, ADR, política de fairness ou liberação de piloto. Aprovação é ato humano registrado.

---

## 1. Matriz de aprovação

| Artefato | Aprovadores obrigatórios | Evidência |
| --- | --- | --- |
| PRD | produto + arquitetura + segurança/privacidade | registro neste arquivo |
| Spec de feature | produto + arquitetura (+ segurança se tocar PII) | status `aprovada` no `spec.md` |
| ADR | arquitetura (+ segurança quando aplicável) | status `aceito` no ADR |
| Política de fairness | produto + pesquisa + segurança | `docs/decisions` + ata |
| DPIA / retenção | jurídico/privacidade + segurança | `docs/security` |
| Release para staging | arquitetura + QA | checklist de release |
| Piloto/produção | fundador + segurança/privacidade + produto | checklist + registro |
| Exceção de gate | fundador + dono do gate | waiver com prazo e mitigação |

---

## 2. Estado das aprovações

| Item | Versão | Status | Aprovadores pendentes | Observação |
| --- | --- | --- | --- | --- |
| PRD v1.0 | 1.0 | aprovado | — | aprovado em 2026-11-20; nenhum código autorizado por si só |
| Decisões Q-01..Q-50 | 1.1 | aprovado | — | registro em `docs/product/decisions-2026-11-20.md`; jurídico pendente antes do piloto |
| foundation-platform | 1.1 | aprovada | — | aprovada em 2026-11-20; tickets FND prontos |
| web-experience | 1.0 | em revisão | Anderson, Victor | depende de foundation |
| local-discovery | 1.0 | em revisão | Victor, Anderson | depende de foundation |
| service-request | 1.0 | em revisão | Anderson, Felipe | depende de foundation |
| trust-reputation | 1.0 | em revisão | Felipe, Victor | depende de service-request |
| fair-visibility | 1.0 | em revisão | Victor, Felipe, Anderson | depende de local-discovery |
| ADR 0001 | 1.0 | aceito | — | aprovado em 2026-11-20 |
| ADR 0002 | 1.0 | aceito | — | SQLite local/Postgres alvo; aprovado em 2026-11-20 |
| ADR 0003 | 1.0 | aceito | — | outbox; aprovado em 2026-11-20 |
| ADR 0004 | 1.0 | aceito | — | sessão/RBAC/MFA; aprovado em 2026-11-20 |
| ADR 0005 | 1.0 | aceito | — | ranking/fairness; ativação depende de simulação |
| ADR 0006 | 1.0 | aceito | — | web server-rendered; aprovado em 2026-11-20 |
| Política de fairness v1 | 1.0 | aprovada (parâmetros) | — | ativação condicionada a simulação offline |
| DPIA | — | pendente | jurídico/privacidade | antes do piloto |

---

## 3. Registro de aprovações

| Data | Item | Aprovador | Decisão | Ressalvas | Prazo |
| --- | --- | --- | --- | --- | --- |
| 2026-11-20 | PRD v1.0 | Humano responsável (instrução registrada) | aprovado | implementação continua condicionada à spec aprovada | — |
| 2026-11-20 | Q-01..Q-50 | Humano responsável (instrução registrada) | aprovado | Q-38–Q-43 exigem validação jurídica/DPIA antes do piloto; Q-01/Q-03/Q-40 pendências operacionais antes de F6 | antes do piloto |
| 2026-11-20 | foundation-platform v1.1 | Humano responsável (instrução registrada) | aprovada | must-fix e should-fix integrados; AC-40..AC-60 adicionados | — |
| 2026-11-20 | ADRs 0001–0006 | Humano responsável (instrução registrada) | aceito | ADR 0005 tem ativação condicionada a simulação | — |
| 2026-11-20 | Política de fairness v1 (parâmetros) | Humano responsável (instrução registrada) | aprovado | ativação só após simulação offline e revisão | antes de F5 |
| — | DPIA | — | pendente | jurídico/privacidade | antes do piloto |

**Decisões possíveis:** `aprovado`, `aprovado com ressalvas`, `rejeitado`, `adiado`.

---

## 4. Aprovação de spec

Uma spec só muda para `aprovada` quando:

- [ ] requisitos rastreáveis ao PRD;
- [ ] critérios de aceite completos;
- [ ] fora de escopo explícito;
- [ ] riscos de segurança/privacidade/fairness avaliados;
- [ ] perguntas abertas com dono e prazo;
- [ ] dependências e tickets definidos;
- [ ] aprovação humana registrada.

---

## 5. Aprovação de release

Uma release só avança quando:

- [ ] pipeline verde;
- [ ] staging com smoke/E2E;
- [ ] migração e rollback;
- [ ] flags com default seguro;
- [ ] observabilidade e alertas;
- [ ] runbook;
- [ ] evidência;
- [ ] aprovação humana.

---

## 6. Waivers

Um waiver é uma exceção registrada, com:

- escopo exato;
- motivo;
- risco aceito;
- mitigação;
- prazo;
- aprovador humano;
- condição de encerramento.

Waivers ativos ficam em `docs/process/waivers/`. Nenhum waiver autoriza violação de segurança, PII, fair ranking ou LGPD.
