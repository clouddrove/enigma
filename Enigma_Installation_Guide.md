
## Installation Guide

Enigma can be installed using the install script provided. This script works for both Linux and macOS platforms, and it supports common shell environments like **bash** and **zsh**.

### Prerequisites

Before you begin, ensure that you have the following installed on your system:

- `curl`: This is required to download the binary.
- `unzip`: Required to extract the downloaded binary from the zip file.

You can install these tools as follows:

- For **Ubuntu/Linux**:
  ```bash
  sudo apt-get install curl unzip
  ```

- For **macOS** (with Homebrew):
  ```bash
  brew install curl unzip
  ```

### Installing Enigma

To install Enigma, run the following command in your terminal:

```bash
curl -L https://raw.githubusercontent.com/clouddrove/enigma/main/install.sh | bash
```

This will automatically:

1. Detect your operating system (Linux or macOS) and architecture (amd64, arm64).
2. Download the appropriate Enigma binary from the latest release on GitHub.
3. Extract the binary from the zip file and install it to `/usr/local/bin/` with the necessary permissions.

You can also specify a specific version during installation:

```bash
curl -L https://raw.githubusercontent.com/clouddrove/enigma/main/install.sh | bash -s v0.0.15
```

### Manual Installation (if required)

If for any reason you prefer to manually build and install Enigma, you can follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/clouddrove/enigma.git
   cd enigma
   ```

2. Build the project using Go:
   ```bash
   go build -o enigma main.go
   ```

3. Move the binary to `/usr/local/bin`:
   ```bash
   sudo mv enigma /usr/local/bin/enigma
   ```

### Shell Compatibility

This install script is compatible with both **bash** and **zsh** shells.

- **Bash**: Bash is the default shell for most Linux distributions and macOS systems. The script works perfectly in this environment.
  
- **Zsh**: Zsh is the default shell for newer macOS versions. The script is also fully compatible with zsh, ensuring users can install Enigma in both environments without issues.

To confirm your shell, you can run:
```bash
echo $SHELL
```

If you're using **bash** or **zsh**, the install script will work seamlessly.

### Verifying Installation

After installation, you can verify that Enigma was installed successfully by running:

```bash
enigma --help
```

This command should display Enigma's available commands and flags. If this works, Enigma has been installed correctly!

---

If you encounter any issues, feel free to open an issue on [GitHub](https://github.com/clouddrove/enigma/issues).
