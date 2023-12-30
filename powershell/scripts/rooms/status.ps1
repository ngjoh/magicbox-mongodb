$mailbox = "room-dk-kb601-23d3@nexigroup.com"
$calendarProcessing = Get-CalendarProcessing $mailbox
$room = Get-Mailbox  $mailbox
$place = Get-Place $mailbox

$result = @{
    "CalendarProcessing" = $calendarProcessing
    "Room" = $room
    "Place" = $place
}



ConvertTo-Json $result