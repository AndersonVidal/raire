# web-experience

**Status:** em revisão  
**Versão:** 1.0  
**RFs:** suporta RF-01..RF-19 (interface); RF-06, RF-08, RF-10, RF-12, RF-14, RF-17  
**RNFs:** RNF-10, RNF-11, RNF-12, RNF-02  
**Módulos:** web, api (contratos), platform  
**Depende de:** `foundation-platform`; ADR 0006  
**Dono:** Anderson  
**Revisores:** Victor (contratos), Felipe (a11y/QA)  
**Aprovação humana:** pendente

---

## 1. Contexto

O RAIRE é mobile-first e precisa funcionar para profissionais com baixa familiaridade digital e clientes em celulares modestos. Esta spec define a experiência web, o design system mínimo, os estados, a acessibilidade e a copy.

Ela não define regras de negócio; consome as specs de domínio. O objetivo é garantir que as jornadas sejam compreensíveis, acessíveis e honestas.

**Problema do usuário:** "não quero aprender um sistema; quero resolver o que preciso em poucos toques, sem me perder."

---

## 2. Objetivo e não objetivo

**Objetivo:** entregar layout, componentes, telas, copy PT-BR, estados e requisitos de acessibilidade para as jornadas críticas.

**Não objetivo:**

- definir regra de negócio;
- criar SPA/app nativo;
- design visual sofisticado;
- animações complexas;
- gamificação;
- dark patterns.

---

## 3. Princípios e restrições

- Mobile-first, funcional em 360 px.
- WCAG 2.2 AA nas jornadas críticas.
- Linguagem simples, sentence case, sem jargão.
- Uma ação principal por tela.
- Rótulos persistentes, ajuda contextual e erro específico.
- Estados de carregando, vazio, erro e sucesso sempre.
- `prefers-reduced-motion` respeitado.
- Baixo consumo de dados; sem autoplay.
- Nenhuma PII, contato ou localização precisa na interface pública.
- Sem urgência falsa, dark pattern ou pay-to-win.

---

## 4. Requisitos

### A. Design system

- **W-01** Tokens de cor, tipografia, espaçamento, foco, borda e movimento documentados.
- **W-02** Contraste mínimo 4,5:1 para texto normal e 3:1 para texto grande.
- **W-03** Alvos de toque com pelo menos 44×44 px.
- **W-04** Componentes: botão, link, input, select, textarea, checkbox, radio, alert, badge, card, paginação, dialog, skeleton.
- **W-05** Componentes têm estados default, hover, focus, active, disabled, loading e error.
- **W-06** Foco visível em todos os elementos interativos.
- **W-07** Componentes não dependem apenas de cor para transmitir informação.

### B. Layout e navegação

- **W-08** Layout base com skip link, header, main, footer e landmarks.
- **W-09** Navegação principal com no máximo 5 itens.
- **W-10** Layout funciona em 360 px, 768 px e 1280 px sem rolagem horizontal.
- **W-11** Header mostra contexto de conta e ação de sair.
- **W-12** Mensagens de erro globais usam `role="alert"`/`aria-live`.
- **W-13** Título de página é único e descritivo.
- **W-14** Idioma declarado como `pt-BR` e `lang` correto.

### C. Jornadas

- **W-15** Cadastro e login com validação inline e mensagem específica.
- **W-16** Confirmação de e-mail e reset com estado de sucesso/erro.
- **W-17** Onboarding de perfil em etapas curtas, com progresso e salvamento.
- **W-18** Prévia pública do perfil antes de publicar.
- **W-19** Busca com campo, categoria, localização e filtros.
- **W-20** Resultados com nome, categorias, distância aproximada, sinais, explicação e ação.
- **W-21** Estado vazio com sugestões de ajuste.
- **W-22** Perfil público com bio, serviços, área aproximada, disponibilidade, portfólio e sinais.
- **W-23** Publicação de pedido com categoria, descrição, região, urgência e orçamento opcional.
- **W-24** Lista de pedidos do cliente e detalhe com interessados.
- **W-25** Feed de oportunidades e detalhe.
- **W-26** Ação de interesse com mensagem curta.
- **W-27** Conexão com contato mascarado antes do aceite e liberado depois.
- **W-28** Avaliação com nota, tags e comentário opcional.
- **W-29** Denúncia com motivo e descrição.
- **W-30** Admin de denúncias, auditoria e fairness.
- **W-31** Página de transparência do ranking.
- **W-32** Página de privacidade, consentimentos, exportação e exclusão.

