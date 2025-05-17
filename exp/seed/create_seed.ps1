#!/usr/bin/env pwsh

# Stop the script when a cmdlet or a native command fails
$ErrorActionPreference = 'Stop'
$PSNativeCommandUseErrorActionPreference = $true

function Get-Master-Passphrase() {
  $password = Read-Host "Please enter a passphrase for the password database" -MaskInput
  $confirmation = Read-Host "Confirm passphrase" -MaskInput

  if ($password -ne $confirmation) {
    Write-Host "FATAL: Passphrase does not match confirmation!"
    exit 1
  }

  return $password
}

$dest = git rev-parse --show-toplevel | split-path -parent
$dest = join-path -path $dest -childpath db.kdbx

Write-Host "This script will create a KeePassXC database at $($dest) and store a master seed in that database."

if (Test-Path -Path $dest) {
  Write-Host "FATAL: a database already exists at $dest. Please delete this file to continue."
  exit 1
}

$passphrase = Get-Master-Passphrase

Write-Host "Creating a database, this may take some time..."
"$passphrase`n$passphrase" | & keepassxc-cli db-create -q -p -t 3000 $dest *> $null
Write-Host "Created!"

$seed = [Convert]::ToBase64String([System.Security.Cryptography.RandomNumberGenerator]::GetBytes(32))

"$passphrase`n$seed" | & keepassxc-cli add $dest "/MASTER_SEED" -p

remove-variable -name passphrase
remove-variable -name seed
