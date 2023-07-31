package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"github.com/koksmat-com/koksmat/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userCollectionName = "user"

type Credential struct {
	Service  string    `json:"service"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Expires  time.Time `json:"exprires" `
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UPN              string       `json:"upn"`
	DisplayName      string       `json:"displayname"`
	Credentials      []Credential `json:"credentials"`
}

func (user *User) Collection() *mgm.Collection {
	// Create new client

	client, err := mgm.NewClient(options.Client().ApplyURI(config.MongoConnectionString()))
	if err != nil {
		panic(err)
	}
	// Get the model's db
	db := client.Database("magicbox")
	// return the model's custom collection
	return mgm.NewCollection(db, userCollectionName)
}

func CreateUser(auth Authorization, upn string, name string) (*User, error) {
	newRecord := &User{

		UPN:         upn,
		DisplayName: name,
		Credentials: []Credential{},
	}

	err := mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}
func UpdateUserCredentials(
	auth Authorization,
	UPN string,
	Credentials []Credential,

) (user *User, err error) {

	return db.UpdateOne[*User](&User{}, bson.D{{"UPN", UPN}}, func(m *User) error {

		m.Credentials = Credentials
		return err
	})

}
func GetUser(auth Authorization, UPN string) (*User, error) {
	return db.FindOne[*User](&User{}, bson.D{{"UPN", UPN}})
}

func GetUsers(auth Authorization) ([]*User, error) {
	return db.GetAll[*User](&User{})
}
