#! /bin/sh

linkerd check --pre
linkerd install | kubectl apply -f -

linkerd viz install | kubectl apply -f -

echo ver los faltantes