# ADR 0004 — Sessão web, CSRF, RBAC e MFA

**Status:** proposto  
**Data:** 2026-11-20  
**Decisores:** arquitetura, segurança, backend  
**Relacionado:** `docs/process/security-privacy.md`, `docs/security/threat-model.md`

## Contexto

O MVP é web, mobile-first e server-rendered. Precisa de autenticação segura, sessão revogável, proteção CSRF e autorização por papel/recurso. Admin lida com dados sensíveis e ações de moderação.

## Opções consideradas

1. **Cookie de sessão HttpOnly + CSRF + RBAC + MFA TOTP** — simples, seguro e compatível com server-rendered.
2. **JWT em localStorage** — pior para XSS, complica revogação.
3. **OAuth delegado** — dependência externa e complexidade para o MVP.

## Decisão

- Sessão em cookie `HttpOnly`, `Secure` (em HTTPS), `SameSite=Lax`, com rotação após login e revogação no logout.
- Token de sessão armazenado como hash.
- CSRF synchronizer token em métodos mutáveis.
- RBAC: `client`, `professional`, `operator`, `admin`.
- Autorização por recurso (ownership) em toda operação.
- MFA TOTP obrigatório para `admin` e `operator`.
- Senhas com Argon2id.

## Consequências

**Positivas:**

- compatível com HTML server-rendered;
- revogação simples;
- menor superfície de XSS que token em localStorage;
- MFA real para admin.

**Negativas:**

- CSRF exige disciplina em formulários;
- cookie requer HTTPS em produção;
- MFA adiciona atrito operacional.

**Mitigações:**

- middleware central de CSRF;
- testes de autorização cruzada;
- rate limit em login;
- auditoria de admin.

## Critérios de reversão

- necessidade de clientes móveis nativos;
- federação de identidade;
- requisito de SSO.

## Status

Proposto. Requer aprovação humana antes da implementação.
