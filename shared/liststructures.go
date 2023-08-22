package shared

import "time"

type LookupReference struct {
	LookupId    int    `json:"LookupId"`
	LookupValue string `json:"LookupValue"`
}
type StandardFields struct {
	OdataEtag       string `json:"@odata.etag"`
	Title           string `json:"Title"`
	LinkTitle       string `json:"LinkTitle"`
	ID              string `json:"id"`
	ContentType     string `json:"ContentType"`
	UIVersionString string `json:"_UIVersionString"`
}
type Item struct {
	OdataEtag            string    `json:"@odata.etag"`
	CreatedDateTime      time.Time `json:"createdDateTime"`
	ETag                 string    `json:"eTag"`
	ID                   string    `json:"id"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	WebURL               string    `json:"webUrl"`
	CreatedBy            struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	LastModifiedBy struct {
		User struct {
			Email       string `json:"email"`
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	ParentReference struct {
		ID     string `json:"id"`
		SiteID string `json:"siteId"`
	} `json:"parentReference"`
	ContentType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"contentType"`
	FieldsOdataContext string `json:"fields@odata.context"`
}
