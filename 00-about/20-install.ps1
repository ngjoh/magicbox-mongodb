<#---
title: Install
---

Run this to get the build in packages ready
#>

Push-Location
Set-Location "$psscriptroot/../.koksmat/app"
go install
Set-Location "$psscriptroot/../.koksmat/web"
pnpm install
pnpm build
