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
        Add-MailboxPermission -Identity $mailbox.ExchangeObjectId  -User $member  -AccessRights FullAccess -InheritanceType All | Out-Null
        Add-RecipientPermission -Identity $mailbox.ExchangeObjectId  -AccessRights SendAs -Trustee $member -confirm:$false | Out-Null
    }
}


