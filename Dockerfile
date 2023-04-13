FROM golang:1.20.3-bullseye

COPY . /app
WORKDIR /app
RUN go mod tidy

RUN go install github.com/pact-foundation/pact-go/v2@2.x.x
RUN pact-go -l DEBUG install

ENTRYPOINT pact-go version && pact-go check && go test ./...
