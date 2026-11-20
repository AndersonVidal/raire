# foundation-platform — mapa de aceite

**Status:** em revisão (testes propostos; caminhos podem mudar na implementação, mantendo a cobertura)  
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
| U | 12 | domínio, validação, catálogo |
| I | 16 | banco, sessão, perfil, admin |
| C | 3 | contrato público |
| E2E | 2 | onboarding e admin |
| A11y | 3 | teclado, 360 px, reduced motion |
| Sec | 5 | PII, authz, rate limit, MFA |
| Perf | 1 | health/readiness |
