param (
    [Parameter(Mandatory = $true)]
    [string]$toolkey
)

$kitchenRoot = $env:KITCHENROOT

if (!$kitchenRoot ){
    Write-Error "Kitchen Root path not set"
    exit
}

$kitchenRootExists = Test-Path $kitchenRoot 
if (!$kitchenRootExists){
    $n = New-Item $kitchenRoot -type directory 
  #  Write-Host "Tools Root path created"

}
$toolsRootPath = "$kitchenRoot.tools"
$toolsRootPathExists = Test-Path $toolsRoot
if (!$toolsRootPathExists){
    $n = New-Item $toolsRootPath -type directory 
  #  Write-Host "Tools Root path created"

}

$toolsRoot = "$toolsRootPath/$toolkey"
$toolsRootExists = Test-Path $toolsRoot
if (!$toolsRootExists){
    $n =  New-Item $toolsRoot -type directory 
   # Write-Host "Tools  path created"

}

$toolsRoot