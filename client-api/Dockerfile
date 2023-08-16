FROM golang:1.20-alpine AS build

WORKDIR /
COPY . /

RUN go build -o ./client-api/cmd/main ./client-api/cmd/main.go

FROM scratch

WORKDIR /app
COPY --from=build /client-api/cmd/main .

ENTRYPOINT ["/app/main"]
