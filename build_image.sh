#!/bin/bash

# BRANCH=$(git rev-parse --abbrev-ref HEAD | tr / - | tr _ -)
# SHORT_COMMIT=$(git log -n 1 --pretty=format:'%h')
DOCKERFILE_PATH=build/images/tf_operator/Dockerfile
IMAGE_TAG=$(python -c "print('v' + '${TRAVIS_TAG}'.split('v')[1])")
IMAGE_NAME=${DOCKER_IMAGE_ORG}/kubeflow-tf-operator:${IMAGE_TAG}
echo "Image name is ${IMAGE_NAME}"
echo "$DOCKER_HUB_ACCESS_TOKEN" | docker login -u "$DOCKER_HUB_USERNAME" --password-stdin && \
  docker build -t $IMAGE_NAME -f $DOCKERFILE_PATH . && \
  docker push $IMAGE_NAME
