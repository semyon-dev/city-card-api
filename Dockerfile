FROM golang:latest as builder
COPY go.mod /go/goods-scanner-api/
COPY go.sum /go/goods-scanner-api/
WORKDIR /go/goods-scanner-api
RUN go mod download
COPY . /go/goods-scanner-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main cmd/goodsScanner.go

FROM alpine:latest
COPY --from=builder /main ./
# ENV enviroment=production
# COPY --from=builder /go/common-email-service/production.env ./
ENV enviroment=dev
COPY --from=builder /go/goods-scanner-api/internal/dev.env ./
# COPY --from=builder /go/goods-scanner-api/internal/resources/migrations ./internal/resources/migrations
RUN chmod +x ./main
ENTRYPOINT ["./main"]

EXPOSE 8081


