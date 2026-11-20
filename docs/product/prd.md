# PRD — RAIRE: Descoberta Local Justa

**Status:** aprovado em 2026-11-20 (instrução humana registrada)  
**Versão:** 1.0  
**Data da revisão:** 2026-11-20  
**Versão anterior:** 0.1 (`docs/product/prd-v0.1.md`)  
**Fase:** Fundação / MVP conceitual / preparação de implementação  
**Responsável de produto:** Anderson  
**Arquitetura:** Victor  
**Qualidade, segurança e privacidade:** Felipe  
**Aprovadores:** fundador, produto, arquitetura, segurança/privacidade e jurídico

> Este PRD organiza o que o RAIRE pretende validar e construir. Ele **não autoriza implementação** por si só. Toda feature nasce de uma spec em `specs/`, com status `aprovada` apenas após revisão humana. Hipóteses continuam sendo hipóteses até existir evidência.
>
> **Estado desta revisão:** organização e planejamento concluídos; nenhum código de produto implementado. A preparação foi autorizada pelo waiver `docs/process/waivers/2026-11-20-mvp-planning.md`, que **não autoriza código, merge, piloto ou produção**.

---

## 0. Como ler esta revisão

Esta versão consolida o PRD 0.1 e o transforma em um documento acionável para três frentes de trabalho: backend/ranking, web/produto e QA/segurança.

| Documento | Para quê |
| --- | --- |
| `docs/product/prd.md` (este) | decisão de produto, escopo, requisitos e critérios de aceite |
| `docs/product/requirements-catalog.md` | catálogo detalhado RF-01 a RF-20, com sub-requisitos |
| `docs/product/personas-and-journeys.md` | personas, jornadas e JTBD |
| `docs/product/metrics-and-experiments.md` | métricas, metas e experimentos |
| `docs/product/risks.md` | riscos, impacto, mitigação e dono |
| `docs/product/roadmap.md` | fases, marcos e critérios de saída |
| `docs/product/open-questions.md` | decisões pendentes, dono e prazo |
| `specs/README.md` | quadro de specs e dependências |
| `docs/team/task-board.md` | todas as tarefas, donos e status |
| `docs/process/pipeline.md` | pipeline de trabalho e gates |

### 0.1 O que mudou em relação à v0.1

1. **Escopo do MVP priorizado.** O MVP foi separado em duas camadas: o **MVP de validação** (piloto) e a **fatia técnica implementável** (primeiro ciclo de engenharia). Nem todo requisito do PRD entra no primeiro ciclo.
2. **Requisitos rastreáveis.** RF-01 a RF-20 foram mantidos e ligados a specs, tickets e critérios de aceite. O detalhamento de sub-requisitos foi movido para o catálogo, evitando duplicação.
3. **Critérios de aceite explícitos.** O PRD passa a ter critérios de maturidade e uma matriz de rastreabilidade; os critérios de feature ficam nas specs.
4. **Fairness com desenho operacional.** A camada de fairness tem política versionada, coortes observáveis, métricas, guardrails e fallback definidos.
5. **Privacidade e LGPD com decisões iniciais.** Classificação de dados, retenção, consentimento, direitos e contato mascarado foram consolidados.
6. **Pipeline e donos.** O trabalho foi dividido entre Victor (arquitetura/backend/ranking), Anderson (web/produto/acessibilidade) e Felipe (QA/segurança/privacidade/DevOps), com gates e handoffs.
7. **Decisões pendentes com default.** Perguntas abertas têm dono, prazo e um default seguro para não bloquear o planejamento.
8. **Documentação enxuta.** O PRD aponta para documentos especializados em vez de repetir tudo.

### 0.2 Decisões adotadas para o planejamento

Estas decisões são **propostas** e precisam de aprovação humana antes da implementação:

| ID | Decisão | Default adotado | Dono |
| --- | --- | --- | --- |
| D-01 | Região piloto | Cidade sintética "Cidade Piloto" nos testes; região real a definir antes do piloto | Produto |
| D-02 | Categorias iniciais | 2 categorias: manutenção/reparos (presencial) e aulas/consultoria (remoto/híbrido) | Produto |
| D-03 | Verificação de identidade | Opcional; e-mail confirmado é obrigatório para publicar/criar pedido; selo de identidade fica fora do primeiro ciclo | Produto + Segurança |
| D-04 | Liberação de contato | Somente após aceite mútuo; canal escolhido pelo titular | Segurança + Produto |
| D-05 | Elegibilidade de avaliação | Conexão marcada como concluída por pelo menos um lado; contestação entra em moderação | Produto |
| D-06 | Monetização | Fora do MVP; nenhuma posição paga | Fundador |
| D-07 | Moderação | Fila manual com SLA inicial de 48 h; automação só como triagem | Trust & Safety |
| D-08 | Retenção | Proposta da seção 11; validar com jurídico antes do piloto | Jurídico/Privacidade |
| D-09 | Stack do MVP técnico | Go modular + SQLite local/CI + PostgreSQL/PostGIS como alvo de produção + web server-rendered | Arquitetura |
| D-10 | Governança de ranking | Produto + pesquisa + segurança aprovam mudanças; política versionada e reversível | Produto + Segurança |

---

## 1. Sumário executivo

