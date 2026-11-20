# service-request

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** RF-08, RF-09, RF-10, RF-11, RF-12, RF-16  
**RNFs:** RNF-02, RNF-05, RNF-06, RNF-13  
**Módulos:** demand, directory (read), catalog (read), identity, moderation (hooks), platform, web  
**Depende de:** `foundation-platform`; ADR 0004  
**Dono:** Anderson  
**Revisores:** Victor (contratos), Felipe (abuso/PII)  
**Aprovação humana:** pendente

---

## 1. Contexto

O cliente precisa publicar uma necessidade e receber interesse de profissionais elegíveis, sem expor contato ou endereço. O profissional precisa ver oportunidades compatíveis e decidir se quer se conectar. Este é o fluxo do "mural de guilda" descrito no PRD.

**Problema do usuário:** "não quero procurar um por um; quero descrever o que preciso e receber opções confiáveis."

---

## 2. Objetivo e não objetivo

**Objetivo:** permitir publicar, editar, fechar e expirar pedidos; listar oportunidades elegíveis; enviar/aceitar/recusar interesses; criar conexão consentida e liberar contato.

**Não objetivo:**

- pagamento e escrow;
- agenda/calendário;
- chat em tempo real;
- leilão de preço;
- contrato digital;
- garantia/seguro;
- matching automático por IA;
- múltiplas cidades.

---

## 3. Princípios e restrições

- contato só é liberado após aceite mútuo;
- localização pública é aproximada; endereço exato nunca aparece;
- interesses são limitados para evitar spam e assédio;
- pedido expira e pode ser fechado pelo cliente;
- notificações respeitam preferência e consentimento;
- toda transição de estado é auditável;
- nenhuma PII em notificação/evento;
- acessibilidade e linguagem simples.

---

## 4. Requisitos

### A. Criação e ciclo do pedido

- **S-01** Cliente autenticado e com e-mail confirmado cria pedido com categoria, descrição, localização aproximada, urgência e preferência de contato.
- **S-02** Faixa de orçamento é opcional; quando informada, é um intervalo.
- **S-03** Pedido tem status `draft`, `open`, `receiving_interests`, `matched`, `closed`, `expired`, `cancelled`.
- **S-04** Cliente edita enquanto `open`/`receiving_interests`, com histórico.
- **S-05** Cliente fecha/cancela; pedido fechado não recebe novos interesses.
- **S-06** Pedido expira automaticamente após janela configurável (default 7 dias), com aviso prévio.
- **S-07** Limite de pedidos ativos por cliente (default 5).
- **S-08** Descrição tem limite, validação e não aceita PII explícita (telefone/e-mail/endereço).
- **S-09** Pedido não expõe endereço exato, coordenada precisa, e-mail ou telefone.
- **S-10** Conteúdo sinalizado entra em moderação conforme risco.

### B. Feed e elegibilidade do profissional

- **S-11** Profissional vê pedidos compatíveis com categoria, tipo e área.
- **S-12** Feed é paginado por cursor e ordenado por proximidade, urgência, recência e compatibilidade.
- **S-13** Profissional filtra por categoria, urgência e distância.
- **S-14** Feed não mostra identidade completa do cliente antes do aceite.
- **S-15** Profissional com perfil incompleto ou suspenso não vê pedidos.

### C. Interesse

- **S-16** Profissional envia interesse com mensagem curta e declara disponibilidade.
- **S-17** Um interesse por profissional/pedido; retirada permitida antes do aceite.
- **S-18** Limite de interesses por profissional/dia (20) e por pedido (10).
- **S-19** Cliente vê lista de interessados com perfil resumido, sinais e mensagem.
- **S-20** Cliente aceita um ou mais interesses (default até 3).
- **S-21** Interesse não aceito expira quando o pedido fecha/expira.
- **S-22** Profissional pode recusar/retirar; cliente é notificado.

### D. Conexão consentida

- **S-23** Aceite cria `connection` com estado `accepted` e registra consentimento dos dois lados.
- **S-24** Contato é liberado apenas após aceite, conforme preferência de cada lado.
- **S-25** Partes marcam conexão como `contacted`, `completed`, `cancelled` ou `disputed`.
- **S-26** Conclusão exige confirmação de ao menos um lado, com contestação.
- **S-27** Bloqueio e denúncia disponíveis em qualquer estado.
- **S-28** Toda transição grava evento de auditoria com ator e timestamp.

