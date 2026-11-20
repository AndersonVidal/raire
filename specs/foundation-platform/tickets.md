# foundation-platform — tickets

**Spec:** `spec.md`  
**Status:** aprovado; tickets prontos para implementação (DoR atendido)  
**Dono:** Victor  
**Revisores:** Anderson (produto/PII), Felipe (segurança/QA)

## Quadro

| Ticket | O quê | Dono | Depende de | Tam. | Status |
| --- | --- | --- | --- | --- | --- |
| FND-01 | Estrutura modular, contratos e regras de dependência | Victor | — | M | pronto |
| FND-02 | Modelo de dados, migrations e seed | Victor | FND-01 | G | pronto |
| FND-03 | Identidade, sessão, RBAC e rate limit | Victor | FND-01, FND-02 | G | pronto |
| FND-04 | Confirmação de e-mail, reset, MFA e recovery codes | Victor | FND-03 | M | pronto |
| FND-05 | Catálogo, sinônimos e admin de catálogo | Victor | FND-02 | M | pronto |
| FND-06 | Perfil profissional, área, disponibilidade e contato | Victor | FND-02, FND-03 | G | pronto |
| FND-07 | Portfólio profissional | Anderson | FND-06 | M | pronto |
| FND-08 | Privacidade: consentimento, PII, exportação/exclusão | Felipe | FND-03, FND-06 | G | pronto |
| FND-09 | Admin mínimo, suspensão e auditoria | Felipe | FND-03, FND-06 | M | pronto |
| FND-10 | Observabilidade, health e feature flags | Felipe | FND-01 | M | pronto |
| FND-11 | Web: cadastro, login e confirmação | Anderson | FND-03, FND-04 | G | pronto |
| FND-12 | Web: onboarding e perfil | Anderson | FND-06, FND-11 | G | pronto |
| FND-13 | E2E, acessibilidade e evidência da fundação | Felipe | FND-11, FND-12 | M | pronto |
| FND-14 | OpenAPI, README e handoff | Victor | todos | P | pronto |
| FND-15 | Segurança de conta e abuso | Victor | FND-03, FND-04 | M | pronto |
| FND-16 | Dados sensíveis: contato, disponibilidade e localização | Victor | FND-06 | M | pronto |
| FND-17 | Admin completo: catálogo, flags e direitos | Felipe | FND-05, FND-09 | M | pronto |
| FND-18 | Web: estados, termos e acessibilidade | Anderson | FND-11, FND-12 | M | pronto |
| FND-19 | QA de hardening e segurança | Felipe | FND-13, FND-15..FND-18 | M | pronto |

Tamanhos: `P` (até 1 dia), `M` (2–3), `G` (4–5), `GG` (quebrar).

> **Integração da revisão 1.1:** os tickets FND-15 a FND-19 cobrem os itens de hardening (segurança de conta, dados sensíveis, admin completo, web/estados e QA). Os tickets FND-03/04/05/06/08/09/10/12 devem ser lidos em conjunto com eles para evitar duplicidade.

---

## Detalhes

### FND-01 — Estrutura modular, contratos e regras de dependência

**Objetivo:** criar o esqueleto do monólito modular, contratos conceituais e regras de dependência.  
**Requisitos:** RNF-01, RNF-05, RNF-09.  
**Aceite:** AC-37.  
**Escopo:**

- `cmd/raire`, `internal/{identity,catalog,directory,demand,discovery,fairness,reputation,moderation,platform,contracts,api,web,app}`.
- `api/openapi.yaml` inicial com auth e perfil.
- Regras de dependência e lint.
- Composição de dependências.
- ADRs 0001, 0002, 0004, 0006 revisados.

**Fora de escopo:** features de negócio.  
**Testes:** build, teste de arquitetura, health.  
**DoD:** build verde, teste de dependência, OpenAPI válido.

---

### FND-02 — Modelo de dados, migrations e seed

**Objetivo:** modelar contas, sessões, consentimentos, perfis, categorias, portfólio, auditoria e outbox.  
**Requisitos:** R-01..R-57, RNF-05.  
**Aceite:** AC-01, AC-09..AC-18, AC-21, AC-24..AC-30.  
**Escopo:**