O RAIRE é uma plataforma de descoberta e conexão entre profissionais autônomos locais e pessoas que buscam serviços. O problema central não é apenas "faltar um marketplace": é a combinação de baixa exposição digital, dependência de indicação e alcance social, dificuldade de comparar confiança e concentração de visibilidade em quem já é popular.

O MVP propõe uma experiência simples e móvel:

1. O profissional cria um perfil, declara serviços, área de atendimento e disponibilidade.
2. O cliente busca por serviço e localização ou publica uma necessidade no "mural de guilda".
3. O RAIRE apresenta profissionais relevantes com explicação de ranking e mecanismos de distribuição justa de visibilidade.
4. O cliente demonstra interesse, o profissional aceita e o contato é liberado com consentimento.
5. Depois da conexão, as duas partes podem registrar o desfecho e avaliar a experiência, construindo reputação contextual.

O diferencial do RAIRE é tratar **fairness, explicabilidade e confiança como requisitos de produto desde o primeiro release**, não como otimizações posteriores. O MVP não terá ranking pago, não usará atributos sensíveis e não prometerá verificações que não consegue entregar.

**Resultado esperado do MVP:** validar, em uma região piloto e em poucas categorias, se é possível melhorar a descoberta de profissionais com baixa exposição sem piorar a relevância percebida pelo cliente e sem gerar riscos de segurança ou privacidade inaceitáveis.

### 1.1 Métrica central

**Conexões consentidas de qualidade por semana na região piloto**, acompanhadas de:

- relevância percebida pelo cliente;
- distribuição de exposição entre profissionais comparáveis;
- ausência de incidentes críticos de privacidade/segurança.

Clique, tempo de tela e volume de cadastros são métricas de apoio, nunca a métrica central.

---

## 2. Contexto e problema

### 2.1 Contexto

Profissionais autônomos locais dependem de círculos sociais, indicação boca a boca, redes sociais fragmentadas e algoritmos de plataformas que nem sempre privilegiam descoberta local. Clientes têm dificuldade para descobrir profissionais confiáveis, comparar opções e avaliar qualidade fora da própria bolha social.

A documentação do projeto já consolidou problema, proposta, princípios, glossário, tipos de serviço e hipóteses iniciais. O que falta é evidência de campo e um produto testável.

### 2.2 Problema observado (hipóteses)

O processo de descoberta e conexão é:

- fragmentado entre canais informais e plataformas desconectadas;
- dependente de visibilidade pré-existente;
- pouco transparente sobre por que um profissional aparece;
- limitado por fatores geográficos e sociais;
- carente de mecanismos de confiança adequados a serviços locais;
- pouco acessível para trabalhadores com baixa familiaridade digital.

### 2.3 Pergunta central

> Como uma plataforma pode melhorar a descoberta e a conexão entre profissionais autônomos locais e demandas por serviços de forma justa, acessível, confiável e escalável?

### 2.4 Evidências que ainda faltam

- Entrevistas com profissionais e clientes reais.
- Auditoria de como plataformas comparáveis distribuem visibilidade.
- Medição de como distância, reputação e disponibilidade afetam conversão.
- Teste de mecanismos de fairness sem sacrificar relevância.
- Validação de que transparência de ranking é compreensível para usuários de baixa literacia digital.

Enquanto essas evidências não existirem, os problemas acima permanecem **hipóteses**, não fatos.

---

## 3. Visão, missão e princípios

### 3.1 Visão

Um ecossistema local onde profissionais competentes possam ser descobertos pelo que sabem fazer, não apenas por quem já conhecem, e onde clientes encontrem serviços com confiança, clareza e segurança.

### 3.2 Missão do MVP

Provar que uma experiência de descoberta local simples, explicável e com regras explícitas de distribuição de visibilidade consegue gerar conexões de qualidade para profissionais com pouca ou nenhuma presença digital.

### 3.3 Princípios de produto

1. **Descoberta justa:** reduzir invisibilidade extrema e concentração excessiva sem prometer igualdade artificial de resultados.
2. **Conexão antes de engajamento:** o sucesso é uma conexão útil, não tempo de tela.
3. **Simplicidade inicial:** resolver o fluxo principal com poucas telas e poucas decisões.
4. **Contexto local importa:** distância, disponibilidade e região são sinais de relevância para serviços presenciais.
5. **Serviços presenciais e remotos são diferentes:** a experiência e o ranking precisam reconhecer essa diferença.
6. **Confiança é construída:** verificações, histórico, transparência e consentimento valem mais do que estrelas isoladas.
7. **Reputação é contextual:** uma nota geral não resume qualidade para todos os serviços e contextos.
8. **Acessibilidade é requisito:** mobile-first, linguagem simples, teclado, leitor de tela e baixo consumo de dados.
9. **Privacidade por padrão:** localização aproximada, contato consentido, minimização de dados.
10. **Sem pay-to-win:** no MVP, dinheiro não compra posição no ranking.
11. **Hipóteses não são fatos:** decisões relevantes devem ser registradas, medidas e revisáveis.
12. **O repositório é memória viva:** spec, decisão, evidência e aprendizado ficam rastreáveis.

### 3.4 Não objetivos

- Substituir redes sociais ou se tornar feed de conteúdo.
- Ser plataforma genérica de freelancers globais.
- Processar pagamentos, cobrar comissão ou reter dinheiro.
- Oferecer agenda completa, videoconferência ou chat em tempo real.
- Fazer background check, garantia de serviço ou seguro.
- Prometer "melhores profissionais" ou "100% verificados".
- Criar aplicativo nativo antes de validar o fluxo web.
- Usar IA generativa para decisões de ranking ou moderação sem revisão humana.
- Expandir para múltiplas cidades antes de validar a densidade local.

