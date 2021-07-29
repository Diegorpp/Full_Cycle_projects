# Kubernets - K8s

Orquestrador/Gerenciador de containers.

- Disponibilizado através de conjunto de APIs
- Normalmente acessamos a API usano a CLI: kubectl
- Tudo é baseado em estado. Você configura o estado de cada objeto
- Kubernetes Master
    - Kube-apiserver
    - Kube-controller-manager
    - Kube-scheduler
- Outros Nodes:
    - Kubelet
    - Kubeproxy

Dentro do Nó do cluste existe os Pods.
Pod representa os processos rodando num cluster.
É possível ter mais um container dentro de um pod, a não ser que os containers estejam muito bem amarrados para justificar isso.

### Replicaset

Estrutura de descreve o numero de pods que devem ser mantidos.


## Kind

Existem outras ferramentas para rodar o kubernets
- Minikube - Ferramenta para rodar localmente kubernets

Ferramenta que simula fielmente o kubernets.
Possibilita configuração de nodes e não exige maquinas virtuais.

[Kind User Guide](https://kind.sigs.k8s.io/docs/user/quick-start#installation)

Para rodar o kind é importante ter tanto o docker para que funcione as coisas, quanto o kubectl para que seja possível gerenciar os recursos.

### Principais comandos
- create
- delete
- get

    kind create cluster - 
    kind get clusters - lista os clusters
    kind delete clusters <cluster_name> - Deleta um cluster

Por padrão o kind cria apenas um nó.

É imprecindível para manipular o container vincular o kubectl a um determinado cluster, no final do processo é informado o comando necessário para vincular ele ao cluster.

    kubectl cluster-info --context kind-kind

Criar um arquivo kind.yaml onde será feita a configuração de como ficara o cluster.

### Criando cluster via YAML

É possível declarar o cluster utilizando um arquivo de manifesto, segue abaixo um exemplo de especificação para um cluster com 1 master e 3 workers.

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4 #Fixo do kind

nodes: # Vai possuir 4 nós
- role: control-plane #Cara que vai gerenciar o kubernets
- role: worker
- role: worker
- role: worker
```

Para executar o comando acima faça:
    kind create cluster --config=file_name.yaml --name=cluster_name

### Comandos uteis kubectl

- kubectl get nodes - Mostra todos os nodes do cluster 
kubernets

O contexto é importante ser definido pois os comandos como por exemplo o kubeset get nodes busca diretamente no contexto, ou seja ele não mostra de fato todos os nodes.

    kubectl config get-clusters - Lista os clusters que existem no sistema.
    kubectl config use-context <context> - Troca o contexto do cluster para que os comandso sejam executados em outro contexto.




### Para ambiente de produção

Para ter estabilidade no ambiente de produção é interessante ter um build multi-master, ou seja mais de um control-plane por cluster para evitar ter ponto de falha se o master morrer.

