# Catálogo de Requisitos Funcionais — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0 (`docs/product/prd.md`)  
**Uso:** detalhamento dos RF-01 a RF-20 para specs, tickets e critérios de aceite.

> Este catálogo preserva os sub-requisitos da versão 0.1 e os organiza para rastreabilidade. Ele não substitui as specs: as regras de negócio, estados, contratos e critérios de aceite de feature ficam em `specs/<feature>/`.

**Prioridade:** M = must; S = should; C = could; W = won't neste ciclo.

---

## RF-01 — Identidade, conta e acesso

**Objetivo:** permitir que usuários criem, acessem e gerenciem suas contas, preservando a identidade e a segurança de sua participação na plataforma.  
**Rastreabilidade:** OBJ-01; jornadas 6.1 e 6.2.  
**Spec:** `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-01.01 | Permitir a criação de uma conta de usuário. | M |
| RF-01.02 | Permitir a autenticação por e-mail e senha. | M |
| RF-01.03 | Permitir a confirmação do endereço de e-mail. | M |
| RF-01.04 | Restringir ações que exigem identidade confirmada até a conclusão das verificações aplicáveis. | M |
| RF-01.05 | Permitir o encerramento da sessão autenticada. | M |
| RF-01.06 | Permitir a recuperação de acesso à conta. | M |
| RF-01.07 | Disponibilizar os termos de utilização e informações de privacidade aplicáveis ao cadastro. | M |
| RF-01.08 | Registrar os aceites obrigatórios relacionados à utilização da plataforma. | M |
| RF-01.09 | Permitir que uma mesma conta exerça os papéis de cliente e profissional. | M |
| RF-01.10 | Permitir a exploração de informações públicas sem autenticação. | M |
| RF-01.11 | Retomar a ação que motivou a autenticação, quando aplicável. | S |

**Critérios transversais:** senha com hash forte; sessão revogável; rate limit; respostas que não enumeram contas; nenhum dado sensível em log.

---

## RF-02 — Gestão do perfil do usuário

**Objetivo:** permitir que usuários gerenciem suas informações pessoais e preferências de participação na plataforma.  
**Rastreabilidade:** OBJ-01, OBJ-06 e OBJ-07; jornadas 6.1 e 6.5.  
**Spec:** `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-02.01 | Permitir a consulta e atualização das informações pessoais editáveis. | M |
| RF-02.02 | Permitir a configuração dos meios de contato autorizados. | M |
| RF-02.03 | Permitir o gerenciamento das preferências de compartilhamento de informações. | M |
| RF-02.04 | Permitir o gerenciamento das preferências de notificações não essenciais. | M |
| RF-02.05 | Permitir que usuários consultem suas conexões e interações registradas. | M |
| RF-02.06 | Permitir que usuários consultem suas informações de reputação, quando disponíveis. | S |
| RF-02.07 | Permitir a solicitação de desativação ou exclusão da conta, observadas as obrigações legais aplicáveis. | M |

---

## RF-03 — Gestão do perfil profissional

