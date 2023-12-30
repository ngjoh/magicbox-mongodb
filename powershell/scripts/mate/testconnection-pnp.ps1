param (

    [Parameter(Mandatory = $true)]
    [string]$applicationname,
    [Parameter(Mandatory = $true)]
    [string]$tenantname
)
$location = Get-Location  ## Should be set to the correct kitchen by caller
Write-Host  "Running Register-PnPAzureADApp in location $location" 

$sharepointPath = "$location/.sharepoint"
$sharepointPathExists = Test-Path $sharepointPath 
if (!$sharepointPathExists){
    New-Item sharepointPath -type directory 
    Write-Host "SharePoint path created"
}
$tenantPath = "$sharepointPath/tenants"
$tenantPathExists = Test-Path $tenantPath 
if (!$tenantPathExists){
    New-Item $tenantPath -type directory 
    Write-Host "Tenant path created"
}

$thistenantPath = "$tenantPath/$tenantname"
$thistenantPathExists = Test-Path $thistenantPath 
if (!$thistenantPathExists){
    New-Item $thistenantPath -type directory 
    Write-Host "This Tenant path created"
}
Set-Location $thistenantPath
$result = Register-PnPAzureADApp -DeviceLogin -ApplicationName "$applicationname" -Tenant "$tenantname.onmicrosoft.com" 
$filename = "connect-pnp.json"
ConvertTo-Json -InputObject $result | Out-File -FilePath $filename -Encoding:utf8NoBOM 
Write-Host  "Done, data written to  $thistenantPath/$filename" 

# -CertificatePath  "$location\certificate.pfx" # -CertificatePassword (ConvertTo-SecureString -String "password" -AsPlainText -Force) 

# return 
# Write-Error "Step 1 ..." 
# Start-Sleep -Seconds 2
# Write-Error  "Step 2 ..."
# Start-Sleep -Seconds 2
# Write-Error  "Step 3 ..."
# Start-Sleep -Seconds 2

# Write-Error  "Connected"

# $result = "Hello World!"



