# Connector: Exchange
# Commands: Set-Mailbox 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string]$DisplayName
)
$mb = Get-MailBox -Identity $ExchangeObjectId
if ($mb -eq $null) {
    write-output "Mailbox $ExchangeObjectId not found"
    exit 1
}
Set-Mailbox -Identity $ExchangeObjectId -DisplayName $DisplayName -Confirm:$false

Write-Output "done"