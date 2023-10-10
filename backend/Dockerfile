# Start from the base Golang image
FROM golang:1.21-alpine3.17 as base

# Set the current working directory inside the container
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
   build-base \
    curl \
    tzdata \
    git \
    vips \
    vips-dev

RUN go install github.com/cosmtrek/air@latest

# Ensure the go.mod and go.sum files are being copied
COPY go.mod go.sum ./

COPY .env /app/.env

COPY . .

RUN go mod download && go mod verify

# DEV IMAGE
FROM base AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD air

#BUILDER IMAGE
FROM base AS BUILDER

WORKDIR /app

RUN go build -o ecomm

# PRODUCTION IMAGE
FROM alpine:latest as PRODUCTION

COPY --from=builder /app/ecomm /usr/local/bin/ecomm

EXPOSE 5000

ENTRYPOINT ["usr/local/bin/ecomm"]

# # Copy go mod and sum files to the workspace
# COPY go.mod go.sum ./

# COPY .env /app/.env

# # Download all dependencies
# RUN go mod download



# # Copy the source code into the container
# COPY . .

# COPY wait-for-it.sh /app/wait-for-it.sh
# RUN chmod +x /app/wait-for-it.sh

# RUN echo "/app/ directory listing"
# RUN ls -l /app/

# # Copy the start-backend.sh script
# COPY start-backend.sh /app/start-backend.sh
# RUN chmod +x /app/start-backend.sh

# # Build the application
# RUN go build -o main .

# # Expose port 8080 to the outside world
# EXPOSE 8080

# Run the executable
# CMD ["./main"]


# # Use the start-backend.sh script as the default command
# CMD ["./start-backend.sh"]
