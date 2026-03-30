#!/bin/bash

set -e

# Version from tag
VERSION=$(git describe --tags --exact-match 2>/dev/null || echo "dev")
BINARY_NAME="ivd"
OUTPUT_DIR="releases"

# Create output directory
mkdir -p $OUTPUT_DIR

# Clean previous builds
rm -f $OUTPUT_DIR/${BINARY_NAME}-*

echo "Building release binaries for version $VERSION..."

# Build for different platforms
platforms=(
    "linux/amd64"
    "linux/arm64" 
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/arm64"
)

for platform in "${platforms[@]}"; do
    GOOS=$(echo $platform | cut -d'/' -f1)
    GOARCH=$(echo $platform | cut -d'/' -f2)
    
    output_name="${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        output_name="${output_name}.exe"
    fi
    
    echo "Building for $GOOS/$GOARCH..."
    
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_DIR/$output_name ./cmd/ivd
    
    # Create compressed archives
    if [ $GOOS = "windows" ]; then
        cd $OUTPUT_DIR
        zip -q ${output_name%.exe}.zip $output_name
        rm $output_name
        cd ..
    else
        cd $OUTPUT_DIR
        tar -czf ${output_name}.tar.gz $output_name
        rm $output_name
        cd ..
    fi
done

echo "Release binaries created in $OUTPUT_DIR/"
ls -la $OUTPUT_DIR/
