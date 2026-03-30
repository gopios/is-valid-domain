#!/bin/bash

# is-valid-domain installer script
# Downloads and installs ivd binary to /usr/local/bin

set -e

VERSION=${1:-latest}
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="ivd"

echo "Installing is-valid-domain (ivd) version: $VERSION"

# Detect architecture
ARCH=$(uname -m)
OS=$(uname -s)

case $OS in
    Linux*)
        OS="linux"
        ;;
    Darwin*)
        OS="darwin"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    armv7l)
        ARCH="arm"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Determine version
if [ "$VERSION" = "latest" ]; then
    # Get latest release tag from GitHub API
    VERSION=$(curl -s https://api.github.com/repos/gopios/is-valid-domain/releases/latest | grep '"tag_name"' | sed -E 's/.*"tag_name": ?"v?([^"]+).*/\1/')
    if [ -z "$VERSION" ]; then
        echo "Failed to get latest version"
        exit 1
    fi
fi

BINARY="${BINARY_NAME}-${VERSION}-${OS}-${ARCH}"
if [ "$OS" = "darwin" ] && [ "$ARCH" = "arm" ]; then
    BINARY="${BINARY_NAME}-${VERSION}-${OS}-arm64"
fi

echo "Downloading: $BINARY"

# Download the binary
DOWNLOAD_URL="https://github.com/gopios/is-valid-domain/releases/download/v${VERSION}/${BINARY}"
echo "From: $DOWNLOAD_URL"

if ! curl -L "$DOWNLOAD_URL" -o "/tmp/$BINARY"; then
    echo "Failed to download binary"
    exit 1
fi

# Make executable
chmod +x "/tmp/$BINARY"

# Check if install directory exists and is writable
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating install directory: $INSTALL_DIR"
    sudo mkdir -p "$INSTALL_DIR"
fi

if [ ! -w "$INSTALL_DIR" ]; then
    echo "Need sudo to install to $INSTALL_DIR"
    sudo mv "/tmp/$BINARY" "$INSTALL_DIR/$BINARY_NAME"
else
    mv "/tmp/$BINARY" "$INSTALL_DIR/$BINARY_NAME"
fi

# Create symlink without version
if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "Installed to: $INSTALL_DIR/$BINARY_NAME"
    echo "Testing installation..."
    
    if "$INSTALL_DIR/$BINARY_NAME" --version 2>/dev/null || "$INSTALL_DIR/$BINARY_NAME" example.com >/dev/null 2>&1; then
        echo "✅ Installation successful!"
        echo "Usage: $BINARY_NAME example.com"
    else
        echo "❌ Installation test failed"
        exit 1
    fi
else
    echo "❌ Installation failed"
    exit 1
fi

# Clean up
rm -f "/tmp/$BINARY"

echo "Done!"
