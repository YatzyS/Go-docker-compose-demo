FROM golang:1.15-alpine3.12
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .
EXPOSE 9090
CMD ["/app/main"]