FROM golang

ADD . /go/src/showreel-api/api_gateway

RUN go install showreel-api/api_gateway

ENTRYPOINT /go/bin/api_gateway

# Expose the application on port 8080
EXPOSE 8080