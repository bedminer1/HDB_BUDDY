FROM golang:1.23.5-alpine

RUN apk update && apk add --no-cache sqlite-dev gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o server

EXPOSE 4000

CMD ["./server"]