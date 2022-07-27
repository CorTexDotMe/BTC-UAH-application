FROM golang:latest

WORKDIR /gses2/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /btcApp btcApp

CMD ["/btcApp"]