---

## 4. Públicos, personas e jobs to be done

Detalhes em `docs/product/personas-and-journeys.md`.

### 4.1 Atores do sistema

| Ator | Descrição | Permissões principais |
| --- | --- | --- |
| Visitante | Pessoa sem conta | Buscar, ver perfis públicos, iniciar cadastro |
| Cliente | Pessoa que busca e contrata serviços | Criar pedido, buscar, demonstrar interesse, avaliar |
| Profissional | Autônomo que oferece serviços | Criar perfil, declarar serviços/área, receber pedidos, aceitar conexão |
| Operador | Equipe de operação e moderação | Moderar denúncias, suspender perfis, gerenciar catálogo |
| Pesquisador | Analista de produto/pesquisa | Ver métricas agregadas, criar experimentos, exportar dados anonimizados |
| Administrador | Dono do sistema | Configurar políticas, papéis e feature flags; auditar ações |

### 4.2 Personas primárias

| ID | Persona | Necessidade central | Critério de sucesso |
| --- | --- | --- | --- |
| P1 | Cliente urgente | resolver um problema hoje, no celular | encontrar em até 2 min e entender o próximo passo |
| P2 | Cliente planejador | comparar opções e reputação | comparar 3–5 profissionais e retomar depois |
| P3 | Profissional invisível | ser descoberto sem presença digital | criar perfil sem ajuda e receber oportunidade compatível |
| P4 | Profissional em crescimento | receber pedidos qualificados | responder rápido e construir reputação |
| P5 | Operador de confiança | analisar denúncias com contexto | decidir com critério, auditoria e possibilidade de recurso |

### 4.3 JTBD

- "Quando preciso de um serviço local, quero encontrar alguém confiável perto de mim sem depender só de indicação."
- "Quando sou autônomo, quero ser encontrado por clientes da minha região sem pagar por visibilidade."
- "Quando recebo uma oportunidade, quero saber se faz sentido para mim antes de expor meu contato."
- "Quando contrato alguém, quero sinais honestos de confiança e um canal seguro para começar."
- "Quando algo dá errado, quero denunciar e entender o que a plataforma fez."

---

## 5. Jornadas principais

Detalhes, fluxos alternativos e fricções em `docs/product/personas-and-journeys.md`.

| Jornada | Resultado esperado | Requisitos principais |
| --- | --- | --- |
| Onboarding do profissional | perfil publicável em poucos minutos, sem documentos obrigatórios | RF-01, RF-02, RF-03, RF-04 |
| Busca e conexão do cliente | encontrar profissional relevante e iniciar contato consentido | RF-05, RF-06, RF-07, RF-11, RF-12 |
| Mural de necessidade | publicar demanda e receber interesse de profissionais elegíveis | RF-08, RF-09, RF-10 |
| Desfecho e reputação | registrar resultado e construir reputação contextual | RF-13, RF-14 |
| Moderação e segurança | denunciar, bloquear e obter resposta proporcional | RF-17, RF-18 |
| Privacidade e controle | entender e exercer direitos sobre dados pessoais | RF-19 |

---

## 6. Escopo do MVP

### 6.1 Escopo funcional incluído

| Área | Incluído no MVP | Spec |
| --- | --- | --- |
| Contas e acesso | cadastro, login, confirmação de e-mail, reset, sessão, papéis | foundation-platform |
| Perfil do usuário | dados editáveis, preferências de contato e privacidade | foundation-platform |
| Perfil profissional | criação, publicação, área, disponibilidade, pausa | foundation-platform |
| Portfólio | trabalhos com título, descrição, categoria e imagem opcional | foundation-platform |
| Catálogo | categorias, subcategorias, sinônimos e seed curado | foundation-platform |
| Descoberta | busca por termo/categoria/localização, filtros, ranking explicável | local-discovery |
| Fairness | caps, exploração, boost de novo, diversidade, métricas | fair-visibility |
| Demanda | publicação de necessidade, feed de oportunidades, interesses e conexões | service-request |
| Confiança | avaliações pós-conexão, resposta, sinais, denúncias e moderação | trust-reputation |
| Operação | admin mínimo, catálogo, moderação, suspensão, auditoria | trust-reputation + foundation |
| Acessibilidade | WCAG 2.2 AA, mobile-first, teclado, leitor de tela | web-experience |
| Privacidade | localização aproximada, contato mascarado, consentimento, direitos | foundation-platform |
| Pesquisa | eventos de funil, métricas agregadas, experimentos | fair-visibility + local-discovery |

### 6.2 Fora do escopo do MVP

- Pagamentos, assinaturas, comissão, carteira ou prêmios.
- Agendamento com calendário e confirmação automática.
- Chat em tempo real, chamadas ou videochamada.
- Aplicativos nativos iOS/Android.
- Feed social, seguidores, curtidas ou conteúdo viral.
- Ranking pago, destaque pago ou leilão de posição.
- Verificação de antecedentes criminais, garantia ou seguro.
- Matching automático por IA generativa.
- Múltiplas cidades, internacionalização e multi-moeda.
- Integrações com ERPs, CRMs ou marketplaces externos.

### 6.3 Duas camadas de MVP

Para evitar confundir validação de produto com escopo de engenharia:

