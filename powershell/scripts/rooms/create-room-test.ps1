$guid = New-Guid
. "$psscriptroot/create-room.ps1"  -Name $guid -Capacity 10 -Prefix "test-"

