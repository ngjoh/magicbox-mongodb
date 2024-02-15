<#---
title: Database Restore
connection: sharepoint
api: post
tag: view

---#>
param ($database="prod-mongos")

$destinationDir = "$env:WORKDIR/download/$database"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}
$backupfile = "$env:WORKDIR/download/$database.tar.gz"

tar -xvzf $backupfile -C $destinationDir



