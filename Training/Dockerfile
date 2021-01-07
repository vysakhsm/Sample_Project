FROM golang:latest
WORKDIR /Project
COPY Calculator.go .
COPY AddPage.html .
RUN go build -o Application .
EXPOSE 8080
CMD ["./Application"]
