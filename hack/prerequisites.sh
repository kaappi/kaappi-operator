#!/bin/bash
DEP_VERSION="0.5.1"
GOLANGCI_VERSION="1.16.0"
SDK_VERSION="0.7.0"
MINIKUBE_VERSION="1.0.0"
KUBE_VERSION="1.14.1"

# Download Dep
curl -Lo dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && chmod +x dep && mv dep $GOPATH/bin/
# Download GolangCI lint
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin/ v${GOLANGCI_VERSION}
# Download Operator SDK
curl -Lo operator-sdk https://github.com/operator-framework/operator-sdk/releases/download/v${SDK_VERSION}/operator-sdk-v${SDK_VERSION}-x86_64-linux-gnu && chmod +x operator-sdk && mv operator-sdk $GOPATH/bin/
# Download kubectl
curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/v${KUBE_VERSION}/bin/linux/amd64/kubectl && chmod +x kubectl && mv kubectl $GOPATH/bin/
# Download Minikube
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v${MINIKUBE_VERSION}/minikube-linux-amd64 && chmod +x minikube && mv minikube $GOPATH/bin/

