# Sinais de Ranking — Nota de Pesquisa

**Status:** referência inicial  
**Versão:** 1.0

---

## 1. Sinais permitidos

| Sinal | Fonte | Uso | Risco |
| --- | --- | --- | --- |
| correspondência de categoria | catálogo + perfil | relevância | catálogo ruim |
| termo/biografia | perfil | relevância | gaming de keyword |
| distância/área | perfil + busca | relevância presencial | localização imprecisa |
| disponibilidade | perfil | relevância | desatualização |
| verificação | identity/contato | confiança | promessa excessiva |
| reputação contextual | avaliações | confiança | manipulação |
| completude | perfil | qualidade | punir P3 |
| atividade | conexões | confiança | feedback loop |
| novidade | data de entrada | fairness | gaming de recadastro |

## 2. Sinais proibidos

- raça, gênero, idade, deficiência, religião, orientação sexual;
- inferência desses atributos;
- proxies como CEP, nome, escola, foto;
- posição paga;
- dados de outros usuários;
- localização precisa.

## 3. Pesos iniciais

```txt
category_match 0.40
text_match     0.20
proximity      0.15
reputation     0.10
availability   0.05
verification   0.05
completeness   0.05
```

## 4. Riscos de gaming

- keyword stuffing;
- recadastro para parecer novo;
- avaliações combinadas;
- cliques artificiais;
- perfis duplicados.

Mitigações: normalização, detecção de padrão, caps, auditoria, revisão.

## 5. Perguntas abertas

- Reputação deve ter peso maior?
- Como medir atividade sem feedback loop?
- Como tratar disponibilidade desatualizada?
- Qual o peso de completude sem punir P3?
