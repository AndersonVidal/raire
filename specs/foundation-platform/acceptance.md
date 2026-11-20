# foundation-platform — mapa de aceite

**Status:** aprovado em 2026-11-20 (AC-40..AC-60 adicionados na revisão 1.1)  
**Spec:** `spec.md`  
**Dono:** Victor  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `A11y` = acessibilidade, `Perf` = performance, `Sec` = segurança.
- Todo cenário Gherkin tem linha na matriz.
- Testes de integração usam banco real (SQLite local/CI; Postgres em staging).
- PII em fixtures é sintética.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | cadastro válido cria conta e consentimentos | R-01, R-15 | U/I | `identity` › "signup persiste conta e versão dos termos" | pendente |
| AC-02 | e-mail duplicado não enumera | R-01, R-05 | I/Sec | `identity` › "e-mail existente responde de forma genérica" | pendente |
| AC-03 | confirmação de e-mail com token de uso único | R-02, R-03 | U/I | `identity` › "token expira e não pode ser reutilizado" | pendente |
| AC-04 | login rotaciona sessão | R-04, R-11 | I | `identity` › "login cria nova sessão e invalida anterior" | pendente |
| AC-05 | logout revoga sessão | R-12 | I | `identity` › "logout revoga apenas a sessão atual" | pendente |
| AC-06 | reset não revela conta | R-05 | I/Sec | `identity` › "reset responde 202 para e-mail inexistente" | pendente |
| AC-07 | rate limit de login | R-08 | I/Sec | `identity` › "6ª tentativa inválida bloqueia temporariamente" | pendente |
| AC-08 | admin exige MFA | R-09, R-51 | I/Sec | `identity` › "login admin sem código TOTP é rejeitado" | pendente |
| AC-09 | onboarding presencial publica perfil | R-20..R-28 | U/I | `directory` › "publica perfil presencial com área e disponibilidade" | pendente |
| AC-10 | onboarding remoto sem área | R-21, R-22 | U | `directory` › "publica perfil remoto sem área" | pendente |
| AC-11 | híbrido exige área | R-22 | U | `directory` › "híbrido sem área retorna INVALID_SERVICE_AREA" | pendente |
| AC-12 | raio inválido é rejeitado | R-23 | U | `directory` › "raio abaixo do mínimo ou acima do máximo é rejeitado" | pendente |
| AC-13 | perfil incompleto não publica | R-28 | U | `directory` › "publicação sem categoria retorna PROFILE_INCOMPLETE" | pendente |
| AC-14 | draft não aparece na busca | R-26, R-27 | I | `directory` › "perfil draft não entra no read model público" | pendente |
| AC-15 | prévia não expõe PII | R-29, R-54 | C/Sec | `directory` › "preview não contém e-mail, telefone, documento ou coordenada" | pendente |
| AC-16 | categoria desativada remove perfil | R-33, R-42 | I | `directory` › "categoria inativa remove perfil da busca" | pendente |
| AC-17 | catálogo ativo/inativo | R-40..R-44 | U/I | `catalog` › "categoria inativa não aparece no onboarding" | pendente |
| AC-18 | consentimento versionado | R-15, R-52 | I | `identity` › "consentimento registra versão e timestamp" | pendente |
| AC-19 | revogação de consentimento | R-16 | I | `identity` › "revogação de pesquisa impede evento analítico" | pendente |
| AC-20 | exportação de dados | R-17 | I | `identity` › "export retorna dados do próprio usuário" | pendente |
| AC-21 | exclusão anonimiza | R-18 | I | `identity` › "delete anonimiza perfil e preserva auditoria" | pendente |
| AC-22 | solicitação de direitos | R-19 | I | `identity` › "account_request é registrada e resolvida" | pendente |
| AC-23 | PII fora de log | R-55 | U/Sec | `platform` › "redaction remove e-mail, telefone, documento e coordenada" | pendente |
| AC-24 | portfólio publica item | R-34..R-37 | I | `directory` › "item publicado aparece no perfil público" | pendente |
| AC-25 | portfólio não é verificado | R-38 | U | `directory` › "item não recebe selo de verificação" | pendente |
| AC-26 | portfólio denunciável | R-39 | I | `moderation` › "report aceita target portfolio" | pendente |
| AC-27 | portfólio sem PII | R-54 | C/Sec | `directory` › "portfólio público não contém PII" | pendente |
| AC-28 | catálogo com sinônimo | R-45 | U | `catalog` › "eletricista encontra instalação elétrica" | pendente |
| AC-29 | seed idempotente | R-43 | I | `catalog` › "seed aplicado duas vezes mantém contagem" | pendente |
| AC-30 | operador gerencia categoria | R-44 | I | `catalog` › "desativação grava auditoria" | pendente |
| AC-31 | admin suspende perfil | R-46..R-48 | I/Sec | `moderation` › "suspender remove da busca e revoga sessão" | pendente |
| AC-32 | auditoria admin | R-49 | I | `moderation` › "ação grava ator, alvo, motivo e correlação" | pendente |
| AC-33 | operador não acessa admin sem MFA | R-51, R-09 | Sec | `identity` › "operator sem MFA recebe FORBIDDEN" | pendente |
| AC-34 | onboarding acessível | R-58..R-63 | A11y | `web` › "teclado, foco, rótulos e 360 px" | pendente |
| AC-35 | estados de UI | R-61 | E2E | `web` › "loading, vazio e erro visíveis" | pendente |
| AC-36 | reduced motion | R-62 | A11y | `web` › "CSS respeita prefers-reduced-motion" | pendente |
| AC-37 | health/readiness | R-67 | I | `platform` › "health 200 e readiness falha sem banco" | pendente |
| AC-38 | logs estruturados | R-64 | U | `platform` › "log contém correlation_id e não contém PII" | pendente |
| AC-39 | feature flag default seguro | R-68 | I | `platform` › "flag desconhecida resolve off" | pendente |
| AC-40 | e-mail não confirmado não publica | R-02 | U/I | `identity` › "publicar sem e-mail confirmado retorna EMAIL_NOT_CONFIRMED" | pendente |
| AC-41 | mesma conta cliente e profissional | R-07 | I | `directory` › "conta com perfil profissional também atua como cliente" | pendente |
| AC-42 | CSRF ausente/inválido | R-10 | Sec | `platform` › "POST sem CSRF retorna CSRF_INVALID" | pendente |
| AC-43 | troca de senha revoga outras sessões | R-12 | I | `identity` › "password change mantém sessão atual e revoga as demais" | pendente |
| AC-44 | preferência de contato e mascaramento | R-13, R-14, R-14.1 | U/I | `identity` › "preferência do perfil sobrescreve a da conta; telefone não verificado não aparece" | pendente |
| AC-45 | selo de verificação honesto | R-31 | U/Sec | `directory` › "selo mostra apenas o que foi verificado" | pendente |
| AC-46 | pausar/reativar perfil | R-32 | I | `directory` › "pause remove da busca; resume recoloca" | pendente |
| AC-47 | reordenar/despublicar portfólio | R-36 | I | `directory` › "reorder e unpublish refletem no público" | pendente |
| AC-48 | categoria/flag auditada | R-44, R-50 | I | `catalog` › "mudança de categoria/flag grava audit_event e métrica" | pendente |
| AC-49 | termos/privacidade exibidos | R-57 | E2E | `web` › "cadastro mostra versão dos termos e privacidade" | pendente |
| AC-50 | métricas de cadastro/publicação | R-65 | I | `platform` › "métricas de funil são registradas sem PII" | pendente |
| AC-51 | rate limit de cadastro/reset/confirmação | R-08 | Sec | `identity` › "limites de signup/reset/resend são aplicados" | pendente |
| AC-52 | idempotência de signup/confirmação/reset | RNF-05 | I | `identity` › "retry com mesma chave não duplica conta/token" | pendente |
| AC-53 | paginação/filtros admin | R-46, R-50 | I | `moderation` › "admin lista com cursor e filtros" | pendente |
| AC-54 | localização precisa cifrada e expurgada | R-24, R-53 | Sec/I | `directory` › "coordenada é cifrada e expurgada após 90 dias" | pendente |
| AC-55 | disponibilidade válida | R-25 | U | `directory` › "janela sobreposta/inválida é rejeitada" | pendente |
| AC-56 | imagem de portfólio validada | R-34 | U/I | `directory` › "imagem acima do limite/tipo inválido retorna INVALID_IMAGE" | pendente |
| AC-57 | sucesso e sessão expirada na UI | R-61 | E2E | `web` › "estado de sucesso e sessão expirada são explícitos" | pendente |
| AC-58 | solicitações de direitos no admin | R-19 | I | `identity` › "operador lista e resolve account_request" | pendente |
| AC-59 | admin de categorias/flags | R-44, R-50 | I | `catalog` › "operador cria/desativa categoria e altera flag" | pendente |
| AC-60 | copy PT-BR | R-63 | U | `web` › "copy revisada, sem jargão ou promessa enganosa" | pendente |

