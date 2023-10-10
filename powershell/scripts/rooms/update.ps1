# Modified 10/9/2023 2:13 PM by niels.johansen@nexigroup.com
$mail = "room-dk-kb601-21a2@nets.eu"
$restrictedTo =  "bente.pedersen@nexigroup.com","julie.hollaender@external.nexigroup.com" 
write-host "Processing" $mail
Set-Mailbox $mail  -MailTip "This room has restrictions on who can book it"
Set-CalendarProcessing $mail  -DeleteComments $false -AutomateProcessing AutoAccept -AllRequestInPolicy $false  -AllBookInPolicy $false -BookInPolicy $restrictedTo -BookingWindowInDays 601 -ResourceDelegates $null       
        
        
# Modified 10/9/2023 2:13 PM by niels.johansen@nexigroup.com
$mail = "room-dk-kb601-21a5@nets.eu"
$restrictedTo =  "bente.pedersen@nexigroup.com","julie.hollaender@external.nexigroup.com" 
write-host "Processing" $mail
Set-Mailbox $mail  -MailTip "This room has restrictions on who can book it"
Set-CalendarProcessing $mail  -DeleteComments $false -AutomateProcessing AutoAccept -AllRequestInPolicy $false  -AllBookInPolicy $false -BookInPolicy $restrictedTo -BookingWindowInDays 601 -ResourceDelegates $null       
 