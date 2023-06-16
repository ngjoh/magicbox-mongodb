package db

import (
	"context"
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

const RecordNotFound = "Not found in database"

func FindOne[M mgm.Model](record M,
	filter bson.D) (foundRecord M, err error) {
	dbResult := mgm.Coll(record).FindOne(context.Background(), filter)
	dbErr := dbResult.Err()
	if dbErr != nil {

		return record, errors.New(RecordNotFound)
	}
	dbResult.Decode(record)
	return record, nil

}

func DeleteOne[M mgm.Model](record M,
	filter bson.D) (err error) {
	dbResult := mgm.Coll(record).FindOne(context.Background(), filter)
	dbErr := dbResult.Err()
	if dbErr != nil {
		return errors.New(RecordNotFound)
	}
	err = mgm.Coll(record).Delete(record)
	return nil

}
func UpdateOne[M mgm.Model](
	record M,
	filter bson.D,
	interact func(record M) error,

) (updateRecord M, err error) {

	dbResult := mgm.Coll(record).FindOne(context.Background(), filter)
	dbErr := dbResult.Err()
	if dbErr != nil {
		return record, errors.New(RecordNotFound)
	}
	dbResult.Decode(record)

	err = interact(record)

	if err != nil {
		log.Println(err)
		return record, err
	}

	err = mgm.Coll(record).Update(record)
	if err != nil {
		log.Println(err)
		return record, err
	}

	return record, nil

}
