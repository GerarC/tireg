FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/tireg ./cmd/tireg

FROM alpine:3.20
RUN adduser -D -u 10001 appuser
COPY --from=builder /out/tireg /usr/local/bin/tireg
USER appuser
EXPOSE 8080
ENTRYPOINT ["tireg"]
