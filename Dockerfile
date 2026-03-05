FROM golang:1.25-alpine

RUN apk add --no-cache build-base ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/nitewatch ./cmd/nitewatch

EXPOSE 8080

ENTRYPOINT ["nitewatch"]
CMD ["worker"]
