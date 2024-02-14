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
        image: ghcr.io/365admin/kubernetes-management:v0.0.2
        command: ["kubernetes-management"]
        args: ["backup","all"]               
        env:
          - name: DEBUG
            value: magicbox*
                                    

"@

Out-File -FilePath $result -InputObject $yaml  -Encoding utf8NoBOM


<#
spec:
  schedule: '*/10 * * * *'
  concurrencyPolicy: Allow
  suspend: false
  jobTemplate:
    metadata:
      creationTimestamp: null
    spec:
      parallelism: 1
      completions: 1
      activeDeadlineSeconds: 2400
      backoffLimit: 3
      template:
        metadata:
          creationTimestamp: null
          labels:
            app: pwsh2
        spec:
          containers:
            - name: koksmat-cli
              image: ghcr.io/365admin/kubernetes-management:v0.0.2
              command:
                - kubernetes-management
              args:
                - backup
                - all

              
              resources: {}
              terminationMessagePath: /dev/termination-log
              terminationMessagePolicy: File
              imagePullPolicy: IfNotPresent
          restartPolicy: Never
          terminationGracePeriodSeconds: 30
          dnsPolicy: ClusterFirst
          securityContext: {}
          schedulerName: default-scheduler
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1


#>