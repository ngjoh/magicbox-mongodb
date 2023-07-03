package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
)

type Country struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `json:"Name" example:"Sweden"`
	Code string `json:"code" example:"SE"`
}

type BusinessGroupUnit struct {
	mgm.DefaultModel `bson:",inline"`

	Name  string `json:"name" example:"Finance"`
	Short string `json:"short" example:"finance"`
}

func Countries() ([]*Country, error) {

	countries, err := db.GetAll[*Country](&Country{})
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func NewCountry(name string, code string) error {
	newRecord := &Country{
		DefaultModel: mgm.DefaultModel{},
		Name:         name,
		Code:         code,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}

func BusinessGroupUnits() ([]*BusinessGroupUnit, error) {

	u, err := db.GetAll[*BusinessGroupUnit](&BusinessGroupUnit{})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NewBusinessGroupUnit(name string, code string) error {
	newRecord := &BusinessGroupUnit{
		DefaultModel: mgm.DefaultModel{},
		Name:         name,
		Short:        code,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}
