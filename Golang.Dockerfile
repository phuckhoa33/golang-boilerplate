# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy environment file
COPY --chown=node:node .env ./

COPY --chown=node:node . .

ARG APP_ENV

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download


CMD ["air", "-c", ".air.toml"]