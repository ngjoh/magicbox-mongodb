<#---
title: Architecture
---
```mermaid

sequenceDiagram

    Operator->>+Database: Database, can you hear me?
    Database-->>-Operator: Hi Operator, I can hear you!
    Operator->>+Database: Make a backup
    Database-->>-Operator: I'm done
    Operator->>+Database: Copy to external store
    Database->>+External Store: Store this
    External Store-->>-Database: Done
    Database-->-Operator: Done
    
``````


#>