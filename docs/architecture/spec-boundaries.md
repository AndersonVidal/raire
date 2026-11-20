# Fronteiras entre Specs — RAIRE

**Status:** aprovado para planejamento em 2026-11-20  
**Versão:** 1.0  
**Uso:** evitar duplicidade de domínio, dependência circular e ticket órfão.  
**Relacionado:** `docs/architecture/architecture-plan.md`, `specs/README.md`

---

## 1. Regra geral

Cada capacidade tem **uma spec dona**. As outras specs consomem por contrato, evento ou read model; não reimplementam a regra.

```txt
foundation-platform
  ├── web-experience (consome)
  ├── local-discovery (consome perfil/catálogo; dono da busca)
  │     └── fair-visibility (dono da política de fairness)
  └── service-request (consome conta/perfil; dono de pedido/interesse/conexão)
        └── trust-reputation (dono de desfecho/avaliação/denúncia/reputação)
```

---

## 2. Matriz de responsabilidade

Legenda: **D** = dona; **C** = consome; **H** = hook/evento; **—** = fora.

| Capacidade | foundation | web | discovery | service-request | trust | fair |
| --- | --- | --- | --- | --- | --- | --- |
| Conta, sessão, papéis | D | C | C | C | C | — |
| Consentimento, exportação, exclusão | D | C | — | — | C | — |
| Perfil profissional, área, disponibilidade | D | C | C | C | C | — |
| Portfólio | D | C | C | — | C | — |
| Catálogo | D | C | C | C | — | — |
| Busca, filtros, paginação | — | C | D | C | — | C |
| Ranking baseline e explicação | — | C | D | — | — | C |
| Política de fairness, ledger, experimentos | — | C | C | — | C | D |
| Pedido, interesse, conexão | — | C | — | D | C | — |
| Contato mascarado/liberado | C | C | — | D | C | — |
| Desfecho | — | C | — | C | D | — |
| Avaliação e reputação | — | C | C | — | D | C |
| Denúncia | H | C | — | H | D | — |
| Moderação, recurso, auditoria de conteúdo | — | C | — | — | D | — |
| Suspensão de perfil | D | C | C | C | C | — |
| Verificação de contato | D | C | C | C | C | — |
| Verificação de identidade | H | C | C | — | D | — |
| Transparência de ranking | C | D | C | — | — | D |
| Acessibilidade e design system | C | D | C | C | C | C |
| Métricas de produto | C | C | D | C | C | D |

---

## 3. Contratos de integração

### 3.1 Foundation → Discovery

- Eventos: `directory.profile_published`, `directory.profile_updated`, `directory.profile_paused`, `directory.profile_suspended`.
- Read model: `professional_search_document` é **dono da discovery**; foundation não escreve nele.
- Foundation expõe dados públicos via API/read model; nunca coordenada precisa.

### 3.2 Foundation → Service Request

- Foundation fornece conta, e-mail confirmado, perfil publicado, preferência de contato e máscara.
- Service Request é dona de pedido, interesse, conexão e liberação de contato.
- Foundation não cria tabela de conexão nem libera contato.

### 3.3 Service Request → Trust

- Service Request emite `demand.connection_accepted` e `demand.connection_completed`.
- Trust é dona de desfecho, avaliação e reputação.
- Trust lê conexão por interface/evento; não altera estado da conexão.

### 3.4 Trust → Foundation

- Trust pode suspender perfil via caso de uso de moderação; foundation executa a suspensão e a auditoria.
- Trust não escreve em `professional_profiles` diretamente.
- Verificação de identidade é da trust; foundation mantém o enum/shell e o selo de contato.

### 3.5 Discovery → Fairness

- Discovery chama a política de fairness ativa; fairness não faz HTTP.
- Fairness devolve ordem/ slots e explicação; discovery monta resposta.
- Fallback: se fairness falhar, discovery usa baseline.

### 3.6 Web → todas

- Web não contém regra de negócio.
- Web consome API `/v1` e páginas server-rendered.
- Web é dona de design system, copy, estados, acessibilidade e testes de UI.

---

## 4. Regras de dependência

1. `foundation` não importa `service-request`, `trust` ou `fairness`.
2. `service-request` pode importar `foundation` (ids/perfil) e `catalog`; não importa `trust`.
3. `trust` pode importar `foundation` e `service-request` (read); não importa `fairness`.
4. `discovery` pode importar `foundation` (read) e `fairness`; não importa `service-request`.
5. `fairness` importa apenas `platform` (config/clock).
6. `web` importa contratos; nunca adapters de contexto.
7. Nenhum contexto importa `adapters` de outro.
8. Eventos são versionados, imutáveis e sem PII.

---

## 5. O que não pode acontecer

- Foundation implementar avaliação/reputação.
- Discovery calcular reputação ou moderação.
- Service Request liberar contato sem consentimento.
- Trust alterar status de perfil diretamente.
- Fairness fazer query HTTP ou acessar PII.
- Web duplicar regra de negócio.
- Duas specs criarem a mesma tabela.
- Um ticket de uma spec exigir mudança de schema de outra sem ADR.

---

## 6. Critérios de review de fronteira

- [ ] A spec declara o que é dona e o que consome.
- [ ] Não há tabela duplicada.
- [ ] Não há regra de negócio em handler HTTP.
- [ ] Eventos têm payload mínimo e sem PII.
- [ ] Contrato de leitura é explícito.
- [ ] Fallback e erro entre contextos estão definidos.
- [ ] Teste de arquitetura cobre a regra.
- [ ] ADR registra mudança estrutural.
