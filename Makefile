# App Config
SERVICE=formscriberapi
VERSION=0.2

# Kubernetes Config
RESOURCE_GROUP=formscriber
REGISTRY=formscriber
CLUSTER=formscriber-cluster
NAMESPACE=cert-manager

run: ssl-local
	cd formscriber && go run .

ssl-local:
	openssl req -x509 -out formscriber/formscriber.com.pem \
	-keyout formscriber/formscriber.key -newkey rsa:2048 -nodes -sha256 \
	-subj '/CN=localhost'

run-docker: build-docker
	docker run -p 8080:80 --name formscriberapi $(REGISTRY).azurecr.io/$(SERVICE)

build-docker:
	docker build -t $(REGISTRY).azurecr.io/$(SERVICE) .

rm-docker:
	docker rm -f formscriberapi

deploy: aks-login push-image
	helm upgrade --install formscriberapi deploy/formscriberapi/ --namespace $(NAMESPACE)

undeploy: aks-login
	helm uninstall $(SERVICE) --namespace $(NAMESPACE)

push-image: acr-login
	az acr build --image $(SERVICE):$(VERSION) --registry $(REGISTRY) --file Dockerfile .

aks-login:
	az aks get-credentials --resource-group $(RESOURCE_GROUP) --name $(CLUSTER)

acr-login:
	az acr login --name $(REGISTRY)

clean:
	find . | grep -E '(\.log|\.pem|\.key)' | xargs rm -rf

.PHONY: clean run