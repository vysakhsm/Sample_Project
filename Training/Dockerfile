FROM golang:latest
WORKDIR /Project
COPY Calculator.go .
RUN go build -o Application .
EXPOSE 8080
CMD ["./Application"]
