FROM golang:1.18-alpine

WORKDIR /

COPY main .
COPY .env .
CMD ["/main"]