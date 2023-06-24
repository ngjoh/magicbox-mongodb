$filename = "$PSScriptRoot/sitepages.json"
$SitePages = Get-PnPListItem -List "Site Pages"
 
$PagesDataColl = @()
#Collect Site Pages data - Title, URL and other properties
ForEach($Page in $SitePages)
{
    $Data = New-Object PSObject -Property ([Ordered] @{
        PageName  = $Page.FieldValues.Title
        RelativeURL = $Page.FieldValues.FileRef     
        CreatedOn = $Page.FieldValues.Created_x0020_Date
        ModifiedOn = $Page.FieldValues.Last_x0020_Modified
        Editor =  $Page.FieldValues.Editor.Email
        ID = $Page.ID
    })
    $PagesDataColl += $Data
}
 
 
#Export data to CSV File
$PagesDataColl 
| Convertto-Json -Depth 10
| Out-File -FilePath $filename -Encoding utf8NoBOM

$result = @{
    siteurl = $siteUrl
    type="sitetemplate"
    filename = $filename
}

ConvertTo-Json -InputObject $result
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM 



#Read more: https://www.sharepointdiary.com/2020/11/sharepoint-online-get-all-pages-using-powershell.html#ixzz85LhfK9I5