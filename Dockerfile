FROM golang:1.26-alpine

RUN adduser -D appuser

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

USER appuser

EXPOSE 8080

CMD ["./main"]