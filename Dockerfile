FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8787

CMD [ "./main" ]