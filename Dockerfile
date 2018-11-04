FROM golang:1.11

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["crawl"]
ENTRYPOINT ["/app/wait-for-it.sh", "ipfs:5001", "--", "/app/wait-for-it.sh", "elasticsearch:9200", "--", "/app/wait-for-it.sh", "rabbitmq:5672", "--", "ipfs-search"]
