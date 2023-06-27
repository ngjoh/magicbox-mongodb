$env = "55037fcd-7d88-4ca7-b347-009a31b5f0c9"
pac  solution list -env $env --json 
| ConvertFrom-Json
| Where-Object { (($_.PublisherUniqueName -eq "adxstudio") `
-or ($_.PublisherUniqueName -eq "Cra8333") `
-or ($_.PublisherUniqueName -eq "microsoftdynamicslabs")) -ne $true }
| ConvertTo-Json -Depth 100
| Out-File -FilePath "$PSScriptRoot/powerapps-solutions-$env.json" -Encoding:utf8NoBOM
