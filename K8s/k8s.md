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

    kind create cluster

Por padrão o kind cria apenas um nó.

É imprecindível para manipular o container vincular o kubectl a um determinado cluster, no final do processo é informado o comando necessário para vincular ele ao cluster.

    kubectl cluster-info --context kind-kind

### Comandos uteis kubectl

- kubectl get nodes - Mostra todos os nodes do cluster kubernets

