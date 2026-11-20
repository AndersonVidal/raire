# Segurança e Privacidade — Processo

**Status:** proposta para revisão humana e jurídica  
**Versão:** 1.0  
**Relacionado:** `docs/security/privacy-and-lgpd.md`, `docs/security/threat-model.md`, `docs/process/review-criteria.md`

---

## 1. Objetivo

Proteger usuários, dados pessoais, localização e a integridade do marketplace. O RAIRE lida com dados que podem expor fisicamente uma pessoa e com dinâmica de reputação que pode causar dano social. Segurança e privacidade são requisitos de produto.

---

## 2. Modelo de ameaças (resumo)

| Ameaça | Ator | Impacto | Controles |
| --- | --- | --- | --- |
| Perfil falso / catfishing | fraudador | confiança, dano financeiro | verificação manual, sinais, denúncia |
| Colheita de contatos / scraping | spammer | assédio, vazamento | rate limit, mascaramento, detecção de enumeração |
| Assédio via contato | usuário malicioso | dano psicológico | consentimento, bloqueio, denúncia, canal seguro |
| Vazamento de localização precisa | stalker/atacante | risco físico | localização aproximada, criptografia, acesso restrito |
| Manipulação de avaliações | profissional/concorrente | reputação distorcida | elegibilidade por conexão, detecção, moderação |
| Spam de pedidos/interesses | spammer | experiência ruim | limites, validação, triagem |
| Account takeover | atacante | fraude, dano | senha forte, rate limit, sessão revogável, MFA admin |
| Abuso de admin | operador | injustiça, vazamento | menor privilégio, MFA, auditoria, dupla aprovação |
| Ranking gaming | profissional | fairness quebrada | caps, logs, detecção de anomalia, revisão |
| Discriminação algorítmica | sistema | injustiça, jurídico | proibir atributos sensíveis, coortes observáveis, auditoria |
| PII em logs/analytics | erro interno | vazamento | redaction, revisão, testes, retenção |
| Dependência vulnerável | atacante | comprometimento | scan, atualização, SBOM |
| Injeção / XSS / CSRF | atacante | comprometimento | validação, encoding, CSP, CSRF token |

Detalhes em `docs/security/threat-model.md`.

---

## 3. LGPD — princípios operacionais

### 3.1 Bases legais por finalidade

| Finalidade | Base legal proposta | Observação |
| --- | --- | --- |
| Criar conta e autenticar | execução de contrato | necessário para o serviço |
| Exibir perfil público | consentimento/execução de contrato | usuário escolhe o que publicar |
| Localização aproximada | execução de contrato | minimização |
| Localização precisa | consentimento | acesso restrito e retenção curta |
| Contato | consentimento | liberado após aceite |
| Avaliações | execução de contrato/interesse legítimo | com direito de resposta |
| Moderação e segurança | legítimo interesse/obrigação legal | registro de auditoria |
| Analytics/pesquisa | consentimento | dados pseudonimizados |
| Comunicação de marketing | consentimento | opt-in separado |

> A validação final é jurídica. Nenhuma base legal é definitiva sem revisão.

### 3.2 Direitos do titular

- confirmação de tratamento;
- acesso;
- correção;
- anonimização/bloqueio/eliminação;
- portabilidade;
- informação sobre compartilhamento;
- revogação de consentimento;
- oposição quando aplicável.

Fluxo operacional:

1. canal de solicitação;
2. verificação de identidade;
3. atendimento em prazo legal;
4. registro da solicitação e resposta;
5. exclusão/anonimização com ressalvas legais.

### 3.3 Retenção

Proposta inicial (validar com jurídico):

- conta ativa: enquanto durar a relação;
- conta excluída: anonimização em até 30 dias;
- localização precisa: até 90 dias;
- eventos analíticos: 13 meses pseudonimizados;
- denúncias/auditoria: 5 anos com acesso restrito;
- documentos de verificação: 90 dias após decisão;
- logs de aplicação: 30 a 90 dias, sem PII.

### 3.4 Compartilhamento

- provedores de e-mail transacional;
- infraestrutura/hospedagem;
- ferramentas de observabilidade (sem PII);
- nenhuma venda de dados;
- contratos e due diligence com operadores;
- transferência internacional avaliada.

### 3.5 Incidentes

1. Detectar e conter.
2. Avaliar risco e dados afetados.
3. Registrar evidências.
4. Comunicar encarregado/direção.
5. Notificar titulares e ANPD quando exigido.
6. Corrigir causa raiz.
7. Postmortem sem culpa.

