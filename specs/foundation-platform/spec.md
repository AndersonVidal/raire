# foundation-platform

**Status:** aprovada  
**Versão:** 1.1  
**RFs:** RF-01, RF-02, RF-03, RF-04, RF-05, RF-15 (contato), RF-19  
**RNFs:** RNF-01, RNF-05, RNF-08, RNF-09, RNF-13, RNF-14, RNF-18  
**Módulos:** identity, catalog, directory, moderation (base), platform, web  
**Depende de:** ADR 0001, ADR 0002, ADR 0004, ADR 0006  
**Dono:** Victor  
**Revisores:** Anderson (produto/PII), Felipe (segurança/QA)  
**Aprovação humana:** aprovada em 2026-11-20 — registro em `docs/process/approvals.md`  
**Decisões aplicadas:** Q-07 e Q-44–Q-50 (`docs/product/decisions-2026-11-20.md`)

---

## 1. Contexto

O RAIRE ainda não existe como produto executável. Esta spec cria a fundação mínima para que um profissional publique um perfil utilizável e um cliente possa encontrar e abrir esse perfil. Sem essa base, busca, pedidos, reputação e fairness não têm onde existir.

A fatia cobre:

1. contas, autenticação e papéis;
2. perfil profissional, área de atendimento e disponibilidade;
3. portfólio básico;
4. catálogo curado de serviços;
5. admin mínimo, MFA e auditoria;
6. base de privacidade, acessibilidade e observabilidade.

**Não cobre:** busca/ranking (`local-discovery`), pedidos (`service-request`), avaliações (`trust-reputation`) e camada completa de fairness (`fair-visibility`).

### 1.1 Fronteiras com outras specs

A fundação é dona da conta, do perfil, do catálogo e do admin base. Ela **não implementa** os fluxos das outras specs; expõe contratos e hooks.

| Tema | `foundation-platform` | Outra spec | Regra |
| --- | --- | --- | --- |
| Conta, sessão, papéis | implementa | — | identity é dono |
| Perfil profissional, área, disponibilidade | implementa | `web-experience` consome | directory é dono |
| Portfólio | implementa CRUD e status | `web-experience` exibe; `trust-reputation` denuncia | foundation não modera |
| Catálogo | implementa catálogo e seed | `local-discovery` consome | catalog é dono |
| Busca e ranking | não implementa | `local-discovery` | foundation emite `profile_published`/`profile_updated` |
| Pedidos, interesses, conexões | não implementa | `service-request` | foundation só fornece conta/perfil/e-mail confirmado |
| Contato liberado | não implementa | `service-request` | foundation guarda preferência e mascara; liberação é da conexão |
| Avaliações, reputação | não implementa | `trust-reputation` | foundation não calcula nota |
| Denúncia e moderação | apenas hook de alvo e suspensão de perfil | `trust-reputation` | fila, ação, recurso e reputação são da trust |
| Verificação de identidade | shell/futuro | `trust-reputation` | fundação entrega `unverified`/`contact_verified`; `identity_verified` fica fora do 1º ciclo |
| Fairness | não implementa | `fair-visibility` | foundation não ordena resultados |
| UI completa | base, conta e onboarding | `web-experience` | componentes e copy são compartilhados |
| Privacidade/LGPD | consentimento, exportação, exclusão, PII | `docs/security/privacy-and-lgpd.md` | fundação implementa os controles; jurídico valida bases |

**Regras de fronteira:**

1. Nenhum contexto importa `adapters` de outro; integração por interface de aplicação, evento de outbox ou read model explícito.
2. A fundação não cria tabela de pedido, interesse, conexão, avaliação, denúncia ou política de fairness.
3. `trust-reputation` pode ler `professional_profiles` e `accounts` por interface, mas não escreve nelas.
4. `local-discovery` lê o read model de perfil; não altera status de perfil.
5. `web-experience` não contém regra de negócio; consome API/HTML.
6. Toda ação administrativa da fundação (suspender perfil, categoria, flag) é auditada; moderação de conteúdo é da trust.

---

## 2. Objetivo e não objetivo

