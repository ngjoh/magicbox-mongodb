<#---
title: Upload
connection: sharepoint

---


#>

param (
    $filename = "backupcontent.json"

)
$destinationDir = "$env:WORKDIR/download"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}
# "$env:WORKDIR/download/"
az storage blob download --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER --name  $filename --file "$destinationDir/$filename"
