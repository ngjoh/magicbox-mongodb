# Connector: Exchange
# Commands: Set-Mailbox 
param (
    [Parameter(Mandatory = $true)]
    [string]$ExchangeObjectId,
    [Parameter(Mandatory = $true)]
    [string]$DisplayName
)

Set-Mailbox -Identity $ExchangeObjectId -DisplayName $DisplayName -Confirm:$false

Write-Output "done"