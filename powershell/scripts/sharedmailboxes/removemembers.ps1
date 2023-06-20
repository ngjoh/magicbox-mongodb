# Connector: Exchange
# Commands: New-Mailbox,set-Mailbox,Add-MailboxPermission,Add-RecipientPermission 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string[]]$Members
)

$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}

if ($members -ne $null -and $members -ne "" ) {
    foreach ($member in $members) {
        write-output "Removing $member from $ExchangeObjectId"
        Remove-MailboxPermission -Identity $ExchangeObjectId  -User $member  -AccessRights FullAccess -Confirm:$false | Out-Null
        Remove-RecipientPermission -Identity $ExchangeObjectId  -AccessRights SendAs -Trustee $member -confirm:$false | Out-Null
    }
}

Start-Sleep -s 2

$resultingSetOfMembers = Get-MailboxPermission -Identity $ExchangeObjectId   
| Where-Object { ($_.User -like '*@*')  -and ($_.AccessRights -like '*FullAccess*')}  
| Select User,AccessRights,IsInherited 

if (!($resultingSetOfMembers -is [array])){
    $resultingSetOfMembers = @($resultingSetOfMembers)
}
$result = @{
    Members = $resultingSetOfMembers
}
ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
