FROM golang:1.19.3-alpine3.16
WORKDIR /app

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY go.mod .
COPY go.sum .
COPY config.yml .
RUN go mod download
RUN apk update && apk add --no-cache git

RUN GOOS=linux GOARCH=amd64 go build -o api .

ENTRYPOINT ["./api"]