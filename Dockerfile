FROM golang:1.11.1-alpine3.8 AS build-env
WORKDIR /go/src/github.com/youtangai/cloud-training
COPY ./ ./
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/youtangai/cloud-training/server /usr/local/bin/server
ENV DB_USER cloud
ENV DB_PASS fun
ENV DB_IP db
ENV DB_PORT 3306
ENV DB_NAME cloud-training

EXPOSE 8080
CMD ["/usr/local/bin/server"]