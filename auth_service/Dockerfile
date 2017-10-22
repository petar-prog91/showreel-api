FROM golang

ADD . /go/src/github.com/petar-prog91/showreel-api/auth_service

RUN go install github.com/petar-prog91/showreel-api/auth_service

ENTRYPOINT /go/bin/auth_service

# Expose the application on port 8080
EXPOSE 8082