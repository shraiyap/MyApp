FROM golang:1.22-alpine
WORKDIR /go-app
COPY . .
RUN go mod tidy
RUN go build -o myapp
EXPOSE 8080
CMD ["./myapp"]