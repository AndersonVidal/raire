# trust-reputation — mapa de aceite

**Status:** em revisão (testes propostos)  
**Spec:** `spec.md`  
**Dono:** Felipe  
**QA:** Felipe

## Convenções

- `U` = unitário, `I` = integração, `C` = contrato, `E2E` = ponta a ponta, `A11y` = acessibilidade, `Sec` = segurança, `Fair` = fairness.
- Dados sintéticos.
- Nenhum documento real.

## Matriz

| ID | Cenário | Requisito | Tipo | Teste | Status |
| --- | --- | --- | --- | --- | --- |
| AC-01 | avaliação exige conexão concluída | T-01 | U | `reputation` › "conexão accepted não é elegível" | pendente |
| AC-02 | uma avaliação por lado | T-02 | U/I | `reputation` › "segunda avaliação do mesmo autor é rejeitada" | pendente |
| AC-03 | autoavaliação bloqueada | T-02 | U/Sec | `reputation` › "autor não pode avaliar a si mesmo" | pendente |
| AC-04 | nota/tags/comentário | T-03 | U | `reputation` › "avaliação estruturada válida" | pendente |
| AC-05 | PII no comentário | T-04 | U/Sec | `reputation` › "telefone/e-mail no comentário é bloqueado" | pendente |
| AC-06 | janela de edição | T-05 | U | `reputation` › "edição após 24h retorna REVIEW_WINDOW_CLOSED" | pendente |
| AC-07 | remoção com auditoria | T-06 | I | `reputation` › "remoção exige motivo e grava auditoria" | pendente |
| AC-08 | resposta única | T-07 | U/I | `reputation` › "segunda resposta é rejeitada" | pendente |
| AC-09 | avaliação pública sem PII | T-08, T-37 | C/Sec | `reputation` › "review público não contém autor PII" | pendente |
| AC-10 | reputação contextual | T-09, T-10 | U | `reputation` › "calcula por categoria e exibe distribuição" | pendente |
| AC-11 | novo com prior neutro | T-11 | U | `reputation` › "novo profissional não recebe zero" | pendente |
| AC-12 | sinais complementares | T-12 | U | `reputation` › "sinais incluem resposta/conclusão/denúncia" | pendente |
| AC-13 | sinais não dominam | T-13 | U | `reputation` › "sinal comportamental tem peso limitado" | pendente |
| AC-14 | decaimento | T-14 | U | `reputation` › "denúncia antiga decai" | pendente |
| AC-15 | recalculável | T-15 | I | `reputation` › "snapshot pode ser recalculado" | pendente |
| AC-16 | verificação de contato | T-16, T-18 | I | `reputation` › "contato verificado muda status" | pendente |
| AC-17 | identidade manual | T-17, T-20 | I/Sec | `reputation` › "documento fica cifrado e restrito" | pendente |
| AC-18 | selo honesto | T-19 | U | `reputation` › "selo descreve o que foi verificado" | pendente |
| AC-19 | rejeição com motivo | T-21 | I | `reputation` › "rejeição permite reenvio" | pendente |
| AC-20 | revogação | T-22 | I | `reputation` › "fraude revoga selo com auditoria" | pendente |
| AC-21 | denúncia de perfil | T-23 | I | `moderation` › "report target profile" | pendente |
| AC-22 | denúncia de avaliação | T-23 | I | `moderation` › "report target review" | pendente |
| AC-23 | prioridade por risco | T-24 | U | `moderation` › "risco físico vira critical" | pendente |
| AC-24 | denunciante protegido | T-25, T-38 | Sec | `moderation` › "denunciado não vê denunciante" | pendente |
| AC-25 | SLA de fila | T-26 | I | `moderation` › "fila ordena por prioridade e tempo" | pendente |
| AC-26 | ações de moderação | T-27 | I | `moderation` › "aplica warning/removal/suspension" | pendente |
| AC-27 | auditoria | T-28 | I | `moderation` › "ação grava ator/motivo/timestamp" | pendente |
| AC-28 | recurso | T-29, T-30 | I | `moderation` › "afetado recorre e outro revisor decide" | pendente |
| AC-29 | comunicação | T-31 | I | `moderation` › "afetado é notificado sem PII de terceiros" | pendente |
| AC-30 | autoavaliação detectada | T-32 | U | `reputation` › "padrão de autoavaliação é sinalizado" | pendente |
| AC-31 | repetição suspeita | T-32, T-33 | I | `reputation` › "avaliações repetidas entram em revisão" | pendente |
| AC-32 | conta nova sinalizada | T-35 | U | `reputation` › "review de conta nova é sinalizada" | pendente |
| AC-33 | denúncia falsa | T-36 | I | `moderation` › "denúncias falsas repetidas são sinalizadas" | pendente |
| AC-34 | acessibilidade avaliação | T-39..T-41 | A11y | `web` › "avaliação por teclado e 360 px" | pendente |
| AC-35 | estados de UI | T-40 | E2E | `web` › "loading/vazio/erro/sucesso" | pendente |
| AC-36 | retenção | T-42 | I | `platform` › "job de expurgo respeita prazos" | pendente |
| AC-37 | crítica legítima preservada | T-06 | I | `moderation` › "review negativa legítima não é removida" | pendente |

## Cenários Gherkin

```gherkin
Cenário: avaliação após conexão concluída
  Dado uma conexão concluída
  Quando o cliente avalia o profissional
  Então a avaliação é publicada com nota, tags e contexto
  E não contém dados pessoais do autor

Cenário: avaliação antes da conclusão é bloqueada
  Dado uma conexão apenas aceita
  Quando tento avaliar
  Então recebo REVIEW_NOT_ELIGIBLE

Cenário: denúncia de risco físico
  Dado uma denúncia de ameaça
  Quando ela entra na fila
  Então recebe prioridade crítica
  E o denunciante não é exposto

Cenário: crítica legítima não é removida
  Dado uma avaliação negativa com crítica fundamentada
  Quando o profissional solicita remoção
  Então a moderação mantém a avaliação
  E registra a decisão com motivo

Cenário: reputação de novo profissional
  Dado um profissional sem avaliações
  Quando seu perfil é exibido
  Então aparece "Novo" e não nota zero
```

## Evidências

- Relatório de testes: pendente
- Segurança/PII: pendente
- Acessibilidade: pendente
- Fairness: pendente
- Moderação: pendente
