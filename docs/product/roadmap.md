# Roadmap de Entrega — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0, seção 13  
**Uso:** planejamento de fases, marcos, donos e critérios de saída.

---

## 1. Princípios de planejamento

- Entregar em fatias verticais, não em camadas horizontais.
- Cada fase tem critério de saída verificável.
- Nenhuma fase começa sem spec aprovada e dependências resolvidas.
- Fairness e privacidade não são fases finais; entram desde a primeira fatia.
- A ordem pode mudar com evidência; mudanças exigem registro.
- Nenhum deploy em produção sem staging e aprovação humana.

---

## 2. Visão geral

```txt
F0  Consolidação e planejamento
F1  Fundação técnica
F2  Descoberta local
F3  Demanda e conexão
F4  Confiança e reputação
F5  Fairness e experimentação
F6  Preparação e execução do piloto
F7  Otimização e evolução
```

| Fase | Objetivo | Dono | Duração indicativa | Depende de |
| --- | --- | --- | --- | --- |
| F0 | PRD, specs, tickets, ADRs, pipeline | Todos | 1 semana | — |
| F1 | contas, perfis, catálogo, admin, observabilidade | Victor | 2–3 semanas | F0 |
| F2 | busca, ranking baseline, perfil público | Victor + Anderson | 2–3 semanas | F1 |
| F3 | mural, interesses, conexões, notificações | Anderson + Felipe | 2–3 semanas | F1 |
| F4 | avaliações, desfechos, denúncias, moderação | Felipe | 2–3 semanas | F3 |
| F5 | fairness, ledger, métricas, experimentos | Victor | 2 semanas | F2 |
| F6 | acessibilidade, segurança, operação, piloto | Todos | 2–4 semanas | F1–F5 |
| F7 | otimização e evolução | Todos | contínuo | F6 |

> Durações são estimativas de planejamento, não compromissos. Ajustar após refinamento dos tickets.

---

## 3. F0 — Consolidação e planejamento

### Objetivo

Transformar o PRD revisado em specs, tickets e pipeline executáveis.

### Entregas

- PRD v1.0 revisado e aprovado.
- Decisões Q-01..Q-50 registradas.
- Catálogo de requisitos, personas, métricas, riscos e questões abertas.
- Specs: foundation, web-experience, local-discovery, service-request, trust-reputation, fair-visibility.
- Critérios de aceite, tickets e planos de implementação.
- ADRs 0001–0006 aceitos.
- Harness, pipeline, RACI, task board e handoffs.
- Templates de review, QA e segurança.

### Critérios de saída

- [x] Specs prontas para revisão humana.
- [x] Dependências entre specs e tickets explícitas.
- [x] Donos definidos para todos os tickets.
- [x] Pipeline e gates documentados.
- [x] Nenhum código implementado.
- [x] Aprovação humana para iniciar F1 (foundation).

---

## 4. F1 — Fundação técnica

### Objetivo

Criar a base para que um profissional publique um perfil e um cliente possa encontrá-lo com segurança.

### Entregas

- Estrutura do monólito modular e contratos internos.
- Migrations, seed de catálogo e auditoria.
- Contas, sessão, confirmação de e-mail, reset e RBAC.
- Perfil profissional, área, disponibilidade e portfólio básico.
- Catálogo e sinônimos.
- Admin mínimo, MFA TOTP + recovery codes e auditoria.
- Privacidade: consentimento, PII, exportação/exclusão.
- Hardening: CSRF, rate limit, idempotência, criptografia de localização, paginação admin.
- Observabilidade: logs, métricas, health e feature flags.
- Web: cadastro, login, onboarding e perfil.

### Critérios de saída

- [ ] Jornada de cadastro → perfil publicado funciona em 360 px.
- [ ] API pública não retorna PII.
- [ ] Admin suspende perfil e auditoria registra.
- [ ] Testes unitários, integração, HTTP e acessibilidade verdes.
- [ ] Migração e rollback testados.
- [ ] Gate de segurança/privacidade sem bloqueios.

---

## 5. F2 — Descoberta local

### Objetivo

Permitir que o cliente encontre profissionais relevantes por serviço e localização, com ranking explicável.

### Entregas

- Busca por termo/categoria/localização.
- Filtros de tipo de atendimento, disponibilidade e verificação.
- Ranking baseline com pesos versionados.
- Explicação curta por resultado e página de transparência.
- Perfil público com portfólio, área aproximada e sinais.
- Eventos de funil e medição de busca.
- Guardrails básicos de fairness (caps, diversidade, fallback).

### Critérios de saída

- [ ] Busca presencial e remota funcionam.
- [ ] Ranking é determinístico para a mesma entrada/política.
- [ ] Nenhum atributo sensível é usado.
- [ ] Estado vazio e erro são claros.
- [ ] p95 da busca ≤ 500 ms no dataset de teste.
- [ ] Acessibilidade dos resultados e filtros validada.

---

## 6. F3 — Demanda e conexão

### Objetivo

Permitir que clientes publiquem necessidades e que profissionais se conectem com consentimento.

### Entregas

- Publicação, edição, fechamento e expiração de pedidos.
- Feed de oportunidades elegíveis.
- Interesses com limites e anti-spam.
- Aceite/recusa e conexão.
- Liberação de contato somente após consentimento.
- Notificações in-app/outbox.
- Bloqueio e denúncia em qualquer etapa.

