# Checklist de Acessibilidade — RAIRE

**Status:** vinculante para jornadas críticas  
**Versão:** 1.0  
**Donos:** Anderson (produto/web), Felipe (QA)  
**Referência:** WCAG 2.2 AA

---

## 1. Jornadas críticas

- cadastro e login;
- confirmação de e-mail e reset;
- onboarding de perfil;
- busca e resultados;
- perfil público;
- publicação de pedido;
- interesse e aceite;
- conexão e contato;
- avaliação;
- denúncia;
- admin.

---

## 2. Estrutura

- [ ] `lang="pt-BR"` no `html`.
- [ ] Um `h1` por página.
- [ ] Hierarquia de headings sem pular níveis.
- [ ] Skip link para o conteúdo principal.
- [ ] Landmarks: `header`, `nav`, `main`, `footer`.
- [ ] Título de página descritivo e único.
- [ ] Links com texto significativo (não "clique aqui").

---

## 3. Teclado

- [ ] Toda ação funciona por teclado.
- [ ] Ordem de tabulação segue a ordem visual.
- [ ] Foco visível e com contraste.
- [ ] Não há armadilha de foco.
- [ ] Dialog prende o foco e devolve ao elemento de origem.
- [ ] Menus e comboboxes funcionam com setas, Enter e Esc.
- [ ] Nenhuma ação depende de hover ou gesto complexo.

---

## 4. Formulários

- [ ] Todo campo tem `label` associado.
- [ ] Campos obrigatórios são indicados em texto, não só por cor/asterisco.
- [ ] Ajuda contextual usa `aria-describedby`.
- [ ] Erro é específico, associado ao campo e anunciado.
- [ ] Foco vai para o primeiro erro após submit.
- [ ] Dados preenchidos são preservados em erro.
- [ ] Autocomplete adequado (`email`, `tel`, `name`).
- [ ] Botão de submit tem estado de carregando e evita duplo envio.

---

## 5. Conteúdo e mídia

- [ ] Imagens informativas têm `alt`; decorativas têm `alt=""`.
- [ ] Ícones não são o único rótulo.
- [ ] Contraste mínimo 4,5:1 (texto normal) e 3:1 (texto grande).
- [ ] Não usar cor como único indicador.
- [ ] Texto redimensiona até 200% sem perda.
- [ ] Conteúdo não exige rolagem horizontal em 360 px.
- [ ] Vídeo/áudio tem legenda/transcrição (quando houver).
- [ ] Tabelas têm `th`, `scope` e legenda quando necessário.

---

## 6. Movimento e tempo

- [ ] `prefers-reduced-motion` desativa animações não essenciais.
- [ ] Nenhum carrossel automático sem controle.
- [ ] Nenhum timeout inesperado; se houver, é avisado e prorrogável.
- [ ] Sem flashing acima do limite.

---

## 7. Leitor de tela

- [ ] Título e heading são anunciados.
- [ ] Erros e sucessos usam `role="alert"`/`aria-live`.
- [ ] Mudança de página/rota é anunciada.
- [ ] Botões têm nome acessível.
- [ ] Estado de expansão/colapso é anunciado.
- [ ] Lista de resultados tem contagem e contexto.

---

## 8. Mobile e toque

- [ ] Alvos de toque ≥ 44×44 px.
- [ ] Espaçamento suficiente entre ações.
- [ ] Layout funciona em 360 px, 768 px e 1280 px.
- [ ] Zoom 200% sem perda.
- [ ] Orientação retrato/paisagem funcional.
- [ ] Teclado virtual não cobre campo focado.

---

## 9. Conteúdo e linguagem

- [ ] PT-BR claro e simples.
- [ ] Sem jargão desnecessário.
- [ ] Instruções antes de ações complexas.
- [ ] Erros dizem o que fazer.
- [ ] Sem urgência falsa ou pressão.
- [ ] O que é público/privado está explícito.

---

## 10. Privacidade na UI

- [ ] Nenhum e-mail/telefone/documento em página pública.
- [ ] Localização aproximada, não exata.
- [ ] Contato mascarado antes do aceite.
- [ ] Consentimento claro e revogável.
- [ ] Usuário entende o que será compartilhado.

---

## 11. Ferramentas e evidência

- axe automatizado nas páginas críticas;
- teste manual de teclado;
- teste com NVDA/VoiceOver;
- Lighthouse para performance/acessibilidade;
- screenshots/vídeo em 360 px;
- registro em `specs/web-experience/evidence/`.

## 12. Critério de bloqueio

Qualquer item não atendido em jornada crítica bloqueia o gate de QA/produto.