## Cenários Gherkin

```gherkin
Cenário: profissional publica perfil presencial
  Dado que estou autenticado como profissional
  E confirmei meu e-mail
  Quando informo nome, bio, categoria, área de atendimento e disponibilidade
  E peço para publicar
  Então meu perfil fica com status "published"
  E aparece na prévia pública sem e-mail, telefone ou coordenada precisa

Cenário: perfil híbrido sem área é rejeitado
  Dado que escolhi o tipo de atendimento "hybrid"
  Quando tento publicar sem informar área de atendimento
  Então recebo o erro "INVALID_SERVICE_AREA"
  E meu perfil continua em rascunho

Cenário: categoria desativada sai da busca
  Dado um perfil publicado na categoria "eletricista"
  Quando a operação desativa essa categoria
  Então o perfil deixa de aparecer na busca
  E o profissional recebe aviso para escolher outra categoria

Cenário: admin suspende perfil
  Dado um perfil publicado
  Quando um admin com MFA suspende o perfil com motivo
  Então o perfil sai da busca imediatamente
  E a ação fica registrada em audit_events
  E a sessão do profissional é revogada no próximo request

Cenário: API pública não vaza PII
  Dado um profissional com telefone, e-mail e localização precisa cadastrados
  Quando qualquer visitante consulta o perfil público
  Então a resposta contém apenas nome de exibição, bio, categorias, área aproximada e sinais públicos
  E não contém telefone, e-mail, documento ou coordenada precisa

Cenário: onboarding acessível em 360 px
  Dado um celular de 360 px e navegação por teclado
  Quando percorro cadastro, confirmação e publicação de perfil
  Então todos os campos têm rótulo, foco visível e mensagem de erro específica
  E não há rolagem horizontal

Cenário: exclusão de conta
  Dado que solicito exclusão da minha conta
  Quando a solicitação é processada
  Então meus dados públicos são anonimizados
  E meus dados de contato são removidos
  E a auditoria é preservada

Cenário: e-mail não confirmado não publica
  Dado que criei conta mas não confirmei o e-mail
  Quando tento publicar um perfil
  Então recebo EMAIL_NOT_CONFIRMED
  E o perfil continua em rascunho

Cenário: e-mail duplicado não enumera
  Dado que já existe uma conta com o e-mail informado
  Quando tento criar outra conta com o mesmo e-mail
  Então recebo a mesma resposta de sucesso genérica
  E nenhum dado da conta existente é revelado

Cenário: CSRF inválido é rejeitado
  Dado que estou autenticado
  Quando envio um POST sem token CSRF válido
  Então recebo CSRF_INVALID
  E nenhuma alteração é persistida

Cenário: troca de senha revoga outras sessões
  Dado que estou autenticado em dois dispositivos
  Quando troco a senha informando a senha atual
  Então a sessão atual permanece
  E a outra sessão é revogada

Cenário: preferência de contato
  Dado que minha conta prefere e-mail
  E meu perfil prefere telefone
  Quando alguém consulta meu perfil público
  Então o contato permanece mascarado até consentimento
  E o canal do perfil é usado quando liberado

Cenário: localização precisa é protegida
  Dado que informei uma coordenada precisa
  Quando consulto meu perfil público
  Então vejo apenas cidade/bairro e faixa de distância
  E a coordenada não aparece na resposta, log ou evento

Cenário: disponibilidade inválida é rejeitada
  Dado que informo uma janela com fim antes do início
  Quando tento salvar o perfil
  Então recebo INVALID_AVAILABILITY
  E a disponibilidade anterior é preservada

Cenário: admin altera categoria e flag com auditoria
  Dado que sou operador com MFA
  Quando desativo uma categoria e desligo uma flag
  Então cada ação grava audit_event e métrica
  E o efeito é imediato na busca/feature

Cenário: direitos do titular
  Dado que solicito exportação/exclusão
  Quando o operador resolve a solicitação
  Então a solicitação é registrada com data e resultado
  E o titular recebe a resposta aplicável
```

## Evidências

- Relatório de testes: `evidence/` (pendente)
- Acessibilidade: pendente
- Segurança: pendente
- Performance: pendente
- Screenshots/vídeo: pendente

## Cobertura por tipo

| Tipo | Quantidade | Observação |
| --- | --- | --- |
| U | 15 | domínio, validação, catálogo, disponibilidade, contato |
| I | 25 | banco, sessão, perfil, catálogo, admin, direitos |
| C | 3 | contrato público |
| E2E | 5 | onboarding, admin, estados, termos |
| A11y | 3 | teclado, 360 px, reduced motion |
| Sec | 8 | PII, authz, CSRF, rate limit, MFA, criptografia |
| Perf | 1 | health/readiness |
