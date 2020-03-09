FROM golang:1.14.0-alpine3.11 AS golang

RUN apk add --no-cache git
RUN mkdir -p /builds/go/src/{{.ModuleName}}/
ENV GOPATH /go
COPY . /builds/go/src/{{.ModuleName}}/
ENV CGO_ENABLED 0
RUN cd /builds/go/src/{{.ModuleName}}/ && TESTRUN=true go test ./... && go build

FROM alpine:3.11
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang /builds/go/src/{{.ModuleName}}/{{.AppName}} /app

EXPOSE 8080
RUN chmod +x /app/{{.AppName}}
WORKDIR /app    
USER app
CMD ["/app/{{.AppName}}"]