**Objetivo:** permitir que um profissional crie conta, confirme e-mail, publique um perfil com categoria/área/disponibilidade e que um visitante veja o perfil público sem PII. A operação deve conseguir suspender um perfil com auditoria.

**Não objetivo:**

- busca e ranking;
- pedidos e conexões;
- avaliação e reputação;
- pagamento;
- app nativo;
- verificação de identidade completa (apenas shell);
- multi-perfil;
- múltiplas cidades.

---

## 3. Princípios e restrições

- Cadastro e onboarding funcionam em 360 px, com uma mão.
- Nenhum dado de contato é público antes de consentimento.
- Localização precisa nunca aparece em resposta pública.
- Papéis explícitos; um usuário pode ter perfil de cliente e profissional, com permissões separadas.
- Catálogo é curado pela operação; usuário não cria categoria livre no MVP.
- Admin usa MFA, menor privilégio e auditoria append-only.
- Nenhuma regra de negócio em handler HTTP.
- Todo endpoint versionado em `/v1` e documentado.
- Senha com Argon2id; sessão revogável; CSRF em métodos mutáveis.
- Nenhum PII em log/evento.

---

## 4. Requisitos

### A. Identidade e sessão

- **R-01** Cadastro com e-mail, senha e aceite explícito de termos e privacidade, com versão do aceite registrada.
- **R-02** E-mail confirmado antes de publicar perfil, criar pedido, enviar interesse ou avaliar.
- **R-03** Confirmação por token de uso único com expiração.
- **R-04** Login cria sessão em cookie `HttpOnly`, `Secure` (em HTTPS) e `SameSite`, com rotação após autenticação e revogação no logout.
- **R-05** Recuperação de senha usa token de uso único, curta duração, e não revela se o e-mail existe.
- **R-06** Papéis: `client`, `professional`, `operator`, `admin`. Autorização por papel e por recurso.
- **R-07** Um usuário pode atuar como cliente e profissional; perfil profissional é recurso separado.
- **R-08** Tentativas de login inválidas são limitadas por conta e IP; resposta não revela qual credencial falhou.
- **R-09** Admin e operador exigem MFA TOTP, com códigos de recuperação de uso único. Ações são auditadas.
- **R-10** CSRF obrigatório em POST/PATCH/DELETE autenticado.
- **R-11** Sessão expirada/revogada não autentica.
- **R-12** Logout revoga a sessão atual; troca de senha autenticada revoga as demais sessões; reset de senha revoga todas.

### B. Conta e preferências

- **R-13** Usuário consulta e atualiza nome de exibição, telefone e preferência de contato.
- **R-14** Preferência de contato: `platform`, `email`, `phone`. A preferência da conta é o padrão; a do perfil pode sobrescrever. O contato público é sempre mascarado até consentimento.
- **R-14.1** Telefone só é considerado verificado após confirmação; enquanto não verificado, não aparece em contato público nem conta como selo.
- **R-15** Consentimentos versionados: termos, privacidade, pesquisa/analytics, marketing.
- **R-16** Usuário revoga consentimento não obrigatório.
- **R-17** Usuário solicita exportação dos próprios dados.
- **R-18** Usuário solicita exclusão/anonimização, com ressalvas legais.
- **R-19** Solicitações de direitos são registradas e tratáveis por operador.

### C. Perfil profissional

- **R-20** Profissional cria perfil com nome de exibição, bio, categorias, tipo de atendimento, disponibilidade e preferência de contato.
- **R-21** Tipo de atendimento: `in_person`, `remote`, `hybrid`.
- **R-22** Presencial/híbrido exige cidade/bairro e raio; remoto não exige.
- **R-23** Raio mínimo/máximo configurável (default 1–50 km).
- **R-24** Localização precisa é armazenada separada, cifrada com AES-256-GCM (ou equivalente), com chave fora do banco, acesso restrito e expiração; nunca pública.
- **R-25** Disponibilidade é representada por janelas de dia da semana, hora de início, hora de fim e fuso; texto livre é apenas complemento. Janelas sobrepostas são normalizadas ou rejeitadas.
- **R-26** Status do perfil: `draft`, `published`, `paused`, `suspended`.
- **R-27** Só `published` e elegível aparece na busca.
- **R-28** Publicação exige nome, ao menos uma categoria ativa, tipo e área (se aplicável).
- **R-29** Prévia pública mostra exatamente o que será exibido, sem PII.
- **R-30** Edição relevante gera evento de auditoria e publica evento de outbox para o read model de descoberta; a fundação não mantém o read model de busca.
- **R-31** Verificação é selo separado: `unverified`, `contact_verified`, `identity_verified`. No primeiro ciclo, a fundação entrega `unverified` e `contact_verified`; `identity_verified` é enum futuro/shell sem fluxo. O selo não promete antecedentes.
- **R-32** Profissional pausa e reativa o perfil.
- **R-33** Categoria desativada remove perfil da busca até ajuste.

