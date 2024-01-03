package connectors

import (
	"encoding/json"
	"time"
)

func AzureSubscriptions() ([]Connector, error) {
	type Account struct {
		CloudName        string `json:"cloudName"`
		HomeTenantID     string `json:"homeTenantId"`
		ID               string `json:"id"`
		IsDefault        bool   `json:"isDefault"`
		ManagedByTenants []any  `json:"managedByTenants"`
		Name             string `json:"name"`
		State            string `json:"state"`
		TenantID         string `json:"tenantId"`
		User             struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"user"`
	}

	bytes, err := Execute("az", *&Options{}, "account", "list")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	accounts := make([]Account, 0)

	err = json.Unmarshal(bytes, &accounts)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		connector := Connector{
			Name:        account.Name,
			Description: account.User.Name,
			IsCurrent:   account.IsDefault,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func AzureStorageAccounts() ([]Connector, error) {
	type StorageAccount struct {
		AccessTier                            string      `json:"accessTier"`
		AccountMigrationInProgress            interface{} `json:"accountMigrationInProgress"`
		AllowBlobPublicAccess                 bool        `json:"allowBlobPublicAccess"`
		AllowCrossTenantReplication           bool        `json:"allowCrossTenantReplication"`
		AllowSharedKeyAccess                  bool        `json:"allowSharedKeyAccess"`
		AllowedCopyScope                      interface{} `json:"allowedCopyScope"`
		AzureFilesIdentityBasedAuthentication interface{} `json:"azureFilesIdentityBasedAuthentication"`
		BlobRestoreStatus                     interface{} `json:"blobRestoreStatus"`
		CreationTime                          time.Time   `json:"creationTime"`
		CustomDomain                          interface{} `json:"customDomain"`
		DefaultToOAuthAuthentication          bool        `json:"defaultToOAuthAuthentication"`
		DNSEndpointType                       string      `json:"dnsEndpointType"`
		EnableHTTPSTrafficOnly                bool        `json:"enableHttpsTrafficOnly"`
		EnableNfsV3                           interface{} `json:"enableNfsV3"`
		Encryption                            struct {
			EncryptionIdentity              interface{} `json:"encryptionIdentity"`
			KeySource                       string      `json:"keySource"`
			KeyVaultProperties              interface{} `json:"keyVaultProperties"`
			RequireInfrastructureEncryption bool        `json:"requireInfrastructureEncryption"`
			Services                        struct {
				Blob struct {
					Enabled         bool      `json:"enabled"`
					KeyType         string    `json:"keyType"`
					LastEnabledTime time.Time `json:"lastEnabledTime"`
				} `json:"blob"`
				File struct {
					Enabled         bool      `json:"enabled"`
					KeyType         string    `json:"keyType"`
					LastEnabledTime time.Time `json:"lastEnabledTime"`
				} `json:"file"`
				Queue interface{} `json:"queue"`
				Table interface{} `json:"table"`
			} `json:"services"`
		} `json:"encryption"`
		ExtendedLocation               interface{} `json:"extendedLocation"`
		FailoverInProgress             interface{} `json:"failoverInProgress"`
		GeoReplicationStats            interface{} `json:"geoReplicationStats"`
		ID                             string      `json:"id"`
		Identity                       interface{} `json:"identity"`
		ImmutableStorageWithVersioning interface{} `json:"immutableStorageWithVersioning"`
		IsHnsEnabled                   interface{} `json:"isHnsEnabled"`
		IsLocalUserEnabled             interface{} `json:"isLocalUserEnabled"`
		IsSftpEnabled                  interface{} `json:"isSftpEnabled"`
		IsSkuConversionBlocked         interface{} `json:"isSkuConversionBlocked"`
		KeyCreationTime                struct {
			Key1 time.Time `json:"key1"`
			Key2 time.Time `json:"key2"`
		} `json:"keyCreationTime"`
		KeyPolicy            interface{} `json:"keyPolicy"`
		Kind                 string      `json:"kind"`
		LargeFileSharesState interface{} `json:"largeFileSharesState"`
		LastGeoFailoverTime  interface{} `json:"lastGeoFailoverTime"`
		Location             string      `json:"location"`
		MinimumTLSVersion    string      `json:"minimumTlsVersion"`
		Name                 string      `json:"name"`
		NetworkRuleSet       struct {
			Bypass              string        `json:"bypass"`
			DefaultAction       string        `json:"defaultAction"`
			IPRules             []interface{} `json:"ipRules"`
			Ipv6Rules           []interface{} `json:"ipv6Rules"`
			ResourceAccessRules interface{}   `json:"resourceAccessRules"`
			VirtualNetworkRules []interface{} `json:"virtualNetworkRules"`
		} `json:"networkRuleSet"`
		PrimaryEndpoints struct {
			Blob               string      `json:"blob"`
			Dfs                string      `json:"dfs"`
			File               string      `json:"file"`
			InternetEndpoints  interface{} `json:"internetEndpoints"`
			MicrosoftEndpoints interface{} `json:"microsoftEndpoints"`
			Queue              string      `json:"queue"`
			Table              string      `json:"table"`
			Web                string      `json:"web"`
		} `json:"primaryEndpoints"`
		PrimaryLocation            string        `json:"primaryLocation"`
		PrivateEndpointConnections []interface{} `json:"privateEndpointConnections"`
		ProvisioningState          string        `json:"provisioningState"`
		PublicNetworkAccess        string        `json:"publicNetworkAccess"`
		ResourceGroup              string        `json:"resourceGroup"`
		RoutingPreference          interface{}   `json:"routingPreference"`
		SasPolicy                  interface{}   `json:"sasPolicy"`
		SecondaryEndpoints         struct {
			Blob               string      `json:"blob"`
			Dfs                string      `json:"dfs"`
			File               interface{} `json:"file"`
			InternetEndpoints  interface{} `json:"internetEndpoints"`
			MicrosoftEndpoints interface{} `json:"microsoftEndpoints"`
			Queue              string      `json:"queue"`
			Table              string      `json:"table"`
			Web                string      `json:"web"`
		} `json:"secondaryEndpoints"`
		SecondaryLocation string `json:"secondaryLocation"`
		Sku               struct {
			Name string `json:"name"`
			Tier string `json:"tier"`
		} `json:"sku"`
		StatusOfPrimary                   string      `json:"statusOfPrimary"`
		StatusOfSecondary                 string      `json:"statusOfSecondary"`
		StorageAccountSkuConversionStatus interface{} `json:"storageAccountSkuConversionStatus"`
		Tags                              struct {
			ApplicationServiceNumber string `json:"Application_Service_Number"`
		} `json:"tags"`
		Type string `json:"type"`
	}

	bytes, err := Execute("az", *&Options{}, "storage", "account", "list")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	accounts := make([]StorageAccount, 0)

	err = json.Unmarshal(bytes, &accounts)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		connector := Connector{
			Name: account.Name,
			//Description: account.,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}
