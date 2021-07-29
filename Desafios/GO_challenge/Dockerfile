FROM golang as builder

COPY . /app

WORKDIR /app

RUN go build -ldflags "-s -w" hello.go

FROM scratch

COPY --from=builder /app/hello .

CMD [ "./hello" ]