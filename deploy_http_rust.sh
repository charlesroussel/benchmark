#!/usr/bin/env bash
export CORE_COUNT=3

kubectl delete pods rusthttp
cat rust_http_server.yaml | envsubst | kubectl apply -f -
kubectl wait --for=condition=Ready --timeout=60s pod/rusthttp