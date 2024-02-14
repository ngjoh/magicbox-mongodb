<#---
title: Backup all databases
description: Backup all databases in the cluster
connection: sharepoint
input: databaseservices.json
tag: all
---#>

$env:WORKDIR = "$psscriptroot/../.koksmat/workdir"

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
    
    kubectl cp $databasename-0:/data/db/dump.tar.gz  $env:WORKDIR/$databasename.tar.gz -n percona
    
    
}

