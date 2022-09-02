FROM golang:1.18.4-alpine as builder
RUN apk add --no-cache make git
WORKDIR /crack-src
COPY . /crack-src
RUN go mod download && \
    make docker && \
    mv ./bin/crack-docker /crack

FROM alpine:latest
COPY --from=builder /crack /

ENTRYPOINT ["/crack"]