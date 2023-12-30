write-host "Hello World from" -NoNewline
write-host  $PSScriptRoot -ForegroundColor DarkYellow

az login --service-principal -u $env:AZURE_CLIENT_ID -p $env:AZURE_CLIENT_SECRET --tenant $env:AZURE_TENANT_ID