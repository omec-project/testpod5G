FROM golang:alpine AS builder

LABEL maintainer="ankur@opennetworking.org"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

#ENV PORT 6000

RUN go build testpod.go

FROM alpine

#RUN apk update

WORKDIR /nf

COPY --from=builder /app/* /nf/

#ENTRYPOINT [ "./testpod" ]

#CMD ["smf"]
