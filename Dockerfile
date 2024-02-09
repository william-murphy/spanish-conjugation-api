FROM golang:latest
WORKDIR /app
COPY ./api/* .
RUN go mod download
RUN go build -o app
EXPOSE 8080
CMD ["./app"]