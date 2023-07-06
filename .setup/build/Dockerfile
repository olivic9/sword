# build image
FROM golang-1.19 AS builder

WORKDIR /usr/src/app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

RUN export CGO_ENABLED=0

RUN make build/apps

# final image
FROM alpine:3.18.2
CMD ["api"]
