# foundation-platform — evidências

**Status:** evidência de testes de 2026-11-20 anexada; E2E/a11y/Postgres pendentes  
**Dono:** Victor  
**Revisores:** Felipe (QA/segurança), Anderson (produto/PII)

## Evidências disponíveis

| Arquivo | Conteúdo |
| --- | --- |
| [`2026-11-20-foundation-tests.md`](2026-11-20-foundation-tests.md) | resumo, comandos, cobertura, smoke e limitações |
| [`go-test-2026-11-20.txt`](go-test-2026-11-20.txt) | saída `go test ./... -v` |
| [`go-vet-2026-11-20.txt`](go-vet-2026-11-20.txt) | saída `go vet ./...` (vazia = sem achados) |
| [`go-build-2026-11-20.txt`](go-build-2026-11-20.txt) | saída `go build ./...` (vazia = verde) |

## O que ainda falta

- `go test -race ./...` em Linux com CGO;
- testes de contrato SQLite × Postgres;
- E2E de cadastro → onboarding → publicação;
- E2E de admin suspendendo/reativando;
- axe, teclado, 360 px, reduced motion;
- screenshots/vídeo;
- performance de API e busca;
- simulação de fairness (fora desta spec);
- pentest e `govulncheck`;
- evidência de migração/rollback em staging.

## Regra

Nenhuma evidência aqui aprova release, piloto ou produção. A aprovação exige
revisão independente e registro humano em `docs/process/approvals.md`.
