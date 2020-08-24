#!/bin/bash

#リバースプロキシを使わずに開発用で
#フロント、バックエンド、DBを個別に立ち上げます
#特にフロントエンドは開発中はホットリロードを使いたいので

#DBコンテナを起動
docker-compose run -d db

popd ./frontend
npm start
pushd

popd ./app
go run -mode develop ./cmd/main.go
pushd

