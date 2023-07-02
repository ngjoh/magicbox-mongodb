package model

import (
	"encoding/json"
	"net/url"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Blob struct {
	mgm.DefaultModel `bson:",inline"`

	Tag     string `json:"Tag"`
	Content string `json:"Content"`
}

type Foo struct {
	X map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

func GetBlob(unencodedTag string) (*map[string]interface{}, error) {
	tag, err := url.QueryUnescape(unencodedTag)
	if err != nil {
		return nil, err
	}
	b, err := db.FindOne[*Blob](&Blob{}, bson.D{{"tag", tag}})
	if err != nil {
		return nil, err
	}
	f := Foo{}

	json.Unmarshal([]byte(b.Content), &f.X)
	return &f.X, nil
}

func SetBlobString(unencodedTag string, content string) error {

	//tag := url.QueryEscape(unencodedTag)
	db.CreateOrUpdate[*Blob](&Blob{}, bson.D{{"tag", unencodedTag}}, func() (*Blob, error) {
		newBlob := &Blob{
			DefaultModel: mgm.DefaultModel{},
			Tag:          unencodedTag,
			Content:      content,
		}

		return newBlob, nil
	},
		func(record *Blob) error {
			record.Content = content
			return nil
		})
	return nil

}

func SetBlobJSON(tag string, data interface{}) error {
	JSON, err := json.Marshal(data)
	if err != nil {

		return err
	}
	return SetBlobString(tag, string(JSON))
}
