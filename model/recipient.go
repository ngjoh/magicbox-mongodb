package model

import (
	"context"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/io"
	"go.mongodb.org/mongo-driver/bson"
)

type Recipient struct {
	mgm.DefaultModel     `bson:",inline"`
	Id                   string   `json:"Identity"`
	Guid                 string   `json:"ExternalDirectoryObjectId"`
	Alias                string   `json:"Alias"`
	RecipientTypeDetails string   `json:"RecipientTypeDetails"`
	EmailAddresses       []string `json:"EmailAddresses"`
	DisplayName          string   `json:"DisplayName"`
	DistinguishedName    string   `json:"DistinguishedName"`
}

func FindRecipientByAddress(address string) (*Recipient, error) {
	//atSymbolEscaped := strings.ReplaceAll(address, "@", "\\@")
	return db.FindOne(&Recipient{}, bson.D{
		{"$or",
			bson.A{
				bson.D{{"guid", address}},
				bson.D{{"emailaddresses", bson.D{
					{"$regex", fmt.Sprintf("smtp:%s", address)},
					{"$options", "i"},
				}}},
				bson.D{{"alias", bson.D{
					{"$regex", address},
					{"$options", "i"},
				}}},
			},
		},
	})
}

func ImportRecipients(inputFile string) {
	io.Waitfor(inputFile)
	data := io.Readfile[Recipient](inputFile)

	for _, rcp := range data {
		log.Println(rcp.DisplayName)

		filter := bson.D{{"id", rcp.Id}}
		result := mgm.Coll(&Recipient{}).FindOne(context.Background(), filter)
		record := &Recipient{}
		result.Decode(record)
		if record.Id == "" {
			newRecord := &Recipient{

				Id:                   rcp.Id,
				Guid:                 rcp.Guid,
				Alias:                rcp.Alias,
				RecipientTypeDetails: rcp.RecipientTypeDetails,
				EmailAddresses:       rcp.EmailAddresses,
				DisplayName:          rcp.DisplayName,
				DistinguishedName:    rcp.DistinguishedName,
			}
			mgm.Coll(newRecord).Create(newRecord)
			log.Println("new")
		} else {
			changedRecord := &Recipient{

				Guid:                 rcp.Guid,
				Alias:                rcp.Alias,
				RecipientTypeDetails: rcp.RecipientTypeDetails,
				EmailAddresses:       rcp.EmailAddresses,
				DisplayName:          rcp.DisplayName,
				DistinguishedName:    rcp.DistinguishedName,
			}
			mgm.Coll(changedRecord).Update(changedRecord)
			log.Println("update")
		}

	}
}

func FindPrimarySMTPAddress(addresses []string) string {
	for _, address := range addresses {
		if len(address) > 4 && address[0:5] == "SMTP:" {
			return address[5:]
		}
	}
	return ""
}

func TranslateRecipients(addresses []string) (emails []string, guids []string) {
	for _, address := range addresses {
		recipient, err := FindRecipientByAddress(address)
		if err == nil {
			emails = append(emails, FindPrimarySMTPAddress(recipient.EmailAddresses))
			guids = append(guids, recipient.Guid)
		}

	}
	return emails, guids
}
