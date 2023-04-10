FROM golang:1.19-alpine AS build

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 G8OOS=linux go build -o /app


FROM alpine:3.16

WORKDIR /

COPY --from=build /app /app

ENTRYPOINT [ "/app" ]