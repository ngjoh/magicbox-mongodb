package db

import (
	"context"
	"errors"
	"log"
	"time"

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

func GetAll[T mgm.Model](record T) (result []T, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	records := []T{}
	defer cancel()

	results, err := mgm.Coll(record).Find(context.TODO(), bson.M{})
	if err != nil {

		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var record T
		if err = results.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
