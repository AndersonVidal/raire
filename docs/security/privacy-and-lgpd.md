# Privacidade e LGPD — RAIRE

**Status:** proposta para revisão jurídica  
**Versão:** 1.1  
**Dono:** Felipe  
**Aprovador:** jurídico/privacidade + fundador  
**Relacionado:** `docs/product/prd.md` (seção 10), `docs/process/security-privacy.md`  
**Decisão de planejamento:** Q-38–Q-43 aprovadas em 2026-11-20; validação jurídica e DPIA continuam obrigatórias antes do piloto (`docs/product/decisions-2026-11-20.md`).

> Bases legais, prazos e textos são propostas. Nada aqui substitui revisão jurídica antes do piloto.

---

## 1. Princípios

- **Minimização:** coletar apenas o necessário para a jornada.
- **Finalidade:** cada dado tem uso declarado.
- **Consentimento:** contato, pesquisa e marketing têm consentimento separado.
- **Segurança:** criptografia em trânsito e repouso; PII separada de dados públicos.
- **Transparência:** o usuário entende o que é público, o que é privado e por quê.
- **Retenção:** prazos definidos e exclusão/anonimização.
- **Responsabilização:** registro de decisões e incidentes.
- **Privacy by design/default:** a jornada mais fácil é a mais privada.

---

## 2. Inventário de dados

| Dado | Finalidade | Base legal proposta | Classificação | Retenção |
| --- | --- | --- | --- | --- |
| e-mail | conta, login, contato | execução de contrato | pessoal | enquanto a conta existir |
| senha (hash) | autenticação | execução de contrato | sensível | enquanto a conta existir |
| nome de exibição | perfil público | consentimento/execução | pública | enquanto publicado |
| bio | perfil público | consentimento/execução | pública | enquanto publicado |
| categorias | descoberta | execução de contrato | pública | enquanto publicado |
| cidade/bairro/raio | descoberta | execução de contrato | pública (aproximada) | enquanto publicado |
| localização precisa | elegibilidade/distância | consentimento | pessoal sensível | 90 dias |
| telefone | contato | consentimento | pessoal | enquanto a conta existir |
| portfólio | apresentação | consentimento | pública | até remoção |
| pedidos | conexão | execução de contrato | interna/pública parcial | 12 meses |
| interesses | conexão | execução de contrato | interna | 12 meses |
| conexões | histórico | execução de contrato | interna | 24 meses |
| avaliações | reputação | execução de contrato/interesse legítimo | pública | 24 meses |
| denúncias | segurança | legítimo interesse/obrigação | sensível | 5 anos |
| auditoria | segurança/compliance | obrigação legal/legítimo interesse | restrita | 5 anos |
| eventos analíticos | produto/fairness | consentimento | pseudonimizada | 13 meses |
| logs | operação | legítimo interesse | sem PII | 30–90 dias |
| documentos de verificação | confiança | consentimento | sensível | 90 dias após decisão |

---

## 3. Bases legais por finalidade

| Finalidade | Base legal proposta | Observação |
| --- | --- | --- |
| Criar conta e autenticar | execução de contrato | necessário para o serviço |
| Exibir perfil público | consentimento/execução | usuário escolhe o que publicar |
| Localização aproximada | execução de contrato | minimização |
| Localização precisa | consentimento | acesso restrito e retenção curta |
| Contato | consentimento | liberado após aceite |
| Avaliações | execução de contrato/interesse legítimo | direito de resposta |
| Moderação e segurança | legítimo interesse/obrigação legal | auditoria |
| Analytics/pesquisa | consentimento | dados pseudonimizados |
| Marketing | consentimento | opt-in separado |

**Validar com jurídico antes do piloto.**

---

## 4. Direitos do titular

| Direito | Como atender | Prazo |
| --- | --- | --- |
| confirmação de tratamento | página de privacidade + solicitação | 15 dias |
| acesso | exportação JSON | 15 dias |
| correção | edição de perfil + solicitação | 15 dias |
| anonimização/bloqueio/eliminação | fluxo de exclusão | 15 dias |
| portabilidade | exportação JSON | 15 dias |
| informação sobre compartilhamento | política + resposta | 15 dias |
| revogação de consentimento | página de consentimentos | imediato |
| oposição | análise | 15 dias |