| Camada | Objetivo | Escopo |
| --- | --- | --- |
| **MVP de validação (piloto)** | validar hipóteses com usuários reais | todas as capacidades da seção 6.1, em uma região e 2–3 categorias |
| **Fatia técnica 1 (primeiro ciclo)** | provar a fundação e o fluxo principal ponta a ponta | contas, perfil, catálogo, busca, pedido, interesse, conexão, avaliação, denúncia, admin mínimo, fairness básica |

A fatia técnica 1 é o que as specs iniciais detalham. O restante do MVP de validação depende de evidência e será detalhado em specs futuras.

### 6.4 Estratégia de piloto

- **Geografia:** uma cidade ou região metropolitana com densidade suficiente; default de planejamento: "Cidade Piloto" sintética nos testes.
- **Categorias:** 2 a 3 categorias com demanda e oferta mapeadas; pelo menos uma presencial e uma remota/híbrida.
- **Oferta:** curadoria manual inicial de profissionais, com convite e suporte de onboarding.
- **Demanda:** parcerias locais, comunidades e campanhas de baixo custo.
- **Duração:** 8 a 12 semanas, com marcos de aprendizado a cada 2 semanas.
- **Critério de continuidade:** metas de conexão, fairness e confiança definidas na seção 10 e em `docs/product/metrics-and-experiments.md`.

### 6.5 Escopo transversal de qualidade

| Área | Compromisso |
| --- | --- |
| Usabilidade | experiência simples, linguagem clara, redução de etapas |
| Acessibilidade | WCAG 2.2 AA nas jornadas críticas |
| Responsividade | mobile-first, funcional em 360 px |
| Desempenho | busca p95 ≤ 500 ms; APIs p95 ≤ 300 ms; LCP p75 ≤ 2,5 s |
| Privacidade | minimização, localização aproximada, contato consentido, LGPD |
| Segurança | OWASP ASVS nível 2 adaptado, authz por recurso, rate limit, auditoria |
| Transparência | explicação de ranking, reputação e limitações |
| Observabilidade | logs estruturados, métricas, correlação e alertas |

---

## 7. Requisitos funcionais

**Prioridade:** M = must para o MVP; S = should; C = could; W = won't neste ciclo.

O catálogo detalhado, com todos os sub-requisitos e prioridades, está em `docs/product/requirements-catalog.md`. A tabela abaixo é a visão de produto para rastreabilidade.

| ID | Requisito | Prioridade | Spec | Dono |
| --- | --- | --- | --- | --- |
| RF-01 | Identidade, conta e acesso | M | foundation-platform | Victor |
| RF-02 | Gestão do perfil do usuário | M | foundation-platform | Victor |
| RF-03 | Gestão do perfil profissional | M | foundation-platform | Victor |
| RF-04 | Portfólio profissional | M | foundation-platform | Anderson |
| RF-05 | Catálogo de serviços | M | foundation-platform | Victor |
| RF-06 | Descoberta de profissionais | M | local-discovery | Victor |
| RF-07 | Relevância, equidade e transparência da descoberta | M | local-discovery + fair-visibility | Victor |
| RF-08 | Publicação e gestão de necessidades | M | service-request | Anderson |
| RF-09 | Descoberta de oportunidades | M | service-request | Anderson |
| RF-10 | Manifestações de interesse | M | service-request | Anderson |
| RF-11 | Solicitações de conexão direta | M | service-request | Anderson |
| RF-12 | Conexões consentidas | M | service-request + trust-reputation | Felipe |
| RF-13 | Registro de desfechos | M | trust-reputation | Felipe |
| RF-14 | Avaliações e reputação bilateral | M | trust-reputation | Felipe |
| RF-15 | Confirmação e verificação de informações | M | trust-reputation + foundation | Felipe |
| RF-16 | Notificações | M | service-request | Anderson |
| RF-17 | Denúncias e proteção do usuário | M | trust-reputation | Felipe |
| RF-18 | Administração e moderação | M | trust-reputation + foundation | Felipe |
| RF-19 | Transparência e controle de dados pessoais | M | foundation-platform | Victor |
| RF-20 | Medição e validação do produto | M | fair-visibility + local-discovery | Victor |

### 7.1 Critérios de aceite de produto por requisito

Cada RF só é considerado aceito quando:

1. a spec correspondente está `aprovada`;
2. todos os cenários de `acceptance.md` estão verdes ou justificados;
3. o gate de segurança/privacidade não tem bloqueios;
4. o gate de fairness foi avaliado quando aplicável;
5. há evidência anexada em `specs/<feature>/evidence/`;
6. a copy em PT-BR foi revisada por produto;
7. o comportamento funciona em 360 px e por teclado.

Os critérios específicos estão nas specs. O PRD não substitui o aceite de feature.

---

## 8. Requisitos não funcionais

Os RNFs são vinculantes para o MVP técnico e para o piloto. Metas quantitativas são referências iniciais e devem ser medidas em condições documentadas.

