# # Get-DistributionGroupMember -Identity smbacl-it-operational-coordination-no  -ResultSize 1  | select *prim*| fl
# $UPN = "morten.torgersen@nexigroup.com"

# Get-DistributionGroupMember -Identity smbacl-it-operational-coordination-no  
# | where { $_.PrimarySmtpAddress -eq $UPN } #| select *prim*| fl
#  #-Filter "PrimarySmtpAddress= '$UPN'" #| select *prim*| fl



# Set-DistributionGroup -Identity 5c4d37c5-4af3-42ba-8ef3-94280f6912c6 -ManagedBy "niels.johansen@nexigroup.com"  -BypassSecurityGroupManagerCheck 

# Add-DistributionGroupMember -Identity 5c4d37c5-4af3-42ba-8ef3-94280f6912c6  -Member "niels.johansen@nexigroup


Get-DistributionGroup -anr "nexi-intra-news-channel-" | Remove-DistributionGroup 