# PRD — RAIRE: Descoberta Local Justa

**Status:** rascunho para revisão humana  
**Versão:** 0.1  
**Data:** 2026-11-20  
**Fase:** Fundação / MVP conceitual  
**Responsável de produto:** a designar  
**Aprovadores:** fundador, pesquisa, arquitetura, segurança e privacidade  

> Este PRD organiza o que o RAIRE pretende validar e construir. Ele **não autoriza implementação** por si só. Toda feature nasce de uma spec em `specs/`, com status `aprovada` apenas após revisão humana. Hipóteses continuam sendo hipóteses até existir evidência.

---

# 1. Sumário executivo

O RAIRE é uma plataforma de descoberta e conexão entre profissionais autônomos e pessoas que buscam serviços em suas regiões, concebida para facilitar o acesso a serviços locais, ampliar oportunidades de trabalho e estabelecer relações de maior confiança entre clientes e profissionais.

Atualmente, a busca por serviços locais é marcada pela fragmentação de informações e pela dependência de indicações pessoais, redes sociais e outros mecanismos informais. Clientes enfrentam dificuldades para identificar profissionais próximos, conhecer seus trabalhos, consultar sua disponibilidade e avaliar sua reputação antes de estabelecer contato.

Paralelamente, profissionais autônomos encontram dificuldades para divulgar seus serviços, alcançar novos clientes e identificar oportunidades de trabalho em suas regiões. A dependência de redes de relacionamento e plataformas generalistas limita sua exposição e torna a aquisição de clientes menos previsível, especialmente para aqueles que ainda não possuem uma presença digital consolidada.

O RAIRE propõe reduzir essas barreiras por meio de um ambiente digital especializado na prestação de serviços locais, proporcionando:

- **Descoberta geolocalizada:** identificação de profissionais e oportunidades de trabalho com base na localização, área de atendimento, categoria de serviço e disponibilidade.
- **Portfólio profissional especializado:** espaço dedicado à apresentação de serviços, trabalhos realizados, experiências e informações relevantes para a decisão do cliente.
- **Conexão direta:** possibilidade de encontrar profissionais e estabelecer contato mediante interesse e consentimento entre as partes.
- **Publicação de demandas:** possibilidade de clientes divulgarem suas necessidades para que profissionais da região possam identificar oportunidades e manifestar interesse na execução dos serviços.
- **Reputação bilateral:** mecanismos de avaliação de profissionais e clientes, contribuindo para a construção de confiança e a redução da assimetria de informações.

A proposta do RAIRE combina essas funcionalidades com princípios de distribuição justa de visibilidade, transparência, explicabilidade, segurança e privacidade desde sua concepção. A plataforma busca evitar que oportunidades sejam determinadas exclusivamente pela popularidade ou capacidade de investimento dos profissionais, preservando a relevância dos resultados apresentados aos clientes.

O MVP será orientado a uma experiência simples, prioritariamente móvel, contemplando a criação de perfis profissionais, divulgação de serviços e portfólios, descoberta geolocalizada, publicação de demandas, estabelecimento de conexões e construção de reputação bilateral.

Nesta primeira versão, a plataforma não oferecerá mecanismos pagos de favorecimento no ranking, não utilizará atributos pessoais sensíveis para classificação de profissionais e não apresentará garantias de segurança ou verificações que não consiga efetivamente realizar.

**Resultado esperado do MVP:** validar, em uma região piloto e em um conjunto limitado de categorias de serviços, a capacidade da plataforma de facilitar a descoberta e conexão entre clientes e profissionais, ampliar a exposição profissional e o acesso a oportunidades de trabalho, reduzir as dificuldades de contratação e promover relações de maior confiança, preservando a relevância, a equidade, a segurança e a privacidade dos usuários.

---

# 2. Contexto e problema

## 2.1 Contexto

O mercado local de prestação de serviços é composto por profissionais autônomos que oferecem competências, conhecimentos e experiências em diferentes categorias, atendendo às necessidades de pessoas e organizações em suas respectivas regiões.

A descoberta desses profissionais e a aquisição de clientes ocorrem frequentemente por meio de indicações pessoais, redes sociais, grupos de mensagens, anúncios e plataformas digitais generalistas. Embora esses canais possibilitem conexões, suas informações e funcionalidades encontram-se distribuídas entre diferentes ambientes, nem sempre especializados nas necessidades do mercado de serviços locais.

Para os clientes, essa fragmentação pode dificultar a identificação de profissionais próximos, a compreensão dos serviços oferecidos, a consulta de disponibilidade e a comparação de alternativas com base em trabalhos realizados, reputação e outros critérios relevantes.

Para os profissionais autônomos, a ausência de canais especializados e acessíveis pode dificultar a divulgação de seus trabalhos, a construção de uma presença digital profissional, a identificação de demandas locais e a aquisição recorrente de clientes.

Nesse contexto, existe uma oportunidade de investigar como uma plataforma digital especializada pode reduzir as barreiras de descoberta, divulgação e conexão entre oferta e demanda no mercado local de serviços.

## 2.2 Definição do problema

Os problemas apresentados nesta seção constituem hipóteses iniciais, fundamentadas na análise preliminar do domínio, e deverão ser validados com os públicos envolvidos.

### 2.2.1 Perspectiva do cliente

Pessoas que necessitam de serviços locais podem enfrentar dificuldades relacionadas a:

- Identificar profissionais que atuem em sua região e atendam às suas necessidades específicas.
- Obter uma visão abrangente da oferta de serviços disponível localmente.
- Consultar informações sobre disponibilidade, especialidades e áreas de atendimento.
- Conhecer trabalhos anteriormente realizados pelos profissionais por meio de portfólios especializados.
- Comparar profissionais com base em critérios relevantes, como experiência, reputação e adequação ao serviço solicitado.
- Avaliar a confiabilidade de profissionais fora de seus círculos sociais.
- Estabelecer contato de maneira prática, segura e respeitosa à privacidade.
- Divulgar necessidades específicas e receber manifestações de interesse de profissionais disponíveis, sem precisar procurar individualmente por cada prestador.

### 2.2.2 Perspectiva do profissional autônomo

Profissionais autônomos podem enfrentar dificuldades relacionadas a:

- Alcançar potenciais clientes em suas regiões de atuação.
- Divulgar seus serviços e trabalhos realizados em ambientes adequados à sua atividade profissional.
- Construir uma presença digital sem depender exclusivamente de redes sociais generalistas.
- Identificar oportunidades de trabalho compatíveis com suas competências, localização e disponibilidade.
- Reduzir a dependência de indicações pessoais e círculos de relacionamento preexistentes.
- Competir por oportunidades sem depender exclusivamente de popularidade, alcance social ou investimento em publicidade.
- Construir uma reputação profissional baseada em experiências reais de prestação de serviços.
- Avaliar a confiabilidade de potenciais clientes antes de estabelecer uma relação comercial.
- Utilizar ferramentas digitais sem enfrentar barreiras excessivas de complexidade, custo ou familiaridade tecnológica.
- Ampliar suas oportunidades de geração de renda de maneira segura e sustentável.

### 2.2.3 Perspectiva do mercado local de serviços

A interação entre oferta e demanda pode apresentar problemas estruturais, entre os quais:

- **Fragmentação da informação:** dados sobre profissionais, serviços e demandas encontram-se distribuídos entre diferentes canais.
- **Assimetria informacional:** clientes e profissionais possuem informações limitadas para avaliar a adequação e a confiabilidade da outra parte.
- **Baixa eficiência de descoberta:** a existência de profissionais disponíveis e clientes interessados em uma mesma região não garante que essas partes consigam se encontrar.
- **Concentração de visibilidade:** oportunidades podem favorecer profissionais com maior exposição digital, independentemente de sua adequação à demanda.
- **Dependência de mecanismos informais:** indicações pessoais e redes de relacionamento podem limitar o alcance das conexões.
- **Barreiras de acessibilidade digital:** ferramentas complexas podem dificultar a participação de profissionais e clientes com diferentes níveis de familiaridade tecnológica.

## 2.3 Consequências do problema

As dificuldades identificadas podem produzir impactos tanto para os participantes quanto para o funcionamento do mercado local de serviços.

Para os clientes, esses impactos podem incluir:

- Maior tempo e esforço para encontrar profissionais adequados.
- Menor diversidade de alternativas disponíveis para comparação.
- Insegurança durante a escolha e o estabelecimento de contato.
- Contratações realizadas com informações insuficientes.

Para os profissionais autônomos, esses impactos podem incluir:

- Menor exposição a potenciais clientes.
- Perda de oportunidades de trabalho.
- Maior esforço de divulgação e aquisição de clientes.
- Imprevisibilidade na geração de renda.
- Dificuldade de estabelecer e consolidar reputação profissional.

Para o mercado local, a consequência potencial é um aproveitamento ineficiente da oferta e da demanda existentes, em que necessidades de serviços permanecem sem atendimento enquanto profissionais qualificados e disponíveis encontram dificuldades para acessar essas oportunidades.

## 2.4 Pergunta central

Como facilitar a descoberta e a conexão entre clientes e profissionais autônomos no mercado local de serviços, reduzindo as barreiras de acesso, divulgação e confiança, e ampliando as oportunidades de trabalho de maneira eficiente, acessível, segura e equitativa?

A partir dessa pergunta central, surgem questões complementares:

- Como facilitar a identificação de profissionais adequados às necessidades dos clientes?
- Como ampliar o acesso dos profissionais a oportunidades de trabalho em suas regiões?
- Como oferecer informações relevantes para a tomada de decisão de ambas as partes?
- Como construir mecanismos de reputação que promovam confiança bilateral?
- Como distribuir visibilidade sem comprometer a relevância dos resultados?
- Como reduzir as barreiras de participação para usuários com diferentes níveis de familiaridade digital?

## 2.5 Evidências necessárias para validação

### 2.5.1 Validação do problema

- Realizar entrevistas com clientes e profissionais autônomos de diferentes categorias de serviços.
- Identificar os canais atualmente utilizados para divulgação, descoberta e contratação de serviços.
- Investigar as principais dificuldades enfrentadas por cada público durante essas jornadas.
- Avaliar a frequência, intensidade e consequências dos problemas identificados.
- Investigar a dependência de indicações pessoais e redes sociais na aquisição de clientes.

### 2.5.2 Validação de mercado

- Identificar plataformas concorrentes diretas e indiretas.
- Analisar as funcionalidades, propostas de valor e limitações das soluções existentes.
- Investigar como essas plataformas promovem a descoberta local e distribuem visibilidade entre profissionais.
- Identificar necessidades que não sejam adequadamente atendidas pelas soluções existentes.

### 2.5.3 Validação da proposta de valor

- Avaliar o interesse dos clientes em ferramentas especializadas de descoberta de profissionais locais.
- Avaliar o interesse dos profissionais em ferramentas especializadas de divulgação e aquisição de clientes.
- Validar a utilidade de portfólios profissionais estruturados.
- Avaliar a demanda pelos dois mecanismos de conexão: busca direta por profissionais e publicação de necessidades de serviços.
- Investigar a aceitação de mecanismos de reputação bilateral.

### 2.5.4 Validação dos mecanismos do produto

- Investigar como localização, disponibilidade, especialidade, portfólio e reputação influenciam a tomada de decisão.
- Avaliar mecanismos de distribuição equitativa de visibilidade sem comprometer a relevância dos resultados.
- Validar a compreensão das informações e explicações apresentadas aos usuários.
- Avaliar a usabilidade da plataforma para diferentes níveis de familiaridade digital.
- Identificar riscos relacionados a fraudes, avaliações abusivas, discriminação, exposição de dados pessoais e segurança das conexões.

---

## 3. Visão, missão e princípios

# 3. Visão, missão e princípios

## 3.1 Visão

Construir um ecossistema de serviços locais acessível, eficiente e confiável, no qual pessoas possam encontrar profissionais adequados às suas necessidades e profissionais autônomos possam demonstrar suas competências, construir reputação e ampliar suas oportunidades de trabalho, independentemente de seus círculos sociais ou de sua exposição digital preexistente.

## 3.2 Missão

Facilitar a descoberta e a conexão entre clientes e profissionais autônomos, reduzindo barreiras de acesso, divulgação e confiança no mercado local de serviços.

## 3.3 Missão do MVP

Validar a capacidade de uma plataforma digital especializada em serviços locais de facilitar a descoberta de profissionais, ampliar o acesso a oportunidades de trabalho e promover conexões relevantes entre clientes e profissionais autônomos.

A validação ocorrerá em um contexto geográfico e operacional delimitado, por meio de uma experiência simples, acessível e orientada à construção de confiança entre as partes.

O MVP deverá produzir evidências sobre a utilidade da plataforma para ambos os públicos, a capacidade de estabelecer conexões relevantes e a viabilidade de distribuir oportunidades de maneira equitativa, sem comprometer a qualidade dos resultados ou a segurança dos usuários.

## 3.4 Princípios de produto

### Geração de valor bilateral

O produto deve atender às necessidades de clientes e profissionais, reconhecendo que sua sustentabilidade depende da geração de valor para ambas as partes.

### Relevância local

A descoberta de profissionais e oportunidades deve considerar o contexto geográfico, as características do serviço, a área de atendimento e a disponibilidade das partes, quando aplicáveis.

### Descoberta justa

A plataforma deve reduzir barreiras de exposição e evitar concentrações excessivas de visibilidade, preservando a relevância dos resultados e sem prometer igualdade artificial de oportunidades ou resultados.

### Autonomia dos usuários

Clientes e profissionais devem possuir controle sobre suas decisões, informações, disponibilidade, interesses e conexões, respeitando os direitos e a privacidade das demais partes.

### Participação ativa

A plataforma deve permitir que profissionais e clientes assumam papéis ativos na descoberta de oportunidades e no estabelecimento de conexões, sem restringir a interação a um único fluxo de oferta e demanda.

### Conexão antes de engajamento

O sucesso do produto deve ser orientado pela geração de conexões úteis e relevantes, e não pela maximização do tempo de permanência, da frequência de acesso ou de interações sem valor efetivo.

### Simplicidade e redução de fricção

A experiência deve minimizar etapas, decisões desnecessárias e barreiras operacionais, permitindo que os usuários alcancem seus objetivos com clareza e eficiência.

### Confiança bilateral

A construção de confiança deve considerar tanto profissionais quanto clientes, utilizando informações relevantes, histórico de interações, reputação, transparência e mecanismos de consentimento.

A plataforma não deverá oferecer garantias ou alegações de verificação que não consiga efetivamente sustentar.

### Reputação contextual

A reputação deve considerar as características e o contexto das experiências avaliadas, reconhecendo que uma classificação geral não representa necessariamente a qualidade de todos os serviços ou interações.

### Transparência e explicabilidade

Os usuários devem compreender os principais critérios que influenciam os resultados apresentados e as decisões automatizadas que os afetam, em linguagem adequada ao contexto de utilização.

### Acessibilidade e inclusão digital

A plataforma deve ser concebida para atender usuários com diferentes capacidades, condições de acesso à tecnologia e níveis de familiaridade digital.

### Privacidade e segurança por padrão

A coleta, o processamento e o compartilhamento de informações devem respeitar os princípios de necessidade, finalidade, minimização de dados, consentimento e proteção da privacidade dos usuários.

## 3.5 Restrições fundamentais do MVP

- Priorizar uma experiência mobile-first.
- Limitar a operação inicial a uma região piloto e a um conjunto reduzido de categorias de serviços.
- Não permitir favorecimento pago no posicionamento de profissionais nos resultados de descoberta.
- Não utilizar atributos pessoais sensíveis como critérios de classificação ou distribuição de visibilidade.
- Não expor informações pessoais de contato sem autorização.
- Não divulgar localização residencial exata como requisito de descoberta.
- Não oferecer garantias de qualidade, identidade ou segurança que não possam ser sustentadas pelos mecanismos disponíveis na plataforma.

## 3.6 Princípios de desenvolvimento e validação

### Desenvolvimento orientado por evidências

Hipóteses de produto não devem ser tratadas como fatos. Decisões relevantes devem ser fundamentadas em evidências, acompanhadas por critérios de validação e revisadas diante de novos aprendizados.

### Evolução incremental

O produto deve evoluir progressivamente, priorizando funcionalidades capazes de validar suas hipóteses fundamentais antes da introdução de mecanismos de maior complexidade.

### Rastreabilidade das decisões

Requisitos, hipóteses, decisões de produto, decisões arquiteturais, evidências e aprendizados devem ser documentados e mantidos de forma rastreável ao longo da evolução do projeto.

---

# 4. Objetivos e não objetivos

## 4.1 Objetivo geral do MVP

Validar a capacidade do RAIRE de facilitar a descoberta e a conexão entre clientes e profissionais autônomos no mercado local de serviços, proporcionando acesso a informações relevantes, oportunidades de trabalho e mecanismos de construção de confiança entre as partes.

O MVP deverá produzir evidências sobre a utilidade da plataforma para ambos os públicos, a eficiência dos mecanismos de descoberta e conexão e a viabilidade de estabelecer um ecossistema local de oferta e demanda.

## 4.2 Objetivos de produto

| ID | Objetivo | Indicadores de sucesso |
| --- | --- | --- |
| OBJ-01 | Permitir que profissionais construam uma presença digital especializada para divulgação de seus serviços e competências. | Taxa de conclusão do cadastro; tempo até publicação do perfil; proporção de perfis com serviços e informações relevantes preenchidos. |
| OBJ-02 | Permitir que profissionais apresentem trabalhos realizados por meio de portfólios especializados. | Taxa de publicação de portfólios; utilização dos portfólios durante a descoberta; utilidade percebida pelos clientes. |
| OBJ-03 | Facilitar a descoberta de profissionais adequados às necessidades dos clientes em suas regiões. | Taxa de buscas com resultados relevantes; tempo até identificação de um profissional adequado; satisfação com os resultados. |
| OBJ-04 | Permitir que clientes publiquem necessidades específicas de serviços e recebam manifestações de interesse. | Taxa de demandas que recebem interesse relevante; tempo até a primeira manifestação de interesse; satisfação com as alternativas recebidas. |
| OBJ-05 | Ampliar a capacidade dos profissionais de identificar e acessar oportunidades de trabalho em suas regiões. | Proporção de profissionais que encontram oportunidades relevantes; taxa de manifestação de interesse; oportunidades que resultam em conexões. |
| OBJ-06 | Facilitar o estabelecimento de conexões relevantes e consentidas entre clientes e profissionais. | Taxa de conversão de interesses em conexões; quantidade de conexões estabelecidas; tempo até a conexão; utilidade percebida pelas partes. |
| OBJ-07 | Promover a construção de confiança bilateral por meio de informações relevantes e mecanismos de reputação. | Confiança percebida pelos usuários; cobertura de avaliações elegíveis; utilidade das avaliações na tomada de decisão; incidência de reclamações. |
| OBJ-08 | Reduzir barreiras de exposição profissional sem comprometer a relevância dos resultados apresentados. | Distribuição de exposição entre profissionais elegíveis e comparáveis; exposição de novos profissionais; satisfação dos clientes com os resultados. |
| OBJ-09 | Validar a capacidade de concentrar oferta e demanda suficientes em uma região piloto. | Quantidade de profissionais ativos por categoria; volume de demandas locais; taxa de atendimento das demandas por conexões relevantes; recorrência de utilização. |

Os indicadores representam instrumentos iniciais de avaliação. Suas definições operacionais, métodos de coleta e metas quantitativas deverão ser estabelecidos no planejamento de validação do MVP.

## 4.3 Objetivos de pesquisa e validação

| ID | Objetivo | Evidência esperada |
| --- | --- | --- |
| VAL-01 | Validar a relevância dos problemas identificados para clientes e profissionais. | Entrevistas, pesquisas e observações das jornadas existentes. |
| VAL-02 | Validar a proposta de valor bilateral do RAIRE. | Evidências de utilidade, intenção de uso e utilização efetiva por ambos os públicos. |
| VAL-03 | Avaliar a efetividade dos dois mecanismos de conexão: descoberta direta e publicação de demandas. | Dados de utilização, conversão e satisfação de cada fluxo. |
| VAL-04 | Investigar os critérios que influenciam a descoberta e a seleção de profissionais e oportunidades. | Evidências sobre a influência de localização, disponibilidade, especialidade, portfólio e reputação. |
| VAL-05 | Avaliar a efetividade dos mecanismos de construção de confiança. | Evidências sobre confiança percebida, utilidade das informações e comportamento dos usuários. |
| VAL-06 | Avaliar mecanismos de distribuição equitativa de visibilidade. | Dados de distribuição de exposição e sua relação com a relevância dos resultados. |
| VAL-07 | Identificar barreiras de acessibilidade, usabilidade e adoção. | Testes de usabilidade com usuários de diferentes níveis de familiaridade digital. |

A coleta de evidências deverá respeitar os princípios de privacidade, transparência, finalidade e minimização de dados.

## 4.4 Objetivos de qualidade

| ID | Dimensão | Objetivo |
| --- | --- | --- |
| QLT-01 | Usabilidade | Permitir a conclusão dos principais fluxos com baixa complexidade e esforço operacional. |
| QLT-02 | Acessibilidade | Adotar as diretrizes WCAG 2.2 AA aplicáveis à experiência web do MVP. |
| QLT-03 | Responsividade | Garantir a utilização adequada dos fluxos principais em dispositivos móveis, incluindo telas de dimensões reduzidas. |
| QLT-04 | Desempenho | Proporcionar tempos de resposta e carregamento adequados em condições limitadas de conectividade e capacidade computacional. |
| QLT-05 | Privacidade | Minimizar a coleta e exposição de dados pessoais, preservando o controle dos usuários sobre o compartilhamento de informações. |
| QLT-06 | Segurança | Proteger contas, informações e interações contra acessos indevidos, abusos e ameaças compatíveis com o contexto operacional. |
| QLT-07 | Transparência | Apresentar informações claras sobre os critérios de descoberta, mecanismos de reputação e limitações da plataforma. |

Os critérios técnicos de aceitação deverão ser detalhados na seção de requisitos não funcionais.

## 4.5 Não objetivos do MVP

### 4.5.1 Escopo de mercado

- Substituir redes sociais generalistas ou tornar-se uma plataforma de produção e consumo contínuo de conteúdo.
- Tornar-se uma plataforma genérica de contratação de freelancers em escala nacional ou internacional.
- Atender inicialmente a todas as categorias de serviços.
- Expandir a operação para múltiplas regiões antes da validação da densidade local de oferta e demanda.

### 4.5.2 Intermediação comercial e financeira

- Processar ou intermediar pagamentos entre clientes e profissionais.
- Cobrar comissões sobre serviços realizados.
- Reter valores, oferecer custódia financeira ou administrar disputas financeiras.
- Gerenciar contratos, obrigações fiscais ou relações trabalhistas entre as partes.
- Implementar mecanismos de monetização que favoreçam artificialmente profissionais nos resultados de descoberta.

### 4.5.3 Gestão da prestação de serviços

- Oferecer ferramentas completas de gerenciamento de projetos ou execução de serviços.
- Disponibilizar sistemas avançados de agendamento e gestão de calendário.
- Implementar videoconferência ou comunicação em tempo real como capacidade central do produto.
- Substituir ferramentas especializadas de gestão profissional.

### 4.5.4 Garantias e responsabilidade

- Garantir a qualidade ou a conclusão dos serviços contratados.
- Realizar investigações de antecedentes pessoais ou profissionais.
- Oferecer seguros ou garantias financeiras.
- Apresentar profissionais como integralmente verificados sem mecanismos que sustentem essa afirmação.
- Eliminar completamente os riscos inerentes às relações entre clientes e profissionais.

### 4.5.5 Complexidade tecnológica

- Desenvolver aplicativos móveis nativos antes da validação da experiência web.
- Implementar mecanismos avançados de recomendação personalizada antes de validar os critérios fundamentais de relevância.
- Utilizar inteligência artificial generativa como autoridade autônoma para decisões de classificação, reputação ou moderação que afetem usuários.
- Implementar mecanismos sofisticados de distribuição de visibilidade sem evidências que justifiquem sua necessidade.

Os não objetivos delimitam o escopo inicial do MVP e não constituem necessariamente restrições permanentes à evolução do RAIRE.

---

# 5. Personas e atores

## 5.1 Atores do sistema

Os atores representam diferentes formas de interação com o RAIRE, considerando suas responsabilidades, necessidades e capacidades dentro da plataforma.

Os papéis de cliente e profissional não são mutuamente exclusivos. Uma mesma pessoa poderá utilizar a plataforma para buscar serviços e, mediante a criação de um perfil profissional, divulgar seus próprios serviços e acessar oportunidades de trabalho.

### 5.1.1 Atores de mercado

| Ator | Descrição | Capacidades principais |
| --- | --- | --- |
| Visitante | Pessoa que acessa a plataforma sem autenticação. | Buscar serviços e profissionais, consultar informações públicas e iniciar cadastro. |
| Cliente | Usuário que busca profissionais ou possui necessidades relacionadas à prestação de serviços. | Buscar e comparar profissionais, consultar portfólios, publicar e gerenciar demandas, receber manifestações de interesse, estabelecer conexões, avaliar profissionais e denunciar irregularidades. |
| Profissional | Usuário que oferece serviços de forma autônoma. | Criar e gerenciar perfil profissional, declarar serviços, disponibilidade e área de atendimento, publicar portfólio, buscar oportunidades, manifestar interesse em demandas, receber solicitações de conexão, avaliar clientes e denunciar irregularidades. |

### 5.1.2 Atores operacionais

| Ator | Descrição | Capacidades principais |
| --- | --- | --- |
| Operador | Responsável pelas atividades operacionais e de suporte da plataforma. | Analisar denúncias, moderar conteúdo, aplicar medidas previstas nas políticas da plataforma e gerenciar informações operacionais autorizadas. |
| Pesquisador | Responsável pela análise de produto e validação de hipóteses. | Consultar métricas agregadas, analisar experimentos e acessar dados de pesquisa adequadamente anonimizados, conforme as políticas de privacidade. |
| Administrador | Responsável pela administração técnica e operacional autorizada da plataforma. | Gerenciar configurações, políticas, permissões e recursos administrativos, respeitando os princípios de menor privilégio e rastreabilidade. |

Os atores operacionais representam responsabilidades conceituais. A implementação de papéis administrativos independentes dependerá das necessidades operacionais do MVP.

As capacidades descritas nesta seção não constituem uma especificação definitiva de autorização. As regras de acesso deverão ser detalhadas nos requisitos funcionais e de segurança.

## 5.2 Personas primárias

As personas representam hipóteses iniciais sobre perfis comportamentais dos públicos do RAIRE.

Sua definição considera necessidades, objetivos, dificuldades, motivações e contextos de utilização, devendo ser revisada a partir de entrevistas, pesquisas e observações com usuários reais.

Uma mesma pessoa poderá apresentar características de diferentes personas conforme seu contexto.

### P1 — Cliente com necessidade imediata

**Contexto**

Possui uma necessidade de serviço com baixa tolerância a atrasos, como reparos elétricos, hidráulicos ou manutenção emergencial.

**Objetivos**

- Encontrar profissionais que atendam à sua região.
- Identificar disponibilidade compatível com sua necessidade.
- Estabelecer contato rapidamente.
- Obter informações suficientes para tomar uma decisão com confiança.

**Dificuldades**

- Informações desatualizadas ou incompletas.
- Dificuldade de identificar profissionais disponíveis.
- Dependência de indicações pessoais.
- Incerteza quanto à confiabilidade do prestador.

**Critérios de sucesso**

- Encontrar alternativas relevantes com baixo esforço.
- Identificar disponibilidade e área de atendimento.
- Compreender claramente como estabelecer contato.
- Preservar informações pessoais até decidir compartilhá-las.

### P2 — Cliente planejador

**Contexto**

Possui uma necessidade de serviço cujo planejamento permite avaliar diferentes profissionais antes de tomar uma decisão.

**Objetivos**

- Obter uma visão abrangente das alternativas disponíveis na região.
- Conhecer trabalhos anteriormente realizados.
- Comparar especialidades, experiência, reputação e disponibilidade.
- Publicar necessidades específicas quando a busca direta não for suficiente.

**Dificuldades**

- Informações distribuídas entre diferentes plataformas.
- Portfólios desestruturados ou de difícil localização.
- Dificuldade de comparar profissionais.
- Dependência de avaliações genéricas ou indicações pessoais.

**Critérios de sucesso**

- Encontrar profissionais adequados à necessidade.
- Comparar informações relevantes de diferentes prestadores.
- Tomar decisões fundamentadas em sinais compreensíveis de confiança.
- Receber manifestações de interesse compatíveis com suas demandas.

### P3 — Profissional com baixa exposição digital

**Contexto**

Possui competências e experiência na prestação de serviços, mas dispõe de presença digital limitada ou depende predominantemente de indicações pessoais e ferramentas informais de comunicação.

**Objetivos**

- Tornar seus serviços conhecidos por clientes da região.
- Apresentar competências e trabalhos realizados.
- Construir reputação profissional.
- Ampliar suas oportunidades de trabalho.

**Dificuldades**

- Alcance limitado de divulgação.
- Dependência de círculos de relacionamento preexistentes.
- Possíveis barreiras de familiaridade digital.
- Dificuldade de competir com profissionais que possuem exposição consolidada.

**Critérios de sucesso**

- Criar e manter um perfil profissional sem complexidade excessiva.
- Apresentar seus serviços e trabalhos adequadamente.
- Receber contatos e oportunidades compatíveis.
- Compreender os principais fatores que influenciam sua visibilidade.

### P4 — Profissional em crescimento

**Contexto**

Já utiliza canais digitais para divulgar seus serviços e busca ampliar sua carteira de clientes, consolidar reputação e profissionalizar sua presença digital.

**Objetivos**

- Alcançar novos clientes na região.
- Apresentar um portfólio profissional especializado.
- Encontrar oportunidades compatíveis com seus serviços.
- Construir e acompanhar sua reputação.
- Reduzir a dependência de mecanismos generalistas de divulgação.

**Dificuldades**

- Concorrência por atenção em redes sociais.
- Fragmentação dos canais de aquisição de clientes.
- Recebimento de contatos pouco relevantes.
- Dificuldade de identificar quais esforços de divulgação geram resultados.

**Critérios de sucesso**

- Receber oportunidades relevantes.
- Estabelecer conexões com potenciais clientes.
- Construir reputação baseada em experiências reais.
- Identificar informações que contribuam para melhorar sua apresentação profissional.

### P5 — Profissional em busca ativa de oportunidades

**Contexto**

Possui disponibilidade para executar serviços e busca ativamente novas demandas, seja para ampliar sua renda, preencher períodos de ociosidade ou manter regularidade na aquisição de clientes.

**Objetivos**

- Identificar demandas disponíveis na região.
- Encontrar oportunidades compatíveis com suas competências e disponibilidade.
- Avaliar informações sobre o serviço e o potencial cliente.
- Manifestar interesse de maneira simples e segura.

**Dificuldades**

- Ausência de canais centralizados de oportunidades locais.
- Dependência de indicações e contatos espontâneos.
- Informações insuficientes sobre demandas.
- Risco de investir tempo em oportunidades incompatíveis ou pouco confiáveis.

**Critérios de sucesso**

- Encontrar demandas relevantes.
- Compreender as características da oportunidade antes de estabelecer contato.
- Manifestar interesse com baixo esforço.
- Converter oportunidades em conexões úteis.

## 5.3 Personas operacionais

### P6 — Operador de confiança e segurança

**Contexto**

Responsável pela manutenção da integridade das interações, análise de denúncias e aplicação das políticas da plataforma.

**Objetivos**

- Identificar comportamentos abusivos, fraudulentos ou incompatíveis com as políticas.
- Analisar denúncias com informações suficientes.
- Aplicar medidas proporcionais e fundamentadas.
- Preservar a segurança e os direitos dos usuários.

**Dificuldades**

- Informações incompletas ou relatos conflitantes.
- Necessidade de proteger dados pessoais durante investigações.
- Risco de decisões inconsistentes ou arbitrárias.
- Possibilidade de utilização abusiva dos mecanismos de denúncia.

**Critérios de sucesso**

- Analisar casos com critérios claros.
- Aplicar medidas proporcionais e rastreáveis.
- Preservar o histórico das decisões.
- Comunicar resultados às partes quando aplicável, respeitando a privacidade e a segurança.

## 5.4 Jobs to Be Done (JTBD)

