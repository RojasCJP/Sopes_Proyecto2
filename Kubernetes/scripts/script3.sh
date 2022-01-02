#! /bin/sh

helm repo add stable https://charts.helm.sh/stable
kubectl create namespace nginx-ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress

helm list -n nginx-ingress

kubectl -n nginx-ingress get deployment nginx-ingress-ingress-nginx-controller -o yaml | linkerd inject --ingress --skip-inbound-ports 443 --skip-outbound-ports 443 - | kubectl apply -f -
kubectl get pods -n nginx-ingress

echo ver los faltantes