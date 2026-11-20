# foundation-platform

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** RF-01, RF-02, RF-03, RF-04, RF-05, RF-15 (contato), RF-19  
**RNFs:** RNF-01, RNF-05, RNF-08, RNF-09, RNF-13, RNF-14, RNF-18  
**Módulos:** identity, catalog, directory, moderation (base), platform, web  
**Depende de:** ADR 0001, ADR 0002, ADR 0004, ADR 0006  
**Dono:** Victor  
**Revisores:** Anderson (produto/PII), Felipe (segurança/QA)  
**Aprovação humana:** pendente

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
- **R-09** Admin e operador exigem MFA TOTP. Ações são auditadas.
- **R-10** CSRF obrigatório em POST/PATCH/DELETE autenticado.
- **R-11** Sessão expirada/revogada não autentica.
- **R-12** Logout revoga a sessão atual; troca de senha revoga as demais.

### B. Conta e preferências

- **R-13** Usuário consulta e atualiza nome de exibição, telefone e preferência de contato.
- **R-14** Preferência de contato: `platform`, `email`, `phone`.
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
- **R-24** Localização precisa é armazenada separada e cifrada; nunca pública.
- **R-25** Disponibilidade é texto estruturado no MVP, com modelo evolutivo.
- **R-26** Status do perfil: `draft`, `published`, `paused`, `suspended`.
- **R-27** Só `published` e elegível aparece na busca.
- **R-28** Publicação exige nome, ao menos uma categoria ativa, tipo e área (se aplicável).
- **R-29** Prévia pública mostra exatamente o que será exibido, sem PII.
- **R-30** Edição relevante gera evento de auditoria e atualiza read model.
- **R-31** Verificação é selo separado: `unverified`, `contact_verified`, `identity_verified`; não promete antecedentes.
- **R-32** Profissional pausa e reativa o perfil.
- **R-33** Categoria desativada remove perfil da busca até ajuste.

### D. Portfólio

- **R-34** Profissional cria item com título, descrição, categoria e imagem opcional.
- **R-35** Item tem status `draft`, `published`, `removed`.
- **R-36** Profissional edita, reordena, despublica e remove item.
- **R-37** Portfólio público mostra apenas itens publicados.
- **R-38** Conteúdo não é apresentado como experiência verificada.
- **R-39** Conteúdo pode ser denunciado.

### E. Catálogo

- **R-40** Catálogo tem categorias/subcategorias com slug, nome, descrição, sinônimos e status.
- **R-41** Cada serviço declara tipo e categoria.
- **R-42** Categorias inativas não aparecem no onboarding nem na busca.
- **R-43** Seed inicial é versionado, idempotente e revisado por produto.
- **R-44** Operador cria/edita/desativa categoria; toda ação é auditada.
- **R-45** Busca no catálogo suporta sinônimo simples.

### F. Admin e auditoria

- **R-46** Operador lista perfis e vê status.
- **R-47** Operador suspende/reativa perfil com motivo e evidência.
- **R-48** Suspensão remove perfil da busca imediatamente e revoga sessões do perfil.
- **R-49** Toda ação administrativa grava ator, alvo, ação, motivo, timestamp e correlação em `audit_events` append-only.
- **R-50** Admin gerencia catálogo e feature flags; política de ranking fica em `fair-visibility`.
- **R-51** Acesso admin exige MFA e menor privilégio.

### G. Privacidade e LGPD

- **R-52** Consentimentos versionados e revogáveis.
- **R-53** Localização precisa cifrada, acesso restrito, retenção definida.
- **R-54** API pública nunca retorna e-mail, telefone, documento ou coordenada precisa.
- **R-55** Logs e eventos não contêm PII; redaction testada.
- **R-56** Exportação e exclusão implementadas.
- **R-57** Termos e política de privacidade exibidos e versionados.

### H. Acessibilidade e experiência

- **R-58** Cadastro, onboarding e perfil funcionam por teclado e leitor de tela.
- **R-59** Layout funciona em 360 px, sem rolagem horizontal.
- **R-60** Formulários têm rótulos persistentes, ajuda contextual e erro específico.
- **R-61** Estados de carregando, vazio, erro e sucesso explícitos.
- **R-62** `prefers-reduced-motion` respeitado.
- **R-63** Textos em PT-BR, linguagem simples, sem jargão.

### I. Observabilidade

- **R-64** Logs JSON com `correlation_id`, sem PII.
- **R-65** Métricas de cadastro, confirmação, publicação e erro.
- **R-66** Auditoria de ações administrativas e mudanças de status.
- **R-67** Health check e readiness do banco.
- **R-68** Feature flags com default seguro.

---

## 5. Regras de negócio e invariantes

- e-mail único normalizado;
- senha nunca em texto;
- token de confirmação/reset de uso único;
- perfil publicado exige categoria ativa e área/remoto;
- área de atendimento válida;
- localização precisa separada do read model público;
- suspensão remove da busca e revoga sessão;
- auditoria append-only;
- categoria inativa não pode ser nova associação;
- perfil incompleto não aparece na busca;
- contato nunca público antes de consentimento.

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
POST   /v1/auth/password-reset/request
POST   /v1/auth/password-reset/confirm

