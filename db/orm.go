package db

import (
	"context"
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOne[M mgm.Model](
	record M,
	filter bson.D,
	interact func(record M) error,

) (updateRecord M, err error) {

	dbResult := mgm.Coll(record).FindOne(context.Background(), filter)
	dbErr := dbResult.Err()
	if dbErr != nil {
		return record, errors.New("Not found in database")
	}
	dbResult.Decode(record)

	err = interact(record)

	if err != nil {
		log.Println(err)
		return record, err
	}
	log.Println("update")
	err = mgm.Coll(record).Update(record)
	if err != nil {
		log.Println(err)
		return record, err
	}
	log.Println("updated")

	return record, nil

}