Os Jobs to Be Done representam os resultados que os usuários procuram alcançar em determinados contextos, independentemente da implementação tecnológica utilizada pelo RAIRE.

### 5.4.1 Cliente

| ID | Job to Be Done |
| --- | --- |
| JTBD-C01 | Quando preciso de um serviço, quero encontrar profissionais que atendam à minha região, para não depender exclusivamente de indicações pessoais. |
| JTBD-C02 | Quando encontro diferentes profissionais, quero conhecer seus trabalhos, competências e reputações, para tomar uma decisão mais fundamentada. |
| JTBD-C03 | Quando possuo uma necessidade específica, quero comunicá-la a profissionais potencialmente interessados, para encontrar alternativas sem precisar procurar cada prestador individualmente. |
| JTBD-C04 | Quando identifico um profissional adequado, quero estabelecer contato de maneira simples e segura, para avançar na negociação do serviço. |
| JTBD-C05 | Quando vivencio uma experiência de prestação de serviço, quero registrar minha avaliação, para contribuir com a reputação do profissional e auxiliar outros clientes. |

### 5.4.2 Profissional

| ID | Job to Be Done |
| --- | --- |
| JTBD-P01 | Quando ofereço serviços, quero ser encontrado por potenciais clientes da minha região, para ampliar minhas oportunidades de trabalho. |
| JTBD-P02 | Quando divulgo meu trabalho, quero apresentar minhas competências e experiências de maneira organizada, para demonstrar meu valor profissional. |
| JTBD-P03 | Quando tenho disponibilidade para trabalhar, quero encontrar demandas compatíveis com meus serviços, para reduzir períodos de ociosidade e ampliar minha renda. |
| JTBD-P04 | Quando identifico uma oportunidade, quero conhecer suas características e informações relevantes sobre o cliente, para avaliar se vale a pena estabelecer contato. |
| JTBD-P05 | Quando realizo um serviço, quero construir uma reputação baseada em experiências reais, para fortalecer minha credibilidade profissional. |
| JTBD-P06 | Quando interajo com um cliente, quero poder avaliar a experiência, para contribuir com a construção de confiança entre profissionais e clientes. |

### 5.4.3 Necessidades transversais

| ID | Job to Be Done |
| --- | --- |
| JTBD-T01 | Quando estabeleço uma conexão, quero controlar o compartilhamento das minhas informações pessoais, para preservar minha privacidade. |
| JTBD-T02 | Quando utilizo mecanismos de descoberta, quero compreender os critérios relevantes apresentados pela plataforma, para tomar decisões conscientes. |
| JTBD-T03 | Quando enfrento uma experiência abusiva ou irregular, quero encontrar um mecanismo de denúncia acessível, para comunicar o problema e obter tratamento adequado. |

### 5.4.4 Operação

| ID | Job to Be Done |
| --- | --- |
| JTBD-O01 | Quando recebo uma denúncia, quero acessar informações autorizadas e relevantes sobre o caso, para tomar uma decisão fundamentada. |
| JTBD-O02 | Quando aplico uma medida de moderação, quero registrar seus fundamentos e consequências, para garantir consistência e rastreabilidade. |

## 5.5 Validação das personas

As personas e os Jobs to Be Done deverão ser validados por meio de:

- Entrevistas com clientes e profissionais de diferentes categorias de serviços.
- Investigação dos comportamentos atuais de descoberta, divulgação e contratação.
- Identificação das motivações e dificuldades recorrentes.
- Avaliação dos níveis de familiaridade digital e das condições de acesso à tecnologia.
- Identificação das necessidades relacionadas à confiança, segurança e privacidade.
- Testes de utilização dos principais fluxos da plataforma.

As evidências coletadas deverão orientar a revisão, consolidação, expansão ou descarte das personas inicialmente propostas.

---

# 6. Jornadas principais

As jornadas descrevem os principais caminhos percorridos pelos usuários para alcançar seus objetivos dentro do RAIRE.

Representam fluxos conceituais de interação e deverão orientar a definição dos requisitos funcionais, da experiência do usuário e das regras de negócio, sem estabelecer prematuramente detalhes de interface ou implementação tecnológica.

## 6.1 Descoberta inicial e criação de conta

**Objetivo:** permitir que novos usuários conheçam a plataforma, compreendam sua proposta de valor e iniciem sua utilização com baixa fricção.

### Fluxo principal

1. O visitante acessa a plataforma.
2. Identifica as possibilidades de encontrar profissionais, publicar necessidades ou divulgar serviços.
3. Pode explorar categorias, realizar buscas e consultar informações públicas de profissionais sem autenticação.
4. Ao iniciar uma ação que exige identificação, é direcionado à criação de conta ou autenticação.
5. Fornece as informações mínimas necessárias para o cadastro.
6. Realiza as confirmações de identidade ou contato exigidas pelo processo de cadastro.
7. Aceita os termos aplicáveis e recebe informações sobre privacidade e tratamento de dados.
8. Retoma, quando possível, a ação que motivou sua autenticação.

### Resultado esperado

O usuário compreende a proposta da plataforma e consegue iniciar sua participação sem etapas desnecessárias.

### Fricções a evitar

- Cadastro obrigatório antes da exploração inicial.
- Solicitação prematura de informações pessoais.
- Escolha permanente entre atuação como cliente e profissional.
- Perda do contexto de navegação após autenticação.

## 6.2 Ativação do perfil profissional

**Objetivo:** permitir que um usuário inicie sua atuação profissional no RAIRE e disponibilize informações suficientes para ser encontrado por potenciais clientes.

### Fluxo principal

1. O usuário seleciona a opção de iniciar sua atuação profissional.
2. Cria seu perfil profissional, informando nome de apresentação e descrição.
3. Seleciona as categorias e os serviços oferecidos.
4. Informa sua região e área de atendimento.
5. Declara sua disponibilidade e condições relevantes de atendimento.
6. Configura suas preferências de contato e compartilhamento de informações.
7. Visualiza uma prévia das informações que serão disponibilizadas publicamente.
8. Publica seu perfil profissional.
9. Recebe orientações sobre divulgação de serviços, construção de portfólio e critérios relevantes de descoberta.

### Resultado esperado

O profissional possui um perfil publicado, com informações suficientes para participar dos mecanismos de descoberta da plataforma.

### Fricções a evitar

- Formulários extensos.
- Exigência de informações que não sejam necessárias à publicação inicial.
- Obrigatoriedade de foto ou portfólio para categorias em que esses elementos não sejam essenciais.
- Exigência prematura de documentos.
- Utilização de linguagem excessivamente técnica.
- Exposição involuntária de informações pessoais.

## 6.3 Gestão do perfil e portfólio profissional

**Objetivo:** permitir que profissionais apresentem suas competências, mantenham suas informações atualizadas e construam uma presença digital especializada.

### Fluxo principal

1. O profissional acessa a área de gerenciamento de seu perfil.
2. Consulta as informações atualmente publicadas.
3. Atualiza sua descrição, serviços, especialidades, área de atendimento e disponibilidade, conforme necessário.
4. Adiciona trabalhos realizados ao seu portfólio.
5. Fornece descrições, imagens e informações contextuais relevantes para cada trabalho.
6. Visualiza como suas informações serão apresentadas aos clientes.
7. Publica as alterações realizadas.
8. Pode editar ou remover informações e trabalhos anteriormente publicados.
9. Pode suspender temporariamente a exposição de seu perfil profissional.

### Resultado esperado

O profissional consegue representar adequadamente seus serviços e experiências, mantendo controle sobre sua apresentação pública e disponibilidade.

### Fricções e riscos a evitar

- Dificuldade de atualizar informações frequentemente.
- Exigência de formatos complexos para publicação de trabalhos.
- Exposição de dados pessoais de clientes anteriores.
- Publicação de conteúdo sem autorização das pessoas envolvidas.
- Informações desatualizadas sobre disponibilidade e atendimento.

## 6.4 Descoberta de profissionais pelo cliente

**Objetivo:** permitir que clientes identifiquem profissionais adequados às suas necessidades e comparem alternativas relevantes em sua região.

### Fluxo principal

1. O cliente informa o serviço desejado ou seleciona uma categoria.
2. Informa a região de atendimento desejada, utilizando localização aproximada ou geolocalização opcional.
3. A plataforma apresenta profissionais compatíveis com o contexto da busca.
4. O cliente pode refinar os resultados utilizando critérios relevantes para o serviço.
5. Os resultados apresentam informações suficientes para uma avaliação inicial, incluindo serviços, área de atendimento, disponibilidade declarada e sinais de reputação.
6. O cliente consulta os perfis profissionais de seu interesse.
7. Analisa informações sobre competências, serviços, portfólios e experiências anteriores.
8. Compara alternativas disponíveis.
9. Identifica um ou mais profissionais com os quais deseja estabelecer contato.

### Resultado esperado

O cliente consegue identificar profissionais relevantes e possui informações suficientes para decidir se deseja avançar para uma conexão.

### Fluxos alternativos

- Quando não existem profissionais compatíveis, a plataforma informa claramente a ausência de resultados.
- O cliente pode revisar os critérios de busca ou consultar regiões de atendimento alternativas.
- O cliente pode publicar uma necessidade para receber manifestações de interesse.

### Fricções a evitar

- Resultados incompatíveis com a região ou categoria solicitada.
- Critérios de classificação incompreensíveis.
- Informações insuficientes para comparação.
- Exigência de localização precisa.
- Favorecimento comercial que comprometa a relevância dos resultados.

## 6.5 Conexão direta entre cliente e profissional

**Objetivo:** permitir que clientes estabeleçam contato com profissionais de seu interesse, preservando a autonomia e a privacidade das partes.

### Fluxo principal

1. O cliente identifica um profissional adequado à sua necessidade.
2. Manifesta interesse em estabelecer uma conexão.
3. Pode fornecer informações básicas sobre o serviço desejado.
4. O profissional recebe a solicitação e consulta as informações disponíveis.
5. O profissional decide aceitar ou recusar a solicitação.
6. Em caso de aceitação, a conexão é estabelecida.
7. Os meios de contato autorizados são disponibilizados conforme as preferências e o consentimento das partes.
8. Cliente e profissional podem iniciar a negociação do serviço.

### Resultado esperado

Uma conexão consentida é estabelecida entre cliente e profissional, permitindo que ambos avancem na negociação.

### Fluxos alternativos

- O profissional recusa a solicitação.
- O cliente cancela uma solicitação pendente.
- A solicitação expira sem resposta.
- Uma das partes denuncia ou bloqueia uma interação inadequada.

A conexão não representa compromisso de contratação, confirmação de pagamento ou garantia de execução do serviço.

## 6.6 Publicação e gestão de necessidades

**Objetivo:** permitir que clientes divulguem necessidades de serviços e recebam manifestações de interesse de profissionais compatíveis.

### Fluxo principal

1. O cliente inicia a publicação de uma necessidade.
2. Seleciona a categoria ou o serviço desejado.
3. Descreve as características relevantes da demanda.
4. Informa a região aproximada de atendimento.
5. Declara informações sobre urgência, prazo ou disponibilidade, quando aplicáveis.
6. Pode informar uma faixa de orçamento, quando desejar.
7. Visualiza as informações que serão disponibilizadas aos profissionais.
8. Publica a necessidade.
9. Recebe manifestações de interesse de profissionais.
10. Consulta os perfis, portfólios, reputações e informações fornecidas pelos interessados.
11. Seleciona os profissionais com os quais deseja estabelecer conexão.
12. Gerencia a necessidade até seu encerramento, cancelamento ou expiração.

### Resultado esperado

O cliente consegue divulgar sua necessidade e acessar alternativas de atendimento sem precisar procurar individualmente por todos os profissionais.

### Fluxos alternativos

- A necessidade não recebe manifestações de interesse.
- O cliente altera informações da demanda.
- O cliente encerra a necessidade após encontrar uma solução.
- A necessidade é cancelada ou expira.

### Fricções e riscos a evitar

- Exposição pública de endereço residencial ou informações de contato.
- Recebimento excessivo de manifestações irrelevantes.
- Publicação de demandas com informações insuficientes.
- Manutenção de demandas desatualizadas.
- Utilização da plataforma para spam ou divulgação indevida.

## 6.7 Descoberta de oportunidades pelo profissional

**Objetivo:** permitir que profissionais assumam uma participação ativa na aquisição de clientes e na identificação de oportunidades de trabalho.

### Fluxo principal

1. O profissional acessa a área de oportunidades.
2. Consulta necessidades disponíveis na região ou área de atendimento relevante.
3. Pode refinar as oportunidades por critérios compatíveis com sua atividade.
4. Identifica uma demanda de interesse.
5. Consulta sua descrição, localização aproximada, condições e informações relevantes sobre o cliente.
6. Avalia a compatibilidade da demanda com suas competências e disponibilidade.
7. Manifesta interesse na oportunidade.
8. Pode fornecer uma mensagem contextualizando sua capacidade de atendimento.
9. O cliente analisa a manifestação de interesse.
10. Caso o cliente aceite a conexão, os meios de contato autorizados são disponibilizados.
11. As partes podem avançar na negociação do serviço.

### Resultado esperado

O profissional consegue identificar oportunidades relevantes e estabelecer conexões com potenciais clientes de maneira ativa.

### Fluxos alternativos

- O profissional retira uma manifestação de interesse pendente.
- O cliente rejeita ou não seleciona a manifestação.
- A necessidade é encerrada antes do estabelecimento da conexão.
- Não existem oportunidades compatíveis com os critérios selecionados.

### Fricções e riscos a evitar

- Demandas incompatíveis com a atuação profissional.
- Informações insuficientes para avaliar oportunidades.
- Exposição prematura de dados pessoais.
- Competição baseada exclusivamente em popularidade ou capacidade de investimento.
- Manifestações automatizadas, excessivas ou abusivas.

## 6.8 Negociação e prestação do serviço

**Objetivo:** estabelecer os limites da participação do RAIRE após a conexão entre cliente e profissional.

### Fluxo principal

1. Cliente e profissional estabelecem contato mediante conexão consentida.
2. As partes discutem as condições do serviço.
3. Negociam valores, prazos, responsabilidades e demais condições aplicáveis.
4. Decidem autonomamente se desejam prosseguir com a contratação.
5. Quando houver contratação, a execução ocorre sob responsabilidade das partes.
6. Os usuários podem registrar informações sobre o desfecho da interação no RAIRE.

### Resultado esperado

As partes possuem autonomia para conduzir a negociação e eventual prestação do serviço.

No MVP, o RAIRE não intermediará pagamentos, formalização contratual ou gerenciamento da execução do serviço.

O registro de um desfecho informado pelos usuários não constitui comprovação independente de contratação ou conclusão do serviço.

## 6.9 Avaliação e construção de reputação

**Objetivo:** permitir que clientes e profissionais registrem experiências relevantes, contribuindo para a construção de confiança bilateral.

### Fluxo principal

1. Após uma interação elegível, o usuário recebe a possibilidade de registrar seu desfecho.
2. Informa se houve negociação, contratação ou prestação do serviço, conforme as opções disponíveis.
3. A plataforma determina a elegibilidade para avaliação segundo as regras de negócio aplicáveis.
4. O cliente pode avaliar aspectos relevantes da experiência com o profissional.
5. O profissional pode avaliar aspectos relevantes da experiência com o cliente.
6. As avaliações são registradas e apresentadas conforme as políticas de publicação e privacidade.
7. As informações contribuem para a construção da reputação contextual dos usuários.
8. Os usuários podem denunciar avaliações que considerem abusivas, fraudulentas ou incompatíveis com as políticas.

### Resultado esperado

Clientes e profissionais conseguem construir reputação a partir de experiências contextualizadas, respeitando os mecanismos de elegibilidade e integridade estabelecidos pela plataforma.

### Fricções e riscos a evitar

- Avaliações sem interação elegível.
- Avaliações retaliatórias.
- Manipulação artificial de reputação.
- Exposição indevida de informações pessoais.
- Utilização de avaliações genéricas como garantia de qualidade.
- Mistura indiscriminada entre reputação como cliente e como profissional.
- Apresentação de experiências autodeclaradas como serviços independentemente verificados.

## 6.10 Denúncia, moderação e segurança

**Objetivo:** oferecer mecanismos de proteção aos usuários e preservar a integridade das interações na plataforma.

### Fluxo principal

1. O usuário identifica conteúdo, comportamento ou interação potencialmente irregular.
2. Acessa o mecanismo de denúncia correspondente.
3. Seleciona o motivo e fornece informações adicionais quando necessário.
4. A denúncia é registrada para análise.
5. O caso recebe prioridade de acordo com sua natureza e risco potencial.
6. Um operador autorizado consulta as informações necessárias e o histórico disponível.
7. O operador avalia o caso segundo as políticas da plataforma.
8. Uma medida é aplicada quando cabível, podendo incluir advertência, remoção de conteúdo, restrição ou suspensão.
9. A decisão e seus fundamentos são registrados para fins de auditoria.
10. As partes são informadas quando aplicável, respeitando requisitos de segurança e privacidade.
11. Mecanismos de contestação são disponibilizados conforme as políticas estabelecidas.

### Resultado esperado

Os usuários conseguem comunicar irregularidades, e a plataforma dispõe de mecanismos proporcionais, consistentes e rastreáveis para tratamento dos casos.

### Fricções e riscos a evitar

- Dificuldade de localizar mecanismos de denúncia.
- Exposição da identidade do denunciante.
- Decisões arbitrárias ou não auditáveis.
- Acesso administrativo excessivo a dados pessoais.
- Utilização abusiva dos mecanismos de denúncia.
- Ausência de tratamento proporcional ao risco.

## 6.11 Considerações transversais

Todas as jornadas deverão observar os princípios de:

- Simplicidade e redução de fricção.
- Acessibilidade e inclusão digital.
- Privacidade e segurança por padrão.
- Autonomia dos usuários.
- Transparência nas interações.
- Consentimento no compartilhamento de informações.
- Tratamento adequado de falhas e estados excepcionais.

Os detalhes referentes a interfaces, estados de domínio, transições, regras de negócio, permissões e critérios de aceitação serão definidos nas especificações correspondentes.

---

# 7. Escopo do MVP

O MVP do RAIRE será constituído pelo conjunto mínimo de capacidades necessárias para validar sua proposta de valor bilateral no mercado local de serviços.

Seu escopo deverá permitir que clientes encontrem profissionais adequados às suas necessidades e que profissionais divulguem seus serviços, apresentem suas competências e acessem oportunidades de trabalho, estabelecendo conexões relevantes em um ambiente orientado à confiança, à segurança e à privacidade.

A implementação deverá priorizar a validação das hipóteses fundamentais do produto, evitando complexidade tecnológica e operacional que não contribua diretamente para esse objetivo.

## 7.1 Escopo funcional incluído

| Área | Capacidades incluídas no MVP |
| --- | --- |
| Identidade e contas | Cadastro, autenticação, confirmação de contato, recuperação de acesso e possibilidade de atuação como cliente e profissional em uma mesma conta. |
| Perfil do usuário | Gerenciamento de informações pessoais, preferências de contato, privacidade e informações relevantes à participação na plataforma. |
| Perfil profissional | Criação, publicação, atualização e suspensão do perfil; apresentação profissional; declaração de serviços, especialidades, área de atendimento e disponibilidade. |
| Portfólio profissional | Publicação, edição e remoção de trabalhos realizados, incluindo imagens, descrições e informações contextuais relevantes. |
| Catálogo de serviços | Organização inicial de categorias e subcategorias de serviços, com classificação suficiente para apoiar a descoberta de profissionais e oportunidades. |
| Localização | Definição de região aproximada, área de atendimento profissional e localização de demandas, preservando informações geográficas sensíveis. |
| Descoberta de profissionais | Busca por serviço, categoria e localização; aplicação de filtros essenciais; apresentação de resultados relevantes e acesso a perfis profissionais públicos. |
| Descoberta de oportunidades | Consulta e filtragem de demandas disponíveis, permitindo que profissionais identifiquem oportunidades compatíveis com sua atuação. |
| Publicação de necessidades | Criação, publicação, atualização, encerramento e cancelamento de demandas por serviços locais. |
| Manifestações de interesse | Registro, consulta, retirada e gerenciamento de manifestações de interesse de profissionais em demandas publicadas. |
| Conexões | Solicitação, aceitação, recusa, cancelamento e expiração de solicitações de conexão, independentemente do mecanismo de descoberta utilizado. |
| Comunicação inicial | Disponibilização de meios de contato autorizados após o estabelecimento de uma conexão consentida. |
| Desfecho das interações | Registro simplificado de informações declaradas pelos usuários sobre os resultados das conexões estabelecidas. |
| Reputação bilateral | Registro e apresentação de avaliações elegíveis de clientes e profissionais, respeitando o contexto das experiências e as políticas de integridade. |
| Segurança e confiança | Denúncia de conteúdo e comportamento inadequado, mecanismos básicos de bloqueio, confirmação de informações quando aplicável e tratamento de irregularidades. |
| Administração operacional | Gerenciamento do catálogo, análise de denúncias, moderação de conteúdo, aplicação de restrições e registro auditável de ações administrativas. |

## 7.2 Mecanismos de descoberta e distribuição de visibilidade

Os mecanismos iniciais de descoberta deverão considerar critérios relevantes ao contexto de cada busca, incluindo, quando aplicável:

- Compatibilidade entre o serviço procurado e os serviços oferecidos.
- Região e área de atendimento.
- Disponibilidade declarada.
- Informações profissionais relevantes.
- Sinais de reputação disponíveis.

A distribuição de visibilidade deverá observar os princípios de relevância, equidade e transparência estabelecidos pelo produto.

O MVP deverá contemplar:

- Critérios iniciais de classificação explícitos, documentados e testáveis.
- Mecanismos básicos para evitar concentração injustificada de exposição.
- Possibilidade de descoberta de profissionais sem histórico consolidado na plataforma.
- Apresentação compreensível dos principais fatores de relevância dos resultados.
- Coleta de informações suficientes para avaliar a distribuição de exposição.

A implementação de mecanismos avançados de otimização, personalização, experimentação e redistribuição de visibilidade dependerá das evidências obtidas durante a utilização do produto.

## 7.3 Escopo transversal de qualidade

| Área | Compromissos do MVP |
| --- | --- |
| Usabilidade | Experiência simples, linguagem clara e redução de etapas desnecessárias nos fluxos principais. |
| Acessibilidade | Atendimento aos critérios aplicáveis das diretrizes WCAG 2.2 AA, incluindo navegação por teclado, tecnologias assistivas e alternativas acessíveis de interação. |
| Responsividade | Experiência web mobile-first, adequada a dispositivos com diferentes dimensões de tela. |
| Desempenho | Carregamento e funcionamento adequados em dispositivos modestos e condições limitadas de conectividade. |
| Privacidade | Minimização de dados, localização aproximada, compartilhamento autorizado de informações pessoais e atendimento aos direitos dos titulares previstos na LGPD. |
| Segurança | Proteção de contas, dados pessoais e operações, incluindo mecanismos proporcionais de prevenção de abuso, controle de acesso e auditoria. |
| Transparência | Comunicação compreensível sobre critérios de descoberta, reputação, verificações realizadas e limitações da plataforma. |
| Observabilidade | Capacidade de identificar falhas técnicas, comportamentos anômalos e problemas nos fluxos essenciais. |

Os critérios técnicos e operacionais de aceitação serão detalhados nas seções de requisitos não funcionais e segurança.

## 7.4 Escopo de pesquisa e validação

O MVP deverá permitir a obtenção de evidências suficientes para avaliar suas hipóteses fundamentais, sem exigir inicialmente uma infraestrutura especializada de experimentação.

O escopo de pesquisa contempla:

- Instrumentação dos principais fluxos de utilização.
- Medição da disponibilidade de oferta e demanda por região e categoria.
- Medição da descoberta, manifestação de interesse e estabelecimento de conexões.
- Registro de informações sobre a distribuição de exposição dos profissionais.
- Pesquisa qualitativa com clientes e profissionais.
- Avaliação da satisfação, confiança percebida e utilidade das conexões.
- Identificação de dificuldades de utilização e barreiras de adoção.
- Registro e acompanhamento das hipóteses de produto.

A coleta e a análise de informações deverão respeitar os princípios de finalidade, necessidade, transparência e proteção de dados pessoais.

Ferramentas administrativas ou analíticas especializadas somente deverão ser desenvolvidas quando sua necessidade operacional estiver demonstrada.

## 7.5 Fora do escopo do MVP

### 7.5.1 Intermediação comercial e financeira

- Processamento e intermediação de pagamentos entre clientes e profissionais.
- Cobrança de comissões sobre serviços realizados.
- Carteiras digitais, custódia de valores ou sistemas de recompensas financeiras.
- Gestão de disputas financeiras.
- Formalização e gerenciamento de contratos de prestação de serviços.
- Implementação de mecanismos comerciais que favoreçam artificialmente posições nos resultados de descoberta.

O modelo de monetização da plataforma permanece como objeto de investigação e não constitui requisito de implementação do MVP.

### 7.5.2 Gestão da prestação de serviços

- Sistemas completos de agendamento e gerenciamento de calendários.
- Confirmação automatizada de compromissos.
- Gerenciamento de projetos, atividades ou execução de serviços.
- Chat em tempo real, chamadas de voz e videoconferência.
- Gestão financeira ou administrativa da atividade profissional.

### 7.5.3 Recursos sociais

- Feed generalista de conteúdo.
- Sistema de seguidores.
- Curtidas e mecanismos de popularidade social.
- Funcionalidades orientadas à viralização de conteúdo.
- Sistemas de gamificação destinados exclusivamente ao aumento de engajamento.

O portfólio profissional não é considerado uma funcionalidade de rede social, mas uma capacidade especializada de apresentação de competências e trabalhos realizados.

### 7.5.4 Garantias e verificações avançadas

- Verificação de antecedentes criminais.
- Garantia da qualidade ou conclusão de serviços.
- Seguros relacionados à prestação de serviços.
- Certificação profissional realizada diretamente pela plataforma.
- Verificação independente de todas as experiências profissionais declaradas.
- Sistemas avançados de investigação e prevenção automatizada de fraudes.

### 7.5.5 Complexidade tecnológica

- Aplicativos móveis nativos para iOS e Android.
- Matching automatizado por inteligência artificial generativa.
- Sistemas avançados de recomendação personalizada.
- Infraestrutura própria de experimentação em larga escala.
- Otimização automatizada de ranking baseada em aprendizado de máquina.
- Integrações com ERPs, CRMs ou marketplaces externos.
- Automação integral de decisões de moderação.

### 7.5.6 Expansão de mercado

- Operação simultânea em múltiplas regiões piloto independentes.
- Atendimento irrestrito a todas as categorias de serviços.
- Internacionalização.
- Suporte a múltiplas moedas.
- Expansão para um mercado genérico de contratação de freelancers sem delimitação geográfica.

Os itens fora do escopo poderão ser reavaliados a partir dos resultados do MVP e das necessidades identificadas durante sua utilização.

## 7.6 Estratégia de piloto

### 7.6.1 Delimitação geográfica

O piloto deverá ocorrer em uma cidade ou região geográfica delimitada, selecionada com base na capacidade de mobilização da equipe, acesso aos públicos envolvidos e potencial de concentração de oferta e demanda.

A seleção deverá considerar:

- Acesso a profissionais autônomos para recrutamento inicial.
- Existência de demanda identificável por serviços locais.
- Possibilidade de acompanhamento e suporte aos participantes.
- Capacidade de estabelecer parcerias com comunidades e organizações locais.
- Viabilidade operacional da pesquisa e da validação do produto.

### 7.6.2 Seleção de categorias

O piloto deverá contemplar um conjunto reduzido de categorias de serviços, inicialmente sugerido entre duas e três.

A seleção deverá considerar:

- Frequência e intensidade da demanda.
- Disponibilidade de profissionais na região.
- Dificuldades existentes de descoberta e divulgação.
- Relevância da localização para a contratação.
- Utilidade de portfólios e informações profissionais.
- Viabilidade de estabelecimento de conexões.
- Riscos relacionados à segurança, regulamentação e responsabilidade.

A inclusão de serviços remotos ou híbridos dependerá de sua aderência às hipóteses de validação do mercado local.

### 7.6.3 Formação da oferta inicial

A aquisição inicial de profissionais poderá ocorrer por meio de:

- Recrutamento direto.
- Parcerias com associações, comunidades e organizações locais.
- Convites a profissionais das categorias selecionadas.
- Apoio inicial na criação de perfis e portfólios.
- Orientação sobre a utilização da plataforma.

A formação da oferta deverá priorizar profissionais reais, disponíveis e compatíveis com as necessidades previstas para o piloto.

### 7.6.4 Formação da demanda inicial

A aquisição inicial de clientes poderá ocorrer por meio de:

- Divulgação em comunidades locais.
- Parcerias com organizações e grupos regionais.
- Campanhas de divulgação de baixo custo.
- Convites a participantes de pesquisas.
- Divulgação direcionada às categorias selecionadas.

As estratégias de aquisição deverão buscar necessidades reais de serviços, evitando a geração artificial de interações como evidência de validação.

### 7.6.5 Execução do piloto

O piloto será realizado durante um período previamente delimitado, inicialmente estimado entre oito e doze semanas, sujeito à revisão conforme as condições operacionais e os objetivos de validação.

Durante sua execução, deverão ocorrer ciclos periódicos de acompanhamento, coleta de evidências e revisão das hipóteses.

O início do período de validação deverá considerar a existência de uma oferta inicial minimamente representativa nas categorias selecionadas.

### 7.6.6 Critérios de continuidade

Ao término do piloto, os resultados deverão fundamentar a decisão sobre continuidade, revisão, redução ou expansão do produto.

A avaliação deverá considerar:

- Relevância dos problemas identificados.
- Utilidade percebida por clientes e profissionais.
- Capacidade de estabelecer conexões relevantes.
- Densidade local de oferta e demanda.
- Eficiência dos mecanismos de descoberta.
- Distribuição de visibilidade.
- Confiança e segurança das interações.
- Barreiras de adoção e utilização.
- Viabilidade operacional da continuidade.

As metas quantitativas e os critérios de decisão deverão ser estabelecidos na seção de métricas e validação, considerando as hipóteses prioritárias do MVP.

---

# 8. Requisitos funcionais

Os requisitos funcionais descrevem os comportamentos e as capacidades que o RAIRE deverá disponibilizar aos seus usuários e operadores.

Os requisitos deverão manter rastreabilidade com os objetivos do produto, as jornadas dos usuários, os princípios estabelecidos e as regras de negócio aplicáveis.

As decisões referentes à arquitetura, representação de dados, tecnologias e mecanismos internos de implementação deverão ser detalhadas nas especificações correspondentes.

## 8.1 Classificação de prioridade

A priorização dos requisitos utiliza a metodologia MoSCoW:

- **M — Must have:** requisito indispensável para a operação ou validação das hipóteses fundamentais do MVP.
- **S — Should have:** requisito importante, cuja ausência inicial não inviabiliza necessariamente a validação do produto.
- **C — Could have:** requisito complementar, condicionado à capacidade de implementação e às evidências disponíveis.
- **W — Won't have:** requisito explicitamente excluído do ciclo atual.

Requisitos obrigatórios de segurança, privacidade, acessibilidade e conformidade não poderão ser dispensados exclusivamente em função da priorização funcional.

## RF-01 — Identidade, conta e acesso

**Objetivo:** permitir que usuários criem, acessem e gerenciem suas contas, preservando a identidade e a segurança de sua participação na plataforma.

**Rastreabilidade:** OBJ-01; jornadas 6.1 e 6.2.

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

Os mecanismos de autenticação, proteção de credenciais, gerenciamento de sessões e autorização deverão observar os requisitos de segurança.

## RF-02 — Gestão do perfil do usuário

**Objetivo:** permitir que usuários gerenciem suas informações pessoais e preferências de participação na plataforma.

**Rastreabilidade:** OBJ-01, OBJ-06 e OBJ-07; jornadas 6.1 e 6.5.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-02.01 | Permitir a consulta e atualização das informações pessoais editáveis. | M |
| RF-02.02 | Permitir a configuração dos meios de contato autorizados. | M |
| RF-02.03 | Permitir o gerenciamento das preferências de compartilhamento de informações. | M |
| RF-02.04 | Permitir o gerenciamento das preferências de notificações não essenciais. | M |
| RF-02.05 | Permitir que usuários consultem suas conexões e interações registradas. | M |
| RF-02.06 | Permitir que usuários consultem suas informações de reputação, quando disponíveis. | S |
| RF-02.07 | Permitir a solicitação de desativação ou exclusão da conta, observadas as obrigações legais aplicáveis. | M |

