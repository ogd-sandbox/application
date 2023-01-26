FROM golang:1.18-alpine AS build-env
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o application ./app/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /build/application .
EXPOSE 3333
CMD [ "./application" ]