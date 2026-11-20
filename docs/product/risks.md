# Riscos e Mitigações — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0, seção 12  
**Uso:** acompanhamento mensal no piloto; bloqueios de release.

---

## 1. Como classificar

| Dimensão | Escala |
| --- | --- |
| Probabilidade | baixa / média / alta |
| Impacto | baixo / médio / alto / crítico |
| Exposição | probabilidade × impacto |
| Dono | pessoa responsável pelo acompanhamento |
| Estado | aberto / mitigado / monitorado / encerrado |

**Crítico** = risco físico, vazamento de PII, discriminação, fraude generalizada ou violação legal.

---

## 2. Riscos estratégicos e de mercado

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-01 | Cold start / baixa liquidez | alta | alto | piloto denso, curadoria de oferta, mural, parcerias locais | Anderson |
| R-02 | Profissionais não veem valor | média | alto | onboarding assistido, oportunidades reais, suporte | Anderson |
| R-03 | Clientes não confiam fora da bolha | média | alto | sinais honestos, avaliação contextual, transparência | Felipe |
| R-04 | Dependência de uma região | média | médio | critérios de expansão por densidade e evidência | Victor |
| R-05 | Concorrência generalista | média | médio | foco local, fairness, acessibilidade | Fundador |

---

## 3. Riscos financeiros e de sustentabilidade

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-06 | Custo de moderação inviável | média | alto | fila priorizada, critérios claros, triagem assistida | Felipe |
| R-07 | Pressão por monetização antes da validação | alta | alto | princípio sem pay-to-win; decisão após piloto | Fundador |
| R-08 | Custo de infraestrutura cresce sem demanda | baixa | médio | arquitetura simples, orçamento, gatilhos | Victor |
| R-09 | Modelo de negócio não definido | alta | médio | pesquisa pós-piloto; não bloquear MVP | Fundador |

---

## 4. Riscos de experiência e inclusão

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-10 | Onboarding longo afasta P3 | alta | alto | formulário curto, salvamento progressivo, teste com usuários | Anderson |
| R-11 | Baixa literacia digital exclui público | alta | alto | linguagem simples, onboarding assistido, canal de suporte | Anderson |
| R-12 | Interface não funciona em 360 px | média | alto | mobile-first, teste em dispositivo real | Anderson |
| R-13 | Acessibilidade tratada como checklist | média | alto | teste com teclado/leitor de tela, auditoria axe | Felipe |

---

## 5. Riscos de confiança, segurança e integridade

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-14 | Perfil falso / fraude | alta | alto | verificação manual, sinais, denúncia, moderação | Felipe |
| R-15 | Assédio via contato | média | crítico | contato mascarado, consentimento, bloqueio, rate limit | Felipe |
| R-16 | Vazamento de contato | média | crítico | mascaramento até aceite, testes de contrato, revisão | Felipe |
| R-17 | Manipulação de avaliações | alta | alto | elegibilidade por conexão, detecção, moderação | Felipe |
| R-18 | Spam de pedidos/interesses | alta | médio | limites, validação, bloqueio, triagem | Anderson |
| R-19 | Account takeover | média | crítico | senha forte, rate limit, sessão revogável, MFA admin | Felipe |
| R-20 | Abuso de admin | baixa | crítico | MFA, menor privilégio, auditoria, recurso | Felipe |
| R-21 | Dependência vulnerável | média | alto | scan, atualização, SBOM | Felipe |
| R-22 | Injeção/XSS/CSRF | média | alto | validação, encoding, CSP, CSRF | Felipe |

---

## 6. Riscos de equidade e integridade da descoberta

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-23 | Ranking concentra visibilidade | alta | alto | caps, exploração, métricas, revisão | Victor |
| R-24 | Fairness reduz relevância | média | alto | piso de relevância, guardrails, experimento | Victor |
| R-25 | Gaming de exposição | média | alto | limites, detecção de anomalia, logs | Victor |
| R-26 | Viés indireto via proxy | média | crítico | proibir atributos sensíveis, auditoria, revisão | Felipe + Victor |
| R-27 | Explicação enganosa | média | alto | derivar de sinais reais, revisão de copy | Anderson |
| R-28 | Feedback loop de popularidade | média | alto | exploração, holdout, monitoramento | Victor |