## RF-03 — Gestão do perfil profissional

**Objetivo:** permitir que usuários apresentem suas competências e disponibilizem informações profissionais relevantes para descoberta e conexão.

**Rastreabilidade:** OBJ-01, OBJ-03 e OBJ-08; jornadas 6.2 e 6.3.

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

Perfis não publicados, suspensos ou inelegíveis não deverão participar dos mecanismos públicos de descoberta.

## RF-04 — Portfólio profissional

**Objetivo:** permitir que profissionais demonstrem suas competências por meio da apresentação organizada de trabalhos e experiências anteriores.

**Rastreabilidade:** OBJ-02, OBJ-03 e OBJ-07; jornadas 6.3 e 6.4.

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

Trabalhos apresentados no portfólio não deverão ser considerados automaticamente experiências verificadas pela plataforma.

## RF-05 — Catálogo de serviços

**Objetivo:** disponibilizar uma estrutura de classificação que permita organizar adequadamente serviços, profissionais e demandas.

**Rastreabilidade:** OBJ-01, OBJ-03, OBJ-04 e OBJ-05; jornadas 6.2, 6.4, 6.6 e 6.7.

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

A desativação de classificações deverá preservar a integridade das informações históricas relacionadas.

## RF-06 — Descoberta de profissionais

**Objetivo:** permitir que clientes identifiquem profissionais adequados às suas necessidades e consultem informações relevantes para tomada de decisão.

**Rastreabilidade:** OBJ-03, OBJ-06 e OBJ-08; jornada 6.4.

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

A descoberta não deverá expor endereços residenciais ou informações privadas de contato.

## RF-07 — Relevância, equidade e transparência da descoberta

**Objetivo:** estabelecer mecanismos de descoberta que preservem relevância, oportunidades de exposição e compreensão dos resultados.

**Rastreabilidade:** OBJ-03 e OBJ-08; jornadas 6.4 e 6.7.

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

A ausência de avaliações ou de histórico de participação não deverá representar, isoladamente, uma condição de exclusão da descoberta.

A classificação não deverá utilizar favorecimento pago nem atributos pessoais sensíveis como critérios de ordenação.

Os sinais, pesos, mecanismos de elegibilidade, estratégias de diversificação e métodos de avaliação deverão ser definidos em uma especificação própria de descoberta e ranking.

## RF-08 — Publicação e gestão de necessidades

**Objetivo:** permitir que clientes comuniquem demandas por serviços e recebam manifestações de interesse de profissionais.

**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornada 6.6.

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

Os critérios de publicação, edição, encerramento, expiração e visibilidade deverão ser definidos nas regras de negócio correspondentes.

## RF-09 — Descoberta de oportunidades

**Objetivo:** permitir que profissionais identifiquem ativamente oportunidades de trabalho compatíveis com sua atuação.

**Rastreabilidade:** OBJ-05 e OBJ-06; jornada 6.7.

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

## RF-10 — Manifestações de interesse

**Objetivo:** permitir que profissionais expressem interesse em necessidades publicadas por clientes.

**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornadas 6.6 e 6.7.

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

Os limites quantitativos de manifestações e os mecanismos de prevenção de abuso deverão ser definidos conforme as regras de negócio e as evidências do piloto.

## RF-11 — Solicitações de conexão direta

**Objetivo:** permitir que clientes iniciem conexões diretamente a partir da descoberta de profissionais, independentemente da publicação de uma necessidade.

**Rastreabilidade:** OBJ-03 e OBJ-06; jornada 6.5.

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

## RF-12 — Conexões consentidas

**Objetivo:** permitir que clientes e profissionais estabeleçam contato mediante manifestação válida de interesse e consentimento das partes.

**Rastreabilidade:** OBJ-06 e OBJ-07; jornadas 6.5, 6.6 e 6.7.

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

Uma conexão estabelecida não constitui confirmação de contratação, pagamento ou realização de serviço.

A revogação de acesso não garante a eliminação de informações que tenham sido anteriormente consultadas, copiadas ou armazenadas por terceiros.

## RF-13 — Registro de desfechos

**Objetivo:** permitir que os usuários informem os resultados das interações originadas na plataforma.

**Rastreabilidade:** OBJ-06 e OBJ-07; jornadas 6.8 e 6.9.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-13.01 | Permitir que participantes de uma conexão informem o desfecho da interação. | M |
| RF-13.02 | Permitir distinguir desfechos relacionados a contato, negociação e eventual realização do serviço. | M |
| RF-13.03 | Permitir que o usuário registe sua percepção sobre a utilidade da conexão. | S |
| RF-13.04 | Preservar a autoria das informações declaradas sobre o desfecho. | M |
| RF-13.05 | Permitir identificar divergências entre informações declaradas pelas partes. | S |
| RF-13.06 | Utilizar informações de desfecho para determinar elegibilidade de avaliações conforme as regras aplicáveis. | M |

Informações declaradas pelos usuários não deverão ser apresentadas como comprovações independentes de contratação ou execução do serviço.

## RF-14 — Avaliações e reputação bilateral

**Objetivo:** disponibilizar mecanismos de construção de reputação contextual para clientes e profissionais.

**Rastreabilidade:** OBJ-07; jornada 6.9.

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

A elegibilidade, o formato, a publicação, a edição e a agregação das avaliações deverão ser definidos em especificação própria.

Os mecanismos de reputação não deverão representar garantias de competência, comportamento ou qualidade de serviços.

## RF-15 — Confirmação e verificação de informações

**Objetivo:** disponibilizar sinais de confiança correspondentes exclusivamente a informações cuja confirmação tenha sido efetivamente realizada.

**Rastreabilidade:** OBJ-07; jornadas 6.1, 6.2 e 6.4.

| ID | Requisito | Prioridade |
| --- | --- | --- |
| RF-15.01 | Registrar o estado das confirmações de informações de contato realizadas. | M |
| RF-15.02 | Apresentar sinais de confirmação correspondentes às informações efetivamente confirmadas. | M |
| RF-15.03 | Explicar o significado e as limitações dos sinais de confirmação apresentados. | M |
| RF-15.04 | Permitir verificações adicionais de informações profissionais, quando justificadas operacionalmente. | C |
| RF-15.05 | Permitir a atualização ou invalidação de verificações que deixem de satisfazer os critérios aplicáveis. | C |

A confirmação de um endereço de e-mail ou meio de contato não deverá ser apresentada como verificação integral de identidade.

A verificação documental de identidade dependerá de decisão específica sobre necessidade, segurança, privacidade e viabilidade operacional.

## RF-16 — Notificações

**Objetivo:** informar os usuários sobre acontecimentos relevantes para suas interações com a plataforma.

**Rastreabilidade:** OBJ-04, OBJ-05 e OBJ-06; jornadas 6.5, 6.6, 6.7 e 6.9.

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

Os canais e mecanismos de entrega deverão ser definidos conforme as necessidades dos fluxos e os requisitos de confiabilidade e privacidade.

## RF-17 — Denúncias e proteção do usuário

**Objetivo:** permitir que usuários comuniquem irregularidades e controlem interações potencialmente indesejadas.

**Rastreabilidade:** OBJ-06 e OBJ-07; jornada 6.10.

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

A identidade e as informações do denunciante deverão ser protegidas contra exposição indevida.

## RF-18 — Administração e moderação

**Objetivo:** disponibilizar capacidades operacionais necessárias à manutenção da integridade, segurança e funcionamento da plataforma.

**Rastreabilidade:** OBJ-07 e OBJ-09; jornada 6.10.

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

As operações administrativas deverão respeitar as políticas de autorização, menor privilégio, proteção de dados e auditoria.

## RF-19 — Transparência e controle de dados pessoais

**Objetivo:** disponibilizar mecanismos funcionais para que usuários compreendam e exerçam controle sobre o tratamento de suas informações.

**Rastreabilidade:** princípios de privacidade, segurança, autonomia e transparência.

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

Os requisitos de conformidade com a LGPD, bases legais, períodos de retenção, segurança do tratamento e processos de resposta a incidentes deverão ser detalhados nas especificações de privacidade e segurança.

## RF-20 — Medição e validação do produto

**Objetivo:** disponibilizar informações suficientes para avaliar o funcionamento do produto e validar suas hipóteses fundamentais.

**Rastreabilidade:** OBJ-01 a OBJ-09; VAL-01 a VAL-07.

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

A instrumentação deverá respeitar os princípios de minimização, finalidade, segurança e privacidade.

Os esquemas dos eventos, indicadores, critérios de qualidade dos dados e métodos de análise deverão ser especificados no plano de medição e validação.

## 8.2 Restrições transversais

Todos os requisitos funcionais deverão observar:

- Os princípios de produto estabelecidos neste documento.
- As regras de negócio e condições de elegibilidade aplicáveis.
- Os requisitos de segurança e autorização.
- Os requisitos de acessibilidade e usabilidade.
- Os requisitos de proteção de dados pessoais.
- Os requisitos de desempenho e confiabilidade.
- As políticas de transparência, confiança e moderação.

## 8.3 Detalhamento complementar

Os requisitos funcionais deverão ser complementados, conforme sua complexidade, por especificações contendo:

- Atores envolvidos.
- Pré-condições e pós-condições.
- Fluxos principais e alternativos.
- Regras de negócio associadas.
- Estados e transições relevantes.
- Dependências funcionais.
- Critérios de aceitação.
- Cenários de falha.
- Considerações de segurança e privacidade.

A implementação de um requisito deverá ser rastreável às respectivas especificações, decisões e evidências de validação.

---

# 9. Requisitos não funcionais

Os requisitos não funcionais estabelecem os atributos de qualidade, restrições operacionais e condições necessárias para que o RAIRE execute suas capacidades funcionais de maneira adequada, segura, confiável e sustentável.

Esses requisitos deverão orientar a avaliação de alternativas arquiteturais, decisões tecnológicas, estratégias de implementação e critérios de aceitação do produto.

As metas quantitativas apresentadas nesta seção representam referências iniciais de qualidade para o MVP e deverão ser refinadas conforme o perfil de utilização, os recursos disponíveis e as evidências obtidas durante o piloto.

## RNF-01 — Manutenibilidade e modularidade

**Objetivo:** permitir a manutenção e evolução do sistema com complexidade controlada e impacto limitado entre capacidades de negócio.

O sistema deverá:

- Possuir responsabilidades claramente delimitadas.
- Preservar alta coesão e reduzir acoplamentos desnecessários.
- Permitir alterações localizadas em regras de negócio.
- Evitar dependências circulares entre componentes.
- Preservar a integridade das fronteiras estabelecidas para o domínio.
- Permitir a compreensão e manutenção da solução pela equipe responsável.
- Manter decisões arquiteturais relevantes documentadas e rastreáveis.

**Critérios de avaliação:**

- Revisões de arquitetura e código.
- Verificação das dependências entre componentes.
- Avaliação do impacto de alterações.
- Conformidade com as fronteiras arquiteturais estabelecidas.

A seleção do estilo arquitetural, linguagem de programação e estrutura de módulos deverá ser fundamentada em decisões arquiteturais específicas.

## RNF-02 — Desempenho

**Objetivo:** proporcionar tempos de resposta e carregamento compatíveis com uma experiência simples e eficiente.

O sistema deverá:

- Apresentar respostas adequadas nos fluxos críticos.
- Evitar processamento e transferência desnecessários de informações.
- Preservar a responsividade da experiência durante operações de maior duração.
- Disponibilizar mecanismos de paginação ou carregamento progressivo quando necessários.
- Otimizar a apresentação e transferência de recursos multimídia.

**Metas preliminares:**

| Indicador | Meta inicial |
| --- | --- |
| Latência de consultas simples à API | p95 ≤ 300 ms |
| Latência de busca de profissionais e oportunidades | p95 ≤ 500 ms |
| Largest Contentful Paint (LCP) | p75 ≤ 2,5 s |
| Interaction to Next Paint (INP) | p75 ≤ 200 ms |
| Cumulative Layout Shift (CLS) | p75 ≤ 0,1 |

As metas de latência deverão ser avaliadas sob condições documentadas de carga, infraestrutura e volume de dados.

As métricas de experiência web deverão considerar dispositivos móveis representativos e condições de conectividade compatíveis com o público-alvo.

Operações com características específicas, como processamento de imagens e tarefas assíncronas, deverão possuir critérios próprios.

## RNF-03 — Capacidade e escalabilidade

**Objetivo:** suportar o volume de utilização previsto para o piloto e permitir crescimento progressivo de acordo com a evolução da demanda.

O sistema deverá:

- Suportar o volume estimado de usuários, profissionais, demandas e conexões.
- Permitir a ampliação progressiva da capacidade de processamento e armazenamento.
- Evitar gargalos que impeçam o crescimento compatível com as necessidades do produto.
- Permitir a identificação dos componentes responsáveis por limitações de capacidade.
- Preservar os requisitos de desempenho dentro do perfil de carga estabelecido.

**Dimensões de capacidade a especificar:**

- Quantidade de usuários cadastrados e ativos.
- Quantidade de perfis profissionais publicados.
- Quantidade de demandas e conexões.
- Concorrência de usuários e operações.
- Volume de buscas e requisições.
- Volume de imagens e demais recursos armazenados.
- Taxa de crescimento dos dados.

As metas quantitativas deverão ser estabelecidas mediante um modelo de carga fundamentado nas expectativas do piloto.

Não constitui requisito inicial suportar volumes de utilização incompatíveis com as projeções do MVP.

## RNF-04 — Disponibilidade

**Objetivo:** manter as capacidades essenciais do RAIRE acessíveis durante sua operação.

O sistema deverá:

- Disponibilizar as jornadas essenciais durante o período operacional estabelecido.
- Permitir a identificação de indisponibilidades.
- Minimizar interrupções associadas a implantações e manutenções.
- Comunicar falhas de maneira compreensível aos usuários.
- Evitar que componentes não essenciais provoquem indisponibilidade generalizada.

**Meta preliminar:**

- Disponibilidade mensal de 99,5% para as capacidades críticas do produto.

O indicador de disponibilidade deverá possuir definição explícita quanto às jornadas monitoradas, janela de avaliação, condições de sucesso e tratamento de manutenções planejadas.

A meta deverá ser revisada conforme o orçamento e a criticidade operacional identificados durante o piloto.

## RNF-05 — Confiabilidade e integridade de dados

**Objetivo:** preservar a correção das informações e o cumprimento das regras de negócio, inclusive diante de falhas e operações concorrentes.

O sistema deverá:

- Preservar as invariantes do domínio.
- Impedir estados inconsistentes decorrentes de operações incompletas.
- Garantir a integridade das informações relacionadas a usuários, profissionais, demandas, interesses, conexões e avaliações.
- Impedir duplicações indevidas de operações sensíveis.
- Tratar adequadamente requisições repetidas e operações concorrentes.
- Preservar a consistência dos estados e transições de negócio.
- Permitir a identificação e tratamento de falhas de processamento.
- Impedir a confirmação de operações que não tenham sido efetivamente concluídas.

**Critérios de avaliação:**

- Testes de invariantes de negócio.
- Testes de concorrência em operações críticas.
- Testes de repetição e recuperação de operações.
- Testes de falhas parciais.
- Verificação da integridade dos dados.

Os mecanismos de controle transacional, isolamento, concorrência e idempotência deverão ser definidos no design da solução.

## RNF-06 — Resiliência e tolerância a falhas

**Objetivo:** limitar o impacto de falhas e preservar as capacidades essenciais quando componentes auxiliares estiverem indisponíveis.

O sistema deverá:

- Tratar falhas de serviços internos e dependências externas.
- Evitar propagação desnecessária de falhas.
- Permitir degradação controlada de funcionalidades não essenciais.
- Preservar dados e operações críticas diante de interrupções.
- Informar adequadamente condições de indisponibilidade.
- Permitir recuperação ou reprocessamento seguro de operações elegíveis.

**Cenários relevantes:**

- Indisponibilidade do mecanismo de notificações.
- Falha no processamento de imagens.
- Falha de componentes auxiliares de classificação.
- Interrupção temporária de serviços externos.
- Repetição de operações após falhas de comunicação.

Estratégias alternativas de classificação deverão preservar os critérios obrigatórios de elegibilidade, privacidade e equidade.

Falhas em mecanismos de autorização ou proteção de dados não deverão resultar na liberação indevida de informações ou operações.

## RNF-07 — Recuperabilidade e continuidade operacional

**Objetivo:** permitir a recuperação do sistema e de suas informações após incidentes técnicos ou perdas de infraestrutura.

O sistema deverá:

- Possuir mecanismos periódicos de proteção e recuperação de dados.
- Permitir a restauração das informações necessárias à continuidade operacional.
- Possuir procedimentos documentados de recuperação.
- Proteger cópias de segurança contra acessos indevidos e perdas evitáveis.
- Validar periodicamente a capacidade de restauração.

**Indicadores a estabelecer:**

| Indicador | Definição |
| --- | --- |
| RPO — Recovery Point Objective | Perda máxima de dados tolerável em caso de recuperação. |
| RTO — Recovery Time Objective | Tempo máximo tolerável para recuperação das capacidades afetadas. |
| Frequência de backup | Periodicidade de proteção dos dados. |
| Retenção de backup | Período de disponibilidade das cópias. |
| Validação de restauração | Frequência e resultado dos testes de recuperação. |

As metas deverão considerar a criticidade das informações, os riscos operacionais e os recursos disponíveis.

## RNF-08 — Segurança da informação

**Objetivo:** proteger usuários, informações e capacidades do sistema contra acesso indevido, comprometimento e abuso.

O sistema deverá:

- Adotar controles de segurança proporcionais aos riscos identificados.
- Proteger credenciais, sessões e mecanismos de autenticação.
- Possuir mecanismos de autorização baseados nas responsabilidades e no contexto de acesso.
- Impedir acesso indevido a recursos pertencentes a outros usuários.
- Aplicar o princípio do menor privilégio.
- Proteger interfaces contra entradas maliciosas e utilização abusiva.
- Proteger informações durante sua transmissão e armazenamento, conforme sua classificação.
- Gerenciar segredos e credenciais de maneira segura.
- Proteger operações administrativas com controles adicionais de autenticação, incluindo MFA.
- Permitir a identificação e tratamento de vulnerabilidades.
- Evitar a exposição de credenciais ou informações privadas em registros técnicos.
- Possuir procedimentos de resposta a incidentes.

**Referência inicial:**

Utilizar o OWASP ASVS Nível 2 como referência para definição e avaliação dos controles de segurança aplicáveis ao MVP.

A versão adotada, os controles aplicáveis e as evidências de atendimento deverão ser registrados.

A segurança deverá ser avaliada mediante análise de ameaças, revisão de código, testes automatizados e verificações específicas dos fluxos críticos.

## RNF-09 — Privacidade e proteção de dados pessoais

**Objetivo:** preservar os direitos dos titulares e assegurar o tratamento adequado de dados pessoais.

O sistema deverá:

- Adotar princípios de privacidade desde a concepção e por padrão.
- Coletar apenas informações necessárias às finalidades estabelecidas.
- Possuir finalidades e bases legais identificadas para as operações de tratamento.
- Preservar a localização precisa e as informações privadas de contato.
- Restringir o acesso a informações pessoais conforme o contexto autorizado.
- Permitir o exercício dos direitos dos titulares previstos na LGPD.
- Estabelecer critérios de retenção, eliminação e descarte de informações.
- Proteger informações utilizadas em atividades de pesquisa e análise.
- Evitar coleta desnecessária de atributos pessoais sensíveis.
- Possuir procedimentos de identificação, avaliação e tratamento de incidentes de privacidade.

Deverá ser avaliada a necessidade de elaboração de Relatório de Impacto à Proteção de Dados Pessoais (RIPD), considerando os riscos e as características das operações de tratamento.

Os requisitos legais, políticas internas e responsabilidades operacionais deverão ser detalhados em documentação específica de privacidade e governança de dados.

## RNF-10 — Acessibilidade

**Objetivo:** permitir que pessoas com diferentes capacidades e necessidades utilizem as jornadas oferecidas pelo RAIRE.

O sistema deverá:

- Atender aos critérios aplicáveis de conformidade WCAG 2.2 nível AA.
- Permitir navegação por teclado.
- Possuir suporte adequado a tecnologias assistivas.
- Disponibilizar contraste e legibilidade apropriados.
- Apresentar indicadores de foco perceptíveis.
- Disponibilizar alternativas acessíveis a conteúdos não textuais.
- Apresentar formulários, mensagens de erro e estados de interação compreensíveis.
- Respeitar preferências de redução de movimento.
- Evitar interações que dependam exclusivamente de cor, gestos complexos ou precisão motora.

**Critérios de avaliação:**

- Auditorias automatizadas de acessibilidade.
- Testes manuais de navegação por teclado.
- Testes com leitores de tela.
- Avaliação dos critérios de conformidade aplicáveis.
- Testes de utilização com pessoas representativas do público-alvo, quando viáveis.

Ferramentas automatizadas não deverão constituir a única evidência de conformidade.

## RNF-11 — Usabilidade e inclusão digital

**Objetivo:** reduzir barreiras de utilização para pessoas com diferentes níveis de familiaridade digital.

O sistema deverá:

- Utilizar linguagem clara e preferencialmente não técnica.
- Minimizar etapas e informações obrigatórias nas jornadas essenciais.
- Apresentar estados de carregamento, sucesso, ausência de resultados e falha de maneira compreensível.
- Permitir a identificação e correção de erros de preenchimento.
- Evitar a perda desnecessária de informações durante interações.
- Apresentar orientações adequadas ao contexto de utilização.
- Permitir a compreensão das consequências de ações relevantes.
- Evitar padrões de interface manipulativos ou enganosos.

**Indicadores de avaliação:**

- Taxa de conclusão das jornadas essenciais.
- Tempo necessário para conclusão das tarefas.
- Frequência de erros e abandonos.
- Necessidade de assistência externa.
- Satisfação e facilidade de utilização percebidas.

As metas quantitativas deverão ser estabelecidas mediante testes de usabilidade e evidências do piloto.

## RNF-12 — Compatibilidade e responsividade

**Objetivo:** proporcionar uma experiência funcional em dispositivos e condições de acesso representativos do público-alvo.

O sistema deverá:

- Adotar uma experiência web mobile-first.
- Preservar as capacidades essenciais em diferentes dimensões de tela.
- Suportar navegadores relevantes para o público-alvo.
- Evitar dependência desnecessária de recursos computacionais específicos.
- Considerar condições limitadas de conectividade.
- Preservar a legibilidade e operabilidade dos elementos de interface.

**Referências iniciais de avaliação:**

- Jornadas essenciais funcionais em telas com largura de 360 pixels CSS.
- Atendimento aos critérios aplicáveis de redimensionamento e refluxo de conteúdo da WCAG.
- Testes em dispositivos móveis representativos.
- Testes nos navegadores definidos na matriz de compatibilidade.

A matriz de dispositivos, navegadores e condições de rede deverá ser estabelecida conforme o perfil de utilização identificado.

## RNF-13 — Observabilidade

**Objetivo:** permitir a compreensão do comportamento técnico do sistema e a identificação de falhas operacionais.

O sistema deverá:

- Disponibilizar registros técnicos estruturados e contextualizados.
- Permitir a correlação de acontecimentos relacionados a uma mesma operação.
- Disponibilizar métricas relevantes de desempenho, disponibilidade e falhas.
- Permitir a identificação de erros recorrentes e comportamentos anômalos.
- Possuir mecanismos de monitoramento das capacidades críticas.
- Disponibilizar alertas proporcionais à criticidade operacional.
- Preservar a privacidade e segurança das informações de observabilidade.

**Indicadores mínimos:**

- Taxa de requisições e operações.
- Latência das capacidades críticas.
- Taxa de erros.
- Disponibilidade.
- Utilização de recursos.
- Falhas em processamentos assíncronos, quando aplicável.
- Falhas de dependências relevantes.

Os mecanismos de rastreamento distribuído deverão ser adotados quando sua necessidade for justificada pela arquitetura e complexidade operacional.

## RNF-14 — Auditabilidade e rastreabilidade

**Objetivo:** preservar evidências suficientes sobre operações relevantes para segurança, integridade, confiança e responsabilização.

O sistema deverá:

- Registrar operações administrativas relevantes.
- Preservar a identificação do responsável, contexto e momento da operação.
- Registrar alterações relevantes em políticas e configurações.
- Preservar evidências necessárias à investigação de incidentes.
- Proteger registros de auditoria contra alterações ou exclusões indevidas.
- Restringir o acesso aos registros conforme responsabilidades autorizadas.
- Estabelecer critérios de retenção e consulta.
- Permitir a rastreabilidade entre decisões, requisitos e evidências de validação.

Os registros de auditoria deverão ser diferenciados dos registros técnicos de diagnóstico e dos eventos destinados à análise de produto.

## RNF-15 — Testabilidade

**Objetivo:** permitir a verificação sistemática dos comportamentos, regras de negócio e atributos de qualidade do sistema.

O sistema deverá:

- Possuir comportamentos verificáveis mediante testes.
- Permitir a validação isolada das regras de negócio relevantes.
- Permitir a verificação da integração entre capacidades e dependências.
- Disponibilizar ambientes e dados de teste adequados.
- Permitir a reprodução de cenários de falha.
- Possuir testes de regressão para comportamentos críticos.
- Permitir a validação das políticas de classificação, elegibilidade e distribuição de visibilidade.
- Permitir a execução automatizada de verificações relevantes.

**Áreas prioritárias de validação:**

- Autenticação e autorização.
- Privacidade das informações de contato e localização.
- Elegibilidade de profissionais e demandas.
- Manifestações de interesse e conexões.
- Integridade e concorrência em operações críticas.
- Avaliações e reputação.
- Moderação.
- Critérios de descoberta e classificação.

A estratégia de testes e os critérios de cobertura deverão ser estabelecidos conforme os riscos e a criticidade dos comportamentos.

Percentuais de cobertura de código não deverão ser utilizados isoladamente como evidência de qualidade.

## RNF-16 — Equidade e explicabilidade

**Objetivo:** preservar critérios de descoberta relevantes, compreensíveis e passíveis de avaliação quanto à distribuição de oportunidades.

O sistema deverá:

- Possuir critérios de classificação documentados e versionados.
- Permitir a identificação dos fatores relevantes utilizados na apresentação dos resultados.
- Permitir a avaliação da distribuição de exposição entre profissionais elegíveis e comparáveis.
- Preservar a participação de profissionais sem histórico consolidado.
- Evitar discriminação e favorecimento comercial indevido.
- Permitir a avaliação dos efeitos de mudanças nas políticas de descoberta.
- Preservar informações suficientes para investigação de comportamentos inesperados.

Os mecanismos deverão ser avaliados considerando conjuntamente relevância, distribuição de exposição e utilidade para os usuários.

A utilização de atributos pessoais sensíveis como critérios de classificação deverá ser proibida.

Possíveis efeitos discriminatórios indiretos deverão ser investigados e tratados conforme as evidências e os riscos identificados.

Mecanismos avançados de simulação, replay de ranking e experimentação dependerão de justificativa específica.

## RNF-17 — Eficiência de custos

**Objetivo:** manter a operação e evolução do MVP compatíveis com os recursos financeiros disponíveis.

O sistema deverá:

- Evitar custos fixos incompatíveis com o estágio de validação do produto.
- Permitir a identificação dos principais componentes de custo.
- Evitar provisionamento excessivo de recursos.
- Permitir a ampliação de capacidade de maneira proporcional à demanda.
- Considerar custos de infraestrutura, armazenamento, transferência, serviços externos e manutenção.
- Avaliar serviços gerenciados considerando custo total, esforço operacional e riscos.

**Indicadores a estabelecer:**

- Orçamento operacional mensal.
- Custo de infraestrutura por período.
- Custo por usuário ativo.
- Custo por conexão estabelecida.
- Custo de armazenamento e distribuição de mídia.
- Custo de serviços externos.

A seleção tecnológica deverá considerar o custo total de propriedade, e não apenas o custo direto da infraestrutura.

## RNF-18 — Implantabilidade e reprodutibilidade

**Objetivo:** permitir que o sistema seja desenvolvido, testado e disponibilizado de maneira consistente e controlada.

O sistema deverá:

- Possuir procedimentos reproduzíveis de configuração dos ambientes.
- Permitir a instalação e execução local pelos integrantes da equipe.
- Manter dependências e configurações necessárias documentadas.
- Permitir a execução de verificações antes da disponibilização de alterações.
- Preservar a rastreabilidade das versões disponibilizadas.
- Possuir mecanismos controlados de evolução das estruturas persistentes.
- Permitir recuperação ou correção de implantações problemáticas.
- Manter isolamento adequado entre ambientes.

As estratégias de empacotamento, integração contínua, entrega e implantação deverão ser estabelecidas no design da solução.

## RNF-19 — Portabilidade e interoperabilidade

**Objetivo:** reduzir dependências tecnológicas desnecessárias e preservar alternativas razoáveis de evolução da infraestrutura.

O sistema deverá:

- Evitar acoplamento desnecessário a fornecedores específicos.
- Possuir dependências externas identificadas e documentadas.
- Utilizar formatos e interfaces interoperáveis quando adequados.
- Permitir a exportação e recuperação das informações persistidas.
- Isolar integrações externas conforme sua responsabilidade e criticidade.
- Considerar custos e riscos de eventual substituição de dependências.

A utilização de serviços proprietários ou gerenciados será aceitável quando os benefícios justificarem os custos e riscos associados.

Portabilidade não deverá implicar a criação de abstrações ou mecanismos de compatibilidade sem necessidade identificada.

## RNF-20 — Evolutividade

**Objetivo:** permitir que o produto incorpore novos conhecimentos e capacidades sem comprometer a estabilidade das funcionalidades existentes.

O sistema deverá:

- Permitir a evolução incremental de regras de negócio.
- Permitir a ampliação das categorias e regiões atendidas.
- Permitir a evolução dos mecanismos de descoberta e reputação.
- Preservar a compatibilidade das informações existentes durante alterações.
- Permitir a substituição ou evolução de componentes relevantes de maneira controlada.
- Manter decisões de evolução fundamentadas em evidências e necessidades identificadas.

A evolutividade não deverá justificar a implementação antecipada de capacidades fora do escopo do MVP.

## 9.1 Validação dos requisitos não funcionais

Os requisitos não funcionais deverão ser complementados por critérios de aceitação verificáveis e evidências proporcionais à sua criticidade.

Sempre que aplicável, a especificação deverá identificar:

- A capacidade ou componente avaliado.
- O cenário de utilização ou evento considerado.
- As condições operacionais.
- O comportamento esperado.
- A métrica correspondente.
- A meta ou limite de aceitação.
- O método de verificação.

As metas ainda não estabelecidas deverão ser registradas como decisões pendentes antes da disponibilização pública do MVP, conforme sua criticidade.

## 9.2 Restrições tecnológicas e arquiteturais

A linguagem de programação, o estilo arquitetural, os mecanismos de persistência, a infraestrutura, os serviços externos e as tecnologias de desenvolvimento não são definidos por esta seção.

Essas decisões deverão considerar os requisitos funcionais, os atributos de qualidade, as restrições da equipe, o orçamento e as necessidades operacionais do produto.

As decisões relevantes deverão ser registradas em Architecture Decision Records (ADRs), incluindo contexto, alternativas avaliadas, justificativa, consequências e trade-offs.

---

# 10. Estrutura de equidade, confiança e transparência

O RAIRE deverá promover um ambiente de descoberta e conexão no qual clientes e profissionais possam participar de maneira acessível, autônoma, segura e equitativa.

A estrutura de equidade, confiança e transparência estabelece os princípios, objetivos, critérios de avaliação e restrições que deverão orientar as políticas de descoberta, reputação, verificação, proteção dos usuários e moderação.

Os mecanismos específicos de implementação deverão ser definidos nas respectivas especificações de domínio e fundamentados em evidências, riscos e necessidades identificadas durante a evolução do produto.

## 10.1 Equidade (fairness)

### 10.1.1 Definição operacional

Equidade corresponde à redução de barreiras e favorecimentos injustificados no acesso, na descoberta e na distribuição de oportunidades entre participantes do mercado local de serviços.