### E. Notificações

- **S-29** Notificações transacionais: pedido criado, novo interesse, interesse aceito, pedido expirando, conexão concluída.
- **S-30** Preferências de notificação são respeitadas; marketing é opt-in separado.
- **S-31** Notificação não contém PII desnecessária nem endereço exato.
- **S-32** Falha de notificação não bloqueia a transição; entra em retry/outbox.

### F. Anti-spam e abuso

- **S-33** Rate limit por cliente/profissional/IP em criar pedido, interessar, aceitar e denunciar.
- **S-34** Detecção básica de duplicidade e conteúdo repetitivo.
- **S-35** Sinais de abuso vão para moderação; não bloqueiam automaticamente sem revisão, salvo risco evidente.
- **S-36** Usuário bloqueado não cria pedido nem envia interesse.

### G. Privacidade e acessibilidade

- **S-37** Localização precisa é usada só para elegibilidade e não é retornada ao outro lado.
- **S-38** Consentimento de contato é registrado com timestamp e versão.
- **S-39** Fluxo funciona em 360 px, teclado e leitor de tela.
- **S-40** Estados de carregando, vazio, erro e sucesso explícitos.
- **S-41** Textos em PT-BR, sem jargão e sem urgência falsa.
- **S-42** Dados de pedido seguem retenção definida; exclusão anonimiza conteúdo público.

---

## 5. Regras de negócio e invariantes

- pedido pertence a um cliente;
- interesse único por profissional/pedido;
- contato só após aceite;
- pedido fechado/expirado não recebe interesse;
- conexão só existe a partir de interesse aceito/solicitação aceita;
- consentimento registrado antes de liberar contato;
- endereço exato nunca no read model público;
- transições válidas apenas;
- cliente não vê identidade completa do profissional antes do interesse? (ver Q-31)
- profissional não vê contato do cliente antes do aceite.

---

## 6. Estados e transições

**Request:**

```txt
draft -> open -> receiving_interests -> matched -> closed
open/receiving_interests -> expired/cancelled
```

**Interest:**

```txt
sent -> viewed -> accepted/declined
sent/viewed -> withdrawn/expired
```

**Connection:**

```txt
pending -> accepted -> contacted -> completed
accepted/contacted -> cancelled/disputed
```

---

## 7. Contratos e dados

### 7.1 Endpoints

```txt
POST   /v1/requests
GET    /v1/requests/{id}
PATCH  /v1/requests/{id}
POST   /v1/requests/{id}/close
POST   /v1/requests/{id}/cancel
GET    /v1/requests/me
GET    /v1/requests/feed
POST   /v1/requests/{id}/interests
GET    /v1/requests/{id}/interests
POST   /v1/interests/{id}/accept
POST   /v1/interests/{id}/decline
POST   /v1/interests/{id}/withdraw
GET    /v1/connections/me
GET    /v1/connections/{id}
POST   /v1/connections/{id}/contacted
POST   /v1/connections/{id}/complete
POST   /v1/connections/{id}/cancel
POST   /v1/connections/{id}/dispute
POST   /v1/reports
GET    /v1/notifications
POST   /v1/notifications/{id}/read
```

### 7.2 Entidades

- `service_requests`
- `request_interests`
- `connections`
- `connection_outcomes` (ver trust-reputation)
- `notifications`
- `outbox_events`
- `audit_events`

### 7.3 Invariantes

- pedido pertence a um cliente;
- interesse único por profissional/pedido;
- contato só após aceite;
- pedido fechado não recebe interesse;
- conexão só existe a partir de interesse aceito;
- consentimento registrado antes de liberar contato;
- endereço exato nunca no read model público;
- transições válidas apenas.

---

## 8. Casos de borda

- cliente tenta aceitar mais interesses que o limite → bloquear com mensagem clara;
- profissional tenta interessar em pedido expirado → `REQUEST_EXPIRED`;
- cliente fecha pedido com interesses pendentes → notificar e expirar;
- profissional retira interesse após aceite → tratar como cancelamento;
- e-mail/notificação falha → estado permanece, retry;
- pedido com PII na descrição → bloquear publicação e orientar;
- denúncia durante conexão → não apaga histórico; moderação decide;
- usuário bloqueado → não cria pedido/interesse;
- pedido duplicado → sinalizar, não bloquear automaticamente.

