FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go build -o enigma main.go

ENTRYPOINT ["/go/src/app/enigma"]