FROM golang:1.22-alpine3.19 AS build

COPY . .

RUN go build -o /aaaallltt

FROM alpine:3.19

COPY --from=build /aaaallltt /aaaallltt

CMD ["/aaaallltt"]
