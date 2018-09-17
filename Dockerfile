FROM golang:1.11-alpine as base
WORKDIR /go/src/github.com/tradziej/wykop-rss
COPY . .

RUN apk --no-cache add git
RUN go get -v ./...
RUN apk --update add ca-certificates
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM scratch
COPY --from=base /go/src/github.com/tradziej/wykop-rss/main /wykop-rss
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 9001
CMD ["/wykop-rss"]