**Objetivo:** permitir que usuários apresentem suas competências e disponibilizem informações profissionais relevantes para descoberta e conexão.  
**Rastreabilidade:** OBJ-01, OBJ-03 e OBJ-08; jornadas 6.2 e 6.3.  
**Spec:** `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-03.01 | Permitir que um usuário inicie sua atuação profissional mediante criação de um perfil. | M |
| RF-03.02 | Permitir o registro de nome de apresentação, descrição e informações profissionais relevantes. | M |
| RF-03.03 | Permitir a declaração de um ou mais serviços oferecidos. | M |
| RF-03.04 | Permitir a associação dos serviços oferecidos às classificações disponíveis no catálogo. | M |
| RF-03.05 | Permitir a declaração e atualização da região e área de atendimento. | M |
| RF-03.06 | Permitir a declaração e atualização de disponibilidade. | M |
| RF-03.07 | Permitir a inclusão de imagem de apresentação profissional. | S |
| RF-03.08 | Permitir a visualização prévia das informações públicas do perfil. | S |
| RF-03.09 | Permitir a publicação do perfil quando satisfeitos os critérios mínimos de elegibilidade. | M |
| RF-03.10 | Permitir a atualização das informações profissionais publicadas. | M |
| RF-03.11 | Permitir a suspensão voluntária da exposição do perfil. | M |
| RF-03.12 | Permitir a reativação do perfil, quando satisfeitas as condições aplicáveis. | M |
| RF-03.13 | Apresentar orientações sobre informações necessárias à publicação e melhoria do perfil. | S |
| RF-03.14 | Permitir a inclusão de links profissionais externos, conforme as políticas de privacidade e segurança. | C |

**Invariantes:** perfil não publicado, suspenso ou inelegível não participa da descoberta pública; perfil presencial/híbrido exige área; remoto não exige.

---

## RF-04 — Portfólio profissional

**Objetivo:** permitir que profissionais demonstrem suas competências por meio da apresentação organizada de trabalhos e experiências anteriores.  
**Rastreabilidade:** OBJ-02, OBJ-03 e OBJ-07; jornadas 6.3 e 6.4.  
**Spec:** `foundation-platform` (com UI em `web-experience`).

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-04.01 | Permitir a criação de registros de trabalhos no portfólio profissional. | M |
| RF-04.02 | Permitir a inclusão de título, descrição e informações contextuais sobre o trabalho. | M |
| RF-04.03 | Permitir a associação de trabalhos aos serviços ou categorias correspondentes. | M |
| RF-04.04 | Permitir a inclusão de imagens relacionadas ao trabalho apresentado. | M |
| RF-04.05 | Permitir a edição de trabalhos publicados. | M |
| RF-04.06 | Permitir a remoção ou despublicação de trabalhos. | M |
| RF-04.07 | Disponibilizar os trabalhos publicados no perfil profissional público. | M |
| RF-04.08 | Permitir a consulta de trabalhos por contexto de serviço, quando aplicável. | S |
| RF-04.09 | Permitir a denúncia de conteúdo inadequado ou de informações indevidamente publicadas. | M |
| RF-04.10 | Permitir a definição da ordem de apresentação dos trabalhos. | C |

**Regra:** trabalho de portfólio não é experiência verificada pela plataforma.

---

## RF-05 — Catálogo de serviços

**Objetivo:** disponibilizar uma estrutura de classificação que permita organizar adequadamente serviços, profissionais e demandas.  
**Rastreabilidade:** OBJ-01, OBJ-03, OBJ-04 e OBJ-05; jornadas 6.2, 6.4, 6.6 e 6.7.  
**Spec:** `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-05.01 | Disponibilizar um catálogo inicial de categorias e serviços. | M |
| RF-05.02 | Permitir a utilização do catálogo na declaração de serviços profissionais. | M |
| RF-05.03 | Permitir a utilização do catálogo na busca de profissionais. | M |
| RF-05.04 | Permitir a utilização do catálogo na publicação e descoberta de demandas. | M |
| RF-05.05 | Permitir o gerenciamento do catálogo por operadores autorizados. | M |
| RF-05.06 | Permitir a ativação e desativação de classificações do catálogo. | M |
| RF-05.07 | Impedir novas associações a classificações indisponíveis. | M |
| RF-05.08 | Reconhecer variações terminológicas e sinônimos relevantes para descoberta de serviços. | S |
| RF-05.09 | Permitir o envio de sugestões de serviços ainda não contemplados pelo catálogo. | C |

**Invariante:** desativação preserva integridade histórica; perfil com categoria desativada sai da busca até ajuste.

---

## RF-06 — Descoberta de profissionais

