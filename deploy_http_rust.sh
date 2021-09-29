#!/usr/bin/env bash
export CORE_COUNT=3
export THREAD_COUNT=3

kubectl delete pods rusthttp
cat rust_http_server.yaml | envsubst | kubectl apply -f -
