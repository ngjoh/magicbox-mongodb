# Connector: Exchange
# Commands: Remove-Mailbox 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId
)
$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}
Remove-Mailbox -Identity $ExchangeObjectId -Confirm:$false
write-output "done"