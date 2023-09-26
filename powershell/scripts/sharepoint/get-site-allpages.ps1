param (
    [Parameter(Mandatory = $true)]
    [string]$Url
)
$filename = "$PSScriptRoot/allpages-template.xml"

# $url = "https://christianiabpos.sharepoint.com/sites/cava3"
Connect-PnPOnline -Url $url  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
# Install-PnPApp -Identity "b0713514-2f12-46d9-833c-990ec907830b" #-Scope Site

Get-PnPSiteTemplate -Out  $filename -force  -Handlers All -Debug -IncludeAllPages:$true -IncludeSiteGroups:$false -IncludeAllTermGroups:$false  -PersistBrandingFiles:$true -PersistMultiLanguageResources:$true


$result = @{
    siteurl = $siteUrl
    type="sitetemplate"
    filename = $filename
}

ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM 

 