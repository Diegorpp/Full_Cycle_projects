FROM golang as builder

COPY . /app

WORKDIR /app

RUN go build hello.go

FROM alpine:latest

COPY --from=builder /app/hello .

CMD [ "./hello" ]