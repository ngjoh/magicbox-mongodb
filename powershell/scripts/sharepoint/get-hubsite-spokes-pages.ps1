param (
    [Parameter(Mandatory = $true)]
    [string]$HubSiteId
)

$childSites = Get-PnPHubSiteChild -Identity $HubSiteId
$sites = @()
foreach ($childSite in $childSites) {
    Connect-PnPOnline -Url $childSite  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

    $site = Get-PnPSite -Includes RootWeb, ServerRelativeUrl
    
    $web = $site.RootWeb

    $SitePages = Get-PnPListItem -List "Site Pages" -Fields "Title", "FileRef", "Created_x0020_Date", "Last_x0020_Modified", "Editor" -PageSize 5000
    Write-Output "Site $($web.Title) has $($SitePages.Count) pages"
    $PagesDataColl = @()
    #Collect Site Pages data - Title, URL and other properties
    ForEach ($Page in $SitePages) {
        $Data =  @{
                HubSiteId    = $HubSiteId
                PageName    = $Page.FieldValues.Title
                RelativeURL = $Page.FieldValues.FileRef     
                CreatedOn   = $Page.FieldValues.Created_x0020_Date
                ModifiedOn  = $Page.FieldValues.Last_x0020_Modified
                Editor      = $Page.FieldValues.Editor.Email
                ID          = $Page.ID
            }
        $PagesDataColl += $Data
    }
    
    $sites += @{
        siteurl = $childSite 
        title   = $web.Title
        HubSiteId = $HubSiteId
        WelcomePage = $web.WelcomePage
        pages   = $PagesDataColl
       
    }
    
    
    
    
}


ConvertTo-Json  -InputObject $sites -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
