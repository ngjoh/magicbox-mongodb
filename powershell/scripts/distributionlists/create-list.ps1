param (
    [Parameter(Mandatory = $true)]
    [string]$name,
    [Parameter(Mandatory = $true)]
    [string]$nameprefix,
    [Parameter(Mandatory = $true)]
    [string]$aliasprefix
)


$alias = $name.Replace(" ","-").ToLower()
$alias = ($alias -replace '[^a-zA-Z0-9\-]', '' )
New-UnifiedGroup -DisplayName "$nameprefix $name" -Alias "$aliasprefix-$alias"  # -AccessType "Public" -AutoSubscribeNewMembers:$false -RequireSenderAuthenticationEnabled:$true -Verbose
# New-DistributionGroup -Name "$nameprefix $name" -Alias "$aliasprefix-$alias" -Type "Security" 