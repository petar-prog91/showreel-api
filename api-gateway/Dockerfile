FROM golang

ADD . /go/src/skolar-api/api_gateway

RUN go install skolar-api/api_gateway

ENTRYPOINT /go/bin/api_gateway

# Expose the application on port 8080
EXPOSE 8080