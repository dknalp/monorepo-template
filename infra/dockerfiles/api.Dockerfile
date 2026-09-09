FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.work go.work.sum* ./
COPY packages/go-api-lib ./packages/go-api-lib
COPY services/api ./services/api
RUN cd services/api && go build -o /bin/api ./cmd/api

FROM alpine:3.20
COPY --from=builder /bin/api /bin/api
EXPOSE 8080
CMD ["/bin/api"]
