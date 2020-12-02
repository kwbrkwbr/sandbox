FROM golang:1.15.2 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY account.json .

COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -mod=readonly -v -o server

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server
COPY --from=builder /app/account.json /app

ENV GOOGLE_APPLICATION_CREDENTIALS="/app/account.json"
# Run the web service on container startup.
CMD ["/app/server"]