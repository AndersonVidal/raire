# ADR 0006 — Web server-rendered mobile-first

**Status:** proposto  
**Data:** 2026-11-20  
**Decisores:** produto, frontend, arquitetura  
**Relacionado:** `docs/architecture/architecture-plan.md`

## Contexto

O público inclui profissionais com baixa familiaridade digital e clientes em celulares modestos. O MVP precisa de web acessível, rápida e simples, sem app nativo. O time é pequeno e o foco é validar o fluxo, não construir uma SPA complexa.

## Opções consideradas

1. **Server-rendered com `html/template` + CSS** — simples, acessível, sem build JS, baixo custo, SEO fácil.
2. **React + Vite (SPA/PWA)** — ecossistema rico, mas mais complexidade, bundle e estado.
3. **Next.js (SSR)** — bom SEO, mas mais infraestrutura e dependência.

## Decisão

Adotar web server-rendered com `html/template`, CSS próprio e JS mínimo quando necessário, no mesmo binário Go. Mobile-first, WCAG 2.2 AA, sem SPA no MVP. PWA e SPA podem ser avaliadas depois, com evidência de necessidade.

## Consequências

**Positivas:**

- menos complexidade e dependências;
- acessibilidade mais fácil de garantir;
- SEO básico sem esforço;
- um único deploy;
- carregamento rápido em conexão limitada.

**Negativas:**

- interações ricas exigem mais cuidado;
- menos ecossistema de componentes;
- pode limitar evolução para app.

**Mitigações:**

- design system simples e testado;
- JS progressivo quando necessário;
- componentes acessíveis;
- avaliar PWA/SPA com evidência.

## Critérios de reversão

- necessidade de interação em tempo real;
- requisito de app/PWA;
- produtividade do time cair.

## Status

Proposto. Requer aprovação humana antes da implementação.
