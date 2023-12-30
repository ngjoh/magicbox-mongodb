
param (
    [Parameter(Mandatory = $true)]
    [string]$SiteURL,
    [Parameter(Mandatory = $true)]
    [string]$tenantDomain
)
. $PSScriptRoot/$tenantDomain-connect-pnp.ps1 -SiteUrl $SiteURL


$script = Get-PnPSiteScriptFromWeb -IncludeAll
Set-Clipboard -Value $script