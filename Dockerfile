FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o go-ecoflow-rest-api .

FROM alpine:edge
WORKDIR /app
COPY --from=builder /app/go-ecoflow-rest-api .
RUN apk --no-cache add ca-certificates tzdata
EXPOSE 8080

ENTRYPOINT ["/app/go-ecoflow-rest-api"]