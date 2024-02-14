<#---
title: Create Sidecar 
input: deployment-backupjob.yaml
connection: sharepoint
tag: sidecarcreate
---

#>

kubectl apply -f deployment-backupjob.yaml --namespace=sandbox
