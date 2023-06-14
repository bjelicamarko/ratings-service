FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./RatingsService

EXPOSE 8084

CMD [ "./RatingsService" ] 