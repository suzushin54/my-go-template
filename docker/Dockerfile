FROM golang:1.12.6

WORKDIR /go/src

# for hot-reload
RUN go get github.com/pilu/fresh

ENV GO111MODULE=on

EXPOSE 8083

CMD ["fresh"]
