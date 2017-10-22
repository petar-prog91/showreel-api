FROM golang

ADD . /go/src/github.com/petar-prog91/showreel-api/api_gateway
ENV D=/go/src/github.com/petar-prog91/showreel-api/api_gateway

RUN go get github.com/golang/dep/cmd/dep
RUN cd $D && dep ensure --vendor-only
RUN go install github.com/petar-prog91/showreel-api/api_gateway

ENTRYPOINT /go/bin/api_gateway

# Expose the application on port 8080
EXPOSE 8080