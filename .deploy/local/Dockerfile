FROM golang:1.18

WORKDIR /var/www/html/dysn/gateway

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

CMD ["go","run","main.go"]