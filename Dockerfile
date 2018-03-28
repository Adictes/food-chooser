FROM golang:alpine as builder
WORKDIR /go/src/github.com/Adictes/food-chooser
COPY . .
RUN go build main.go

FROM alpine:latest
COPY --from=builder /go/src/github.com/Adictes/food-chooser .
EXPOSE 8080
ENTRYPOINT [ "./main" ]
