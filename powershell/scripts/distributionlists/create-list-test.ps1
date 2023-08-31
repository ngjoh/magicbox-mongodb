

# . "$psscriptroot/create-list.ps1" -Name "Group News" -NamePrefix "Nexi Intra News Channel" -AliasPrefix "nexi-intra-news-channel"

Import-Csv -LiteralPath "$psscriptroot/News Channels.csv" | ForEach-Object {
    
     . "$psscriptroot/create-list.ps1" -Name  $_."Channel Name" -NamePrefix "Intra News" -AliasPrefix "nexi-intra-news-channel"
}