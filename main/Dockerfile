FROM golang:1.14-alpine
ENV VERSION=1.0.0
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
