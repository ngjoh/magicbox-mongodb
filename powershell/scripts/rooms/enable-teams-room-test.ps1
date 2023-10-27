#$guid = New-Guid
#$mail = . "$psscriptroot/create-room.ps1"  -Name $guid -Capacity 10 -Prefix "test-"
$mail = "room-dk-kb601-21d4@nets.eu"
. "$psscriptroot/enable-teams-room.ps1"   -Mail $mail -Password "n!ND12217001100"

#. "$psscriptroot/remove-room.ps1"  -Mail $mail 

