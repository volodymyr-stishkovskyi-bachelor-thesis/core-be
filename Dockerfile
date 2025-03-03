FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o core-be ./cmd/main.go

EXPOSE 8080

CMD ["/app/core-be"]