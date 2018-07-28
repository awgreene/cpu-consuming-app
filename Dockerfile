FROM golang:1.10.3

LABEL maintainer="agreene@redhat.com"

EXPOSE 8000

COPY . /app

RUN go get -u github.com/gorilla/mux

CMD go run /app/main.go
