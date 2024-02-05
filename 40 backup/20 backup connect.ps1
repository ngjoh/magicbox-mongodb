# kubectl get pods -n percona

$vars = kubectl exec svc/miller-mongos -n percona -- "env" 
Write-Host $vars
foreach ($var in $vars) {
    $s = $var -split "="
    if ($s[0] -eq "MONGODB_DATABASE_ADMIN_PASSWORD") {
     
        Write-Host $s[1] 
        $password = $s[1]
    }
}


$cmd = "mongodump --username databaseAdmin --password $($password) --authenticationDatabase admin "

Set-Clipboard -Value $cmd


