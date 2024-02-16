---
title: kubernetes-management
description: Describe the main purpose of this kitchen
---

# kubernetes-management




We define a timer job which runs as a sidecar to any given POD needing backup. The timer job will run a backup job at a given interval. The backup job will create a backup of the database and store it in an external store.

Virtual Machine (pod)
Internal Storage (kubernetes)
External Storage (e.g. Azure blob)

The pod mount the internal storage and the external storage. The backup job will download to internal storage, then move the backup to the external storage.

## Recovery

The pod will copy the backup from the external storage to the internal storage. The recovery job will then restore the database from the backup.


#

```bash
apk add neofetch
neofetch
```
