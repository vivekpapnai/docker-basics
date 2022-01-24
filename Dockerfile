FROM golang:1.16-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-example

EXPOSE 8080

CMD [ "/docker-example" ]