| ID | Requisito | Meta / critério | Verificação |
| --- | --- | --- | --- |
| RNF-01 | Manutenibilidade e modularidade | monólito modular, contextos delimitados, sem dependência circular | teste de arquitetura + review |
| RNF-02 | Desempenho | busca p95 ≤ 500 ms; APIs p95 ≤ 300 ms; LCP p75 ≤ 2,5 s; INP ≤ 200 ms; CLS ≤ 0,1 | benchmark + teste de carga |
| RNF-03 | Capacidade | 5 mil profissionais, 50 mil usuários ativos, 1 região piloto | modelo de carga + teste |
| RNF-04 | Disponibilidade | 99,5% mensal nas capacidades críticas | monitoramento + SLO |
| RNF-05 | Confiabilidade e integridade | invariantes, transações, idempotência, concorrência | testes de invariantes e concorrência |
| RNF-06 | Resiliência | fallback de ranking/notificação; nenhuma falha auxiliar derruba o fluxo crítico | testes de falha + caos leve |
| RNF-07 | Recuperabilidade | RPO/RTO definidos, backup testado, runbook | teste de restore |
| RNF-08 | Segurança | OWASP ASVS nível 2 adaptado; authz por recurso; rate limit; segredos fora do código | SAST, DAST, review |
| RNF-09 | Privacidade e LGPD | privacy by design, minimização, retenção, direitos do titular | DPIA + review + testes |
| RNF-10 | Acessibilidade | WCAG 2.2 AA nas jornadas críticas | axe + teclado + leitor de tela |
| RNF-11 | Usabilidade e inclusão digital | linguagem simples, baixa fricção, estados claros | teste com usuários |
| RNF-12 | Compatibilidade | 360 px, navegadores relevantes, conectividade limitada | matriz de dispositivos |
| RNF-13 | Observabilidade | logs JSON, correlation ID, métricas, traces quando justificado | revisão + dashboards |
| RNF-14 | Auditabilidade | ações administrativas append-only, ator, motivo, timestamp | testes + review |
| RNF-15 | Testabilidade | domínio testável, fixtures, cobertura de aceite, sem mock de banco em integração | matriz de testes |
| RNF-16 | Equidade e explicabilidade | critérios versionados, coortes observáveis, sem atributo sensível | simulação + review |
| RNF-17 | Eficiência de custos | custo fixo baixo; serviços gerenciados justificados | orçamento + revisão |
| RNF-18 | Implantabilidade | ambientes reproduzíveis, migrations versionadas, rollback | CI + runbook |
| RNF-19 | Portabilidade | containers, formatos abertos, dependências isoladas | revisão de arquitetura |
| RNF-20 | Evolutividade | evolução incremental, compatibilidade de dados | ADR + roadmap |

### 8.1 Requisitos transversais

Todos os RFs devem observar:

- autorização por papel e por recurso;
- validação e normalização de entrada;
- ausência de PII em logs, eventos e analytics;
- acessibilidade e linguagem simples;
- estados de carregamento, vazio, erro e sucesso;
- rollback e feature flag quando a mudança for relevante;
- rastreabilidade para spec, ticket e evidência.

---

## 9. Estrutura de equidade, confiança e transparência

Detalhes técnicos em `docs/architecture/ranking-and-fairness.md`.

### 9.1 Definição operacional de fairness

Fairness é a busca por reduzir invisibilidade extrema e concentração excessiva de oportunidades, sem garantir igualdade artificial de resultados e sem sacrificar a relevância para o cliente.

**Coortes observáveis permitidas:** profissional novo, profissional com poucas avaliações, profissional com baixa atividade recente, região/bairro, categoria, tipo de serviço.

**Proibido:** raça, gênero, idade, deficiência, religião, orientação sexual ou inferência desses atributos.

### 9.2 Mecanismos iniciais

- limite de exposição por profissional por sessão/dia;
- slots de exploração (percentual da primeira página);
- boost de novo profissional com decaimento;
- diversidade de resultados por bairro/categoria;
- ordenação por bandas de relevância antes de diversificar;
- explicação transparente do motivo da posição;
- fallback para ranking baseline se a camada de fairness falhar.

**Defaults propostos para o primeiro ciclo:**

| Parâmetro | Default | Dono |
| --- | --- | --- |
| Exploração | 20% dos slots da primeira página | Victor |
| Cap por sessão | 2 impressões por profissional | Victor |
| Cap por dia | 10 impressões por profissional | Victor |
| Boost de novo | 30 dias ou 5 conexões, com decaimento | Victor |
| Piso de relevância | percentil 40 do baseline | Victor + Produto |
| Guardrail de relevância | queda ≤ 2 p.p. no CTR e ≤ 5% no nDCG | Produto + Pesquisa |

### 9.3 Métricas de fairness

- participação de exposição por coorte;
- coeficiente de Gini de impressões;
- share do top 10/20 por sessão e por dia;
- share de novos profissionais nas primeiras posições;
- posição média por coorte;
- conversão impressão → clique → interesse → conexão;
- satisfação e relevância percebida.

### 9.4 Confiança

**Camadas:**

1. Conta verificada: e-mail confirmado.
2. Contato verificado: telefone/e-mail validado.
3. Identidade verificada: documento analisado manualmente, com retenção mínima (fora do primeiro ciclo).
4. Histórico: conexões, avaliações, resposta e conclusão.

**Regras:**

- verificação é opcional para participar, mas exibida com clareza;
- nunca prometer background check;
- avaliação só após conexão elegível;
- sinais comportamentais têm peso menor que relevância;
- profissionais novos não são penalizados por falta de histórico além do necessário;
- denúncias têm prazo e critério públicos;
- decisões de moderação são auditáveis e recorríveis.

### 9.5 Transparência

