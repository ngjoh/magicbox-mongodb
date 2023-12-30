package connectors

import (
	"encoding/json"
)

func M365Context() ([]Connector, error) {
	type Context struct {
		TenantID            string      `json:"tenantId"`
		FederationBrandName interface{} `json:"federationBrandName"`
		DisplayName         string      `json:"displayName"`
		DefaultDomainName   string      `json:"defaultDomainName"`
	}

	bytes, err := Execute("m365", *&Options{}, "tenant", "info", "get")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	config := Context{}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	connector := Connector{
		Name: config.DisplayName,
		ID:   config.TenantID,

		IsCurrent: true,
	}
	connectors = append(connectors, connector)

	return connectors, nil
}

func M365Sites() ([]Connector, error) {
	type Site struct {
		ObjectType                                  string        `json:"_ObjectType_"`
		ObjectIdentity                              string        `json:"_ObjectIdentity_"`
		AllowDownloadingNonWebViewableFiles         bool          `json:"AllowDownloadingNonWebViewableFiles"`
		AllowEditing                                bool          `json:"AllowEditing"`
		AllowSelfServiceUpgrade                     bool          `json:"AllowSelfServiceUpgrade"`
		AnonymousLinkExpirationInDays               int           `json:"AnonymousLinkExpirationInDays"`
		ArchiveStatus                               string        `json:"ArchiveStatus"`
		AuthContextStrength                         interface{}   `json:"AuthContextStrength"`
		AuthenticationContextLimitedAccess          bool          `json:"AuthenticationContextLimitedAccess"`
		AuthenticationContextName                   interface{}   `json:"AuthenticationContextName"`
		AverageResourceUsage                        int           `json:"AverageResourceUsage"`
		BlockDownloadLinksFileType                  int           `json:"BlockDownloadLinksFileType"`
		BlockDownloadMicrosoft365GroupIds           interface{}   `json:"BlockDownloadMicrosoft365GroupIds"`
		BlockDownloadPolicy                         bool          `json:"BlockDownloadPolicy"`
		BlockGuestsAsSiteAdmin                      int           `json:"BlockGuestsAsSiteAdmin"`
		BonusDiskQuota                              int           `json:"BonusDiskQuota"`
		ClearRestrictedAccessControl                bool          `json:"ClearRestrictedAccessControl"`
		CommentsOnSitePagesDisabled                 bool          `json:"CommentsOnSitePagesDisabled"`
		CompatibilityLevel                          int           `json:"CompatibilityLevel"`
		ConditionalAccessPolicy                     int           `json:"ConditionalAccessPolicy"`
		CurrentResourceUsage                        int           `json:"CurrentResourceUsage"`
		DefaultLinkPermission                       int           `json:"DefaultLinkPermission"`
		DefaultLinkToExistingAccess                 bool          `json:"DefaultLinkToExistingAccess"`
		DefaultLinkToExistingAccessReset            bool          `json:"DefaultLinkToExistingAccessReset"`
		DefaultShareLinkRole                        int           `json:"DefaultShareLinkRole"`
		DefaultShareLinkScope                       int           `json:"DefaultShareLinkScope"`
		DefaultSharingLinkType                      int           `json:"DefaultSharingLinkType"`
		DenyAddAndCustomizePages                    int           `json:"DenyAddAndCustomizePages"`
		Description                                 string        `json:"Description"`
		DisableAppViews                             int           `json:"DisableAppViews"`
		DisableCompanyWideSharingLinks              int           `json:"DisableCompanyWideSharingLinks"`
		DisableFlows                                int           `json:"DisableFlows"`
		ExcludeBlockDownloadPolicySiteOwners        bool          `json:"ExcludeBlockDownloadPolicySiteOwners"`
		ExcludeBlockDownloadSharePointGroups        []interface{} `json:"ExcludeBlockDownloadSharePointGroups"`
		ExcludedBlockDownloadGroupIds               []interface{} `json:"ExcludedBlockDownloadGroupIds"`
		ExternalUserExpirationInDays                int           `json:"ExternalUserExpirationInDays"`
		GroupID                                     string        `json:"GroupId"`
		GroupOwnerLoginName                         interface{}   `json:"GroupOwnerLoginName"`
		HasHolds                                    bool          `json:"HasHolds"`
		HubSiteID                                   string        `json:"HubSiteId"`
		IBMode                                      interface{}   `json:"IBMode"`
		IBSegments                                  interface{}   `json:"IBSegments"`
		IBSegmentsToAdd                             interface{}   `json:"IBSegmentsToAdd"`
		IBSegmentsToRemove                          interface{}   `json:"IBSegmentsToRemove"`
		IsGroupOwnerSiteAdmin                       bool          `json:"IsGroupOwnerSiteAdmin"`
		IsHubSite                                   bool          `json:"IsHubSite"`
		IsTeamsChannelConnected                     bool          `json:"IsTeamsChannelConnected"`
		IsTeamsConnected                            bool          `json:"IsTeamsConnected"`
		LastContentModifiedDate                     string        `json:"LastContentModifiedDate"`
		Lcid                                        int           `json:"Lcid"`
		LimitedAccessFileType                       int           `json:"LimitedAccessFileType"`
		ListsShowHeaderAndNavigation                bool          `json:"ListsShowHeaderAndNavigation"`
		LockIssue                                   interface{}   `json:"LockIssue"`
		LockState                                   string        `json:"LockState"`
		LoopDefaultSharingLinkRole                  int           `json:"LoopDefaultSharingLinkRole"`
		LoopDefaultSharingLinkScope                 int           `json:"LoopDefaultSharingLinkScope"`
		LoopOverrideSharingCapability               bool          `json:"LoopOverrideSharingCapability"`
		LoopSharingCapability                       int           `json:"LoopSharingCapability"`
		MediaTranscription                          int           `json:"MediaTranscription"`
		OverrideBlockUserInfoVisibility             int           `json:"OverrideBlockUserInfoVisibility"`
		OverrideSharingCapability                   bool          `json:"OverrideSharingCapability"`
		OverrideTenantAnonymousLinkExpirationPolicy bool          `json:"OverrideTenantAnonymousLinkExpirationPolicy"`
		OverrideTenantExternalUserExpirationPolicy  bool          `json:"OverrideTenantExternalUserExpirationPolicy"`
		Owner                                       string        `json:"Owner"`
		OwnerEmail                                  interface{}   `json:"OwnerEmail"`
		OwnerLoginName                              interface{}   `json:"OwnerLoginName"`
		OwnerName                                   interface{}   `json:"OwnerName"`
		PWAEnabled                                  int           `json:"PWAEnabled"`
		ReadOnlyAccessPolicy                        bool          `json:"ReadOnlyAccessPolicy"`
		ReadOnlyForBlockDownloadPolicy              bool          `json:"ReadOnlyForBlockDownloadPolicy"`
		ReadOnlyForUnmanagedDevices                 bool          `json:"ReadOnlyForUnmanagedDevices"`
		RelatedGroupID                              string        `json:"RelatedGroupId"`
		RequestFilesLinkEnabled                     bool          `json:"RequestFilesLinkEnabled"`
		RequestFilesLinkExpirationInDays            int           `json:"RequestFilesLinkExpirationInDays"`
		RestrictedAccessControl                     bool          `json:"RestrictedAccessControl"`
		RestrictedAccessControlGroups               interface{}   `json:"RestrictedAccessControlGroups"`
		RestrictedAccessControlGroupsToAdd          interface{}   `json:"RestrictedAccessControlGroupsToAdd"`
		RestrictedAccessControlGroupsToRemove       interface{}   `json:"RestrictedAccessControlGroupsToRemove"`
		RestrictedToRegion                          int           `json:"RestrictedToRegion"`
		SandboxedCodeActivationCapability           int           `json:"SandboxedCodeActivationCapability"`
		SensitivityLabel                            string        `json:"SensitivityLabel"`
		SensitivityLabel2                           interface{}   `json:"SensitivityLabel2"`
		SetOwnerWithoutUpdatingSecondaryAdmin       bool          `json:"SetOwnerWithoutUpdatingSecondaryAdmin"`
		SharingAllowedDomainList                    interface{}   `json:"SharingAllowedDomainList"`
		SharingBlockedDomainList                    interface{}   `json:"SharingBlockedDomainList"`
		SharingCapability                           int           `json:"SharingCapability"`
		SharingDomainRestrictionMode                int           `json:"SharingDomainRestrictionMode"`
		SharingLockDownCanBeCleared                 bool          `json:"SharingLockDownCanBeCleared"`
		SharingLockDownEnabled                      bool          `json:"SharingLockDownEnabled"`
		ShowPeoplePickerSuggestionsForGuestUsers    bool          `json:"ShowPeoplePickerSuggestionsForGuestUsers"`
		SiteDefinedSharingCapability                int           `json:"SiteDefinedSharingCapability"`
		SiteID                                      string        `json:"SiteId"`
		SocialBarOnSitePagesDisabled                bool          `json:"SocialBarOnSitePagesDisabled"`
		Status                                      string        `json:"Status"`
		StorageMaximumLevel                         int           `json:"StorageMaximumLevel"`
		StorageQuotaType                            interface{}   `json:"StorageQuotaType"`
		StorageUsage                                int           `json:"StorageUsage"`
		StorageWarningLevel                         int           `json:"StorageWarningLevel"`
		TeamsChannelType                            int           `json:"TeamsChannelType"`
		Template                                    string        `json:"Template"`
		TimeZoneID                                  int           `json:"TimeZoneId"`
		Title                                       string        `json:"Title"`
		TitleTranslations                           interface{}   `json:"TitleTranslations"`
		URL                                         string        `json:"Url"`
		UserCodeMaximumLevel                        int           `json:"UserCodeMaximumLevel"`
		UserCodeWarningLevel                        int           `json:"UserCodeWarningLevel"`
		WebsCount                                   int           `json:"WebsCount"`
	}
	bytes, err := Execute("m365", *&Options{}, "spo", "site", "list", "-o", "json")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	sites := []Site{}

	err = json.Unmarshal(bytes, &sites)
	if err != nil {
		return nil, err
	}
	for _, site := range sites {
		connector := Connector{
			Name:        site.Title,
			Description: site.Description,
			Url:         site.URL,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}
