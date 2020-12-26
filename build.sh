#!/bin/sh
export GIN_MODE=release
export GOOS=linux
export GOARCH=amd64
go build -o ./dist/main main.go

git submodule update --remote --init ui

cd ui && npm install

npm run build:prod

cp -r dist/blog-cms ../dist/public

cd ..

tar -czf dist.tar ./dist