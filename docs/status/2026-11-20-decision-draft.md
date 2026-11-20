# Minuta de Decisão — Q-07 e Q-44 a Q-50

**Status:** aprovado e integrado em 2026-11-20  
**Data:** 2026-11-20  
**Dono:** Victor  
**Revisores:** Felipe (segurança/QA), Anderson (produto/web)  
**Regra:** as recomendações foram aprovadas por humano e registradas em `docs/product/decisions-2026-11-20.md`.

---

## 1. Resumo

Estas 8 decisões destravam o início da implementação da fundação (`foundation-platform`). A recomendação busca o caminho mais simples que ainda seja seguro, testável e reversível.

| ID | Decisão recomendada | Quem aprova |
| --- | --- | --- |
| Q-07 | Um perfil profissional por conta no MVP | Victor |
| Q-44 | SQLite local/CI + PostgreSQL/PostGIS em produção | Victor + Felipe |
| Q-45 | Web server-rendered no MVP | Victor + Anderson |
| Q-46 | Sessão por cookie `HttpOnly` + CSRF | Victor + Felipe |
| Q-47 | MFA TOTP para admin/operador, com códigos de recuperação | Felipe |
| Q-48 | Outbox interno, sem broker externo | Victor |
| Q-49 | Busca geoespacial atrás de interface; Haversine local e PostGIS em produção | Victor |
| Q-50 | Feature flags em tabela no banco, com default seguro | Victor + Felipe |

---

## 2. Q-07 — Um profissional pode ter mais de um perfil?

**Recomendação:** **não no MVP; um perfil por conta.**

**Por quê:**

- o modelo fica simples (`professional_profiles.account_id` único);
- evita confusão de qual perfil aparece na busca;
- simplifica autorização, moderação e reputação;
- multi-perfil pode ser adicionado depois com migração e evidência.

**Impacto na spec:**

- `foundation-platform` R-07 e R-20 permanecem como estão;
- `data-model.md` mantém `account_id UNIQUE`;
- tickets FND-06 e FND-09 não precisam tratar múltiplos perfis.

**O que fica fora:** multi-perfil, estúdio com vários profissionais, CNPJ.

---

## 3. Q-44 — SQLite local e PostgreSQL em produção?

**Recomendação:** **sim.**

**Por quê:**

- SQLite permite rodar testes e o app local sem instalar banco;
- PostgreSQL + PostGIS é o alvo de produção para distância, índices e concorrência;
- a diferença fica confinada ao adaptador de repositório;
- reduz atrito para o time e para o CI.

**Condições:**

- domínio e casos de uso não importam `database/sql` diretamente;
- testes de contrato do repositório rodam nos dois bancos quando possível;
- staging usa PostgreSQL desde a F1;
- migrations revisadas para os dois dialetos;
- query geoespacial isolada (ver Q-49).

**Risco aceito:** um teste pode passar em SQLite e falhar em Postgres; mitigação com staging e testes de contrato.

---

## 4. Q-45 — Web server-rendered ou SPA?

**Recomendação:** **server-rendered no MVP.**

**Por quê:**

- menos JavaScript e menos bundle em celular modesto;
- mais simples de garantir acessibilidade e SEO;
- um único deploy (Go + HTML);
- o time é pequeno e o foco é validar o fluxo, não construir uma SPA.

**Condições:**

- `html/template` + CSS + JS progressivo quando necessário;
- formulários com CSRF;
- componentes acessíveis;
- ADR 0006 atualizado para `aceito`.

**O que fica fora:** SPA, PWA completa, app nativo.

---

## 5. Q-46 — Sessão por cookie ou token?

**Recomendação:** **cookie de sessão `HttpOnly` + CSRF.**

**Por quê:**

- funciona bem com página renderizada no servidor;
- permite revogar sessão no logout, troca de senha e suspensão;
- evita token em `localStorage`, que amplia risco de XSS;
- CSRF é necessário, mas é um controle conhecido e testável.

**Condições:**

- cookie `HttpOnly`, `Secure` em HTTPS, `SameSite=Lax`;
- token de sessão armazenado como hash;
- rotação de sessão após login;
- CSRF token em POST/PATCH/DELETE;
- expiração e revogação testadas.

