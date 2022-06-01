# From Code to Kubernetes
## Workshop preparation

Installing kind:

	curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.14.0/kind-darwin-amd64
	chmod +x ./kind
	sudo mv kind /usr/local/bin

Installing helm

	curl -LO https://get.helm.sh/helm-v3.9.0-darwin-amd64.tar.gz
	tar zxf helm-v3.9.0-darwin-amd64.tar.gz
	sudo mv darwin-amd64/helm /usr/local/bin/


