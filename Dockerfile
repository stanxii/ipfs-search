FROM golang:1.11-alpine AS build

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# This results in a single layer image
FROM alpine
COPY wait-for /usr/local/bin/wait-for
COPY --from=build /go/bin/ipfs-search /usr/local/bin/ipfs-search

CMD ["crawl"]
# TODO: Replace by decent run script
ENTRYPOINT ["wait-for", "ipfs:5001", "--", "wait-for", "elasticsearch:9200", "--", "wait-for", "rabbitmq:5672", "--", "/usr/local/bin/ipfs-search"]
