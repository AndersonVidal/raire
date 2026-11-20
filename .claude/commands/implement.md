---
description: Implementa um ticket aprovado, testes primeiro
argument-hint: <feature> <ticket>
---
1. Abra a spec, o aceite e o ticket. Se a spec não estiver `aprovada`, pare.
2. Confirme o cenário de aceite e o DoR.
3. Escreva os testes do cenário (devem falhar).
4. Implemente até passarem.
5. Rode build, testes, vet, acessibilidade e segurança aplicáveis.
6. Atualize OpenAPI, migration e documentação.
7. Anexe evidência em `specs/<feature>/evidence/`.
8. Abra PR em `feat/<feature>` com link para spec/ticket.
9. Não faça merge.
