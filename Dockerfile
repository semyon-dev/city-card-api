FROM golang:latest as builder
COPY go.mod /go/city-card-api/
COPY go.sum /go/city-card-api/
WORKDIR /go/city-card-api
RUN go mod download
COPY . /go/city-card-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main cmd/main.go

FROM alpine:latest
COPY --from=builder /main ./
# ENV enviroment=dev
# COPY --from=builder /go/goods-scanner-api/internal/dev.env ./
# COPY --from=builder /go/goods-scanner-api/internal/resources/migrations ./internal/resources/migrations
RUN chmod +x ./main
ENTRYPOINT ["./main"]

EXPOSE 8081