### 4.1 Fluxo operacional

1. canal de solicitação;
2. verificação de identidade;
3. registro em `account_requests`;
4. atendimento no prazo;
5. resposta e encerramento;
6. exclusão/anonimização com ressalvas legais.

---

## 5. Consentimento

- Termos e privacidade aceitos no cadastro, com versão.
- Pesquisa/analytics com consentimento separado.
- Marketing com opt-in separado.
- Contato liberado somente após aceite da conexão.
- Revogação disponível a qualquer momento.
- Registro de timestamp, versão e canal.

---

## 6. Localização

- Público: cidade/bairro e faixa de distância.
- Preciso: cifrado, acesso restrito, retenção curta.
- Nunca exibir endereço residencial.
- Nunca retornar coordenada precisa em API pública, evento ou log.
- Arredondar/ruído para área pública.
- Teste de contrato obrigatório.

---

## 7. Portfólio e conteúdo de terceiros

- Profissional declara ter autorização para publicar imagens.
- Conteúdo denunciável.
- Remoção por violação de direitos, com auditoria.
- Sem reconhecimento facial ou biometria.
- Metadados de imagem removidos no upload (quando houver upload).

---

## 8. Compartilhamento e fornecedores

- Provedores de e-mail transacional.
- Infraestrutura/hospedagem.
- Observabilidade sem PII.
- Nenhuma venda de dados.
- Contratos e due diligence com operadores.
- Transferência internacional avaliada por fornecedor.
- Lista de subprocessadores pública.

---

## 9. Retenção e descarte

| Dado | Prazo | Ação |
| --- | --- | --- |
| conta ativa | relação | manter |
| conta excluída | 30 dias | anonimizar |
| localização precisa | 90 dias | expurgar |
| pedidos/interesses | 12 meses | anonimizar/agregar |
| conexões | 24 meses | anonimizar |
| avaliações | 24 meses | manter pública com autor pseudonimizado |
| denúncias/auditoria | 5 anos | manter restrito |
| documentos | 90 dias após decisão | excluir |
| eventos | 13 meses | agregar/expurgar |
| logs | 30–90 dias | expurgar |

Jobs de expurgo devem ser auditáveis e testados.

---

## 10. Segurança do tratamento

- Argon2id para senha;
- TLS em trânsito;
- criptografia em repouso para PII sensível;
- acesso por menor privilégio;
- MFA para admin;
- logs sem PII;
- backup cifrado;
- teste de restore;
- auditoria append-only.

---

## 11. Incidentes de privacidade

1. Detectar e conter.
2. Avaliar risco e dados afetados.
3. Registrar evidências.
4. Comunicar encarregado/direção.
5. Notificar titulares e ANPD quando exigido.
6. Corrigir causa raiz.
7. Postmortem sem culpa.
8. Atualizar controles.

**Prazo interno para avaliação inicial:** 24 h.  
**Comunicação externa:** conforme lei e orientação jurídica.

---

## 12. DPIA / RIPD

Antes do piloto, avaliar necessidade de RIPD considerando:

- localização e contato;
- moderação e denúncias;
- reputação e potencial dano social;
- ranking e fairness;
- dados de vulneráveis;
- escala e retenção.

A DPIA deve registrar:

- fluxos de dados;
- bases legais;
- riscos aos titulares;
- medidas de mitigação;
- riscos residuais;
- aprovação.

---

## 13. Governança

| Papel | Responsabilidade |
| --- | --- |
| Fundador | aprovar política e liberar piloto |
| Felipe | operar privacidade, revisar incidentes, manter registro |
| Jurídico/DPO | validar bases legais, retenção e DPIA |
| Victor | implementar controles técnicos |
| Anderson | garantir transparência e copy compreensível |

Revisão trimestral da política e dos processos.

---

## 14. Critérios mínimos para o piloto

- [ ] política de privacidade publicada e versionada;
- [ ] termos de uso publicados;
- [ ] bases legais revisadas;
- [ ] DPIA concluída ou justificada;
- [ ] retenção e expurgo definidos;
- [ ] direitos do titular operacionais;
- [ ] contato e localização protegidos;
- [ ] plano de incidente;
- [ ] DPO/encarregado definido;
- [ ] fornecedores mapeados;
- [ ] treinamento da operação.
