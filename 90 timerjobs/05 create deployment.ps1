<#---
title: Create Deployment
output: deployment-backupjob.yaml
connection: sharepoint
tag: create-deployment
---
#>

param(
    $name = "goworker",
    $namespace = "christianiabpos"
)

$yaml = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $name
spec:
  selector:
    matchLabels:
      app: $name
  replicas: 1
  template:
    metadata:
      labels:
        app: $name
    spec: 
      containers:
      - name: koksmat-cli
        image: ghcr.io/365admin/kubernetes-management:v0.0.1
        command: ["kubernetes-management"]
        args: ["backup","run"]               
        env:
          - name: DEBUG
            value: magicbox*
                                    

"@

Out-File -FilePath $result -InputObject $yaml  -Encoding utf8NoBOM