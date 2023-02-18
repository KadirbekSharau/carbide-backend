# Start from the latest golang base image.
FROM golang:latest

# Add maintainer information
LABEL maintainer="sharaukadr2001@gmail.com"

# Set the current working directory inside an image.
WORKDIR /app

# Copy Go module dependency requirements file.
COPY go.mod .

# Copy Go Modules expected hashes file.
COPY go.sum .

# Download dependencies.
RUN go mod download

# Copy all sources.
COPY . .

# Build the application.
RUN go build -o /apprentice

# Delete source files.
RUN find . -name "*.go" -type f -delete

# Run the application.
CMD ["/apprentice"]