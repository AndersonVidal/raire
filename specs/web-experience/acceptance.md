# web-experience — mapa de aceite

**Status:** em revisão (testes propostos)  
**Spec:** `spec.md`  
**Dono:** Anderson  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `A11y` = acessibilidade, `Perf` = performance, `Sec` = segurança.
- Todo cenário Gherkin tem linha na matriz.
- PII em fixtures é sintética.
- Páginas críticas: home, cadastro, login, onboarding, busca, perfil, pedido, conexão, avaliação, denúncia, admin.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | tokens documentados | W-01 | U | `web` › "tokens têm cor, tipo, espaço, foco e movimento" | pendente |
| AC-02 | contraste AA | W-02 | A11y | `web` › "pares de cor passam 4,5:1 e 3:1" | pendente |
| AC-03 | alvo de toque | W-03 | A11y | `web` › "botões e links têm área ≥ 44×44 px" | pendente |
| AC-04 | componentes e estados | W-04, W-05 | U | `web` › "componentes têm estados default/focus/error" | pendente |
| AC-05 | foco visível | W-06, W-07 | A11y | `web` › "foco é visível e não depende só de cor" | pendente |
| AC-06 | landmarks e skip link | W-08 | A11y | `web` › "página tem skip link, header, main, footer" | pendente |
| AC-07 | navegação principal | W-09, W-11 | U | `web` › "nav tem no máximo 5 itens e ação de sair" | pendente |
| AC-08 | responsividade | W-10 | E2E | `web` › "360/768/1280 sem rolagem horizontal" | pendente |
| AC-09 | erro global | W-12 | A11y | `web` › "erro global tem role=alert" | pendente |
| AC-10 | título e idioma | W-13, W-14 | A11y | `web` › "title único e html lang=pt-BR" | pendente |
| AC-11 | cadastro | W-15 | E2E | `web` › "cadastro valida inline e preserva dados" | pendente |
| AC-12 | login | W-15 | E2E | `web` › "login mostra erro específico sem enumerar" | pendente |
| AC-13 | confirmação/reset | W-16 | E2E | `web` › "confirmação e reset mostram sucesso/erro" | pendente |
| AC-14 | onboarding em etapas | W-17 | E2E | `web` › "onboarding tem progresso e salvamento" | pendente |
| AC-15 | prévia do perfil | W-18 | E2E | `web` › "prévia mostra exatamente o público" | pendente |
| AC-16 | edição/pausa | W-17 | E2E | `web` › "editar e pausar perfil refletem no público" | pendente |
| AC-17 | formulário de busca | W-19 | E2E | `web` › "busca tem campo, categoria, localização e filtros" | pendente |
| AC-18 | resultados | W-20 | E2E | `web` › "card mostra distância, sinais, explicação e ação" | pendente |
| AC-19 | estado vazio | W-21 | E2E | `web` › "vazio sugere ajustar região/categoria" | pendente |
| AC-20 | perfil público | W-22 | E2E | `web` › "perfil mostra bio, área aproximada e sinais" | pendente |
| AC-21 | portfólio | W-22 | E2E | `web` › "portfólio publicado aparece no perfil" | pendente |
| AC-22 | transparência | W-20, W-31 | U | `web` › "explicação e página de transparência são coerentes" | pendente |
| AC-23 | formulário de pedido | W-23 | E2E | `web` › "pedido valida PII e campos obrigatórios" | pendente |
| AC-24 | lista/detalhe de pedido | W-24 | E2E | `web` › "cliente vê pedidos e interessados" | pendente |
| AC-25 | feed de oportunidades | W-25 | E2E | `web` › "profissional vê pedidos elegíveis" | pendente |
| AC-26 | interesse | W-26 | E2E | `web` › "interesse com mensagem curta" | pendente |
| AC-27 | contato mascarado | W-27 | Sec/E2E | `web` › "contato não aparece antes do aceite" | pendente |
| AC-28 | contato liberado | W-27 | E2E | `web` › "contato aparece após aceite" | pendente |
| AC-29 | estados de conexão | W-27 | E2E | `web` › "conexão mostra estado e próxima ação" | pendente |
| AC-30 | avaliação | W-28 | E2E | `web` › "avaliação com nota/tags/comentário" | pendente |
| AC-31 | denúncia | W-29 | E2E | `web` › "denúncia com motivo e descrição" | pendente |
| AC-32 | admin denúncias | W-30 | E2E | `web` › "fila de denúncias com prioridade" | pendente |
| AC-33 | admin auditoria | W-30 | E2E | `web` › "auditoria lista ações com ator e motivo" | pendente |
| AC-34 | admin fairness | W-30 | E2E | `web` › "fairness mostra métricas agregadas" | pendente |
| AC-35 | privacidade | W-32 | E2E | `web` › "privacidade explica dados e direitos" | pendente |
| AC-36 | consentimento/exportação/exclusão | W-32 | E2E | `web` › "consentimento, exportação e exclusão funcionam" | pendente |
| AC-37 | notificações | W-30 | E2E | `web` › "histórico de notificações" | pendente |
| AC-38 | teclado | W-33 | A11y | `web` › "jornadas críticas completas por teclado" | pendente |
| AC-39 | leitor de tela | W-34 | A11y | `web` › "título, erro e sucesso são anunciados" | pendente |
| AC-40 | labels e erros | W-35 | A11y | `web` › "todo input tem label e aria-describedby" | pendente |
| AC-41 | imagens | W-36, W-37 | A11y | `web` › "img tem alt; decorativa alt vazio" | pendente |
| AC-42 | zoom 200% | W-38 | A11y | `web` › "zoom 200% sem perda" | pendente |
| AC-43 | reduced motion | W-39 | A11y | `web` › "media query respeitada" | pendente |
| AC-44 | dialog | W-40 | A11y | `web` › "dialog prende e devolve foco" | pendente |
| AC-45 | copy PT-BR | W-42..W-47 | U | `web` › "copy sem jargão, urgência falsa ou promessa enganosa" | pendente |
| AC-46 | transparência de privacidade | W-47 | U | `web` › "perfil explica o que é público" | pendente |
| AC-47 | orçamento | W-48, W-49 | Perf | `web` › "HTML+CSS+JS ≤ 200 KB; imagens com dimensões" | pendente |
| AC-48 | Core Web Vitals | W-50..W-51 | Perf | `web` › "LCP ≤ 2,5 s; INP ≤ 200 ms; CLS ≤ 0,1" | pendente |
| AC-49 | no-JS | W-52 | E2E | `web` › "leitura e formulários básicos sem JS" | pendente |

## Cenários Gherkin

```gherkin
Cenário: busca acessível em 360 px
  Dado um celular de 360 px e navegação por teclado
  Quando informo categoria e região
  E envio a busca
  Então vejo resultados com nome, distância aproximada e explicação
  E não há rolagem horizontal
  E o foco vai para o título dos resultados

Cenário: contato protegido
  Dado um interesse ainda não aceito
  Quando abro a conexão
  Então o contato aparece como protegido
  E não vejo e-mail ou telefone

Cenário: erro de formulário preserva dados
  Dado que preencho o cadastro com e-mail inválido
  Quando envio
  Então vejo erro específico no campo
  E os demais campos permanecem preenchidos

Cenário: leitor de tela anuncia erro
  Dado um formulário com erro
  Quando o erro é exibido
  Então a região de alerta é anunciada
  E o foco vai para o primeiro campo com erro
```

## Evidências

- Relatório de testes: pendente
- Acessibilidade: pendente
- Performance: pendente
- Screenshots/vídeo: pendente
