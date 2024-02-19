<#---
title: Download all backups
connection: sharepoint
input: bloblist.json
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
$databases = Get-Content -Path "$env:WORKDIR/bloblist.json" | ConvertFrom-Json
foreach ($database in $databases) {
    $filename = $database.name
    $sourceDirectory = [System.IO.Path]::GetDirectoryName($filename)
    $targetDirectory = "$destinationDir/$sourceDirectory"
    if (-not (Test-Path $targetDirectory)) {
        $x = New-Item -Path $targetDirectory -ItemType Directory 
    }
    az storage blob download --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER --name  $filename --file "$destinationDir/$filename"

}

