FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update

# build go app
RUN go mod download
RUN go build -o crm-warehouse ./cmd/main.go

RUN chmod +x crm-warehouse

CMD ["./crm-warehouse"]