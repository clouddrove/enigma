FROM golang:1.23 as builder

# Install dependencies and CLI tools in a single layer to minimize build time
RUN apt-get update && apt-get install -y --no-install-recommends \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    bash \
    unzip && \
    # Install Docker CLI
    curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" \
    | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
    apt-get update && apt-get install -y --no-install-recommends docker-ce-cli && \
    # Install AWS CLI
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && ./aws/install && rm -rf aws awscliv2.zip && \
    # Install GCP CLI
    curl -o /tmp/google-cloud-sdk.tar.gz https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz && \
    mkdir -p /usr/local/gcloud && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz && \
    /usr/local/gcloud/google-cloud-sdk/install.sh && \
    # Install Azure CLI
    curl -sL https://aka.ms/InstallAzureCLIDeb | bash && \
    # Clean up to reduce image size
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*

# Set environment variables for GCP CLI
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin

# Set up Go environment
WORKDIR /go/src/app
COPY . .

# Build Go application in builder stage
RUN go build -o enigma main.go

# Create a minimal final image without build dependencies
FROM golang:1.23 as runner
WORKDIR /go/src/app

# Copy the binary from the builder stage
COPY --from=builder /go/src/app/enigma /go/src/app/enigma
COPY --from=builder /go/src/app/entrypoint.sh /go/src/app/entrypoint.sh

# Add executable permissions to the entrypoint script
RUN chmod +x /go/src/app/entrypoint.sh

# Set entrypoint to the shell script
ENTRYPOINT ["/go/src/app/entrypoint.sh"]
