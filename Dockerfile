From golang:1.24

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

CMD ["./main"]