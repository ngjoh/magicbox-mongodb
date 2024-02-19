<#---
title: Database Restore
connection: sharepoint
api: post
tag: view

---#>
param ($database="booking-mongos")

$destinationDir = "$env:WORKDIR/mongodb/$database"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}
$backupfile = "$env:WORKDIR/download/mongodb/$database.tar.gz"

tar -xvzf $backupfile -C $destinationDir



