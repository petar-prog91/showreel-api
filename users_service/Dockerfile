FROM golang

ADD . /go/src/github.com/petar-prog91/showreel-api/users_service
ENV D=/go/src/github.com/petar-prog91/showreel-api/users_service

RUN go get github.com/golang/dep/cmd/dep
RUN cd $D && dep ensure --vendor-only
RUN go install github.com/petar-prog91/showreel-api/users_service

ENTRYPOINT /go/bin/users_service

# Expose the application on port 8081
EXPOSE 8081