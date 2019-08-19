GCR=

.PHONY: build
build:
	docker build -t rudder:dev .
	docker tag rudder:dev $(GCR)/rudder:dev
	docker push $(GCR)/rudder:dev

install:
	kubectl apply -f ./artifact/rbac
	kubectl apply -f ./artifact/stg

.PHONY: reload
reload:
	kubectl delete po rudder || echo skip
	cd ./artifact/stg; kubectl apply -f rudder.yaml

