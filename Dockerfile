FROM golang:1.11

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["crawl"]
ENTRYPOINT ["ipfs-search"]
