FROM golang:1.11-alpine3.8 AS golang

RUN apk add --no-cache git
RUN mkdir -p /builds/go/src/{{.ModuleName}}/
ENV GOPATH /go
COPY . /builds/go/src/{{.ModuleName}}/
ENV CGO_ENABLED 0
RUN cd /builds/go/src/{{.ModuleName}}/ && TESTRUN=true go test ./... && go build

FROM alpine:3.8
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
COPY --from=golang /builds/go/src/{{.ModuleName}}/{{.AppName}} /
RUN update-ca-certificates

EXPOSE 8080
RUN chmod +x /{{.AppName}}
CMD ["/{{.AppName}}"]
