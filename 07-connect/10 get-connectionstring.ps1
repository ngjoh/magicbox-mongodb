<#---
title: Get connection strings
connection: sharepoint
output: connectionstrings.json
api: post
tag: connectionstrings
---#>
param ($databasename = "prod2")
if ($null -eq $env:WORKDIR) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
}
$password = ""
$username = ""
$destinationDir = "$env:WORKDIR/dump"
if (-not (Test-Path $destinationDir)) {
    $x = New-Item -Path $destinationDir -ItemType Directory 
}

$vars = kubectl exec "pod/$databasename-mongos-0" -n percona -- "env" 
    
foreach ($var in $vars) {
    $s = $var -split "="
    if ($s[0] -eq "MONGODB_DATABASE_ADMIN_PASSWORD") {
        $password = $s[1]
    }

    if ($s[0] -eq "MONGODB_DATABASE_ADMIN_USER") {
        $username = $s[1]
    }
}
    
    

$connectionStrings = @{
    direct = "mongodb://"+$username+":"+$password+ "@databasename-mongos.percona:27017/?directConnection=true&authMechanism=DEFAULT&tls=false"
    forwarded = "mongodb://"+$username+":"+$password+ "@localhost:27017/?directConnection=true&authMechanism=DEFAULT&tls=false"
}

Write-Output $connectionString
 
ConvertTo-Json -Depth 10 -InputObject $connectionString
| Out-File "$env:WORKDIR/connectionstrings.json"

