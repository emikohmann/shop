FROM golang:1.18

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go mod download

WORKDIR /app/cmd/items-api

RUN go build -o /users-api

EXPOSE 8081

CMD [ "/users-api" ]
