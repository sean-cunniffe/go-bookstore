FROM golang:1.18-alpine as builder

ENV APP_HOME /go/src/go-bookstore

WORKDIR "$APP_HOME"
COPY src/ src/
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify
RUN go build -o bookstore src/cmd/main/main.go

FROM alpine
COPY --from=builder /go/src/go-bookstore/bookstore ./
ENTRYPOINT ["./bookstore"]