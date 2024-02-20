<#---
title: List backup blobs
connection: sharepoint
tag: list
api: post
output: bloblist.json
---


#>

$result = "$env:WORKDIR/bloblist.json"
$output = az storage blob list --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER   #--file $env:WORKDIR/backupcontent.json --name backupcontent.json
$raw = $output | ConvertFrom-Json 
$blobs = @()
foreach ($blob in $raw) {
    $blobs  += @{
        name = $blob.name
        lastModified = $blob.properties.lastModified
        size = $blob.properties.contentLength
    }
}

$blobs | ConvertTo-Json -Depth 10 |  Out-File $result -Encoding utf8NoBOM