### D. Portfólio

- **R-34** Profissional cria item com título, descrição, categoria e imagem opcional.
- **R-35** Item tem status `draft`, `published`, `removed`.
- **R-36** Profissional edita, reordena, despublica e remove item.
- **R-37** Portfólio público mostra apenas itens publicados.
- **R-38** Conteúdo não é apresentado como experiência verificada.
- **R-39** Conteúdo pode ser denunciado: a fundação aceita o alvo `portfolio` no hook de denúncia; a fila, ação e recurso são de `trust-reputation`.

### E. Catálogo

- **R-40** Catálogo tem categorias/subcategorias com slug, nome, descrição, sinônimos e status.
- **R-41** Cada serviço declara tipo e categoria.
- **R-42** Categorias inativas não aparecem no onboarding nem na busca.
- **R-43** Seed inicial é versionado, idempotente e revisado por produto.
- **R-44** Operador cria/edita/desativa categoria; toda ação é auditada e registra métrica.
- **R-45** Busca no catálogo suporta sinônimo simples.

### F. Admin e auditoria

- **R-46** Operador lista perfis e vê status.
- **R-47** Operador suspende/reativa perfil com motivo e evidência.
- **R-48** Suspensão remove perfil da busca imediatamente e revoga sessões do perfil.
- **R-49** Toda ação administrativa grava ator, alvo, ação, motivo, timestamp e correlação em `audit_events` append-only.
- **R-50** Admin gerencia catálogo e feature flags, com paginação, filtros, auditoria e métrica; política de ranking fica em `fair-visibility`.
- **R-51** Acesso admin exige MFA e menor privilégio.

### G. Privacidade e LGPD

- **R-52** Consentimentos versionados e revogáveis.
- **R-53** Localização precisa cifrada, acesso restrito, retenção máxima de 90 dias e job de expurgo auditável.
- **R-54** API pública nunca retorna e-mail, telefone, documento ou coordenada precisa.
- **R-55** Logs e eventos não contêm PII; redaction testada.
- **R-56** Exportação e exclusão implementadas, com registro de solicitação e resolução.
- **R-57** Termos e política de privacidade exibidos e versionados no cadastro.

### H. Acessibilidade e experiência

- **R-58** Cadastro, onboarding e perfil funcionam por teclado e leitor de tela.
- **R-59** Layout funciona em 360 px, sem rolagem horizontal.
- **R-60** Formulários têm rótulos persistentes, ajuda contextual e erro específico.
- **R-61** Estados de carregando, vazio, erro e sucesso explícitos.
- **R-62** `prefers-reduced-motion` respeitado.
- **R-63** Textos em PT-BR, linguagem simples, sem jargão.

### I. Observabilidade

- **R-64** Logs JSON com `correlation_id`, sem PII.
- **R-65** Métricas de cadastro, confirmação, publicação, erro, catálogo e feature flags.
- **R-66** Auditoria de ações administrativas, mudanças de status, catálogo e flags.
- **R-67** Health check e readiness do banco.
- **R-68** Feature flags com default seguro e cache curto.

---

## 5. Regras de negócio e invariantes

