FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /client cmd/client/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /server /server
COPY --from=builder /client /client

EXPOSE 8080