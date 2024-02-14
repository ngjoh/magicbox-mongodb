<#---
title: Sidecar Pattern
output: deployment-backupjob.yaml
connection: sharepoint
tag: sidecar
---

This scripts creates a deployment with two containers. The first container is the main container and the second container is the sidecar container. The main container is a busybox container that runs a command to cat the access.log file every 30 seconds. The sidecar container is also a busybox container that runs a command to cat the access.log file every 30 seconds. Both containers share the same volume mount to the /var/log/nginx directory. The volume mount is an emptyDir volume. The deployment is named sample-datahost-with-sidecar and has the label app: datahostwothsidecar.

#>

param(
    # $name = "goworker",
    # $namespace = "christianiabpos"
)
$result = "deployment-backupjob.yaml"
$yaml = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: datahostwithsidecar
spec:
  selector:
    matchLabels:
      app: datahostwithsidecar
  replicas: 1
  template:
    metadata:
      labels:
        app: datahostwithsidecar
    spec:
      containers:
        - name: datahost
          image: busybox
          command: ["sh","-c","while true;  sleep 30; done"]
          volumeMounts:
            - name: backup
              mountPath: /data/backup
        - name: sidecar
          image: busybox
          command: ["sh","-c","while true; sleep 30; done"]
          volumeMounts:
            - name: backup
              mountPath: /data/backup
      volumes:
        - name: backup
          emptyDir: {}
                                    


                                    

"@

Out-File -FilePath $result -InputObject $yaml  -Encoding utf8NoBOM