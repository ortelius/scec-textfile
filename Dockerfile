FROM cgr.dev/chainguard/go@sha256:1272cf2beb5d088abe4904da6cc587dbf9011aecfa61939d476a13be75c8ebe3 AS builder

WORKDIR /app
COPY . /app

RUN go mod tidy; \
    go build -o main .

FROM cgr.dev/chainguard/glibc-dynamic@sha256:5915b6b2c9ff77916fc534d9c0676eaaf776964f674e4a6aeb6b43426c7db79a

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs docs

ENV ARANGO_HOST localhost
ENV ARANGO_USER root
ENV ARANGO_PASS rootpassword
ENV ARANGO_PORT 8529
ENV MS_PORT 8080

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
