FROM golang

ADD . /go/src/skolar-api/users_service

RUN go install skolar-api/users_service

ENTRYPOINT /go/bin/users_service

# Expose the application on port 8080
EXPOSE 8081