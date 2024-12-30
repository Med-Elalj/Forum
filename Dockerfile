FROM golang:1.23.4-alpine3.21
WORKDIR /app
COPY . .

RUN apk add gcc
RUN apk add musl-dev

RUN go mod download
RUN go build -o forum-project .

CMD ["./forum-project"]