O RAIRE deverá buscar condições adequadas de participação para profissionais com diferentes níveis de experiência, presença digital, reputação e histórico de utilização da plataforma.

A equidade não implica:

- Igualdade artificial de exposição.
- Garantia de contratação ou geração de renda.
- Distribuição uniforme de oportunidades entre profissionais com diferentes níveis de compatibilidade.
- Priorização indiscriminada de profissionais novos.
- Comprometimento da relevância ou da segurança das conexões.
- Interferência indevida na autonomia de escolha dos usuários.

### 10.1.2 Dimensões de equidade

| Dimensão | Objetivo |
| --- | --- |
| Acesso | Reduzir barreiras injustificadas à participação na plataforma. |
| Elegibilidade | Aplicar critérios consistentes e pertinentes ao contexto de descoberta. |
| Exposição | Reduzir concentrações injustificadas de visibilidade entre profissionais comparáveis. |
| Oportunidades | Facilitar o acesso a demandas compatíveis e possibilidades relevantes de conexão. |
| Procedimental | Assegurar consistência, transparência e possibilidade de contestação das decisões da plataforma. |

### 10.1.3 Princípios de descoberta equitativa

Os mecanismos de descoberta deverão:

- Priorizar a compatibilidade entre as necessidades dos usuários e os serviços oferecidos.
- Aplicar critérios de elegibilidade antes de mecanismos de distribuição de exposição.
- Permitir a participação de profissionais sem histórico consolidado.
- Evitar dependência excessiva de popularidade ou volume histórico de interações.
- Evitar mecanismos que incentivem artificialmente a frequência de utilização da plataforma.
- Considerar a relevância contextual das informações profissionais.
- Preservar a autonomia dos usuários na escolha de profissionais e oportunidades.
- Impedir favorecimento comercial por meio da aquisição de posições nos resultados durante o MVP.

A ausência de avaliações, portfólio extenso ou histórico de participação não deverá, isoladamente, impedir a descoberta de profissionais que satisfaçam os critérios mínimos de elegibilidade.

### 10.1.4 Dimensões de análise e coortes observáveis

A avaliação dos mecanismos de descoberta poderá considerar agrupamentos relacionados a características operacionais pertinentes ao domínio.

Exemplos:

- Tempo de participação na plataforma.
- Quantidade de avaliações recebidas.
- Histórico de exposição.
- Histórico de oportunidades e conexões.
- Disponibilidade declarada.
- Categoria ou serviço oferecido.
- Região e área de atendimento.
- Tipo de atendimento, quando aplicável.

A utilização dessas dimensões deverá possuir finalidade explícita, justificativa e avaliação de possíveis impactos adversos.

Atributos pessoais sensíveis, protegidos ou irrelevantes ao contexto do serviço, como raça, gênero, idade, deficiência, religião e orientação sexual, não deverão ser utilizados como critérios de favorecimento ou penalização nos mecanismos de classificação.

A plataforma não deverá inferir esses atributos para personalização ou ordenação de resultados.

Possíveis efeitos discriminatórios indiretos, inclusive aqueles associados à localização, atividade, reputação e características socioeconômicas, deverão ser considerados na avaliação das políticas de descoberta.

A ausência de atributos sensíveis nos mecanismos de classificação não deverá ser interpretada como comprovação de ausência de discriminação.

### 10.1.5 Indicadores de equidade

Os indicadores deverão permitir a avaliação conjunta de elegibilidade, exposição, oportunidades e relevância.

| Dimensão | Indicadores candidatos |
| --- | --- |
| Elegibilidade | Proporção de profissionais elegíveis que participam dos resultados. |
| Cobertura de exposição | Proporção de profissionais elegíveis que recebem exposição em contextos compatíveis. |
| Distribuição de exposição | Participação relativa de impressões entre profissionais e grupos comparáveis. |
| Concentração | Coeficiente de Gini ou outros indicadores adequados de concentração. |
| Visibilidade inicial | Exposição de profissionais sem histórico consolidado. |
| Posição | Distribuição de posições e participação nas primeiras posições dos resultados. |
| Oportunidades | Distribuição de manifestações de interesse e conexões entre profissionais comparáveis. |
| Conversão | Conversão entre impressão, consulta, interesse e conexão. |
| Relevância | Compatibilidade e utilidade percebida dos resultados apresentados. |
| Satisfação | Percepção de clientes e profissionais sobre a experiência de descoberta. |

A interpretação dos indicadores deverá considerar:

- Volume e composição da demanda.
- Critérios de elegibilidade.
- Compatibilidade dos serviços.
- Disponibilidade dos profissionais.
- Região e contexto da descoberta.
- Período de observação.
- Tamanho e representatividade das amostras.

Nenhum indicador deverá ser utilizado isoladamente como comprovação de equidade.

Os métodos de cálculo, critérios de comparação, metas e limites deverão ser estabelecidos no plano de medição e validação.

### 10.1.6 Mecanismos de distribuição de exposição

Os mecanismos de descoberta deverão iniciar com políticas compreensíveis, testáveis e proporcionais à maturidade do produto.

Para o MVP, deverão ser priorizados:

- Critérios explícitos de elegibilidade.
- Ordenação contextual por relevância.
- Participação de profissionais sem histórico consolidado.
- Mecanismos básicos de prevenção de favorecimento injustificado.
- Instrumentação necessária à avaliação da distribuição de exposição.
- Explicações compreensíveis sobre os principais fatores de classificação.

Mecanismos adicionais poderão ser avaliados conforme as evidências disponíveis:

- Diversificação entre alternativas de relevância semelhante.
- Exploração controlada de profissionais elegíveis.
- Rotação entre resultados comparáveis.
- Incentivos temporários à descoberta de novos profissionais.
- Limites de concentração de exposição.
- Ajustes de classificação fundamentados no histórico de distribuição de oportunidades.

Esses mecanismos constituem alternativas candidatas, e não obrigações simultâneas de implementação.

A adoção de cada mecanismo deverá considerar relevância, impactos sobre os usuários, complexidade operacional e evidências de benefício.

### 10.1.7 Limites e salvaguardas

As políticas de descoberta deverão observar as seguintes salvaguardas:

- Não apresentar resultados incompatíveis com critérios obrigatórios de elegibilidade.
- Não comprometer injustificadamente a relevância para aumentar indicadores de exposição.
- Não utilizar favorecimento pago durante o MVP.
- Não incentivar comportamentos manipulativos ou engajamento artificial.
- Não aplicar penalizações injustificadas por ausência de histórico.
- Avaliar possíveis impactos adversos antes de alterações relevantes.
- Manter alterações de políticas identificáveis e passíveis de revisão.
- Permitir a reversão controlada de alterações problemáticas.
- Preservar os requisitos de privacidade, segurança e autonomia.

Os limites quantitativos de degradação de relevância e demais critérios de interrupção deverão ser definidos antes da realização de experimentos que possam afetar significativamente os usuários.

## 10.2 Estrutura de confiança

### 10.2.1 Definição

Confiança corresponde à capacidade de clientes e profissionais tomarem decisões informadas sobre possíveis interações, considerando informações relevantes, evidências disponíveis, histórico contextual e mecanismos de proteção.

O RAIRE deverá fornecer condições que reduzam incertezas e riscos evitáveis, sem apresentar garantias que não possam ser sustentadas pelas capacidades da plataforma.

A confiança deverá ser bilateral, considerando tanto a perspectiva do cliente quanto a perspectiva do profissional.

### 10.2.2 Dimensões de confiança

| Dimensão | Finalidade |
| --- | --- |
| Identificação | Reduzir incertezas sobre informações de identificação e contato. |
| Competência | Permitir a apresentação e avaliação de capacidades profissionais. |
| Reputação | Disponibilizar referências contextualizadas de experiências anteriores. |
| Comportamento | Informar características relevantes de interações observadas ou declaradas. |
| Segurança | Reduzir riscos de exposição indevida, abuso e interações indesejadas. |
| Transparência | Permitir a compreensão da origem, significado e limitações das informações. |

Nenhuma dimensão deverá ser interpretada isoladamente como comprovação de segurança, competência ou confiabilidade integral de um usuário.

### 10.2.3 Origem e natureza das informações

As informações apresentadas pela plataforma deverão ser distinguíveis de acordo com sua origem e natureza.

| Classificação | Descrição |
| --- | --- |
| Declarada | Informação fornecida pelo próprio usuário. |
| Demonstrada | Informação acompanhada de conteúdo ou evidência apresentada pelo usuário, como trabalhos de portfólio. |
| Confirmada | Informação submetida a um processo específico de confirmação. |
| Observada | Informação decorrente de acontecimentos efetivamente registrados pela plataforma. |
| Avaliada | Informação originada de avaliações elegíveis realizadas por outros usuários. |

A apresentação de evidências fornecidas pelo usuário não deverá ser confundida com validação independente pela plataforma.

Informações declaradas sobre contratações e serviços realizados fora da plataforma deverão preservar sua natureza declaratória.

### 10.2.4 Confirmação e verificação

Os mecanismos de confirmação e verificação deverão possuir critérios explícitos, finalidade legítima e proporcionalidade em relação aos riscos envolvidos.

| Mecanismo | Significado | Escopo inicial |
| --- | --- | --- |
| Confirmação de e-mail | Evidência de acesso ao endereço informado. | MVP |
| Confirmação de telefone | Evidência de acesso ao número informado. | A avaliar |
| Verificação de identidade | Verificação de informações de identidade mediante procedimento específico. | Evolutivo |
| Verificação de qualificações | Verificação de documentos ou informações profissionais específicas. | Evolutivo |

A confirmação de e-mail ou telefone não deverá ser apresentada como verificação integral de identidade.

As confirmações obrigatórias para determinadas ações deverão ser diferenciadas das verificações adicionais e opcionais.

A necessidade de verificações documentais deverá ser avaliada considerando:

- Riscos associados ao contexto de utilização.
- Benefícios efetivos para clientes e profissionais.
- Fricção introduzida no processo de participação.
- Capacidade operacional de análise.
- Segurança e proteção dos documentos.
- Obrigações legais aplicáveis.
- Necessidade e proporcionalidade da retenção dos dados.

A plataforma não deverá apresentar declarações genéricas de usuário totalmente verificado ou garantias de antecedentes, competência ou comportamento sem evidências que as sustentem.

### 10.2.5 Competência e portfólio profissional

O portfólio deverá constituir um dos principais instrumentos de apresentação das competências profissionais.

A plataforma deverá favorecer:

- Apresentação de trabalhos relevantes.
- Contextualização dos trabalhos conforme os serviços oferecidos.
- Distinção entre experiências declaradas e qualificações verificadas.
- Consulta de informações que auxiliem a comparação entre profissionais.
- Transparência quanto à origem dos conteúdos apresentados.
- Proteção contra utilização indevida de imagens, dados pessoais e trabalhos de terceiros.

A ausência de um portfólio consolidado não deverá representar automaticamente ausência de competência profissional.

### 10.2.6 Reputação bilateral e contextual

O RAIRE deverá permitir a construção de reputação por clientes e profissionais a partir de experiências elegíveis originadas na plataforma.

Os mecanismos de reputação deverão:

- Distinguir avaliações recebidas na atuação como cliente e como profissional.
- Associar avaliações ao contexto da interação.
- Impedir avaliações sem vínculo com experiências consideradas elegíveis.
- Impedir duplicações e autoavaliações indevidas.
- Reduzir possibilidades de manipulação, retaliação e avaliações abusivas.
- Diferenciar ausência de histórico de histórico negativo.
- Apresentar contexto suficiente para interpretação das avaliações.
- Permitir denúncia e revisão conforme as políticas aplicáveis.

A elegibilidade das avaliações deverá considerar a natureza da experiência avaliada.

Uma conexão estabelecida poderá fundamentar uma avaliação da experiência de contato ou interação, quando permitido pelas regras aplicáveis, mas não deverá, isoladamente, fundamentar uma avaliação apresentada como experiência de serviço realizado.

A avaliação de qualidade de um serviço deverá depender de critérios específicos de elegibilidade e preservar a natureza das evidências disponíveis.

As regras de publicação, edição, contestação e agregação das avaliações deverão ser estabelecidas em uma política própria de reputação.

### 10.2.7 Histórico e sinais comportamentais

A plataforma poderá apresentar informações relevantes sobre interações anteriores, desde que sua origem, confiabilidade e significado estejam suficientemente estabelecidos.

Exemplos de informações candidatas:

- Quantidade de conexões estabelecidas.
- Histórico de respostas a solicitações.
- Avaliações recebidas em contextos elegíveis.
- Desfechos voluntariamente informados.
- Recorrência de interações relevantes.

Essas informações deverão ser avaliadas considerando volume de observações, recência, contexto e possibilidade de manipulação.

Sinais comportamentais não deverão penalizar injustificadamente usuários com menor frequência de utilização, disponibilidade limitada ou pouco histórico.

Denúncias não confirmadas não deverão ser apresentadas como evidência de comportamento inadequado.

A ausência de informações observáveis não deverá ser preenchida por inferências apresentadas como fatos.

### 10.2.8 Segurança e autonomia das interações

Os mecanismos de confiança deverão ser acompanhados de capacidades de proteção dos usuários.

A plataforma deverá:

- Preservar informações privadas de contato até autorização adequada.
- Evitar exposição desnecessária de localização precisa.
- Permitir aceitação, recusa e cancelamento de solicitações conforme as regras aplicáveis.
- Permitir bloqueio de interações indesejadas.
- Disponibilizar mecanismos de denúncia.
- Informar orientações relevantes de segurança.
- Permitir que os usuários decidam autonomamente com quem desejam estabelecer contato.
- Informar as limitações da atuação da plataforma em interações realizadas externamente.

O RAIRE não deverá apresentar a existência de uma conexão, avaliação ou verificação como garantia de segurança de uma futura interação.

## 10.3 Transparência

### 10.3.1 Transparência da descoberta

Os usuários deverão possuir acesso a informações compreensíveis sobre o funcionamento dos mecanismos de descoberta.

A plataforma deverá:

- Informar os principais critérios de elegibilidade e relevância.
- Explicar os fatores pertinentes à apresentação dos resultados.
- Informar a existência de mecanismos de distribuição equitativa de exposição.
- Diferenciar ordenação por relevância de ordenações alternativas, quando disponíveis.
- Informar a inexistência de favorecimento pago no ranking durante o MVP.
- Disponibilizar orientações sobre informações profissionais relevantes para descoberta.

As explicações deverão corresponder aos critérios efetivamente utilizados.

A transparência não implica exposição de componentes internos, dados de terceiros ou detalhes cuja divulgação comprometa a segurança e favoreça manipulação abusiva.

### 10.3.2 Transparência das informações de confiança

A plataforma deverá permitir que os usuários compreendam:

- Quais informações foram declaradas pelos próprios usuários.
- Quais informações foram efetivamente confirmadas.
- Quais informações foram observadas pela plataforma.
- O contexto e a origem das avaliações apresentadas.
- As limitações dos sinais de confiança.
- As consequências das ações de compartilhamento de informações.

A interface não deverá induzir interpretações de segurança ou qualidade superiores às evidências disponíveis.

### 10.3.3 Contestação e revisão

Os usuários deverão possuir mecanismos adequados para comunicar:

- Informações incorretas em seus perfis.
- Erros de classificação ou elegibilidade.
- Avaliações incompatíveis com as políticas da plataforma.
- Decisões de moderação consideradas inadequadas.
- Possíveis comportamentos discriminatórios.
- Irregularidades relacionadas ao tratamento de suas informações.

A contestação deverá permitir análise proporcional à natureza e à criticidade do caso.

A possibilidade de contestação não implica direito a uma posição específica nos resultados, remoção automática de avaliações desfavoráveis ou reversão obrigatória de decisões legítimas.

## 10.4 Governança das políticas de equidade e confiança

As políticas relacionadas à descoberta, reputação, verificação e moderação deverão possuir governança compatível com seu impacto sobre os usuários.

Alterações relevantes deverão considerar:

- Contexto e justificativa.
- Objetivo pretendido.
- Evidências disponíveis.
- Usuários e capacidades potencialmente afetados.
- Riscos e efeitos adversos.
- Indicadores de avaliação.
- Salvaguardas aplicáveis.
- Critérios de revisão e reversão.

As políticas deverão possuir versões identificáveis e responsáveis por sua aprovação e revisão.

Decisões administrativas que afetem significativamente a participação de usuários deverão possuir justificativa e registro adequado, observadas as restrições de privacidade.

A equipe deverá acompanhar periodicamente os indicadores de equidade, confiança, relevância e segurança.

A divulgação externa de resultados deverá considerar a maturidade dos dados, sua representatividade, a privacidade dos participantes e a possibilidade de interpretações inadequadas.

## 10.5 Escopo inicial e evolução

O MVP deverá priorizar mecanismos simples e verificáveis que permitam validar as hipóteses fundamentais de equidade e confiança.

| Capacidade | Direção inicial |
| --- | --- |
| Elegibilidade e relevância contextual | Obrigatória |
| Participação de profissionais sem histórico | Obrigatória |
| Distribuição básica de exposição | Obrigatória |
| Instrumentação de exposição e conexões | Obrigatória |
| Explicabilidade básica da descoberta | Obrigatória |
| Confirmação de e-mail | Obrigatória |
| Portfólio profissional | Obrigatória |
| Reputação bilateral contextual | Obrigatória |
| Proteção de informações de contato | Obrigatória |
| Denúncia, bloqueio e moderação | Obrigatória |
| Verificação adicional de contato | A avaliar |
| Verificação documental de identidade | Evolutiva |
| Mecanismos avançados de diversificação | Evolutivos |
| Simulação e replay de políticas de ranking | Evolutivos |
| Detecção avançada de fraude reputacional | Evolutiva |
| Relatórios públicos de equidade | Condicionados à maturidade das evidências |

A classificação deverá ser reconciliada com os requisitos funcionais, requisitos não funcionais e o escopo consolidado do MVP.

A evolução dos mecanismos deverá ser fundamentada em necessidades reais, avaliação de riscos e evidências obtidas durante a validação do produto.

---

# 11. Governança de dados, privacidade e proteção de dados pessoais

O RAIRE deverá adotar práticas de governança e proteção de dados compatíveis com a natureza das informações tratadas, os riscos associados às suas operações e os direitos dos titulares.

A utilização de dados deverá contribuir para as finalidades legítimas do produto, preservando a autonomia, a privacidade, a segurança e a capacidade de decisão dos usuários.

As políticas e operações de tratamento deverão observar a Lei Geral de Proteção de Dados Pessoais (LGPD) e demais obrigações legais e regulatórias aplicáveis.

## 11.1 Princípios

O tratamento de dados deverá observar os seguintes princípios:

- **Finalidade:** tratar informações para propósitos legítimos, específicos, explícitos e informados.
- **Adequação:** assegurar compatibilidade entre o tratamento realizado e as finalidades estabelecidas.
- **Necessidade e minimização:** limitar a coleta, utilização e retenção ao mínimo necessário.
- **Livre acesso:** permitir que os titulares consultem informações sobre o tratamento de seus dados.
- **Qualidade dos dados:** favorecer a exatidão, clareza, relevância e atualização das informações.
- **Transparência:** disponibilizar informações compreensíveis sobre coleta, utilização, exposição, compartilhamento e retenção.
- **Segurança:** proteger informações contra acesso indevido, perda, alteração, divulgação e tratamento não autorizado.
- **Prevenção:** identificar e reduzir riscos associados às operações de tratamento.
- **Não discriminação:** impedir tratamentos destinados a finalidades discriminatórias ilícitas ou abusivas.
- **Responsabilização e prestação de contas:** preservar evidências das medidas adotadas e das decisões relevantes.
- **Privacidade desde a concepção:** incorporar requisitos de proteção de dados durante a definição e evolução do produto.
- **Privacidade por padrão:** evitar exposição e coleta desnecessárias nas configurações iniciais do sistema.

## 11.2 Governança das operações de tratamento

As operações de tratamento de dados pessoais deverão ser identificadas, documentadas e avaliadas conforme suas finalidades e riscos.

O registro das operações deverá contemplar, quando aplicável:

- Identificação da operação de tratamento.
- Finalidade.
- Categorias de titulares envolvidos.
- Categorias de dados tratados.
- Base legal aplicável.
- Origem das informações.
- Responsáveis pelo tratamento.
- Destinatários e compartilhamentos.
- Critérios de acesso.
- Período e justificativa de retenção.
- Medidas de proteção.
- Riscos identificados.
- Procedimentos de eliminação ou anonimização.

As bases legais deverão ser definidas conforme a natureza de cada operação, não sendo o consentimento considerado fundamento universal para todo tratamento.

A aceitação dos termos de uso, a ciência da política de privacidade, o consentimento para tratamento de dados e a autorização de compartilhamento deverão ser tratados como manifestações distintas, conforme suas respectivas finalidades.

## 11.3 Classificação de dados

A classificação das informações deverá considerar dimensões complementares de natureza jurídica, finalidade operacional e nível de acesso.

Uma mesma informação poderá pertencer simultaneamente a diferentes classificações.

### 11.3.1 Natureza jurídica

| Categoria | Descrição | Exemplos |
| --- | --- | --- |
| Dados pessoais | Informações relacionadas a pessoas naturais identificadas ou identificáveis. | Nome, telefone, e-mail, localização e identificadores de conta. |
| Dados pessoais sensíveis | Informações enquadradas nas categorias sensíveis previstas na LGPD. | Dados de saúde, origem racial ou étnica, biometria vinculada à pessoa e convicção religiosa. |
| Dados anonimizados | Informações cuja associação a um titular não seja razoavelmente possível, observados os critérios legais. | Estatísticas adequadamente anonimizadas. |
| Dados não pessoais | Informações que não se relacionam a pessoas naturais identificadas ou identificáveis. | Configurações técnicas e definições genéricas do catálogo. |

A pseudonimização deverá ser considerada uma técnica de proteção, não uma alteração automática da natureza jurídica dos dados.

### 11.3.2 Classificação de acesso

| Classe | Descrição | Exemplos |
| --- | --- | --- |
| Pública | Informações destinadas à divulgação autorizada. | Perfil profissional, serviços, portfólio publicado e avaliações públicas. |
| Contextual | Informações disponibilizadas apenas mediante condições específicas de negócio e autorização. | Contato compartilhado após conexão e informações restritas de uma demanda. |
| Interna | Informações necessárias à operação, sem divulgação pública. | Configurações, eventos operacionais e histórico administrativo. |
| Confidencial | Informações cujo acesso indevido pode causar prejuízos aos usuários ou à organização. | E-mail, telefone privado, identificadores internos e histórico de interações. |
| Altamente restrita | Informações associadas a riscos elevados de exposição ou utilização indevida. | Documentos de identidade, localização precisa, evidências de denúncias e credenciais. |

A classificação deverá considerar os riscos associados ao acesso, compartilhamento, alteração e exposição das informações.

Dados públicos poderão continuar sujeitos às obrigações de proteção de dados pessoais.

### 11.3.3 Finalidade operacional

As informações poderão ser agrupadas conforme sua utilização no domínio:

- Identificação e acesso.
- Perfil profissional.
- Catálogo e descoberta.
- Portfólio profissional.
- Demandas de serviço.
- Interesses e conexões.
- Avaliações e reputação.
- Confirmação e verificação.
- Moderação e segurança.
- Auditoria e operação.
- Análise de produto e pesquisa.

Essa classificação deverá auxiliar na definição das responsabilidades, controles de acesso, políticas de retenção e fronteiras de tratamento.

## 11.4 Visibilidade e compartilhamento contextual

O RAIRE deverá estabelecer regras explícitas sobre quais informações podem ser acessadas, por quem e em quais circunstâncias.

| Informação | Visibilidade inicial |
| --- | --- |
| Nome de exibição profissional | Pública após publicação autorizada do perfil. |
| Serviços, bio e portfólio | Públicos conforme configuração e publicação. |
| Região aproximada de atendimento | Pública quando necessária à descoberta. |
| Localização precisa | Não pública, com tratamento restrito à finalidade necessária. |
| E-mail cadastral | Privado. |
| Telefone privado | Compartilhado apenas nas condições autorizadas. |
| Demanda de serviço | Contextual, conforme critérios de elegibilidade e publicação. |
| Identidade civil | Restrita, salvo divulgação específica autorizada. |
| Resultado de confirmação ou verificação | Divulgado de maneira proporcional e compreensível. |
| Histórico interno de conexões | Restrito às partes autorizadas e às operações legítimas da plataforma. |
| Avaliações | Conforme as regras de publicação e reputação. |
| Denúncias e evidências | Restritas às responsabilidades autorizadas. |
| Eventos analíticos | Restritos às finalidades de análise estabelecidas. |

A plataforma deverá informar previamente as consequências da publicação ou do compartilhamento de informações.

A autorização de compartilhamento de contato deverá ser associada à conexão correspondente.

O usuário deverá compreender que informações disponibilizadas a terceiros poderão ser utilizadas em canais externos e que sua recuperação posterior poderá não ser tecnicamente possível.

Os mecanismos de bloqueio, cancelamento e revogação deverão impedir novas disponibilizações indevidas, sem prometer a recuperação de informações anteriormente compartilhadas.

## 11.5 Privacidade geográfica

A utilização de informações geográficas deverá ser proporcional às necessidades de descoberta, elegibilidade e conexão local.

A plataforma deverá:

- Permitir a seleção manual de cidade, bairro ou região.
- Disponibilizar geolocalização automática como mecanismo opcional.
- Evitar a exigência de endereço residencial para participação.
- Priorizar informações aproximadas de localização e área de atendimento.
- Restringir o tratamento de coordenadas precisas às finalidades justificadas.
- Evitar a persistência de coordenadas precisas quando o processamento temporário for suficiente.
- Impedir a exposição pública de endereços residenciais e coordenadas privadas.
- Considerar riscos de inferência de localização mediante consultas sucessivas ou combinação de informações.
- Evitar coleta contínua de localização e monitoramento de deslocamentos sem necessidade demonstrada.

Endereços comerciais poderão possuir regras específicas de divulgação voluntária, observadas as finalidades e autorizações aplicáveis.

A utilização de serviços externos de mapas, geocodificação ou cálculo de distância deverá considerar os dados compartilhados com esses fornecedores e os riscos decorrentes do tratamento.

## 11.6 Portfólio, imagens e conteúdo de terceiros

A publicação de conteúdo profissional deverá observar os direitos dos titulares e de terceiros eventualmente representados.

A plataforma deverá:

- Informar as responsabilidades associadas à publicação de conteúdo.
- Orientar usuários quanto à exposição indevida de informações pessoais.
- Disponibilizar mecanismos de denúncia e remoção de conteúdo irregular.
- Evitar a exposição desnecessária de metadados de localização e identificação presentes em arquivos.
- Restringir a utilização de conteúdo para finalidades diferentes daquelas informadas.
- Considerar riscos associados a imagens de clientes, residências, documentos, crianças e ambientes privados.
- Permitir a correção, substituição e remoção de conteúdo conforme as regras aplicáveis.

A publicação de trabalhos em um portfólio não deverá implicar autorização automática para utilização desse conteúdo em campanhas promocionais, pesquisas ou treinamento de modelos de inteligência artificial.

## 11.7 Consentimento, autorizações e preferências

Quando o consentimento constituir a base legal aplicável, sua obtenção deverá observar os requisitos previstos na LGPD.

A plataforma deverá, quando aplicável:

- Apresentar finalidades específicas em linguagem compreensível.
- Permitir manifestações livres, informadas e inequívocas.
- Evitar autorizações genéricas para finalidades distintas.
- Registrar a manifestação, finalidade, momento e versão das informações apresentadas.
- Permitir a revogação por mecanismos acessíveis.
- Respeitar as consequências da revogação conforme a natureza da operação e demais fundamentos legais aplicáveis.
- Evitar condicionar funcionalidades essenciais a tratamentos opcionais desnecessários.

Autorizações de negócio relacionadas à publicação de perfis, compartilhamento de contato e estabelecimento de conexões deverão possuir regras próprias.

Preferências de comunicação, pesquisa, personalização e tratamento opcional deverão ser diferenciadas das operações essenciais ao funcionamento do produto.

## 11.8 Direitos dos titulares

O RAIRE deverá disponibilizar mecanismos adequados para o exercício dos direitos previstos na legislação aplicável, incluindo:

- Confirmação da existência de tratamento.
- Acesso aos dados pessoais.
- Correção de informações incompletas, inexatas ou desatualizadas.
- Anonimização, bloqueio ou eliminação de dados desnecessários, excessivos ou tratados irregularmente.
- Portabilidade, conforme regulamentação e condições aplicáveis.
- Eliminação de dados tratados com consentimento, ressalvadas as hipóteses legais de conservação.
- Informação sobre compartilhamento com entidades públicas e privadas.
- Informação sobre a possibilidade de não fornecer consentimento e suas consequências.
- Revogação de consentimento.
- Oposição ao tratamento nas hipóteses legalmente aplicáveis.
- Solicitação de revisão de decisões tomadas unicamente com base em tratamento automatizado de dados pessoais, quando aplicável.

Os procedimentos deverão considerar:

- Identificação proporcional do solicitante.
- Proteção contra solicitações fraudulentas.
- Registro e acompanhamento da solicitação.
- Prazos legais aplicáveis.
- Comunicação clara sobre o resultado.
- Justificativa para eventuais limitações.
- Proteção dos dados e direitos de terceiros.

A plataforma deverá disponibilizar um canal acessível para comunicação relacionada ao tratamento de dados pessoais.

## 11.9 Ciclo de vida e retenção de dados

As informações deverão possuir critérios de retenção fundamentados em sua finalidade, necessidade operacional, riscos e obrigações legais.

O ciclo de vida deverá considerar:

1. Coleta.
2. Utilização.
3. Armazenamento.
4. Compartilhamento autorizado.
5. Arquivamento, quando necessário.
6. Eliminação ou anonimização.
7. Expiração das cópias residuais, conforme políticas aplicáveis.

### 11.9.1 Diretrizes iniciais de retenção

| Categoria | Diretriz inicial |
| --- | --- |
| Dados cadastrais | Manutenção durante a relação com o usuário, observadas as finalidades e obrigações aplicáveis. |
| Perfis e portfólios | Manutenção enquanto publicados ou necessários à relação com o usuário. |
| Localização precisa | Processamento temporário por padrão; persistência somente mediante justificativa. |
| Demandas e conexões | Retenção compatível com as finalidades de prestação, histórico e segurança. |
| Avaliações e reputação | Retenção conforme política de reputação, direitos dos titulares e integridade das informações. |
| Eventos analíticos | Retenção limitada ao período necessário para análise, com posterior eliminação ou anonimização adequada. |
| Registros de acesso | Retenção conforme obrigações legais aplicáveis. |
| Logs técnicos | Retenção proporcional às necessidades de diagnóstico e segurança. |
| Auditoria | Retenção conforme criticidade, finalidade e obrigações aplicáveis. |
| Denúncias e evidências | Retenção conforme natureza do caso, riscos, necessidade de defesa e obrigações legais. |
| Documentos de verificação | Não coletados por padrão no MVP; eventual retenção limitada à finalidade justificada. |
| Backups | Expiração conforme política de continuidade e recuperação operacional. |

Os prazos definitivos deverão ser estabelecidos em uma matriz de retenção antes do tratamento de dados reais no piloto.

Nenhuma categoria deverá possuir retenção indefinida sem fundamento identificado.

### 11.9.2 Encerramento de conta e eliminação

O encerramento de uma conta deverá resultar na interrupção de sua participação ativa e na aplicação dos procedimentos pertinentes de eliminação, anonimização ou conservação justificada.

Os procedimentos deverão considerar:

- Remoção de exposição pública quando aplicável.
- Revogação de sessões e credenciais.
- Interrupção de tratamentos que perderam sua finalidade.
- Identificação das informações sujeitas à conservação legalmente justificada.
- Restrição de acesso aos dados conservados.
- Propagação das medidas aos fornecedores pertinentes.
- Tratamento de caches, índices, arquivos e demais cópias.
- Expiração controlada de cópias de segurança.
- Prevenção da reintrodução indevida de dados eliminados após restaurações.

O encerramento de conta não deverá ser confundido com eliminação imediata e irrestrita de todos os registros associados.

As exceções de conservação deverão possuir fundamento, finalidade e acesso controlado.

