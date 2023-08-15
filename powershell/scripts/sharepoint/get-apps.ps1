# Get-PnPList

$PnPList = Get-PnPList -Identity "Documents"
$ListFields = Get-PnPProperty -ClientObject $PnPList -Connection $Connection -Property "Fields"
$ListFields | ConvertTo-Json