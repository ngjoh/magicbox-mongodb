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
        Add-MailboxPermission -Identity $mailbox.ExchangeObjectId  -User $member  -AccessRights ReadPermission -InheritanceType All | Out-Null
    }
}

