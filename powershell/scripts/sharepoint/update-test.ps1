
    Connect-PnPOnline -Url "https://christianiabpos.sharepoint.com/sites/intra365"  -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH"
    Update-PnPApp -Identity "b0713514-2f12-46d9-833c-990ec907830b" #-Scope Site


