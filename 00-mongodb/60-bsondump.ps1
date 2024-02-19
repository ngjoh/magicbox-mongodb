param (
    $database="booking-mongos",
    $collection="booking-cro"
)
$destinationDir = "$env:WORKDIR/mongodb/$database/data/db/dump/$collection"
Push-Location
set-location $destinationDir
$files = Get-ChildItem -Path $destinationDir -Filter "*.bson"  #-Include *.bson 
for ($i = 0; $i -lt $files.Count; $i++) {
    $file = $files[$i]
    $file = $file.Name
    bsondump $file --outFile="$file.json"
}
Pop-Location