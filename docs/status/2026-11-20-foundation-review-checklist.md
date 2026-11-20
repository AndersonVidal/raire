# Review da Spec `foundation-platform` — Checklist

**Status:** integrado; spec aprovada em 2026-11-20  
**Data:** 2026-11-20  
**Spec:** `specs/foundation-platform/spec.md`  
**Reviewers recomendados:** Victor (arquitetura/backend), Anderson (produto/PII), Felipe (segurança/QA)  
**Veredito:** **aprovada com ressalvas resolvidas**; must-fix e should-fix integrados na versão 1.1.

---

## 1. Resumo

A spec está bem estruturada: cobre contas, perfil, portfólio, catálogo, admin, privacidade, acessibilidade e observabilidade. O maior risco não é falta de requisito, e sim **ambiguidade de fronteira** com outras specs e **lacunas de aceite** em pontos de segurança.

Antes de aprovar, é necessário:

1. fechar as decisões Q-07 e Q-44–Q-50;
2. corrigir 8 pontos must-fix;
3. adicionar critérios de aceite para comportamentos já exigidos;
4. confirmar as fronteiras com `trust-reputation`, `service-request` e `local-discovery`.

### 1.1 Status de integração (2026-11-20)

| Item | Status | Onde foi integrado |
| --- | --- | --- |
| MF-01 verificação shell | resolvido | spec R-31, AC-45 |
| MF-02 troca de senha | resolvido | spec R-12, endpoint `/v1/me/password`, FND-15, AC-43 |
| MF-03 e-mail duplicado | resolvido | spec seção 9, AC-02/AC-40, FND-15 |
| MF-04 endpoints admin | resolvido | spec seção 7.1, API contracts, FND-17, AC-58/AC-59 |
| MF-05 contato | resolvido | spec R-13/R-14/R-14.1, data model, FND-16, AC-44 |
| MF-06 disponibilidade | resolvido | spec R-25, data model, FND-16, AC-55 |
| MF-07 localização precisa | resolvido | spec R-24/R-53, data model, FND-16, AC-54 |
| MF-08 fronteira trust | resolvido | spec 1.1, `docs/architecture/spec-boundaries.md` |
| SF-01..SF-08 | resolvido | spec, API contracts, data model, FND-15..FND-19 |
| AC-40..AC-60 | adicionados | `acceptance.md` |

**Resultado:** spec aprovada em 2026-11-20; registro em `docs/process/approvals.md`.

---

## 2. Must-fix (bloqueiam aprovação)

### MF-01 — Verificação: shell ou completa?

**Onde:** R-31, R-53, entidade `verifications`, não objetivo “verificação de identidade completa (apenas shell)”.

**Problema:** a spec cita `unverified`, `contact_verified`, `identity_verified`, mas diz que verificação de identidade é apenas shell. Não está claro o que a fundação entrega.

**Recomendação:** a fundação entrega:

- `unverified`;
- `contact_verified` (e-mail confirmado; telefone opcional);
- `identity_verified` fica **fora** do primeiro ciclo, como enum futuro/shell sem fluxo.

**Aceite:** adicionar AC explícito de que o selo exibe apenas o que foi verificado e que não promete antecedentes.

---

### MF-02 — Troca de senha sem endpoint

**Onde:** R-12 (“troca de senha revoga as demais”), seção 7.1.

**Problema:** não existe endpoint autenticado de troca de senha; existe apenas reset por token.

**Recomendação:** adicionar `POST /v1/me/password` com senha atual + nova, ou remover a menção a “troca de senha” do escopo.

**Aceite:** teste de que a troca revoga as outras sessões e mantém a atual.

---

### MF-03 — `EMAIL_ALREADY_USED` enumera conta

**Onde:** seção 9, código `EMAIL_ALREADY_USED`.

**Problema:** R-01/R-05 dizem para não revelar se o e-mail existe; um código específico de e-mail duplicado contradiz isso.

**Recomendação:** remover o código público. Cadastro com e-mail existente responde como sucesso genérico e envia aviso interno; o erro só pode existir em log interno sem PII.

