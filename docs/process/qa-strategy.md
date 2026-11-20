# Estratégia de QA — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Relacionado:** `docs/process/review-criteria.md`, `specs/_template/acceptance.md`, `docs/team/raci.md`

---

## 1. Objetivo

Garantir que o MVP seja confiável, acessível, seguro e justo o suficiente para um piloto controlado, sem construir uma suíte cara e lenta antes de existir produto.

A estratégia segue o risco:

- quanto mais perto de dinheiro, privacidade ou fairness, mais rigor;
- quanto mais exploratório, mais aprendizado e menos automação prematura;
- todo critério de aceite tem um teste ou uma justificativa explícita.

---

## 2. Pirâmide de testes

```txt
        E2E (poucos, jornadas críticas)
      Integração (API + banco real)
    Contrato (OpenAPI + fixtures)
  Unitário (domínio, ranking, políticas)
```

### 2.1 Unitários

- invariantes de domínio;
- máquinas de estado;
- funções de ranking e fairness;
- cálculo de reputação;
- validação e normalização;
- regras de autorização;
- parsers e formatadores.

**Meta inicial:** 80% de cobertura nos pacotes de domínio; 90% nas funções puras de fairness/ranking.

### 2.2 Integração

- repositórios com banco real (SQLite local/CI; Postgres em staging);
- transações e outbox;
- migrations em banco vazio e com seed;
- concorrência e idempotência;
- consultas geoespaciais;
- jobs/worker.

**Nunca** usar mock de banco em teste de integração.

### 2.3 Contrato

- OpenAPI validado contra implementação;
- fixtures de request/response;
- compatibilidade de campos;
- erros padronizados;
- geração de cliente web quando aplicável.

### 2.4 E2E

Jornadas críticas:

1. Profissional cria conta, publica perfil e aparece na busca.
2. Cliente busca por serviço/localização e abre perfil.
3. Cliente publica pedido; profissional demonstra interesse; cliente aceita; contato é liberado.
4. Avaliação pós-conexão.
5. Denúncia e moderação.
6. Admin suspende perfil e auditoria registra.

Ferramenta proposta: Playwright, viewport 360×640 e desktop, teclado, `prefers-reduced-motion`.

### 2.5 Acessibilidade

- axe automatizado em todas as páginas críticas;
- teste manual de teclado;
- leitor de tela (NVDA/VoiceOver) nas jornadas principais;
- contraste e foco;
- formulários com erro específico;
- zoom 200%;
- movimento reduzido.

Checklist detalhado: `docs/process/accessibility-checklist.md`.

### 2.6 Performance

- k6 ou ferramenta equivalente;
- busca p95 < 500 ms;
- APIs p95 < 300 ms;
- carga de 2× o piloto;
- LCP < 2,5 s em 4G mid-range;
- teste de degradação quando ranking/fairness falha.

### 2.7 Segurança

- `govulncheck`;
- `gosec`/SAST;
- secret scan;
- dependency scan;
- DAST baseline em staging;
- teste de autorização por recurso;
- teste de rate limit;
- teste de PII em logs/eventos.

### 2.8 Fairness

- golden set de consultas;
- simulação offline de políticas;
- testes de propriedade (caps, elegibilidade, determinismo);
- métricas por coorte;
- comparação baseline vs. política candidata;
- guardrails de relevância;
- revisão humana dos resultados.

---

## 3. Matriz de aceite

Toda spec tem `acceptance.md` com:

| Cenário | Requisito | Tipo | Arquivo de teste | Status |
| --- | --- | --- | --- | --- |
| ... | RF-xx/R-xx | U/I/C/E2E/A11y/Perf/Fairness/Sec | caminho | pendente/verde |

Regras:

- todo cenário Gherkin tem teste ou justificativa;
- cenário sem teste é bloqueio de review;
- nome do teste referencia o ID do cenário quando possível;
- evidência é anexada em `evidence/`.

---

## 4. Dados de teste

- Dados sintéticos, nunca PII real.
- Seed reproduzível com perfis, categorias, áreas e avaliações.
- Fixtures de localização com coordenadas fictícias.
- Cenários de borda: sem resultado, área vazia, perfil incompleto, localização inválida, muitos resultados.
- Dados de fairness: coortes balanceadas e casos de concentração.
- Nenhum segredo real em fixture.

---

## 5. Ambientes

| Ambiente | Uso | Dados |
| --- | --- | --- |
| Local | desenvolvimento e debugging | sintéticos |
| CI | testes automatizados | efêmeros, sintéticos |
| Staging | aceite, E2E, segurança, performance | sintéticos + anonimizados |
| Piloto | uso real controlado | reais, com consentimento e LGPD |

---

## 6. Evidência

Cada entrega anexa:

- comando e saída dos testes;
- commit testado;
- relatório de acessibilidade;
- relatório de performance;
- simulação de fairness (quando aplicável);
- screenshots/vídeo das jornadas;
- riscos residuais e falhas conhecidas.

Formato sugerido:

```txt
specs/<feature>/evidence/<data>/
├── README.md
├── unit.txt
├── integration.txt
├── e2e.txt
├── a11y.json
├── perf.json
├── fairness.json
└── screenshots/
```

---

## 7. Critérios de saída do piloto

- Jornadas críticas com taxa de sucesso ≥ 90%.
- Nenhum incidente crítico de privacidade/segurança.
- WCAG 2.2 AA nas jornadas críticas.
- Performance dentro do orçamento.
- Fairness monitorada e dentro dos guardrails acordados.
- Falhas conhecidas priorizadas e comunicadas.
- Relatório de aprendizado publicado em `docs/experiments`.

---

## 8. O que não fazer

- Automatizar tudo antes de validar o fluxo.
- Usar dados reais em teste.
- Esconder falha com `.skip`.
- Testar só o caminho feliz.
- Medir clique como sucesso sem medir conexão.
- Rodar experimento de fairness sem guardrail.
- Declarar "pronto" sem evidência.

---

## 9. Responsabilidades

| Atividade | Dono | Apoio |
| --- | --- | --- |
| Plano de teste por spec | Felipe | autor da feature |
| Testes unitários de domínio | autor | Victor (backend) / Anderson (web) |
| Testes de integração | autor | Felipe |
| E2E e acessibilidade | Felipe | Anderson |
| Performance | Victor | Felipe |
| Fairness | Victor | Felipe |
| Evidência e relatório | Felipe | autor |
