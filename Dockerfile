# Step 1: Build the Go binary
FROM golang:1.18-alpine as build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .

# Step 2: Create the final image with the binary
FROM alpine:latest

WORKDIR /app/

COPY --from=build app/main .

EXPOSE 8080

CMD ["./main"]
