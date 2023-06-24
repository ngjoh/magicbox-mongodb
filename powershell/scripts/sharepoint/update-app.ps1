$childSites = Get-PnPHubSiteChild -Identity b80f09f2-c5e5-4f69-9944-33e8fe18a96c 
foreach ($childSite in $childSites) {
    write-host "Installing app on $($childSite)"
    Connect-PnPOnline -Url $childSite  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
    Update-PnPApp -Identity "b0713514-2f12-46d9-833c-990ec907830b" #-Scope Site
}