---

## 9. Erros e mensagens

| Código | Quando | Mensagem |
| --- | --- | --- |
| `PII_NOT_ALLOWED` | descrição com PII | "Remova telefone, e-mail ou endereço da descrição." |
| `REQUEST_EXPIRED` | pedido expirado | "Este pedido expirou." |
| `INTEREST_LIMIT_REACHED` | limite | "Você atingiu o limite de interesses." |
| `ALREADY_INTERESTED` | duplicado | "Você já demonstrou interesse." |
| `CONTACT_PROTECTED` | contato antes do aceite | "O contato é liberado após o aceite." |
| `INVALID_STATE_TRANSITION` | transição inválida | "Esta ação não está disponível agora." |
| `FORBIDDEN` | sem permissão | "Você não tem acesso a este pedido." |
| `RATE_LIMITED` | limite | "Muitas ações. Aguarde alguns minutos." |

---

## 10. Privacidade, segurança e fairness

- **PII:** contato/endereço nunca antes do aceite.
- **Authz:** dono do pedido, profissional elegível, partes da conexão.
- **Rate limit:** pedido 5/dia, interesse 20/dia, denúncia 10/dia.
- **Fairness:** feed ordena por relevância; sem posição paga.
- **Acessibilidade:** S-39..S-41.

---

## 11. Observabilidade e métricas

- eventos: `request_created`, `request_closed`, `request_expired`, `interest_sent`, `interest_accepted`, `interest_declined`, `connection_created`, `connection_contacted`, `connection_completed`.
- métricas: pedidos com 3+ interesses, tempo até primeiro interesse, aceite, desfecho.
- alertas: pico de interesses, falha de notificação, pedido sem interesse.

---

## 12. Fora de escopo

- pagamento/escrow;
- agenda/calendário;
- chat;
- leilão;
- contrato;
- garantia/seguro;
- matching por IA;
- múltiplas cidades.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| spam/assédio | experiência ruim | limites, bloqueio, moderação |
| vazamento de contato | risco a usuários | consentimento e mascaramento |
| baixa liquidez | poucos interesses | curadoria, notificações, piloto denso |
| pedido sem resposta | frustração | expiry, lembretes, sugestão de busca |
| aceite sem intenção real | conexão falsa | confirmação, sinais, denúncia |
| PII na descrição | vazamento | validação e orientação |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- criação/edição/fechamento/expiração;
- feed elegível;
- interesse e limites;
- aceite e liberação de contato;
- recusa/retirada;
- bloqueio/denúncia;
- notificações;
- PII;
- acessibilidade.

---

## 15. Gates de review específicos

- [ ] contato mascarado até aceite;
- [ ] authz por dono do pedido/interesse;
- [ ] rate limit e anti-spam;
- [ ] PII fora de feed/notificação/evento;
- [ ] transições de estado testadas;
- [ ] acessibilidade das jornadas;
- [ ] notificação e retry;
- [ ] auditoria.

---

## 16. Perguntas para o humano

- [ ] Q-25: cliente pode aceitar mais de um profissional?
- [ ] Q-26: limites de interesses.
- [ ] Q-27: expiração do pedido.
- [ ] Q-28: contato por e-mail, telefone ou ambos?
- [ ] Q-29: cliente publica sem localização exata?
- [ ] Q-30: como confirmar serviço?
- [ ] Q-31: moderação prévia ou pós-publicação?

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| S-01..S-10 | RF-08 | SRV-01, SRV-02 | AC-01..AC-10 |
| S-11..S-15 | RF-09 | SRV-03 | AC-11..AC-15 |
| S-16..S-22 | RF-10 | SRV-04 | AC-16..AC-22 |
| S-23..S-28 | RF-11, RF-12 | SRV-05, SRV-06 | AC-23..AC-29 |
| S-29..S-32 | RF-16 | SRV-07 | AC-30..AC-33 |
| S-33..S-36 | RF-10, RF-17 | SRV-04, SRV-08 | AC-34..AC-37 |
| S-37..S-42 | RF-12, RF-19 | SRV-06, SRV-09 | AC-38..AC-40 |
