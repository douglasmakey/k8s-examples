FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]