**Objetivo:** permitir que clientes identifiquem profissionais adequados às suas necessidades e consultem informações relevantes para tomada de decisão.  
**Rastreabilidade:** OBJ-03, OBJ-06 e OBJ-08; jornada 6.4.  
**Spec:** `local-discovery`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-06.01 | Permitir a busca de profissionais por serviço, termo ou categoria. | M |
| RF-06.02 | Permitir a indicação da região desejada para atendimento. | M |
| RF-06.03 | Identificar profissionais elegíveis conforme os critérios aplicáveis à busca. | M |
| RF-06.04 | Apresentar resultados ordenados segundo critérios de relevância. | M |
| RF-06.05 | Permitir o refinamento dos resultados por critérios essenciais de descoberta. | M |
| RF-06.06 | Apresentar informações resumidas sobre cada profissional encontrado. | M |
| RF-06.07 | Permitir a consulta ao perfil profissional completo. | M |
| RF-06.08 | Permitir a consulta aos trabalhos publicados no portfólio. | M |
| RF-06.09 | Apresentar informações de disponibilidade e área de atendimento, quando declaradas. | M |
| RF-06.10 | Apresentar informações de reputação e verificação, quando disponíveis. | S |
| RF-06.11 | Informar os principais fatores de relevância dos resultados apresentados. | M |
| RF-06.12 | Apresentar orientações quando não existirem resultados compatíveis. | M |
| RF-06.13 | Permitir o acesso ao fluxo de publicação de necessidade a partir da descoberta. | S |
| RF-06.14 | Permitir a utilização opcional da localização do dispositivo para auxiliar a busca. | S |
| RF-06.15 | Permitir a ordenação alternativa por proximidade, quando aplicável. | S |
| RF-06.16 | Permitir salvar profissionais para consulta posterior. | C |

**Invariante:** nenhuma localização precisa, endereço residencial ou contato privado é exposto.

---

## RF-07 — Relevância, equidade e transparência da descoberta

**Objetivo:** estabelecer mecanismos de descoberta que preservem relevância, oportunidades de exposição e compreensão dos resultados.  
**Rastreabilidade:** OBJ-03 e OBJ-08; jornadas 6.4 e 6.7.  
**Specs:** `local-discovery` + `fair-visibility`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-07.01 | Ordenar resultados de descoberta conforme critérios documentados de relevância. | M |
| RF-07.02 | Considerar a compatibilidade entre serviço, região e critérios aplicáveis à busca. | M |
| RF-07.03 | Permitir a participação de profissionais elegíveis que ainda não possuam histórico consolidado na plataforma. | M |
| RF-07.04 | Aplicar mecanismos básicos destinados a reduzir concentrações injustificadas de exposição entre profissionais comparáveis. | M |
| RF-07.05 | Apresentar explicações compreensíveis sobre os principais fatores de classificação. | M |
| RF-07.06 | Disponibilizar informações gerais sobre os critérios utilizados nos mecanismos de descoberta. | M |
| RF-07.07 | Registrar informações necessárias à avaliação da distribuição de exposição. | M |
| RF-07.08 | Permitir a utilização de mecanismos adicionais de diversificação, quando justificados por evidências. | C |
| RF-07.09 | Permitir a avaliação controlada de mudanças nos mecanismos de classificação. | C |

**Proibições:** favorecimento pago; atributos pessoais sensíveis; exclusão automática por ausência de avaliações.

---

## RF-08 — Publicação e gestão de necessidades

**Objetivo:** permitir que clientes comuniquem demandas por serviços e recebam manifestações de interesse de profissionais.  
**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornada 6.6.  
**Spec:** `service-request`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-08.01 | Permitir a criação de uma necessidade de serviço por usuário autorizado. | M |
| RF-08.02 | Permitir a classificação da necessidade conforme o catálogo disponível. | M |
| RF-08.03 | Permitir a descrição das características relevantes do serviço procurado. | M |
| RF-08.04 | Permitir a indicação da região aproximada de atendimento. | M |
| RF-08.05 | Permitir a declaração de urgência, prazo ou condições temporais relevantes. | M |
| RF-08.06 | Permitir a indicação opcional de faixa de orçamento. | S |
| RF-08.07 | Permitir a publicação da necessidade quando satisfeitas as regras aplicáveis. | M |
| RF-08.08 | Permitir que o responsável consulte e gerencie suas necessidades. | M |
| RF-08.09 | Permitir a alteração de necessidades conforme seu estado e as regras de negócio. | M |
| RF-08.10 | Permitir o encerramento voluntário de uma necessidade. | M |
| RF-08.11 | Permitir o cancelamento de uma necessidade. | M |
| RF-08.12 | Encerrar a disponibilidade de necessidades expiradas. | M |
| RF-08.13 | Permitir que o cliente consulte e compare os profissionais interessados. | M |
| RF-08.14 | Impedir a exposição pública de informações privadas de contato e localização precisa. | M |

