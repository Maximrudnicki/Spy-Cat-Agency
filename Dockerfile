FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/test_rudnytskyi ./cmd/.

EXPOSE 8000

CMD ["./bin/test_rudnytskyi"]