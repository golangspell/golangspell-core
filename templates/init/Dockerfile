FROM golang:1.19.0-alpine3.16 AS golang

ARG LOC=/builds/go/src/{{.ModuleName}}/
RUN apk add --no-cache git
RUN mkdir -p $LOC
ENV GOPATH /go
COPY . $LOC
ENV CGO_ENABLED 0
RUN cd $LOC && TESTRUN=true go test ./... && go build

FROM alpine:3.16
ARG LOC=/builds/go/src/{{.ModuleName}}
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang $LOC/{{.AppName}} /app

EXPOSE 8080
RUN chmod +x /app/{{.AppName}}
WORKDIR /app    
USER app
CMD ["/app/{{.AppName}}"]
