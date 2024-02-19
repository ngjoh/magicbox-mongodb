<#---
title: Backup all databases
description: Backup all databases in the cluster
connection: sharepoint
input: databaseservices.json
api: post
tag: all
---#>

if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
}
$destinationDir = "$env:WORKDIR/dump"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}
$databases = Get-Content -Path "$env:WORKDIR/databaseservices.json" | ConvertFrom-Json
foreach ($database in $databases) {
    $databasename = $database.name
    $vars = kubectl exec "pod/$databasename-0" -n percona -- "env" 
    Write-Host $vars
    foreach ($var in $vars) {
        $s = $var -split "="
        if ($s[0] -eq "MONGODB_DATABASE_ADMIN_PASSWORD") {
         
            Write-Host $s[1] 
            $password = $s[1]
        }
    }
    
    
    $cmd = "mongodump --username databaseAdmin --password $($password) --authenticationDatabase admin "
    
    kubectl exec "pod/$databasename-0" -n percona -- mongodump --username databaseAdmin --password $($password) --authenticationDatabase admin -o /data/db/dump
    
    kubectl exec "pod/$databasename-0" -n percona -- tar -czvf dump.tar.gz /data/db/dump
    
    kubectl cp $databasename-0:/data/db/dump.tar.gz  $destinationDir/$databasename.tar.gz -n percona
    $timestamp = get-date -f "yyyy-MM-dd-HH"
    az storage blob upload  --account-name $env:AZURE_STORAGE_ACCOUNT --account-key $env:AZURE_STORAGE_KEY --container-name $env:AZURE_STORAGE_CONTAINER --overwrite $true  --file $destinationDir/$databasename.tar.gz --name mongodb/$timestamp/$databasename.tar.gz 
    
}

