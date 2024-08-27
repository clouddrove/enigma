# Use the official Golang image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the entire project into the container
COPY . .

# Install dependencies and build the Go application
RUN go build -o enigma main.go

# Install Docker CLI
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/debian \
    $(lsb_release -cs) \
    stable" && \
    apt-get update && apt-get install -y docker-ce-cli

# Set the entry point for the container
ENTRYPOINT ["/go/src/app/enigma"]
