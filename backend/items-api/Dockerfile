FROM golang:1.18

COPY . .

WORKDIR /cmd/items-api

RUN go mod download
RUN go build -o /items-api

EXPOSE 8080

CMD [ "/items-api" ]