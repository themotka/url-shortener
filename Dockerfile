FROM golang:1.22-alpine AS builder

WORKDIR /usr/docker/src

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

RUN go build -o bin/app cmd/app/main.go

FROM alpine AS runner

COPY --from=builder usr/docker/src/bin/app .
COPY internal/configs /internal/configs
COPY migrations /migrations

CMD ["/app", "-d"]