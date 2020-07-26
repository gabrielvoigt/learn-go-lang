FROM golang:1.14-alpine
LABEL maintainer="Gabriel Voigt <gabrielvoigt@gmail.com>"
WORKDIR /go/src/
COPY cmd .
RUN GOOS=linux go build main.go
CMD ["./main"]