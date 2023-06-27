param (
    [Parameter(Mandatory = $true)]
    [string]$Mail,
    [Parameter(Mandatory = $true)]
    [int]$BookingWindowInDays
)


write-host "Processing" $Mail

Set-Mailbox $Mail  -MailTip ""
Set-CalendarProcessing $Mail  -DeleteComments $false -AutomateProcessing AutoAccept  -AllBookInPolicy:$true -BookInPolicy $null -BookingWindowInDays $BookingWindowInDays       
 