FROM golang:alpine as builder
WORKDIR /go/src/app
COPY go.mod .
COPY cmd/*.go .
COPY parser/* .
COPY *.go .
RUN go mod download
RUN CGO_ENABLED=0 go build .

FROM alpine:3.13.5
COPY --from=builder /go/src/app/opening_hours .
CMD ["./opening_hours"]
