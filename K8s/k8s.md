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


## Namespace

O namespace permite uma separação lógica da aplicação dessa forma é possível setar:
- permissão por usuario
- Limite de recurso por namespace

É normalmente utilizado para separação entre projetos e é recomendado que o ambiente de dev fique em um cluster separado.

### Comandos uteis

    kubectl get ns - lista todos os namespaces.
    kubectl create ns <nome> - cria um namespace.
    kubectl set-context --namespace=<nome> - Muda o contexto atual para o namespace informado.

    kubectl apply -f <yaml_file> -n=<namespace_name> - executa o arquivo de manifesto no namespace informado.

    kubectl get view - mostra um conjunto de informações do cluster

É possível inserir o namespace também no arquivo de manifesto, o que é mais seguro que deixar para colocar no comando.

No k8s é possível criar varios contextos diferentes com isso entrando no cluster com uma pré-configuração. 

    kubectl config current-context - Mostra o contexto atual
    kubectl config set-context dev --namespace=dev --cluster=<cluster_name> --user=<usuario>
    kubectl config use-context dev - muda para o contexto solicitado.

#### Service account

Atribuir uma conta de serviço para um pod dessa forma ele fica limitado as permissões. <br>
Se não criar isso e uam aplicação for invadida é possível escalar privilégio dentro do cluster e VRAU.

    kubectl get serviceaccounts - lista todos os services account do cluster.

É recomendado criar uma pelo menos por projeto.

O service account padrão mapeia um diretório dentro de cada instancia para saber o grau de permissão da aplicação, dessa forma se alguem invadir a aplicação é possível ler as credenciais do cluster e executar comando pra API do k8s, para evitar isso é importante trocar o serviceaccount.

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: server

---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: server-read
rules:
  - apiGroups: [""] #Determina quais os recursos ele pode acessar.
    resources: ["pods","services"] #Tipos de serviços que ele podem acessar.
    verbs: ["get","watch","list"] #São os comandos que podem ser executados nos recursos acima.
  - apiGroups: ["apps"] #Determina quais os recursos o deployments pode usar.
    resources: ["deployments"] #Tipos de serviços que ele podem acessar.
    verbs: ["get","watch","list"] #São os comandos que podem ser executados nos recursos acima.

--- 

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: server-read-bind
subjects:
  - kind: ServiceAccount
    name: server
    #namespace: prod
  roleRef:
    kind: Role
    name: server-read
    apiGroups: rbac.authorization.k8s.io


```

A Role é responsável por atribuir as permissões do serviceaccount;

Para verificar as APIs de cada serviço é possível utilizar o comando abaixo

    kubectl api-resources

Após criar o serviceaccount e a Role é necessário "bindar" a Role ao service.

Após a criação de tudo isso ainda é necessário associar essas informações ao Deployment. Para fazer isso dentro da especificação do dpeloyment em specs.templates.spec.serviceAccount deve estar referenciado corretamente.

É possível fazer as regras em nivel de cluster utilizando "ClusterRole", "ClusterRoleBinding" para que as permissões funcionam entre os namespaces.

[Codigos do curso](https://github.com/codeedu/fc2-k8s)

## Service Mesh - Istio

É uma camada extra adiciona junto a seu cluster para monitorar e modificar em tempo real o trafego das aplicações, bem como elevar o nível de segurança e confiabilidade de todo ecossistema.

Istio é um projeto open-source que implementa service mesh visando diminuir a complexidade no gerenciamento de aplicações distribuidas independente de qual linguagem ou tecnologia elas foram desenvolvidas.

- Roda em cima de diversas soluções.

É possível rodar ele em cima do kubernets, Apache Mesos("Concorrente do k8s + algumas features), Consul(Service Discovery) e Nomad(Hashcorp). 