---

## 4. Segurança de aplicação

### 4.1 Autenticação

- hash de senha com Argon2id (ou bcrypt com custo adequado);
- política de senha moderna (comprimento, bloqueio de senhas comuns);
- rate limit por conta e IP;
- sessão com cookie `HttpOnly`, `Secure`, `SameSite`;
- rotação de sessão após login;
- revogação em logout e troca de senha;
- MFA obrigatório para admin;
- tokens de reset de uso único e curta duração, sem revelar existência da conta.

### 4.2 Autorização

- RBAC + ownership;
- toda operação verifica ator e recurso;
- admin não acessa dados sem necessidade e com auditoria;
- testes de autorização cruzada;
- negar por padrão.

### 4.3 Entrada e saída

- validação server-side;
- normalização de texto e limites;
- query parametrizada;
- output encoding;
- CSP restritiva;
- upload futuro: validação de tipo, tamanho, antivírus e storage isolado.

### 4.4 Rate limiting

| Ação | Limite inicial proposto |
| --- | --- |
| Login | 5 tentativas/min por conta + IP |
| Cadastro | 3/h por IP |
| Busca | 60/min por sessão |
| Pedido | 5/dia por cliente |
| Interesse | 20/dia por profissional |
| Denúncia | 10/dia por usuário |
| Contato | 20/dia por conexão |
| Admin | auditoria e limite por operador |

Valores a calibrar com dados reais e sem penalizar usuários legítimos.

### 4.5 Segredos e dependências

- secret manager ou variáveis de ambiente seguras;
- nunca commitar `.env`;
- scan de segredos no CI;
- SBOM e scan de dependências;
- atualização com prazo para críticas;
- assinatura de imagens/containers quando possível.

### 4.6 Logs e observabilidade

- JSON estruturado;
- `correlation_id`;
- nunca logar senha, token, documento, telefone, e-mail completo, localização precisa ou conteúdo de denúncia;
- redaction testada;
- acesso a logs restrito e auditado.

---

## 5. Segurança de produto e trust & safety

### 5.1 Contato

- contato mascarado até aceite;
- preferência de canal explícita;
- opção de não receber contato;
- bloqueio e denúncia em todas as etapas;
- registro de consentimento com timestamp e versão.

### 5.2 Localização

- público: cidade/bairro ou raio aproximado;
- preciso: criptografado, acesso restrito, retenção curta;
- nunca exibir endereço residencial;
- arredondamento/ruído para área pública;
- teste de que a API pública não retorna coordenada precisa.

### 5.3 Avaliações

- só após conexão confirmada;
- uma por lado por conexão;
- detecção de autoavaliação e redes de avaliação;
- direito de resposta;
- moderação com critério e auditoria;
- não remover avaliação negativa legítima.

### 5.4 Moderação

- critérios públicos;
- fila priorizada por risco;
- prazo inicial de 48 h;
- ações auditáveis;
- recurso e revisão humana;
- proteção contra abuso de denúncia.

---

## 6. Fairness como controle de segurança

Fairness não é só ética: ranking opaco e discriminatório é risco jurídico e reputacional.

Controles:

- nenhum atributo sensível em features;
- coortes observáveis e agregadas;
- política versionada;
- simulação e guardrails;
- logs de decisão;
- revisão humana;
- canal de contestação;
- monitoramento contínuo.

---

## 7. Checklist de review de segurança/privacidade

- [ ] Autenticação e sessão seguras.
- [ ] Autorização por papel e recurso.
- [ ] Rate limit nas ações sensíveis.
- [ ] Entrada validada.
- [ ] Saída encoded.
- [ ] CSRF/CSP quando aplicável.
- [ ] Sem segredo no código.
- [ ] Sem PII em log/evento/analytics.
- [ ] Localização precisa protegida.
- [ ] Contato consentido e mascarado.
- [ ] Retenção e exclusão definidas.
- [ ] Dependências sem vulnerabilidade crítica/alta.
- [ ] Admin com MFA e auditoria.
- [ ] Backup e restore.
- [ ] Plano de incidente.
- [ ] DPIA quando aplicável.

---

## 8. Bloqueios absolutos

- vazamento de dado de outro usuário;
- exposição de localização precisa ou contato sem consentimento;
- bypass de autenticação/autorização;
- segredo commitado;
- injeção explorável;
- PII em log/evento;
- ranking com atributo sensível;
- remoção silenciosa de avaliação legítima;
- ação administrativa sem auditoria.
