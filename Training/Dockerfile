FROM golang:latest
WORKDIR /Project
RUN git clone https://github.com/vysakhsm/Sample_Project.git
RUN go build -o .
EXPOSE 8080
CMD [./]
