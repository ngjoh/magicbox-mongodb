
$result = Get-AcceptedDomain # | Select DomainName,DomainType,IsValid 

 if (!($result -is [array])){
    $result = @($result)
 }

ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM

 


