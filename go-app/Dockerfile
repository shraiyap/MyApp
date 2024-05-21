FROM golang:1.22-alpine
WORKDIR /go-app
COPY . .
RUN go mod tidy
RUN go build -o myapp
EXPOSE 9080
CMD ["./myapp"]