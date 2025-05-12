#!/usr/bin/env pwsh

$cold_pubkey = $Env:COLD_PUBLIC_KEY
$server_pubkey = "028234EE392897A902F48E2BBEADE662C587C7BB81F131C18B9DED156853DB4169"

$rpcconnect = "127.0.0.1:18443"
$rpcuser = "username"
$rpcpassword = "password"

$desc = "wsh(multi(2,$($cold_pubkey),$($server_pubkey)))"

function call_cmd($cmd) {
    $btc_cli = "bitcoin-cli -rpcconnect=""$($rpcconnect)"" -rpcpassword=$($rpcpassword) -rpcuser=$($rpcuser) -regtest"
    invoke-expression "$($btc_cli) $($cmd)"
}

$with_checksum = call_cmd("getdescriptorinfo ""$($desc)""") | convertfrom-json
write-host $with_checksum.descriptor

call_cmd "deriveaddresses ""$($with_checksum.descriptor)"""
