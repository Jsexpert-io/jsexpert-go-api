# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.26 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o myapp specifies the output name of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official lightweight Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/
FROM alpine:3.14
WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp .
EXPOSE 8080
# Run the web service on container startup.
CMD ["./myapp"]
