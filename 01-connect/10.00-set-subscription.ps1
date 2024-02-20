<#---
title: Set Subscription
---
#>
param (
    $subscriptionName = "Office365 admin"
)
az account set --subscription $subscriptionName  -o table
