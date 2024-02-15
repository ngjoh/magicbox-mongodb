<#---
title: List
connection: sharepoint
tag: list
api: post
output: bloblist.json
---


#>

$result = "$env:WORKDIR/bloblist.json"
$output = az storage blob list --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER   #--file $env:WORKDIR/backupcontent.json --name backupcontent.json

$output | Out-File $result 
