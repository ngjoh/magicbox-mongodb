
$guid = New-Guid
. "$psscriptroot/create.ps1"  -Name "test5-$guid" -DisplayName "Test5 $guid"  -Alias "test5-$guid" -Members "s" -Readers "s" -Owner="s"


