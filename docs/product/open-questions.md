# Questões Abertas e Decisões Pendentes — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0, seção 16  
**Uso:** desbloquear planejamento e implementação sem inventar requisito.

> Recomendações técnicas para Q-07 e Q-44–Q-50: `docs/status/2026-11-20-decision-draft.md`. Ainda **não são decisões**; aguardam aprovação humana.

---

## 1. Classificação

| Classe | Significado | Prazo para decisão |
| --- | --- | --- |
| Estratégica | muda escopo, mercado ou modelo | antes de F1 |
| Técnica | muda arquitetura, custo ou operação | antes da spec afetada |
| Produto | muda jornada ou prioridade | antes do ticket afetado |
| Jurídica | muda base legal, retenção ou risco | antes do piloto |
| Pesquisa | precisa de evidência de campo | antes de concluir o piloto |

---

## 2. Questões estratégicas e de mercado

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-01 | Qual cidade/região piloto? | densidade e aquisição | Anderson | antes de F6 | Cidade sintética nos testes; região real antes do piloto |
| Q-02 | Quais categorias iniciais? | escopo e catálogo | Anderson | antes de F1 | manutenção/reparos + aulas/consultoria |
| Q-03 | Qual o tamanho mínimo de oferta? | liquidez | Anderson | antes de F6 | 150 perfis completos |
| Q-04 | Qual o apetite para moderação manual? | custo e segurança | Felipe | antes de F4 | fila pequena, SLA 48 h |
| Q-05 | Qual o modelo de monetização futuro? | conflito com fairness | Fundador | pós-piloto | definir depois; nunca ranking pago |

---

## 3. Questões de identidade e perfil

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-06 | Verificação de identidade é obrigatória para aparecer? | confiança vs. fricção | Anderson + Felipe | antes de F4 | opcional; e-mail confirmado obrigatório |
| Q-07 | Um profissional pode ter mais de um perfil? | modelo de dados | Victor | antes de F1 | um perfil por conta no MVP |
| Q-08 | Área por raio é suficiente ou precisa bairro/polígono? | busca e privacidade | Victor | antes de F2 | cidade/bairro + raio |
| Q-09 | Como confirmar contato? | verificação | Felipe | antes de F4 | e-mail; telefone opcional |
| Q-10 | Qual o prazo de retenção de documentos? | LGPD | Jurídico | antes de F4 | 90 dias após decisão |

---

## 4. Questões de catálogo e portfólio

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-11 | Catálogo inicial por produto ou pesquisa? | relevância | Anderson | antes de F1 | seed curado por produto, revisado por pesquisa |
| Q-12 | Portfólio aceita imagem no MVP? | custo e moderação | Anderson | antes de F1 | sim, com limite e validação; sem upload obrigatório |
| Q-13 | Quem pode remover conteúdo de portfólio? | moderação | Felipe | antes de F4 | operador, com auditoria |

---

## 5. Questões de descoberta e ranking

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-14 | Busca inicial exige localização? | UX | Victor | antes de F2 | cidade/bairro; remoto permite "qualquer lugar" |
| Q-15 | Qual o raio padrão e máximo? | relevância | Victor | antes de F2 | 10 km padrão; 50 km máximo |
| Q-16 | Quantos resultados por página? | mobile e performance | Anderson | antes de F2 | 10 no mobile, 20 no desktop |
| Q-17 | Distância exata, faixa ou bairro? | privacidade | Felipe | antes de F2 | faixa aproximada ("até 2 km") |
| Q-18 | Reputação aparece na busca no MVP? | confiança | Anderson | antes de F2 | sim, resumida, com contagem |
| Q-19 | Percentual de exploração? | fairness | Victor | antes de F5 | 20% |
| Q-20 | Cap por sessão/dia? | fairness | Victor | antes de F5 | 2/10 |
| Q-21 | Boost de novo por quanto tempo? | fairness | Victor | antes de F5 | 30 dias ou 5 conexões |
| Q-22 | Piso de relevância? | fairness | Victor + Anderson | antes de F5 | percentil 40 |
| Q-23 | Guardrail de relevância? | fairness | Anderson | antes de F5 | queda ≤ 2 p.p. CTR / 5% nDCG |
| Q-24 | Quem aprova política de ranking? | governança | Fundador | antes de F5 | produto + pesquisa + segurança |

---

## 6. Questões de demandas, interesses e conexões

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-25 | Cliente pode aceitar mais de um profissional por pedido? | conexão | Anderson | antes de F3 | sim, até 3 |
| Q-26 | Limite de interesses por pedido/profissional? | anti-spam | Felipe | antes de F3 | 10 por pedido; 20/dia por profissional |
| Q-27 | Pedido expira em quantos dias? | liquidez | Anderson | antes de F3 | 7 dias |
| Q-28 | Contato é liberado por e-mail, telefone ou ambos? | privacidade | Felipe | antes de F3 | canal escolhido pelo titular; e-mail default |
| Q-29 | Cliente pode publicar sem localização exata? | privacidade | Victor | antes de F3 | sim, cidade/bairro |
| Q-30 | Como confirmar que o serviço ocorreu? | avaliação | Anderson | antes de F4 | confirmação de um lado + contestação |
| Q-31 | Moderação prévia ou pós-publicação? | operação | Felipe | antes de F3 | pós-publicação, com fila por risco |

