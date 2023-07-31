# [9:11 AM] Mele Fabrizio

# •    https://christianiabpos.sharepoint.com/sites/intranets-hr/SitePages/On-boarding.aspx

# •    https://christianiabpos.sharepoint.com/sites/intranets-hr/SitePages/Buddy-Program.aspx

# •    https://christianiabpos.sharepoint.com/sites/intranets-hr/SitePages/New2Nets.aspx

# [9:11 AM] Mele Fabrizio

# Inside:

# [9:11 AM] Mele Fabrizio

#https://christianiabpos.sharepoint.com/sites/nexiintra-unit-gf-hr

#Parameters
$SourceSiteURL = "https://christianiabpos.sharepoint.com/sites/intranets-hr"
$DestinationSiteURL = "https://christianiabpos.sharepoint.com/sites/nexiintra-unit-gf-hr"
$PageName =  "New2Nets.aspx"
 
#Connect to Source Site
Connect-PnPOnline -Url $SourceSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 

#Export the Source page
$TempFile = [System.IO.Path]::GetTempFileName()
Export-PnPPage -Force -Identity $PageName -Out $TempFile
 
#Import the page to the destination site

Connect-PnPOnline -Url $DestinationSiteURL -ClientId $PNPAPPID -Tenant $PNPTENANTID -CertificatePath "$PNPCERTIFICATEPATH" 

Invoke-PnPSiteTemplate -Path $TempFile


#Read more: https://www.sharepointdiary.com/2020/07/sharepoint-online-copy-pages-to-another-site-using-powershell.html#ixzz88YGGWHXQ