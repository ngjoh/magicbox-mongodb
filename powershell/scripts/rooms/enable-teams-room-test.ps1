$guid = New-Guid
$mail = . "$psscriptroot/create-room.ps1"  -Name $guid -Capacity 10 -Prefix "test-"

. "$psscriptroot/enable-teams-room.ps1"   -Mail $mail -Password "sdaf90jnc$sz!!a"

. "$psscriptroot/remove-room.ps1"  -Mail $mail 

