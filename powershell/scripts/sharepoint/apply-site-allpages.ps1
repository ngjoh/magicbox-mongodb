param (
    [Parameter(Mandatory = $true)]
    [string]$DestinationSiteURL
)
$TempFile = "$PSScriptRoot/allpages-template.xml"

Connect-PnPOnline -Url $DestinationSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 

Invoke-PnPSiteTemplate -Path $TempFile -ExcludeHandlers SiteFooter,ApplicationLifecycleManagement
#Read more: https://www.sharepointdiary.com/2020/07/sharepoint-online-copy-pages-to-another-site-using-powershell.html#ixzz88YGGWHXQ