- migrations forward-only;
- constraints e índices;
- seed idempotente de categorias e bairros sintéticos;
- `audit_events` append-only;
- `outbox_events`;
- `feature_flags`;
- teste em banco vazio e com seed.

**Fora de escopo:** dados reais; PostGIS nesta fatia.  
**Testes:** migration, seed idempotente, constraints, rollback lógico.  
**Riscos:** diferença SQLite/Postgres; mitigar com testes de contrato.

---

### FND-03 — Identidade, sessão, RBAC e rate limit

**Objetivo:** cadastro, login, logout, sessão, papéis e limites.  
**Requisitos:** R-01..R-12.  
**Aceite:** AC-01..AC-08.  
**Escopo:**

- Argon2id;
- cookie `HttpOnly`/`SameSite`/`Secure` configurável;
- rotação de sessão;
- CSRF;
- RBAC e ownership;
- rate limit;
- respostas que não enumeram contas.

**Fora de escopo:** MFA (FND-04), confirmação de e-mail (FND-04).  
**Testes:** unitários de token/senha, integração de sessão, authz cruzada, rate limit.  
**Riscos:** CSRF mal configurado; mitigar com middleware central e testes.

---

### FND-04 — Confirmação de e-mail, reset e MFA

**Objetivo:** confirmar e-mail, recuperar senha e proteger admin/operador com TOTP.  
**Requisitos:** R-02, R-03, R-05, R-09.  
**Aceite:** AC-03, AC-06, AC-08, AC-33.  
**Escopo:**

- token de confirmação de uso único;
- token de reset de uso único;
- e-mail transacional via outbox/inbox de dev;
- TOTP para admin/operador;
- rate limit;
- resposta genérica.

**Fora de escopo:** SMS, verificação de identidade.  
**Testes:** token expirado/reutilizado, reset genérico, TOTP válido/inválido, rate limit.  
**Riscos:** vazamento de token em log; mitigar com redaction.

---

### FND-05 — Catálogo, sinônimos e admin de catálogo

**Objetivo:** disponibilizar categorias ativas, sinônimos e gestão por operador.  
**Requisitos:** R-40..R-45.  
**Aceite:** AC-17, AC-28..AC-30.  
**Escopo:**

- CRUD de categoria;
- sinônimos normalizados;
- ativação/desativação;
- seed versionado;
- auditoria de ação.

**Fora de escopo:** criação livre por usuário; sugestão de categoria (C).  
**Testes:** normalização, sinônimo, categoria inativa, seed idempotente, auditoria.  
**Riscos:** catálogo mal curado; mitigar com revisão de produto.

---

### FND-06 — Perfil profissional, área e disponibilidade

**Objetivo:** CRUD do perfil, publicação, área, disponibilidade e verificação de contato.  
**Requisitos:** R-20..R-33.  
**Aceite:** AC-09..AC-16.  
**Escopo:**

- `draft`/`published`/`paused`/`suspended`;
- cidade/bairro/raio;
- localização precisa cifrada;
- disponibilidade em texto estruturado;
- prévia pública sem PII;
- evento `directory.profile_published`;
- categoria inativa remove da busca.

**Fora de escopo:** busca/ranking; portfólio (FND-07).  
**Testes:** invariantes de publicação, área, disponibilidade, contrato público.  
**Riscos:** perfil incompleto; mitigar com DoR e validação.

---

### FND-07 — Portfólio profissional

**Objetivo:** permitir que o profissional apresente trabalhos.  
**Requisitos:** R-34..R-39.  
**Aceite:** AC-24..AC-27.  
**Escopo:**

- criar/editar/reordenar/despublicar;
- imagem opcional com validação;
- status;
- exibição pública;
- denúncia.

**Fora de escopo:** verificação de trabalho; antes/depois; vídeo.  
**Testes:** CRUD, status, imagem inválida, PII, contrato.  
**Riscos:** imagem de terceiros; mitigar com declaração e denúncia.

---

