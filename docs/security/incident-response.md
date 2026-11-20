# Resposta a Incidentes — RAIRE

**Status:** proposta  
**Versão:** 1.0  
**Dono:** Felipe  
**Aprovador:** fundador + segurança/privacidade

---

## 1. Objetivo

Detectar, conter, corrigir e aprender com incidentes de segurança, privacidade, abuso ou disponibilidade, preservando evidências e cumprindo prazos legais.

---

## 2. Classificação

| Severidade | Exemplos | Prazo de resposta |
| --- | --- | --- |
| S0 — crítico | vazamento de PII/localização, bypass de auth, fraude generalizada | imediato |
| S1 — alto | acesso indevido a recurso, manipulação de avaliação em escala, indisponibilidade | 1 h |
| S2 — médio | spam, abuso contido, falha de notificação | 4 h |
| S3 — baixo | bug sem exposição, ruído operacional | 1 dia |

---

## 3. Fluxo

```txt
1. Detectar e registrar
2. Conter
3. Avaliar escopo e risco
4. Preservar evidências
5. Erradicar causa
6. Recuperar
7. Comunicar (interno/externo)
8. Postmortem sem culpa
9. Atualizar controles
```

### 3.1 Detectar

- alerta de monitoramento;
- denúncia;
- revisão de segurança;
- relato de usuário;
- auditoria.

Registrar: data/hora UTC, detector, sistema afetado, sintoma, correlação.

### 3.2 Conter

- desligar feature flag;
- revogar sessões;
- suspender conta/perfil;
- bloquear IP/rota;
- isolar serviço;
- preservar logs.

### 3.3 Avaliar

- quais dados?
- quantos titulares?
- há risco físico?
- há exposição pública?
- há obrigação legal de notificar?
- qual a causa provável?

### 3.4 Preservar evidência

- logs;
- snapshots;
- auditoria;
- hashes;
- cadeia de custódia;
- acesso restrito.

### 3.5 Erradicar e recuperar

- corrigir causa;
- testar;
- restaurar backup se necessário;
- monitorar recorrência;
- validar dados.

### 3.6 Comunicar

- interno: fundador, segurança, produto;
- externo: titulares/ANPD quando exigido;
- imprensa: somente fundador;
- usuários: linguagem clara, sem detalhe que aumente risco.

### 3.7 Postmortem

- linha do tempo;
- causa raiz;
- impacto;
- o que funcionou;
- o que falhou;
- ações com dono e prazo;
- atualização de risco/controle.

---

## 4. Papéis

| Papel | Responsável |
| --- | --- |
| Incident commander | Felipe |
| Comunicação | Fundador |
| Produto | Anderson |
| Engenharia | Victor |
| Jurídico/DPO | a designar |
| Aprovação de comunicação | Fundador + jurídico |

---

## 5. Runbooks rápidos

### 5.1 Suspeita de vazamento de contato/localização

1. desligar exposição/feature;
2. revogar sessões envolvidas;
3. preservar logs e auditoria;
4. identificar escopo;
5. corrigir e testar;
6. comunicar conforme risco;
7. postmortem.

### 5.2 Account takeover

1. revogar sessões;
2. forçar reset;
3. verificar MFA;
4. revisar logs;
5. comunicar;
6. investigar causa.

### 5.3 Ranking injusto / atributo sensível

1. desligar fairness;
2. voltar política anterior;
3. preservar ledger;
4. revisar features;
5. comunicar;
6. postmortem.

### 5.4 Fraude de avaliação

1. sinalizar avaliações;
2. suspender contas envolvidas;
3. preservar evidência;
4. revisar reputação;
5. corrigir;
6. comunicar.

### 5.5 Indisponibilidade

1. verificar health/readiness;
2. verificar banco/dependências;
3. rollback se necessário;
4. comunicar status;
5. restaurar;
6. postmortem.

---

## 6. Prazos legais

- Avaliação inicial de incidente de privacidade: 24 h.
- Comunicação à ANPD/titulares: conforme LGPD e orientação jurídica.
- Registro de decisão e evidência: imediato.
- Revisão pós-incidente: 7 dias.

---

## 7. Prevenção

- revisão de segurança em todo PR;
- rate limit;
- MFA admin;
- auditoria;
- backups testados;
- alertas;
- treinamento;
- simulação de incidente antes do piloto.

---

## 8. Critérios de encerramento

- [ ] causa contida;
- [ ] dados corrigidos/restaurados;
- [ ] comunicação feita;
- [ ] evidência preservada;
- [ ] ações com dono e prazo;
- [ ] risco atualizado;
- [ ] aprovação humana.
