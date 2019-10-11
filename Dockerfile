FROM golang:1.13.1-alpine as base

LABEL maintainer="Rain <rainlay@gmail.com>"

RUN apk update
RUN apk add bash git
WORKDIR /go/src/app

# cache module
COPY go.mod .
COPY go.sum .

COPY . .
ENV GOOS=linux
ENV GOARCH=amd64
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

#### Put the binary onto base image
FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
ENV GIN_MODE=release
EXPOSE 8080
COPY --from=base /go/src/app/main .
#ENTRYPOINT ["./main"]
CMD ["./main"]