### FND-08 — Privacidade: consentimento, PII, exportação/exclusão

**Objetivo:** cumprir os requisitos de privacidade do MVP.  
**Requisitos:** R-13..R-19, R-52..R-57.  
**Aceite:** AC-18..AC-23.  
**Escopo:**

- consentimentos versionados;
- preferência de contato;
- exportação JSON;
- exclusão/anonimização;
- redaction de logs;
- teste de contrato público;
- registro de solicitações.

**Fora de escopo:** DPIA completa (documento separado); documentos.  
**Testes:** integração, redaction, exportação, exclusão, autorização.  
**Riscos:** PII em log; mitigar com teste e revisão.

---

### FND-09 — Admin mínimo, suspensão e auditoria

**Objetivo:** operação consegue suspender/reativar perfil e auditar.  
**Requisitos:** R-46..R-51.  
**Aceite:** AC-31..AC-33.  
**Escopo:**

- listagem de perfis;
- suspender/reativar com motivo;
- MFA;
- `audit_events`;
- revogação de sessão;
- efeito imediato na busca.

**Fora de escopo:** moderação de denúncias (trust-reputation).  
**Testes:** autorização, auditoria, efeito na busca, MFA.  
**Riscos:** abuso de admin; mitigar com auditoria e menor privilégio.

---

### FND-10 — Observabilidade, health e feature flags

**Objetivo:** operar e diagnosticar a fundação.  
**Requisitos:** R-64..R-68.  
**Aceite:** AC-37..AC-39.  
**Escopo:**

- logs JSON, correlation ID;
- métricas de cadastro/publicação;
- health/readiness;
- feature flags com default seguro;
- redaction.

**Fora de escopo:** tracing distribuído.  
**Testes:** health, log redaction, flag default, métricas.  
**Riscos:** log com PII; mitigar com teste.

---

### FND-11 — Web: cadastro, login e confirmação

**Objetivo:** fluxo de conta acessível e mobile-first.  
**Requisitos:** R-01..R-12, R-58..R-63.  
**Aceite:** AC-01..AC-08, AC-34..AC-36.  
**Escopo:**

- telas de cadastro, login, confirmação e reset;
- estados de erro e carregamento;
- teclado e leitor de tela;
- CSRF em formulários.

**Fora de escopo:** onboarding de perfil (FND-12).  
**Testes:** componentes/render, E2E, axe, 360 px.  
**Riscos:** formulário longo; mitigar com etapas curtas.

---

### FND-12 — Web: onboarding e perfil

**Objetivo:** profissional publica perfil sem ajuda.  
**Requisitos:** R-20..R-33, R-58..R-63.  
**Aceite:** AC-09..AC-16, AC-34..AC-36.  
**Escopo:**

- formulário em etapas curtas;
- seleção de categoria e área;
- disponibilidade;
- prévia;
- edição e pausa.

**Fora de escopo:** portfólio completo (FND-07); busca.  
**Testes:** render, E2E, acessibilidade, 360 px.  
**Riscos:** abandono; mitigar com salvamento progressivo.

---

### FND-13 — E2E, acessibilidade e evidência da fundação

**Objetivo:** provar as jornadas da fundação.  
**Requisitos:** R-58..R-63, RNF-10, RNF-15.  
**Aceite:** AC-34..AC-36, AC-37.  
**Escopo:**

- E2E de cadastro → onboarding → publicação;
- E2E de admin suspendendo perfil;
- axe e teclado;
- evidência em `evidence/`;
- relatório de QA.

**Fora de escopo:** performance de busca.  
**Testes:** Playwright/Go E2E, axe, screenshots.  
**Riscos:** ambiente instável; mitigar com seed e dados sintéticos.

---

### FND-14 — OpenAPI, README e handoff

**Objetivo:** deixar a fundação utilizável pelas próximas frentes.  
**Requisitos:** RNF-01, RNF-18.  
**Aceite:** todos.  
**Escopo:**

- OpenAPI atualizado;
- README de execução local;
- ADRs;
- runbook mínimo;
- revisão de spec e tickets.

**Testes:** validação OpenAPI, links, build.  
**DoD:** documentação revisada e mergeada.

