# foundation-platform — QA report

**Status:** QA automatizado de backend executado; web/E2E/a11y pendentes  
**Spec:** `spec.md`  
**Dono:** Victor  
**Execução:** 2026-11-20  
**Ambiente:** Windows 10/11, Go 1.27.1, SQLite (modernc), sem CGO  
**Commit:** alterações locais desta entrega

## 1. Sumário

- funções de teste: 72 em 23 arquivos;
- cenários de aceite cobertos por automação: 43 de 60;
- parciais/fora da fatia: 17;
- verdes: todos os testes executados;
- falhas: 0;
- cobertura total (pacotes internos): 64,6%;
- cobertura do pacote `internal/api`: 84,9%;
- bloqueios: E2E, a11y, Postgres e race não executados.

## 2. Matriz de aceite

Cobertos:

`AC-01..AC-19`, `AC-23`, `AC-28..AC-32`, `AC-37..AC-46`, `AC-48`, `AC-50..AC-55`, `AC-59`.

Parciais/fora da fatia:

| AC | Status | Dono |
| --- | --- | --- |
| AC-20..AC-22, AC-58 | não implementado | FND-08/FND-17 — Felipe |
| AC-24..AC-27, AC-47, AC-56 | não implementado | FND-07 — Anderson |
| AC-33 | login sem MFA rejeitado; endpoint sem sessão MFA não é criado | FND-04/09 |
| AC-34..AC-36, AC-49, AC-57, AC-60 | não implementado (web) | FND-11/12/13/18 |
| AC-41 | sem fluxo de cliente/pedido na fundação | `service-request` |

## 3. Testes executados

```txt
gofmt -l cmd internal db         -> sem saída
go vet ./...                     -> sem achados
go test ./... -count=1           -> todos os pacotes verdes
go build ./...                   -> verde
smoke do binário + curl          -> health 200, ready 200, signup, confirmação e login OK
```

Cobertura:

```txt
go test ./... -coverpkg=./internal/... -coverprofile=...  -> total 64,6%
go test ./internal/api -cover                            -> 84,9%
```

`go test -race ./...` não rodou porque o ambiente não tem CGO habilitado. O CI
Linux deve rodar com `CGO_ENABLED=1`.

## 4. Acessibilidade

- axe: não executado (sem web);
- teclado: não executado;
- leitor de tela: não executado;
- 360 px: não executado;
- contraste: não executado.

## 5. Performance

- API p95: não medida; testes locais em SQLite não representam produção;
- busca p95: fora desta spec;
- LCP: sem web;
- carga: não executada.

## 6. Fairness

- não se aplica à fundação; a spec não ordena resultados;
- nenhuma métrica de exposição foi implementada nesta fatia.

## 7. Falhas conhecidas

Nenhuma falha automatizada. Há lacunas funcionais deliberadas (portfólio, web,
direitos, Postgres, outbox worker), registradas em follow-ups.

## 8. Riscos residuais

- SQLite não valida o comportamento de produção em Postgres/PostGIS.
- Sem teste de concorrência real de sessão/rate limit entre instâncias.
- Sem E2E, acessibilidade ou teste de contrato com consumidor web.
- Sem teste de carga, caos ou restauração de backup.
- Evidência de código ainda não foi revisada por Felipe de forma independente.

## 9. Conclusão

`backend aprovado no QA automatizado; release não aprovado`. A revisão
independente em contexto separado foi executada e suas correções estão na suíte;
a aprovação formal de Felipe/Anderson continua pendente. A fundação está pronta
para o adaptador Postgres e para a web, não para piloto.
