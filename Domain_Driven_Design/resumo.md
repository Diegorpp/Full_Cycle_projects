# DDD - Domain Driven Design

## Dominio e Subdominio
- Dominio: Escopo Macro de complexidade que é muito dificil se enxerga como unidade, logo são subdividos em pedaços menores.
    - Core Domain: Seria o domínio principal, ou seja, não faz sentido sequer ter uma aplicação se não existir esse cara, com isso é possível notar que ele está diretamente ligado as regras de negocio. (Engloba o diferencial competitivo)

- Subdominio: Escopo menor dentro do dominio onde você atribui um pedaço do software que é mais facil de agrupar.
    - Support subdomain: apoia o domínio no dia a dia, ou seja ajuda na operação do domínio.
    - Generic subdomain: São os softwares auxiliares, são recursos que podem ser trocados sem problemas. Não possui diferencial competitivo.


## Problema e Solução

    ...

## Contexto é Rei

Nomes iguais com perspectivas diferentes demonstram que esses elementos estão em contextos diferentes.

Por exemplo: É possível que hajam varias classes de Cliente de acordo com a pespectiva de contexto, por exemplo uma classe "Usuário" pode ser criado para o contexto de suporte, contexto de atendimento ao cliente, etc. Não faz sentido tentar transformar a classe cliente para englobar tudo possível e imaginavel de um sistema, pois ao tentar separar isso futuramente para tornar mais escalável irá gerar MUITA refatoração.

## Context Mapping

