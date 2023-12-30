

param (
    [Parameter(Mandatory = $true)]
    [string]$SiteURL,
    [Parameter(Mandatory = $true)]
    [string]$tenantDomain
)
. $PSScriptRoot/connect-pnp.ps1 -SiteUrl $SiteURL -tenantDomain $tenantDomain


