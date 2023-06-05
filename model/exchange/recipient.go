package model

type RecipientsType struct {
	Id                   string   `json:"Id"`
	Guid                 string   `json:"Guid"`
	Alias                string   `json:"Alias"`
	RecipientTypeDetails string   `json:"RecipientTypeDetails"`
	EmailAddresses       []string `json:"EmailAddresses"`
	DisplayName          string   `json:"DisplayName"`
	DistinguishedName    string   `json:"DistinguishedName"`
}
