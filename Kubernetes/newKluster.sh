#! /bin/sh

gcloud container clusters create k8s-demo --num-nodes=1 --tags=allin,allout --enable-legacy-authorization --issue-client-certificate --preemptible --machine-type=n1-standard-2

# curl -fsL https://run.linkerd.io/install | sh
# export PATH=$PATH:$HOME/.linkerd2/bin
linkerd check --pre
linkerd install | kubectl apply -f -

linkerd viz install | kubectl apply -f -

# wget https://get.helm.sh/helm-v3.7.2-linux-amd64.tar.gz
# tar -zxvf helm-v3.7.2-linux-amd64.tar.gz
# mv linux-amd64/helm /usr/local/bin/helm
# helm version

helm repo add stable https://charts.helm.sh/stable
kubectl create namespace nginx-ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress

helm list -n nginx-ingress

kubectl -n nginx-ingress get deployment nginx-ingress-ingress-nginx-controller -o yaml | linkerd inject --ingress --skip-inbound-ports 443 --skip-outbound-ports 443 - | kubectl apply -f -
kubectl get pods -n nginx-ingress
# kubectl describe pods nginx-ingress-ingress-nginx-controller-64665dd6bc-dgkj4 -n nginx-ingress | grep "linkerd.io/inject: ingress"


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