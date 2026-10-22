FROM golang:1.27.1-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags migrate -o app ./cmd/app

CMD ["./app"]