# service-request — mapa de aceite

**Status:** em revisão (testes propostos)  
**Spec:** `spec.md`  
**Dono:** Anderson  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `A11y` = acessibilidade, `Sec` = segurança, `Perf` = performance.
- Dados sintéticos.
- Contato sempre mascarado em fixture pública.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | criar pedido válido | S-01 | U/I | `demand` › "cria pedido com categoria, região e urgência" | pendente |
| AC-02 | orçamento opcional | S-02 | U | `demand` › "orçamento é intervalo opcional" | pendente |
| AC-03 | status inicial | S-03 | U | `demand` › "pedido nasce open" | pendente |
| AC-04 | editar pedido aberto | S-04 | I | `demand` › "edita enquanto open e registra histórico" | pendente |
| AC-05 | fechar/cancelar | S-05 | I | `demand` › "pedido fechado não recebe interesse" | pendente |
| AC-06 | expiração | S-06 | I | `demand` › "pedido expira após 7 dias" | pendente |
| AC-07 | limite de pedidos | S-07 | U | `demand` › "6º pedido ativo é rejeitado" | pendente |
| AC-08 | PII na descrição | S-08 | U/Sec | `demand` › "telefone/e-mail/endereço bloqueiam publicação" | pendente |
| AC-09 | pedido sem PII pública | S-09 | C/Sec | `demand` › "resposta pública não tem endereço/contato" | pendente |
| AC-10 | conteúdo sinalizado | S-10 | I | `demand` › "pedido sinalizado entra na fila" | pendente |
| AC-11 | feed elegível | S-11 | I | `demand` › "profissional vê pedidos da sua categoria/área" | pendente |
| AC-12 | feed paginado | S-12 | I | `demand` › "cursor avança sem duplicar" | pendente |
| AC-13 | filtros do feed | S-13 | I | `demand` › "filtra por categoria/urgência/distância" | pendente |
| AC-14 | identidade do cliente | S-14 | C/Sec | `demand` › "feed não mostra identidade completa antes do aceite" | pendente |
| AC-15 | perfil inelegível | S-15 | U | `demand` › "perfil draft/suspenso não acessa feed" | pendente |
| AC-16 | enviar interesse | S-16 | U/I | `demand` › "interesse com mensagem curta" | pendente |
| AC-17 | interesse único | S-17 | U | `demand` › "segundo interesse retorna ALREADY_INTERESTED" | pendente |
| AC-18 | limite de interesses | S-18 | U | `demand` › "11º interesse no pedido é rejeitado" | pendente |
| AC-19 | lista de interessados | S-19 | I | `demand` › "cliente vê interessados com perfil resumido" | pendente |
| AC-20 | aceitar múltiplos | S-20 | I | `demand` › "aceita até 3 interesses" | pendente |
| AC-21 | expiração de interesse | S-21 | I | `demand` › "interesses pendentes expiram com o pedido" | pendente |
| AC-22 | retirar/recusar | S-22 | I | `demand` › "profissional retira e cliente é notificado" | pendente |
| AC-23 | aceite cria conexão | S-23 | I | `demand` › "aceite cria connection accepted" | pendente |
| AC-24 | contato protegido | S-24 | Sec | `demand` › "contato não aparece antes do aceite" | pendente |
| AC-25 | contato liberado | S-24, S-38 | I/Sec | `demand` › "após aceite, contato do canal escolhido aparece" | pendente |
| AC-26 | estados da conexão | S-25 | U | `demand` › "transições válidas de connection" | pendente |
| AC-27 | conclusão/contestação | S-26 | I | `demand` › "um lado conclui, outro contesta" | pendente |
| AC-28 | bloqueio/denúncia | S-27 | I | `demand` › "bloqueio impede novo interesse" | pendente |
| AC-29 | auditoria de transição | S-28 | I | `demand` › "transição grava ator e timestamp" | pendente |
| AC-30 | notificação de pedido | S-29 | I | `demand` › "novo pedido notifica profissional elegível" | pendente |
| AC-31 | preferências | S-30 | I | `demand` › "preferência desliga notificação não essencial" | pendente |
| AC-32 | notificação sem PII | S-31 | Sec | `demand` › "notificação não contém endereço/contato" | pendente |
| AC-33 | falha de notificação | S-32 | I | `demand` › "falha de notificação não bloqueia transição" | pendente |
| AC-34 | rate limit | S-33 | Sec | `demand` › "limite de pedido/interesse/denúncia" | pendente |
| AC-35 | duplicidade | S-34 | U | `demand` › "pedido repetido é sinalizado" | pendente |
| AC-36 | abuso | S-35 | I | `demand` › "sinal de abuso vai para moderação" | pendente |
| AC-37 | usuário bloqueado | S-36 | I | `demand` › "bloqueado não cria pedido nem interesse" | pendente |
| AC-38 | localização precisa | S-37 | Sec | `demand` › "localização precisa não é retornada" | pendente |
| AC-39 | consentimento | S-38 | I | `demand` › "consentimento de contato tem timestamp/versão" | pendente |
| AC-40 | acessibilidade | S-39..S-41 | A11y | `web` › "fluxo por teclado e 360 px" | pendente |

## Cenários Gherkin

```gherkin
Cenário: cliente publica pedido e profissional demonstra interesse
  Dado que sou cliente com e-mail confirmado
  Quando publico um pedido de eletricista no Centro
  Então o pedido fica disponível para profissionais elegíveis
  E um profissional da categoria vê a oportunidade
  Quando o profissional envia interesse com mensagem
  Então o cliente vê o interessado com perfil resumido

Cenário: contato só é liberado após aceite
  Dado um interesse pendente
  Quando o cliente abre o pedido
  Então o contato do profissional está protegido
  Quando o cliente aceita o interesse
  Então a conexão é criada
  E o contato do canal escolhido é liberado para ambos

Cenário: pedido expirado não recebe interesse
  Dado um pedido com data de expiração no passado
  Quando um profissional tenta enviar interesse
  Então recebo REQUEST_EXPIRED
  E o pedido não muda de estado

Cenário: PII na descrição é bloqueada
  Dado que escrevo meu telefone na descrição
  Quando tento publicar
  Então recebo PII_NOT_ALLOWED
  E o pedido não é publicado
```

## Evidências

- Relatório de testes: pendente
- Segurança (contato/PII): pendente
- Acessibilidade: pendente
- E2E: pendente
