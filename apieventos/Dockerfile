FROM golang:1.23-alpine

WORKDIR /app

COPY . .


RUN go mod tidy
RUN go build -o server .


EXPOSE 8080


CMD ["./server"]
