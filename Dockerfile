FROM golang:alpine AS build

WORKDIR /go/src/app

COPY . .

COPY --from=quay.io/goswagger/swagger /usr/bin/swagger /usr/bin/swagger

RUN swagger generate server --exclude-spec --name=core --spec=./swagger/swagger.yaml

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app ./cmd/core-server

FROM alpine:latest AS alpine

RUN apk --no-cache add ca-certificates

FROM scratch

COPY --from=build /go/bin/app /app

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/app"]
