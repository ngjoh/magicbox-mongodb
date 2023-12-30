
$AZCERTIFICATEPATH = "/Users/nielsgregersjohansen/tmp50wlsl0z.pem" # "$($PSScriptRoot)/az.pfx"
$AZURE_TENANT_ID=$env:AZURE_TENANT_ID
$MAGICBOX_SERVICE_PRINCIPAL_ID = $env:MAGICBOX_SERVICE_PRINCIPAL_ID
#$PNPSITE=$env:PNPSITE
#$bytes = [Convert]::FromBase64String($ENV:PNPCERTIFICATE)
#[IO.File]::WriteAllBytes($PNPCERTIFICATEPATH, $bytes)



az login --service-principal --username $MAGICBOX_SERVICE_PRINCIPAL_ID --tenant $AZURE_TENANT_ID --password $AZCERTIFICATEPATH
