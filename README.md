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


## Quickstart 

1. Create a KIND cluster by running `network kind` from fabric-operator/sample-network
```shell
network kind 
```

2. Create the "test network," `mychannel`, and install the basic asset transfer smart contract:
```shell
sail apply -f samples/asset-transfer-basic.yaml
```

3. Invoke chaincode:
```shell
TODO
```

4. Run a gateway client application: 
```shell
TODO
```
