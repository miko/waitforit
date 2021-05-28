FROM golang:alpine
RUN apk add git
RUN go get -u -v github.com/miko/waitforit/v2
ENTRYPOINT ["/go/bin/waitforit"]
WORKDIR /app
RUN which waitforit
CMD ["-address", "https://jsonplaceholder.typicode.com/todos/1", "-body", "-expected_json", "{\"userId\":1}", "-debug"]

