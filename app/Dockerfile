FROM golang:1.18 as build-env

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /

CMD ["/app"]