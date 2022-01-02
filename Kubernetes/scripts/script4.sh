#! /bin/sh

kubectl create -f ns.yaml
kubectl create -f dummy.yaml
kubectl create -f backend.yaml
kubectl create -f frontend.yaml
kubectl create -f redis-pub.yaml
kubectl create -f redis-sub.yaml
kubectl create -f grpc-client.yaml
kubectl create -f grpc-server.yaml
kubectl create -f traffic-splitter.yaml

kubectl -n project get deploy -o yaml | linkerd inject - | kubectl apply -f -
kubectl -n project get deploy -o yaml | linkerd inject - | kubectl apply -f -

kubectl get all -n nginx-ingress

echo revisar que todo funcione