---

## RF-09 — Descoberta de oportunidades

**Objetivo:** permitir que profissionais identifiquem ativamente oportunidades de trabalho compatíveis com sua atuação.  
**Rastreabilidade:** OBJ-05 e OBJ-06; jornada 6.7.  
**Spec:** `service-request`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-09.01 | Disponibilizar uma área de descoberta de necessidades publicadas. | M |
| RF-09.02 | Apresentar oportunidades elegíveis conforme os serviços e a área de atuação profissional. | M |
| RF-09.03 | Permitir a busca ou filtragem de oportunidades por critérios relevantes. | M |
| RF-09.04 | Permitir a consulta aos detalhes de uma necessidade disponível. | M |
| RF-09.05 | Apresentar informações autorizadas e relevantes sobre o cliente responsável pela necessidade. | M |
| RF-09.06 | Apresentar informações de reputação do cliente, quando disponíveis e autorizadas. | S |
| RF-09.07 | Informar quando uma oportunidade estiver encerrada, indisponível ou expirada. | M |
| RF-09.08 | Apresentar estados vazios e orientações quando não existirem oportunidades compatíveis. | M |
| RF-09.09 | Permitir notificações sobre novas oportunidades compatíveis, conforme as preferências do profissional. | C |

---

## RF-10 — Manifestações de interesse

**Objetivo:** permitir que profissionais expressem interesse em necessidades publicadas por clientes.  
**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornadas 6.6 e 6.7.  
**Spec:** `service-request`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-10.01 | Permitir que um profissional elegível manifeste interesse em uma necessidade disponível. | M |
| RF-10.02 | Permitir a inclusão de uma mensagem contextual na manifestação de interesse. | M |
| RF-10.03 | Permitir que o cliente consulte as manifestações recebidas em suas necessidades. | M |
| RF-10.04 | Permitir que o cliente consulte os perfis dos profissionais interessados. | M |
| RF-10.05 | Permitir que o cliente aceite ou recuse manifestações de interesse. | M |
| RF-10.06 | Permitir que o profissional retire uma manifestação ainda pendente. | M |
| RF-10.07 | Informar o estado da manifestação às partes envolvidas. | M |
| RF-10.08 | Impedir manifestações incompatíveis com as regras de elegibilidade ou com o estado da necessidade. | M |
| RF-10.09 | Aplicar controles proporcionais contra manifestações abusivas, automatizadas ou repetitivas. | M |
| RF-10.10 | Originar uma conexão quando uma manifestação válida for aceita e as condições de consentimento forem satisfeitas. | M |

---

## RF-11 — Solicitações de conexão direta

**Objetivo:** permitir que clientes iniciem conexões diretamente a partir da descoberta de profissionais, independentemente da publicação de uma necessidade.  
**Rastreabilidade:** OBJ-03 e OBJ-06; jornada 6.5.  
**Spec:** `service-request`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-11.01 | Permitir que um cliente solicite conexão com um profissional elegível. | M |
| RF-11.02 | Permitir a inclusão de informações contextuais sobre o serviço desejado. | M |
| RF-11.03 | Permitir que o profissional consulte as solicitações recebidas. | M |
| RF-11.04 | Permitir que o profissional aceite ou recuse uma solicitação pendente. | M |
| RF-11.05 | Permitir que o cliente cancele uma solicitação ainda pendente. | M |
| RF-11.06 | Informar às partes o estado da solicitação. | M |
| RF-11.07 | Encerrar solicitações expiradas conforme as regras aplicáveis. | M |
| RF-11.08 | Impedir solicitações incompatíveis com restrições ou bloqueios existentes. | M |
| RF-11.09 | Originar uma conexão quando uma solicitação válida for aceita e as condições de consentimento forem satisfeitas. | M |

