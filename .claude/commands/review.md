---
description: Executa review em sequência
argument-hint: <feature>
---
1. Leia a spec, o aceite, o ticket e o diff.
2. Rode os gates na ordem:
   - reviewer (código/arquitetura);
   - security-privacy;
   - trust-safety;
   - qa;
   - product-manager.
3. Registre bloqueios com arquivo:linha e requisito.
4. Consolide em `specs/<feature>/review.md`.
5. Não edite o código revisado.
6. Bloqueio impede merge.
