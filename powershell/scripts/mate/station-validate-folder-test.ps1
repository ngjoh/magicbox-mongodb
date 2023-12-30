    $root= "/Users/nielsgregersjohansen/kitchens"
    $kitchenName = "danish"
    $stationName = "icing"
    
    $repourl = "https://github.com/koksmat-com/sharepoint.git"


. "$psscriptroot/station-validate-folder.ps1" -root $root -kitchenName $kitchenName -stationName $stationName -repourl $repourl
