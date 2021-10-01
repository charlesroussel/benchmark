#!/usr/bin/env bash
export CORE_COUNT=3
export THREAD_COUNT=3

kubectl delete pods ginserver
cat ginserver.yaml | envsubst | kubectl apply -f -
