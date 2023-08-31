
$anr = "nexi-intra-news-channel-" 


$result = Get-DistributionGroup -Anr $anr | select ExchangeObjectId,DisplayName,WindowsEmailAddress 
| Export-Csv  $PSScriptRoot/dls.csv -Encoding:utf8NoBOM 