- e-mail único normalizado; cadastro com e-mail existente responde de forma genérica;
- senha nunca em texto;
- token de confirmação/reset de uso único;
- um perfil profissional por conta no MVP;
- perfil publicado exige categoria ativa e área/remoto;
- área de atendimento válida;
- disponibilidade com janelas válidas, sem sobreposição não resolvida;
- localização precisa separada, cifrada, com retenção e expurgo;
- perfil de contato: conta define padrão, perfil pode sobrescrever;
- telefone só conta como verificado após confirmação;
- suspensão remove da busca e revoga sessão;
- auditoria append-only;
- categoria inativa não pode ser nova associação;
- perfil incompleto não aparece na busca;
- contato nunca público antes de consentimento;
- toda mutação autenticada exige CSRF;
- feature flag desconhecida resolve `off`.

---

## 6. Estados e transições

**Profile:**

```txt
draft -> published -> paused -> published
published/paused -> suspended -> published (reativação com revisão)
```

**Portfolio item:**

```txt
draft -> published -> removed
published -> draft (despublicar)
```

**Account:**

```txt
active -> suspended -> active
active -> deleted (anonimizado)
```

**Verification:**

```txt
not_started -> pending -> verified/rejected -> expired
```

---

## 7. Contratos e dados

### 7.1 Endpoints

```txt
POST   /v1/auth/signup
POST   /v1/auth/login
POST   /v1/auth/logout
POST   /v1/auth/confirm-email
POST   /v1/auth/resend-confirmation
POST   /v1/auth/password-reset/request
POST   /v1/auth/password-reset/confirm

GET    /v1/me
PATCH  /v1/me
POST   /v1/me/password
GET    /v1/me/consents
POST   /v1/me/consents/{kind}/revoke
GET    /v1/me/export
POST   /v1/me/delete
GET    /v1/me/requests

GET    /v1/categories
GET    /v1/categories/{slug}

GET    /v1/professionals/me
POST   /v1/professionals/me
PATCH  /v1/professionals/me
POST   /v1/professionals/me/publish
POST   /v1/professionals/me/pause
POST   /v1/professionals/me/resume
GET    /v1/professionals/{id}
GET    /v1/professionals/{id}/preview          # dono
GET    /v1/professionals/{id}/portfolio
POST   /v1/professionals/me/portfolio
PATCH  /v1/portfolio/{id}                     # dono
DELETE /v1/portfolio/{id}                     # dono

GET    /v1/admin/professionals                # cursor, status, q
POST   /v1/admin/professionals/{id}/suspend   # reason, evidence
POST   /v1/admin/professionals/{id}/reactivate
GET    /v1/admin/categories
POST   /v1/admin/categories
PATCH  /v1/admin/categories/{id}
POST   /v1/admin/categories/{id}/deactivate
GET    /v1/admin/feature-flags
PATCH  /v1/admin/feature-flags/{key}
GET    /v1/admin/account-requests
POST   /v1/admin/account-requests/{id}/resolve
GET    /v1/admin/audit-events                 # cursor, actor, target, period
GET    /v1/healthz
GET    /v1/readyz
```

### 7.2 Entidades

- `accounts`, `sessions`, `email_tokens`, `mfa_recovery_codes`, `consents`, `account_requests`
- `professional_profiles`, `profile_categories`, `portfolio_items`
- `categories`, `category_synonyms`
- `verifications` (shell)
- `audit_events`, `outbox_events`, `feature_flags`, `idempotency_keys`

### 7.3 Payloads

Detalhes em `docs/architecture/api-contracts.md`.

---

## 8. Casos de borda