### Critérios de saída

- [ ] Fluxo pedido → interesse → aceite → contato funciona ponta a ponta.
- [ ] Contato mascarado antes do aceite é testado.
- [ ] Limites de interesse e pedido são aplicados.
- [ ] Estados de conexão são válidos e auditáveis.
- [ ] Nenhuma PII em feed/notificação/evento.

---

## 7. F4 — Confiança e reputação

### Objetivo

Construir sinais honestos de confiança e moderação auditável.

### Entregas

- Registro de desfecho.
- Avaliação bilateral pós-conexão.
- Reputação contextual por categoria/papel.
- Resposta do avaliado.
- Verificação de contato.
- Denúncias, fila de moderação, ações e recurso.
- Detecção básica de autoavaliação e repetição.

### Critérios de saída

- [ ] Avaliação só é possível em conexão elegível.
- [ ] Uma avaliação por lado; autoavaliação bloqueada.
- [ ] Crítica legítima não é removida.
- [ ] Toda ação de moderação tem ator, motivo e timestamp.
- [ ] Denunciante não é exposto.
- [ ] Reputação de novo profissional não é zero.

---

## 8. F5 — Fairness e experimentação

### Objetivo

Tornar a distribuição de visibilidade mensurável, governável e reversível.

### Entregas

- Política de fairness versionada.
- Ledger de exposição pseudonimizado.
- Caps, exploração, boost de novo e diversidade.
- Métricas de Gini, top 10/20, share de novos e posição por coorte.
- Dashboard administrativo.
- Simulação offline e experimentos com guardrails.
- Relatório de fairness.

### Critérios de saída

- [ ] Política publicada é imutável.
- [ ] Ativação exige aprovação registrada.
- [ ] Fallback baseline funciona.
- [ ] Nenhum atributo sensível no schema.
- [ ] Simulação offline anexada.
- [ ] Guardrail de relevância monitorado.

---

## 9. F6 — Preparação e execução do piloto

### Objetivo

Validar o produto em uma região e categorias delimitadas, com segurança e evidência.

### Preparação

- Região, categorias e metas definidas.
- DPIA e revisão jurídica concluídas.
- Política de privacidade e termos publicados.
- Moderação operante com SLA.
- Acessibilidade auditada.
- Segurança revisada (ASVS, DAST, dependências).
- Observabilidade e alertas.
- Runbook e plano de incidente.
- Feature flags e rollback.
- Suporte e onboarding assistido.

### Execução

- 8–12 semanas, com marcos a cada 2 semanas.
- Acompanhamento de métricas e riscos.
- Experimentos com guardrails.
- Registro de aprendizados.

### Critérios de saída

- [ ] Metas de conexão e fairness avaliadas.
- [ ] Nenhum incidente crítico sem correção.
- [ ] WCAG 2.2 AA nas jornadas críticas.
- [ ] Relatório de aprendizado publicado.
- [ ] Decisão de continuidade registrada.

---

## 10. F7 — Otimização e evolução

### Objetivo

 Evoluir com base em evidência, sem comprometer fairness, privacidade ou confiança.

### Entregas potenciais

- Melhoria de matching e ranking.
- Expansão de categorias/regiões.
- Verificação de identidade.
- Notificações e canais adicionais.
- Ferramentas de moderação.
- Modelo de sustentabilidade.
- Integrações e parcerias.

### Gatilhos

- Busca p95 > 800 ms ou relevância limitada → avaliar índice dedicado.
- Volume de eventos > outbox → broker externo.
- Time > 8 pessoas ou deploy independente → extrair serviço.
- Moderação manual inviável → triagem assistida com revisão humana.
- Necessidade de ML → avaliação de fairness e feature store.

---

## 11. Marcos de entrega

| Marco | Critério | Dono |
| --- | --- | --- |
| M0 — Planejamento pronto | PRD + specs + tickets + pipeline | Todos |
| M1 — Fundação | cadastro → perfil publicado | Victor |
| M2 — Descoberta | busca explicável e acessível | Victor + Anderson |
| M3 — Conexão | pedido → aceite → contato | Anderson + Felipe |
| M4 — Confiança | avaliação + moderação | Felipe |
| M5 — Fairness | política + métricas + simulação | Victor |
| M6 — Piloto pronto | gates verdes + aprovação humana | Todos |

---

## 12. Riscos de cronograma

| Risco | Mitigação |
| --- | --- |
| Escopo do MVP cresce | separar fatia técnica de validação; cortar S/C |
| Dependência entre frentes | contratos e interfaces acordados antes; fatias verticais |
| Moderação subestimada | começar cedo; critérios e fila simples |
| Acessibilidade no fim | incluir por jornada, não por fase |
| Fairness tratada como extra | guardrails desde F2 |
| Migração SQLite/Postgres | isolar repositórios e testar contrato |

---

## 13. Governança do roadmap

- Revisão a cada 2 semanas.
- Mudança de fase exige critério de saída atendido ou waiver.
- Toda mudança de escopo atualiza PRD, spec e task board.
- Nenhuma fase pula gate de segurança/fairness.
- Aprendizados viram decisão registrada em `docs/decisions`.
