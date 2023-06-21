package model

import (
	"context"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/powershell"
	"go.mongodb.org/mongo-driver/bson"
)

type Domain struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"Name"`
}

func SyncDomains() error {

	domains, err := powershell.GetDomains()
	if err != nil {
		return err
	}

	for _, domain := range *domains {
		log.Println(domain.DomainName)

		_, err := db.FindOne[*Domain](&Domain{}, bson.D{{"name", domain.DomainName}})

		if err != nil {
			newRecord := &Domain{
				Name: domain.DomainName,
			}
			mgm.Coll(newRecord).Create(newRecord)
			log.Println("new")
		}

	}
	return nil
}

func GetDomains() (domains []Domain, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	results, err := mgm.Coll(&Domain{}).Find(context.TODO(), bson.M{})
	if err != nil {

		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var domain Domain
		if err = results.Decode(&domain); err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}
	return domains, nil
}