- página pública "Como o RAIRE ordena os resultados";
- explicação curta em cada resultado;
- canal para contestar posição ou avaliação;
- registro versionado de mudanças de política;
- relatório periódico de fairness.

---

## 10. Privacidade, dados e LGPD

Detalhes em `docs/security/privacy-and-lgpd.md`.

### 10.1 Princípios

- minimização;
- finalidade;
- consentimento granular quando exigido;
- segurança em trânsito e repouso;
- PII separada de dados públicos;
- transparência sobre o que é público;
- retenção e anonimização;
- responsabilização.

### 10.2 Classificação de dados

| Classe | Exemplos | Tratamento |
| --- | --- | --- |
| Pública | nome de exibição, bio, categorias, área aproximada | visível na plataforma |
| Interna | status de verificação, histórico de edição | restrita à operação |
| Pessoal | e-mail, telefone, localização precisa | criptografada, acesso mínimo |
| Sensível | documento de identidade, evidência de denúncia | acesso restrito, retenção curta, auditoria |
| Analítica | eventos pseudonimizados, impressões | sem PII, retenção agregada |

### 10.3 Direitos do titular

- acesso e portabilidade;
- correção;
- exclusão, ressalvadas obrigações legais;
- revogação de consentimento;
- informação sobre compartilhamento;
- canal de contato e prazo de resposta.

### 10.4 Retenção inicial proposta

| Dado | Retenção |
| --- | --- |
| Conta ativa | enquanto durar a relação |
| Conta excluída | anonimização em até 30 dias |
| Localização precisa | até 90 dias, salvo necessidade de segurança |
| Eventos analíticos | 13 meses pseudonimizados |
| Denúncias e auditoria | 5 anos, acesso restrito |
| Documentos de verificação | 90 dias após decisão, salvo obrigação legal |
| Logs de aplicação | 30–90 dias, sem PII |

> Prazos propostos; validar com jurídico antes do piloto.

### 10.5 Contato e localização

- contato público sempre mascarado até consentimento;
- contato liberado somente após aceite mútuo;
- localização pública aproximada (cidade/bairro/raio);
- localização precisa separada, criptografada e com acesso restrito;
- nenhum endereço residencial em resposta pública, log ou evento.

---

## 11. Métricas, validação e experimentação

Detalhes em `docs/product/metrics-and-experiments.md`.

### 11.1 Métricas principais

| Dimensão | Métrica | Meta inicial do piloto |
| --- | --- | --- |
| Oferta | perfis completos publicados | 150–300 na região piloto |
| Descoberta | buscas com resultado relevante | ≥ 70% |
| Descoberta | busca → perfil | ≥ 25% |
| Demanda | pedidos com ≥ 3 interesses | ≥ 50% |
| Conexão | interesse → conexão aceita | ≥ 30% |
| Conexão | tempo até primeiro interesse | < 24 h |
| Confiança | conexões com avaliação | ≥ 20% no primeiro ciclo |
| Fairness | Gini de impressões | redução vs. baseline |
| Fairness | novos profissionais no top 20 | ≥ 20% |
| Segurança | incidentes de privacidade | 0 |
| Segurança | denúncias procedentes | < 2% das conexões |
| Acessibilidade | jornadas críticas WCAG 2.2 AA | 100% |
| Pesquisa | hipóteses com evidência registrada | 100% das priorizadas |

### 11.2 Experimentos prioritários

| ID | Hipótese | Desenho | Critério de sucesso |
| --- | --- | --- | --- |
| EXP-01 | boost de novos profissionais aumenta diversidade sem perder relevância | A/B com guardrail de CTR e satisfação | +20% de novos no top 20; queda de relevância ≤ 2 p.p. |
| EXP-02 | explicação do ranking aumenta confiança e clique | A/B com e sem explicação | +10% de clique no perfil; confiança autorrelatada maior |
| EXP-03 | distância é decisiva em serviços presenciais | análise de funil por faixa de distância | queda de conversão > 30% acima de X km |
| EXP-04 | mural de necessidade gera conexões de qualidade | piloto com 2 categorias | ≥ 50% dos pedidos com 3 interesses; ≥ 30% de aceite |
| EXP-05 | reputação contextual melhora decisão | distribuição + tags vs. média | maior confiança e menor arrependimento |

Nenhum experimento começa sem revisão de privacidade e fairness. Todo experimento tem hipótese, métrica, guardrail, duração e critério de parada.

---

## 12. Riscos

Detalhes, donos e acompanhamento em `docs/product/risks.md`.

| Risco | Impacto | Mitigação |
| --- | --- | --- |
| Cold start / baixa liquidez | produto parece vazio | piloto denso, curadoria de oferta, mural de necessidade, parcerias locais |
| Perfis falsos ou fraude | perda de confiança | verificação manual, sinais de comportamento, denúncia, moderação |
| Assédio ou vazamento de contato | dano a usuários | contato mascarado, consentimento, bloqueio, rate limit, resposta a incidentes |
| Vazamento de localização precisa | risco físico e jurídico | localização aproximada, criptografia, acesso restrito, retenção curta |
| Manipulação de avaliações | reputação distorcida | avaliação pós-conexão, detecção de padrões, moderação, direito de resposta |
| Viés ou discriminação no ranking | injustiça e risco reputacional | sem atributos sensíveis, coortes observáveis, auditoria, revisão humana |
| Gaming de visibilidade | fairness quebrada | limites, detecção de tráfego anômalo, logs auditáveis |
| Baixa literacia digital | exclusão do público-alvo | UX simples, onboarding assistido, linguagem clara, teste com usuários |
| Pressão por monetização | conflito com fairness | sem pay-to-win no MVP; princípio explícito; revisão humana |
| Custo operacional de moderação | falha de segurança | fila priorizada, critérios claros, automação só como triagem |
| Dependência de uma região | aprendizado limitado | critérios de expansão baseados em densidade e evidência |

