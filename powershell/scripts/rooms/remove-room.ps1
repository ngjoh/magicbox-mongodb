param (
    [Parameter(Mandatory = $true)]
    [string]$Mail
)

write-host "Deleting" $mail
Remove-Mailbox $Mail -Confirm:$false
