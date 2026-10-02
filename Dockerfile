FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]