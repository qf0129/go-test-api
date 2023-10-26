# Stage 1: compile the program
FROM golang:1.20-alpine AS builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
COPY --from=builder /app/server /server
CMD ["/server"]  