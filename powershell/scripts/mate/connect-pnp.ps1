
param (

    [Parameter(Mandatory = $true)]
    [string]$kitchen,
    [Parameter(Mandatory = $true)]
    [string]$SiteURL,
    [Parameter(Mandatory = $true)]
    [string]$tenantDomain
)
$ErrorActionPreference = "Stop"
$kitchenRoot = $env:KITCHENROOT
set-location "$kitchenRoot/$kitchen"
$location = get-location 
. $location/.sharepoint/tenants/$tenantDomain/env.ps1
Connect-PnPOnline -Url $SiteURL  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

