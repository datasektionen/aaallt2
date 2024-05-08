FROM golang:1.22-alpine3.19 AS build

WORKDIR /src

COPY . .

RUN go build -o /aaallt2

FROM alpine:3.19

COPY --from=build /aaallt2 /aaallt2

CMD ["/aaallt2"]
