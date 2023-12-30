$loc = get-location
set-location "/Users/nielsgregersjohansen/kitchens/danish"
. $PSScriptRoot/connect-pnp.ps1 `
-SiteUrl "https://365adm.sharepoint.com/sites/koksmat" `
 -tenantDomain "365adm" `
 -kitchen "danish"

 Get-PnPTenantAppCatalogUrl