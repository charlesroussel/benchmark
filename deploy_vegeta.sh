#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods httpserver -o jsonpath='{.status.podIP}'`
export QPS=50000
export DATA="{}"

CONCURRENCY=32
#export DATA=`cat mopub_c.json`
#export JEST=`jq -n --arg b "${DATA}" '$b'`

read -r -d '' QUERY << EndOfMessage
POST http://${SERVER_IP}:8080/ad
Content-Type: application/json
EndOfMessage

kubectl delete pods vegeta1
kubectl run vegeta1 --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=${CONCURRENCY} | tee results.bin | vegeta report"
kubectl wait --for=condition=Ready --timeout=600s pod/vegeta1
kubectl logs -f vegeta1