## 11.10 Dados analíticos, pesquisa e equidade

A instrumentação do produto deverá permitir a avaliação de suas hipóteses, qualidade e impactos sem introduzir tratamentos desproporcionais de dados pessoais.

Os mecanismos analíticos deverão:

- Possuir finalidades documentadas.
- Limitar informações coletadas ao necessário.
- Evitar inclusão de credenciais, contatos privados, localização precisa e conteúdo pessoal desnecessário em eventos.
- Utilizar pseudonimização quando adequada.
- Restringir mecanismos de associação entre eventos e titulares.
- Priorizar agregação e anonimização para relatórios e divulgação.
- Considerar riscos de reidentificação em grupos pequenos.
- Respeitar as bases legais e preferências aplicáveis.
- Evitar utilização secundária incompatível com as finalidades informadas.

A avaliação de equidade deverá priorizar características operacionais pertinentes, conforme estabelecido na seção 10.

A coleta de dados pessoais sensíveis para pesquisas ou avaliações de discriminação não deverá ocorrer no MVP sem avaliação específica de necessidade, proporcionalidade, base legal e riscos.

Relatórios externos não deverão permitir a identificação indevida de indivíduos por combinação de métricas, regiões, categorias ou características de participação.

Dados reais de usuários não deverão ser utilizados em ambientes de desenvolvimento, demonstrações ou ferramentas externas de inteligência artificial sem avaliação e controles adequados.

## 11.11 Compartilhamento e fornecedores externos

A utilização de serviços externos deverá considerar suas responsabilidades e as operações de tratamento realizadas.

Exemplos de fornecedores potencialmente envolvidos:

- Infraestrutura e hospedagem.
- Armazenamento e distribuição de imagens.
- Autenticação.
- Envio de e-mails e notificações.
- Mapas e geocodificação.
- Monitoramento e observabilidade.
- Análise de produto.
- Ferramentas de suporte operacional.

A seleção e utilização desses fornecedores deverá considerar:

- Finalidade e necessidade do compartilhamento.
- Categorias de dados envolvidos.
- Responsabilidades das partes.
- Medidas de proteção.
- Condições contratuais.
- Retenção e eliminação.
- Subcontratação.
- Localização do tratamento.
- Transferências internacionais de dados.
- Procedimentos de resposta a incidentes.
- Capacidade de atendimento aos direitos dos titulares.

Transferências internacionais deverão observar as condições e os mecanismos legalmente aplicáveis.

## 11.12 Segurança e controle de acesso

A proteção dos dados deverá observar sua classificação, finalidade e criticidade.

O sistema deverá:

- Aplicar o princípio do menor privilégio.
- Restringir acessos administrativos conforme responsabilidades autorizadas.
- Proteger informações em trânsito e em repouso de maneira proporcional aos riscos.
- Evitar exposição de credenciais, tokens e informações pessoais desnecessárias em registros técnicos.
- Registrar acessos e operações sensíveis quando necessário.
- Preservar mecanismos seguros de autenticação e autorização.
- Restringir exportações de informações pessoais.
- Proteger cópias de segurança.
- Permitir a identificação e investigação de acessos indevidos.
- Considerar riscos de coleta automatizada, enumeração e extração massiva de informações públicas.

Os mecanismos técnicos de proteção deverão ser definidos em conformidade com os requisitos não funcionais de segurança e privacidade.

## 11.13 Gestão de incidentes de privacidade

O RAIRE deverá possuir procedimentos proporcionais aos riscos de suas operações para identificação, avaliação e tratamento de incidentes envolvendo dados pessoais.

Os procedimentos deverão contemplar:

- Identificação e registro do incidente.
- Avaliação da natureza e extensão do comprometimento.
- Contenção e mitigação.
- Preservação segura de evidências.
- Identificação dos dados e titulares potencialmente afetados.
- Avaliação de riscos e danos relevantes.
- Comunicação às autoridades e titulares, quando legalmente exigida.
- Registro das medidas adotadas.
- Revisão das causas e prevenção de recorrência.

As responsabilidades, critérios de escalonamento e prazos deverão observar as obrigações legais e regulatórias aplicáveis.

## 11.14 Responsabilidades e governança de privacidade

Antes da disponibilização do MVP para usuários reais, deverão ser identificados:

- O agente responsável pelas decisões relativas ao tratamento.
- Os operadores de dados pessoais envolvidos.
- As responsabilidades internas de proteção de dados.
- O canal de comunicação com titulares.
- A necessidade de indicação de encarregado, conforme as regras aplicáveis.
- As responsabilidades relacionadas a fornecedores externos.
- Os procedimentos de aprovação de novas operações de tratamento.
- A necessidade de elaboração de Relatório de Impacto à Proteção de Dados Pessoais (RIPD).

Alterações relevantes nas finalidades, categorias de dados ou mecanismos de exposição deverão ser avaliadas antes de sua implementação.

A governança deverá ser proporcional à maturidade do produto, sem dispensar as obrigações legais aplicáveis.

## 11.15 Condições mínimas para o piloto

Antes do tratamento de dados pessoais de usuários reais, deverão estar estabelecidos:

- Inventário inicial das operações de tratamento.
- Identificação das finalidades e bases legais aplicáveis.
- Política de privacidade e termos de uso.
- Classificação das informações e regras de visibilidade.
- Matriz de acesso e responsabilidades.
- Política inicial de retenção e eliminação.
- Procedimentos de atendimento aos direitos dos titulares.
- Procedimentos de resposta a incidentes.
- Avaliação dos fornecedores envolvidos.
- Avaliação dos riscos de privacidade das jornadas principais.
- Controles essenciais de segurança e proteção de dados.

As definições jurídicas e regulatórias deverão ser avaliadas por profissional qualificado antes do início do piloto.

A conformidade deverá ser tratada como uma responsabilidade contínua, acompanhando a evolução das operações do RAIRE.

---

# 12. Métricas de sucesso, validação e experimentação

A estratégia de medição do RAIRE deverá permitir a avaliação da proposta de valor, da utilidade das capacidades oferecidas, da eficiência dos mecanismos de descoberta e conexão e dos impactos produzidos sobre clientes e profissionais.

As métricas deverão orientar decisões de produto, identificação de problemas, priorização de melhorias e validação das hipóteses estabelecidas.

A instrumentação não deverá incentivar crescimento artificial de engajamento, coleta desnecessária de dados ou otimização de indicadores em detrimento da utilidade, equidade, confiança, segurança e autonomia dos usuários.

## 12.1 Objetivos da medição

A estratégia de medição deverá permitir responder às seguintes perguntas:

- O problema identificado possui relevância suficiente para justificar a existência do produto?
- Clientes conseguem encontrar profissionais adequados às suas necessidades?
- Profissionais conseguem acessar oportunidades de trabalho relevantes?
- Os mecanismos de descoberta resultam em conexões úteis?
- Informações profissionais, portfólios e reputação contribuem para decisões informadas?
- Existe densidade suficiente de oferta e demanda nos segmentos atendidos?
- A plataforma reduz barreiras injustificadas de participação e exposição?
- Os usuários percebem valor, confiança e autonomia durante a utilização?
- A operação é sustentável considerando os recursos disponíveis?

## 12.2 Hierarquia de métricas

As métricas deverão ser organizadas nas seguintes categorias:

| Categoria | Finalidade |
| --- | --- |
| Métrica central de valor | Representar a entrega de valor fundamental do produto. |
| Métricas de resultado | Avaliar os benefícios percebidos ou declarados pelos participantes. |
| Métricas direcionadoras | Identificar fatores que contribuem para a geração de valor. |
| Métricas diagnósticas | Identificar fricções, gargalos e oportunidades de melhoria. |
| Guardrails | Identificar impactos adversos e preservar propriedades essenciais. |
| Critérios de aceitação | Verificar requisitos obrigatórios de qualidade e conformidade. |
| Métricas de aprendizado | Avaliar a evolução das evidências sobre as hipóteses do produto. |

Nenhuma métrica deverá ser interpretada isoladamente como comprovação de sucesso.

## 12.3 Métrica central de valor

A métrica central candidata do RAIRE será a quantidade semanal de conexões úteis estabelecidas entre clientes e profissionais.

Uma conexão útil corresponde a uma conexão consentida para a qual existam indicações de benefício relacionado à descoberta de um profissional ou ao acesso a uma oportunidade de serviço.

A definição operacional deverá estabelecer:

- Critérios de identificação de uma conexão.
- Critérios de avaliação da utilidade.
- Origem e confiabilidade das evidências.
- Janela temporal de observação.
- Tratamento de conexões duplicadas.
- Tratamento de respostas ausentes ou divergentes.
- Distinção entre resultados observados e declarados.

A quantidade de conexões estabelecidas não deverá ser automaticamente interpretada como quantidade de contratações ou serviços realizados.

Enquanto os mecanismos de avaliação de utilidade não estiverem suficientemente estabelecidos, o volume de conexões consentidas poderá ser utilizado como indicador intermediário, acompanhado por pesquisas qualitativas e métricas de qualidade.

A adequação da métrica central deverá ser reavaliada durante o piloto.

## 12.4 Métricas de oferta e participação profissional

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-OFE-01 | Profissionais com perfil publicado | Avaliar a formação inicial da oferta. |
| MET-OFE-02 | Taxa de conclusão do onboarding | Identificar barreiras de participação. |
| MET-OFE-03 | Tempo para publicação do perfil | Avaliar a eficiência do processo de entrada. |
| MET-OFE-04 | Profissionais ativos por região e categoria | Avaliar a densidade da oferta. |
| MET-OFE-05 | Profissionais disponíveis para atendimento | Avaliar a capacidade efetiva de atendimento. |
| MET-OFE-06 | Profissionais com portfólio publicado | Avaliar a adoção dos mecanismos de apresentação profissional. |
| MET-OFE-07 | Profissionais que recebem oportunidades relevantes | Avaliar o alcance do valor oferecido. |
| MET-OFE-08 | Continuidade da participação profissional | Avaliar a permanência de profissionais elegíveis ao longo do tempo. |

A quantidade total de profissionais cadastrados não deverá ser considerada suficiente para avaliar a disponibilidade efetiva de serviços.

## 12.5 Métricas de demanda

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-DEM-01 | Clientes com intenção de encontrar serviços | Avaliar a demanda observável. |
| MET-DEM-02 | Buscas realizadas por região e categoria | Identificar necessidades e distribuição da demanda. |
| MET-DEM-03 | Demandas de serviço publicadas | Avaliar a utilização do mural de necessidades. |
| MET-DEM-04 | Taxa de conclusão da publicação de demanda | Identificar fricções no processo. |
| MET-DEM-05 | Demandas que recebem interesse relevante | Avaliar a capacidade de aproximação entre oferta e demanda. |
| MET-DEM-06 | Demandas encerradas sem conexão | Identificar necessidades não atendidas pela plataforma. |
| MET-DEM-07 | Recorrência de utilização quando surgem novas necessidades | Avaliar a utilidade continuada do produto. |

A recorrência deverá ser interpretada considerando a natureza eventual de determinadas necessidades de serviço.

A ausência de utilização frequente não deverá ser automaticamente interpretada como falta de valor.

## 12.6 Métricas de descoberta

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-DES-01 | Buscas com profissionais elegíveis | Avaliar a cobertura da oferta. |
| MET-DES-02 | Buscas com resultados relevantes | Avaliar a adequação dos mecanismos de descoberta. |
| MET-DES-03 | Buscas sem resultados | Identificar lacunas de oferta ou problemas de classificação. |
| MET-DES-04 | Tempo até identificação de alternativa adequada | Avaliar a eficiência da descoberta. |
| MET-DES-05 | Conversão de busca em consulta de perfil | Avaliar o interesse inicial nos resultados. |
| MET-DES-06 | Conversão de consulta de perfil em manifestação de interesse | Avaliar a progressão em direção à conexão. |
| MET-DES-07 | Utilização de informações de portfólio | Avaliar sua contribuição para decisões. |
| MET-DES-08 | Relevância percebida pelos usuários | Avaliar a compatibilidade entre resultados e necessidades. |
| MET-DES-09 | Profissionais que encontram demandas compatíveis | Avaliar a eficiência da descoberta de oportunidades. |

A existência de resultados não deverá ser confundida com relevância percebida.

Cliques e visualizações deverão ser tratados como indicadores intermediários, e não como resultados finais de sucesso.

## 12.7 Métricas de conexão e resultados

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-CON-01 | Conexões consentidas estabelecidas | Avaliar a capacidade de aproximação entre as partes. |
| MET-CON-02 | Conversão de interesse em conexão | Avaliar a progressão das manifestações de interesse. |
| MET-CON-03 | Tempo até primeiro interesse relevante | Avaliar a velocidade de resposta à demanda. |
| MET-CON-04 | Tempo até estabelecimento de conexão | Avaliar a eficiência das jornadas. |
| MET-CON-05 | Conexões úteis declaradas | Avaliar a geração de valor percebido. |
| MET-CON-06 | Contratações declaradas | Investigar resultados posteriores à conexão. |
| MET-CON-07 | Serviços realizados declarados | Investigar desfechos das interações. |
| MET-CON-08 | Satisfação bilateral com as conexões | Avaliar a qualidade percebida pelas partes. |
| MET-CON-09 | Conexões canceladas ou expiradas | Identificar dificuldades de progressão. |
| MET-CON-10 | Cobertura de feedback sobre desfechos | Avaliar a representatividade das informações coletadas. |

As métricas relacionadas a contratações, pagamentos e serviços realizados deverão preservar sua natureza declaratória quando os acontecimentos ocorrerem fora da plataforma.

A ausência de confirmação não deverá ser automaticamente interpretada como contratação malsucedida.

## 12.8 Métricas de liquidez do mercado

A liquidez corresponde à capacidade de oferta e demanda encontrarem contrapartes compatíveis em condições e prazos adequados.

O RAIRE deverá avaliar a liquidez considerando a segmentação geográfica, as categorias de serviço e as características das necessidades atendidas.

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-LIQ-01 | Oferta elegível por segmento de mercado | Avaliar disponibilidade de profissionais. |
| MET-LIQ-02 | Volume de demandas por segmento | Avaliar densidade da demanda. |
| MET-LIQ-03 | Proporção de demandas com conexão estabelecida | Avaliar a capacidade de atendimento. |
| MET-LIQ-04 | Tempo até identificação de contraparte compatível | Avaliar a eficiência do mercado. |
| MET-LIQ-05 | Profissionais que recebem oportunidades relevantes | Avaliar a distribuição efetiva da demanda. |
| MET-LIQ-06 | Necessidades sem oferta compatível | Identificar desequilíbrios e lacunas. |
| MET-LIQ-07 | Profissionais elegíveis sem acesso a oportunidades | Identificar insuficiência ou concentração de demanda. |

A avaliação deverá evitar agregações que ocultem desequilíbrios relevantes entre categorias, regiões ou períodos de utilização.

A expansão do produto deverá considerar evidências de liquidez nos segmentos inicialmente atendidos.

## 12.9 Métricas de confiança e reputação

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-TRU-01 | Confiança percebida pelos clientes | Avaliar segurança e clareza percebidas na decisão. |
| MET-TRU-02 | Confiança percebida pelos profissionais | Avaliar segurança e qualidade percebidas das oportunidades. |
| MET-TRU-03 | Cobertura de avaliações elegíveis | Avaliar a formação do histórico reputacional. |
| MET-TRU-04 | Utilidade percebida das avaliações | Avaliar sua contribuição para decisões. |
| MET-TRU-05 | Utilidade percebida do portfólio | Avaliar a apresentação das competências profissionais. |
| MET-TRU-06 | Avaliações contestadas | Identificar problemas nas políticas de reputação. |
| MET-TRU-07 | Denúncias relacionadas a manipulação reputacional | Identificar riscos de integridade. |
| MET-TRU-08 | Compreensão dos sinais de verificação | Avaliar a transparência das informações de confiança. |

A confiança percebida deverá ser analisada em conjunto com indicadores de segurança e evidências observáveis.

Uma percepção elevada de confiança não constitui, isoladamente, comprovação de segurança.

## 12.10 Métricas de equidade

A avaliação de equidade deverá observar os princípios e restrições estabelecidos na seção 10.

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-EQU-01 | Cobertura de exposição entre profissionais elegíveis | Identificar invisibilidade persistente. |
| MET-EQU-02 | Distribuição de impressões entre profissionais comparáveis | Avaliar concentração de exposição. |
| MET-EQU-03 | Coeficiente de Gini de exposição | Avaliar desigualdade em contextos comparáveis. |
| MET-EQU-04 | Exposição de profissionais sem histórico | Avaliar barreiras iniciais de descoberta. |
| MET-EQU-05 | Distribuição de posições nos resultados | Avaliar diferenças de visibilidade. |
| MET-EQU-06 | Distribuição de interesses e conexões | Avaliar oportunidades além das impressões. |
| MET-EQU-07 | Conversão por grupos operacionais comparáveis | Identificar diferenças na progressão das jornadas. |
| MET-EQU-08 | Relevância percebida após mudanças de distribuição | Identificar possíveis efeitos adversos. |

Os indicadores deverão considerar elegibilidade, composição da oferta, disponibilidade, contexto geográfico, categoria e volume de demanda.

Metas relativas à exposição de profissionais novos deverão considerar sua participação no conjunto de profissionais elegíveis.

Nenhum percentual fixo de exposição ou valor de concentração deverá ser estabelecido sem justificativa contextual.

## 12.11 Indicadores de segurança, privacidade e moderação

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-SEG-01 | Incidentes de segurança e privacidade identificados | Monitorar acontecimentos adversos. |
| MET-SEG-02 | Gravidade e impacto dos incidentes | Avaliar riscos efetivos aos usuários. |
| MET-SEG-03 | Tentativas de acesso não autorizado detectadas | Avaliar ameaças e controles preventivos. |
| MET-SEG-04 | Denúncias recebidas por natureza | Identificar padrões de abuso. |
| MET-SEG-05 | Denúncias confirmadas | Avaliar ocorrências fundamentadas. |
| MET-SEG-06 | Tempo de tratamento de denúncias | Avaliar capacidade operacional. |
| MET-SEG-07 | Reincidência de comportamentos abusivos identificados | Avaliar efetividade das medidas aplicadas. |
| MET-SEG-08 | Solicitações de titulares atendidas nos prazos aplicáveis | Avaliar procedimentos de proteção de dados. |
| MET-SEG-09 | Falhas identificadas em controles críticos | Avaliar riscos e necessidades de correção. |

O objetivo deverá ser prevenir incidentes e reduzir riscos evitáveis.

A ausência de incidentes ou denúncias registradas não deverá ser interpretada como comprovação de ausência de problemas.

A redução artificial de denúncias, a criação de barreiras à comunicação de problemas ou a desconsideração de incidentes não deverão ser utilizadas para melhorar indicadores.

Incidentes críticos deverão possuir critérios próprios de escalonamento e resposta, independentemente de metas agregadas.

## 12.12 Métricas de usabilidade e inclusão

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-USA-01 | Taxa de conclusão das jornadas essenciais | Avaliar a operabilidade do produto. |
| MET-USA-02 | Tempo de conclusão de tarefas | Avaliar eficiência de utilização. |
| MET-USA-03 | Frequência de erros e abandonos | Identificar fricções. |
| MET-USA-04 | Necessidade de assistência externa | Identificar barreiras de autonomia. |
| MET-USA-05 | Facilidade de utilização percebida | Avaliar a experiência dos usuários. |
| MET-USA-06 | Barreiras de acessibilidade identificadas | Orientar correções e melhorias. |
| MET-USA-07 | Desempenho das jornadas em dispositivos representativos | Avaliar adequação às condições de acesso. |

Os requisitos de acessibilidade, desempenho, segurança e privacidade deverão possuir critérios de aceitação próprios, conforme estabelecido nos requisitos não funcionais.

O atendimento aos critérios aplicáveis de acessibilidade não deverá ser tratado como variável opcional de crescimento ou experimentação.

## 12.13 Métricas de sustentabilidade operacional

| ID | Indicador | Finalidade |
| --- | --- | --- |
| MET-OPE-01 | Custo operacional por período | Avaliar sustentabilidade financeira. |
| MET-OPE-02 | Custo de aquisição de participantes, quando mensurável | Avaliar estratégias de formação de oferta e demanda. |
| MET-OPE-03 | Custo por conexão estabelecida | Avaliar eficiência operacional. |
| MET-OPE-04 | Esforço de suporte e moderação | Identificar necessidades operacionais. |
| MET-OPE-05 | Custo de infraestrutura por usuário ativo | Avaliar eficiência tecnológica. |
| MET-OPE-06 | Esforço de onboarding assistido | Avaliar barreiras de participação e expansão. |

Esses indicadores deverão orientar decisões de sustentabilidade, expansão e avaliação futura de modelos de negócio.

O MVP não pressupõe a existência de receitas, taxas ou mecanismos de monetização.

## 12.14 Métricas de aprendizado

A evolução do produto deverá ser acompanhada pela qualidade das evidências obtidas durante suas atividades de validação.

| Indicador | Finalidade |
| --- | --- |
| Hipóteses prioritárias investigadas | Avaliar a execução do plano de validação. |
| Hipóteses com evidências registradas | Preservar rastreabilidade do aprendizado. |
| Hipóteses corroboradas | Identificar proposições apoiadas pelas evidências disponíveis. |
| Hipóteses refutadas | Identificar proposições incompatíveis com os resultados observados. |
| Hipóteses inconclusivas | Identificar insuficiência ou inconsistência das evidências. |
| Decisões fundamentadas em evidências | Avaliar a utilização efetiva do aprendizado. |
| Problemas relevantes identificados | Orientar revisão de escopo e estratégias. |

A refutação de uma hipótese deverá ser considerada um resultado legítimo de aprendizado.

A equipe não deverá utilizar a quantidade de hipóteses corroboradas como incentivo para confirmação artificial das expectativas iniciais.

## 12.15 Definição e governança das métricas

Cada métrica deverá possuir uma definição operacional documentada.

A especificação deverá contemplar, quando aplicável:

- Identificador e nome.
- Finalidade.
- Hipótese ou objetivo associado.
- Unidade de análise.
- Fórmula de cálculo.
- Numerador e denominador.
- Eventos ou fontes de informação.
- Critérios de inclusão e exclusão.
- Janela temporal de observação.
- Segmentações relevantes.
- Tratamento de dados ausentes.
- Baseline.
- Meta ou faixa esperada.
- Limitações conhecidas.
- Responsável pela definição.
- Periodicidade de avaliação.
- Versão da definição.

Alterações relevantes nas definições deverão ser identificáveis, evitando comparações entre períodos calculados mediante critérios incompatíveis.

## 12.16 Estratégia de experimentação

A experimentação deverá ser orientada pelas incertezas mais relevantes para a continuidade e evolução do RAIRE.

Os métodos deverão ser selecionados conforme a natureza das hipóteses, maturidade do produto, disponibilidade de dados e riscos envolvidos.

Métodos possíveis:

- Entrevistas semiestruturadas.
- Observação de comportamentos.
- Pesquisas qualitativas e quantitativas.
- Testes de conceito.
- Testes de protótipo.
- Testes de usabilidade.
- Pilotos assistidos.
- Análise de funil.
- Análise de grupos operacionais.
- Simulações com dados sintéticos.
- Avaliações offline.
- Experimentos controlados, quando viáveis.

Testes A/B não deverão constituir o método padrão de validação quando não houver volume de utilização suficiente ou condições adequadas de comparação.

## 12.17 Hipóteses e experimentos prioritários

### Fase 1 — Validação do problema

| ID | Hipótese | Método sugerido | Evidência esperada |
| --- | --- | --- | --- |
| EXP-01 | Clientes possuem dificuldades relevantes para encontrar profissionais locais adequados. | Entrevistas e análise de jornadas existentes. | Problemas recorrentes, consequências e limitações dos canais atuais. |
| EXP-02 | Profissionais enfrentam barreiras relevantes para divulgação e aquisição de clientes. | Entrevistas e investigação de canais de aquisição. | Evidências de oportunidades perdidas e necessidades não atendidas. |
| EXP-03 | Informações disponíveis atualmente são insuficientes para decisões informadas. | Entrevistas e comparação de experiências existentes. | Identificação de lacunas de informação, confiança e comparação. |

### Fase 2 — Validação da proposta de valor

| ID | Hipótese | Método sugerido | Evidência esperada |
| --- | --- | --- | --- |
| EXP-04 | Uma experiência especializada de descoberta local facilita a identificação de profissionais adequados. | Testes de protótipo e tarefas orientadas. | Melhoria percebida de facilidade, relevância e eficiência. |
| EXP-05 | Profissionais percebem valor em um perfil especializado e portfólio contextualizado. | Testes de conceito e onboarding assistido. | Utilidade percebida e disposição para utilização. |
| EXP-06 | Clientes percebem valor na publicação de necessidades. | Protótipo ou piloto assistido. | Disposição para publicar demandas e avaliar interessados. |
| EXP-07 | Profissionais percebem valor na descoberta ativa de oportunidades. | Piloto assistido. | Interesse em consultar e responder a demandas compatíveis. |
| EXP-08 | O compartilhamento consentido de contato reduz barreiras de confiança sem fricção excessiva. | Testes de usabilidade e entrevistas. | Compreensão, aceitação e segurança percebida. |

### Fase 3 — Validação do mercado e das conexões

| ID | Hipótese | Método sugerido | Evidência esperada |
| --- | --- | --- | --- |
| EXP-09 | Existe densidade suficiente de oferta e demanda nos segmentos selecionados. | Piloto regional e análise segmentada. | Existência recorrente de contrapartes compatíveis. |
| EXP-10 | A descoberta direta gera conexões úteis. | Piloto operacional e análise de jornadas. | Conexões estabelecidas e utilidade bilateral declarada. |
| EXP-11 | O mural de necessidades gera oportunidades e conexões úteis. | Piloto operacional. | Demandas atendidas por interesses relevantes e conexões. |
| EXP-12 | Portfólio e reputação contextual contribuem para decisões informadas. | Testes comparativos e feedback qualitativo. | Maior clareza e utilidade percebida das informações. |
| EXP-13 | Os usuários voltam a utilizar o produto quando surgem necessidades ou oportunidades relevantes. | Observação longitudinal. | Recorrência compatível com as características do mercado. |

### Fase 4 — Otimização dos mecanismos

| ID | Hipótese | Método sugerido | Evidência esperada |
| --- | --- | --- | --- |
| EXP-14 | A exploração controlada de novos profissionais reduz invisibilidade sem comprometer relevância. | Avaliação offline e experimento controlado, quando viável. | Melhoria de cobertura de exposição com relevância preservada. |
| EXP-15 | Explicações da classificação aumentam compreensão e confiança. | Testes de usabilidade e comparação controlada. | Melhor compreensão e maior confiança informada. |
| EXP-16 | A importância da proximidade varia conforme a categoria e o contexto do serviço. | Análise segmentada de comportamento. | Evidências sobre a relação entre distância, compatibilidade e conexão. |
| EXP-17 | Reputação contextual contribui mais para decisões do que indicadores exclusivamente agregados. | Testes comparativos. | Melhor compreensão, utilidade e qualidade percebida da decisão. |
| EXP-18 | Mecanismos de distribuição de exposição reduzem concentrações injustificadas. | Simulação, avaliação offline e experimentação controlada. | Melhoria de equidade sem impactos adversos relevantes. |

A execução dos experimentos deverá respeitar dependências entre hipóteses.

Experimentos de otimização não deverão substituir a investigação de problemas fundamentais de adoção, utilidade ou liquidez.

## 12.18 Protocolo de experimentação

Cada experimento deverá possuir documentação proporcional à sua complexidade e impacto.

A especificação deverá estabelecer:

- Identificador.
- Contexto e problema investigado.
- Hipótese.
- Justificativa e prioridade.
- Público ou unidade experimental.
- Método de investigação.
- Métrica principal.
- Métricas secundárias.
- Baseline, quando disponível.
- Critérios de sucesso.
- Salvaguardas e limites aceitáveis.
- Duração ou condição de encerramento.
- Requisitos de amostragem, quando aplicáveis.
- Riscos e possíveis vieses.
- Avaliação de privacidade e equidade.
- Critérios de interrupção.
- Responsável.
- Resultados e limitações.
- Conclusão e decisões decorrentes.

Os critérios de sucesso deverão ser definidos antes da execução, evitando alteração retrospectiva motivada pelos resultados observados.

Experimentos envolvendo usuários reais deverão possuir avaliação proporcional aos riscos e respeitar as restrições de tratamento de dados.

## 12.19 Metas iniciais e estabelecimento de baseline

As metas quantitativas deverão ser fundamentadas em evidências, expectativas justificadas, características do mercado ou requisitos técnicos e regulatórios aplicáveis.

A estratégia inicial deverá contemplar:

1. Definição operacional dos indicadores.
2. Identificação das fontes de informação.
3. Instrumentação mínima necessária.
4. Coleta inicial de evidências.
5. Estabelecimento de baseline.
6. Definição de metas e limites.
7. Avaliação periódica dos resultados.
8. Revisão de objetivos e decisões de produto.

Metas preliminares poderão ser utilizadas como referências de planejamento, desde que identificadas como estimativas e acompanhadas de suas justificativas.

A definição de metas após o estabelecimento de baseline não deverá resultar na alteração retrospectiva dos critérios de experimentos já iniciados.

## 12.20 Critérios de avaliação do piloto

A avaliação do piloto deverá considerar conjuntamente:

- Relevância do problema investigado.
- Utilidade percebida por clientes e profissionais.
- Capacidade de formação de oferta e demanda.
- Liquidez nos segmentos atendidos.
- Eficiência das jornadas de descoberta e conexão.
- Qualidade das conexões estabelecidas.
- Segurança e confiança das interações.
- Distribuição de oportunidades.
- Acessibilidade e inclusão.
- Sustentabilidade operacional.
- Qualidade e suficiência das evidências obtidas.

Os resultados poderão fundamentar as seguintes decisões:

| Decisão | Condição geral |
| --- | --- |
| Continuar | Evidências suficientes de valor e viabilidade, sem impedimentos críticos. |
| Ajustar | Evidências de valor acompanhadas de problemas relevantes e potencialmente solucionáveis. |
| Reduzir escopo | Valor demonstrado apenas em determinados segmentos ou capacidades. |
| Revisar hipótese | Evidências incompatíveis ou insuficientes em relação à proposta inicial. |
| Expandir | Evidências consistentes de valor, liquidez e capacidade operacional. |
| Suspender | Riscos críticos, inviabilidade operacional ou ausência persistente de evidências de valor. |

Os critérios específicos deverão ser estabelecidos antes do início do piloto e registrados no plano de validação.

O sucesso do piloto não deverá ser interpretado exclusivamente como crescimento de usuários, volume de interações ou cumprimento de metas isoladas.

## 12.21 Governança dos resultados

Os resultados das atividades de medição e experimentação deverão ser registrados de forma rastreável.

A documentação deverá preservar:

- Hipóteses investigadas.
- Métodos utilizados.
- Evidências observadas.
- Dados e critérios de interpretação.
- Limitações e possíveis vieses.
- Resultados inconclusivos ou negativos.
- Decisões tomadas.
- Questões que permanecem abertas.

A divulgação de resultados deverá respeitar as políticas de privacidade e proteção de dados.

As conclusões deverão ser proporcionais à qualidade, representatividade e suficiência das evidências disponíveis.

---

# 13. Riscos e estratégias de mitigação

O RAIRE está sujeito a riscos estratégicos, mercadológicos, operacionais, tecnológicos, jurídicos e sociais que poderão comprometer sua capacidade de entregar valor, preservar a segurança dos participantes e sustentar sua evolução.

A gestão de riscos deverá orientar decisões de produto, desenvolvimento, operação, validação e expansão.

As estratégias de mitigação deverão ser proporcionais à probabilidade, ao impacto, à capacidade de detecção e aos recursos disponíveis, considerando o estágio de maturidade do produto.

