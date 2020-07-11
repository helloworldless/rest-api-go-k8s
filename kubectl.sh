#!/usr/bin/env bash

kubectl create -f db-service.yaml,db-deployment.yaml,posts-api-service.yaml,posts-api-claim0-persistentvolumeclaim.yaml,posts-api-deployment.yaml