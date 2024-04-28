FROM golang:latest AS build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app .
RUN apk --no-cache add tzdata

CMD [ "./app" ]

EXPOSE 8080