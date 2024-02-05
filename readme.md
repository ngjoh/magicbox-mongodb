---
title: kubernetes-management
description: Describe the main purpose of this kitchen
---

# kubernetes-management


```mermaid

sequenceDiagram

    Operator->>+Database: Database, can you hear me?
    Database-->>-Operator: Hi Operator, I can hear you!
    Operator->>+Database: Make a backup
    Database-->>-Operator: I'm done
    Operator->>+Database: Upload it to Azure Blob
    Database->>+Azure: Store this
    Azure-->>-Database: Done
    Database-->-Operator: Done
    
``````