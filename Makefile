run: ssl-local
	cd formscriber && go run .

ssl-local:
	openssl req -x509 -out formscriber/formscriber.com.pem \
	-keyout formscriber/formscriber.key -newkey rsa:2048 -nodes -sha256 \
	-subj '/CN=localhost'

build-docker:
	docker build -t formscriber.azurecr.io/formscriber_df .

run-docker:
	docker run formscriber.azurecr.io/formscriber_df

push-image: acr-login
	az acr build --image formscriberapi --registry formscriber --file Dockerfile .

acr-login:
	az acr login --name formscriber

aks-login:
	az aks get-credentials --resource-group formscriber --name formscriber-cluster

deploy: aks-login
	helm install formscriberapi deploy/

undeploy: aks-login
	helm uninstall formscriberapi

clean:
	find . | grep -E '(\.log|\.pem|\.key)' | xargs rm -rf