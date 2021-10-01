#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods rusthttp -o jsonpath='{.status.podIP}'`
export QPS=50000
export DATA="{}"
export WORKERS=32
export DURATION="180s"
#export DATA=`cat mopub_c.json`
#export JEST=`jq -n --arg b "${DATA}" '$b'`

kubectl delete job vegeta
cat vegeta.yaml | envsubst | kubectl apply -f -