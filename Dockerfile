FROM golang:latest

WORKDIR /gses2/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./test/...
RUN go build -o /btcApp ./cmd/...

CMD ["/btcApp"]