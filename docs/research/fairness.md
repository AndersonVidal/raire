# Fairness — Nota de Pesquisa

**Status:** referência inicial  
**Versão:** 1.0  
**Uso:** orientar desenho de ranking e experimentos.

---

## 1. Definições úteis

- **Fairness de exposição:** distribuição de visibilidade entre profissionais comparáveis.
- **Fairness de oportunidade:** acesso a conexões, não apenas impressões.
- **Fairness procedimental:** critérios consistentes, transparentes e contestáveis.
- **Fairness de resultado:** qualidade das conexões para ambos os lados.

O RAIRE prioriza exposição + oportunidade + procedimental, sem prometer igualdade de resultado.

---

## 2. Riscos

- concentração de visibilidade;
- feedback loop de popularidade;
- punição de novos;
- proxy de atributo sensível;
- gaming;
- perda de relevância;
- opacidade.

---

## 3. Métricas

- Gini de impressões;
- share top 10/20;
- share de novos;
- posição média por coorte;
- conversão por coorte;
- cobertura;
- diversidade de bairro/categoria;
- satisfação e confiança.

---

## 4. Mecanismos

- caps;
- exploração;
- boost de novo;
- diversidade;
- piso de relevância;
- explicação;
- fallback;
- política versionada.

---

## 5. Referências conceituais

- fair ranking / exposure fairness;
- multi-sided fairness;
- market design;
- reputation systems;
- algorithmic auditing;
- accessibility.

> Referências formais devem ser adicionadas em `references/` com link e resumo. Não citar paper sem leitura.

---

## 6. Perguntas abertas

- Qual o peso justo de reputação sem punir novos?
- Como medir oportunidade além de impressão?
- Qual guardrail de relevância é aceitável?
- Como evitar gaming sem prejudicar legítimos?
- Como comunicar fairness sem facilitar manipulação?
