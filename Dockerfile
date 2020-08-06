FROM golang:1.14 as builder
WORKDIR $GOPATH/src/github.com/CSUOS/rabums
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -mod=readonly -o /go/bin/rabums cmd/*.go

FROM alpine:3.11.3
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/bin/rabums /bin/rabums
CMD ["/bin/rabums"]