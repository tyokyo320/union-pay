# Start from the latest golang base image
FROM golang:alpine AS builder

# Add Maintainer Info
LABEL maintainer="tyokyo320 <contact@tyokyo320.com>"

# Set the Current Working Directory inside the container
WORKDIR /union-pay

# Copy go mod and sum files
COPY app.go .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo cmd/server/main.go -o app .

# RUN
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /union-pay/app .
EXPOSE 8080
CMD ["./app"]  