**Condições impeditivas para o piloto:** incidente crítico de privacidade sem correção; ranking com atributo sensível; ausência de moderação; ausência de DPIA; ausência de plano de incidente; ausência de aprovação humana.

---

## 13. Roadmap e pipeline

Detalhes em `docs/product/roadmap.md` e `docs/process/pipeline.md`.

| Fase | Objetivo | Entregas principais | Dono | Critério de saída |
| --- | --- | --- | --- | --- |
| F0 | Consolidação e planejamento | PRD v1.0, specs, tickets, ADRs, pipeline | Todos | specs prontas para revisão humana |
| F1 | Fundação técnica | contas, perfis, catálogo, admin, observabilidade | Victor | jornada de perfil publicada + testes |
| F2 | Descoberta local | busca, ranking baseline, explicação, perfil público | Victor + Anderson | busca com resultado explicável + acessível |
| F3 | Demanda e conexão | mural, interesses, conexões, notificações | Anderson + Felipe | fluxo pedido → aceite → contato |
| F4 | Confiança e reputação | avaliações, desfechos, denúncias, moderação | Felipe | avaliação elegível + moderação auditável |
| F5 | Fairness e experimentação | política, ledger, métricas, experimentos | Victor | política versionada + simulação |
| F6 | Preparação do piloto | acessibilidade, segurança, operação, evidências | Todos | gates verdes + aprovação humana |

> A ordem pode mudar com evidência. Fairness não é "fase final": os guardrails básicos entram já na F2; a F5 aprofunda mecanismos e medição.

### 13.1 Pipeline de trabalho

```txt
1. Refinar requisito no PRD
2. Escrever spec + aceite + tickets + plano
3. Revisão humana da spec
4. Refinar tickets e dependências
5. Implementar com testes primeiro
6. Review independente
7. Gate de segurança/privacidade
8. Gate de trust & safety
9. Gate de QA
10. Gate de produto
11. Evidência + release candidate
12. Aprovação humana para piloto
```

Detalhes de papéis, artefatos, SLAs e critérios de bloqueio em `docs/process/pipeline.md`.

---

## 14. Equipe, RACI e divisão de trabalho

Detalhes em `docs/team/raci.md`, `docs/team/task-board.md` e `docs/team/README.md`.

| Pessoa | Frente | Entrega principal | Também responde por |
| --- | --- | --- | --- |
| **Victor** | Arquitetura + backend + ranking | foundation, discovery, fairness, ADRs, contratos | revisão de arquitetura, performance, migrations |
| **Anderson** | Web + produto + acessibilidade | web-experience, jornadas, design system, copy PT-BR | product acceptance, UX, acessibilidade |
| **Felipe** | QA + segurança + privacidade + DevOps | trust-reputation, moderação, pipeline, evidências | gates de segurança/fairness, CI/CD, runbooks |

### 14.1 Regra de independência

- Quem implementa não aprova o próprio código.
- Victor revisa arquitetura/backend; Anderson revisa produto/web; Felipe revisa segurança/QA/trust & safety.
- Em caso de conflito, o gate de segurança/privacidade bloqueia até decisão humana.
- Toda decisão de fairness, privacidade ou escopo tem registro.

---

## 15. Critérios de aceite e maturidade do PRD

**Status:** aprovado em 2026-11-20 para planejamento; região real do piloto e DPIA continuam pendentes antes de F6.

O PRD será considerado pronto para orientar a implementação quando:

- [x] Problema, visão, personas e jornadas estiverem revisados por produto e pesquisa.
- [x] RFs e RNFs tiverem prioridade e critérios verificáveis.
- [x] Escopo do MVP, fatia técnica e não objetivos estiverem explícitos.
- [ ] Região real, metas finais e operação do piloto estiverem definidas.
- [ ] Política de fairness e privacidade tiver revisão de segurança/jurídico final.
- [x] Riscos tiverem dono e mitigação.
- [x] Perguntas abertas tiverem decisão ou plano de investigação.
- [x] Specs iniciais estiverem derivadas e rastreadas aos RFs.
- [x] Harness, pipeline e divisão de trabalho estiverem operacionais.
- [x] Nenhuma implementação tiver começado sem aprovação humana das specs.

### 15.1 Estados de maturidade

| Estado | Significado |
| --- | --- |
| `rascunho` | em elaboração; não orienta implementação |
| `em revisão` | pronto para revisão humana; comentários aceitos |
| `aprovado` | aprovado por humano; orienta implementação |
| `em execução` | implementação em andamento |
| `concluído` | critérios de aceite e evidências completos |
| `arquivado` | substituído por nova versão |

### 15.2 Condições de aprovação

A aprovação exige registro humano de produto, arquitetura e segurança/privacidade. Nenhum agente de IA pode mover o PRD para `aprovado`.

---

## 16. Decisões abertas

**Status:** decididas em 2026-11-20. Registro canônico em `docs/product/decisions-2026-11-20.md`; a tabela abaixo mantém o default histórico.

