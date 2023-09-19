param (
    [Parameter(Mandatory = $true)]
    [string]$SourceSiteURL ,
    [Parameter(Mandatory = $true)]
    [string]$DestinationSiteURL ,
    [Parameter(Mandatory = $true)]
    [string]$PageName
    [Parameter(Mandatory = $true)]
    [string]$DestPageName
    
)



Connect-PnPOnline -Url $SourceSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 


$TempFile = "$PSScriptRoot/copypagetemplate.xml"
$TempFile2 = "$PSScriptRoot/copypagetemplate2.xml"
Export-PnPPage -Force -Identity $PageName -Out $TempFile -PersistBrandingFiles

$newPageName = $DestPageName
$replaceTag = "PageName=""$($PageName)"""
$replaceWith = "PageName=""$($newPageName)"""

[string]$text = Get-Content $TempFile
$text.Replace($replaceTag, $replaceWith) | Set-Content $TempFile2 




Connect-PnPOnline -Url $DestinationSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 

Invoke-PnPSiteTemplate -Path $TempFile2
$copyFileResult = @{

    NewPageURL = "$DestinationSiteURL/SitePages/$newPageName"
  }


ConvertTo-Json  -InputObject $copyFileResult -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
