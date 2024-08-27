FROM golang:1.23

# Install dependencies
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# Add Docker’s official GPG key
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Set up the Docker repository
RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install specific Docker version
RUN apt-get update && apt-get install -y docker-ce=5:27.0.3~3-0~debian-buster

# Verify Docker version
RUN docker --version

WORKDIR /go/src/app
COPY . .
RUN go build -o enigma main.go

ENTRYPOINT ["/go/src/app/entrypoint.sh"]
