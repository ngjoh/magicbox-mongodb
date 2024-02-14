<#---
title: Database Discovery
description: Discover databases in the cluster
connection: sharepoint
api: post
tag: discover
output: databaseservices.json
---#>
if ($env:WORKDIR -eq $null) {
    $env:WORKDIR = "$psscriptroot/../.koksmat/workdir"
}

$services = kubectl get services -n percona -o json | ConvertFrom-Json

$result = @()
$databaseservices = @()
foreach ($item in $services.items) {
    $match = [string]$item.metadata.name.EndsWith("-mongos")
    if ($match -eq "True") {
        $result += $item
        $databaseservices += @{
            name = $item.metadata.name
            namespace = $item.metadata.namespace
           
        }
    }
}







ConvertTo-Json -Depth 10 -InputObject $databaseservices
| Out-File "$env:WORKDIR/databaseservices.json"
