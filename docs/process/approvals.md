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
| PRD v1.0 | 1.0 | em revisão | produto, arquitetura, segurança | nenhum código autorizado |
| foundation-platform | 1.0 | em revisão | Victor, Felipe | depende de Q-07, Q-44–Q-50 |
| web-experience | 1.0 | em revisão | Anderson, Victor | depende de foundation |
| local-discovery | 1.0 | em revisão | Victor, Anderson | depende de foundation |
| service-request | 1.0 | em revisão | Anderson, Felipe | depende de foundation |
| trust-reputation | 1.0 | em revisão | Felipe, Victor | depende de service-request |
| fair-visibility | 1.0 | em revisão | Victor, Felipe, Anderson | depende de local-discovery |
| ADR 0001 | 1.0 | proposto | Victor | aguarda decisão |
| ADR 0002 | 1.0 | proposto | Victor, Felipe | SQLite local/Postgres alvo |
| ADR 0003 | 1.0 | proposto | Victor | outbox |
| ADR 0004 | 1.0 | proposto | Victor, Felipe | sessão/RBAC/MFA |
| ADR 0005 | 1.0 | proposto | Victor, Anderson, Felipe | ranking/fairness |
| ADR 0006 | 1.0 | proposto | Anderson, Victor | web server-rendered |
| Política de fairness v1 | 1.0 | proposta | produto + pesquisa + segurança | defaults do PRD |
| DPIA | — | não iniciada | jurídico/privacidade | antes do piloto |

---

## 3. Registro de aprovações

| Data | Item | Aprovador | Decisão | Ressalvas | Prazo |
| --- | --- | --- | --- | --- | --- |
| — | — | — | — | — | — |

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
