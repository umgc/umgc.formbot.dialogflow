# App Config
SERVICE=formscriberapi
VERSION=0.2

# Kubernetes Config
RESOURCE_GROUP=formscriber
REGISTRY=formscriber
CLUSTER=formscriber-cluster
NAMESPACE=cert-manager
STATIC_IP=

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

###############################################################
# Makefile for Advance Development Factory (ADF) - Dialogflow
###############################################################

#Pull the latest ADF Dialogflow image from docker.io
adf-docker-pull:
	docker pull umgccaps/advance-development-factory-formbot-dialogflow:latest

#Run the ADF docker container
adf-docker-run: adf-docker-pull
	docker run -t -d -v $(PWD):/usr/src/formscriber --name adfcontainer umgccaps/advance-development-factory-formbot-dialogflow

#Login to Azure using docker container
adf-az-login: adf-docker-run
	docker exec adfcontainer az login

#Create Azure Resource Group using docker container
adf-az-rg-create: adf-az-login
	docker exec adfcontainer az group create --name formscriber --location eastus
	
#Create Azure ACR using docker container
adf-az-acr-create: adf-az-rg-create
	docker exec adfcontainer az acr create --resource-group formscriber --name formscriber1 --sku Basic
	
#Create Azure AKS cluster using docker container
adf-az-aks-create: adf-az-acr-create
	docker exec adfcontainer az aks create --resource-group $(RESOURCE_GROUP) --name $(CLUSTER) \
  				--enable-addons monitoring \
  				--node-count 1\
  				--generate-ssh-keys \
  				--attach-acr $(REGISTRY)
	
#Create Azure Public IP using docker container
adf-az-aks-all-create: adf-az-aks-create
	docker exec adfcontainer az network public-ip create \
		--resource-group MC_formscriber_formscriber-cluster_eastus \
		--name formscriberPublicIp --sku Standard --allocation-method static \
		--query publicIp.ipAddress -o tsv
		
##########################################################################################################

#Get Credential from AKS cluster
adf-aks-cred-get:
	docker exec adfcontainer az aks get-credentials --resource-group $(RESOURCE_GROUP) --name $(CLUSTER)

#Create AKS namespace using docker container
adf-aks-namespace-create: adf-aks-cred-get
	docker exec adfcontainer kubectl create namespace cert-manager
	
#Add the ingress-nginx repository using docker container
adf-helm-repo-add: adf-aks-namespace-create
	docker exec adfcontainer helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

#Use Helm to deploy an NGINX ingress controller using docker container
adf-helm-controller-deploy: adf-helm-repo-add
	docker exec adfcontainer helm install nginx-ingress ingress-nginx/ingress-nginx \
    					--namespace cert-manager \
    					--set controller.replicaCount=2 \
    					--set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
    					--set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux \
    					--set controller.admissionWebhooks.patch.nodeSelector."beta\.kubernetes\.io/os"=linux \
    					--set controller.service.loadBalancerIP="$(STATIC_IP)" \
    					--set controller.service.annotations."service\.beta\.kubernetes\.io/azure-dns-label-name"="formscriber-dso"

#Label the cert-manager namespace using docker container
adf-aks-namespace-label: adf-helm-controller-deploy
	docker exec adfcontainer kubectl label namespace cert-manager cert-manager.io/disable-validation=true

#Add the Jetstack Helm repository using docker container
adf-helm-jetstack-add: adf-aks-namespace-label
	docker exec adfcontainer helm repo add jetstack https://charts.jetstack.io

#Update your local Helm chart repository cache using docker container
adf-helm-repo-update: adf-helm-jetstack-add
	docker exec adfcontainer helm repo update
	
#Install the cert-manager Helm chart using docker container
adf-helm-cert-install: adf-helm-repo-update
	docker exec adfcontainer helm install cert-manager \
  				--namespace cert-manager \
  				--version v0.16.1 \
  				--set installCRDs=true \
  				--set nodeSelector."beta\.kubernetes\.io/os"=linux jetstack/cert-manager
	docker exec adfcontainer sleep 20
	
#Create the issuer cert using docker container
adf-kube-cert-create: adf-helm-cert-install 
	docker exec adfcontainer kubectl apply -f /usr/src/formscriber/deploy/ssl/cluster-issuer-staging.yaml
	
#Build image using docker container
adf-az-acr-build: adf-kube-cert-create
	docker exec -w /usr/src/formscriber adfcontainer az acr build --image $(SERVICE):$(VERSION) --registry $(REGISTRY) --file Dockerfile .

#Deploy image to AKS using docker container
adf-image-aks-deploy: adf-az-acr-build
	docker exec adfcontainer helm upgrade --install formscriberapi /usr/src/formscriber/deploy/formscriberapi/ --namespace $(NAMESPACE)

#################################################################################################################################################

#Run Go "build" from docker container
adf-go-build:
	docker run --rm -v "$$PWD"/formscriber:/usr/src/formscriber -w /usr/src/formscriber umgccaps/advance-development-factory-formbot-dialogflow:latest go build -v
	
#Run Go "test" from docker container
adf-go-test:
	docker run --rm -v "$$PWD"/formscriber:/usr/src/formscriber -w /usr/src/formscriber umgccaps/advance-development-factory-formbot-dialogflow:latest go test

#Delete all docker containers
adf-docker-all-delete:
	docker rm $$(docker ps -aq) --force
	
#Delete AKS Cluster
adf-aks-delete:
	docker exec adfcontainer az group delete --name formscriber --yes --no-wait