## 13.1 Riscos estratégicos e de mercado

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-MKT-01 | Problema insuficientemente relevante | Baixa demanda pela solução. | Validação do problema mediante entrevistas, observação de comportamentos e investigação dos canais existentes. |
| RSK-MKT-02 | Ausência de product-market fit | Baixa adoção, retenção e geração de valor. | Validação incremental da proposta de valor, testes de conceito e pilotos assistidos. |
| RSK-MKT-03 | Resistência à adoção | Preferência por indicações, redes sociais e canais de comunicação existentes. | Reduzir barreiras de entrada e demonstrar vantagens concretas de descoberta, comparação e acesso a oportunidades. |
| RSK-MKT-04 | Cold start e baixa liquidez | Incapacidade de conectar oferta e demanda em condições adequadas. | Concentração geográfica, seleção criteriosa de categorias, curadoria inicial da oferta e aquisição direcionada de demanda. |
| RSK-MKT-05 | Desequilíbrio entre oferta e demanda | Excesso de competição profissional ou insuficiência de alternativas para clientes. | Monitoramento segmentado da liquidez e ajuste das estratégias de aquisição. |
| RSK-MKT-06 | Baixa recorrência de utilização | Redução da participação e dificuldade de consolidação do marketplace. | Desenvolver valor continuado por meio de portfólio, reputação, descoberta e acesso a novas oportunidades. |
| RSK-MKT-07 | Desintermediação | Perda de visibilidade sobre resultados e redução das razões para retornar ao produto. | Fortalecer os benefícios recorrentes da plataforma sem impor restrições artificiais às relações externas. |
| RSK-MKT-08 | Baixo diferencial competitivo | Dificuldade de aquisição e retenção diante de alternativas existentes. | Investigar concorrentes e validar diferenciais percebidos pelos usuários. |
| RSK-MKT-09 | Generalização indevida dos resultados do piloto | Expansão fundamentada em evidências pouco representativas. | Documentar características do mercado piloto e validar novos segmentos antes de expandir. |
| RSK-MKT-10 | Expansão prematura | Fragmentação da oferta, da demanda e da capacidade operacional. | Condicionar a expansão a evidências de liquidez, utilidade e capacidade de operação. |

## 13.2 Riscos financeiros e de sustentabilidade

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-FIN-01 | Ausência de modelo de negócio sustentável | Incapacidade de financiar a continuidade do produto. | Investigar alternativas de monetização e disposição a pagar sem comprometer a proposta de valor. |
| RSK-FIN-02 | Custo elevado de aquisição | Crescimento economicamente inviável. | Priorizar estratégias locais, parcerias, aquisição orgânica e mensuração de eficiência. |
| RSK-FIN-03 | Custo operacional desproporcional | Necessidade de recursos superiores à capacidade da equipe. | Monitorar custos de infraestrutura, suporte, onboarding e moderação. |
| RSK-FIN-04 | Conflito entre monetização e equidade | Deterioração da relevância, confiança e distribuição de oportunidades. | Preservar a independência entre capacidade financeira e classificação orgânica dos profissionais. |
| RSK-FIN-05 | Insuficiência de recursos financeiros | Interrupção do desenvolvimento ou da operação. | Planejamento orçamentário, limitação de custos fixos e priorização de investimentos conforme evidências. |

A ausência de monetização no MVP não deverá ser interpretada como ausência de necessidade de investigação da sustentabilidade econômica do produto.

## 13.3 Riscos de experiência e inclusão

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-UX-01 | Baixa familiaridade digital | Exclusão de profissionais e clientes com maior necessidade da solução. | Linguagem simples, jornadas reduzidas, onboarding assistido e testes com públicos representativos. |
| RSK-UX-02 | Fricção excessiva nas jornadas | Abandono antes da geração de valor. | Divulgação progressiva de informações e redução de etapas desnecessárias. |
| RSK-UX-03 | Incompatibilidade com dispositivos ou conexões limitadas | Restrição de acesso ao produto. | Design mobile-first, otimização de recursos e testes em condições representativas. |
| RSK-UX-04 | Barreiras de acessibilidade | Impedimento de participação de pessoas com deficiência. | Aplicação dos requisitos de acessibilidade e testes manuais e automatizados. |
| RSK-UX-05 | Informações profissionais insuficientes ou desatualizadas | Decisões inadequadas e conexões pouco relevantes. | Orientação de preenchimento, atualização simplificada e mecanismos de controle de disponibilidade. |
| RSK-UX-06 | Interpretação incorreta dos sinais de confiança | Expectativas indevidas de qualidade ou segurança. | Explicações claras sobre o significado e as limitações de verificações, avaliações e indicadores. |

## 13.4 Riscos de confiança, segurança e integridade

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-TRU-01 | Perfis falsos ou representação indevida | Fraude, perda de confiança e danos aos participantes. | Confirmações proporcionais ao risco, mecanismos de denúncia e análise de comportamento suspeito. |
| RSK-TRU-02 | Fraudes comerciais ou financeiras | Prejuízos econômicos aos usuários. | Orientações preventivas, comunicação dos limites da plataforma e tratamento de denúncias. |
| RSK-TRU-03 | Assédio, ameaças ou perseguição | Danos emocionais, sociais ou físicos. | Compartilhamento contextual de informações, bloqueio, denúncia e procedimentos de resposta. |
| RSK-TRU-04 | Riscos em interações presenciais | Danos à integridade física ou patrimonial dos participantes. | Orientações de segurança, minimização de exposição e mecanismos de comunicação de ocorrências. |
| RSK-TRU-05 | Manipulação de avaliações | Distorção da reputação e prejuízo à tomada de decisão. | Regras de elegibilidade, limites de participação, análise de padrões e moderação. |
| RSK-TRU-06 | Avaliações retaliatórias ou coercitivas | Inibição de avaliações legítimas e danos reputacionais. | Políticas de publicação, contestação, proteção dos participantes e mecanismos de denúncia. |
| RSK-TRU-07 | Spam e manifestações de interesse abusivas | Sobrecarga de clientes e profissionais. | Limites proporcionais, rate limiting e mecanismos de restrição de comportamentos abusivos. |
| RSK-TRU-08 | Publicação de serviços ou conteúdo irregular | Exposição a riscos legais e operacionais. | Catálogo controlado, políticas de conteúdo, denúncia e moderação. |
| RSK-TRU-09 | Uso indevido de conteúdo de portfólio | Violação de direitos, fraude de identidade profissional ou exposição de terceiros. | Políticas de publicação, orientação e mecanismos de denúncia e remoção. |
| RSK-TRU-10 | Discriminação ou tratamento abusivo entre participantes | Exclusão, danos pessoais e comprometimento da proposta de valor. | Políticas de conduta, mecanismos de denúncia e decisões proporcionais de moderação. |

A plataforma deverá comunicar os limites de suas capacidades de verificação, prevenção e intervenção.

A existência de mecanismos de confiança não deverá ser apresentada como garantia de legitimidade, qualidade profissional ou segurança das interações.

## 13.5 Riscos de equidade e integridade da descoberta

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-EQU-01 | Concentração excessiva de exposição | Redução das oportunidades para profissionais elegíveis. | Monitoramento de exposição em contextos comparáveis e diversificação proporcional. |
| RSK-EQU-02 | Penalização de profissionais sem histórico | Barreiras de entrada e perpetuação de vantagens acumuladas. | Tratamento adequado da ausência de histórico e exploração controlada de perfis elegíveis. |
| RSK-EQU-03 | Viés direto ou indireto nos mecanismos de classificação | Distribuição injustificada de oportunidades. | Avaliação dos sinais utilizados, auditoria, revisão humana e análise de efeitos adversos. |
| RSK-EQU-04 | Manipulação de mecanismos de classificação | Distorção da relevância e distribuição de exposição. | Validação dos sinais, monitoramento de padrões anômalos e controles contra abuso. |
| RSK-EQU-05 | Degradação de relevância por mecanismos de equidade | Redução da utilidade da descoberta para os clientes. | Guardrails de relevância, avaliação offline e experimentação controlada. |
| RSK-EQU-06 | Utilização inadequada de métricas de equidade | Otimização artificial de indicadores sem benefícios reais. | Análise contextual de exposição, elegibilidade, conversão e utilidade. |
| RSK-EQU-07 | Classificação incompreensível para os usuários | Perda de confiança e percepção de arbitrariedade. | Explicações proporcionais, documentação de políticas e canais de contestação. |
| RSK-EQU-08 | Ciclos de retroalimentação reputacional | Concentração progressiva de oportunidades nos mesmos profissionais. | Limitação da influência de popularidade e histórico acumulado, acompanhada de monitoramento. |

A ausência de utilização direta de atributos pessoais sensíveis não deverá ser considerada evidência suficiente de ausência de discriminação.

Os mecanismos de equidade deverão preservar a relevância, a autonomia dos usuários e a integridade das informações utilizadas.

## 13.6 Riscos de privacidade e proteção de dados

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-PRV-01 | Exposição indevida de informações de contato | Assédio, fraude e perda de controle sobre dados pessoais. | Autorização contextual, controle de acesso e minimização da exposição. |
| RSK-PRV-02 | Exposição ou inferência de localização precisa | Riscos à integridade física e à privacidade. | Localização aproximada, redução da precisão pública e não persistência desnecessária. |
| RSK-PRV-03 | Coleta excessiva de dados pessoais | Ampliação de riscos jurídicos e de segurança. | Inventário de dados, avaliação de necessidade e privacidade por padrão. |
| RSK-PRV-04 | Acesso administrativo indevido | Divulgação, alteração ou utilização irregular de informações. | Menor privilégio, controles de autenticação e auditoria. |
| RSK-PRV-05 | Reidentificação de dados analíticos | Exposição de participantes por combinação de informações. | Agregação, pseudonimização, anonimização adequada e restrição de segmentações. |
| RSK-PRV-06 | Retenção excessiva de informações | Aumento da superfície de exposição e descumprimento de obrigações. | Políticas justificadas de retenção, eliminação e revisão periódica. |
| RSK-PRV-07 | Tratamento inadequado por fornecedores | Exposição de dados fora do controle direto da plataforma. | Avaliação de fornecedores, condições contratuais e limitação dos compartilhamentos. |
| RSK-PRV-08 | Exposição de dados por ferramentas de desenvolvimento ou IA | Compartilhamento não autorizado de informações reais. | Dados sintéticos, sanitização de registros e restrições ao envio de informações externas. |
| RSK-PRV-09 | Descumprimento dos direitos dos titulares | Prejuízos aos usuários e riscos regulatórios. | Procedimentos de atendimento, registro de solicitações e controle dos prazos aplicáveis. |

As estratégias de mitigação deverão observar as diretrizes de governança e proteção de dados estabelecidas na seção 11.

## 13.7 Riscos operacionais

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-OPE-01 | Capacidade insuficiente de moderação | Acúmulo de denúncias e exposição prolongada a comportamentos abusivos. | Priorização por risco, dimensionamento da operação e limitação da expansão. |
| RSK-OPE-02 | Decisões inconsistentes de moderação | Tratamento arbitrário e perda de confiança. | Políticas explícitas, registro de justificativas e mecanismos de revisão. |
| RSK-OPE-03 | Dependência excessiva de processos manuais | Dificuldade de expansão e aumento de custos. | Identificação de gargalos e automação incremental de atividades apropriadas. |
| RSK-OPE-04 | Concentração de responsabilidades | Interrupção de atividades devido à indisponibilidade de integrantes da equipe. | Documentação, distribuição de responsabilidades e compartilhamento de conhecimento. |
| RSK-OPE-05 | Onboarding assistido não sustentável | Incapacidade de reproduzir os resultados do piloto em maior escala. | Medição do esforço operacional e simplificação progressiva das jornadas. |
| RSK-OPE-06 | Resposta inadequada a incidentes | Agravamento de danos e perda de confiança. | Procedimentos de resposta, responsabilidades definidas e critérios de escalonamento. |
| RSK-OPE-07 | Ausência de processos de suporte | Abandono, conflitos não resolvidos e redução da satisfação. | Canais acessíveis de suporte e tratamento proporcional das solicitações. |

A capacidade operacional deverá ser considerada uma restrição explícita para a aquisição de usuários e expansão do produto.

## 13.8 Riscos tecnológicos e de engenharia

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-TEC-01 | Complexidade arquitetural prematura | Aumento de custos, prazo e dificuldade de manutenção. | Arquitetura proporcional ao estágio do produto, modularidade e evolução incremental. |
| RSK-TEC-02 | Crescimento descontrolado do escopo | Atraso na entrega e comprometimento da validação. | Priorização por valor, controle de mudanças e preservação do escopo do MVP. |
| RSK-TEC-03 | Vulnerabilidades de autenticação e autorização | Comprometimento de contas e exposição de informações. | Desenvolvimento seguro, testes de autorização e revisão dos controles críticos. |
| RSK-TEC-04 | Inconsistência de estados e operações concorrentes | Conexões inválidas, duplicidades e violações de regras de negócio. | Invariantes explícitas, controle de concorrência, idempotência e testes de integração. |
| RSK-TEC-05 | Indisponibilidade ou perda de dados | Interrupção de jornadas e perda de confiança. | Monitoramento, backups, procedimentos de recuperação e testes de restauração. |
| RSK-TEC-06 | Degradação de desempenho | Abandono das jornadas e aumento de custos operacionais. | Metas proporcionais, testes representativos e otimização fundamentada em evidências. |
| RSK-TEC-07 | Dependência excessiva de fornecedores | Dificuldade de migração ou continuidade operacional. | Avaliação de dependências, portabilidade proporcional e estratégias de substituição quando justificadas. |
| RSK-TEC-08 | Déficit de qualidade decorrente de geração automatizada de código | Introdução de vulnerabilidades, inconsistências e dificuldades de manutenção. | Revisão humana, testes automatizados e validação de decisões produzidas com assistência de IA. |
| RSK-TEC-09 | Ausência de rastreabilidade técnica | Dificuldade de investigar falhas e compreender decisões. | Documentação de arquitetura, registros de decisões e observabilidade. |
| RSK-TEC-10 | Acúmulo de dívida técnica crítica | Redução da capacidade de evolução e aumento de riscos operacionais. | Identificação, classificação e tratamento proporcional de dívida técnica. |

A mitigação de riscos tecnológicos não deverá justificar investimentos prematuros em infraestrutura, arquitetura ou automação sem necessidade demonstrada.

## 13.9 Riscos jurídicos e regulatórios

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-LEG-01 | Descumprimento de obrigações de proteção de dados | Sanções, danos aos titulares e perda de confiança. | Governança de dados, avaliação jurídica e implementação dos controles aplicáveis. |
| RSK-LEG-02 | Interpretação inadequada das responsabilidades da plataforma | Conflitos e expectativas incompatíveis com a atuação do RAIRE. | Definição clara do papel da plataforma e avaliação das responsabilidades legais aplicáveis. |
| RSK-LEG-03 | Divulgação de serviços sujeitos a requisitos profissionais específicos | Facilitação de atividades irregulares ou exposição a riscos legais. | Avaliação das categorias disponibilizadas e adoção de restrições proporcionais. |
| RSK-LEG-04 | Violação de direitos de imagem ou propriedade intelectual | Reclamações, remoção de conteúdo e responsabilização. | Políticas de publicação, denúncia, análise e remoção de conteúdo irregular. |
| RSK-LEG-05 | Conflitos decorrentes de avaliações públicas | Danos reputacionais e disputas entre participantes. | Políticas de avaliação, contestação e moderação fundamentada. |
| RSK-LEG-06 | Participação inadequada de menores de idade | Exposição de públicos vulneráveis e riscos regulatórios. | Definição de critérios etários e regras de participação antes do piloto. |

Os termos de uso e comunicações da plataforma deverão refletir suas capacidades efetivas e as responsabilidades estabelecidas pela legislação aplicável.

A avaliação jurídica deverá considerar as características das categorias de serviço, do modelo de negócio e das operações realizadas.

## 13.10 Riscos de pesquisa e validação

| ID | Risco | Impacto potencial | Estratégia de mitigação |
| --- | --- | --- | --- |
| RSK-VAL-01 | Viés de confirmação | Manutenção de hipóteses incompatíveis com as evidências. | Critérios prévios de avaliação e documentação de resultados negativos. |
| RSK-VAL-02 | Amostras pouco representativas | Conclusões inadequadas sobre os públicos atendidos. | Diversificação dos participantes e documentação das limitações amostrais. |
| RSK-VAL-03 | Confusão entre engajamento e valor | Otimização de cadastros, cliques e impressões sem benefícios efetivos. | Priorização das métricas de utilidade e resultados estabelecidas na seção 12. |
| RSK-VAL-04 | Instrumentação incompleta ou incorreta | Decisões fundamentadas em indicadores inconsistentes. | Definições operacionais, validação de eventos e testes de instrumentação. |
| RSK-VAL-05 | Experimentos sem condições metodológicas adequadas | Interpretação incorreta de efeitos observados. | Avaliação de amostragem, vieses, interferências e adequação do método. |
| RSK-VAL-06 | Dependência de resultados autodeclarados | Visibilidade limitada sobre a qualidade e os desfechos das conexões. | Combinação de eventos observáveis, feedback bilateral e investigação qualitativa. |
| RSK-VAL-07 | Resultados dependentes de intervenção manual | Superestimação da capacidade autônoma do produto. | Registro do esforço de assistência e avaliação de jornadas sem intervenção. |

Resultados inconclusivos deverão ser reconhecidos como tais, evitando apresentar ausência de evidência como confirmação de uma hipótese.

## 13.11 Avaliação e classificação de riscos

Cada risco deverá ser avaliado considerando:

- Probabilidade de ocorrência.
- Impacto potencial.
- Exposição dos usuários e da organização.
- Capacidade de detecção.
- Efetividade dos controles existentes.
- Esforço necessário para mitigação.
- Risco residual após aplicação dos controles.

A classificação inicial poderá utilizar os seguintes níveis:

| Nível | Diretriz de tratamento |
| --- | --- |
| Crítico | Exige tratamento prioritário e poderá impedir o lançamento ou determinar suspensão de operações afetadas. |
| Alto | Exige estratégia de mitigação, acompanhamento e responsabilidade definida. |
| Moderado | Exige acompanhamento e tratamento proporcional. |
| Baixo | Poderá ser aceito mediante justificativa e revisão periódica. |

A classificação deverá considerar o contexto e o estágio de maturidade do produto.

Riscos relacionados à integridade física, exposição de dados pessoais, fraude grave ou obrigações legais não deverão ser aceitos exclusivamente em razão de sua baixa probabilidade estimada.

## 13.12 Estratégias de tratamento

O tratamento dos riscos poderá envolver:

- **Evitar:** remover a atividade ou condição que origina o risco.
- **Mitigar:** reduzir sua probabilidade ou impacto.
- **Transferir ou compartilhar:** distribuir determinadas responsabilidades por mecanismos contratuais ou operacionais adequados.
- **Aceitar:** reconhecer um risco residual mediante justificativa, responsabilidade e acompanhamento.

A transferência de atividades a terceiros não deverá ser interpretada como eliminação automática das responsabilidades da plataforma.

A aceitação de riscos deverá observar as obrigações legais, os princípios do produto e os limites de exposição dos participantes.

## 13.13 Registro e acompanhamento

Os riscos relevantes deverão ser mantidos em um registro atualizado.

Cada registro deverá contemplar, quando aplicável:

- Identificador.
- Descrição do risco.
- Causas e condições de ocorrência.
- Impactos potenciais.
- Probabilidade estimada.
- Classificação de criticidade.
- Controles preventivos existentes.
- Estratégia de tratamento.
- Indicadores de alerta.
- Procedimentos de resposta.
- Responsável pelo acompanhamento.
- Risco residual.
- Estado atual.
- Data de revisão.

O registro deverá ser atualizado diante de mudanças relevantes no escopo, arquitetura, modelo de negócio, categorias atendidas ou condições operacionais.

Durante o piloto, a revisão deverá acompanhar os marcos periódicos de aprendizado e validação.

## 13.14 Condições impeditivas para o piloto

A disponibilização do MVP para usuários reais deverá considerar a existência de impedimentos críticos, incluindo:

- Falhas conhecidas nos controles essenciais de autenticação e autorização.
- Exposição indevida de informações pessoais.
- Ausência de mecanismos mínimos de comunicação e tratamento de abusos.
- Ausência de procedimentos de resposta a incidentes.
- Incapacidade de preservar a integridade das conexões e autorizações.
- Ausência das definições legais e de privacidade necessárias.
- Riscos críticos sem estratégia de tratamento adequada.
- Capacidade operacional incompatível com o volume de participantes admitidos.

O piloto deverá possuir critérios de interrupção ou restrição de funcionalidades diante da materialização de riscos relevantes.

## 13.15 Evolução da gestão de riscos

A gestão de riscos deverá acompanhar a evolução do RAIRE, sem estabelecer processos desproporcionais à maturidade do projeto.

A identificação de novos riscos deverá resultar em avaliação de impacto sobre:

- Proposta de valor.
- Escopo do MVP.
- Requisitos funcionais e não funcionais.
- Regras de negócio.
- Arquitetura.
- Políticas operacionais.
- Estratégias de validação.
- Sustentabilidade do produto.

O objetivo da gestão de riscos não será eliminar toda incerteza, mas permitir decisões fundamentadas, limitar exposições injustificadas e preservar a capacidade de evolução responsável do produto.

---

# 14. Roadmap de entrega e evolução do produto

O roadmap do RAIRE deverá orientar a evolução incremental do produto, considerando a redução de incertezas, a entrega de capacidades funcionais, a validação das hipóteses e o estabelecimento das condições necessárias para uma operação responsável e sustentável.

A progressão entre fases deverá considerar evidências, dependências, riscos, disponibilidade de recursos e critérios de aceitação, não exclusivamente o cumprimento de cronogramas ou a conclusão de atividades técnicas.

## 14.1 Princípios de planejamento

O roadmap deverá observar os seguintes princípios:

- **Entrega incremental:** disponibilizar capacidades em incrementos verificáveis.
- **Orientação a valor:** priorizar funcionalidades que contribuam diretamente para os objetivos do produto.
- **Orientação a evidências:** revisar prioridades conforme os resultados das atividades de pesquisa e validação.
- **Redução de riscos:** investigar antecipadamente as incertezas capazes de comprometer a continuidade do projeto.
- **Proporcionalidade técnica:** evitar complexidade arquitetural e operacional sem necessidade demonstrada.
- **Qualidade contínua:** incorporar segurança, privacidade, acessibilidade, testabilidade e observabilidade desde as primeiras entregas.
- **Rastreabilidade:** relacionar entregas aos objetivos, requisitos, hipóteses e decisões correspondentes.
- **Flexibilidade:** permitir revisão de escopo, sequência e abordagem conforme novas evidências.
- **Sustentabilidade:** considerar a capacidade financeira, técnica e operacional da equipe.

As fases não deverão ser interpretadas como etapas necessariamente isoladas ou estritamente sequenciais.

## 14.2 F0 — Discovery e definição do produto

### Objetivo

Reduzir as principais incertezas relacionadas ao problema, público-alvo, proposta de valor, mercado inicial e viabilidade do produto.

### Entregas esperadas

- Consolidação inicial do PRD.
- Identificação e priorização das hipóteses do produto.
- Pesquisa com clientes e profissionais autônomos.
- Investigação de alternativas e concorrentes.
- Refinamento das personas e jornadas.
- Definição preliminar da região e categorias do piloto.
- Delimitação do escopo do protótipo e do MVP.
- Identificação dos requisitos prioritários.
- Definição preliminar das regras de negócio e conceitos do domínio.
- Identificação dos principais riscos e dependências.
- Planejamento inicial da estratégia de validação.
- Avaliação preliminar de alternativas de sustentabilidade financeira.
- Identificação das restrições técnicas, financeiras e operacionais.

### Critérios de saída

- Problema e proposta de valor suficientemente delimitados.
- Hipóteses prioritárias identificadas e acompanhadas de estratégias de investigação.
- Jornadas essenciais definidas.
- Escopo inicial acordado pela equipe.
- Riscos críticos identificados.
- Incertezas impeditivas tratadas ou acompanhadas de planos de investigação.

A conclusão desta fase não pressupõe a validação definitiva do mercado ou do modelo de negócio.

## 14.3 F1 — Fundação técnica e operacional

### Objetivo

Estabelecer as condições mínimas necessárias para o desenvolvimento consistente, seguro e sustentável das capacidades do produto.

### Entregas esperadas

- Definição da arquitetura inicial.
- Identificação das fronteiras funcionais e responsabilidades do sistema.
- Registro das decisões arquiteturais relevantes.
- Padronização dos ambientes de desenvolvimento.
- Configuração dos repositórios e processos de colaboração.
- Integração contínua e mecanismos de entrega.
- Estratégia inicial de testes automatizados.
- Harness de desenvolvimento e validação com assistência de IA.
- Fundação de identidade, autenticação, sessões e autorização.
- Suporte à coexistência dos papéis de cliente e profissional.
- Design system inicial e padrões de acessibilidade.
- Configuração segura de segredos e variáveis de ambiente.
- Fundação de observabilidade e auditoria.
- Definição inicial das políticas de tratamento e acesso a dados.
- Capacidades administrativas necessárias às primeiras entregas.

### Critérios de saída

- Ambiente de desenvolvimento reproduzível.
- Processos essenciais de build, teste e entrega operacionais.
- Arquitetura inicial documentada.
- Controles fundamentais de autenticação e autorização testados.
- Padrões mínimos de qualidade definidos.
- Equipe capaz de desenvolver e integrar novas capacidades de maneira consistente.

A arquitetura e os mecanismos de engenharia deverão permanecer proporcionais às necessidades das entregas previstas.

## 14.4 F2 — Oferta e presença profissional

### Objetivo

Permitir que profissionais autônomos construam uma presença digital especializada, apresentem suas competências e disponibilizem informações necessárias à descoberta local.

### Entregas esperadas

- Cadastro e ativação do perfil profissional.
- Catálogo inicial de categorias e serviços.
- Associação de profissionais aos serviços oferecidos.
- Definição de região e área de atendimento.
- Declaração de disponibilidade.
- Publicação e atualização do perfil.
- Criação e gerenciamento de portfólio profissional.
- Publicação de imagens e descrições contextualizadas.
- Visualização pública das informações autorizadas.
- Controles de visibilidade e suspensão voluntária.
- Administração básica de categorias e conteúdo.
- Instrumentação das jornadas de onboarding.
- Mecanismos mínimos de denúncia e tratamento de conteúdo inadequado antes da exposição a usuários reais.

### Hipóteses investigadas

- Profissionais percebem valor em uma presença digital especializada.
- O processo de publicação é acessível ao público pretendido.
- Informações de serviços e portfólio representam adequadamente as competências profissionais.
- A estrutura de categorias é compreensível e adequada ao mercado inicial.

### Critérios de saída

- Profissionais conseguem criar, publicar, atualizar e suspender seus perfis.
- Profissionais conseguem apresentar serviços e trabalhos realizados.
- As informações públicas respeitam as regras de visibilidade.
- As jornadas essenciais podem ser concluídas em dispositivos representativos.
- As informações necessárias à descoberta encontram-se disponíveis.

## 14.5 F3 — Descoberta local e conexão direta

### Objetivo

Permitir que clientes identifiquem profissionais adequados às suas necessidades e estabeleçam conexões mediante autorização das partes.

### Entregas esperadas

- Busca por serviço e região.
- Identificação de profissionais elegíveis.
- Filtros essenciais de descoberta.
- Ranking inicial baseado em relevância.
- Tratamento proporcional de profissionais sem histórico.
- Controles básicos contra concentração injustificada de exposição.
- Explicações dos critérios de descoberta.
- Consulta de perfis e portfólios.
- Tratamento de buscas sem resultados.
- Solicitações diretas de conexão.
- Aceitação, recusa, cancelamento e expiração de solicitações.
- Compartilhamento contextual e autorizado de contato.
- Notificações essenciais.
- Mecanismos de bloqueio e prevenção de abuso.
- Instrumentação dos eventos de descoberta e conexão.
- Avaliação inicial da relevância dos resultados.

### Hipóteses investigadas

- A descoberta especializada facilita a identificação de profissionais adequados.
- Informações de perfil e portfólio contribuem para decisões informadas.
- A proximidade geográfica possui relevância para os serviços selecionados.
- O mecanismo de conexão consentida é compreensível e útil.

### Critérios de saída

- Clientes conseguem encontrar profissionais compatíveis.
- Resultados respeitam os critérios de elegibilidade.
- A classificação possui comportamento documentado e verificável.
- Profissionais elegíveis sem histórico não são excluídos exclusivamente por essa condição.
- Clientes conseguem solicitar conexões.
- Profissionais conseguem avaliar e responder às solicitações.
- Informações privadas não são disponibilizadas antes das autorizações necessárias.
- Eventos essenciais da jornada são registrados.

## 14.6 F4 — Demanda e descoberta de oportunidades

### Objetivo

Permitir que clientes publiquem necessidades de serviço e que profissionais assumam uma participação ativa na descoberta de oportunidades de trabalho.

### Entregas esperadas

- Criação e publicação de demandas de serviço.
- Classificação por categoria e contexto geográfico.
- Gerenciamento, atualização e encerramento de demandas.
- Descoberta de oportunidades por profissionais.
- Consulta de detalhes e informações autorizadas.
- Manifestação e retirada de interesse.
- Comparação dos profissionais interessados.
- Aceitação ou recusa de manifestações de interesse.
- Estabelecimento de conexões consentidas.
- Notificações relacionadas às demandas.
- Controles contra spam, fraude e exposição indevida.
- Instrumentação dos eventos relacionados à demanda e às oportunidades.

### Hipóteses investigadas

- Clientes percebem valor na publicação de necessidades.
- Profissionais possuem interesse em buscar oportunidades ativamente.
- A descoberta de demandas complementa a descoberta direta de profissionais.
- O mecanismo permite reduzir barreiras de aquisição de clientes.

### Critérios de saída

- Clientes conseguem publicar e gerenciar demandas válidas.
- Profissionais conseguem identificar oportunidades compatíveis.
- Manifestações de interesse respeitam as regras de elegibilidade.
- Clientes conseguem avaliar interessados e estabelecer conexões.
- Os estados das demandas, interesses e conexões permanecem consistentes.
- A privacidade e a autonomia das partes são preservadas.
- A jornada completa pode ser avaliada mediante eventos e feedback.

## 14.7 F5 — Confiança, reputação e avaliação de resultados

### Objetivo

Estabelecer mecanismos de confiança e aprendizado sobre as interações realizadas, sem atribuir à plataforma capacidades de verificação ou garantia que ela não possua.

### Entregas esperadas

- Registro declaratório dos desfechos das conexões.
- Coleta de feedback sobre utilidade.
- Definição e aplicação das regras de elegibilidade para avaliações.
- Avaliações bilaterais contextualizadas.
- Apresentação de informações de reputação.
- Distinção entre reputação como cliente e como profissional.
- Comunicação clara dos sinais de confirmação e verificação.
- Mecanismos de contestação de avaliações.
- Evolução dos controles de denúncia e moderação.
- Registro das decisões administrativas.
- Instrumentação das métricas de confiança e reputação.
- Avaliação dos riscos de manipulação, retaliação e discriminação.

Mecanismos adicionais de verificação de identidade deverão depender de avaliação de necessidade, riscos, custos operacionais e requisitos legais.

### Hipóteses investigadas

- A reputação contextual contribui para decisões informadas.
- Avaliações bilaterais geram valor para ambos os participantes.
- Os usuários compreendem o significado e os limites dos sinais de confiança.
- O feedback sobre desfechos permite investigar a utilidade das conexões.

### Critérios de saída

- Avaliações somente podem ser realizadas em interações elegíveis.
- Regras de reputação são documentadas e verificáveis.
- Existem mecanismos de comunicação e tratamento de abusos.
- Os sinais de confiança não induzem interpretações indevidas.
- Resultados declarados são distinguíveis de acontecimentos observados pela plataforma.
- A equipe possui capacidade mínima de acompanhar interações problemáticas.

Os controles fundamentais de segurança, privacidade, denúncia e proteção dos usuários não deverão depender da conclusão desta fase.

## 14.8 F6 — Preparação e execução do piloto

### Objetivo

Avaliar a capacidade do RAIRE de gerar valor para clientes e profissionais em condições reais de utilização, dentro de um mercado inicial delimitado.

### Preparação

- Confirmação da região e categorias selecionadas.
- Definição dos critérios de participação.
- Planejamento da aquisição inicial de oferta e demanda.
- Estabelecimento de parcerias locais, quando pertinentes.
- Recrutamento e onboarding de profissionais.
- Validação das jornadas essenciais.
- Revisão dos riscos críticos.
- Avaliação das condições de segurança e privacidade.
- Definição das responsabilidades operacionais.
- Preparação dos procedimentos de suporte e moderação.
- Definição operacional das métricas prioritárias.
- Estabelecimento dos critérios de continuidade, ajuste e interrupção.
- Verificação da capacidade de instrumentação.

