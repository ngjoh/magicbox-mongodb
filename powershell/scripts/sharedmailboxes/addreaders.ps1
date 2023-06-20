# Connector: Exchange
# Commands: New-Mailbox,set-Mailbox,Add-MailboxPermission,Add-RecipientPermission 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string[]]$readers
)

$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}

if ($readers -ne $null -and $readers -ne "" ) {
    foreach ($reader in $readers) {
        write-output "Adding $reader to $ExchangeObjectId"
        Add-MailboxPermission -Identity $ExchangeObjectId  -User $reader  -AccessRights ReadPermission -InheritanceType All | Out-Null
    }
}

Start-Sleep -s 2

$resultingSetOfMembers = Get-MailboxPermission -Identity $ExchangeObjectId     
| Where-Object { ($_.User -like '*@*')  -and ($_.AccessRights -like '*ReadPermission*')}
| Select User,AccessRights,IsInherited 

if (!($resultingSetOfMembers -is [array])){
    $resultingSetOfMembers = @($resultingSetOfMembers)
}
$result = @{
    Members = $resultingSetOfMembers
}
ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
