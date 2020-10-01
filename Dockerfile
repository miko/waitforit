FROM golang:alpine
RUN GO111MODULE=on go get github.com/miko/waitforit
ENTRYPOINT ["/go/bin/waitforit"]
WORKDIR /app
RUN which waitforit
CMD ["-address", "https://jsonplaceholder.typicode.com/todos/1", "-body", "-expected_json", "'{\"userId\":1}'", "-debug"]
