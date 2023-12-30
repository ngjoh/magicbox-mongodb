
#az ad app create

$name  = "magicbox-azure"
$SubscriptionID = $env:AZURE_SUBSCRIPTION_ID
$ResourceGroupName = $env:AZURE_RESOUREGROUP_NAME
$TenantId = $env:AZURE_TENANT_ID
$roleName = "Contributor"
$result = "OK"

az ad sp create-for-rbac --name $name --role $roleName --scopes /subscriptions/$SubscriptionID/resourceGroups/$ResourceGroupName --create-cert


# az login --service-principal `
#          --username myServicePrincipalID `
#          --tenant myOwnerOrganizationId `
#          --password /path/to/cert