---

## RF-12 — Conexões consentidas

**Objetivo:** permitir que clientes e profissionais estabeleçam contato mediante manifestação válida de interesse e consentimento das partes.  
**Rastreabilidade:** OBJ-06 e OBJ-07; jornadas 6.5, 6.6 e 6.7.  
**Specs:** `service-request` + `trust-reputation`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-12.01 | Registrar conexões originadas de solicitações diretas ou manifestações de interesse aceitas. | M |
| RF-12.02 | Associar cada conexão ao contexto que motivou seu estabelecimento. | M |
| RF-12.03 | Disponibilizar informações de contato exclusivamente às partes autorizadas, conforme as preferências e permissões aplicáveis. | M |
| RF-12.04 | Impedir a disponibilização de informações privadas de contato antes do estabelecimento das condições de autorização. | M |
| RF-12.05 | Permitir que as partes consultem suas conexões estabelecidas. | M |
| RF-12.06 | Permitir que usuários denunciem irregularidades relacionadas a uma conexão. | M |
| RF-12.07 | Permitir o bloqueio de novas interações com outro usuário. | M |
| RF-12.08 | Permitir a interrupção de acessos futuros a informações compartilhadas, quando tecnicamente aplicável. | S |

**Limite:** conexão não confirma contratação, pagamento ou realização do serviço.

---

## RF-13 — Registro de desfechos

**Objetivo:** permitir que os usuários informem os resultados das interações originadas na plataforma.  
**Rastreabilidade:** OBJ-06 e OBJ-07; jornadas 6.8 e 6.9.  
**Spec:** `trust-reputation`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-13.01 | Permitir que participantes de uma conexão informem o desfecho da interação. | M |
| RF-13.02 | Permitir distinguir desfechos relacionados a contato, negociação e eventual realização do serviço. | M |
| RF-13.03 | Permitir que o usuário registre sua percepção sobre a utilidade da conexão. | S |
| RF-13.04 | Preservar a autoria das informações declaradas sobre o desfecho. | M |
| RF-13.05 | Permitir identificar divergências entre informações declaradas pelas partes. | S |
| RF-13.06 | Utilizar informações de desfecho para determinar elegibilidade de avaliações conforme as regras aplicáveis. | M |

---

## RF-14 — Avaliações e reputação bilateral

**Objetivo:** disponibilizar mecanismos de construção de reputação contextual para clientes e profissionais.  
**Rastreabilidade:** OBJ-07; jornada 6.9.  
**Spec:** `trust-reputation`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-14.01 | Permitir avaliações exclusivamente em contextos considerados elegíveis pelas regras de negócio. | M |
| RF-14.02 | Permitir que clientes avaliem experiências com profissionais. | M |
| RF-14.03 | Permitir que profissionais avaliem experiências com clientes. | M |
| RF-14.04 | Permitir o registro de uma avaliação estruturada da experiência. | M |
| RF-14.05 | Permitir a inclusão opcional de comentário textual. | S |
| RF-14.06 | Associar avaliações ao contexto e ao papel desempenhado pelo usuário avaliado. | M |
| RF-14.07 | Impedir avaliações duplicadas ou autoavaliações indevidas. | M |
| RF-14.08 | Apresentar informações de reputação acompanhadas de contexto suficiente para sua interpretação. | M |
| RF-14.09 | Distinguir a reputação decorrente da atuação como cliente daquela decorrente da atuação profissional. | M |
| RF-14.10 | Permitir a denúncia de avaliações incompatíveis com as políticas da plataforma. | M |
| RF-14.11 | Permitir a resposta do usuário avaliado, conforme as políticas aplicáveis. | C |
| RF-14.12 | Permitir a correção ou retirada de avaliações nas condições estabelecidas pelas regras de negócio. | S |
| RF-14.13 | Identificar padrões potencialmente fraudulentos ou abusivos para análise operacional. | C |

