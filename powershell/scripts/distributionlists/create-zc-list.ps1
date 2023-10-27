$upn = "niels.johansen@nexigroup.com"

$ErrorActionPreference = "SilentlyContinue"

$dl = Get-DistributionGroup "zc-dl-company-31de442c0e5c09d941d23300f0819e8536bb5f63"
if ($dl -eq $null){
    $ErrorActionPreference = "Continue"
    New-DistributionGroup -Name "zc-dl-company-31de442c0e5c09d941d23300f0819e8536bb5f63" -DisplayName "All employees Centrum Rozliczeń El. [Nets]" -ManagedBy ${upn} 
    $ErrorActionPreference = "SilentlyContinue"
}