GET    /v1/me
PATCH  /v1/me
GET    /v1/me/consents
POST   /v1/me/consents/{kind}/revoke
GET    /v1/me/export
POST   /v1/me/delete

GET    /v1/categories
GET    /v1/categories/{slug}

GET    /v1/professionals/me
POST   /v1/professionals/me
PATCH  /v1/professionals/me
POST   /v1/professionals/me/publish
POST   /v1/professionals/me/pause
POST   /v1/professionals/me/resume
GET    /v1/professionals/{id}
GET    /v1/professionals/{id}/preview
GET    /v1/professionals/{id}/portfolio
POST   /v1/professionals/me/portfolio
PATCH  /v1/portfolio/{id}
DELETE /v1/portfolio/{id}

GET    /v1/admin/professionals
POST   /v1/admin/professionals/{id}/suspend
POST   /v1/admin/professionals/{id}/reactivate
GET    /v1/admin/audit-events
GET    /v1/healthz
GET    /v1/readyz
```

### 7.2 Entidades

- `accounts`, `sessions`, `email_tokens`, `consents`, `account_requests`
- `professional_profiles`, `profile_categories`, `portfolio_items`
- `categories`, `category_synonyms`
- `verifications` (shell)
- `audit_events`, `outbox_events`, `feature_flags`

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
- E-mail duplicado → resposta genérica para não enumerar contas.
- Localização precisa ausente → usar cidade/bairro; não inventar coordenada.
- Upload de imagem inválida → rejeitar com mensagem específica.
- Token expirado → pedir novo.
- Usuário excluído tenta login → resposta genérica.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `VALIDATION_ERROR` | campo inválido | "Confira os campos destacados." |
| `UNAUTHENTICATED` | sem sessão | "Entre para continuar." |
| `FORBIDDEN` | sem permissão | "Você não tem acesso a esta ação." |
| `PROFILE_INCOMPLETE` | publicação incompleta | "Complete nome, categoria e área para publicar." |
| `INVALID_SERVICE_AREA` | área/raio inválido | "Informe uma área de atendimento válida." |
| `EMAIL_ALREADY_USED` | e-mail existente (resposta genérica) | "Se o e-mail existir, enviaremos instruções." |
| `TOKEN_EXPIRED` | token inválido | "Este link expirou. Solicite outro." |
| `RATE_LIMITED` | limite | "Muitas tentativas. Tente novamente em alguns minutos." |
| `CONTACT_PROTECTED` | contato antes do aceite | "O contato é liberado após o aceite." |
| `INTERNAL_ERROR` | erro interno | "Algo deu errado. Tente novamente." |

---

## 10. Privacidade, segurança e fairness

- **PII:** e-mail, telefone, documento e localização precisa nunca em resposta pública.
- **Authz:** toda operação verifica ator e recurso; admin com MFA.
- **Rate limit:** login 5/min, cadastro 3/h, reset 3/h.
- **Fairness:** esta spec não ordena resultados; apenas garante elegibilidade.
- **Acessibilidade:** R-58 a R-63.

---

## 11. Observabilidade e métricas

- eventos: `account_signup_started`, `account_signup_completed`, `email_confirmed`, `profile_created`, `profile_published`, `profile_updated`.
- métricas: cadastro, confirmação, publicação, erro, tempo de onboarding.
- logs: `correlation_id`, ator pseudonimizado, sem PII.
- alertas: taxa de erro de cadastro, falha de e-mail, falha de migration.

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
- [ ] authz por recurso em perfil e admin;
- [ ] rate limit em auth;
- [ ] MFA e auditoria de admin;
- [ ] acessibilidade das jornadas de onboarding;
- [ ] migração e rollback;
- [ ] seed de catálogo idempotente;
- [ ] OpenAPI atualizado;
- [ ] SQLite/Postgres testados.

---

## 16. Perguntas para o humano

- [ ] Q-07: um perfil por conta no MVP?
- [ ] Q-08: área por raio é suficiente?
- [ ] Q-09: confirmação de contato por e-mail, telefone ou ambos?
- [ ] Q-10: retenção de documentos?
- [ ] Q-12: portfólio com imagem no MVP?
- [ ] Q-44: SQLite local/Postgres produção?
- [ ] Q-45: web server-rendered?
- [ ] Q-46: sessão por cookie?
- [ ] Q-47: MFA TOTP?
- [ ] Q-50: feature flags em banco?

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| R-01..R-12 | RF-01 | FND-03, FND-04 | AC-01..AC-08 |
| R-13..R-19 | RF-02, RF-19 | FND-08 | AC-18..AC-23 |
| R-20..R-33 | RF-03 | FND-06 | AC-09..AC-16 |
| R-34..R-39 | RF-04 | FND-07 | AC-24..AC-27 |
| R-40..R-45 | RF-05 | FND-05 | AC-17, AC-28..AC-30 |
| R-46..R-51 | RF-18 | FND-09 | AC-31..AC-33 |
| R-52..R-57 | RF-19 | FND-08 | AC-18..AC-23 |
| R-58..R-63 | RNF-10..12 | FND-11, FND-12 | AC-34..AC-36 |
| R-64..R-68 | RNF-13, RNF-14 | FND-10 | AC-37..AC-39 |
