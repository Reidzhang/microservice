From golang:1.9.0 as builder

WORKDIR /go/src/github.com/Reidzhang/microservice/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo database.go handler.go repository.go token_service.go main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/Reidzhang/microservice/user-service/user-service .

CMD ["./user-service"]