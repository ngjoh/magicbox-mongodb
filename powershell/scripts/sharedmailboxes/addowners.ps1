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


if ($owner -ne $null -and $owner -ne "" ) {
   
    Set-Mailbox -Identity $mailbox.ExchangeObjectId -CustomAttribute1 $mb.CustomAttribute1 + ","+$Owners
}
    
