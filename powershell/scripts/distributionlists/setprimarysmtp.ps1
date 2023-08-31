anr = "nexi-intra-news-channel-" 

foreach ($dl in Get-DistributionGroup -Anr $anr ) {
    $m = $dl.Alias
   Set-DistributionGroup -Identity $dl.ExchangeObjectId -PrimarySmtpAddress  "$m@nexigroup.com" #-WhatIf
}