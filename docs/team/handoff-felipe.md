# Handoff — Felipe

**Papel:** QA + segurança + privacidade + DevOps  
**Ownership:** `trust-reputation`; gates de segurança, privacidade, QA e release; CI/CD e observabilidade  
**Revisa:** segurança, LGPD, PII, abuso, moderação, reputação, evidências  
**Não aprova:** o próprio código; exceção de segurança sozinho

---

## 1. Leitura obrigatória

1. `AGENTS.md` e `CLAUDE.md`.
2. `docs/product/prd.md` — visão, escopo, RFs, RNFs e decisões.
3. `docs/product/metrics-and-experiments.md` — métricas e guardrails.
4. `docs/product/risks.md` — riscos e donos.
5. `docs/security/threat-model.md` — ameaças e controles.
6. `docs/security/privacy-and-lgpd.md` — bases legais, retenção e direitos.
7. `docs/process/security-privacy.md` — checklist e bloqueios.
8. `docs/process/qa-strategy.md` — pirâmide, matriz e evidências.
9. `docs/process/pipeline.md` — fluxo, gates e SLAs.
10. `specs/trust-reputation/` e `specs/foundation-platform/` (admin/privacidade).

---

## 2. Foco da primeira semana

### Dia 1 — Contexto

- [ ] Ler PRD, riscos, segurança e specs da sua frente.
- [ ] Listar dúvidas de segurança, privacidade, moderação e QA.
- [ ] Verificar Q-06, Q-09, Q-10, Q-30–Q-36, Q-38–Q-43.

### Dia 2 — Modelo de ameaças

- [ ] Validar ameaças e controles.
- [ ] Priorizar riscos críticos: contato, localização, account takeover, admin.
- [ ] Definir o que bloqueia release.
- [ ] Revisar rate limits e authz.

### Dia 3 — Privacidade e LGPD

- [ ] Validar classificação de dados, bases legais e retenção.
- [ ] Definir consentimento, exportação, exclusão e incidente.
- [ ] Definir o que entra na DPIA.
- [ ] Revisar PII em logs/eventos.

### Dia 4 — QA e evidências

- [ ] Validar matriz de aceite das specs.
- [ ] Definir testes críticos, acessibilidade e performance.
- [ ] Definir formato de evidência.
- [ ] Revisar dados sintéticos e fixtures.

### Dia 5 — Review cruzado e handoff

- [ ] Revisar specs de Victor e Anderson sob ótica de segurança/QA.
- [ ] Atualizar task board.
- [ ] Preparar parecer para aprovação humana.
- [ ] Registrar riscos e bloqueios.

---

## 3. Entregáveis da fase de planejamento

- [ ] Parecer de segurança/privacidade das 6 specs.
- [ ] Modelo de ameaças revisado.
- [ ] Checklist de LGPD e retenção.
- [ ] Plano de QA e matriz de aceite consolidada.
- [ ] Pipeline de CI e critérios de evidência.
- [ ] Runbook de incidente.
- [ ] Riscos e bloqueios de release.
- [ ] Plano de moderação.

---

## 4. Primeiros tickets após aprovação

| Ordem | Ticket | Entrega |
| --- | --- | --- |
| 1 | FND-08 | consentimento, PII, exportação/exclusão |
| 2 | FND-09 | admin mínimo, suspensão e auditoria |
| 3 | FND-10 | observabilidade, health e feature flags |
| 4 | TRS-01 | modelo de denúncia e migrations |
| 5 | TRS-02 | desfecho de conexão |
| 6 | TRS-03..TRS-05 | elegibilidade, avaliação e reputação |

Detalhes no [`task-board.md`](task-board.md) e em `specs/*/tickets.md`.

---

## 5. Responsabilidades de review

- Gate de segurança/privacidade em todo PR.
- Gate de trust & safety em moderação/reputação/fairness.
- Gate de QA em toda entrega.
- Bloquear merge se houver:
  - PII em log/evento;
  - contato/localização expostos;
  - authz quebrada;
  - rate limit ausente em auth;
  - moderação sem auditoria;
  - teste escondido ou critério sem cobertura;
  - dependência crítica vulnerável.

---

## 6. Perguntas que você precisa responder

- Q-06: verificação de identidade obrigatória?
- Q-09: confirmação de contato por e-mail/telefone?
- Q-10: retenção de documentos.
- Q-30: como confirmar serviço?
- Q-31: moderação prévia ou pós-publicação?
- Q-32..Q-36: avaliação, edição, verificação, denúncia anônima, SLA.
- Q-38..Q-43: bases legais, retenção, DPO, transferência, incidente, pesquisa.

Registre a decisão em `docs/product/open-questions.md` e, se jurídico, em `docs/security/privacy-and-lgpd.md`.

---

## 7. Critérios de sucesso

- nenhum incidente crítico de privacidade;
- authz e rate limit testados;
- PII fora de log/evento;
- moderação auditável e recorrível;
- avaliação só após conexão elegível;
- crítica legítima preservada;
- evidência completa e reproduzível;
- nenhum código antes da aprovação humana.
