# CI/CD e Operação — RAIRE

**Status:** proposta para revisão humana  
**Versão:** 1.0  
**Relacionado:** `docs/process/pipeline.md`, `docs/process/qa-strategy.md`, `docs/process/security-privacy.md`

---

## 1. Objetivo

Definir um pipeline reproduzível, seguro e barato para o MVP, com evidência antes de merge e aprovação humana antes de piloto/produção.

---

## 2. Ambientes

| Ambiente | Uso | Dados | Deploy |
| --- | --- | --- | --- |
| Local | desenvolvimento e debugging | sintéticos | manual |
| CI | testes automatizados | efêmeros, sintéticos | automático por PR |
| Staging | aceite, E2E, segurança, performance | sintéticos/anonimizados | automático no merge |
| Piloto | uso real controlado | reais, com consentimento | aprovação humana |

---

## 3. Pipeline de PR

```txt
1. gofmt -l
2. go vet ./...
3. go test ./... -count=1
4. go test ./... -coverprofile=coverage.out
5. go build ./...
6. govulncheck ./...
7. gosec ./...
8. secret scan
9. migration test (banco vazio + seed)
10. contrato/OpenAPI (quando API)
11. acessibilidade (quando web)
12. build do binário
```

**Falha em qualquer etapa bloqueia merge.**

### 3.1 Evidência do PR

- saída dos testes com data e commit;
- cobertura por pacote;
- resultado de `govulncheck` e `gosec`;
- relatório de acessibilidade (quando web);
- migration e rollback;
- riscos residuais.

---

## 4. Pipeline de staging

```txt
1. build da imagem/binário
2. aplicar migrations
3. seed sintético
4. smoke test
5. E2E das jornadas críticas
6. DAST baseline
7. auditoria de acessibilidade
8. teste de performance
9. publicar evidência
```

### 4.1 Smoke test mínimo

- health e readiness;
- cadastro e login;
- publicação de perfil;
- busca com resultado;
- pedido → interesse → aceite;
- avaliação;
- denúncia;
- admin suspende perfil.

---

## 5. Release

### 5.1 Checklist

- [ ] pipeline verde;
- [ ] staging com smoke e E2E verdes;
- [ ] migração aplicada e rollback testado;
- [ ] feature flags com default seguro;
- [ ] dashboards e alertas;
- [ ] runbook e contato de incidente;
- [ ] comunicação de mudança;
- [ ] aprovação humana registrada.

### 5.2 Estratégia

- canário ou rolling;
- monitorar erro, latência e funil;
- rollback por flag ou imagem anterior;
- janela de observação antes de ampliar.

### 5.3 Rollback

- feature flag off;
- imagem anterior;
- migration compatível (expand/contract);
- comunicação e registro.

---

## 6. Migrations

- forward-only, versionadas e revisadas;
- expand/contract para mudanças destrutivas;
- toda migration tem plano de rollback lógico;
- seed idempotente;
- teste em banco vazio e com seed;
- backup antes de migração destrutiva.

---

## 7. Feature flags

| Flag | Default | Uso | Dono |
| --- | --- | --- | --- |
| `foundation.enabled` | off | libera cadastro/perfil | Victor |
| `discovery.enabled` | off | libera busca | Victor |
| `discovery.fairness.enabled` | off | libera camada de fairness | Victor |
| `requests.enabled` | off | libera mural/interesses | Anderson |
| `reviews.enabled` | off | libera avaliações | Felipe |
| `moderation.manual_queue` | on | fila manual | Felipe |
| `analytics.research_consent` | off | analytics/pesquisa | Felipe |
| `admin.enabled` | off | área admin | Felipe |

Toda flag tem dono, descrição, default seguro e data de revisão.

---

## 8. Observabilidade

### 8.1 Pilares

- **Logs:** JSON estruturado, correlation ID, nível, contexto, sem PII.
- **Métricas:** latência, erro, saturação, funil de busca, conexão, fairness.
- **Traces:** OpenTelemetry quando justificado.
- **Auditoria:** eventos de negócio e ações administrativas append-only.

### 8.2 SLOs iniciais

| SLO | Meta |
| --- | --- |
| Disponibilidade da API | 99,5% |
| Latência p95 de busca | < 500 ms |
| Latência p95 de APIs | < 300 ms |
| Erro 5xx | < 0,5% |
| Freshness de perfil na busca | < 60 s |
| Entrega de notificação in-app | > 99% em 5 min |

### 8.3 Alertas

- queda de conversão de busca;
- aumento de 5xx ou latência;
- pico de denúncias/spam;
- falha de outbox/worker;
- drift de fairness além do guardrail;
- falha de backup/restore;
- exposição de PII em logs (teste periódico).

---

## 9. Backup e continuidade

- backup diário do banco;
- retenção definida;
- criptografia em repouso;
- teste de restore trimestral;
- RPO/RTO documentados antes do piloto;
- runbook de recuperação.

---

## 10. Runbook mínimo

1. **Serviço fora:** verificar health, logs, banco, flag; rollback se necessário.
2. **Erro alto:** identificar endpoint, latência, dependência; mitigar.
3. **Denúncias em pico:** priorizar risco físico; comunicar.
4. **Suspeita de vazamento:** conter, preservar evidência, acionar segurança/DPO.
5. **Falha de migration:** parar deploy, restaurar backup, rollback lógico.
6. **Fairness fora do guardrail:** desligar flag, voltar política anterior, revisar.

---

## 11. Custos

- evitar serviços gerenciados sem justificativa;
- monitorar custo por usuário ativo e por conexão;
- orçamento mensal definido antes do piloto;
- revisar custos mensalmente.

---

## 12. Critérios de saída do pipeline

- [ ] PR com testes e evidência.
- [ ] Gates de review, segurança, trust & safety, QA e produto.
- [ ] Staging com smoke/E2E.
- [ ] Migração e rollback.
- [ ] Flags e monitoramento.
- [ ] Aprovação humana.