---

## RF-15 — Confirmação e verificação de informações

**Objetivo:** disponibilizar sinais de confiança correspondentes exclusivamente a informações cuja confirmação tenha sido efetivamente realizada.  
**Rastreabilidade:** OBJ-07; jornadas 6.1, 6.2 e 6.4.  
**Specs:** `trust-reputation` + `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-15.01 | Registrar o estado das confirmações de informações de contato realizadas. | M |
| RF-15.02 | Apresentar sinais de confirmação correspondentes às informações efetivamente confirmadas. | M |
| RF-15.03 | Explicar o significado e as limitações dos sinais de confirmação apresentados. | M |
| RF-15.04 | Permitir verificações adicionais de informações profissionais, quando justificadas operacionalmente. | C |
| RF-15.05 | Permitir a atualização ou invalidação de verificações que deixem de satisfazer os critérios aplicáveis. | C |

**Limite:** confirmação de e-mail/telefone não é verificação integral de identidade; background check está fora do MVP.

---

## RF-16 — Notificações

**Objetivo:** informar os usuários sobre acontecimentos relevantes para suas interações com a plataforma.  
**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornadas 6.5, 6.6, 6.7 e 6.9.  
**Spec:** `service-request`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-16.01 | Notificar profissionais sobre novas solicitações de conexão. | M |
| RF-16.02 | Notificar clientes sobre manifestações de interesse recebidas. | M |
| RF-16.03 | Notificar usuários sobre a aceitação ou recusa de suas solicitações e manifestações de interesse. | M |
| RF-16.04 | Notificar as partes sobre o estabelecimento de uma conexão. | M |
| RF-16.05 | Informar acontecimentos relevantes relacionados à segurança da conta. | M |
| RF-16.06 | Notificar usuários sobre a possibilidade de registrar avaliações elegíveis. | S |
| RF-16.07 | Respeitar as preferências aplicáveis às comunicações não essenciais. | M |
| RF-16.08 | Disponibilizar um histórico interno de notificações. | S |

---

## RF-17 — Denúncias e proteção do usuário

**Objetivo:** permitir que usuários comuniquem irregularidades e controlem interações potencialmente indesejadas.  
**Rastreabilidade:** OBJ-06 e OBJ-07; jornada 6.10.  
**Spec:** `trust-reputation`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-17.01 | Permitir a denúncia de usuários e perfis profissionais. | M |
| RF-17.02 | Permitir a denúncia de necessidades e conteúdos de portfólio. | M |
| RF-17.03 | Permitir a denúncia de avaliações e interações originadas na plataforma. | M |
| RF-17.04 | Permitir a indicação de motivo e informações adicionais relevantes à denúncia. | M |
| RF-17.05 | Registrar denúncias para tratamento operacional. | M |
| RF-17.06 | Permitir o bloqueio de interações indesejadas com usuários específicos. | M |
| RF-17.07 | Permitir que o denunciante receba informações pertinentes sobre o tratamento do caso, respeitadas as restrições de privacidade. | S |
| RF-17.08 | Disponibilizar orientações claras sobre comportamentos incompatíveis com as políticas da plataforma. | M |

---

## RF-18 — Administração e moderação

**Objetivo:** disponibilizar capacidades operacionais necessárias à manutenção da integridade, segurança e funcionamento da plataforma.  
**Rastreabilidade:** OBJ-07 e OBJ-09; jornada 6.10.  
**Specs:** `trust-reputation` + `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-18.01 | Permitir que operadores autorizados consultem denúncias registradas. | M |
| RF-18.02 | Permitir a organização dos casos conforme sua natureza, estado e risco potencial. | M |
| RF-18.03 | Permitir a consulta às informações autorizadas necessárias à análise de um caso. | M |
| RF-18.04 | Permitir o registro de decisões e fundamentos de moderação. | M |
| RF-18.05 | Permitir a aplicação de advertências, restrições e suspensões previstas nas políticas. | M |
| RF-18.06 | Permitir a remoção ou indisponibilização de conteúdo incompatível com as políticas. | M |
| RF-18.07 | Permitir a revisão e reversão de medidas administrativas quando aplicável. | M |
| RF-18.08 | Registrar a autoria e o contexto das ações administrativas relevantes. | M |
| RF-18.09 | Permitir a gestão operacional do catálogo de serviços. | M |
| RF-18.10 | Permitir o tratamento de contestações apresentadas por usuários afetados. | M |
| RF-18.11 | Disponibilizar indicadores operacionais básicos. | S |
| RF-18.12 | Disponibilizar ferramentas especializadas de configuração e controle de funcionalidades. | C |

