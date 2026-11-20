# Follow-ups da Fundação — 2026-11-20

**Status:** aberto  
**Origem:** implementação `FND-01..FND-06`, `FND-14..FND-16` e fatias de `FND-09/10/17`  
**Regra:** follow-up não é ticket aprovado; precisa entrar no board/CSV e passar pelo fluxo normal.

---

## 1. Bloqueantes para produção/piloto

| ID | Follow-up | Prioridade | Dono sugerido | Critério de aceite |
| --- | --- | --- | --- | --- |
| FUP-01 | Adaptador PostgreSQL/PostGIS atrás das interfaces de repositório | P0 | Victor | migrations e contratos rodam em Postgres; testes de contrato SQLite/Postgres no CI; staging com PostGIS |
| FUP-02 | Worker de outbox com consumo idempotente e retry | P0 | Victor/Felipe | eventos pendentes são processados; falha não perde evento; métrica de backlog e alerta |
| FUP-03 | CSRF de login e verificação de `Origin`/`Referer` em mutações | P0 | Victor/Felipe | login cross-site bloqueado; teste de origem ausente/estranha; SameSite documentado |
| FUP-04 | Rate limit compartilhado (Postgres/Redis) e limites por operador | P1 | Felipe | duas instâncias compartilham contador; teste de concorrência; `Retry-After` correto |
| FUP-05 | Enrollment, rotação e regeneração de recovery codes de MFA | P1 | Felipe | operador consegue ativar/trocar MFA com auditoria; recovery codes antigos invalidados |
| FUP-06 | Wiring de feature flags e exportação de métricas | P1 | Felipe | flags gateiam features sensíveis; métricas exportadas; default `off` testado |
| FUP-07 | Limpeza de `idempotency_keys` e `email_tokens` expirados | P2 | Felipe | job diário remove expirados; auditoria/métrica; teste |
| FUP-08 | Política de PII para `evidence` administrativo e mapeamento de erro 5xx | P2 | Felipe | evidência não aceita PII desnecessária; erro de banco não vira 400 |
| FUP-09 | Fluxo de verificação de telefone (OTP/token) e emissão do selo | P1 | Felipe | telefone só fica verificado após confirmação; teste de preferência efetiva |
| FUP-10 | Backfill de recovery codes para contas anteriores à migration 0004 | P2 | Felipe | operador antigo recebe códigos novos ou é forçado a reenrollment |

---

## 2. Entregas de outras specs

| ID | Follow-up | Dono | Observação |
| --- | --- | --- | --- |
| FND-07 | Portfólio, imagem validada, reordenação, denúncia | Anderson | tabela `portfolio_items` já existe; falta CRUD/HTTP/testes |
| FND-08 | Exportação, exclusão/anonimização, solicitações de direitos | Felipe | `account_requests` já existe; falta fluxo e redaction de export |
| FND-11/FND-12 | Web de conta, onboarding e perfil | Anderson | nenhum template/handler HTML implementado |
| FND-13/FND-18 | E2E, axe, 360 px, termos, estados, copy | Felipe/Anderson | depende da web |
| FND-19 | QA independente de hardening | Felipe | suíte de Victor executada; falta revisão independente, race, E2E e evidência final |

---

## 3. Riscos e decisões em aberto

| Item | Risco | Mitigação atual | Próximo passo |
| --- | --- | --- | --- |
| SQLite em produção | diferença de dialeto, concorrência e geo | documentado; repositórios isolados | FUP-01 |
| Outbox sem worker | eventos acumulam sem consumidor | payload mínimo e tabela indexada | FUP-02 |
| Login CSRF | login cross-site pode fixar sessão do atacante | SameSite=Lax + cookies HttpOnly | FUP-03 |
| Rate limit em memória | bypass ao escalar horizontalmente | limites por instância | FUP-04 |
| MFA sem enrollment | operador depende de seed/CLI | `seed-admin` em dev | FUP-05 |
| Telefone sem fluxo de verificação | selo de telefone não é emitido | campo e mascaramento existem | FUP-09 |
| Flags não gateiam features | flag não protege rollout | default `off` e tabela | FUP-06 |
| Evidência admin | pode conter PII | acesso admin + auditoria | FUP-08 |
| Sem HSTS | downgrade em produção | `Secure` configurável | configurar TLS/HSTS no deploy |
| Sem Postgres/DPIA/piloto | não pode ir a produção | nenhum deploy autorizado | jurídico + FUP-01 |

---

## 4. Dívida técnica aceitável no MVP

- Rate limit em memória e métricas em memória, desde que o deploy seja uma instância.
- Sem tracing distribuído.
- Sem broker externo.
- Sem busca/ranking nesta spec.
- Sem cache de sessão; cada request valida no banco.
- Sem jobs assíncronos além do expurgo diário.

---

## 5. Critério para fechar a fundação

- [ ] FUP-01 e FUP-02 resolvidos ou aceitos formalmente como risco de produção.
- [ ] review independente de Anderson e Felipe registrado.
- [ ] `go test -race ./...` verde no CI Linux.
- [ ] E2E e acessibilidade da web anexados.
- [ ] evidência final em `specs/foundation-platform/evidence/`.
- [ ] aprovação humana de produto, segurança e QA.
- [ ] DPIA/jurídico/DPO concluídos antes de qualquer piloto.
