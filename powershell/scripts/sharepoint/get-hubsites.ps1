Get-PnPHubSite | select ID,Description,Title,SiteUrl
| ConvertTo-Json   -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
