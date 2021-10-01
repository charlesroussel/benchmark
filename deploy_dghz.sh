#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods grpcserver -o jsonpath='{.status.podIP}'`
export QPS=60000
export CONCURRENCY=16

kubectl delete job dghz
cat dghz.yaml | envsubst | kubectl apply -f -
kubectl logs -f -l job-name=dghz --all-containers=true