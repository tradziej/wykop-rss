FROM golang:1.11-alpine as base

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM scratch
COPY --from=base /app/main /wykop-rss
COPY --from=base /app/html ./html
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 9001
CMD ["/wykop-rss"]