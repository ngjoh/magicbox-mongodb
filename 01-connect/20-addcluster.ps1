<#---
title: Add Cluster
---
#>
param (
    $resourcegroup = "magicbox",
    $clustername = "magicbox-prod")

az aks get-credentials --resource-group $resourcegroup --name $clustername