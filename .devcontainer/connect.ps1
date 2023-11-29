# TODO - Figure out if all of these are needed - until then comment out what is not needed
az login --use-device-code --allow-no-subscriptions
az aks install-cli
az account set --subscription "Office365 admin" -o table
az aks get-credentials --resource-group magicbox --name magicbox-prod

kubectl port-forward services/prod-mongos 27017:27017 -n percona