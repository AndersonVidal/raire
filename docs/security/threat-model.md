# Modelo de Ameaças — RAIRE

**Status:** proposta para revisão  
**Versão:** 1.0  
**Dono:** Felipe  
**Relacionado:** `docs/process/security-privacy.md`, `docs/security/privacy-and-lgpd.md`

---

## 1. Ativos

| Ativo | Sensibilidade | Impacto se comprometido |
| --- | --- | --- |
| credenciais e sessões | crítica | account takeover |
| e-mail e telefone | alta | assédio, spam, fraude |
| localização precisa | crítica | risco físico |
| documentos de verificação | crítica | fraude, dano legal |
| denúncias e evidências | crítica | exposição do denunciante |
| avaliações e reputação | alta | manipulação, dano social |
| ranking e políticas | alta | injustiça, gaming |
| eventos analíticos | média | reidentificação |
| logs | média | vazamento de PII |
| catálogo e perfis | pública | desinformação |

---

## 2. Atores de ameaça

- fraudador / perfil falso;
- spammer / scraper;
- atacante externo;
- usuário malicioso (assédio, retaliação);
- concorrente (avaliação combinada, gaming);
- operador abusivo;
- insider;
- dependência comprometida.

---

## 3. Ameaças e controles

| ID | Ameaça | Ator | Impacto | Controles | Bloqueia release? |
| --- | --- | --- | --- | --- | --- |
| T-01 | Perfil falso / catfishing | fraudador | alto | verificação, sinais, denúncia, moderação | sim se generalizado |
| T-02 | Scraping de contatos/perfis | spammer | alto | rate limit, mascaramento, detecção de enumeração | sim |
| T-03 | Assédio via contato | usuário malicioso | crítico | consentimento, bloqueio, denúncia, canal seguro | sim |
| T-04 | Vazamento de localização precisa | stalker | crítico | localização aproximada, cifra, acesso restrito | sim |
| T-05 | Manipulação de avaliações | concorrente | alto | elegibilidade, detecção, moderação | sim se sistemático |
| T-06 | Spam de pedidos/interesses | spammer | médio | limites, validação, triagem | não (se contido) |
| T-07 | Account takeover | atacante | crítico | senha forte, rate limit, sessão revogável, MFA | sim |
| T-08 | Abuso de admin | operador | crítico | MFA, menor privilégio, auditoria, recurso | sim |
| T-09 | Ranking gaming | profissional | alto | caps, detecção de anomalia, revisão | sim se injusto |
| T-10 | Discriminação algorítmica | sistema | crítico | proibir atributo sensível, coortes, auditoria | sim |
| T-11 | PII em log/evento | erro interno | crítico | redaction, testes, revisão | sim |
| T-12 | Dependência vulnerável | atacante | alto | scan, atualização, SBOM | sim se crítica |
| T-13 | Injeção / XSS / CSRF | atacante | alto | validação, encoding, CSP, CSRF | sim |
| T-14 | Enumeração de contas | atacante | médio | resposta genérica, rate limit | não |
| T-15 | Upload malicioso | atacante | alto | validação, storage isolado, antivírus | sim |
| T-16 | Abuso de denúncia | usuário | médio | limite, triagem, histórico | não |
| T-17 | Vazamento de denunciante | insider/log | crítico | acesso restrito, cifra, auditoria | sim |
| T-18 | Reidentificação em analytics | insider | alto | pseudonimização, agregação, k-anonimato | sim |
| T-19 | Falha de backup/restore | operação | crítico | backup cifrado, teste de restore | sim |
| T-20 | Dependência de fornecedor | operação | médio | isolamento, portabilidade | não |

---

## 4. Superfícies de ataque

1. cadastro e login;
2. confirmação de e-mail e reset;
3. busca e perfil público;
4. pedidos e interesses;
5. conexão e liberação de contato;
6. avaliações e respostas;
7. denúncias e moderação;
8. admin e feature flags;
9. upload de portfólio;
10. eventos e logs;
11. migrations e deploy.

---

## 5. Cenários prioritários

### 5.1 Acesso indevido a recurso de outro usuário

- **Cenário:** cliente A tenta ler pedido/conexão de B.
- **Controle:** authz por recurso; teste cruzado; 404 em vez de 403 quando revelar existência.
- **Bloqueia:** sim.

### 5.2 Contato exposto antes do aceite

- **Cenário:** API ou HTML retorna e-mail/telefone antes do aceite.
- **Controle:** campo derivado `contact.visible`; teste de contrato; mascaramento.
- **Bloqueia:** sim.

### 5.3 Localização precisa em resposta pública

- **Cenário:** perfil público retorna lat/lng exata.
- **Controle:** read model público sem coordenada; faixa de distância; teste.
- **Bloqueia:** sim.

### 5.4 Autoavaliação / avaliação combinada

- **Cenário:** profissional avalia a si mesmo ou combina com cliente.
- **Controle:** autor ≠ avaliado; uma por lado; detecção de padrão; moderação.
- **Bloqueia:** sim se sistemático.

### 5.5 Admin sem MFA

- **Cenário:** operador acessa admin só com senha.
- **Controle:** MFA TOTP obrigatório; teste.
- **Bloqueia:** sim.

### 5.6 Ranking com proxy sensível

- **Cenário:** CEP/nome/escola usados como proxy.
- **Controle:** revisão de features; proibição; auditoria.
- **Bloqueia:** sim.

---

## 6. Controles por camada

| Camada | Controles |
| --- | --- |
| Borda | TLS, headers, CSP, rate limit, WAF quando disponível |
| Auth | Argon2id, sessão, CSRF, MFA, reset seguro |
| App | validação, authz, erros tipados, logging sem PII |
| Dados | PII separada, cifra, backup cifrado, retenção |
| Operação | auditoria, menor privilégio, runbook, alertas |
| Produto | consentimento, mascaramento, bloqueio, denúncia |
| Fairness | sem atributo sensível, política versionada, revisão |

---

## 7. Testes de segurança obrigatórios

- autorização cruzada (A não acessa B);
- CSRF em formulário/API;
- rate limit em login/cadastro/interesse/denúncia;
- PII fora de resposta pública, log e evento;
- contato antes/depois do aceite;
- localização precisa ausente;
- autoavaliação bloqueada;
- admin sem MFA bloqueado;
- SQL injection e XSS;
- dependência crítica/alta;
- restore de backup.

---

## 8. Riscos residuais

- moderação manual limitada no piloto;
- SQLite local não replica todas as garantias do Postgres;
- e-mail transacional pode falhar;
- scraping sofisticado pode contornar limites;
- denúncia anônima não é permitida no MVP;
- verificação de identidade fica fora do primeiro ciclo.

---

## 9. Revisão

- Revisão trimestral do modelo.
- Atualização após incidente.
- Atualização antes de nova superfície (upload, pagamento, app).
- Aprovação de segurança registrada.
