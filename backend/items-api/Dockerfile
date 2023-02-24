FROM golang:1.18

WORKDIR /items-api

COPY . .

RUN go mod download

WORKDIR /items-api/cmd/items-api

RUN go build -o items-api

EXPOSE 8080

CMD [ "/items-api" ]