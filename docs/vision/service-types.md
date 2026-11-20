# Tipos de Serviço — RAIRE

**Status:** referência de domínio  
**Versão:** 1.0

---

## 1. Por localidade

| Tipo | Definição | Exige área? | Distância no ranking? | Exemplo |
| --- | --- | --- | --- | --- |
| Presencial (`in_person`) | execução no local do cliente | sim | sim | eletricista, encanador, diarista |
| Remoto (`remote`) | execução à distância | não | não | aula online, consultoria, design |
| Híbrido (`hybrid`) | parte remota, parte presencial | sim | sim | arquiteto, personal trainer |

**Regra:** perfil presencial/híbrido sem área não publica. Remoto sem área é permitido.

---

## 2. Por relação de trabalho

| Tipo | Definição | Observação |
| --- | --- | --- |
| Autônomo individual | pessoa física | foco do MVP |
| Estúdio/pequeno negócio | equipe pequena | pode usar perfil de profissional; multi-perfil fora do MVP |
| Organização | empresa | fora do MVP |

---

## 3. Por urgência

| Tipo | Definição | Efeito no produto |
| --- | --- | --- |
| Imediato | hoje/24h | destaque de urgência no pedido; sem urgência falsa |
| Semana | próximos dias | padrão |
| Flexível | sem prazo rígido | planejamento |

---

## 4. Por complexidade

| Tipo | Definição | Implicação |
| --- | --- | --- |
| Tarefa simples | uma visita, escopo claro | conexão direta |
| Projeto | múltiplas etapas | fora do MVP (sem gestão de projeto) |
| Recorrente | periodicidade | fora do MVP (sem agenda) |

---

## 5. Categorias iniciais propostas

| Categoria | Tipo | Exemplos | Por que |
| --- | --- | --- | --- |
| Manutenção e reparos | presencial | eletricista, encanador, montador | demanda frequente, distância importa |
| Aulas e consultoria | remoto/híbrido | professor particular, consultor | valida serviço remoto e agenda flexível |

> Categorias finais dependem de Q-01/Q-02 e validação de campo.

---

## 6. O que não é serviço no MVP

- produto físico;
- pagamento;
- contrato;
- obra de grande porte;
- serviço que exige verificação regulatória específica;
- emergência médica/jurídica.

---

## 7. Implicações de produto

- tipo de serviço afeta busca, ranking e formulário;
- remoto não deve ser penalizado por distância;
- híbrido precisa de área, mas pode ter parte remota;
- urgência não pode virar dark pattern;
- categorias reguladas exigem análise jurídica antes de entrar.