**Aceite:** AC-02 cobre resposta genérica; adicionar teste de que a resposta não diferencia e-mail novo de existente.

---

### MF-04 — Endpoints administrativos ausentes

**Onde:** R-19, R-44, R-50 vs seção 7.1.

**Problema:** a spec exige operador gerenciando catálogo (`R-44`), feature flags (`R-50`) e solicitações de direitos (`R-19`), mas não lista endpoints para isso.

**Recomendação:** adicionar ou referenciar explicitamente:

```txt
GET/POST/PATCH  /v1/admin/categories
POST            /v1/admin/categories/{id}/deactivate
GET/PATCH       /v1/admin/feature-flags
GET             /v1/admin/account-requests
POST            /v1/admin/account-requests/{id}/resolve
```

**Aceite:** autorização, auditoria e paginação para cada um.

---

### MF-05 — Fonte de verdade do contato

**Onde:** `accounts.contact_email`, `accounts.contact_phone`, `accounts.contact_preference` e `professional_profiles.contact_preference`; R-13/R-14/R-24.

**Problema:** há duplicação de preferência de contato e não está claro qual vence; não há estado de verificação de telefone.

**Recomendação:**

- `accounts.contact_preference` é a preferência da conta;
- `professional_profiles.contact_preference` é a preferência pública do perfil e pode sobrescrever a da conta;
- contato público sempre mascarado até consentimento;
- telefone tem `phone_verified_at` ou registro em `verifications`;
- documentar precedência e fallback.

**Aceite:** AC de precedência e de mascaramento.

---

### MF-06 — Disponibilidade ambígua

**Onde:** R-25 (`texto estruturado`), `availability_text` e `availability_json`.

**Problema:** “texto estruturado” não é verificável e há dois campos.

**Recomendação:** escolher um formato para o MVP:

```txt
weekday (0–6) + start_time + end_time + timezone
```

com lista de janelas e texto livre apenas como complemento. Validar sobreposição e fuso.

**Aceite:** AC de janela válida, sobreposição e fuso.

---

### MF-07 — Criptografia e retenção da localização precisa

**Onde:** R-24, R-53; `precise_lat_enc`, `precise_lng_enc`.

**Problema:** a spec diz “cifrada” e “retenção definida”, mas não define algoritmo, gestão de chave, acesso ou job de expurgo.

**Recomendação:**

- cifrar com AES-256-GCM ou equivalente;
- chave fora do banco (env/secret manager);
- acesso restrito a casos de elegibilidade;
- retenção máxima de 90 dias;
- job de expurgo auditável;
- nunca retornar em API pública, log, evento ou read model.

**Aceite:** teste de que a API pública não retorna coordenada e de que o job expurga.

---

### MF-08 — Fronteira com `trust-reputation`

**Onde:** R-39 (denúncia de portfólio), R-46–R-51 (admin), R-31 (verificação).

**Problema:** a fundação cita ações de moderação e verificação que pertencem a `trust-reputation`. Sem fronteira, os tickets duplicam trabalho.

**Recomendação:**

- fundação: perfil, portfólio, catálogo, suspensão de perfil, auditoria base, shell de verificação de contato;
- `trust-reputation`: denúncias, fila, ações de moderação, recurso, verificação completa de identidade, reputação;
- R-39 vira “hook”: o alvo `portfolio` é aceito pelo módulo de denúncia, mas o fluxo completo é da outra spec.

**Aceite:** nenhum ticket de foundation implementa fila de denúncia/reputação.

---

## 3. Should-fix (não bloqueiam, mas evitam retrabalho)

### SF-01 — Paginação e filtros administrativos

`GET /v1/admin/professionals` e `GET /v1/admin/audit-events` precisam de cursor, filtro por status/período e limite. Sem isso, a operação não escala e a performance cai.

### SF-02 — Idempotência de cadastro/confirmação

Definir `Idempotency-Key` ou comportamento idempotente para signup, confirmação e reset. Evita duplicidade em retry de celular/rede instável.

### SF-03 — Rate limit de cadastro, reset e confirmação

A spec cita login, cadastro e reset na seção 10, mas o aceite só cobre login. Adicionar testes para cadastro (3/h por IP), reset (3/h) e confirmação (token).

