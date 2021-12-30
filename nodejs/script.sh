#! /bin/sh

cd ..
cd home
ls
cd Frontend
npm install http-server -g
npm install -g @angular/cli
npm install
ng build --prod
cd dist
http-server Frontend/
top