---

## 7. Riscos de privacidade e proteção de dados

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-29 | Vazamento de localização precisa | baixa | crítico | localização aproximada, criptografia, acesso restrito | Felipe |
| R-30 | PII em log/evento | média | crítico | redaction, testes, revisão | Felipe |
| R-31 | Retenção além do necessário | média | alto | prazos, jobs de expurgo, auditoria | Felipe |
| R-32 | Base legal inadequada | média | crítico | revisão jurídica, DPIA | Jurídico |
| R-33 | Direitos do titular não atendidos | média | alto | fluxo de solicitação, prazo, registro | Felipe |
| R-34 | Uso secundário sem consentimento | média | alto | consentimento granular, governança | Felipe |

---

## 8. Riscos operacionais

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-35 | Fila de denúncias sem prioridade | média | alto | prioridade por risco, SLA 48 h | Felipe |
| R-36 | Decisão de moderação inconsistente | média | alto | critérios públicos, dupla revisão, auditoria | Felipe |
| R-37 | Falha de notificação | média | médio | outbox, retry, fallback in-app | Anderson |
| R-38 | Migração sem rollback | baixa | alto | expand/contract, backup, runbook | Victor |
| R-39 | Perda de dados | baixa | crítico | backup, restore testado, RPO/RTO | Felipe |

---

## 9. Riscos tecnológicos e de engenharia

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-40 | Acoplamento entre contextos | média | médio | regras de dependência, review, teste de arquitetura | Victor |
| R-41 | Ranking lento | média | alto | índices, paginação, benchmark, fallback | Victor |
| R-42 | Complexidade de eventos | média | médio | outbox simples, contratos versionados | Victor |
| R-43 | SQLite local ≠ Postgres produção | média | médio | repositórios isolados, testes de contrato, migration review | Victor |
| R-44 | Dependência de fornecedor | baixa | médio | containers, formatos abertos | Victor |

---

## 10. Riscos jurídicos e regulatórios

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-45 | LGPD/DPIA incompleta | média | crítico | revisão jurídica antes do piloto | Jurídico |
| R-46 | Responsabilidade por serviço prestado | média | alto | termos claros; plataforma não intermedeia | Jurídico |
| R-47 | Conteúdo de terceiros no portfólio | média | médio | consentimento, denúncia, remoção | Felipe |
| R-48 | Discriminação algorítmica | baixa | crítico | fairness auditável, sem atributos sensíveis | Victor + Felipe |

---

## 11. Riscos de pesquisa e validação

| ID | Risco | Prob. | Impacto | Mitigação | Dono |
| --- | --- | --- | --- | --- | --- |
| R-49 | Métrica de clique mascara baixa conexão | alta | alto | métrica central de conexão | Anderson |
| R-50 | Amostra pequena gera conclusão falsa | alta | alto | registrar limitações, não generalizar | Pesquisa |
| R-51 | Viés de recrutamento | média | alto | diversificar canais e perfis | Pesquisa |
| R-52 | Experimento sem guardrail | média | crítico | protocolo obrigatório | Felipe |
| R-53 | Evidência não registrada | média | alto | template e revisão | Todos |

---

## 12. Condições impeditivas para o piloto

O piloto **não começa** se:

- houver incidente crítico de privacidade sem correção;
- o ranking usar atributo sensível ou proxy não avaliado;
- não houver moderação operante;
- a DPIA não estiver concluída;
- não houver plano de resposta a incidentes;
- a acessibilidade das jornadas críticas não estiver validada;
- não houver aprovação humana registrada.

---

## 13. Acompanhamento

- Revisão mensal do registro de riscos.
- Toda mudança de exposição exige atualização de dono e mitigação.
- Riscos críticos viram bloqueio de release.
- Riscos encerrados mantêm histórico e evidência.
- Aprendizados vão para `docs/experiments` e `docs/decisions`.
