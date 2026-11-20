# foundation-platform — security review

**Status:** revisão interna e independente executada; aprovação de Felipe pendente  
**Spec:** `spec.md`  
**Dono:** Victor  
**Revisor designado:** Felipe  
**Implementação:** commit local `2602cba`  
**Data:** 2026-11-20  
**Escopo:** código local da fundação; sem deploy, dados reais ou piloto

## Veredito

`aceitável para desenvolvimento local após correções; bloqueado para produção/piloto`.
A revisão independente encontrou enumeração no signup, lost update na suspensão e
consumo não atômico de recovery code; os três foram corrigidos e cobertos por
teste. Continuam riscos altos de produção (Postgres, outbox, login CSRF) e a
validação jurídica/DPIA.

## 3.1 Correções da revisão independente

| Achado | Correção |
| --- | --- |
| Signup revelava e-mail existente (201 vs 202) | resposta sempre `202 {"status":"ok"}`; token de dev só via header e só para conta criada |
| PATCH de perfil podia desfazer suspensão concorrente | `version` + update condicional; conflito retorna 409 |
| Suspensão e revogação de sessão não eram atômicas | mesma transação para status, auditoria, outbox e sessões |
| Recovery code podia ser consumido duas vezes em corrida | tabela dedicada com `UPDATE ... WHERE used_at IS NULL` no login |
| Token de e-mail não checava linhas afetadas | confirmação/reset retornam `TOKEN_INVALID` no perdedor da corrida |
| Telefone re-salvo perdia verificação | só invalida quando o número muda |
| Bio/disponibilidade podiam expor contato | validação de PII em texto público |
| API lia repositório direto | consultas passaram por métodos do `App` |
| `RAIRE_ENV` ausente caía em dev | default seguro de produção, chave obrigatória |

## 1. Superfície avaliada

- auth: cadastro, login, logout, confirmação, reset, sessão, MFA;
- authz: conta, perfil, admin, catálogo, flags, auditoria;
- PII: e-mail, telefone, localização precisa, logs, eventos;
- upload: fora da fatia (portfólio);
- admin: suspensão, categoria, flag, auditoria;
- eventos/logs: outbox, audit, access log, correlation id;
- dependências: Go stdlib, `x/crypto`, `modernc.org/sqlite`, `yaml.v3`.

## 2. Checklist

- [x] autenticação/sessão: token opaco de 32 bytes, hash SHA-256 no banco, expiração, rotação no login, revogação;
- [x] autorização por recurso: perfil do dono, preview por dono, admin por papel + MFA de sessão;
- [x] CSRF/CSP/headers: CSRF em mutação autenticada, CSP, nosniff, frame deny, no-store;
- [x] rate limit: login, signup, reset, confirmação/reenvio (em memória);
- [x] validação/encoding: JSON estrito, limites de tamanho, validação de domínio;
- [x] segredos: senha Argon2id, TOTP cifrado, recovery codes com hash, chave fora do banco;
- [x] PII em log/evento: access log sem query/corpo; eventos de domínio só com IDs; auditoria administrativa tem motivo/evidência livre, restrita a operador;
- [x] localização/contato: AES-256-GCM, retenção 90 dias, expurgo auditado, contato mascarado;
- [x] retenção/consentimento: consentimentos versionados e revogáveis;
- [x] dependências: `go vet`; sem CVE scan executado;
- [x] auditoria: append-only por trigger, ator/alvo/motivo/correlação;
- [ ] incidente: runbook e simulação antes do piloto (FUP/jurídico).

## 3. Controles implementados e testados

| Controle | Evidência |
| --- | --- |
| Argon2id com parâmetros no hash | `security_test.go` |
| TOTP + recovery code de uso único | `TestOperatorRequiresMFAAndRecovery` |
| Segredo TOTP cifrado no banco | teste consulta `mfa_secret` e exige `enc:v1:` |
| CSRF em mutação autenticada | `TestCSRFRequired` |
| Rate limit de login/signup/reset | `TestLoginRateLimit`, `TestSignupRateLimit`, `TestResetRateLimit` |
| Idempotência de signup/confirmação | `TestSignupIdempotency`, `TestConfirmIdempotency` |
| E-mail duplicado genérico | `TestDuplicateSignupIsGeneric` |
| Localização precisa cifrada/expurgada | `TestPreciseLocationEncryptedAndPurged` |
| API pública sem PII | `TestPublishedProfileHasNoPII` |
| Logs sem PII | `TestLogsHaveCorrelationAndNoPII` |
| Auditoria append-only | `TestAuditAppendOnly` |
| Categoria inativa oculta perfil | `TestCategoryDeactivatedHidesProfile` |
| Suspensão revoga sessão + audita | `TestAdminSuspendsReactivatesAndAudits` |

## 4. Achados

### Altos

| # | Achado | Impacto | Follow-up |
| --- | --- | --- | --- |
| S-01 | Sem adaptador Postgres/PostGIS | produção não suportada; diferenças de dialeto/geo | FUP-01 |
| S-02 | Outbox sem worker | eventos não consumidos; backlog silencioso | FUP-02 |
| S-03 | Login sem CSRF/origin check | login CSRF pode fixar sessão do atacante | FUP-03 |

### Médios

| # | Achado | Impacto | Follow-up |
| --- | --- | --- | --- |
| S-04 | Rate limit em memória | bypass ao escalar horizontalmente | FUP-04 |
| S-05 | Sem enrollment/rotação de MFA | operação depende de seed/CLI | FUP-05 |
| S-06 | Sem fluxo de verificação de telefone | selo de telefone não é emitido na prática | FUP-09 |
| S-07 | Flags não gateiam features | rollout não protegido por flag | FUP-06 |
| S-08 | Evidência de admin pode conter PII | exposição interna indevida | FUP-08 |
| S-09 | Sem HSTS/TLS documentado no deploy | downgrade em produção | configurar no staging |

### Baixos

| # | Achado | Impacto | Follow-up |
| --- | --- | --- | --- |
| S-10 | Idempotency/tokens expirados não são limpos | crescimento de tabela | FUP-07 |
| S-11 | TOTP aceita ±1 passo de 30 s | replay curto dentro da janela | padrão aceito; documentar |
| S-12 | `LastSeenAt` grava no banco a cada 5 min | carga desnecessária | monitorar/ajustar |

## 5. Bloqueios

- Nenhuma aprovação de segurança independente foi registrada.
- Nenhum DPIA, jurídico ou DPO concluído.
- Nenhum teste de dependências (`govulncheck`), pentest ou simulação de incidente.
- Nenhum deploy autorizado.

## 6. Riscos residuais

- PII de contato permanece mascarada, mas o dono vê o próprio e-mail no `/v1/me`; isso é esperado.
- O admin pode inserir evidência livre; política de PII e retenção precisam de revisão.
- SQLite local não representa locks/concorrência de Postgres.
- A chave de dev é pública no código; produção exige chave própria, ainda não configurada no deploy.

## 7. O que o humano deve olhar

- S-01..S-03 antes de staging/produção.
- Política de retenção de evidência e logs.
- DPIA e bases legais (Q-38–Q-43).
- Se a decisão de usar SQLite no CI e Postgres em produção está refletida no pipeline real.
- Aprovação independente de Felipe; esta revisão foi feita pelo implementador.
