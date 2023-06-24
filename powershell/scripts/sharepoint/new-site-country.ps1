param (
    [Parameter(Mandatory = $true)]
    [string]$countryShortCode,
    [Parameter(Mandatory = $true)]
    [string[]]$countryName,
    [Parameter(Mandatory = $true)]
    [string[]]$owner   
)

$url = "https://christianiabpos.sharepoint.com/sites/nexiintra-country-$countryShortCode"
write-host "Creating country site $countryName with url $url"
New-PnPSite -Type CommunicationSite -Title "$countryName" -Url "$url" -HubSiteId "b80f09f2-c5e5-4f69-9944-33e8fe18a96c" -Owner $owner

Connect-PnPOnline -Url $url  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

Invoke-PnPSiteTemplate -Path "$PSScriptRoot/template-filtered.xml" -ResourceFolder "$PSScriptRoot" 
