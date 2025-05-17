#!/usr/bin/env bash

# Exit on error, on undefined variable, when a pipeline fails
set -euo pipefail

dest="$(pwd)/db.kdbx"
echo "This script will create a KeePassXC database at $dest and store a master seed in this database."

if [[ -f "$dest" ]]; then
  echo "FATAL: File $dest exists! Please delete this file to continue."
  exit 1
fi

printf "Please enter a passphrase for the password database (will not be echoed): "
read -s passphrase
echo

echo -n "Confirm passphrase: "
read -s confirmation
echo

if [[ "$passphrase" != "$confirmation" ]]; then
  echo "FATAL: Passphrase does not match confirmation!"
  exit 1
fi

echo "Creating a KeePassXC database, this may take some time..."
printf "$passphrase\n$passphrase" | keepassxc-cli db-create -q -p -t 3000 $dest >/dev/null 2>&1
echo "Created database. Generating seed..."
seed="$(openssl rand -base64 32)"

printf "$passphrase\n$seed" | keepassxc-cli add $dest "/MASTER_SEED" -p >/dev/null 2>&1
