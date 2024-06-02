FROM golang:alpine

ENV GOOS=linux
ENV GOARCH=amd64

COPY cfs.go .

RUN go build -o cfs cfs.go

CMD ["./cfs"]
