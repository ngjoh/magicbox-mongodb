package officegraph

import (
	"time"

	"github.com/koksmat-com/koksmat/util"
)

type Device struct {
	ID                            string    `json:"id"`
	DeletedDateTime               any       `json:"deletedDateTime"`
	AccountEnabled                bool      `json:"accountEnabled"`
	ApproximateLastSignInDateTime time.Time `json:"approximateLastSignInDateTime"`
	ComplianceExpirationDateTime  any       `json:"complianceExpirationDateTime"`
	CreatedDateTime               time.Time `json:"createdDateTime"`
	DeviceCategory                any       `json:"deviceCategory"`
	DeviceID                      string    `json:"deviceId"`
	DeviceMetadata                any       `json:"deviceMetadata"`
	DeviceOwnership               string    `json:"deviceOwnership"`
	DeviceVersion                 int       `json:"deviceVersion"`
	DisplayName                   string    `json:"displayName"`
	DomainName                    any       `json:"domainName"`
	EnrollmentProfileName         any       `json:"enrollmentProfileName"`
	EnrollmentType                string    `json:"enrollmentType"`
	ExternalSourceName            any       `json:"externalSourceName"`
	IsCompliant                   bool      `json:"isCompliant"`
	IsManaged                     bool      `json:"isManaged"`
	IsRooted                      bool      `json:"isRooted"`
	ManagementType                string    `json:"managementType"`
	Manufacturer                  string    `json:"manufacturer"`
	MdmAppID                      any       `json:"mdmAppId"`
	Model                         string    `json:"model"`
	OnPremisesLastSyncDateTime    any       `json:"onPremisesLastSyncDateTime"`
	OnPremisesSyncEnabled         any       `json:"onPremisesSyncEnabled"`
	OperatingSystem               string    `json:"operatingSystem"`
	OperatingSystemVersion        string    `json:"operatingSystemVersion"`
	PhysicalIds                   []any     `json:"physicalIds"`
	ProfileType                   any       `json:"profileType"`
	RegistrationDateTime          time.Time `json:"registrationDateTime"`
	SourceType                    any       `json:"sourceType"`
	SystemLabels                  []any     `json:"systemLabels"`
	TrustType                     string    `json:"trustType"`
	ExtensionAttributes           struct {
		ExtensionAttribute1  any `json:"extensionAttribute1"`
		ExtensionAttribute2  any `json:"extensionAttribute2"`
		ExtensionAttribute3  any `json:"extensionAttribute3"`
		ExtensionAttribute4  any `json:"extensionAttribute4"`
		ExtensionAttribute5  any `json:"extensionAttribute5"`
		ExtensionAttribute6  any `json:"extensionAttribute6"`
		ExtensionAttribute7  any `json:"extensionAttribute7"`
		ExtensionAttribute8  any `json:"extensionAttribute8"`
		ExtensionAttribute9  any `json:"extensionAttribute9"`
		ExtensionAttribute10 any `json:"extensionAttribute10"`
		ExtensionAttribute11 any `json:"extensionAttribute11"`
		ExtensionAttribute12 any `json:"extensionAttribute12"`
		ExtensionAttribute13 any `json:"extensionAttribute13"`
		ExtensionAttribute14 any `json:"extensionAttribute14"`
		ExtensionAttribute15 any `json:"extensionAttribute15"`
	} `json:"extensionAttributes"`
	AlternativeSecurityIds []struct {
		Type             int    `json:"type"`
		IdentityProvider any    `json:"identityProvider"`
		Key              string `json:"key"`
	} `json:"alternativeSecurityIds"`
}

func GetDevices() (*[]Device, error) {
	_, token, err := GetClient()
	if err != nil {
		return nil, err
	}
	endPoint := `https://graph.microsoft.com/v1.0/devices`

	items, err := util.HttpGet[Device](token, endPoint)
	if err != nil {
		return nil, err
	}
	return items, nil
}
