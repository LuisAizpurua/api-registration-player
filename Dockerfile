FROM golang:1.26-alpine

RUN adduser -D appuser

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

ENV NODE_IP=192.168.49.2
ENV SVC_PORT=30090

USER appuser

EXPOSE 8080

CMD ["./main"]