#!/bin/bash

# Enigma Install Script
# This script installs the latest version of Enigma on Linux or macOS systems.

set -e  # Exit immediately if a command exits with a non-zero status.

# Function to display error messages
error_exit() {
    echo ""
    echo "🚫 Error: $1"
    exit 1
}

echo "🔍 Detecting operating system..."

# Detect OS
OS=$(uname -s)

case "$OS" in
    Linux*)     OS_TYPE="Linux";;
    Darwin*)    OS_TYPE="macOS";;
    *)          error_exit "Unsupported OS: $OS";;
esac

echo "✅ Detected OS: $OS_TYPE"

echo "🔍 Detecting architecture..."

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64|amd64) ARCH_TYPE="amd64";;
    arm64|aarch64) ARCH_TYPE="arm64";;
    *) error_exit "Unsupported architecture: $ARCH";;
esac

echo "✅ Detected architecture: $ARCH_TYPE"

echo "🌐 Fetching the latest Enigma version..."

# Allow the user to specify a version, otherwise fetch the latest version from GitHub API
VERSION=${1:-$(curl -s https://api.github.com/repos/clouddrove/enigma/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')}

if [ -z "$VERSION" ]; then
    error_exit "Failed to fetch the latest version."
fi

echo "✅ Using Enigma version: $VERSION"

echo "🔗 Constructing download URL..."

# Construct download URL based on OS and architecture
if [ "$OS_TYPE" = "Linux" ]; then
    BINARY_NAME="enigma-linux-$ARCH_TYPE.zip"
elif [ "$OS_TYPE" = "macOS" ]; then
    BINARY_NAME="enigma-darwin-$ARCH_TYPE.zip"
fi

URL="https://github.com/clouddrove/enigma/releases/download/$VERSION/$BINARY_NAME"

echo "✅ Download URL: $URL"

echo "⬇️  Downloading Enigma binary..."

# Download the zip file
if curl --output /dev/null --silent --head --fail "$URL"; then
    curl -L "$URL" -o enigma.zip || error_exit "Failed to download Enigma binary."
else
    error_exit "Binary for $OS_TYPE ($ARCH_TYPE) not found at $URL."
fi

# Verify download
if [ ! -f "enigma.zip" ]; then
    error_exit "Download failed or file not found."
fi

echo "✅ Downloaded Enigma binary zip."

echo "📦 Extracting Enigma binary..."

# Unzip the file
unzip enigma.zip || error_exit "Failed to extract Enigma binary."

# Remove the zip file after extraction
rm enigma.zip

# Check if the binary was extracted
if [ ! -f "enigma" ]; then
    error_exit "Enigma binary not found after extraction."
fi

echo "✅ Extracted Enigma binary."

echo "🔑 Setting executable permissions..."

# Make binary executable
chmod +x enigma

echo "🚚 Installing Enigma to /usr/local/bin..."

# Move binary to /usr/local/bin
if [ -w "/usr/local/bin" ]; then
    mv enigma /usr/local/bin/enigma
else
    sudo mv enigma /usr/local/bin/enigma || error_exit "Failed to move Enigma binary to /usr/local/bin."
fi

echo "✅ Enigma installed successfully!"

echo "🧪 Verifying installation..."

# Verify installation using 'enigma --help'
echo "🚀 Verifying Enigma installation with 'enigma --help':"
enigma --help || error_exit "Failed to run Enigma."

echo "🎉 Installation completed successfully!"

exit 0
