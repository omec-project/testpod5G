# SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
# SPDX-License-Identifier: Apache-2.0
FROM golang:alpine AS builder

LABEL maintainer="ankur@opennetworking.org"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 6000

RUN go build dummysmf.go

FROM alpine

#RUN apk update

WORKDIR /smf

COPY --from=builder /app/* /smf/

CMD ["./dummyudm"]
