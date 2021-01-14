# Start from the latest golang base image
FROM golang:alpine AS builder

# Add Maintainer Info
LABEL maintainer="tyokyo320 <contact@tyokyo320.com>"

# Set the Current Working Directory inside the container
WORKDIR /union-pay

# Copy all files
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o app cmd/server/main.go

# RUN
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /union-pay/templates templates
COPY --from=builder /union-pay/assets assets
COPY --from=builder /union-pay/app .
EXPOSE 8080

RUN groupadd -r union_pay && useradd --no-log-init -r -g union_pay union_pay
USER union_pay
CMD ["./app"]  




