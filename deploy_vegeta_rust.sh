#!/usr/bin/env bash
export SERVER_IP=`kubectl get pods rusthttp -o jsonpath='{.status.podIP}'`
export QPS=50000
export DATA="{}"
export DATA=`cat payload_rust.json`
#export JEST=`jq -n --arg b "${DATA}" '$b'`

read -r -d '' QUERY << EndOfMessage
POST http://${SERVER_IP}:8080/echo
Content-Type: application/json
EndOfMessage

kubectl delete pods vegeta vegeta1 vegeta2 vegeta3 vegeta4
# | "
kubectl run vegeta1 --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=32 | tee results.bin | vegeta report"
kubectl run vegeta2 --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=32 | tee results.bin | vegeta report"
kubectl run vegeta3 --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=32 | tee results.bin | vegeta report"
kubectl run vegeta4 --image="peterevans/vegeta:latest" --restart=Never -- sh -c "echo '$DATA' > body.json && echo '$QUERY' | vegeta attack -body body.json -rate=$QPS -duration=60s -max-workers=32 | tee results.bin | vegeta report"

kubectl wait --for=condition=Ready --timeout=600s pod/vegeta
kubectl logs -f vegeta