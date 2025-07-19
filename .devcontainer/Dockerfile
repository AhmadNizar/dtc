FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Set the Current Working Directory inside the container
WORKDIR /workspaces

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# Install PostgreSQL client and go-migrate
RUN apt-get update && \
    apt-get install -y postgresql-client postgresql-contrib rabbitmq-server

RUN apt-get update && apt-get install -y \
    curl \
    tar \
    sudo

# Download the latest Go Migrate binary
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz

# Extract the binary
RUN tar xzf migrate.tar.gz

# Move the binary to /usr/local/bin
RUN mv migrate /usr/local/bin/

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of your application
COPY . .

# Expose the port the app runs on (adjust as needed)
EXPOSE 8080
