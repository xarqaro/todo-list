FROM golang:1.27
WORKDIR /usr/src/app
COPY . .
RUN go mod tidy