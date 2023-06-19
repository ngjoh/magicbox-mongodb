# Connector: Exchange
# Commands: Set-Mailbox 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string]$Email
)
$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}

write-output "Setting $Email as primary email address on $ExchangeObjectId"
Set-Mailbox -Identity $ExchangeObjectId  -WindowsEmailAddress $Email -Confirm:$false


Write-Output "done"