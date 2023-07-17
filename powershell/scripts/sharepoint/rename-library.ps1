param (
    [Parameter(Mandatory = $true)]
    [string]$SourceSiteURL ,
    [Parameter(Mandatory = $true)]
    [string]$oldListName ,
    [Parameter(Mandatory = $true)]
    [string]$newListName ,
    [Parameter(Mandatory = $true)]
    [string]$newListUrl 
)

Connect-PnPOnline -Url $SourceSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 

# Get the SharePoint list
$list = Get-PnPList -Identity $oldListName

# Move SharePoint list to the new URL
$list.Rootfolder.MoveTo($newListUrl)
Invoke-PnPQuery

# Rename List

Set-PnPList -Identity $oldListName -Title $newListName