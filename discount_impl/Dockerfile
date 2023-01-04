FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go clean --cache
RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env ./

EXPOSE 9003

CMD ./main