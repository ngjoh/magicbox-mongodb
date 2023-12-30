

$process = Get-Process -Name centrifugo -ErrorAction SilentlyContinue #| Stop-Process -Force
$centrifugoRunning = $process -ne $null
write-host "Centrifugo process running: $centrifugoRunning"

if ($centrifugoRunning ) {return}

$centrifugoRoot = . "$PSScriptRoot/get-toollocation.ps1" -toolkey "centrifugo"

Set-Location $centrifugoRoot
$centrifugoInstalled  = Test-Path "$centrifugoRoot/centrifugo"
write-host "Centrifugo installed: $centrifugoInstalled"

if (!$centrifugoInstalled){
    curl -sSLf https://centrifugal.dev/install.sh | sh
$config = @"
{
    "token_hmac_secret_key": "93027222-bea5-4637-96e6-cf843b65b1ee",
    "admin_password": "356534df-45ba-4ee9-b27a-519eb539afa6",
    "admin_secret": "97621e09-f4e9-4b44-8736-ed8e75a399e7",
    "api_key": "913f84d9-797c-49e7-b2ac-8bacb40f7637",
    "log_level" : "debug",
    "allowed_origins": ["*"]
  }
"@  
    Out-File -FilePath $centrifugoRoot/config.json -InputObject $config
}


if (!$centrifugoRunning){
    write-host "Centrifugo process starting" 
    Start-Process $centrifugoRoot/centrifugo -ArgumentList "--swagger --sockjs --admin --health  --client_insecure --admin_insecure"
} 