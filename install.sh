#!/bin/bash

OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH_TYPE=$(uname -m)

if [ "$ARCH_TYPE" == "x86_64" ]; then
    BIN_NAME="paping-linux-amd64"
elif [ "$ARCH_TYPE" == "aarch64" ] || [ "$ARCH_TYPE" == "arm64" ]; then
    BIN_NAME="paping-arm64"
else
    echo "Architecture non supportée ($ARCH_TYPE)."
    exit 1
fi

echo "Téléchargement de paping..."
URL="https://raw.githubusercontent.com/axelhc2/paping/main/bin/$BIN_NAME"

sudo curl -L $URL -o /usr/bin/paping
sudo chmod +x /usr/bin/paping

echo "Succès ! Tape 'paping' pour commencer. Credit: Axel Chetail (infrawire.fr)"
