#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods httpserver -o jsonpath='{.status.podIP}'`
export QPS=26000
export DATA="{}"
#export DATA=`cat mopub_c.json`
#export JEST=`jq -n --arg b "${DATA}" '$b'`

read -r -d '' QUERY << EndOfMessage
POST http://${SERVER_IP}:8080/ad
Content-Type: application/json
EndOfMessage

kubectl delete pods vegeta
# | "
kubectl run vegeta --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=256 | tee results.bin | vegeta report"
kubectl wait --for=condition=Ready --timeout=600s pod/vegeta
kubectl logs -f vegeta