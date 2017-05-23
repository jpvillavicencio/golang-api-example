FROM golang
MAINTAINER jpvillavicencio

WORKDIR /go/src/app
ADD . .
EXPOSE 8080
RUN go-wrapper download
CMD go run main.go
