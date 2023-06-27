param (
    [Parameter(Mandatory = $true)]
    [string]$Name,
    [Parameter(Mandatory = $true)]
    [int]$Capacity

)

$alias = $Name.Split("(")[0].Trim().Replace(" ", "-").ToLower()
$mailbox = New-Mailbox -Name  "room-$alias" -DisplayName "$Name" -Room -ResourceCapacity  $Capacity


$result = @{
    MailAddress = $mailbox.WindowsEmailAddress
}
ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM


$mailbox.WindowsEmailAddress