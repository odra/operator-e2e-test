SHELL = /bin/bash
REG = docker.io
ORG = odranoel
IMAGE = stuff
TAG = latest
NS = stuff-test
SA = stuff

prepare:
	@oc apply -f deploy/sa.yaml -n ${NS}
	@oc apply -f deploy/rbac.yaml -n ${NS}

build:
	@echo Building operator image "${REG}/${ORG}/${IMAGE}:${TAG}"...
	@operator-sdk build --enable-tests ${REG}/${ORG}/${IMAGE}:${TAG}
	@echo Done!

push:
	@echo Pushing operator image to registry "${REG}/${ORG}/${IMAGE}:${TAG}"...
	@docker push ${REG}/${ORG}/${IMAGE}:${TAG}
	@echo Done!

test-e2e: build push
	@echo Running e2e tests against remote cluster...
	@operator-sdk test cluster --namespace ${NS} --service-account ${SA} ${REG}/${ORG}/${IMAGE}:${TAG}
	@echo Done!