### SF-04 — Métricas e auditoria de catálogo/flags

R-65 cita métricas de cadastro/publicação, mas não de catálogo/flags. Adicionar `category_changed`, `feature_flag_changed` e auditoria correspondente.

### SF-05 — Imagem de portfólio

Definir limite de tamanho, tipos aceitos, remoção de metadados e storage. Se não houver upload no primeiro ciclo, remover a imagem do escopo de FND-07 e deixar para depois.

### SF-06 — Termos e privacidade

R-57 exige exibição/versão, mas o aceite não testa a exibição dos documentos. Adicionar teste de que os links/versões aparecem no cadastro.

### SF-07 — Estados de UI

R-61 exige loading/vazio/erro/sucesso; o aceite cobre loading/vazio/erro. Adicionar sucesso e estado de sessão expirada.

### SF-08 — `CONTACT_PROTECTED` na fundação

O código existe na tabela de erros, mas a fundação não tem conexão/contato. Mover para `service-request` ou marcar como compartilhado.

---

## 4. Lacunas de aceite (adicionar ACs)

| Requisito | O que falta testar | AC sugerido |
| --- | --- | --- |
| R-02 | e-mail não confirmado não publica | AC-40 |
| R-07 | mesma conta como cliente e profissional | AC-41 |
| R-10 | CSRF ausente/inválido é rejeitado | AC-42 |
| R-12 | troca de senha revoga outras sessões | AC-43 |
| R-13/R-14 | preferência de contato e mascaramento | AC-44 |
| R-31 | selo de verificação honesto | AC-45 |
| R-32 | pausar/reativar perfil | AC-46 |
| R-36 | reordenar/despublicar portfólio | AC-47 |
| R-44/R-50 | categoria/flag auditada | AC-48 |
| R-57 | termos/privacidade exibidos | AC-49 |
| R-65 | métricas de cadastro/publicação | AC-50 |

> Renumerar conforme a matriz final; o importante é cobrir o comportamento.

---

## 5. Fronteiras a confirmar

| Tema | Foundation | Outra spec |
| --- | --- | --- |
| Busca/ranking | não faz | `local-discovery` |
| Pedidos/interesses/conexões | não faz | `service-request` |
| Denúncia/moderação/recurso | apenas hook e suspensão de perfil | `trust-reputation` |
| Reputação/avaliação | não faz | `trust-reputation` |
| Fairness | não faz | `fair-visibility` |
| UI de todas as jornadas | base e onboarding | `web-experience` |
| Verificação de identidade | shell/futuro | `trust-reputation` |

---

## 6. Dependências de decisão

Antes de aprovar:

- **Q-07** — um perfil por conta.
- **Q-44** — SQLite/Postgres.
- **Q-45** — server-rendered.
- **Q-46** — cookie + CSRF.
- **Q-47** — MFA TOTP.
- **Q-48** — outbox.
- **Q-49** — geo isolado.
- **Q-50** — flags em banco.

Também afetam a fundação, mas são de outros donos:

- Q-02/Q-11 (catálogo inicial) — Anderson;
- Q-08 (área) — Victor;
- Q-09 (confirmação de contato) — Felipe;
- Q-12/Q-13 (portfólio/imagem/remoção) — Anderson/Felipe;
- Q-10 (retenção de documento) — jurídico (fora do 1º ciclo).

---

## 7. Critérios para aprovar

- [x] 8 must-fix resolvidos;
- [x] ACs faltantes adicionados ou justificados;
- [x] fronteiras com trust/service/discovery confirmadas;
- [x] Q-07 e Q-44–Q-50 decididas;
- [x] ADRs 0001, 0002, 0004 e 0006 aceitos;
- [x] tickets FND-01..FND-19 refinados e com DoR;
- [x] riscos de SQLite/Postgres, PII e admin revisados por Felipe;
- [x] aprovação humana registrada em `docs/process/approvals.md`.

## 8. Veredito

**Aprovada em 2026-11-20** após a integração dos 8 must-fix e dos 8 should-fix. A spec está implementável e a fundação pode começar pelo `FND-01`.
