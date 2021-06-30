FROM golang:alpine AS builder

LABEL maintainer="ankur@opennetworking.org"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

#ENV PORT 6000

RUN go build dummyamf.go

FROM alpine

WORKDIR /amf

COPY --from=builder /app/* /amf/

#RUN apk add --no-cache bash

#RUN chmod +x /pod/client.sh

#CMD ["./client.sh"]

CMD ["./dummyamf"]
