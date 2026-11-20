# Critérios de Review — RAIRE

**Status:** vinculante para todo PR a partir da aprovação das specs  
**Versão:** 1.0  
**Relacionado:** `docs/process/harness.md`, `docs/process/pipeline.md`, `docs/process/qa-strategy.md`, `docs/process/security-privacy.md`

---

## 1. Como usar

Todo PR passa por gates na ordem abaixo. Um bloqueio em qualquer gate impede merge. O parecer final é consolidado em `specs/<feature>/review.md`.

Formato de saída:

```txt
Veredito: aprovado | aprovado com ressalvas | bloqueado
Gate: produto | arquitetura | código | segurança | privacidade | fairness | QA | release

Bloqueios:
- [arquivo:linha] descrição objetiva e correção esperada

Sugestões:
- ...

Riscos residuais:
- ...

O que o humano deve olhar:
- ...
```

Regras:

- quem implementou não aprova;
- bloqueio precisa citar requisito/spec;
- sugestão não bloqueia, mas deve ser registrada;
- "aprovado com ressalvas" exige lista de ressalvas e prazo;
- nenhum gate pode ser pulado por urgência.

---

## 2. Gate de produto

| Critério | Bloqueia? |
| --- | --- |
| Spec aprovada e vinculada ao PR | Sim |
| Ticket corresponde a um RF/RNF do PRD | Sim |
| Critérios de aceite cobrem caminho feliz, erros e bordas | Sim |
| Fora de escopo está explícito | Sim |
| Copy em PT-BR, clara e sem promessa enganosa | Sim |
| Estados de carregando, vazio, erro e sucesso existem | Sim |
| Jornada funciona em 360 px | Sim |
| Não introduz dark pattern, urgência falsa ou pay-to-win | Sim |
| Métrica de sucesso e instrumentação definidas | Sim |
| Perguntas abertas resolvidas ou aceitas por humano | Sim |

**Perguntas de review:**

- O que o usuário consegue fazer que não conseguia antes?
- Qual hipótese isso valida?
- O que acontece se o usuário errar?
- O que acontece se ranking/notificação cair?
- Isso aumenta confiança ou apenas engajamento?

---

## 3. Gate de arquitetura

| Critério | Bloqueia? |
| --- | --- |
| Contextos e dependências respeitam o plano | Sim |
| Nenhum contexto importa adapter de outro | Sim |
| Contratos versionados e compatíveis | Sim |
| Modelo de dados e invariantes documentados | Sim |
| Migração tem plano de rollback lógico | Sim |
| Eventos são imutáveis, versionados e sem PII | Sim |
| Decisão relevante tem ADR | Sim |
| Complexidade nova tem justificativa e gatilho | Sim |
| Ranking/fairness permanece isolado e testável | Sim |
| Fallback e degradação graciosa definidos | Sim |

**Perguntas de review:**

- Isso precisa existir agora?
- Qual a alternativa mais simples?
- O acoplamento criado é reversível?
- O que quebra se essa parte for extraída depois?
- Há dependência circular ou shared kernel excessivo?

---

## 4. Gate de código

| Critério | Bloqueia? |
| --- | --- |
| Testes escritos antes e cobrindo aceite | Sim |
| Testes verdes, sem `.skip`, `.only` ou TODO órfão | Sim |
| Cobertura não caiu abaixo do mínimo acordado | Sim |
| Lint, vet, typecheck e build verdes | Sim |
| Sem casting inseguro ou erro ignorado | Sim |
| Erros de domínio tipados e mensagens seguras | Sim |
| Entrada validada e normalizada | Sim |
| Autorização por recurso (ownership) | Sim |
| Transações e idempotência onde necessário | Sim |
| Sem log de PII ou debug | Sim |
| Nomes claros, funções pequenas, sem duplicação relevante | Não (sugestão) |
| Comentário explica "por quê", não "o quê" | Não (sugestão) |

**Proibido:**

- regra de negócio em handler HTTP;
- SQL concatenado com input;
- segredo hardcoded;
- mock de banco em teste de integração;
- teste que passa por acidente;
- dependência nova sem justificativa.

---

## 5. Gate de segurança

Checklist mínimo (OWASP ASVS nível 2, adaptado):

