#!/usr/bin/env bash
kubectl delete pods grpcserver && kubectl apply -f grpc_server.yaml
