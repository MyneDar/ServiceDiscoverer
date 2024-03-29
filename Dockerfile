FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN ls
RUN go build -v -o cmd ./...
EXPOSE 10000

CMD ["go", "run", "cmd/main.go"]