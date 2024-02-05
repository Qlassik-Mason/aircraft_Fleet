FROM golang:1.21.6

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "app/bin" ]


