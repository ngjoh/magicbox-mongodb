Set-Location "$PSScriptRoot/.."
docker build --pull --rm -f "Dockerfile" -t kubernetesmanagement:latest "." 