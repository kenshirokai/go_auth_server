#!/bin/bash

pushd ./frontend
npm run build
cp ./dist/* ../app/static/
popd
# pushd ./app
# go run ./cmd/main.go
docker-compose build
docker-compose up -d
