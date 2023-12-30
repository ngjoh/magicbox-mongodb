$AZURE_TOKEN_ID = az account get-access-token --resource-type ms-graph --query accessToken --output tsv
curl --header "Authorization: Bearer $AZURE_TOKEN_ID" --request GET 'https://graph.microsoft.com/v1.0/domains' | jq -r '.value[] | select(.isDefault == true) | {id}[]'

