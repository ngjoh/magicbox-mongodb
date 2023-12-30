
param (
    [Parameter(Mandatory = $true)]
    [string]$SiteURL,
    [Parameter(Mandatory = $true)]
    [string]$tenantDomain
)
. $PSScriptRoot/$tenantDomain-connect-pnp.ps1 -SiteUrl $SiteURL
$filename = "$PSScriptRoot/allpages-template.xml"

Get-PnPSiteTemplate -Out  $filename -force  -Handlers All -Debug -IncludeAllPages:$true -IncludeSiteGroups:$false -IncludeAllTermGroups:$false  -PersistBrandingFiles:$true -PersistMultiLanguageResources:$false

$filename

