# Catálogo Inicial — Proposta

**Status:** proposta para revisão  
**Versão:** 1.0  
**Dono:** Anderson  
**Apoio:** Victor (dados), Felipe (moderação)  
**Decisão:** Q-02, Q-11

> Seed sintético para o MVP técnico. Categorias reais do piloto dependem de validação de campo.

---

## 1. Categorias propostas

| Categoria | Slug | Tipo | Descrição | Sinônimos |
| --- | --- | --- | --- | --- |
| Manutenção e reparos | `manutencao-reparos` | presencial | pequenos reparos e instalações | reparo, conserto, manutenção |
| Eletricista | `eletricista` | presencial | instalações e reparos elétricos | instalação elétrica, elétrica, tomada |
| Encanador | `encanador` | presencial | vazamentos e instalações hidráulicas | hidráulica, vazamento, encanamento |
| Montador de móveis | `montador-moveis` | presencial | montagem e desmontagem | montagem, móveis |
| Aulas e consultoria | `aulas-consultoria` | remoto/híbrido | aulas e consultoria | aula, consultoria, mentoria |
| Professor particular | `professor-particular` | remoto | aulas de reforço | reforço, aula particular |
| Consultoria | `consultoria` | remoto | consultoria profissional | mentoria, assessoria |
| Personal trainer | `personal-trainer` | híbrido | treino personalizado | treino, academia |

---

## 2. Bairros sintéticos

| Cidade | Bairro | Coordenada aproximada |
| --- | --- | --- |
| Cidade Piloto | Centro | -30.0000, -51.0000 |
| Cidade Piloto | Jardim Norte | -29.9900, -51.0100 |
| Cidade Piloto | Vila Sul | -30.0100, -51.0200 |
| Cidade Piloto | Bairro Leste | -30.0050, -50.9900 |

> Coordenadas fictícias para teste; não representam endereço real.

---

## 3. Regras

- categoria precisa estar ativa para aparecer em onboarding/busca;
- sinônimos normalizados (lowercase, sem acento);
- desativação preserva histórico;
- usuário não cria categoria livre;
- sugestão de categoria é could;
- catálogo é revisado por produto antes do piloto.

---

## 4. Critérios de aceite do seed

- [ ] idempotente;
- [ ] slugs únicos;
- [ ] sinônimos normalizados;
- [ ] categoria inativa não aparece;
- [ ] revisão de produto;
- [ ] sem dado real.
