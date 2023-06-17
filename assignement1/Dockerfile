FROM golang:alpine AS builder

WORKDIR /channel
COPY . .
RUN go mod vendor
RUN go mod tidy
RUN go build -o build/bin -v .

FROM alpine
WORKDIR /app


COPY --from=builder /channel/build/bin .

CMD [ "/app/bin" ]