- Publicar perfil sem categoria → `PROFILE_INCOMPLETE`.
- Área com raio fora do limite → `INVALID_SERVICE_AREA`.
- Perfil remoto sem área → permitido.
- Perfil híbrido sem área → bloqueado.
- Categoria desativada após publicação → perfil sai da busca até ajuste.
- Suspensão durante sessão ativa → sessão revogada no próximo request.
- E-mail duplicado → resposta genérica de sucesso; não revelar existência.
- E-mail não confirmado tenta publicar/criar pedido/interesse/avaliar → `EMAIL_NOT_CONFIRMED`.
- Token expirado/reutilizado → `TOKEN_EXPIRED`/`TOKEN_INVALID`.
- Telefone não verificado → não entra em contato público nem selo.
- Preferência de contato do perfil ausente → usar a da conta.
- Disponibilidade com sobreposição → normalizar ou `INVALID_AVAILABILITY`.
- Localização precisa ausente → usar cidade/bairro; não inventar coordenada.
- Localização precisa expirada → expurgar e manter apenas área aproximada.
- Upload de imagem inválida/grande → `INVALID_IMAGE`.
- Troca de senha sem senha atual → `INVALID_PASSWORD`.
- CSRF ausente/inválido → `CSRF_INVALID`.
- Usuário excluído tenta login → resposta genérica.
- Feature flag desconhecida → `off`.
- Operador tenta ação de admin → `FORBIDDEN`.
- Admin sem MFA → `MFA_REQUIRED`.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `VALIDATION_ERROR` | campo inválido | "Confira os campos destacados." |
| `UNAUTHENTICATED` | sem sessão | "Entre para continuar." |
| `FORBIDDEN` | sem permissão | "Você não tem acesso a esta ação." |
| `PROFILE_INCOMPLETE` | publicação incompleta | "Complete nome, categoria e área para publicar." |
| `INVALID_SERVICE_AREA` | área/raio inválido | "Informe uma área de atendimento válida." |
| `EMAIL_NOT_CONFIRMED` | e-mail não confirmado | "Confirme seu e-mail para continuar." |
| `TOKEN_EXPIRED` | token expirado | "Este link expirou. Solicite outro." |
| `TOKEN_INVALID` | token inválido/reutilizado | "Este link não é válido. Solicite outro." |
| `INVALID_AVAILABILITY` | janela inválida/sobreposta | "Confira os dias e horários." |
| `INVALID_IMAGE` | imagem inválida/grande | "Escolha uma imagem JPG, PNG ou WebP de até 2 MB." |
| `INVALID_PASSWORD` | senha atual errada | "A senha atual não confere." |
| `CSRF_INVALID` | CSRF ausente/inválido | "Atualize a página e tente novamente." |
| `MFA_REQUIRED` | admin sem MFA | "Confirme o código de verificação." |
| `RATE_LIMITED` | limite | "Muitas tentativas. Tente novamente em alguns minutos." |
| `INTERNAL_ERROR` | erro interno | "Algo deu errado. Tente novamente." |

---

## 10. Privacidade, segurança e fairness

- **PII:** e-mail, telefone, documento e localização precisa nunca em resposta pública.
- **Authz:** toda operação verifica ator e recurso; admin com MFA TOTP e recovery codes.
- **Rate limit:** login 5/min; cadastro 3/h; reset 3/h; confirmação/reenvio 5/h; admin por operador.
- **CSRF:** obrigatório em POST/PATCH/DELETE autenticado; token sincronizado com a sessão.
- **Idempotência:** signup, confirmação, reset e criação de portfólio aceitam retry seguro.
- **Criptografia:** localização precisa com AES-256-GCM e chave fora do banco; retenção 90 dias; expurgo auditável.
- **Contato:** preferência da conta como padrão; perfil pode sobrescrever; telefone só verificado após confirmação.
- **Fairness:** esta spec não ordena resultados; apenas garante elegibilidade e emite eventos.
- **Acessibilidade:** R-58 a R-63.

---

## 11. Observabilidade e métricas

- eventos: `account_signup_started`, `account_signup_completed`, `email_confirmed`, `profile_created`, `profile_published`, `profile_updated`, `category_changed`, `feature_flag_changed`, `password_changed`.
- métricas: cadastro, confirmação, publicação, erro, tempo de onboarding, catálogo, flags.
- logs: `correlation_id`, ator pseudonimizado, sem PII.
- alertas: taxa de erro de cadastro, falha de e-mail, falha de migration, backlog de outbox, pico de rate limit.
- auditoria: ações administrativas, mudanças de status, catálogo, flags e solicitações de direitos.

---

## 12. Fora de escopo

