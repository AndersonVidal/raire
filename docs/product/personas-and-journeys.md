# Personas, Atores e Jornadas — RAIRE

**Status:** em revisão  
**Versão:** 1.0  
**Fonte:** PRD v1.0, seções 4 e 5  
**Uso:** orientar UX, copy, critérios de aceite e testes de usabilidade.

---

## 1. Atores do sistema

| Ator | Descrição | Pode | Não pode |
| --- | --- | --- | --- |
| Visitante | pessoa sem conta | buscar, ver perfis públicos, iniciar cadastro | publicar perfil, criar pedido, enviar interesse |
| Cliente | pessoa que busca serviço | buscar, criar pedido, receber interesses, aceitar conexão, avaliar | acessar dados privados de profissionais |
| Profissional | autônomo que oferece serviço | criar/publicar perfil, portfólio, ver oportunidades, enviar interesse, aceitar conexão, avaliar | acessar contato de cliente antes do aceite |
| Operador | equipe de operação | analisar denúncias, moderar conteúdo, aplicar restrições, auditar | alterar ranking sem aprovação; acessar PII sem necessidade |
| Administrador | dono do sistema | configurar catálogo, papéis, flags, políticas; auditar | vender posição; usar atributo sensível |
| Pesquisador | analista de produto/pesquisa | ver métricas agregadas, criar experimentos, exportar dados anonimizados | acessar PII ou reidentificar usuários |

**Regra:** uma mesma conta pode atuar como cliente e profissional, mas os perfis e permissões são separados.

---

## 2. Personas primárias

### P1 — Cliente urgente

**Contexto:** precisa resolver um problema hoje (eletricista, encanador, chaveiro).  
**Dispositivo:** celular, conexão instável, pouca paciência.  
**Perguntas:** atende na minha região? é confiável? como falo?  
**Sucesso:** encontrar em até 2 minutos, entender o próximo passo, não expor endereço.  
**Frustrações:** formulário longo, resultado irrelevante, contato escondido sem explicação.

### P2 — Cliente planejador

**Contexto:** reforma ou serviço recorrente (arquiteto, personal trainer, diarista).  
**Comportamento:** compara opções, preço e reputação; volta depois.  
**Perguntas:** qual a diferença entre eles? o que dizem as avaliações?  
**Sucesso:** comparar 3–5 profissionais, salvar/retomar, publicar necessidade.  
**Frustrações:** nota única sem contexto, avaliação genérica, ausência de portfólio.

### P3 — Profissional invisível

**Contexto:** experiência real, pouca ou nenhuma presença digital.  
**Dependência:** indicação e WhatsApp; usa celular.  
**Barreiras:** formulário longo, termos técnicos, upload obrigatório, medo de exposição.  
**Sucesso:** criar perfil sem ajuda, receber oportunidade compatível, entender por que aparece ou não.  
**Frustrações:** jargão de marketing, exigência de documento no primeiro acesso, ranking opaco.

### P4 — Profissional em crescimento

**Contexto:** já usa redes sociais e quer mais clientes.  
**Expectativa:** aprender, mas não depender de anúncios caros.  
**Perguntas:** como estou performando? o que melhora meu perfil?  
**Sucesso:** receber pedidos qualificados, responder rápido, construir reputação.  
**Frustrações:** competir com quem paga por posição, não entender a distribuição de visibilidade.

### P5 — Operador de confiança

**Contexto:** analisa denúncias, perfis suspeitos e avaliações.  
**Necessidade:** contexto, histórico, critérios claros e ações auditáveis.  
**Sucesso:** resolver casos com consistência, sem abuso de poder.  
**Frustrações:** fila sem prioridade, ferramenta que não registra decisão, pressão por remover crítica legítima.

---

## 3. Personas secundárias / operacionais

| Persona | Necessidade | Risco se falhar |
| --- | --- | --- |
| Administrador | configurar catálogo, flags e políticas com segurança | mudança sem auditoria, abuso de poder |
| Pesquisador | medir funil, fairness e satisfação sem PII | decisão baseada em métrica distorcida |
| Parceiro local | indicar profissionais e acompanhar o piloto | expectativa desalinhada, dados expostos |
| Suporte | ajudar usuários com acesso, perfil e denúncia | atendimento lento, exposição de dados |