Detalhes e prazos em `docs/product/open-questions.md`.

| ID | Pergunta | Dono | Default seguro |
| --- | --- | --- | --- |
| Q-01 | Qual cidade/região piloto? | Produto | Cidade sintética nos testes; região real antes do piloto |
| Q-02 | Quais categorias iniciais? | Produto | manutenção/reparos + aulas/consultoria |
| Q-03 | Verificação de identidade é obrigatória? | Produto + Segurança | opcional; e-mail confirmado obrigatório |
| Q-04 | Contato pode ser liberado por telefone/WhatsApp? | Segurança + Produto | sim, após aceite, com mascaramento |
| Q-05 | Como confirmar que o serviço ocorreu? | Produto | confirmação de um lado + contestação |
| Q-06 | Qual o modelo de monetização futuro? | Fundador | definir depois; nunca ranking pago |
| Q-07 | Qual o apetite para moderação manual? | Trust & Safety | fila pequena, SLA 48 h |
| Q-08 | Quais prazos de retenção o jurídico exige? | Jurídico | proposta da seção 10.4 |
| Q-09 | Pesquisa pode usar dados agregados? | Privacidade | consentimento separado + anonimização |
| Q-10 | Qual stack de frontend e hospedagem? | Arquitetura | server-rendered + SQLite/Postgres |
| Q-11 | Qual o limite aceitável de queda de relevância? | Produto + Pesquisa | 2 p.p. CTR / 5% nDCG |
| Q-12 | Quem aprova mudanças de política de ranking? | Governança | produto + pesquisa + segurança |

---

## 17. Rastreabilidade

### 17.1 Hipóteses → requisitos

| Hipótese | Requisitos relacionados |
| --- | --- |
| H1/H2/H3 — descoberta e visibilidade | RF-06, RF-07, RF-20 |
| H4/H5/H6 — confiança e reputação | RF-14, RF-15, RF-17 |
| H7/H8/H9 — contexto geográfico | RF-03, RF-06, RF-07 |
| H10/H11/H12 — comportamento | RF-08, RF-10, RF-16 |
| H13/H14/H15 — matching e relevância | RF-06, RF-07 |
| H16/H17 — IA e recomendação | RF-07, RF-20 (sem IA generativa no MVP) |

### 17.2 Requisitos → specs

| Spec | Requisitos | Dono |
| --- | --- | --- |
| `foundation-platform` | RF-01, RF-02, RF-03, RF-04, RF-05, RF-19; RNF-01, RNF-05, RNF-08, RNF-09, RNF-14 | Victor |
| `web-experience` | RNF-10, RNF-11, RNF-12; jornadas 6.1–6.10 | Anderson |
| `local-discovery` | RF-06, RF-07, RF-20; RNF-02, RNF-13, RNF-16 | Victor |
| `service-request` | RF-08, RF-09, RF-10, RF-11, RF-12, RF-16; RNF-05, RNF-06 | Anderson |
| `trust-reputation` | RF-12, RF-13, RF-14, RF-15, RF-17, RF-18; RNF-08, RNF-14 | Felipe |
| `fair-visibility` | RF-07, RF-20; RNF-16 | Victor |

### 17.3 Requisitos → tickets → testes

Cada spec mantém:

- `tickets.md` com IDs de ticket, requisito, dono e dependência;
- `acceptance.md` com matriz cenário → requisito → tipo de teste → status;
- `implementation-plan.md` com fatias, flags, migrations e riscos;
- `review.md`, `qa-report.md` e `security-review.md` após execução.

### 17.4 Documentos relacionados

- `docs/product/requirements-catalog.md`
- `docs/product/personas-and-journeys.md`
- `docs/product/metrics-and-experiments.md`
- `docs/product/risks.md`
- `docs/product/roadmap.md`
- `docs/product/open-questions.md`
- `docs/product/catalog-seed.md`
- `docs/product/transparency-ranking.md`
- `docs/vision/problem.md`
- `docs/vision/proposal.md`
- `docs/vision/principles.md`
- `docs/vision/hypothesis.md`
- `docs/vision/service-types.md`
- `docs/vision/glossary.md`
- `docs/architecture/architecture-plan.md`
- `docs/architecture/ranking-and-fairness.md`
- `docs/security/privacy-and-lgpd.md`
- `docs/security/threat-model.md`
- `docs/process/harness.md`
- `docs/process/pipeline.md`
- `docs/process/review-criteria.md`
- `docs/process/qa-strategy.md`
- `docs/decisions/README.md`
- `specs/README.md`

---

## 18. Changelog

### v1.0 — 2026-11-20

- Revisão completa do PRD 0.1.
- Separação entre MVP de validação e fatia técnica.
- Requisitos RF-01..RF-20 consolidados e rastreados.
- Catálogo detalhado movido para `requirements-catalog.md`.
- Personas/jornadas, métricas, riscos, roadmap e questões abertas separados.
- Fairness com defaults operacionais, guardrails e governança.
- Privacidade/LGPD com classificação, retenção e direitos.
- Pipeline, RACI, task board e handoffs por pessoa.
- Templates de spec, aceite, ticket, review, QA e segurança.
- Nenhum código implementado na data da v1.0; a fundação foi implementada depois, em 2026-11-20.

### v0.1 — 2026-11-20

- Versão inicial, preservada em `docs/product/prd-v0.1.md`.