### D. Acessibilidade

- **W-33** Todas as jornadas críticas funcionam por teclado.
- **W-34** Leitor de tela anuncia título, erro, sucesso e mudança de contexto.
- **W-35** Formulários têm `label` associado, `aria-describedby` para ajuda e erro.
- **W-36** Imagens têm `alt` apropriado; decorativas com `alt=""`.
- **W-37** Vídeos/áudio (se houver) têm legenda/transcrição.
- **W-38** Zoom 200% sem perda de conteúdo.
- **W-39** `prefers-reduced-motion` desativa animações não essenciais.
- **W-40** Não há armadilha de foco; dialog devolve foco.
- **W-41** Tabelas têm cabeçalho e escopo quando usadas.

### E. Copy e conteúdo

- **W-42** Copy em PT-BR, sem jargão técnico.
- **W-43** Não prometer "melhor", "garantido", "100% verificado" ou antecedentes.
- **W-44** Explicar o próximo passo em cada tela.
- **W-45** Mensagens de erro dizem o que fazer, não só o que falhou.
- **W-46** Não usar urgência falsa, contagem regressiva ou pressão.
- **W-47** Explicar o que é público e o que é privado no perfil.

### F. Performance e responsividade

- **W-48** Orçamento: HTML+CSS+JS inicial ≤ 200 KB comprimido por página.
- **W-49** Imagens otimizadas e com dimensões declaradas.
- **W-50** LCP p75 ≤ 2,5 s em 4G mid-range.
- **W-51** INP p75 ≤ 200 ms; CLS p75 ≤ 0,1.
- **W-52** Funciona sem JS para leitura e formulários básicos; JS é progressivo.

---

## 5. Regras de negócio e invariantes

- interface nunca mostra PII pública;
- contato só aparece após aceite;
- localização é aproximada;
- erro de formulário não perde dados preenchidos;
- ação destrutiva pede confirmação;
- admin é inacessível sem MFA;
- denúncia não expõe denunciante;
- copy não promete verificação inexistente.

---

## 6. Estados e transições de UI

Cada tela deve ter:

```txt
loading | empty | error | success | partial
```

Estados globais:

- sessão expirada → mensagem e link de login;
- sem permissão → explicação e ação segura;
- rate limit → mensagem e tempo de espera;
- offline → aviso e retry quando possível.

---

## 7. Contratos e dados

A web consome os endpoints de `docs/architecture/api-contracts.md`. Não acessa banco.

### 7.1 Páginas

```txt
/                              home
/criar-conta                   cadastro
/entrar                        login
/confirmar-email               confirmação
/recuperar-senha               reset
/painel                        painel
/perfil                        perfil do profissional
/buscar                        busca
/profissional/{id}             perfil público
/pedidos                       pedidos do cliente
/pedidos/novo                  novo pedido
/pedidos/{id}                  detalhe
/oportunidades                 feed do profissional
/conexoes/{id}                 conexão
/avaliar/{connectionId}        avaliação
/denunciar                     denúncia
/notificacoes                  notificações
/minha-privacidade             privacidade
/transparencia                 ranking
/admin/denuncias               moderação
/admin/auditoria               auditoria
/admin/fairness                fairness
```

### 7.2 Componentes mínimos

- `Button`, `Link`, `Input`, `Select`, `Textarea`, `Checkbox`, `Radio`;
- `Alert`, `Badge`, `Card`, `Pagination`, `Dialog`, `Skeleton`;
- `SearchForm`, `ProfileCard`, `RequestCard`, `ConnectionPanel`, `ReviewForm`, `ReportForm`.

---

## 8. Casos de borda

- usuário sem JS;
- conexão lenta;
- sessão expirada no meio do formulário;
- erro de validação em campo específico;
- imagem quebrada;
- lista vazia;
- muitos resultados;
- texto longo;
- nome com acento/emoji;
- leitor de tela em formulário de múltiplas etapas;
- zoom 200%;
- `prefers-reduced-motion`.

---

## 9. Erros e mensagens