---

## 4. Jobs to be done

### Cliente

- "Quando preciso de um serviço local, quero encontrar alguém confiável perto de mim sem depender só de indicação."
- "Quando comparo profissionais, quero entender o que cada um faz e o que outras pessoas viveram."
- "Quando falo com alguém, quero escolher como meu contato é compartilhado."
- "Quando algo dá errado, quero denunciar e entender o que a plataforma fez."

### Profissional

- "Quando sou autônomo, quero ser encontrado por clientes da minha região sem pagar por visibilidade."
- "Quando recebo uma oportunidade, quero saber se faz sentido para mim antes de expor meu contato."
- "Quando construo reputação, quero que meu histórico e meu contexto contem, não só uma nota."
- "Quando algo injusto acontece, quero contestar e receber uma resposta."

### Operação

- "Quando recebo uma denúncia, quero priorizar risco e decidir com evidência."
- "Quando aplico uma medida, quero que ela seja auditável e recorrível."
- "Quando mudo uma política, quero simular, aprovar e reverter."

### Pesquisa

- "Quando avalio fairness, quero comparar coortes observáveis sem usar atributos sensíveis."
- "Quando meço sucesso, quero conexão de qualidade, não clique isolado."

---

## 5. Jornadas principais

### 5.1 Descoberta inicial e criação de conta

**Fluxo principal:**

1. Visitante acessa a home e entende a proposta em uma frase.
2. Escolhe "Sou profissional" ou "Preciso de um serviço".
3. Busca sem conta ou inicia cadastro.
4. Informa e-mail, senha, nome de exibição e aceita termos/privacidade.
5. Confirma o e-mail.
6. Retoma a ação original.

**Resultado:** conta criada, e-mail confirmado, contexto preservado.  
**Fricções a evitar:** cadastro antes de ver valor; senha com regras absurdas; termos inacessíveis; perda da busca ao autenticar.

### 5.2 Ativação do perfil profissional

**Fluxo principal:**

1. Profissional escolhe "Ofereço serviços".
2. Informa nome de exibição, bio curta e categorias.
3. Escolhe tipo de atendimento: presencial, remoto ou híbrido.
4. Se presencial/híbrido, define cidade/bairro e raio.
5. Informa disponibilidade e preferência de contato.
6. Vê prévia pública.
7. Publica.
8. Recebe orientação sobre ranking e próximos passos.

**Resultado:** perfil publicado, elegível para busca.  
**Fricções a evitar:** documento obrigatório, foto obrigatória, formulário em uma tela gigante, jargão.

### 5.3 Gestão de perfil e portfólio

**Fluxo principal:**

1. Profissional acessa "Meu perfil".
2. Edita bio, categorias, área, disponibilidade e contato.
3. Adiciona trabalho ao portfólio (título, descrição, categoria, imagem opcional).
4. Reordena, edita ou remove trabalhos.
5. Vê prévia e publica alterações.
6. Recebe confirmação e evento de auditoria.

**Resultado:** perfil atualizado, read model de busca consistente.  
**Riscos:** expor PII em imagem/descrição; portfólio de terceiros sem autorização; informação desatualizada.

### 5.4 Descoberta de profissionais pelo cliente

**Fluxo principal:**

1. Cliente informa serviço ou categoria.
2. Informa região aproximada (cidade/bairro; geolocalização opcional).
3. Vê resultados com distância aproximada, sinais de confiança e explicação.
4. Refina por tipo de atendimento, disponibilidade e verificação.
5. Abre perfil e compara.
6. Decide avançar para contato.

**Resultado:** cliente identifica profissional relevante e entende por que ele aparece.  
**Fluxos alternativos:** sem resultados → sugestões e ampliar raio; erro → tentar novamente; localização ausente → pedir cidade/bairro, não assumir.

### 5.5 Conexão direta

**Fluxo principal:**

