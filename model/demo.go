package model

import "github.com/kamva/mgm/v3"

type Demo struct {
	mgm.DefaultModel `bson:",inline"`
	Hello            string `json:"hello"`
}

func NewDemo(hello string) error {

	newRecord := &Demo{
		DefaultModel: mgm.DefaultModel{},
		Hello:        hello,
	}

	return mgm.Coll(newRecord).Create(newRecord)

}
