FROM golang:1.19 AS build

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . ./
RUN CGO_ENABLED=0 go build -o ./build/goapp ./cmd

FROM alpine

COPY --from=build /app/build/goapp /goapp
EXPOSE 9997 8888

CMD ["./goapp"]