1. Cliente clica em "Tenho interesse".
2. Informa contexto do serviço e canal de contato preferido.
3. Profissional recebe solicitação, vê perfil resumido e decide.
4. Profissional aceita.
5. Contato é liberado para ambos, conforme preferência.
6. Partes negociam fora da plataforma.
7. Podem registrar desfecho.

**Fluxos alternativos:** recusa, cancelamento, expiração, bloqueio, denúncia.  
**Resultado:** conexão consentida, sem exposição prematura.

### 5.6 Mural de necessidade

**Fluxo principal:**

1. Cliente publica necessidade com categoria, descrição, região aproximada, urgência e orçamento opcional.
2. Profissionais elegíveis veem a oportunidade.
3. Profissional envia interesse com mensagem curta.
4. Cliente compara interessados.
5. Cliente aceita um ou mais.
6. Conexão é registrada e contato liberado.
7. Pedido é fechado ou expira.

**Riscos:** spam de interesses, PII na descrição, pedido desatualizado, poucos interessados.

### 5.7 Descoberta de oportunidades pelo profissional

**Fluxo principal:**

1. Profissional acessa "Oportunidades".
2. Vê pedidos compatíveis com categoria/área.
3. Filtra por urgência, distância e categoria.
4. Abre detalhe.
5. Envia interesse ou retira.
6. Acompanha estado.

**Resultado:** profissional encontra demanda compatível sem expor contato antes do aceite.

### 5.8 Negociação e prestação

**Fluxo principal:**

1. Contato liberado após aceite.
2. Partes discutem condições.
3. Contratam e executam fora da plataforma.
4. Registram desfecho no RAIRE.

**Limite:** o RAIRE não intermedeia pagamento, contrato ou execução no MVP.

### 5.9 Avaliação e reputação

**Fluxo principal:**

1. Conexão é marcada como concluída.
2. Partes recebem convite para avaliar.
3. Avaliam nota, tags contextuais e comentário opcional.
4. Avaliação é publicada com contexto.
5. Avaliado pode responder.
6. Reputação é agregada por categoria e papel.

**Riscos:** retaliação, autoavaliação, avaliação combinada, PII no comentário.

### 5.10 Denúncia e moderação

**Fluxo principal:**

1. Usuário denuncia perfil, pedido, avaliação ou conexão.
2. Escolhe motivo e descreve.
3. Caso entra na fila com prioridade por risco.
4. Operador analisa evidências e histórico.
5. Aplica medida: nenhuma, advertência, remoção, suspensão, bloqueio.
6. Ação é auditada; afetado pode recorrer.

**Resultado:** caso tratado com critério, prazo e rastreabilidade.

---

## 6. Mapa jornada → requisitos

| Jornada | RFs | Spec principal |
| --- | --- | --- |
| 5.1 Conta | RF-01, RF-02, RF-19 | foundation-platform |
| 5.2 Perfil | RF-03, RF-05, RF-15 | foundation-platform |
| 5.3 Portfólio | RF-04 | foundation-platform |
| 5.4 Busca | RF-06, RF-07 | local-discovery |
| 5.5 Conexão direta | RF-11, RF-12, RF-16 | service-request |
| 5.6 Mural | RF-08, RF-09, RF-10 | service-request |
| 5.7 Oportunidades | RF-09, RF-10 | service-request |
| 5.8 Negociação | RF-12, RF-13 | service-request + trust-reputation |
| 5.9 Avaliação | RF-13, RF-14 | trust-reputation |
| 5.10 Moderação | RF-17, RF-18 | trust-reputation |

---

## 7. Critérios de validação das personas

- Entrevistar pelo menos 5 clientes e 5 profissionais por categoria piloto.
- Observar a jornada real, não só a opinião declarada.
- Testar onboarding com 3 profissionais de baixa familiaridade digital.
- Testar busca e comparação com 3 clientes em celular.
- Registrar evidência em `docs/research` e `docs/experiments`.
- Revisar personas após o primeiro ciclo do piloto.

> Personas são hipóteses até evidência de campo. Não usar persona para justificar requisito sem evidência.
