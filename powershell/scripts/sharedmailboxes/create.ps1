# Connector: Exchange
# Commands: New-Mailbox,set-Mailbox,Add-MailboxPermission,Add-RecipientPermission 
param (
    [Parameter(Mandatory = $true)]
    [string]$Name,
    [Parameter(Mandatory = $true)]
    [string]$DisplayName,
    [Parameter(Mandatory = $true)]
    [string]$Alias,
    #[Parameter(Mandatory = $true)]
    [string[]]$Owners="",
    #[Parameter(Mandatory = $true)]
    [string[]]$Members="",
    #[Parameter(Mandatory = $true)]
    [string[]]$Readers=""
)


$mailbox = New-Mailbox -Shared -Name $name -DisplayName $displayName -Alias $alias
if ($mailbox -eq $null) {
    throw "Failed to create mailbox"
}
Start-Sleep -s 5

if ($owner -ne $null -and $owner -ne "" ) {
Set-Mailbox -Identity $mailbox.ExchangeObjectId -CustomAttribute1 ($Owners -join "," )
}


if ($members -ne $null -and $members -ne "" ) {
    foreach ($member in $members) {
        Add-MailboxPermission -Identity $mailbox.ExchangeObjectId  -User $member  -AccessRights FullAccess -InheritanceType All | Out-Null
        Add-RecipientPermission -Identity $mailbox.ExchangeObjectId  -AccessRights SendAs -Trustee $member -confirm:$false | Out-Null
    }
}


if ($readers -ne $null -and $readers -ne "" ) {
    foreach ($reader in $readers) {
        Add-MailboxPermission -Identity $mailbox.ExchangeObjectId  -User $member  -AccessRights ReadPermission -InheritanceType All | Out-Null
    }
}

write-output $mailbox | Select name,displayname,Identity,PrimarySmtpAddress,ExchangeObjectId | ConvertTo-Json | Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
