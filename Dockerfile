FROM golang:1.23

WORKDIR /go/src/app

COPY . .

RUN go build -o enigma main.go

ENTRYPOINT ["/go/src/app/enigma"]