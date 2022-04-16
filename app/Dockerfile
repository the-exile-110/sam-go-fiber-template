FROM public.ecr.aws/bitnami/golang:1.18

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app main.go

CMD ["./app"]