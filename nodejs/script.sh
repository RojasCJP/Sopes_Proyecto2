#! /bin/sh

cd ..
cd home
ls
cd Frontend
npm install -g @angular/cli
npm install
export NODE_OPTIONS=--openssl-legacy-provider
ng serve