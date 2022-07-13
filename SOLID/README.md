# SOLID

SOLID é um conjunto de pilares para melhoras a forma de codificar.

Criado por: (Uncle Bob)Robert C. Martin no livro Agile software development.

## S - Single Responsability Principle

Uma classe deve ter uma única responsabilidade. Ela não deve fazer coisa a mais do que ela foi projetada para fazer.

Uma classe deve ter um e apenas um motivo para mudar a classe. (caso ela não seja cumprida é provavel que ela esteja com mais atribuições do que deveria).



## O - Open/Closed Principle

Toda classe deve ser aberta para extensão e fechada para modificação.

Criar uma base da classe de interesse e apartir do momento que você queria adicionar algum comportamento, você deve extender essa classe e fazer a nova alteração a partir de um outro componente e não da classe principal.

Um exemplo seria criar um interface para calcular interesse e para cada novo interesse é criado uma nova classe para aquele respectivo interesse. A classe deve extender a classe abstrata e implementar seus métodos da forma que for conveniênte com a regra de negócio especificada.

## L - Liskov Substitution Principle

Subclasses podem ser substituidas por suas classes pai.

Ou seja todos os métodos da classe pai devem poder ser executados a partir da classe filho. Dito isso, caso algum método não seja possível de ser executado a partir da classe filho significa que essa classe não deveria ter sido extendida dessa classe.

## I - Interface segregation Principle

## D - Dependency inversion Principle

