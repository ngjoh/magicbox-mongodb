
param (
    [Parameter(Mandatory = $true)]
    [string]$UPN,
    [Parameter(Mandatory = $true)]
    [string[]]$Memberships,
    [Parameter(Mandatory = $true)]
    [string[]]$DistributionLists
)

$result = @()
foreach ($exchangeObjectId in $DistributionLists) {
    $dg =  Get-DistributionGroup -Identity $exchangeObjectId
   
    $add = $false
    
    foreach ($membership in $Memberships) {
        # write-host "E" exchangeObjectId  "M" $membership
        if ($exchangeObjectId -eq $membership) {
            $add = $true
        }
    }

    $members = Get-DistributionGroupMember -Identity $exchangeObjectId | where { $_.PrimarySmtpAddress -eq $UPN } 

    foreach ($member in $members) {
   
        if ($member.PrimarySmtpAddress -eq $UPN) {
            $add = $false
        }
       
    }
    if ($add){
        write-host "Adding $UPN to $exchangeObjectId"
        Add-DistributionGroupMember -Identity $exchangeObjectId -Member $UPN #-WhatIf
      
        $result += "Added to $exchangeObjectId"
    }
    $remove = $false
    foreach ($member in $members) {
   
        if ($member.PrimarySmtpAddress -eq $UPN) {
            $remove = $true
        }
       
    }
    if ($remove){
        write-host "Removing $UPN from $exchangeObjectId"
        Remove-DistributionGroupMember -Identity $exchangeObjectId -Member $UPN #-WhatIf
       
        $result += "Removed from $exchangeObjectId"
    }

    
}

# $result = Get-DistributionGroup -Anr $anr | select ExchangeObjectId,DisplayName,WindowsEmailAddress | fl
# ConvertTo-Json -InputObject $result
# | Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
# return
# $DistributionGroups= Get-DistributionGroup -Anr $anr 
# | where { (Get-DistributionGroupMember $_.Name 
#     | foreach {$_.PrimarySmtpAddress}) -contains "$member"}

#     $DistributionGroups

$result
return

ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM

