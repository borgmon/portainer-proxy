FROM golang:1.18-alpine
WORKDIR /app
RUN apk add git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go","run","."]