- busca e ranking;
- pedidos, interesses e conexões;
- avaliações completas;
- verificação de identidade completa;
- pagamento;
- app nativo;
- chat;
- múltiplas cidades/idiomas;
- multi-perfil.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| onboarding longo afasta profissionais | baixa oferta | formulário curto, salvamento progressivo, teste com usuários |
| localização precisa vaza | risco físico | PII separada, cifra, testes de contrato |
| catálogo mal curado | busca ruim | seed revisado, sinônimos, feedback |
| admin abusa do poder | injustiça | MFA, auditoria, menor privilégio, recurso |
| perfil incompleto polui busca | relevância | status draft + requisitos de publicação |
| SQLite ≠ Postgres | bug em produção | repositórios isolados, testes de contrato, staging |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- cadastro, confirmação, login, logout e reset;
- onboarding presencial, remoto e híbrido;
- validação de área e disponibilidade;
- perfil draft vs. published;
- catálogo ativo/inativo e sinônimos;
- admin suspende e auditoria registra;
- API pública sem PII;
- exportação/exclusão;
- acessibilidade e mobile.

---

## 15. Gates de review específicos

- [ ] nenhuma PII em resposta pública, log ou evento;
- [ ] authz por recurso em perfil, portfólio, admin e direitos;
- [ ] rate limit em auth, cadastro, reset, confirmação e admin;
- [ ] CSRF em toda mutação autenticada;
- [ ] MFA TOTP + recovery codes e auditoria de admin;
- [ ] e-mail duplicado não enumera conta;
- [ ] localização precisa cifrada, com retenção e expurgo;
- [ ] disponibilidade validada (dia/hora/fuso/sobreposição);
- [ ] acessibilidade das jornadas de onboarding, admin e privacidade;
- [ ] migração e rollback;
- [ ] seed de catálogo idempotente;
- [ ] OpenAPI atualizado;
- [ ] SQLite/Postgres testados;
- [ ] paginação/filtros de admin;
- [ ] métricas e auditoria de catálogo/flags;
- [ ] fronteiras com trust/service/discovery respeitadas.

---

## 16. Perguntas para o humano

**Status:** decididas em 2026-11-20; ver `docs/product/decisions-2026-11-20.md`.

- [x] Q-07: um perfil por conta no MVP.
- [x] Q-08: cidade/bairro + raio.
- [x] Q-09: e-mail; telefone opcional.
- [x] Q-10: 90 dias após decisão (fora do 1º ciclo).
- [x] Q-12: portfólio com imagem opcional e validada.
- [x] Q-44: SQLite local/Postgres produção.
- [x] Q-45: web server-rendered.
- [x] Q-46: sessão por cookie + CSRF.
- [x] Q-47: MFA TOTP + recovery codes.
- [x] Q-50: feature flags em banco.

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| R-01..R-12 | RF-01 | FND-03, FND-04, FND-15 | AC-01..AC-08, AC-40..AC-43, AC-51, AC-52 |
| R-13..R-19 | RF-02, RF-19 | FND-08, FND-15, FND-16, FND-17 | AC-18..AC-23, AC-44, AC-58 |
| R-20..R-33 | RF-03 | FND-06, FND-16 | AC-09..AC-16, AC-45, AC-46, AC-54, AC-55 |
| R-34..R-39 | RF-04 | FND-07, FND-18 | AC-24..AC-27, AC-47, AC-56 |
| R-40..R-45 | RF-05 | FND-05, FND-17 | AC-17, AC-28..AC-30, AC-48, AC-59 |
| R-46..R-51 | RF-18 | FND-09, FND-17 | AC-31..AC-33, AC-48, AC-53, AC-59 |
| R-52..R-57 | RF-19 | FND-08, FND-16, FND-17 | AC-18..AC-23, AC-44, AC-49, AC-54, AC-58 |
| R-58..R-63 | RNF-10..12 | FND-11, FND-12, FND-18 | AC-34..AC-36, AC-57, AC-60 |
| R-64..R-68 | RNF-13, RNF-14 | FND-10, FND-17 | AC-37..AC-39, AC-50, AC-53 |

### 17.1 Cobertura da revisão

- Must-fix MF-01..MF-08 integrados na versão 1.1.
- Should-fix SF-01..SF-08 integrados na versão 1.1.
- ACs adicionais AC-40..AC-60 cobrem os comportamentos que faltavam.
- Fronteiras com `trust-reputation`, `service-request`, `local-discovery` e `web-experience` explicitadas na seção 1.1 e em `docs/architecture/spec-boundaries.md`.
- Decisões Q-07 e Q-44–Q-50 registradas em `docs/product/decisions-2026-11-20.md`.
