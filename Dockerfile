FROM node:12 as nuxt-builder
WORKDIR /usr/src/app
COPY ./rabums_view/package.json ./
RUN yarn
COPY ./rabums_view/. .
RUN yarn generate

FROM golang:1.14 as go-builder
WORKDIR $GOPATH/src/github.com/CSUOS/rabums
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /go/bin/rabums main.go

FROM alpine:3.11.3
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=go-builder /go/bin/rabums /bin/rabums
COPY --from=nuxt-builder /usr/src/app/dist/. /dist/.
CMD ["/bin/rabums"]