
#stage 1
FROM golang:alpine AS builder

WORKDIR /test
RUN apk update

COPY . .

COPY go.mod go.sum /test/
RUN go mod tidy

RUN go build -o build/bin -v .

#stage 2

FROM alpine

WORKDIR /app

COPY --from=builder /test/build/bin .


EXPOSE 3030

CMD [ "/app/bin" ]