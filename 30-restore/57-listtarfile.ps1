<#---
title: Database Restore
connection: sharepoint
api: post
tag: listtar
output: backupcontent.json
---#>
param ($database="booking-mongos")
#if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
#}

$backupfile = "$env:WORKDIR/download/mongodb/$database.tar.gz"

$output = tar -ztvf $backupfile 


 $output
| ConvertTo-Json -Depth 10 
| Out-File "$env:WORKDIR/backupcontent.json"
