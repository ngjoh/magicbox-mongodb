param (
    [Parameter(Mandatory = $true)]
    [string]$root,
    [Parameter(Mandatory = $true)]
    [string]$kitchenName ,
    [Parameter(Mandatory = $true)]
    $stationName ,
    [Parameter(Mandatory = $true)]
    $repourl 
)
# $root = "/Users/nielsgregersjohansen/kitchens"
# $kitchenName = "noma"
# $stationName = "ui"
# $repourl = ""
Write-Host "Checking station workspace: $root"
$kitchenPath = "$root/$kitchenName"
$kitchenPathExists = Test-Path $kitchenPath
if (!$kitchenPathExists){
    New-Item $kitchenPath -type directory 
    Write-Host "Kitchen path created"

}else
{
    Write-Host "Kitchen path already exists"
}

$stationPath = "$root/$kitchenName/$stationName"

$stationPathExists = Test-Path $stationPath
if (!$stationPathExists){
    New-Item $stationPath -type directory 
    Write-Host "Station path created"
}else{
    Write-Host "Station path already exists"

}

$sourceCodePath = "$root/$kitchenName/$stationName/sourcecode"
$sourceCodePath = Test-Path $sourceCodePath

if (!$sourceCodePath){
   set-location $stationPath
   git clone $repourl "sourcecode"
}else {
    Write-Host "Source code already exists"
}