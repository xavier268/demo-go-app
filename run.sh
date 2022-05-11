#!/bin/bash

go version 
GOARCH=wasm GOOS=js go build -o web/app.wasm
go build
firefox localhost:8080 &
echo "Ctrl-C to exit"
./goapp

