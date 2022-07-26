# fabric-sail

Fabric `sail` is a lightweight, companion driver for the [fabric-operator](https://github.com/hyperledger-labs/fabric-operator).

When building networks with `sail`, a blockchain is specified as a concise yaml document describing the topology of
the target network.  `sail` reads the topology document, applying updates to the Kubernetes API controller, where
custom resources will be reflected on the cluster and realized with the operator.  This project aims to determine the
_minimal_ specification necessary to convey the Fabric network topology, hiding much or all of the complexity of k8s
interactions behind a simple veneer.

`sail` is a proof-of-concept, inspired by: 

- minikube
- fabric-test-operator 
- The Fabric Smart Client integration test suite
- Hyperledger Fabric "in a box"
- fabric-hyperkube

`sail` can be used to quickly spin up networks with varying topologies.  It can serve as a replacement for the
fabric-samples test network, the operator sample-network, or a one-step mechanism to quickly spin up a network with
a target topology for a performance benchmarks or integration test.

`sail` is a prototype, NOT a production system.  _Caveat emptor._



## Updated approach:  Use Argo Workflows 

There isn't (yet) a good way to programmatically invoke an administrative API for all steps necessary to 
set up, admin, and manage a network.   Instead of writing a golang SDK to - effectively - fork+exec peer 
binaries behind the scenes, let's try a really simple approach: 

1. Install Kubernetes 
2. Install Argo Workflows 
3. Launch an Argo Workflow to construct the Fabric Network 
4. Launch an Argo Workflow to construct channels, join peers, etc. 
5. Launch an Argo Workflow to install chaincode
6. ... 

Channel construction, chaincode installation, etc. can be refactored as sub-workflows or workflow templates. 


## Quickstart (SCRATCH NOTES)

```shell
export NS=test-network

network kind 
network cluster init
```

```shell
network operator
```

```shell
kubectl -n $NS apply -f https://raw.githubusercontent.com/argoproj/argo-workflows/master/manifests/quick-start-postgres.yaml
kubectl -n $NS apply -f kube/config/base/sail-role.yaml
kubectl -n $NS apply -f kube/config/base/sail-rolebinding.yaml

sleep 10

kubectl -n $NS rollout status deploy argo-server
```

Argo GUI (terminal II): 
```shell
kubectl -n $NS port-forward deployment/argo-server 2746:2746 & 
open https://localhost:2746 
```

Test network: 
```shell
argo -n $NS submit --watch kube/workflows/network-up.yaml 
argo -n $NS submit --watch kube/workflows/network-channel-create.yaml
```

```shell
argo -n $NS submit --watch kube/workflows/network-chaincode-conga.yaml
```

Invoke / query chaincode: 
```shell
TODO: do ... 
```

## Reset: 

```shell
network down
```
or 
```shell
network unkind
```