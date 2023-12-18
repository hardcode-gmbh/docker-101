FROM golang:alpine AS builder

WORKDIR /usr/src/app

# Separate dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY main.go .

# Build it
RUN go build -v  ./...




FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache fortune

#RUN adduser user
#USER user

# Document env + ports
ENV DEFAULT_NAME=Peter
EXPOSE 1234

# Add healtcheck
HEALTHCHECK --interval=10s --timeout=5s --start-period=3s --retries=3 CMD wget http://localhost:1234/health -q -O -

# Copy build result from first stage
COPY --from=builder /usr/src/app/hello-docker-go /hello-docker-go

# Define start command
CMD ["/hello-docker-go"]