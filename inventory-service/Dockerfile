FROM golang:1.20

WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -o api ./cmd

EXPOSE 8080

CMD ["./api"]