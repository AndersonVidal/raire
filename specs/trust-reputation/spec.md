# trust-reputation

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** RF-12, RF-13, RF-14, RF-15, RF-17, RF-18  
**RNFs:** RNF-05, RNF-08, RNF-09, RNF-14  
**Módulos:** reputation, moderation, demand (read), directory (read), identity, web, platform  
**Depende de:** `foundation-platform`, `service-request`; ADR 0004  
**Dono:** Felipe  
**Revisores:** Victor (dados), Anderson (copy/jornada)  
**Aprovação humana:** pendente

---

## 1. Contexto

O cliente precisa de sinais honestos de confiança para decidir. O profissional precisa construir reputação sem ser refém de uma nota isolada ou de avaliações falsas. A plataforma precisa moderar abuso sem apagar críticas legítimas.

Esta spec cobre desfechos, avaliações pós-conexão, resposta, verificação de contato, sinais de confiança, denúncias, moderação e recurso.

**Problema do usuário:** "não sei se posso confiar nessa pessoa; a nota sozinha não me diz o suficiente."

---

## 2. Objetivo e não objetivo

**Objetivo:** permitir avaliação bilateral após conexão elegível, reputação contextual, denúncia e moderação auditável.

**Não objetivo:**

- background check;
- garantia/seguro;
- mediação de pagamento;
- arbitragem jurídica;
- reputação entre plataformas;
- score público único;
- verificação automática por IA no MVP.

---

## 3. Princípios e restrições

- avaliação só após conexão confirmada;
- uma avaliação por lado por conexão;
- reputação é contextual, não média global;
- verificação é honesta e não promete antecedentes;
- crítica legítima não é removida;
- toda ação de moderação é auditável e recorrível;
- nenhuma PII de avaliador ou denunciante exposta;
- profissional novo não é punido por falta de histórico.

---

## 4. Requisitos

### A. Elegibilidade e avaliação

- **T-01** Avaliação só é permitida quando a conexão está `completed` ou quando a regra de confirmação for satisfeita.
- **T-02** Cada lado avalia uma vez por conexão; não há autoavaliação.
- **T-03** Avaliação contém nota 1–5, tags contextuais, comentário opcional e categoria da conexão.
- **T-04** Comentário tem limite, validação de PII e não aceita links maliciosos.
- **T-05** Janela de edição curta (default 24 h); depois, imutável, salvo moderação.
- **T-06** Avaliação pode ser removida apenas por violação de política, com motivo e auditoria.
- **T-07** Avaliado pode responder uma vez; resposta segue as mesmas regras de conteúdo.
- **T-08** Avaliação e resposta são públicas após publicação, sem e-mail/telefone/documento.

### B. Reputação contextual

- **T-09** Reputação é calculada por categoria e contexto, com média bayesiana e distribuição.
- **T-10** Exibir número de avaliações, distribuição e recência; não apenas média.
- **T-11** Profissional novo recebe prior neutro, não zero.
- **T-12** Sinais complementares: taxa de resposta, conclusão, cancelamento e denúncias procedentes.
- **T-13** Sinais não podem dominar relevância nem criar punição permanente por evento isolado.
- **T-14** Denúncia procedente tem peso limitado e decai com o tempo.
- **T-15** Reputação é recalculável e auditável a partir dos eventos.

### C. Verificação

- **T-16** Verificação de contato: e-mail/telefone validado.
- **T-17** Verificação de identidade: documento analisado manualmente, com retenção mínima (fora do primeiro ciclo).
- **T-18** Status: `unverified`, `contact_verified`, `identity_verified`, `rejected`, `expired`.
- **T-19** Selo exibe o que foi verificado e a data; nunca promete antecedentes criminais.
- **T-20** Documento é cifrado, acesso restrito e excluído após decisão/prazo.
- **T-21** Rejeição tem motivo e possibilidade de reenvio.
- **T-22** Verificação pode ser revogada por fraude, com auditoria.

### D. Denúncias e moderação

