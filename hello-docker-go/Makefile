
VERSION=latest

build:
	go build ./...

docker-build:
	docker build -t hello-docker-go:$(VERSION) .
	minikube image load hello-docker-go:$(VERSION)

docker-run:
	docker compose up

setup-minikube:
	minikube start --driver=hyperkit
	minikube addons enable ingress
	minikube addons enable dashboard
	minikube addons enable metrics-server

set-version:
	kustomize edit set image latest

dashboard:
	minikube dashboard

deploy-test:
	cd k8s && kubectl apply -k overlay/test

deploy-prod:
	cd k8s && kubectl apply -k overlay/prod

test:
	#minikube --namespace hello-docker-go service hello-docker-go-service --url
	curl --resolve test.hello-docker-go.local:80:$(shell minikube ip) -i http://test.hello-docker-go.local/health
