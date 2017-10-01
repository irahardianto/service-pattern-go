FROM golang:1.9.0 as builder
WORKDIR $GOPATH/src/github.com/irahardianto/service-pattern-go/
RUN go get -u github.com/golang/dep/cmd/dep
COPY . ./
RUN $GOPATH/bin/dep ensure
RUN go build --tags service-pattern-go --ldflags '-extldflags "-lm -lstdc++ -static"'

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/irahardianto/service-pattern-go/service-pattern-go .
CMD ["./service-pattern-go"]

EXPOSE 8080