### Execução

- Disponibilização controlada do produto.
- Aquisição de clientes com necessidades reais.
- Acompanhamento da densidade de oferta e demanda.
- Monitoramento da descoberta e das conexões.
- Coleta de feedback bilateral.
- Avaliação da relevância e utilidade dos resultados.
- Monitoramento da distribuição de exposição.
- Identificação de barreiras de acessibilidade e usabilidade.
- Acompanhamento de incidentes e denúncias.
- Avaliação dos custos e esforços operacionais.
- Revisão periódica dos resultados e implementação de ajustes prioritários.

### Critérios de saída

- Evidências registradas sobre as hipóteses prioritárias.
- Métricas iniciais de utilização, liquidez e utilidade estabelecidas.
- Problemas e limitações identificados.
- Capacidade operacional avaliada.
- Resultados de confiança, segurança e equidade analisados.
- Relatório de aprendizado produzido.
- Decisão fundamentada sobre continuidade, ajuste, redução, expansão ou suspensão do produto.

A duração do piloto deverá ser definida no plano de validação, considerando o volume de participação, a natureza dos serviços e a suficiência das evidências.

## 14.9 F7 — Otimização, sustentabilidade e evolução

### Objetivo

Evoluir o produto com base nas evidências obtidas, melhorando sua capacidade de gerar conexões úteis, ampliar oportunidades e operar de maneira responsável e sustentável.

### Entregas potenciais

- Refinamento dos mecanismos de relevância.
- Otimização da descoberta geográfica.
- Evolução dos mecanismos de fairness.
- Simulação e avaliação offline de políticas de classificação.
- Experimentação controlada, quando metodologicamente viável.
- Evolução da reputação contextual.
- Aprimoramento dos mecanismos de prevenção de fraude.
- Redução de fricções nas jornadas.
- Evolução dos mecanismos de suporte e moderação.
- Automação proporcional de atividades operacionais.
- Desenvolvimento de capacidades analíticas adicionais.
- Investigação de modelos de monetização.
- Avaliação de novas categorias e regiões.
- Otimização de desempenho e custos.
- Evolução arquitetural conforme necessidades demonstradas.

As funcionalidades desta fase não constituem compromissos antecipados de implementação.

Sua priorização deverá depender dos resultados observados, dos riscos identificados, da capacidade da equipe e das oportunidades de geração de valor.

A expansão deverá ocorrer mediante evidências suficientes de liquidez, utilidade, segurança e sustentabilidade operacional.

## 14.10 Responsabilidades transversais

As seguintes atividades deverão acompanhar continuamente a evolução do produto:

| Dimensão | Responsabilidade |
| --- | --- |
| Pesquisa | Investigar hipóteses e necessidades dos participantes. |
| Produto | Revisar prioridades, escopo e proposta de valor. |
| Engenharia | Preservar qualidade, consistência e capacidade de evolução. |
| Segurança | Identificar ameaças e implementar controles proporcionais. |
| Privacidade | Avaliar tratamentos e preservar os direitos dos titulares. |
| Acessibilidade | Identificar e reduzir barreiras de participação. |
| Equidade | Preservar relevância e evitar distribuição injustificada de oportunidades. |
| Observabilidade | Permitir diagnóstico e acompanhamento da operação. |
| Validação | Coletar evidências e avaliar resultados. |
| Documentação | Preservar requisitos, regras, decisões e aprendizados. |
| Gestão de riscos | Identificar, acompanhar e tratar exposições relevantes. |

Nenhuma dessas dimensões deverá ser considerada exclusivamente responsabilidade de uma fase posterior.

## 14.11 Marcos de entrega

O roadmap deverá distinguir os seguintes marcos:

| Marco | Resultado esperado |
| --- | --- |
| M1 — Fundamentos definidos | Problema, proposta de valor, escopo e direção inicial estabelecidos. |
| M2 — Fundação operacional de desenvolvimento | Equipe capaz de implementar, testar e entregar incrementos. |
| M3 — Presença profissional funcional | Profissionais conseguem apresentar seus serviços e portfólios. |
| M4 — Primeira jornada vertical funcional | Cliente consegue descobrir um profissional e estabelecer uma conexão consentida. |
| M5 — Marketplace bilateral funcional | Clientes publicam demandas e profissionais descobrem oportunidades. |
| M6 — Confiança e reputação integradas | Interações elegíveis produzem feedback e informações reputacionais. |
| M7 — MVP apto ao piloto | Condições funcionais, operacionais, legais e de segurança atendidas. |
| M8 — Piloto avaliado | Evidências e decisões sobre a evolução do produto documentadas. |

Os marcos representam resultados verificáveis e não necessariamente versões públicas do produto.

## 14.12 Protótipo, MVP e produto em evolução

### Protótipo demonstrável

O protótipo deverá comunicar a proposta de valor e demonstrar as jornadas representativas do produto em condições controladas.

Poderá utilizar:

- Dados sintéticos.
- Categorias e regiões limitadas.
- Interfaces simplificadas.
- Regras iniciais de descoberta.
- Simulações de interações.
- Integrações substituídas por mecanismos demonstrativos.

Limitações e funcionalidades simuladas deverão ser explicitadas.

Um protótipo demonstrável não deverá ser interpretado automaticamente como produto apto ao tratamento de dados reais ou à operação pública.

### MVP operacional

O MVP deverá implementar as capacidades mínimas necessárias para investigar a proposta de valor em condições reais de utilização.

Deverá atender:

- Requisitos funcionais prioritários.
- Controles essenciais de segurança e privacidade.
- Critérios aplicáveis de acessibilidade.
- Regras de negócio fundamentais.
- Instrumentação necessária à validação.
- Condições mínimas de suporte e moderação.
- Obrigações legais e regulatórias aplicáveis.
- Critérios de entrada no piloto.

### Produto em evolução

A evolução posterior deverá incorporar melhorias e novas capacidades conforme evidências de valor, necessidades dos usuários e sustentabilidade da operação.

## 14.13 Planejamento de implementação

Cada incremento de entrega deverá possuir, quando aplicável:

- Objetivo.
- Capacidades contempladas.
- Requisitos relacionados.
- Hipóteses investigadas.
- Dependências.
- Critérios de aceitação.
- Riscos relevantes.
- Estratégia de validação.
- Responsáveis.
- Estimativa de esforço.
- Condições de conclusão.

O detalhamento em épicos, histórias, tarefas e atividades técnicas deverá ser realizado no backlog de desenvolvimento.

O roadmap não deverá substituir o backlog, a especificação de requisitos, o plano de validação ou a documentação arquitetural.

## 14.14 Revisão e governança do roadmap

O roadmap deverá ser revisado regularmente e sempre que ocorrerem mudanças relevantes nas condições do projeto.

A revisão deverá considerar:

- Evidências produzidas pelas atividades de pesquisa.
- Resultados das entregas realizadas.
- Alterações nas hipóteses do produto.
- Riscos identificados ou materializados.
- Capacidade da equipe.
- Dependências técnicas ou operacionais.
- Custos e recursos disponíveis.
- Alterações regulatórias relevantes.
- Oportunidades de simplificação.
- Feedback dos participantes.

Alterações relevantes deverão possuir justificativa registrada.

Prazos definitivos deverão ser estabelecidos após a avaliação de escopo, dependências, capacidade e esforço.

A conclusão de uma fase não deverá impedir revisões posteriores de seus resultados ou decisões.

O roadmap deverá permanecer um instrumento de orientação e aprendizado, e não um compromisso inflexível com funcionalidades cuja necessidade ainda não foi demonstrada.

---

# 15. Critérios de aceite e maturidade do PRD

O PRD será considerado suficientemente maduro quando estabelecer uma compreensão compartilhada sobre o problema, a proposta de valor, os objetivos, o público-alvo, as capacidades esperadas, as restrições e as principais incertezas do RAIRE.

Sua aprovação deverá indicar que existe fundamentação suficiente para orientar a elaboração de especificações, decisões de engenharia e planejamento incremental.

A aprovação do PRD não pressupõe a validação definitiva das hipóteses de mercado, a conclusão da arquitetura, a implementação das funcionalidades ou a prontidão do produto para operação pública.

## 15.1 Clareza estratégica e fundamentação

- [ ] O contexto e o problema estão explicitamente definidos.
- [ ] Problemas de clientes e profissionais estão representados.
- [ ] Hipóteses estão distinguíveis de fatos sustentados por evidências.
- [ ] Visão, missão e proposta de valor são consistentes entre si.
- [ ] Os diferenciais pretendidos estão identificados.
- [ ] Os princípios fundamentais do produto estão estabelecidos.
- [ ] As principais restrições estratégicas estão documentadas.
- [ ] Os objetivos possuem relação identificável com os problemas apresentados.

## 15.2 Usuários e jornadas

- [ ] Os atores relevantes estão identificados.
- [ ] As personas estão documentadas como hipóteses ou perfis sustentados por pesquisa.
- [ ] As necessidades e os Jobs to Be Done estão identificados.
- [ ] As jornadas essenciais estão descritas.
- [ ] As jornadas de clientes e profissionais possuem representação adequada.
- [ ] Os mecanismos de descoberta direta e publicação de necessidades estão contemplados.
- [ ] Os pontos críticos de confiança, segurança e privacidade estão identificados.
- [ ] As principais barreiras de participação e acessibilidade estão consideradas.

## 15.3 Objetivos e escopo

- [ ] Os objetivos do MVP estão explicitamente definidos.
- [ ] As hipóteses prioritárias de validação estão identificadas.
- [ ] As capacidades contempladas pelo MVP estão delimitadas.
- [ ] Os não objetivos e exclusões de escopo estão documentados.
- [ ] As diferenças entre protótipo, MVP operacional e evolução futura estão estabelecidas.
- [ ] As funcionalidades essenciais estão distinguíveis das capacidades evolutivas.
- [ ] As restrições de recursos e capacidade conhecidas estão consideradas.
- [ ] As decisões de escopo possuem justificativa compatível com os objetivos do produto.

## 15.4 Requisitos funcionais e não funcionais

- [ ] Os requisitos possuem identificadores únicos.
- [ ] Os requisitos funcionais expressam capacidades ou comportamentos esperados.
- [ ] Os requisitos não funcionais expressam atributos de qualidade e restrições relevantes.
- [ ] Os requisitos estão relacionados aos objetivos e jornadas correspondentes, quando aplicável.
- [ ] Os requisitos possuem priorização inicial.
- [ ] Os requisitos são suficientemente claros para orientar especificações detalhadas.
- [ ] Os requisitos possuem condições verificáveis ou mecanismos previstos de verificação.
- [ ] Dependências e incompatibilidades relevantes estão identificadas.
- [ ] Decisões de implementação estão separadas das necessidades do produto, salvo restrições técnicas justificadas.
- [ ] Incertezas relevantes estão explicitamente registradas.

Os critérios detalhados de aceitação, fluxos alternativos, regras de negócio, invariantes e contratos deverão ser desenvolvidos nas especificações correspondentes.

## 15.5 Equidade, confiança e responsabilidade

- [ ] Os princípios de equidade estão estabelecidos.
- [ ] A relação entre relevância e distribuição de oportunidades está definida.
- [ ] Restrições aplicáveis aos mecanismos de classificação estão documentadas.
- [ ] Os princípios de confiança e reputação estão estabelecidos.
- [ ] Os limites das capacidades de verificação da plataforma estão explícitos.
- [ ] Os princípios de segurança, privacidade e proteção de dados estão documentados.
- [ ] Os requisitos de acessibilidade estão identificados.
- [ ] As questões dependentes de avaliação jurídica ou especializada estão registradas.
- [ ] Não existem promessas de segurança, conformidade ou verificação sem fundamentação adequada.

As avaliações especializadas necessárias deverão ser realizadas antes da implementação ou operação das capacidades afetadas, conforme a natureza e a gravidade dos riscos envolvidos.

## 15.6 Validação e métricas

- [ ] Existe uma estratégia inicial de validação do problema e da proposta de valor.
- [ ] As hipóteses prioritárias possuem métodos de investigação propostos.
- [ ] Existe uma métrica central candidata de geração de valor.
- [ ] Os indicadores essenciais de oferta, demanda, descoberta, conexão e liquidez estão identificados.
- [ ] Os indicadores de confiança, equidade, segurança e acessibilidade estão contemplados.
- [ ] Métricas observáveis estão diferenciadas de resultados autodeclarados.
- [ ] Metas preliminares estão diferenciadas de metas sustentadas por evidências.
- [ ] Existe uma estratégia para estabelecimento de baseline.
- [ ] Os critérios gerais de continuidade, revisão e interrupção do piloto estão identificados.
- [ ] As limitações conhecidas da estratégia de medição estão documentadas.

A ausência de dados empíricos suficientes não deverá resultar na criação artificial de metas quantitativas.

## 15.7 Riscos e questões abertas

- [ ] Os principais riscos estratégicos, operacionais, tecnológicos e jurídicos estão identificados.
- [ ] Os riscos prioritários possuem estratégias iniciais de tratamento.
- [ ] Os riscos críticos possuem responsáveis pelo acompanhamento.
- [ ] As dependências externas relevantes estão identificadas.
- [ ] As questões abertas possuem classificação de prioridade.
- [ ] As decisões pendentes possuem responsável ou estratégia de investigação.
- [ ] As questões impeditivas para a elaboração das especificações prioritárias estão resolvidas.
- [ ] As questões não impeditivas estão registradas para tratamento posterior.
- [ ] Não existem contradições críticas conhecidas entre as seções do documento.

A existência de questões abertas não impede automaticamente a aprovação do PRD, desde que suas consequências, dependências e estratégias de tratamento estejam identificadas.

## 15.8 Roadmap e planejamento

- [ ] Existe uma estratégia de entrega incremental.
- [ ] As capacidades prioritárias estão identificadas.
- [ ] Os principais marcos de evolução estão estabelecidos.
- [ ] As dependências relevantes entre entregas estão identificadas.
- [ ] Os critérios gerais de progressão entre fases estão documentados.
- [ ] As atividades de pesquisa e validação acompanham a evolução do produto.
- [ ] Os controles essenciais de qualidade, segurança e privacidade são considerados desde as primeiras entregas.
- [ ] O roadmap distingue preparação técnica, implementação, validação e operação.
- [ ] Capacidades evolutivas não constituem compromissos antecipados de implementação.

## 15.9 Qualidade documental e rastreabilidade

- [ ] A terminologia utilizada é suficientemente consistente.
- [ ] Conceitos relevantes possuem definição ou referência identificável.
- [ ] Identificadores de objetivos, requisitos, hipóteses e riscos consistem entre as seções.
- [ ] As referências a documentos complementares estão identificadas.
- [ ] Decisões relevantes possuem justificativa registrada.
- [ ] Hipóteses, propostas, restrições e decisões aprovadas são distinguíveis.
- [ ] O documento possui identificação de versão e estado de aprovação.
- [ ] Existe um responsável pela manutenção do documento.
- [ ] Existe um procedimento proporcional para revisão e atualização.
- [ ] A equipe possui acesso à versão vigente.

## 15.10 Estados de maturidade do PRD

O PRD poderá assumir os seguintes estados:

| Estado | Definição |
| --- | --- |
| Draft | Documento em elaboração, sujeito a alterações significativas. |
| In Review | Documento submetido à revisão de consistência, escopo e fundamentação. |
| Approved | Documento suficientemente maduro para orientar as atividades subsequentes. |
| Under Revision | Documento em revisão devido a novas evidências ou mudanças relevantes. |
| Superseded | Versão substituída por uma versão posterior. |

A aprovação deverá estabelecer uma baseline documental, sem tornar o produto ou suas especificações imutáveis.

## 15.11 Condições de aprovação

O PRD poderá ser aprovado quando:

1. Os critérios obrigatórios aplicáveis estiverem atendidos.
2. Não existirem questões impeditivas para as próximas atividades planejadas.
3. Os riscos críticos conhecidos possuírem tratamento ou plano de investigação.
4. As pendências não impeditivas estiverem documentadas.
5. A equipe possuir compreensão compartilhada sobre os objetivos, escopo e restrições.
6. As decisões e hipóteses relevantes estiverem suficientemente explícitas.
7. O documento oferecer fundamentação suficiente para derivação das especificações prioritárias.

Critérios considerados não aplicáveis ou temporariamente dispensados deverão possuir justificativa registrada.

A aprovação não deverá depender exclusivamente da quantidade de critérios atendidos, mas da relevância das questões pendentes e de seus impactos.

## 15.12 Transição para as especificações

Após a aprovação inicial do PRD, as especificações deverão ser elaboradas incrementalmente conforme as prioridades do roadmap.

As especificações deverão possuir rastreabilidade com os requisitos, objetivos e jornadas correspondentes.

O detalhamento poderá contemplar:

- Comportamentos esperados.
- Atores e responsabilidades.
- Pré-condições e pós-condições.
- Fluxos principais e alternativos.
- Regras de negócio.
- Invariantes do domínio.
- Estados e transições.
- Critérios de aceitação.
- Restrições de segurança e privacidade.
- Dependências funcionais e técnicas.

A conclusão antecipada de todas as especificações não constitui condição necessária para iniciar a implementação das capacidades suficientemente definidas.

## 15.13 Quality gates das etapas subsequentes

As condições de avanço do projeto deverão ser avaliadas por critérios específicos de maturidade.

| Gate | Finalidade | Condição geral |
| --- | --- | --- |
| G1 — Product Readiness | Aprovação do PRD. | Direção, objetivos, requisitos, escopo e incertezas suficientemente definidos. |
| G2 — Specification Readiness | Preparação de uma capacidade para implementação. | Comportamentos, regras e critérios de aceitação suficientemente especificados. |
| G3 — Engineering Readiness | Preparação técnica para desenvolvimento. | Ambientes, arquitetura inicial, testes, processos e ferramentas adequados. |
| G4 — Release Readiness | Preparação de uma entrega. | Funcionalidades implementadas, verificadas e adequadas ao ambiente de disponibilização. |
| G5 — Pilot Readiness | Preparação para utilização por usuários reais. | Condições funcionais, operacionais, jurídicas, de segurança e privacidade atendidas. |

Os gates poderão ser avaliados incrementalmente por capacidade ou entrega, quando aplicável.

A existência de um gate aprovado não implica aprovação automática dos demais.

## 15.14 Revisão contínua

O PRD deverá permanecer sujeito a revisão durante todo o ciclo de vida do produto.

Novas evidências, alterações de escopo, mudanças regulatórias, riscos identificados e resultados de validação poderão justificar modificações em versões anteriormente aprovadas.

Alterações relevantes deverão ser acompanhadas de:

- Descrição da mudança.
- Justificativa.
- Evidências ou decisões associadas.
- Avaliação de impacto.
- Identificação dos requisitos e especificações afetados.
- Atualização das referências e versões pertinentes.

O objetivo do aceite do PRD é estabelecer uma base confiável para decisões e entregas, sem impedir a evolução do conhecimento sobre o produto.

---

# 16. Questões abertas e decisões pendentes

Esta seção consolida as incertezas, decisões pendentes e questões que necessitam de investigação, validação ou avaliação especializada para orientar a evolução do RAIRE.

A existência de uma questão aberta não implica necessariamente impedimento para aprovação do PRD ou continuidade do desenvolvimento.

Cada questão deverá ser avaliada conforme seu impacto, urgência, dependências e consequências sobre o produto.

Recomendações iniciais representam hipóteses de decisão e não deverão ser interpretadas como decisões aprovadas.

## 16.1 Classificação das questões

| Classificação | Definição |
| --- | --- |
| Bloqueante | Necessita de resolução antes do início da atividade ou especificação afetada. |
| Pré-piloto | Necessita de resolução antes da disponibilização do produto para usuários reais. |
| Investigativa | Depende da obtenção de evidências ou aprofundamento de conhecimento. |
| Evolutiva | Pode ser postergada sem comprometer a entrega inicial, desde que suas consequências sejam conhecidas. |

Uma questão poderá ser reclassificada conforme a evolução do projeto.

## 16.2 Questões estratégicas e de mercado

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-MKT-01 | Qual região deverá receber o piloto? | Liquidez, aquisição e capacidade operacional. | Priorizar uma região na qual a equipe consiga mobilizar oferta e demanda. |
| Q-MKT-02 | Quais categorias deverão compor o catálogo inicial? | Escopo, relevância e densidade do mercado. | Selecionar poucas categorias mediante pesquisa e capacidade de recrutamento. |
| Q-MKT-03 | Quais problemas possuem maior intensidade e frequência? | Proposta de valor e priorização. | Investigar separadamente clientes e profissionais. |
| Q-MKT-04 | Quais segmentos de profissionais apresentam maior necessidade de aquisição de clientes? | Posicionamento e formação da oferta. | Realizar entrevistas e investigar os canais utilizados atualmente. |
| Q-MKT-05 | O MVP deverá contemplar serviços remotos? | Posicionamento, classificação e descoberta geográfica. | Priorizar a hipótese de serviços locais; incluir serviços remotos somente mediante justificativa. |
| Q-MKT-06 | Empresas poderão participar como clientes ou prestadores? | Modelo de identidade, catálogo e posicionamento. | Delimitar inicialmente o público principal, evitando ampliar o modelo sem necessidade. |
| Q-MKT-07 | Como formar a oferta e a demanda iniciais? | Cold start e liquidez. | Investigar parcerias, recrutamento direcionado e aquisição local. |
| Q-MKT-08 | Quais diferenciais são efetivamente percebidos pelos usuários? | Competitividade e adoção. | Validar mediante comparação com canais e alternativas existentes. |

## 16.3 Questões de sustentabilidade e modelo de negócio

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-FIN-01 | Qual modelo de negócio poderá sustentar o RAIRE? | Continuidade financeira e estratégia. | Investigar alternativas sem assumir implementação imediata. |
| Q-FIN-02 | Quem possui disposição a pagar e por qual benefício? | Proposta comercial e geração de receita. | Investigar clientes, profissionais e possíveis parceiros. |
| Q-FIN-03 | Quais funcionalidades poderão ser monetizadas sem comprometer a equidade? | Acessibilidade, confiança e competitividade. | Preservar a independência da classificação orgânica em relação ao pagamento. |
| Q-FIN-04 | Qual orçamento estará disponível para desenvolvimento e operação? | Escopo, infraestrutura e sustentabilidade. | Estabelecer limites financeiros acordados pela equipe. |
| Q-FIN-05 | Qual custo operacional será considerado aceitável? | Viabilidade do piloto e expansão. | Estabelecer indicadores iniciais de custo e esforço operacional. |

## 16.4 Questões de identidade e perfil profissional

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-IDN-01 | Quais confirmações serão obrigatórias para utilização das capacidades do produto? | Segurança, fraude e fricção. | Exigir confirmações proporcionais ao risco de cada operação. |
| Q-IDN-02 | Quais serão os critérios etários de participação? | Proteção de menores e obrigações legais. | Definir antes da disponibilização para usuários reais. |
| Q-IDN-03 | A verificação documental de identidade será necessária no MVP? | Confiança, privacidade e custo operacional. | Não torná-la obrigatória sem evidências ou requisitos legais que a justifiquem. |
| Q-IDN-04 | Quais informações serão obrigatórias para publicação de um perfil profissional? | Elegibilidade, experiência e qualidade da oferta. | Exigir apenas informações necessárias à descoberta e apresentação profissional. |
| Q-IDN-05 | Um profissional poderá oferecer serviços em múltiplas categorias? | Catálogo, descoberta e reputação contextual. | Permitir múltiplos serviços, respeitando regras específicas das categorias. |
| Q-IDN-06 | Como representar disponibilidade profissional? | Compatibilidade dos resultados e atualização. | Utilizar inicialmente informações declaradas, sem agenda automatizada. |
| Q-IDN-07 | Como definir a área de atendimento? | Relevância geográfica e privacidade. | Avaliar alternativas proporcionais ao contexto dos serviços iniciais. |
| Q-IDN-08 | Quais informações do cliente serão públicas? | Confiança, reputação e privacidade. | Limitar a exposição ao necessário para as jornadas correspondentes. |
| Q-IDN-09 | Links externos poderão ser publicados nos perfis? | Autonomia, segurança e compartilhamento de contato. | Avaliar sua compatibilidade com as políticas de exposição e conexão consentida. |

## 16.5 Questões de catálogo e portfólio

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-CAT-01 | Qual será a estrutura inicial do catálogo? | Classificação de serviços, busca e evolução. | Adotar estrutura simples e adequada às categorias selecionadas. |
| Q-CAT-02 | Profissionais poderão sugerir novos serviços ou categorias? | Cobertura e integridade do catálogo. | Considerar sugestões sujeitas à avaliação administrativa. |
| Q-CAT-03 | Quais categorias exigem comprovação de qualificação profissional? | Segurança e responsabilidade jurídica. | Avaliar requisitos aplicáveis antes da inclusão dessas categorias. |
| Q-CAT-04 | Quais formatos de conteúdo serão permitidos no portfólio? | Experiência, armazenamento e moderação. | Priorizar imagens e descrições contextualizadas. |
| Q-CAT-05 | Quais limites serão aplicados ao conteúdo publicado? | Custos, desempenho e prevenção de abuso. | Definir limites proporcionais às necessidades do MVP. |
| Q-CAT-06 | Informações indicativas de preço serão apresentadas nos perfis? | Transparência, comparação e negociação. | Investigar sua utilidade conforme a natureza dos serviços. |
| Q-CAT-07 | O portfólio será obrigatório para publicação do perfil? | Qualidade da oferta e inclusão. | Não exigir universalmente sem justificativa por categoria. |

## 16.6 Questões de descoberta e classificação

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-DES-01 | Quais critérios determinarão a elegibilidade de um profissional em uma busca? | Relevância, equidade e regras de negócio. | Priorizar compatibilidade entre serviço, região e condições essenciais de atendimento. |
| Q-DES-02 | Como a proximidade deverá influenciar a classificação? | Relevância e descoberta local. | Investigar sua importância por categoria de serviço. |
| Q-DES-03 | Quais sinais deverão compor a classificação inicial? | Qualidade, transparência e manipulação. | Adotar critérios simples, documentados e verificáveis. |
| Q-DES-04 | Como tratar informações ausentes ou desatualizadas? | Inclusão e confiabilidade dos resultados. | Diferenciar ausência de informação de evidência negativa. |
| Q-DES-05 | Quais filtros e ordenações serão disponibilizados? | Autonomia e facilidade de descoberta. | Priorizar os critérios considerados relevantes pelos usuários. |
| Q-DES-06 | Como proceder quando não houver oferta compatível? | Experiência, liquidez e conversão. | Comunicar a ausência de resultados e oferecer alternativas pertinentes. |

## 16.7 Questões de demandas, interesses e conexões

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-CON-01 | Quais informações serão obrigatórias para publicação de uma demanda? | Compatibilidade, privacidade e qualidade. | Exigir categoria, descrição e contexto geográfico ou modalidade de atendimento aplicável. |
| Q-CON-02 | Uma demanda poderá ser atendida por múltiplos profissionais? | Estados, interesses e regras de encerramento. | Definir conforme a natureza das necessidades atendidas. |
| Q-CON-03 | Quais alterações serão permitidas após o recebimento de interesses? | Integridade das expectativas dos participantes. | Preservar o contexto original e comunicar alterações relevantes. |
| Q-CON-04 | Qual será a validade de uma demanda publicada? | Atualização da oferta de oportunidades. | Estabelecer regras de expiração e possibilidades de renovação. |
| Q-CON-05 | Quantos interesses poderão ser registrados por demanda ou profissional? | Spam, competitividade e experiência. | Utilizar limites somente quando justificados por segurança ou capacidade operacional. |
| Q-CON-06 | Um profissional poderá manifestar interesse novamente após retirada ou recusa? | Autonomia e prevenção de abuso. | Definir regras conforme o estado da demanda e o histórico da interação. |
| Q-CON-07 | Quais condições estabelecerão uma conexão consentida? | Privacidade, autorização e consistência do domínio. | Exigir manifestação contextual de interesse e aceitação pela contraparte, com autorizações aplicáveis ao compartilhamento. |
| Q-CON-08 | Quais informações de contato poderão ser compartilhadas? | Privacidade e continuidade da interação. | Considerar telefone, e-mail ou outros canais expressamente autorizados. |
| Q-CON-09 | Como tratar a revogação de autorização após o compartilhamento? | Direitos dos usuários e limitações técnicas. | Impedir acessos futuros quando possível e comunicar os limites relativos a dados já obtidos. |
| Q-CON-10 | Haverá comunicação interna entre participantes? | Escopo, custos e segurança. | Não incluir comunicação em tempo real inicialmente; investigar a necessidade de mensagens contextuais. |
| Q-CON-11 | Como tratar conexões repetidas entre os mesmos participantes? | Histórico, métricas e reputação. | Preservar o contexto específico de cada necessidade ou solicitação. |
| Q-CON-12 | Quais serão os efeitos de um bloqueio sobre interações existentes? | Segurança e consistência. | Impedir novas interações e definir o tratamento das conexões e registros anteriores. |

## 16.8 Questões de confiança e reputação

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-TRU-01 | Quais acontecimentos tornarão uma interação elegível para avaliação? | Integridade reputacional. | Considerar evidências adicionais à existência de uma conexão aceita. |
| Q-TRU-02 | Como registrar a declaração de realização de um serviço? | Métricas, reputação e confiabilidade. | Permitir declarações identificáveis de ambas as partes, sem apresentá-las como comprovação independente. |
| Q-TRU-03 | Como tratar declarações divergentes? | Justiça, reputação e contestação. | Preservar a autoria das declarações e estabelecer regras específicas de tratamento. |
| Q-TRU-04 | Qual modelo de avaliação será utilizado? | Clareza e qualidade das informações. | Investigar avaliações estruturadas, comentários e indicadores contextuais. |
| Q-TRU-05 | Quando as avaliações serão publicadas? | Retaliação e manipulação. | Avaliar publicação simultânea ou após encerramento de uma janela de avaliação. |
| Q-TRU-06 | Avaliações poderão ser editadas ou removidas? | Integridade histórica e direitos dos participantes. | Definir condições, prazos e rastreabilidade das alterações. |
| Q-TRU-07 | Como será calculada e apresentada a reputação? | Decisão dos usuários e classificação. | Preservar contexto, volume de avaliações e distinção entre papéis. |
| Q-TRU-08 | A reputação de clientes será publicamente acessível? | Privacidade, confiança e exposição. | Avaliar visibilidade proporcional ao contexto da interação. |
| Q-TRU-09 | Qual será a influência da reputação sobre a descoberta? | Relevância e equidade. | Evitar que ausência de histórico e popularidade acumulada produzam barreiras injustificadas. |
| Q-TRU-10 | Quais sinais de verificação serão apresentados? | Confiança informada. | Comunicar apenas confirmações efetivamente realizadas e suas limitações. |

## 16.9 Questões de equidade

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-EQU-01 | Como identificar profissionais comparáveis para avaliação de equidade? | Validade das métricas. | Considerar elegibilidade, categoria, região, disponibilidade e contexto da busca. |
| Q-EQU-02 | Como caracterizar um profissional novo? | Cold start individual e exposição. | Definir mediante critérios temporais ou de histórico observável. |
| Q-EQU-03 | Quais mecanismos iniciais de distribuição de exposição serão utilizados? | Relevância, complexidade e oportunidades. | Priorizar regras simples e justificáveis. |
| Q-EQU-04 | Quais limites de concentração serão considerados aceitáveis? | Distribuição de oportunidades. | Estabelecer baseline antes de definir limites quantitativos. |
| Q-EQU-05 | Qual degradação de relevância poderá ser considerada aceitável em experimentos? | Equilíbrio entre relevância e equidade. | Definir limites contextualizados antes de cada experimento. |
| Q-EQU-06 | Como investigar efeitos discriminatórios indiretos? | Justiça e responsabilidade. | Avaliar sinais, contexto e possíveis vieses sem introduzir tratamentos indevidos de dados pessoais. |
| Q-EQU-07 | Quem aprovará mudanças de política de classificação? | Governança e responsabilização. | Exigir revisão humana e registro da justificativa, dos riscos e das evidências. |
| Q-EQU-08 | Quais explicações serão apresentadas aos usuários? | Transparência e prevenção de manipulação. | Comunicar fatores relevantes sem prometer precisão explicativa inexistente ou facilitar abusos. |

