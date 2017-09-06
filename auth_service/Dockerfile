FROM golang

ADD . /go/src/skolar-api/auth_service

RUN go install skolar-api/auth_service

ENTRYPOINT /go/bin/auth_service

# Expose the application on port 8080
EXPOSE 8082