---

## RF-19 — Transparência e controle de dados pessoais

**Objetivo:** disponibilizar mecanismos funcionais para que usuários compreendam e exerçam controle sobre o tratamento de suas informações.  
**Rastreabilidade:** princípios de privacidade, segurança, autonomia e transparência.  
**Spec:** `foundation-platform`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-19.01 | Disponibilizar informações acessíveis sobre o tratamento de dados pessoais. | M |
| RF-19.02 | Permitir a consulta das informações pessoais disponibilizadas pelo usuário. | M |
| RF-19.03 | Permitir a atualização ou solicitação de correção de informações pessoais. | M |
| RF-19.04 | Permitir a solicitação de acesso e fornecimento de dados pessoais, conforme as regras legais aplicáveis. | M |
| RF-19.05 | Permitir a solicitação de eliminação de dados pessoais, observadas as hipóteses legais de retenção. | M |
| RF-19.06 | Permitir o gerenciamento e a revogação de consentimentos quando essa for a base legal aplicável. | M |
| RF-19.07 | Registrar solicitações relacionadas ao exercício de direitos dos titulares. | M |
| RF-19.08 | Permitir o tratamento das solicitações por operadores autorizados. | M |
| RF-19.09 | Permitir que o usuário identifique quais informações de seu perfil são disponibilizadas publicamente. | M |

---

## RF-20 — Medição e validação do produto

**Objetivo:** disponibilizar informações suficientes para avaliar o funcionamento do produto e validar suas hipóteses fundamentais.  
**Rastreabilidade:** OBJ-01 a OBJ-09; VAL-01 a VAL-07.  
**Specs:** `fair-visibility` + `local-discovery`.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-20.01 | Registrar acontecimentos relevantes dos fluxos de cadastro e publicação de perfis. | M |
| RF-20.02 | Registrar acontecimentos relevantes dos mecanismos de descoberta. | M |
| RF-20.03 | Registrar acontecimentos relacionados à publicação e utilização de necessidades. | M |
| RF-20.04 | Registrar acontecimentos relacionados a solicitações, interesses e conexões. | M |
| RF-20.05 | Registrar acontecimentos relacionados a desfechos e avaliações elegíveis. | M |
| RF-20.06 | Disponibilizar informações necessárias à análise da distribuição de exposição profissional. | M |
| RF-20.07 | Permitir a obtenção de indicadores agregados de oferta, demanda e interação. | M |
| RF-20.08 | Permitir a coleta voluntária de informações qualitativas sobre a experiência dos usuários. | S |
| RF-20.09 | Disponibilizar mecanismos especializados de visualização de indicadores. | C |
| RF-20.10 | Disponibilizar infraestrutura especializada de experimentação controlada. | C |

---

## Matriz de prioridade e ciclo

| Ciclo | RFs must | RFs should/could |
| --- | --- | --- |
| Fatia técnica 1 | RF-01, RF-02, RF-03, RF-05, RF-06, RF-07 (guardrails), RF-08, RF-09, RF-10, RF-11, RF-12, RF-13, RF-14, RF-15 (e-mail), RF-17, RF-18, RF-19, RF-20 (mínimo) | RF-04 (portfólio básico), RF-16 (in-app) |
| Piloto | todos os M | S conforme evidência |
| Pós-piloto | S/C priorizados por evidência | — |

> A fatia técnica 1 é uma decisão de engenharia para validar o fluxo; não substitui o escopo de validação do piloto.
