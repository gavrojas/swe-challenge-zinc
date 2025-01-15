FROM golang:1.21.13-alpine

RUN mkdir app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ARG PORT=:8080 \
# ARG PORT=:443 \
    GIN_MODE=release

ENV PORT=${PORT} \
    GIN_MODE=${GIN_MODE}

COPY . .
# COPY fullchain.pem /etc/ssl/certs/ 
# COPY privkey.pem /etc/ssl/private/

RUN go build -o main .

CMD ["./main"]