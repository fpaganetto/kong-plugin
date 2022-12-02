#!/bin/sh

cd /app/main
go mod download
go build -o ../bin/go-jwt-aud
