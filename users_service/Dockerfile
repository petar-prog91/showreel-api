FROM golang

ADD . /go/src/github.com/petar-prog91/showreel-api/users_service

RUN go install github.com/petar-prog91/showreel-api/users_service

ENTRYPOINT /go/bin/users_service

# Expose the application on port 8080
EXPOSE 8081