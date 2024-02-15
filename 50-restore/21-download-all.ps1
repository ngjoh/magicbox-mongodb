<#---
title: Download all backups
connection: sharepoint
input: databaseservices.json
api: post
tag: all
---#>

if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
}
$destinationDir = "$env:WORKDIR/download"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}
$databases = Get-Content -Path "$env:WORKDIR/databaseservices.json" | ConvertFrom-Json
foreach ($database in $databases) {
    $filename = "$($database.name).tar.gz"
    az storage blob download --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER --name  "mongodb/$filename" --file "$destinationDir/$filename"

      
}

