# Connector: Exchange
# Commands: Remove-Mailbox 
param (
    $ExchangeObjectId
)

Remove-Mailbox -Identity $ExchangeObjectId -Confirm:$false
write-output "done"