.PHONY: requirement image* dev act sync

IMAGE_NAME=cnsumi/xen-miner
IMAGE_TAG=latest
IMAGE="${IMAGE_NAME}:${IMAGE_TAG}"

requirements:
	pipreqs . --force

image:
	docker build -t ${IMAGE} .

image-go:
	docker build -f Dockerfile-go .

image-run:
	docker run -it --rm ${IMAGE}

image-buildx:
	docker buildx build \
	--platform linux/amd64,linux/arm64 \
	.

dev:
	python3 miner.py

act:
	act workflow_dispatch \
	--actor CNSumi \
	-s DOCKERHUB_USERNAME \
	-s DOCKERHUB_PASSWORD

act-image-py:
	act workflow_dispatch \
	--actor CNSumi \
	-j image-py \
	-s DOCKERHUB_USERNAME \
	-s DOCKERHUB_PASSWORD

act-image-go:
	act workflow_dispatch \
	--actor CNSumi \
	-j image-go \
	-s DOCKERHUB_USERNAME \
	-s DOCKERHUB_PASSWORD

sync: 
	rsync -avh . \
	--exclude=".git" \
	--exclude=".vscode" \
	r86s:~/xen-miner