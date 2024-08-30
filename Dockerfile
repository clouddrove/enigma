FROM golang:1.23

# Install Docker CLI and other dependencies
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    unzip

# Add Docker's official GPG key
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Set up the Docker repository
RUN echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
    $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker CE CLI
RUN apt-get update && apt-get install -y docker-ce-cli

# Install AWS CLI v2
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf aws awscliv2.zip

# Install Google Cloud CLI
RUN curl -O https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz && \
    tar -xzf google-cloud-sdk.tar.gz && \
    ./google-cloud-sdk/install.sh --quiet && \
    rm google-cloud-sdk.tar.gz

# Add Google Cloud SDK to PATH
ENV PATH $PATH:/google-cloud-sdk/bin

WORKDIR /go/src/app
COPY . .
RUN go build -o enigma main.go

ENTRYPOINT ["/go/src/app/entrypoint.sh"]
