param (
    [Parameter(Mandatory = $true)]
    [string]$Url
)

Remove-PnPTenantSite -Url $Url -SkipRecycleBin -Force
Remove-PnPTenantDeletedSite -Identity $Url  -Force

