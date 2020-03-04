

install:
	cd cmd/ds-circleci-plugin && go install

image-push:
	docker build -t quay.io/dotmesh/dotscience-circleci-plugin:latest -f Dockerfile .
	docker push quay.io/dotmesh/dotscience-circleci-plugin:latest