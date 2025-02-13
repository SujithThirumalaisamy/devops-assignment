FROM golang:1.23.0-alpine AS base

WORKDIR /app

COPY . .

RUN go build -o main .

FROM base AS runner

WORKDIR /app

COPY --from=base /app/main /app/main

EXPOSE 8080

CMD ["./main"]