---

## 7. Questões de confiança e reputação

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-32 | Avaliação exige confirmação mútua? | elegibilidade | Anderson | antes de F4 | não; um lado confirma, outro pode contestar |
| Q-33 | Janela de edição de avaliação? | integridade | Felipe | antes de F4 | 24 h |
| Q-34 | Verificação de identidade manual ou terceirizada? | custo | Felipe | antes de F4 | manual no MVP |
| Q-35 | Denúncias anônimas são permitidas? | segurança | Felipe | antes de F4 | não; identidade protegida |
| Q-36 | SLA de moderação para risco físico? | segurança | Felipe | antes de F4 | 4 h |
| Q-37 | Reputação aparece na busca ou só no perfil? | descoberta | Anderson | antes de F2 | resumida na busca; detalhada no perfil |

---

## 8. Questões de segurança, privacidade e conformidade

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-38 | Quais bases legais por finalidade? | LGPD | Jurídico | antes de F6 | proposta em `docs/security/privacy-and-lgpd.md` |
| Q-39 | Quais prazos de retenção? | LGPD | Jurídico | antes de F6 | proposta do PRD |
| Q-40 | Quem é o encarregado/DPO? | governança | Fundador | antes de F6 | a designar |
| Q-41 | Haverá transferência internacional? | fornecedores | Jurídico | antes de F6 | avaliar por fornecedor |
| Q-42 | Qual o plano de incidente? | resposta | Felipe | antes de F6 | runbook inicial |
| Q-43 | Pesquisa pode usar dados agregados? | ética | Felipe | antes de F6 | consentimento separado + anonimização |

---

## 9. Questões de arquitetura e engenharia

| ID | Pergunta | Impacto | Dono | Prazo | Default seguro |
| --- | --- | --- | --- | --- | --- |
| Q-44 | SQLite local e Postgres produção? | implementação | Victor | antes de F1 | sim; repositórios isolados |
| Q-45 | Frontend server-rendered ou SPA? | stack | Victor + Anderson | antes de F1 | server-rendered no MVP |
| Q-46 | Sessão por cookie ou token? | segurança | Victor | antes de F1 | cookie HttpOnly + CSRF |
| Q-47 | MFA admin via TOTP ou outro? | segurança | Felipe | antes de F1 | TOTP |
| Q-48 | Outbox interno ou broker? | operação | Victor | antes de F1 | outbox interno |
| Q-49 | Como isolar busca geoespacial? | portabilidade | Victor | antes de F2 | interface de repositório; Haversine no MVP, PostGIS depois |
| Q-50 | Feature flags em banco ou arquivo? | operação | Victor | antes de F1 | tabela + default seguro |

---

## 10. Registro e acompanhamento

Toda questão deve ter:

- ID e pergunta;
- classe;
- impacto;
- dono;
- prazo;
- default seguro;
- estado: `aberta`, `em análise`, `decidida`, `bloqueada`;
- decisão e data;
- link para ADR/spec quando aplicável.

### 10.1 Estados

| Estado | Significado |
| --- | --- |
| aberta | sem análise iniciada |
| em análise | dono está investigando |
| decidida | decisão registrada |
| bloqueada | depende de terceiro/humano |
| revisada | decisão reavaliada após evidência |

### 10.2 Processo de resolução

1. Dono investiga e propõe opções.
2. Alternativas e trade-offs são registrados.
3. Humano responsável decide.
4. Decisão vai para ADR/spec/PRD.
5. Task board é atualizado.

### 10.3 Participação de agentes de IA

Agentes podem:

- pesquisar e resumir alternativas;
- apontar riscos e lacunas;
- preparar minuta de decisão.

Agentes **não podem**:

- aprovar decisão final;
- marcar questão como decidida sem humano;
- tratar default como decisão.

---

## 11. Questões que bloqueiam a implementação

Antes de iniciar F1:

- Q-07 (um perfil por conta)
- Q-44 (SQLite/Postgres)
- Q-45 (server-rendered)
- Q-46 (sessão)
- Q-47 (MFA)
- Q-48 (outbox)
- Q-50 (feature flags)

Antes de F2:

- Q-08, Q-14, Q-15, Q-16, Q-17, Q-18, Q-37

Antes de F3:

- Q-25, Q-26, Q-27, Q-28, Q-29, Q-31

Antes de F4:

- Q-06, Q-09, Q-10, Q-30, Q-32, Q-33, Q-34, Q-35, Q-36

Antes de F5:

- Q-19, Q-20, Q-21, Q-22, Q-23, Q-24

Antes de F6:

- Q-01, Q-03, Q-38, Q-39, Q-40, Q-41, Q-42, Q-43
