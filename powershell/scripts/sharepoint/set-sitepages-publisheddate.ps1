param (
    [Parameter(Mandatory = $true)]
    [string]$Url,
    [Parameter(Mandatory = $true)]
    [string]$listname
)
$filename = "$PSScriptRoot/itemdata.xml"

# $url = "https://christianiabpos.sharepoint.com/sites/cava3"
Connect-PnPOnline -Url $url  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
# Install-PnPApp -Identity "b0713514-2f12-46d9-833c-990ec907830b" #-Scope Site

$result = Get-PnPListItem -List $listname #-Fields "Title","FileLeafRef","FileRef","FileDirRef","File_x0020_Type"
foreach ($pageItem in $result) {
    <# $currentItemName is the current item #>
    # write-host $pageItem.FieldValues.PromotedState  $pageItem.FieldValues.FileRef
    if ( $pageItem.FieldValues.PromotedState -eq 2) {
  
        $NewsID = $pageItem.FieldValues.ID
        # $Date = "2023-01-01T00:00:00.000"  
        $Date = $pageItem.FieldValues.OriginalPublishedDate
        $PubDate = $pageItem.FieldValues.FirstPublishedDate

        if ($Date -ne $null) {

            $file = $pageItem.FieldValues.FileRef
            $page = Get-PnPPage -Identity $file.Split("SitePages/")[1]
            write-host $Date $PubDate $pageItem.FieldValues.FileRef -NoNewline
            Set-PnPListItem -List $ListName -Identity $NewsID -Values @{"FirstPublishedDate" = $Date; "Created" = $Date } -UpdateType SystemUpdate
            $page.publish("Published by script")
            write-host " updated"
        } 
    }
}
return
$result = @{
    siteurl  = $siteUrl
    type     = "sitetemplate"
    filename = $filename
}

ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM 

 