#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods grpcserver -o jsonpath='{.status.podIP}'`
export QPS=60000
export CONCURRENCY=32

kubectl delete pods burster
cat burster.yaml | envsubst | kubectl apply -f -

kubectl wait --for=condition=Ready --timeout=600s pod/burster
kubectl logs -f burster