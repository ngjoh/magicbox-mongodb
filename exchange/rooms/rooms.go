package rooms

type Room struct {
	Capacity           int    `json:"Capacity"`
	ProvisioningStatus string `json:"Provisioning_x0020_Status"`
	Title              string `json:"Title"`
	Metadata           struct {
		Etag string `json:"etag"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"__metadata"`
}
