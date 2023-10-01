
$childSites = Get-PnPHubSiteChild -Identity b80f09f2-c5e5-4f69-9944-33e8fe18a96c 
foreach ($childSite in $childSites) {
    write-host "Installing templates on $($childSite)"
    
    . "$psscriptroot/apply-site-allpages.ps1" -DestinationSiteURL $childSite
}
