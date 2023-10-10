FROM golang:1.21.2 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /stock-price-api

FROM golang:1.21.2

COPY --from=build /stock-price-api /stock-price-api

EXPOSE 8080

CMD [ "/stock-price-api" ]