---

### FND-15 — Segurança de conta e abuso

**Objetivo:** fechar lacunas de autenticação, CSRF, rate limit, idempotência e recuperação de MFA.  
**Requisitos:** R-08..R-12, R-14.1.  
**Aceite:** AC-40..AC-44, AC-51, AC-52.  
**Escopo:**

- `POST /v1/me/password` com senha atual;
- revogação das outras sessões;
- códigos de recuperação TOTP (hash, uso único);
- CSRF em toda mutação autenticada;
- rate limit de signup/reset/resend/confirmação;
- idempotência de signup/confirmação/reset;
- resposta genérica para e-mail duplicado.

**Fora de escopo:** WebAuthn, SSO, MFA por SMS.  
**Testes:** unitários de token/senha, integração de sessão/CSRF, rate limit, idempotência, recuperação.  
**Riscos:** lockout de admin; mitigar com recovery codes e auditoria.

---

### FND-16 — Dados sensíveis: contato, disponibilidade e localização

**Objetivo:** fechar as ambiguidades de contato, disponibilidade e localização precisa.  
**Requisitos:** R-13..R-14.1, R-24..R-25, R-53.  
**Aceite:** AC-44, AC-45, AC-54, AC-55.  
**Escopo:**

- precedência conta → perfil;
- `phone_verified_at` e selo de contato;
- janelas de disponibilidade com dia/hora/fuso e validação de sobreposição;
- cifra AES-256-GCM para localização precisa;
- chave fora do banco;
- retenção de 90 dias e job de expurgo auditável;
- mascaramento de contato.

**Fora de escopo:** polígono, geocoding, verificação de identidade.  
**Testes:** unitários de disponibilidade, integração de contato, segurança/criptografia, job de expurgo.  
**Riscos:** vazamento de coordenada; mitigar com testes de contrato e acesso restrito.

---

### FND-17 — Admin completo: catálogo, flags e direitos

**Objetivo:** completar a operação administrativa da fundação.  
**Requisitos:** R-19, R-44, R-46..R-50.  
**Aceite:** AC-48, AC-53, AC-58, AC-59.  
**Escopo:**

- CRUD de categorias no admin;
- ativação/desativação com auditoria e métrica;
- feature flags com default `off`, dono e revisão;
- listagem/resolução de `account_requests`;
- paginação e filtros em perfis/auditoria;
- métricas de catálogo/flags.

**Fora de escopo:** moderação de denúncias (trust), política de ranking (fairness).  
**Testes:** integração, authz, auditoria, paginação, métricas.  
**Riscos:** abuso de admin; mitigar com MFA, menor privilégio e auditoria.

---

### FND-18 — Web: estados, termos e acessibilidade

**Objetivo:** fechar estados de UI, exibição de termos e copy da fundação.  
**Requisitos:** R-57, R-58..R-63.  
**Aceite:** AC-49, AC-57, AC-60.  
**Escopo:**

- estado de sucesso e sessão expirada;
- exibição/versão de termos e privacidade;
- validação e mensagens de imagem;
- copy PT-BR revisada;
- teclado, foco, 360 px e reduced motion.

**Fora de escopo:** telas de busca/pedido/avaliação (web-experience).  
**Testes:** E2E, a11y, copy review, 360 px.  
**Riscos:** copy enganosa; mitigar com revisão de produto.

---

### FND-19 — QA de hardening e segurança

**Objetivo:** provar os novos comportamentos de segurança, privacidade e admin.  
**Requisitos:** RNF-08, RNF-09, RNF-14.  
**Aceite:** AC-40..AC-60 (aplicáveis).  
**Escopo:**

- testes de CSRF, rate limit, idempotência;
- PII em resposta/log/evento;
- criptografia/expurgo de localização;
- paginação/filtros admin;
- SQLite/Postgres contract;
- evidência de hardening.

**Fora de escopo:** pentest externo.  
**Testes:** unit, integração, contrato, segurança, a11y.  
**Riscos:** cobertura incompleta; mitigar com matriz de aceite e revisão de Felipe.