- autenticação e sessão seguras;
- senha com hash forte e política adequada;
- autorização por papel e por recurso;
- rate limit em login, cadastro, busca, pedidos, interesses e denúncias;
- validação de entrada e output encoding;
- proteção CSRF quando aplicável;
- headers de segurança e CSP;
- segredos fora do código;
- dependências sem vulnerabilidade crítica/alta;
- logs sem PII/segredos;
- backup e restore testados;
- admin com MFA e auditoria;
- erro não vaza stack ou detalhe interno.

**Bloqueia:** vulnerabilidade crítica/alta alcançável; acesso indevido a dado de outro usuário; injeção; bypass de autenticação; exposição de segredo.

---

## 6. Gate de privacidade e LGPD

| Critério | Bloqueia? |
| --- | --- |
| Finalidade e base legal documentadas | Sim |
| Coleta mínima necessária | Sim |
| Consentimento granular quando exigido | Sim |
| Localização precisa protegida e com retenção | Sim |
| Contato mascarado até consentimento | Sim |
| Direitos do titular implementados | Sim |
| PII fora de logs/eventos/analytics | Sim |
| Retenção e anonimização definidas | Sim |
| DPIA quando houver risco | Sim |
| Incidente tem plano de resposta | Sim |

**Perguntas de review:**

- Precisamos desse dado para a jornada?
- O usuário entende o que será público?
- Onde esse dado aparece em logs, eventos e backups?
- Quem acessa e por quanto tempo?
- Como corrigir/excluir?

---

## 7. Gate de fairness e ética

Aplicável a busca, ranking, recomendação, reputação, moderação e experimentos.

| Critério | Bloqueia? |
| --- | --- |
| Nenhum atributo sensível usado direta ou indiretamente | Sim |
| Coortes são observáveis e agregadas | Sim |
| Mudança de ranking tem avaliação offline | Sim |
| Guardrails de relevância e fairness definidos | Sim |
| Política versionada, auditável e reversível | Sim |
| Explicação deriva de sinais reais | Sim |
| Não há posição paga ou vantagem oculta | Sim |
| Profissionais novos não são punidos sem justificativa | Sim |
| Revisão humana antes de ativar | Sim |
| Plano de monitoramento pós-ativação | Sim |

**Sinais de alerta:**

- melhora de clique com queda de conexão;
- aumento de concentração no topo;
- queda de diversidade de bairro/categoria;
- reclamações de profissionais invisíveis;
- explicação genérica ou enganosa;
- uso de proxies de atributos sensíveis (CEP, nome, escola etc.) sem análise.

---

## 8. Gate de QA

| Critério | Bloqueia? |
| --- | --- |
| Matriz de aceite → teste completa | Sim |
| Testes unitários e de integração verdes | Sim |
| E2E das jornadas críticas verdes | Sim |
| Acessibilidade automatizada + manual | Sim |
| Regressão das áreas afetadas | Sim |
| Performance dentro do orçamento | Sim |
| Teste em 360 px e teclado | Sim |
| Evidência anexada | Sim |
| Falha conhecida registrada com issue | Sim |

Detalhes em `docs/process/qa-strategy.md`.

---

## 9. Gate de release

| Critério | Bloqueia? |
| --- | --- |
| Pipeline completo verde | Sim |
| Smoke em staging | Sim |
| Migração aplicada e rollback testado | Sim |
| Feature flag com default seguro | Sim |
| Dashboard e alertas | Sim |
| Runbook e contato de incidente | Sim |
| Comunicação de mudança | Sim |
| Aprovação humana para piloto/produção | Sim |

---

## 10. Checklist de PR

```md
## Spec e escopo
- [ ] Spec aprovada: specs/<feature>/spec.md
- [ ] Ticket: <ID>
- [ ] RF/RNF: <IDs>
- [ ] Fora de escopo respeitado

## Testes
- [ ] Testes de aceite escritos antes
- [ ] Unitários
- [ ] Integração
- [ ] E2E
- [ ] Acessibilidade
- [ ] Performance
- [ ] Fairness (se aplicável)

## Segurança e privacidade
- [ ] Authz por recurso
- [ ] Rate limit
- [ ] Sem PII em log/evento
- [ ] Retenção/consentimento
- [ ] Dependências verificadas

## Operação
- [ ] Migração + rollback
- [ ] Feature flag
- [ ] Métricas/alertas
- [ ] Runbook

## Evidência
- [ ] Relatório de testes
- [ ] Screenshots/vídeo
- [ ] Simulação de fairness (se aplicável)
- [ ] Riscos residuais
```
