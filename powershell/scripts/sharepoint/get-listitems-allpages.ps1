param (
    [Parameter(Mandatory = $true)]
    [string]$Url
)

# $url = "https://christianiabpos.sharepoint.com/sites/cava3"
Connect-PnPOnline -Url $url  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
# Install-PnPApp -Identity "b0713514-2f12-46d9-833c-990ec907830b" #-Scope Site

Get-PnPListItem -List "Site Pages" -PageSize 5000 #| Select-Object -Property Id,


ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM 

 