# service-request — security review

**Status:** não iniciado  
**Spec:** `spec.md`  
**Dono:** Anderson  
**Revisor:** Felipe  
**PR:** —  
**Commit:** —  
**Data:** —

## Veredito

`pendente`

## 1. Superfície avaliada

- auth:
- authz:
- PII:
- upload:
- admin:
- eventos/logs:
- dependências:

## 2. Checklist

- [ ] autenticação/sessão;
- [ ] autorização por recurso;
- [ ] CSRF/CSP/headers;
- [ ] rate limit;
- [ ] validação/encoding;
- [ ] segredos;
- [ ] PII em log/evento;
- [ ] localização/contato;
- [ ] retenção/consentimento;
- [ ] dependências;
- [ ] auditoria;
- [ ] incidente.

## 3. Bloqueios

- Nenhum até a implementação.

## 4. Riscos residuais

- Spec em revisão.

## 5. O que o humano deve olhar

- Aprovação da spec.
- Modelo de ameaças em `docs/security/threat-model.md`.
