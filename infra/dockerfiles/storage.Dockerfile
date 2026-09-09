FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.work go.work.sum* ./
COPY packages/go-api-lib ./packages/go-api-lib
COPY services/storage ./services/storage
RUN cd services/storage && go build -o /bin/storage ./cmd/storage

FROM alpine:3.20
COPY --from=builder /bin/storage /bin/storage
EXPOSE 8082
CMD ["/bin/storage"]