- **T-23** Usuário denuncia perfil, avaliação, pedido, interesse, portfólio ou conexão.
- **T-24** Denúncia tem motivo, descrição e prioridade por risco.
- **T-25** Denunciante não é exposto ao denunciado.
- **T-26** Fila de moderação com SLA inicial de 48 h e prioridade para risco físico.
- **T-27** Ações: nenhuma, advertência, remoção de conteúdo, suspensão temporária, bloqueio.
- **T-28** Toda ação grava ator, alvo, motivo, evidência e timestamp.
- **T-29** Usuário afetado pode recorrer em prazo definido.
- **T-30** Recurso é analisado por pessoa diferente da que decidiu, quando possível.
- **T-31** Ação de moderação é comunicada ao afetado sem expor dados de terceiros.

### E. Anti-manipulação

- **T-32** Detectar autoavaliação, avaliação cruzada combinada e repetição suspeita.
- **T-33** Limitar avaliações por IP/dispositivo/conexão sem usar dados sensíveis.
- **T-34** Sinais de manipulação vão para revisão; não removem automaticamente sem critério.
- **T-35** Avaliações de contas recém-criadas podem ser sinalizadas.
- **T-36** Denúncias falsas repetidas podem ser sinalizadas.

### F. Privacidade e acessibilidade

- **T-37** Avaliação pública não contém PII do autor.
- **T-38** Denúncia e evidências são restritas à moderação.
- **T-39** Fluxo de avaliação, resposta e denúncia funciona em 360 px, teclado e leitor de tela.
- **T-40** Estados de carregando, vazio, erro e sucesso explícitos.
- **T-41** Textos em PT-BR, sem linguagem que incentive retaliação.
- **T-42** Retenção de avaliações, denúncias e evidências segue política.

---

## 5. Regras de negócio e invariantes

- uma avaliação por autor/conexão;
- autor ≠ avaliado;
- avaliação exige conexão elegível;
- resposta única;
- remoção só com motivo e auditoria;
- documento restrito e com retenção;
- denúncia não expõe denunciante;
- ação de moderação tem ator e timestamp;
- crítica legítima preservada;
- novo profissional com prior neutro.

---

## 6. Estados e transições

**Outcome:**

```txt
contacted -> negotiated -> completed
contacted/negotiated -> not_completed/disputed
```

**Review:**

```txt
eligible -> submitted -> published
submitted/published -> under_review -> removed/restored
```

**Report:**

```txt
open -> triaged -> investigating -> resolved/dismissed
resolved/dismissed -> appealed -> resolved/dismissed
```

**Verification:**

```txt
not_started -> pending -> verified/rejected -> expired/revoked
```

---

## 7. Contratos e dados

### 7.1 Endpoints

```txt
POST /v1/connections/{id}/outcomes
POST /v1/connections/{id}/reviews
GET  /v1/professionals/{id}/reviews
POST /v1/reviews/{id}/response
POST /v1/reviews/{id}/report
POST /v1/reports

GET  /v1/me/verification
POST /v1/me/verification/contact
POST /v1/me/verification/identity
GET  /v1/admin/reports
POST /v1/admin/reports/{id}/resolve
POST /v1/admin/moderation-actions
POST /v1/admin/moderation-actions/{id}/appeal
GET  /v1/admin/audit-events
```

### 7.2 Entidades

- `connection_outcomes`
- `reviews`, `review_responses`
- `reputation_snapshots`
- `verifications`
- `reports`, `moderation_actions`, `appeals`
- `audit_events`

### 7.3 Invariantes

- uma avaliação por autor/conexão;
- autor ≠ avaliado;
- avaliação exige conexão elegível;
- resposta única;
- remoção só com motivo e auditoria;
- documento restrito e com retenção;
- denúncia não expõe denunciante;
- ação de moderação tem ator e timestamp.

---

## 8. Casos de borda

