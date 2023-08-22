package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/officegraph"
	"github.com/koksmat-com/koksmat/officegraph/sites"
	"github.com/koksmat-com/koksmat/shared"
	"github.com/koksmat-com/koksmat/sharepoint/sites/nexiintra_home"
)

type CountryListItem struct {
	shared.Item `bson:",inline"`
	CountryData *nexiintra_home.SP_Countries `json:"fields,inline"`
}

type Country struct {
	mgm.DefaultModel `bson:",inline"`
	Item             CountryListItem `bson:",inline"`
}

func CreateCountry(sharepoint_Country CountryListItem) (*Country, error) {

	newRecord := &Country{}
	newRecord.Item = sharepoint_Country

	err := mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}

func ImportCountries() error {

	_, token, err := officegraph.GetClient()
	if err != nil {
		return err
	}

	got, err := sites.GetListItems[CountryListItem](token, "sites/nexiintra-home", "Countries", "")
	if err != nil {
		return err
	}

	for _, channel := range *got {

		_, err := CreateCountry(channel)
		if err != nil {
			return err
		}

	}

	return nil
}
