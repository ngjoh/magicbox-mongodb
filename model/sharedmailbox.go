package model

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/io"
	"github.com/koksmat-com/koksmat/powershell"

	"go.mongodb.org/mongo-driver/bson"
)

const SharedMailboxPrimaryKey = "exchangeobjectid"

type SharedMailbox struct {
	mgm.DefaultModel `bson:",inline"`
	ExchangeObjectId string `json:"ExchangeObjectId"`

	PrimarySmtpAddress string   `json:"PrimarySmtpAddress"`
	DisplayName        string   `json:"DisplayName"`
	CustomAttribute1   string   `json:"CustomAttribute1"` // stores a comma separated list of owners upn's
	Members            []string `bson:"members,truncate"`
	Owners             []string `bson:"owners,truncate"`
	Readers            []string `bson:"readers,truncate"`
}
type SharedMailboxNewRequest struct {
	DisplayName string   `json:"displayName" binding:"required"`
	Alias       string   `json:"alias" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Members     []string `json:"members"`
	Owners      []string `json:"owners"`
	Readers     []string `json:"readers"`
}
type SharedMailboxNewResponce struct {
	*SharedMailboxNewRequest
	PrimarySmtpAddress string `json:"primarySmtpAddress" binding:"required"`
}

type access struct {
	Identity     string   `json:"Identity"`
	User         string   `json:"User"`
	AccessRights []string `json:"AccessRights"` // https://learn.microsoft.com/en-us/previous-versions/office/developer/exchange-server-2010/ff321296(v=exchg.140)
}

type permission struct {
	Identity          string   `json:"Identity"`
	Trustee           string   `json:"Trustee"`
	AccessControlType string   `json:"AccessControlType"`
	AccessRights      []string `json:"AccessRights"` //https://learn.microsoft.com/en-us/powershell/module/exchange/add-recipientpermission?view=exchange-ps#-accessrights
}

type statistics struct {
	WhenCreated  time.Time `json:"WhenCreated"`
	LastSent     string    `json:"LastSent"`
	LastReceived string    `json:"LastReceived"`
}

func DeleteSharedMailbox(Identity string) error {
	_, err := powershell.DeleteSharedMailbox(Identity)

	err2 := db.DeleteOne[*SharedMailbox](&SharedMailbox{}, bson.D{{SharedMailboxPrimaryKey, Identity}})
	if err2 != nil {
		if err != nil {
			return errors.New("Exchange deleted, but database not: " + fmt.Sprint(err2))
		}
	}
	if err != nil {
		return errors.New("Not found in Exchange but deleted in database " + fmt.Sprint(err))
	}
	return nil
}

func CreateSharedMailbox(DisplayName string,
	Alias string,
	Name string,
	Members []string,
	Owners []string,
	Readers []string,
) (sharedMailbox SharedMailbox, err error) {

	newMailbox, err := powershell.CreateSharedMailbox(
		Name,
		DisplayName,
		Alias,
		Members,
		Owners,
		Readers,
	)

	if err != nil {
		return sharedMailbox, err
	}
	newRecord := &SharedMailbox{
		ExchangeObjectId: newMailbox.ExchangeObjectId,

		PrimarySmtpAddress: newMailbox.PrimarySmtpAddress,
		DisplayName:        newMailbox.DisplayName,
		Members:            Members,
		Owners:             Owners,
		Readers:            Readers,
	}
	log.Println("insert")
	err = mgm.Coll(newRecord).Create(newRecord)
	log.Println("inserted")

	return *newRecord, err

}
func UpdateSharedMailbox(
	Identity string,
	DisplayName string,

) (sharedMailbox *SharedMailbox, err error) {
	return db.UpdateOne[*SharedMailbox](&SharedMailbox{}, bson.D{{SharedMailboxPrimaryKey, Identity}}, func(m *SharedMailbox) error {
		_, err = powershell.UpdateSharedMailbox(Identity, DisplayName)

		m.DisplayName = DisplayName
		return err
	})

}

func AddSharedMailboxMembers(
	Identity string,
	Members []string,

) (sharedMailbox *SharedMailbox, err error) {
	r := &SharedMailbox{}

	if err != nil {
		return r, err
	}
	return db.UpdateOne[*SharedMailbox](r, bson.D{{SharedMailboxPrimaryKey, Identity}}, func(m *SharedMailbox) error {
		err = powershell.AddSharedMailboxMembers(Identity, Members)

		m.Members = append(m.Members, Members...)
		return err
	})

}

func AddSharedMailboxReaders(
	Identity string,
	Readers []string,

) (sharedMailbox *SharedMailbox, err error) {
	r := &SharedMailbox{}

	if err != nil {
		return r, err
	}
	return db.UpdateOne[*SharedMailbox](r, bson.D{{SharedMailboxPrimaryKey, Identity}}, func(m *SharedMailbox) error {
		err = powershell.AddSharedMailboxReaders(Identity, Readers)

		m.Readers = append(m.Readers, Readers...)
		return err
	})

}

func AddSharedMailboxOwners(
	Identity string,
	Owners []string,

) (sharedMailbox *SharedMailbox, err error) {
	r := &SharedMailbox{}

	if err != nil {
		return r, err
	}
	return db.UpdateOne[*SharedMailbox](r, bson.D{{SharedMailboxPrimaryKey, Identity}}, func(m *SharedMailbox) error {
		err = powershell.AddSharedMailboxOwners(Identity, Owners)

		m.Owners = append(m.Owners, Owners...)
		return err
	})

}

// filter := bson.D{{"exchangeobjectid", Identity}}
// dbResult := mgm.Coll(&SharedMailbox{}).FindOne(context.Background(), filter)
// record := &SharedMailbox{}
// dbResult.Decode(record)
// id := record.GetID()
// if id == nil {
// 	return sharedMailbox, errors.New("Not found in database")
// } else {
// 	_, err := powershell.UpdateSharedMailbox(Identity, DisplayName)
// 	if err != nil {
// 		return sharedMailbox, err
// 	}
// 	record.DisplayName = DisplayName

// 	log.Println("update")
// 	err = mgm.Coll(record).Update(record)
// 	if err != nil {
// 		log.Println(err)
// 		return *record, errors.New("Exchange updated, but database not: " + fmt.Sprint(err))
// 	}
// 	log.Println("updated")
// }

// if err != nil {
// 	return *record, err
// }

// return *record, nil

//}

func GetSharedMailboxes() (sharedMailboxes []SharedMailbox, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	results, err := mgm.Coll(&SharedMailbox{}).Find(context.TODO(), bson.M{})
	if err != nil {

		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var sharedMailbox SharedMailbox
		if err = results.Decode(&sharedMailbox); err != nil {
			return nil, err
		}

		sharedMailboxes = append(sharedMailboxes, sharedMailbox)
	}
	return sharedMailboxes, nil
}
func ReadSharedMailboxes(inputFile string) {
	io.Waitfor(inputFile)
	data := io.Readfile[SharedMailbox](inputFile)
	max := len(data)
	i := 0
	for _, smt := range data {
		i++
		log.Println(i, max, smt.PrimarySmtpAddress)
		dir := path.Dir(inputFile)
		sharedmailboxpermissionsPath := path.Join(dir, "sharedmailboxpermissions-"+smt.ExchangeObjectId+".json")
		sharedmailboxRecipientpermissionsPath := path.Join(dir, "sharedmailboxrecipientPermission-"+smt.ExchangeObjectId+".json")
		members := []string{}
		owners := strings.Split(smt.CustomAttribute1, ",")
		readers := []string{}
		tester := []string{}
		mailboxAccess := io.Readfile[access](sharedmailboxpermissionsPath)
		mailboxPermission := io.Readfile[permission](sharedmailboxRecipientpermissionsPath)

		for _, mbp := range mailboxAccess {
			if strings.Contains(strings.Join(mbp.AccessRights, ","), "FullAccess") {
				members = append(members, mbp.User)
			} else {
				readers = append(readers, mbp.User)
			}

		}
		for _, mbp := range mailboxPermission {
			tester = append(tester, mbp.Trustee)

		}
		//log.Println(tester)
		filter := bson.D{{"exchangeobjectid", smt.ExchangeObjectId}}
		result := mgm.Coll(&SharedMailbox{}).FindOne(context.Background(), filter)
		record := &SharedMailbox{}
		result.Decode(record)
		if record.ExchangeObjectId == "" {
			newRecord := &SharedMailbox{
				ExchangeObjectId: smt.ExchangeObjectId,

				PrimarySmtpAddress: smt.PrimarySmtpAddress,
				DisplayName:        smt.DisplayName,
				Members:            members,
				Owners:             owners,
				Readers:            readers,
			}
			log.Println("insert")
			mgm.Coll(newRecord).Create(newRecord)
			log.Println("inserted")
		} else {
			changedRecord := &SharedMailbox{

				PrimarySmtpAddress: smt.PrimarySmtpAddress,
				DisplayName:        smt.DisplayName,
				Members:            members,
				Owners:             owners,
				Readers:            readers,
			}
			log.Println("update")
			mgm.Coll(changedRecord).Update(changedRecord)
			log.Println("updated")
		}

	}

}
