package model

import (
	"context"
	"log"
	"path"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SharedMailbox struct {
	mgm.DefaultModel
	ExchangeObjectId   string   `json:"ExchangeObjectId"`
	Identity           string   `json:"Identity"`
	PrimarySmtpAddress string   `json:"PrimarySmtpAddress"`
	DisplayName        string   `json:"DisplayName"`
	Members            []string `bson:"members,truncate"`
	Owners             []string `bson:"owners,truncate"`
	Readers            []string `bson:"readers,truncate"`
}

type access struct {
	Identity     string   `json:"Identity"`
	User         string   `json:"User"`
	AccessRights []string `json:"AccessRights"`
}

type permission struct {
	Identity          string   `json:"Identity"`
	Trustee           string   `json:"Trustee"`
	AccessControlType string   `json:"AccessControlType"`
	AccessRights      []string `json:"AccessRights"`
}

func GetSharedMailboxes() (cur *mongo.Cursor, err error) {
	return mgm.Coll(&SharedMailbox{}).Find(context.TODO(), bson.M{})
}
func ReadSharedMailboxes(inputFile string) {
	data := io.Readfile[SharedMailbox](inputFile)

	for _, smt := range data {
		log.Println(smt.PrimarySmtpAddress)

		sharedmailboxpermissionsPath := path.Join(path.Dir(inputFile), "sharedmailboxpermissions-"+smt.PrimarySmtpAddress+".json")
		sharedmailboxRecipientpermissionsPath := path.Join(path.Dir(inputFile), "sharedmailboxrecipientPermissions-"+smt.Identity+".json")
		members := []string{}
		owners := []string{}
		readers := []string{}
		tester := []string{}
		mailboxAccess := io.Readfile[access](sharedmailboxpermissionsPath)
		mailboxPermission := io.Readfile[permission](sharedmailboxRecipientpermissionsPath)

		for _, mbp := range mailboxAccess {
			members = append(members, mbp.User)

		}
		for _, mbp := range mailboxPermission {
			tester = append(members, mbp.Trustee)

		}
		log.Println(tester)
		filter := bson.D{{"identity", smt.Identity}}
		result := mgm.Coll(&SharedMailbox{}).FindOne(context.Background(), filter)
		record := &SharedMailbox{}
		result.Decode(record)
		if record.Identity == "" {
			newRecord := &SharedMailbox{

				Identity:           smt.Identity,
				PrimarySmtpAddress: smt.PrimarySmtpAddress,
				DisplayName:        smt.DisplayName,
				Members:            members,
				Owners:             owners,
				Readers:            readers,
			}
			mgm.Coll(newRecord).Create(newRecord)
			log.Println("new")
		} else {
			changedRecord := &SharedMailbox{

				PrimarySmtpAddress: smt.PrimarySmtpAddress,
				DisplayName:        smt.DisplayName,
				Members:            members,
				Owners:             owners,
				Readers:            readers,
			}
			mgm.Coll(changedRecord).Update(changedRecord)
			log.Println("update")
		}

		//recipientPermissions := io.Readfile[model.SharedMailboxType](sharedmailboxRecipientpermissionsPath)

	}

}
