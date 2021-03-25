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

###############################################################
# Makefile for Advance Development Factory (ADF) - Dialogflow
###############################################################

#pull the latest ADF Dialogflow image from docker.io
adf-docker-pull
	docker pull umgccaps/advance-development-factory-formbot-dialogflow:latest

#run the ADF docker container
adf-docker-run
	docker run -t -d --name adfcontainer umgccaps/advance-development-factory-formbot-dialogflow

#login to Azure using docker container
adf-az-login
	docker exec adfcontainer read -p "Azure UserName: " AZ_USER && echo && read -sp "Azure password: " AZ_PASS && echo && az login -u $AZ_USER -p $AZ_PASS

#create Azure Resource Group using docker container
adf-az-rg-create
	docker exec adfcontainer az group create --name $(RESOURCE_GROUP) --location eastus
	
#create Azure ACR using docker container
adf-az-acr-create
	docker exec adfcontainer az acr create --resource-group formscriber --name $(REGISTRY) --sku Basic
	
#create Azure AKS cluster using docker container
adf-az-aks-create
	docker exec adfcontainer az aks create -resource-group $(RESOURCE_GROUP) --name $(CLUSTER) \
  				--enable-addons monitoring, http_application_routing \
  				--node-count 1\
  				--generate-ssh-keys \
  				--attach-acr $(REGISTRY)

#create Azure Public IP
adf-az-ip-create
	docker exec adfcontainer az network public-ip create \
		--resource-group MC_formscriber_formscriber-cluster_eastus \
		--name formscriberPublicIp --sku Standard --allocation-method static \
		--query publicIp.ipAddress -o tsv

#create AKS namespace
adf-aks-namespace-create 
	docker exec adfcontainer kubectl create namespace cert-manager
	
#
	
