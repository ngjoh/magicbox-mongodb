package model

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"path"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/io"
	"github.com/koksmat-com/koksmat/magicbox"

	"go.mongodb.org/mongo-driver/bson"
)

type SharedMailbox struct {
	mgm.DefaultModel
	ExchangeObjectId   string   `json:"ExchangeObjectId"`
	Identity           string   `json:"Identity"`
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

	return errors.New("Not implemented")
}

func CreateSharedMailbox(DisplayName string,
	Alias string,
	Name string,
	Members []string,
	Owners []string,
	Readers []string,
) (sharedMailbox SharedMailboxNewResponce, err error) {
	request := SharedMailboxNewRequest{
		DisplayName: DisplayName,
		Alias:       Alias,
		Name:        Name,
		Members:     Members,
		Owners:      Owners,
		Readers:     Readers,
	}

	result, err := magicbox.Powerpack(request)
	if err != nil {
		return sharedMailbox, err
	}

	response := SharedMailboxNewResponce{}
	json.Unmarshal(result, &response)
	return response, nil

}
func UpdateSharedMailbox(
	Identity string,
	DisplayName string,

) (sharedMailbox SharedMailboxNewResponce, err error) {
	return sharedMailbox, errors.New("Not implemented")

}

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
				ExchangeObjectId:   smt.ExchangeObjectId,
				Identity:           smt.Identity,
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
				Identity:           smt.Identity,
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
