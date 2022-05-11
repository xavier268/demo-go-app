#!/bin/bash

go version 
GOARCH=wasm GOOS=js go build -o web/app.wasm
go build

echo "Ctrl-C to exit"

google-chrome http://localhost:8080 &
./demo-go-app 

