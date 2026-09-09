FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.work go.work.sum* ./
COPY packages/go-api-lib ./packages/go-api-lib
COPY services/auth ./services/auth
RUN cd services/auth && go build -o /bin/auth ./cmd/auth

FROM alpine:3.20
COPY --from=builder /bin/auth /bin/auth
EXPOSE 8081
CMD ["/bin/auth"]