---

## 6. Q-47 — MFA do admin via TOTP?

**Recomendação:** **sim, TOTP para `admin` e `operator`, com códigos de recuperação de uso único.**

**Por quê:**

- não depende de SMS ou serviço pago;
- funciona offline;
- é um segundo fator real;
- códigos de recuperação evitam lockout operacional.

**Condições:**

- segredo TOTP cifrado no banco;
- janela de tolerância curta (±1 período);
- rate limit no login admin;
- auditoria de ativação/uso/recuperação;
- códigos de recuperação armazenados como hash e uso único.

**Quem decide:** Felipe; Victor implementa.

---

## 7. Q-48 — Outbox interno ou broker externo?

**Recomendação:** **outbox interno no MVP.**

**Por quê:**

- evento é gravado na mesma transação do dado;
- não exige Kafka/NATS/RabbitMQ;
- worker interno processa com retry e idempotência;
- suficiente para notificação, indexação e métricas do piloto.

**Condições:**

- tabela `outbox_events` com `event_id`, `type`, `payload`, `attempts`, `processed_at`;
- processamento idempotente por `event_id`;
- alerta de backlog;
- expurgo/arquivamento;
- broker só com gatilho de volume/latência.

---

## 8. Q-49 — Como isolar a busca geoespacial?

**Recomendação:** **interface de repositório; Haversine no adaptador local e PostGIS no adaptador de produção.**

**Por quê:**

- o domínio não conhece SQL nem PostGIS;
- testes locais não precisam de extensão geoespacial;
- produção usa índice GiST e funções nativas;
- troca de banco/search engine fica isolada.

**Condições:**

- interface `GeoSearch`/`ProfileSearchRepository` com `FindWithinRadius`;
- cálculo de distância documentado;
- precisão pública sempre aproximada;
- coordenada precisa fora do read model público;
- teste de contrato para os dois adaptadores.

---

## 9. Q-50 — Feature flags em banco ou arquivo?

**Recomendação:** **tabela no banco, com default seguro.**

**Por quê:**

- permite ligar/desligar sem deploy;
- registra dono, descrição e data de revisão;
- permite auditoria de mudança;
- default `off` reduz risco.

**Condições:**

- flag desconhecida resolve `off`;
- mudança de flag gera `audit_event`;
- flags críticas: `foundation.enabled`, `admin.enabled`, `privacy.export_enabled`, `catalog.use_synonyms`;
- cache curto para não consultar o banco em toda request;
- rollback por flag.

---

## 10. Decisões complementares (aprovadas em 2026-11-20)

Estas não eram decisões do Victor, mas afetam a fundação. Foram aprovadas junto com o pacote:

| ID | Pergunta | Dono | Decisão |
| --- | --- | --- | --- |
| Q-02 | Categorias iniciais | Anderson | manutenção/reparos + aulas/consultoria |
| Q-08 | Área por raio ou bairro/polígono | Victor | cidade/bairro + raio |
| Q-09 | Confirmação de contato | Felipe | e-mail; telefone opcional |
| Q-10 | Retenção de documentos | Jurídico | 90 dias após decisão (fora do 1º ciclo) |
| Q-11 | Catálogo inicial | Anderson | seed curado por produto |
| Q-12 | Portfólio com imagem | Anderson | sim, opcional e validada |
| Q-13 | Remoção de portfólio | Felipe | operador, com auditoria |

---

## 11. Registro da decisão

As recomendações foram aprovadas em 2026-11-20 e registradas em:

1. `docs/product/decisions-2026-11-20.md` (registro canônico);
2. `docs/process/approvals.md` (log de aprovação);
3. `docs/product/open-questions.md` (status decidido);
4. ADRs 0001–0006 (status `aceito`);
5. `specs/foundation-platform/spec.md` (versão 1.1 aprovada).

## 12. Veredito

As 8 recomendações foram aprovadas em 2026-11-20. A spec `foundation-platform` v1.1 foi aprovada após a integração dos must-fix e should-fix. A implementação pode começar pelo `FND-01`; o piloto continua bloqueado até DPIA/jurídico.
