param (
    [Parameter(Mandatory = $true)]
    [string]$Mail,
    [Parameter(Mandatory = $true)]
    [string]$Password
)

Set-Mailbox $Mail -EnableRoomMailboxAccount $true  -RoomMailboxPassword (ConvertTo-SecureString -String $Password -AsPlainText -Force)

