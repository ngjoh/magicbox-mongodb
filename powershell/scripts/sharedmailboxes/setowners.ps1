# Connector: Exchange
# Commands: New-Mailbox,set-Mailbox,Add-MailboxPermission,Add-RecipientPermission 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string[]]$owners
)

$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}


if ($owners -ne $null -and $owners -ne "" ) {
    $attr = "$($owners -join ",")"
   write-output "Setting CustomAttribute1 to $attr on $ExchangeObjectId"
    Set-Mailbox -Identity $ExchangeObjectId -CustomAttribute1 $attr
}
    
$result = @{
    Owners = $attr
}

 ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