- conexão contestada → avaliação fica pendente até resolução;
- avaliação com PII → bloqueada com orientação;
- avaliação ofensiva → moderação, não remoção automática;
- profissional sem avaliações → exibir "Novo" e não nota zero;
- denúncia durante avaliação → avaliação pode ficar `under_review`;
- recurso após prazo → não aceito, salvo erro da plataforma;
- verificação expirada → selo removido, sem apagar histórico;
- duas avaliações da mesma conexão → permitir um por lado;
- avaliação de conexão cancelada → bloquear.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `REVIEW_NOT_ELIGIBLE` | conexão não elegível | "A avaliação fica disponível após a conexão ser concluída." |
| `REVIEW_ALREADY_EXISTS` | duplicada | "Você já avaliou esta conexão." |
| `SELF_REVIEW` | autoavaliação | "Não é possível avaliar a si mesmo." |
| `PII_NOT_ALLOWED` | comentário com PII | "Remova dados pessoais do comentário." |
| `REVIEW_WINDOW_CLOSED` | edição expirada | "O prazo de edição terminou." |
| `REPORT_LIMIT_REACHED` | limite | "Você atingiu o limite de denúncias." |
| `FORBIDDEN` | sem permissão | "Você não tem acesso a esta ação." |

---

## 10. Privacidade, segurança e fairness

- **PII:** autor/denunciante protegidos; documento cifrado.
- **Authz:** partes da conexão, operador com MFA.
- **Rate limit:** denúncia 10/dia, avaliação 1/conexão.
- **Fairness:** reputação contextual; novo com prior neutro; sem atributo sensível.
- **Acessibilidade:** T-39..T-41.

---

## 11. Observabilidade e métricas

- eventos: `outcome_recorded`, `review_submitted`, `review_published`, `report_created`, `moderation_action_taken`.
- métricas: cobertura de avaliação, distribuição, denúncias, tempo de triagem/resolução, recursos.
- alertas: pico de denúncias, tempo de fila, manipulação.

---

## 12. Fora de escopo

- background check;
- garantia/seguro;
- mediação de pagamento;
- arbitragem jurídica;
- reputação entre plataformas;
- score público único;
- verificação automática por IA.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| avaliações falsas | confiança | elegibilidade, detecção, moderação |
| retaliação | medo de avaliar | anonimato do autor, prazo, política |
| remoção de crítica legítima | perda de confiança | critério público e auditoria |
| exposição de documento | risco grave | cifra, retenção, acesso restrito |
| moderação lenta | dano | prioridade por risco e SLA |
| viés de reputação | injustiça | contexto, prior neutro, decaimento |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- elegibilidade;
- uma avaliação por lado;
- resposta;
- reputação contextual;
- verificação;
- denúncia e moderação;
- recurso;
- anti-manipulação;
- PII;
- acessibilidade.

---

## 15. Gates de review específicos

- [ ] avaliação só após conexão elegível;
- [ ] PII de autor/denunciante protegida;
- [ ] documento cifrado e com retenção;
- [ ] moderação auditável e recorrível;
- [ ] crítica legítima preservada;
- [ ] anti-manipulação testado;
- [ ] acessibilidade;
- [ ] fairness de reputação.

---

## 16. Perguntas para o humano

- [ ] Q-06: verificação de identidade obrigatória?
- [ ] Q-09: confirmação de contato.
- [ ] Q-10: retenção de documentos.
- [ ] Q-30: como confirmar serviço?
- [ ] Q-32: confirmação mútua?
- [ ] Q-33: janela de edição.
- [ ] Q-34: verificação manual/terceirizada.
- [ ] Q-35: denúncia anônima?
- [ ] Q-36: SLA de risco físico.

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| T-01..T-08 | RF-14 | TRS-02, TRS-03, TRS-04 | AC-01..AC-09 |
| T-09..T-15 | RF-14 | TRS-05 | AC-10..AC-15 |
| T-16..T-22 | RF-15 | TRS-09 | AC-16..AC-20 |
| T-23..T-31 | RF-17, RF-18 | TRS-06, TRS-07 | AC-21..AC-29 |
| T-32..T-36 | RF-14 | TRS-08 | AC-30..AC-33 |
| T-37..T-42 | RF-17, RNF-09 | TRS-06, TRS-10 | AC-34..AC-37 |
