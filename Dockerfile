FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/app
EXPOSE 8888
CMD ["./main"]