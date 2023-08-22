package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/officegraph"
	"github.com/koksmat-com/koksmat/officegraph/sites"
	"github.com/koksmat-com/koksmat/shared"
	"github.com/koksmat-com/koksmat/sharepoint/sites/nexiintra_home"
)

type CompanyListItem struct {
	shared.Item `bson:",inline"`
	CompanyData *nexiintra_home.SP_Companies `json:"fields,inline"`
}

type Company struct {
	mgm.DefaultModel `bson:",inline"`
	Item             CompanyListItem `bson:",inline"`
}

func CreateCompany(sharepoint_company CompanyListItem) (company *Company, err error) {

	newRecord := &Company{}
	newRecord.Item = sharepoint_company

	err = mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err

}

func ImportCompanies() error {

	_, token, err := officegraph.GetClient()
	if err != nil {
		return err
	}

	got, err := sites.GetListItems[CompanyListItem](token, "sites/nexiintra-home", "Companies", "")
	if err != nil {
		return err
	}

	for _, channel := range *got {

		_, err := CreateCompany(channel)
		if err != nil {
			return err
		}

	}

	return nil
}
