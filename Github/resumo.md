Git Flow
Configurações dos nossos branches
PR / Templates para PR
Code Review
Plugin do VSCode
SemVer - Semantical Versioning - Como versionar o código [Major, Minor, Patch]
Conventional Commits
CODEEOWNERS

## GitFlow

O gitflow tem como objetivo definir boas praticas de gerenciamento de repositório. Basicamente
temos como fluxo o **Master** que tem como objetivo tem sempre a versão ESTÁVEL, ou seja, ela é sempre
a versão que vai estar rodando em produção.

- Develop: Branch responsável por contentrar a features em desenvolvimento/teste

- Feature: Branch utilizada para implementar cada feature

- Release: Utilizado para concentrar as funcionalidades que forem ficando prontas antes de ir para a master. Aqui é solicitado para que o setor de teste possa entrar.

- Hotfix: branch separa para fazer alterações imediatas de produção, ou seja, a casa ta caindo e vc precisa resolver.

[Installing git flow](https://github.com/petervanderdoes/gitflow-avh/wiki/Installation)


### Segurança

Para evitar que outras pessoas commitem utilizando assinaturas iguais a sua. Para evitar isso é importante criar uma assinatura digital, atribuindo um determinado padrão de segurança.

- Assinatura

É necessario possuir uma chave publica e uma chave privada.

#### GNU private guard - gpg

Projeto open source que possui diversas funcionalidades de segurança e criação de chaves.

- Lista todos as chaves que ja foram criadas no sistema utilizando o gpg
```
    gpg --list-secret-keys --keyid-format LONG
```
- Comando para iniciar a criação do tipo de chave
```
    gpg --full-generate-key
    Tipode chave: 1 (RSA = 1 = Padrão)
    Numero de bits: 4096
    Tempo de expiração: 1y
    Aceitar os termos anteriores: y
    usuario: Usuario como ficará registrado seus commits
    email: email a ser utilizado
    senha: utilizado para acessar a chave.
```
- Obtendo chave publica do gpg
```
    gpg --armor --export <KEY_ID>
```

Após gerar a chave publica, ela deve ser inserida no campo gpg dentro das configurações do github.

Adicionar a chave nas configuração do seu git local.

```
    git config --global user.signingkey <KEY_ID>
```

Inserir o comando a baixo no seu .bashprofile
```
    export GPG_TTY=$(tty)
```

#### Configuração opcinal

Para assinar por padrão com a chave gpg sempre que for interessante, utilize:

O parâmetro "--global" indica que irá assinar assim para todos os repositorios, se quiser apenas no repositório atual, remova esse parâmetro.
```
    git config --global commit.gpgsign true
    git config --global commit.tagSign true
```

## Boas praticas

Mudar em configuração a branch para develop, dessa forma evita que commits sejam feitos na master caso alguem clonou e por alguma razão não alterou a branch.

É possível criar novas regras de proteção de branch.
```
    feature/*
```
Regra se aplica a todas as branches que começam com feature/

É possível exigit que os commits sejam assinados, é possível restringir usuarios ou grupos para poder commitar em uma determinada branch

O Ideal é proteger os projetos de commits diretos na develop e na master, dessa forma a unica possibilidade de alteração seria utilizando.

O processo de proteção de branches que restringe quem poderá realizar o push é uma funcionalidade que está disponível apenas para repositórios associados a uma organização no Github (o que é o mais comum quando trabalhamos com github em uma empresa).

Logo, se você não estiver vendo a opção: "Restrict who can push to matching branches" e quiser testar o recurso, crie uma organização no Github e crie um repositório associado a ela.



