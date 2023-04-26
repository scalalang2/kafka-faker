FROM golang:1.20-alpine

WORKDIR /src
COPY . /src

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add --no-cache git gcc librdkafka-dev libc-dev
RUN go mod download
RUN go build --tags musl -o /src/kafka-faker .

CMD ["/src/kafka-faker"]
