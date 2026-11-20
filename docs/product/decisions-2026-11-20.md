# Decisões Aprovadas — 2026-11-20

**Status:** aprovado por humano em 2026-11-20  
**Aprovador:** responsável humano do projeto (instrução registrada no pedido de trabalho)  
**Escopo:** decisões de produto, arquitetura, segurança e operação para o MVP  
**Ressalva:** decisões jurídicas (Q-38–Q-43) valem para planejamento; a validação jurídica/DPIA continua obrigatória antes do piloto.

> Este documento registra a aprovação das recomendações apresentadas em `docs/status/2026-11-20-decision-draft.md` e dos defaults de `docs/product/open-questions.md`. A partir daqui, as questões Q-01..Q-50 passam a ser referência de decisão, não de hipótese.

---

## 1. Produto, mercado e piloto

| ID | Decisão aprovada |
| --- | --- |
| Q-01 | Região piloto real a definir antes de F6; testes usam a cidade sintética "Cidade Piloto". |
| Q-02 | Categorias iniciais: manutenção/reparos (presencial) e aulas/consultoria (remoto/híbrido). |
| Q-03 | Mínimo de 150 perfis completos para iniciar o piloto. |
| Q-04 | Moderação manual com fila pequena e SLA inicial de 48 h; risco físico em 4 h. |
| Q-05 | Modelo de monetização definido após o piloto; nunca ranking pago. |

## 2. Identidade, perfil e catálogo

| ID | Decisão aprovada |
| --- | --- |
| Q-06 | Verificação de identidade opcional no MVP; e-mail confirmado é obrigatório para publicar/criar pedido/interesse/avaliar. |
| Q-07 | Um perfil profissional por conta no MVP. |
| Q-08 | Área de atendimento por cidade/bairro + raio; polígono fica para depois. |
| Q-09 | Confirmação de contato por e-mail; telefone opcional. |
| Q-10 | Retenção de documentos de verificação: 90 dias após decisão; fora do primeiro ciclo. |
| Q-11 | Catálogo inicial curado por produto, revisado por pesquisa. |
| Q-12 | Portfólio aceita imagem opcional, com limite e validação; upload não é obrigatório. |
| Q-13 | Remoção de conteúdo de portfólio por operador, com auditoria. |

## 3. Descoberta e ranking

| ID | Decisão aprovada |
| --- | --- |
| Q-14 | Busca presencial exige cidade/bairro; remoto permite "qualquer lugar". |
| Q-15 | Raio padrão de 10 km; máximo de 50 km. |
| Q-16 | 10 resultados por página no mobile; 20 no desktop. |
| Q-17 | Distância pública em faixa aproximada ("até 2 km"), nunca coordenada exata. |
| Q-18 | Reputação resumida na busca, com contagem; detalhada no perfil. |
| Q-19 | Exploração de 20% dos slots da primeira página. |
| Q-20 | Cap de 2 impressões por sessão e 10 por dia por perfil. |
| Q-21 | Boost de novo profissional por 30 dias ou 5 conexões, com decaimento. |
| Q-22 | Piso de relevância no percentil 40 do baseline. |
| Q-23 | Guardrail de relevância: queda ≤ 2 p.p. no CTR e ≤ 5% no nDCG. |
| Q-24 | Política de ranking aprovada por produto + pesquisa + segurança. |
| Q-37 | Reputação resumida na busca; completa no perfil. |

## 4. Demandas, interesses e conexões

| ID | Decisão aprovada |
| --- | --- |
| Q-25 | Cliente pode aceitar até 3 profissionais por pedido. |
| Q-26 | Limite de 10 interesses por pedido e 20/dia por profissional. |
| Q-27 | Pedido expira em 7 dias. |
| Q-28 | Contato liberado pelo canal escolhido pelo titular; e-mail é o default. |
| Q-29 | Cliente pode publicar pedido com cidade/bairro, sem endereço exato. |
| Q-30 | Conclusão do serviço por confirmação de um lado, com contestação do outro. |
| Q-31 | Moderação pós-publicação, com fila por risco. |

## 5. Confiança, reputação e moderação

| ID | Decisão aprovada |
| --- | --- |
| Q-32 | Avaliação não exige confirmação mútua; um lado confirma e o outro pode contestar. |
| Q-33 | Janela de edição de avaliação: 24 h. |
| Q-34 | Verificação de identidade manual no MVP. |
| Q-35 | Denúncia não é anônima; identidade do denunciante é protegida. |
| Q-36 | SLA de moderação para risco físico: 4 h. |

## 6. Segurança, privacidade e conformidade

| ID | Decisão aprovada | Observação |
| --- | --- | --- |
| Q-38 | Bases legais propostas em `docs/security/privacy-and-lgpd.md`. | Validar com jurídico antes do piloto. |
| Q-39 | Prazos de retenção propostos no PRD. | Validar com jurídico antes do piloto. |
| Q-40 | Encarregado/DPO a designar antes do piloto. | Pendência humana. |
| Q-41 | Transferência internacional avaliada por fornecedor. | Antes do piloto. |
| Q-42 | Plano de incidente inicial em `docs/security/incident-response.md`. | Simular antes do piloto. |
| Q-43 | Pesquisa com consentimento separado e dados anonimizados. | Antes de qualquer experimento. |

## 7. Arquitetura e engenharia

| ID | Decisão aprovada |
| --- | --- |
| Q-44 | SQLite local/CI + PostgreSQL/PostGIS em produção, atrás de interfaces de repositório. |
| Q-45 | Web server-rendered no MVP, com JS progressivo. |
| Q-46 | Sessão por cookie `HttpOnly`, `Secure` em HTTPS, `SameSite=Lax` + CSRF. |
| Q-47 | MFA TOTP para admin/operador, com códigos de recuperação de uso único. |
| Q-48 | Outbox interno, sem broker externo no MVP. |
| Q-49 | Busca geoespacial atrás de interface; Haversine no local e PostGIS em produção. |
| Q-50 | Feature flags em tabela no banco, com default seguro (`off`). |

---

## 8. Consequências

- `foundation-platform` aprovada em 2026-11-20; implementação pode começar pelo `FND-01`.
- ADRs 0001–0006 passam a `aceito`.
- As demais specs continuam em revisão; suas questões já têm decisão, mas falta revisão humana da spec.
- A política de fairness tem parâmetros aprovados; ativação continua condicionada a simulação offline e revisão humana.
- DPIA, jurídico e DPO continuam pendentes antes do piloto.
- Nenhum código foi implementado por este registro.

## 9. Rastreabilidade

- Recomendações: `docs/status/2026-11-20-decision-draft.md`.
- Review da spec: `docs/status/2026-11-20-foundation-review-checklist.md`.
- Aprovações: `docs/process/approvals.md`.
- Spec aprovada: `specs/foundation-platform/spec.md`.
- ADRs: `docs/decisions/`.
