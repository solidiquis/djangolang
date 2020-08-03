### Go ###
FROM golang:latest as go_builder
COPY ./server /app
WORKDIR /app/server
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
CMD CompileDaemon -log-prefix=false -directory="." -build="go build ./main.go"  -command="./main"

### React ###
FROM node:alpine as node_builder
COPY ./client /app
WORKDIR /app/client
RUN yarn
CMD yarn start
