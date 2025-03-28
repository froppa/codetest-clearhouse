FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o service ./main.go

FROM scratch

COPY --from=build /app/service /service
COPY --from=build /app/config/base.yml /config/base.yml

EXPOSE 8081

CMD ["./service"]
