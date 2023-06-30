param (
    [Parameter(Mandatory = $true)]
    [string]$HubSiteId
)

$childSites = Get-PnPHubSiteChild -Identity $HubSiteId

$sites = @()
foreach ($childSite in $childSites) {
    Connect-PnPOnline -Url $childSite  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"

    $site = Get-PnPSite -Includes RootWeb, ServerRelativeUrl
    
    $web = Get-PnPWeb -Includes Title, ServerRelativeUrl,Navigation
    $nav = Get-PnPNavigationNode -Location QuickLaunch  #-Web $web

  
    Write-Output "Site $($web.Title) "
    


    function IterateNav($nav){  
        $DataColl = @()
        Foreach($n in $nav){
            $node = Get-PnPNavigationNode -Id $n.Id 
            $Data =  @{
               
                Title    = $node.Title
                RelativeURL = $node.Url     
              
            }
            
            if($node.Children.Count -gt 0){
                $Data.Childs = IterateNav($node.Children)
            }
            $DataColl += $Data
        }
        return $DataColl
    }
    $Navigation = IterateNav($nav)

    
    $sites += @{
        siteurl = $childSite 
        title   = $web.Title
        navigation   = $Navigation
       
    }
    
    
    
    
}


ConvertTo-Json  -InputObject $sites -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
