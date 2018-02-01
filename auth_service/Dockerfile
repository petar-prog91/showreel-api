FROM golang

ADD . /go/src/github.com/petar-prog91/showreel-api/auth_service
ENV D=/go/src/github.com/petar-prog91/showreel-api/auth_service

RUN go get github.com/golang/dep/cmd/dep
RUN cd $D && dep ensure --vendor-only
RUN go install github.com/petar-prog91/showreel-api/auth_service

ENTRYPOINT /go/bin/auth_service

# Expose the application on port 8082
EXPOSE 8082