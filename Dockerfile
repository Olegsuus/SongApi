FROM golang:1.23.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

EXPOSE 4444

CMD ["./main"]