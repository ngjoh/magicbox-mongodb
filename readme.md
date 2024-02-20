---
title: Magicbox MongoDB
description: Everything you need to be able to do to deploy, use and operate MongoDB on the Magicbox platform
---

# Magicbox MongoDB

## Install

## Connect

## Use 

## Recover
The foundation for being able to recover data is to have a backup. We define a timer job which runs as a sidecar to any given POD needing backup. The timer job will run a backup job at a given interval. The backup job will create a backup of the database and store it in an external store.

All jobs are organized in a subfolder pr type of system to backup. Initially we have a "mongodb" folder. 

If the folder, you find a number of scripts organized with names making it sorted by the order they should be run.


Virtual Machine (pod)
Internal Storage (kubernetes)
External Storage (e.g. Azure blob)

The pod mount the internal storage and the external storage. The backup job will download to internal storage, then move the backup to the external storage.

## Recovery

The pod will copy the backup from the external storage to the internal storage. The recovery job will then restore the database from the backup.


