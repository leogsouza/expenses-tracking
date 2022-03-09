FROM golang:alpine as build

RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/tracking



FROM alpine:latest
WORKDIR /app
COPY --from=build /go/bin .
COPY --from=build /app/db/migrations /app/db/migrations

EXPOSE 8888

CMD ["./tracking"]