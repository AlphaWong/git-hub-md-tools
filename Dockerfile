FROM golang:1.9-alpine
RUN apk update
RUN apk add git
RUN mkdir -p /go/src/github.com/git-hub-md-tools
WORKDIR /go/src/github.com/git-hub-md-tools
COPY main.go .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/git-hub-md-tools/main .
CMD ["./main"] 