## 16.10 Questões de segurança, privacidade e conformidade

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-PRV-01 | Quais dados pessoais são necessários para cada finalidade? | Minimização e segurança. | Elaborar inventário de tratamentos. |
| Q-PRV-02 | Quais bases legais serão aplicáveis? | Conformidade com a LGPD. | Avaliar cada finalidade e tratamento individualmente. |
| Q-PRV-03 | Será necessário coletar ou persistir localização precisa? | Privacidade e descoberta geográfica. | Evitar persistência quando a finalidade puder ser atendida por informação aproximada. |
| Q-PRV-04 | Quais prazos de retenção serão aplicados? | Risco jurídico e exposição. | Fundamentar os prazos em finalidade, necessidade e obrigações aplicáveis. |
| Q-PRV-05 | Como serão atendidos os direitos dos titulares? | Conformidade e experiência. | Estabelecer procedimentos, responsabilidades e canais apropriados. |
| Q-PRV-06 | Quais dados poderão ser utilizados para análise e pesquisa? | Privacidade, ética e validação. | Priorizar minimização, agregação e anonimização adequada, com base legal aplicável. |
| Q-PRV-07 | Quais fornecedores terão acesso a informações pessoais? | Dependências externas e proteção de dados. | Avaliar necessidade, responsabilidades, segurança e condições contratuais. |
| Q-PRV-08 | Será necessária a elaboração de um RIPD? | Governança e avaliação de impacto. | Avaliar a necessidade conforme características e riscos dos tratamentos. |
| Q-PRV-09 | Como tratar registros relacionados a fraude, abuso e incidentes? | Investigação, retenção e direitos dos participantes. | Estabelecer finalidade, acesso restrito e critérios proporcionais de retenção. |
| Q-PRV-10 | Quais informações poderão ser disponibilizadas a ferramentas e agentes de IA? | Confidencialidade e segurança. | Priorizar dados sintéticos e impedir compartilhamentos não autorizados. |

## 16.11 Questões operacionais e de moderação

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-OPE-01 | Quem será responsável pelo suporte e moderação? | Segurança e continuidade operacional. | Distribuir responsabilidades explicitamente antes do piloto. |
| Q-OPE-02 | Qual volume de participantes poderá ser atendido pela equipe? | Capacidade operacional e expansão. | Estabelecer limites conforme recursos disponíveis. |
| Q-OPE-03 | Como serão classificadas as denúncias? | Priorização e resposta a incidentes. | Classificar por natureza, gravidade, urgência e risco. |
| Q-OPE-04 | Quais serão os prazos de atendimento? | Proteção dos usuários e expectativa de serviço. | Definir prazos diferenciados conforme criticidade. |
| Q-OPE-05 | Quais medidas administrativas poderão ser aplicadas? | Integridade, proporcionalidade e governança. | Estabelecer políticas de advertência, restrição, remoção e suspensão. |
| Q-OPE-06 | Como funcionará a contestação de decisões? | Consistência e justiça procedimental. | Permitir revisão fundamentada e registro das decisões. |
| Q-OPE-07 | Quais atividades poderão ser automatizadas? | Eficiência e responsabilidade. | Priorizar assistência operacional, preservando revisão apropriada em decisões de alto impacto. |
| Q-OPE-08 | Quais situações determinarão suspensão do piloto ou de funcionalidades? | Segurança e continuidade. | Estabelecer critérios de interrupção antes da operação real. |

## 16.12 Questões de arquitetura e engenharia

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-TEC-01 | Qual stack tecnológica será utilizada? | Capacidade de entrega e manutenção. | Avaliar alternativas conforme experiência da equipe e requisitos. |
| Q-TEC-02 | Qual estrutura arquitetural atenderá ao estágio inicial? | Complexidade, testabilidade e evolução. | Priorizar uma arquitetura proporcional ao escopo e aos recursos disponíveis. |
| Q-TEC-03 | Como serão delimitadas as responsabilidades dos componentes? | Coesão, acoplamento e evolução. | Derivar fronteiras das capacidades, regras e conceitos do domínio. |
| Q-TEC-04 | Qual estratégia será utilizada para persistência e consulta geográfica? | Consistência, relevância e desempenho. | Avaliar alternativas conforme os comportamentos especificados. |
| Q-TEC-05 | Qual infraestrutura de hospedagem será utilizada? | Custos, segurança e disponibilidade. | Avaliar custo total, complexidade operacional e necessidades reais. |
| Q-TEC-06 | Qual estratégia será utilizada para armazenamento de imagens? | Custos, desempenho e privacidade. | Avaliar armazenamento, processamento, acesso e remoção de conteúdo. |
| Q-TEC-07 | Como serão preservadas as invariantes em operações concorrentes? | Integridade das demandas, interesses e conexões. | Definir invariantes nas especificações e avaliar mecanismos técnicos correspondentes. |
| Q-TEC-08 | Quais controles e ferramentas comporão o harness de desenvolvimento com IA? | Qualidade, segurança e produtividade. | Estabelecer orientação contextual, verificações automatizadas e revisão humana. |
| Q-TEC-09 | Quais objetivos de recuperação e disponibilidade serão aplicados? | Resiliência e continuidade. | Definir metas proporcionais aos riscos e estágio operacional. |
| Q-TEC-10 | Como serão organizadas as estratégias de testes e entrega? | Confiabilidade e produtividade. | Estabelecer práticas incrementais orientadas aos riscos e requisitos. |

As decisões arquiteturais deverão ser encaminhadas aos artefatos de engenharia correspondentes, incluindo ADRs quando apropriado.

## 16.13 Questões de pesquisa, métricas e validação

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-VAL-01 | Qual definição operacional será utilizada para conexão útil? | Métrica central do produto. | Investigar formas de identificar valor percebido por ambas as partes. |
| Q-VAL-02 | Como avaliar a relevância dos resultados? | Descoberta, ranking e experimentação. | Combinar critérios de compatibilidade com avaliações dos usuários. |
| Q-VAL-03 | Qual densidade mínima será necessária por segmento? | Liquidez e preparação do piloto. | Fundamentar mediante pesquisa e observação inicial. |
| Q-VAL-04 | Quais indicadores deverão ser instrumentados desde a primeira entrega? | Capacidade de aprendizado. | Priorizar eventos essenciais das jornadas e controles críticos. |
| Q-VAL-05 | Como coletar informações sobre desfechos externos? | Qualidade e representatividade das evidências. | Utilizar feedback voluntário e preservar a natureza declaratória dos dados. |
| Q-VAL-06 | Qual será a duração do piloto? | Planejamento e suficiência das evidências. | Definir conforme características do mercado e condições de avaliação. |
| Q-VAL-07 | Quais metas quantitativas serão consideradas adequadas? | Avaliação de sucesso. | Estabelecer baseline e justificar metas antes da avaliação correspondente. |
| Q-VAL-08 | Quais critérios determinarão continuidade, revisão ou suspensão? | Governança e investimento. | Estabelecer critérios anteriores ao início do piloto. |
| Q-VAL-09 | Quando experimentos controlados serão metodologicamente viáveis? | Confiabilidade das conclusões. | Avaliar amostra, dependências, unidade experimental e interferências. |
| Q-VAL-10 | Como avaliar a representatividade dos participantes? | Generalização e qualidade da pesquisa. | Documentar características da amostra e limitações das conclusões. |

## 16.14 Questões de governança e planejamento

| ID | Questão | Impacto | Recomendação inicial |
| --- | --- | --- | --- |
| Q-GOV-01 | Qual será o escopo da entrega para o hackathon? | Prazo, esforço e apresentação. | Priorizar a demonstração das jornadas centrais e dos diferenciais do produto. |
| Q-GOV-02 | Como serão distribuídas as responsabilidades entre integrantes? | Capacidade de execução e continuidade. | Definir responsabilidades e responsáveis por decisões relevantes. |
| Q-GOV-03 | Quem será responsável pela aprovação e manutenção do PRD? | Consistência documental. | Estabelecer responsabilidade explícita e revisão colaborativa. |
| Q-GOV-04 | Como serão aprovadas mudanças de escopo? | Controle de prioridades e esforço. | Avaliar impactos, justificativas e dependências. |
| Q-GOV-05 | Quais decisões poderão ser delegadas a agentes de IA? | Autonomia, qualidade e responsabilidade. | Delegar investigação e execução controlada, preservando aprovação humana das decisões críticas. |
| Q-GOV-06 | Como serão registradas decisões e evidências? | Rastreabilidade e evolução. | Manter registros versionados e relacionados aos artefatos afetados. |
| Q-GOV-07 | Quais critérios deverão autorizar cada release? | Qualidade, segurança e previsibilidade. | Definir critérios proporcionais ao ambiente e público da entrega. |
| Q-GOV-08 | Qual será o processo de resolução de divergências entre integrantes? | Governança e continuidade do projeto. | Estabelecer responsabilidades decisórias e mecanismos de discussão fundamentada. |

## 16.15 Registro e acompanhamento das questões

Cada questão deverá possuir, quando aplicável:

- Identificador único.
- Descrição.
- Categoria.
- Contexto e justificativa.
- Impacto potencial.
- Classificação de prioridade.
- Dependências.
- Atividade ou marco afetado.
- Alternativas conhecidas.
- Recomendação inicial.
- Evidências necessárias.
- Responsável.
- Estado atual.
- Prazo ou condição para resolução.
- Decisão resultante.
- Referência ao registro de decisão correspondente.

Os registros deverão permanecer disponíveis em um catálogo de questões e decisões pendentes.

## 16.16 Estados das questões

| Estado | Definição |
| --- | --- |
| Aberta | Questão identificada, ainda não investigada. |
| Em investigação | Evidências ou alternativas estão sendo analisadas. |
| Aguardando decisão | Informações suficientes estão disponíveis para deliberação. |
| Resolvida | Decisão tomada e registrada. |
| Adiada | Resolução postergada mediante justificativa e condição de revisão. |
| Superada | Questão deixou de ser relevante devido a mudanças no produto ou contexto. |

Questões resolvidas deverão permanecer rastreáveis quando possuírem relevância histórica.

## 16.17 Processo de resolução

A resolução de uma questão deverá considerar:

1. Definição do problema decisório.
2. Identificação das alternativas.
3. Investigação das evidências disponíveis.
4. Avaliação dos impactos e trade-offs.
5. Identificação de riscos e dependências.
6. Deliberação pelos responsáveis.
7. Registro da decisão e justificativa.
8. Atualização dos artefatos afetados.
9. Definição das condições de revisão, quando aplicável.

Questões poderão ser resolvidas por decisões provisórias quando existir justificativa e possibilidade de revisão.

## 16.18 Participação de agentes de IA

Agentes de IA poderão auxiliar nas atividades de:

- Identificação de inconsistências e questões abertas.
- Investigação e organização de informações.
- Comparação de alternativas.
- Análise de impactos e trade-offs.
- Identificação de riscos.
- Formulação de recomendações.
- Preparação de propostas de decisão.
- Atualização controlada da documentação.

Agentes não deverão interpretar recomendações iniciais, hipóteses ou questões abertas como decisões aprovadas.

Decisões relacionadas ao posicionamento do produto, alterações relevantes de escopo, tratamento de dados pessoais, políticas de equidade, segurança dos participantes, obrigações legais e compromissos financeiros deverão possuir aprovação humana apropriada.

A ausência de uma decisão não deverá autorizar a criação silenciosa de regras de negócio ou restrições que modifiquem o comportamento esperado do produto.

Quando uma questão impeditiva for identificada durante a elaboração de especificações ou implementação, ela deverá ser encaminhada aos responsáveis, acompanhada de contexto, alternativas e análise de impacto.

## 16.19 Revisão contínua

O catálogo de questões abertas deverá acompanhar a evolução do RAIRE.

Novas questões poderão surgir durante atividades de pesquisa, especificação, arquitetura, desenvolvimento, operação e validação.

A resolução de uma questão deverá resultar na atualização dos documentos, requisitos, regras ou decisões correspondentes.

O objetivo desta seção não é eliminar antecipadamente todas as incertezas, mas garantir que decisões relevantes sejam tomadas de maneira consciente, fundamentada e rastreável.

---

## 17. Rastreabilidade

# 17. Rastreabilidade

A rastreabilidade deverá estabelecer relações explícitas entre os problemas identificados, hipóteses, evidências, objetivos, jornadas, requisitos, regras de negócio, especificações, decisões, entregas e atividades de verificação e validação do RAIRE.

Seu objetivo é preservar a coerência entre a fundamentação estratégica do produto e sua realização técnica, permitindo identificar lacunas, inconsistências, dependências e impactos decorrentes de mudanças.

## 17.1 Objetivos da rastreabilidade

A estratégia de rastreabilidade deverá permitir:

- Identificar a origem e a justificativa dos requisitos.
- Relacionar funcionalidades aos objetivos e necessidades dos usuários.
- Identificar objetivos sem capacidades correspondentes.
- Identificar requisitos sem fundamentação suficiente.
- Identificar requisitos ainda não especificados.
- Verificar a cobertura de regras de negócio e critérios de aceitação.
- Relacionar decisões arquiteturais aos requisitos e restrições que as motivaram.
- Identificar requisitos ainda não implementados ou verificados.
- Relacionar hipóteses às atividades de pesquisa e validação.
- Avaliar impactos decorrentes de alterações.
- Preservar o histórico das decisões relevantes.
- Fornecer contexto consistente para atividades humanas e assistidas por agentes de IA.

## 17.2 Princípios

A rastreabilidade deverá observar os seguintes princípios:

- **Bidirecionalidade:** permitir identificar tanto a origem quanto os elementos derivados de um artefato.
- **Identificação estável:** utilizar identificadores únicos e persistentes.
- **Granularidade proporcional:** registrar relações relevantes sem introduzir complexidade administrativa desnecessária.
- **Consistência:** evitar referências inexistentes, obsoletas ou contraditórias.
- **Versionamento:** preservar alterações relevantes e seus impactos.
- **Fonte de verdade:** evitar a manutenção de múltiplas definições independentes para o mesmo requisito.
- **Atualização incremental:** estabelecer relações conforme os artefatos sejam produzidos.
- **Fundamentação:** distinguir relações propostas de relações confirmadas.
- **Verificabilidade:** permitir avaliação humana e automatizada da integridade das referências.

## 17.3 Identificação dos artefatos

| Artefato | Identificador | Exemplo |
| --- | --- | --- |
| Problema | PRB | PRB-01 |
| Hipótese | H | H-01 |
| Evidência | EVD | EVD-01 |
| Objetivo | OBJ | OBJ-03 |
| Objetivo de validação | VAL | VAL-03 |
| Objetivo de qualidade | QLT | QLT-01 |
| Jornada | JRN | JRN-04 |
| Job to Be Done | JTBD | JTBD-C01 |
| Requisito funcional | RF | RF-06.01 |
| Requisito não funcional | RNF | RNF-08 |
| Regra de negócio | BR | BR-CON-01 |
| Invariante de domínio | INV | INV-CON-01 |
| Cenário de atributo de qualidade | QAS | QAS-01 |
| Especificação | SPEC | SPEC-DES-01 |
| Critério de aceitação | AC | AC-DES-01 |
| Decisão de produto | PDR | PDR-001 |
| Decisão arquitetural | ADR | ADR-001 |
| Risco | RISK | RISK-01 |
| Questão aberta | Q | Q-CON-07 |
| Experimento | EXP | EXP-01 |
| Métrica | MET | MET-01 |
| Caso de teste | TC | TC-CON-01 |

Os identificadores existentes deverão ser preservados ou formalmente reconciliados em caso de alteração.

A convenção apresentada deverá ser aplicada conforme os artefatos correspondentes sejam definidos.

## 17.4 Objetivos → requisitos funcionais

A matriz inicial deverá relacionar os objetivos estabelecidos na seção 4 às famílias de requisitos funcionais da seção 8.

| Objetivo | Resultado pretendido | Famílias de requisitos relacionadas |
| --- | --- | --- |
| OBJ-01 | Presença digital profissional especializada | RF-01, RF-02, RF-03, RF-05 |
| OBJ-02 | Apresentação contextualizada de competências e trabalhos | RF-03, RF-04 |
| OBJ-03 | Descoberta relevante de profissionais | RF-03, RF-04, RF-05, RF-06, RF-07, RF-14 |
| OBJ-04 | Publicação de necessidades e recebimento de interesses | RF-08, RF-10, RF-16 |
| OBJ-05 | Descoberta ativa de oportunidades profissionais | RF-05, RF-09, RF-10 |
| OBJ-06 | Estabelecimento de conexões relevantes e consentidas | RF-08, RF-09, RF-10, RF-11, RF-12, RF-13, RF-16, RF-17 |
| OBJ-07 | Confiança e reputação bilateral | RF-04, RF-12, RF-13, RF-14, RF-15, RF-17, RF-18, RF-19 |
| OBJ-08 | Redução de barreiras injustificadas de exposição | RF-03, RF-06, RF-07, RF-20 |
| OBJ-09 | Validação da liquidez e viabilidade do mercado inicial | RF-18, RF-20 |

Os requisitos de instrumentação, proteção dos usuários, privacidade e administração poderão contribuir transversalmente para múltiplos objetivos.

Esta matriz representa uma associação inicial em nível de família. A rastreabilidade detalhada deverá utilizar os identificadores individuais dos requisitos.

## 17.5 Hipóteses → objetivos de validação → requisitos

As hipóteses do produto deverão ser relacionadas aos objetivos de validação e às capacidades que permitem investigá-las.

| Linha de investigação | Objetivo de validação | Famílias de requisitos relacionadas |
| --- | --- | --- |
| Existência e intensidade dos problemas identificados | VAL-01 | Pesquisa primária; não depende necessariamente da implementação de requisitos. |
| Geração de valor para clientes e profissionais | VAL-02 | RF-03, RF-04, RF-06, RF-08, RF-09, RF-10, RF-11, RF-12, RF-13 |
| Complementaridade entre descoberta direta e publicação de necessidades | VAL-03 | RF-06, RF-08, RF-09, RF-10, RF-11, RF-12 |
| Relevância de localização, disponibilidade, especialidade, portfólio e reputação | VAL-04 | RF-03, RF-04, RF-05, RF-06, RF-07, RF-09, RF-14 |
| Contribuição dos mecanismos de confiança | VAL-05 | RF-12, RF-13, RF-14, RF-15, RF-17, RF-18 |
| Distribuição de exposição sem comprometimento indevido da relevância | VAL-06 | RF-06, RF-07, RF-20 |
| Usabilidade e acessibilidade das jornadas | VAL-07 | Requisitos das jornadas avaliadas e RNFs correspondentes. |

A associação definitiva com hipóteses individuais deverá ser estabelecida após a consolidação do catálogo de hipóteses.

Um requisito implementado não constitui evidência suficiente de confirmação de uma hipótese.

A validação deverá depender de métodos de investigação apropriados, resultados observados e análise das limitações das evidências.

## 17.6 Jornadas → requisitos funcionais

| Jornada | Referência no PRD | Famílias de requisitos relacionadas |
| --- | --- | --- |
| Exploração inicial e cadastro | 6.1 | RF-01, RF-02, RF-06 |
| Ativação profissional | 6.2 | RF-01, RF-02, RF-03, RF-05 |
| Gerenciamento de perfil e portfólio | 6.3 | RF-03, RF-04 |
| Descoberta de profissionais | 6.4 | RF-03, RF-04, RF-05, RF-06, RF-07, RF-14, RF-15 |
| Conexão direta | 6.5 | RF-11, RF-12, RF-16, RF-17 |
| Publicação e gerenciamento de demandas | 6.6 | RF-08, RF-10, RF-12, RF-16 |
| Descoberta de oportunidades | 6.7 | RF-09, RF-10, RF-12, RF-16 |
| Interação e registro de desfechos | 6.8 | RF-12, RF-13 |
| Avaliação e reputação | 6.9 | RF-13, RF-14, RF-17 |
| Denúncia e moderação | 6.10 | RF-17, RF-18 |

Os requisitos transversais de privacidade, instrumentação, segurança e administração deverão ser associados adicionalmente às jornadas nas quais sejam aplicáveis.

## 17.7 Requisitos → especificações

As especificações deverão detalhar comportamentos esperados, regras de negócio, estados, restrições, invariantes, fluxos alternativos e critérios de aceitação.

A seguinte organização representa uma proposta inicial de agrupamento por capacidade funcional.

| Especificação proposta | Famílias de requisitos relacionadas |
| --- | --- |
| identity-access | RF-01 |
| user-profile | RF-02 |
| professional-profile | RF-03 |
| professional-portfolio | RF-04 |
| service-catalog | RF-05 |
| local-discovery | RF-06 |
| fair-visibility | RF-07 |
| service-request | RF-08 |
| opportunity-discovery | RF-09 |
| service-interest | RF-10 |
| direct-connection-request | RF-11 |
| consented-connection | RF-12 |
| interaction-outcome | RF-13 |
| reviews-reputation | RF-14 |
| verification-trust-signals | RF-15 |
| notifications | RF-16 |
| user-protection | RF-17 |
| administration-moderation | RF-18 |
| privacy-data-rights | RF-19 |
| measurement-analytics | RF-20 |

Esses agrupamentos não representam necessariamente documentos individuais, módulos de software, serviços independentes ou bounded contexts.

Uma especificação poderá contemplar múltiplos requisitos, e um requisito poderá possuir comportamentos detalhados em mais de uma especificação.

A decomposição definitiva deverá considerar coesão funcional, dependências, complexidade e estratégia de entrega.

## 17.8 Requisitos não funcionais → áreas de aplicação

Os requisitos não funcionais deverão ser relacionados às capacidades, componentes e cenários nos quais seus atributos de qualidade possam ser definidos e verificados.

| Requisito | Atributo de qualidade | Áreas iniciais de aplicação |
| --- | --- | --- |
| RNF-01 | Manutenibilidade | Engenharia e arquitetura. |
| RNF-02 | Desempenho | Descoberta, interfaces e operações críticas. |
| RNF-03 | Capacidade | Infraestrutura, dados e processamento. |
| RNF-04 | Disponibilidade | Jornadas críticas e operação. |
| RNF-05 | Confiabilidade e integridade | Identidade, demandas, interesses, conexões e avaliações. |
| RNF-06 | Resiliência | Integrações, notificações e dependências. |
| RNF-07 | Recuperabilidade | Persistência, backups e recuperação operacional. |
| RNF-08 | Segurança | Todas as capacidades aplicáveis. |
| RNF-09 | Privacidade | Tratamento, acesso, exposição e descarte de dados. |
| RNF-10 | Acessibilidade | Interfaces e jornadas de interação. |
| RNF-11 | Usabilidade e inclusão | Jornadas de clientes, profissionais e operadores. |
| RNF-12 | Responsividade | Interfaces web e dispositivos representativos. |
| RNF-13 | Observabilidade | Aplicação, infraestrutura e operação. |
| RNF-14 | Auditabilidade | Operações críticas, moderação e alterações de políticas. |
| RNF-15 | Testabilidade | Domínio, aplicação, integrações e interfaces. |
| RNF-16 | Equidade e explicabilidade | Descoberta, classificação e distribuição de exposição. |
| RNF-17 | Eficiência de custos | Infraestrutura, armazenamento e operação. |
| RNF-18 | Deployability | Ambientes, processos de entrega e evolução de dados. |
| RNF-19 | Portabilidade | Dados, integrações e dependências tecnológicas. |
| RNF-20 | Evolutividade | Arquitetura, domínio e políticas do produto. |

A verificação dos requisitos não funcionais deverá considerar critérios mensuráveis ou cenários de atributos de qualidade, conforme aplicável.

## 17.9 Especificações → regras de negócio e invariantes

Cada especificação deverá identificar as regras de negócio e invariantes necessárias à preservação do comportamento esperado.

Deverão ser diferenciadas:

- Regras de elegibilidade.
- Regras de autorização.
- Regras de visibilidade.
- Regras de classificação.
- Regras de transição de estado.
- Restrições temporais.
- Restrições de integridade.
- Políticas configuráveis.
- Invariantes do domínio.

Regras compartilhadas por múltiplas especificações deverão possuir uma definição canônica, evitando duplicação e divergências.

## 17.10 Requisitos e especificações → decisões arquiteturais

Decisões arquiteturais relevantes deverão identificar:

- Requisitos que motivaram a decisão.
- Atributos de qualidade envolvidos.
- Restrições aplicáveis.
- Alternativas consideradas.
- Trade-offs.
- Consequências sobre as capacidades afetadas.
- Condições que poderão justificar revisão futura.

Não será obrigatório produzir um ADR para cada requisito ou escolha local de implementação.

ADRs deverão ser utilizados para decisões arquiteturalmente significativas.

## 17.11 Requisitos → verificação

A verificação deverá estabelecer evidências de que os comportamentos implementados atendem aos requisitos e especificações aprovados.

| Elemento | Evidência de verificação |
| --- | --- |
| Requisito funcional | Critérios de aceitação e testes correspondentes. |
| Regra de negócio | Testes dos comportamentos permitidos e proibidos. |
| Invariante | Testes de integridade, concorrência e condições excepcionais, quando aplicáveis. |
| Requisito não funcional | Medições, testes, inspeções ou cenários de qualidade. |
| Controle de segurança | Testes de autorização, abuso e vulnerabilidades aplicáveis. |
| Controle de privacidade | Testes de acesso, visibilidade, tratamento e descarte. |
| Acessibilidade | Avaliações automatizadas e manuais pertinentes. |
| Decisão arquitetural | Revisão e evidências de atendimento às restrições estabelecidas. |

As evidências deverão identificar a versão ou configuração avaliada e as condições relevantes da verificação.

A existência de um teste não implica automaticamente cobertura suficiente de um requisito.

## 17.12 Hipóteses → experimentos → evidências

As hipóteses deverão possuir rastreabilidade com as atividades utilizadas para investigá-las.

A relação deverá contemplar, quando aplicável:

- Hipótese investigada.
- Objetivo de validação.
- Método de pesquisa ou experimento.
- Participantes ou população observada.
- Métricas utilizadas.
- Critérios de avaliação.
- Resultados obtidos.
- Limitações.
- Evidências produzidas.
- Conclusão.
- Decisão resultante.

Os métodos poderão incluir entrevistas, testes de usabilidade, observação, estudos comparativos, análise de funil, avaliação offline e experimentos controlados.

A existência de um experimento não pressupõe a confirmação da hipótese correspondente.

Resultados inconclusivos ou contrários à hipótese também deverão ser registrados.

## 17.13 Riscos → controles e evidências

Os riscos relevantes deverão ser associados aos mecanismos utilizados para preveni-los, detectá-los, mitigá-los ou responder à sua materialização.

| Risco | Elementos potencialmente relacionados |
| --- | --- |
| Baixa liquidez | Estratégia de aquisição, objetivos de validação e métricas do piloto. |
| Fraude e perfis falsos | RF-01, RF-15, RF-17, RF-18, RNF-08. |
| Assédio e exposição indevida de contato | RF-11, RF-12, RF-17, RF-18, RF-19, RNF-08, RNF-09. |
| Exposição indevida de localização | RF-03, RF-06, RF-08, RF-09, RF-19, RNF-08, RNF-09. |
| Manipulação de avaliações | RF-13, RF-14, RF-17, RF-18, RNF-05. |
| Distribuição injustificada de exposição | RF-06, RF-07, RF-20, RNF-16. |
| Exclusão por barreiras de acesso | RNF-10, RNF-11, RNF-12. |
| Insuficiência de capacidade operacional | RF-18, RNF-03, RNF-17 e planejamento operacional. |

A associação de um risco a um controle não comprova sua mitigação efetiva.

A conclusão do tratamento deverá depender de avaliação e evidências apropriadas.

## 17.14 Matriz de rastreabilidade detalhada

A matriz operacional deverá permitir identificar a situação de cada requisito individual.

Poderá contemplar os seguintes campos:

| Campo | Finalidade |
| --- | --- |
| Identificador | Identificar unicamente o requisito. |
| Origem | Identificar o problema, objetivo, jornada, risco ou restrição que o motivou. |
| Prioridade | Representar sua importância para a entrega correspondente. |
| Especificações | Identificar os artefatos que detalham seu comportamento. |
| Regras e invariantes | Identificar comportamentos e restrições associados. |
| Critérios de aceitação | Identificar as condições verificáveis. |
| Decisões relacionadas | Identificar decisões de produto ou arquitetura pertinentes. |
| Entregas | Identificar épicos, histórias ou incrementos responsáveis pela implementação. |
| Evidências de verificação | Identificar testes, inspeções ou medições. |
| Estado de aprovação | Indicar a maturidade documental. |
| Estado de implementação | Indicar a situação da entrega. |
| Estado de verificação | Indicar a situação das evidências de conformidade. |
| Dependências | Identificar requisitos ou decisões relacionados. |
| Versão | Identificar a versão vigente do requisito. |

A matriz poderá ser produzida a partir das referências mantidas nos próprios artefatos, evitando duplicação manual de informações.

## 17.15 Integridade e cobertura

A integridade da rastreabilidade deverá considerar:

- Ausência de identificadores duplicados.
- Existência dos elementos referenciados.
- Ausência de referências a elementos removidos sem tratamento adequado.
- Identificação de requisitos sem origem ou justificativa.
- Identificação de requisitos prioritários sem especificação quando necessária.
- Identificação de requisitos sem critérios de aceitação ou estratégia de verificação.
- Identificação de controles críticos sem evidências.
- Identificação de funcionalidades entregues sem relação com requisitos ou decisões aprovadas.
- Identificação de hipóteses prioritárias sem estratégia de investigação.
- Identificação de questões impeditivas ainda não resolvidas.

Indicadores de cobertura deverão considerar o escopo e o estágio da entrega.

A cobertura quantitativa da rastreabilidade não deverá ser interpretada isoladamente como evidência de qualidade, conformidade ou geração de valor.

## 17.16 Análise de impacto de mudanças

Alterações relevantes deverão permitir identificar os artefatos potencialmente afetados.

A análise poderá contemplar:

- Objetivos e hipóteses.
- Jornadas.
- Requisitos.
- Regras de negócio.
- Invariantes.
- Especificações.
- Decisões de produto.
- Decisões arquiteturais.
- Contratos e modelos de dados.
- Implementações existentes.
- Testes automatizados.
- Métricas e experimentos.
- Políticas de segurança e privacidade.
- Riscos e dependências.
- Documentação operacional.

A necessidade de atualização deverá ser avaliada conforme a natureza da alteração.

## 17.17 Rastreabilidade no desenvolvimento assistido por IA

Agentes de IA deverão utilizar referências explícitas para identificar os requisitos, especificações, regras de negócio e decisões aplicáveis às atividades sob sua responsabilidade.

Antes de implementar uma capacidade, deverão identificar, quando aplicável:

- Requisitos relacionados.
- Especificações vigentes.
- Regras de negócio e invariantes.
- Critérios de aceitação.
- Restrições arquiteturais.
- Controles de segurança e privacidade.
- Questões abertas impeditivas.
- Dependências relevantes.

Durante a execução, não deverão criar silenciosamente requisitos, regras de negócio ou decisões de produto ausentes da documentação aprovada.

Após a implementação, deverão disponibilizar informações que permitam identificar:

- Requisitos atendidos.
- Alterações realizadas.
- Testes executados.
- Evidências produzidas.
- Limitações conhecidas.
- Divergências identificadas.
- Questões dependentes de revisão humana.

A rastreabilidade produzida por agentes deverá permanecer sujeita à verificação.

## 17.18 Manutenção e governança

A responsabilidade pela rastreabilidade deverá acompanhar a responsabilidade pelos artefatos correspondentes.

A equipe deverá estabelecer mecanismos proporcionais para:

- Preservar identificadores.
- Atualizar referências.
- Registrar alterações.
- Verificar integridade.
- Revisar relações relevantes.
- Identificar lacunas de cobertura.
- Avaliar impactos.
- Preservar evidências históricas.

A automação deverá ser adotada quando reduzir esforço de manutenção ou aumentar a confiabilidade do processo.

A rastreabilidade deverá apoiar o desenvolvimento e a tomada de decisões, evitando transformar-se em uma atividade burocrática desvinculada da geração de valor.
