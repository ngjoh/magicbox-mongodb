#$ENV:PNPPSSITE="https://christianiabpos.sharepoint.com/sites/nexiintra-country-se"
$Web = Get-PnPWeb -Includes Title,Url,Id,Lists,Webs,SiteUsers,SiteGroups,AssociatedOwnerGroup,AssociatedMemberGroup,AssociatedVisitorGroup,AssociatedOwnerGroup.Users,AssociatedMemberGroup.Users,AssociatedVisitorGroup.Users
$Web | fl