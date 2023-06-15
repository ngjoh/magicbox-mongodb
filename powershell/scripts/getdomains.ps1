

 Get-AcceptedDomain | Select DomainName,DomainType,IsValid | ConvertTo-Json | Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
