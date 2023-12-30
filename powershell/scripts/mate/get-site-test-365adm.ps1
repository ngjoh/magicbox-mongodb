$loc = get-location
set-location "/Users/nielsgregersjohansen/kitchens/danish"
. $PSScriptRoot/get-site.ps1 `
-SiteUrl "https://365adm.sharepoint.com/sites/koksmat" `
 -tenantDomain "365adm"

 Set-Location $loc