| Situação | Mensagem |
| --- | --- |
| campo obrigatório | "Preencha este campo." |
| e-mail inválido | "Informe um e-mail válido." |
| senha curta | "Use pelo menos 10 caracteres." |
| sessão expirada | "Sua sessão expirou. Entre novamente." |
| sem resultado | "Nenhum profissional encontrado. Tente ampliar a região ou mudar a categoria." |
| rate limit | "Muitas tentativas. Aguarde alguns minutos." |
| contato protegido | "O contato é liberado após o aceite." |
| erro interno | "Algo deu errado. Tente novamente." |

---

## 10. Privacidade, segurança e fairness

- **PII:** nunca exibir e-mail/telefone/documento/coordenada em página pública.
- **Authz:** esconder ações não permitidas; ainda validar no servidor.
- **Rate limit:** feedback claro, sem revelar detalhes.
- **Fairness:** explicação curta e honesta; link para transparência.
- **Acessibilidade:** W-33 a W-41.

---

## 11. Observabilidade e métricas

- eventos de funil: cadastro, busca, perfil, pedido, interesse, conexão, avaliação;
- métricas de usabilidade: conclusão, erro, abandono;
- não coletar PII;
- respeitar consentimento de analytics.

---

## 12. Fora de escopo

- SPA/app nativo;
- tema escuro completo;
- animações complexas;
- internacionalização;
- chat;
- pagamento;
- mapa interativo avançado;
- gamificação.

---

## 13. Dependências e riscos

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| onboarding longo | abandono | etapas curtas, salvamento progressivo |
| acessibilidade no fim | retrabalho | testar por jornada |
| bundle grande | performance | orçamento e JS mínimo |
| copy enganosa | perda de confiança | revisão de produto |
| contato exposto | risco | teste de contrato e revisão |
| responsividade quebrada | exclusão | teste em 360 px |

---

## 14. Critérios de aceite

Ver `acceptance.md`. Cobertura mínima:

- design tokens e componentes;
- navegação e landmarks;
- cadastro/login/onboarding;
- busca/resultados/perfil;
- pedido/interesse/conexão;
- avaliação/denúncia;
- admin/fairness/transparência;
- acessibilidade e 360 px;
- performance.

---

## 15. Gates de review específicos

- [ ] WCAG 2.2 AA nas jornadas críticas;
- [ ] 360 px sem rolagem horizontal;- [ ] teclado e foco;
- [ ] leitor de tela;
- [ ] contraste;
- [ ] estados de UI;
- [ ] copy PT-BR honesta;
- [ ] sem PII;
- [ ] orçamento de performance;
- [ ] `prefers-reduced-motion`.

---

## 16. Perguntas para o humano

- [ ] Q-11: catálogo inicial.
- [ ] Q-14: busca exige localização?
- [ ] Q-15: raio padrão/máximo.
- [ ] Q-16: resultados por página.
- [ ] Q-17: distância exata/faixa/bairro.
- [ ] Q-18: reputação na busca?
- [ ] Q-25..Q-31: pedido/interesse/conexão.
- [ ] Q-37: reputação na busca ou só no perfil.
- [ ] Q-45: server-rendered no MVP.

---

## 17. Rastreabilidade

| Requisito | RF/RNF | Ticket | Teste |
| --- | --- | --- | --- |
| W-01..W-07 | RNF-10, RNF-11 | WEB-01 | AC-01..AC-05 |
| W-08..W-14 | RNF-10, RNF-12 | WEB-02 | AC-06..AC-10 |
| W-15..W-18 | RF-01, RF-02, RF-03 | WEB-03, WEB-04 | AC-11..AC-16 |
| W-19..W-22 | RF-06, RF-07 | WEB-05, WEB-06 | AC-17..AC-22 |
| W-23..W-27 | RF-08..RF-12 | WEB-07, WEB-08 | AC-23..AC-29 |
| W-28..W-30 | RF-14, RF-17, RF-18 | WEB-08, WEB-09 | AC-30..AC-34 |
| W-31..W-32 | RF-07, RF-19 | WEB-10 | AC-35..AC-37 |
| W-33..W-41 | RNF-10 | WEB-11 | AC-38..AC-43 |
| W-42..W-47 | RNF-11 | WEB-12 | AC-44..AC-46 |
| W-48..W-52 | RNF-02, RNF-12 | WEB-11 | AC-47..AC-49 |
