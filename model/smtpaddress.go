package model

import "github.com/kamva/mgm/v3"

type SmtpAddress struct {
	mgm.DefaultModel `bson:",inline"`
	Guid             string `json:"guid"`
	Displayname      string `json:"displayname"`
	RecipientType    string `json:"recipienttype"`
	Address          string `json:"address"`
}

func TestSMTP() error {
	smtp := &SmtpAddress{
		Guid:          "123",
		Displayname:   "kjjl",
		RecipientType: "sdfds",
		Address:       "asdfds",
	}
	return mgm.Coll(smtp).Create(smtp)

}
