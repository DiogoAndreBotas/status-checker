FROM golang:1.21.3-alpine
LABEL maintainer="Diogo Botas"

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8082

CMD ["./main"]
