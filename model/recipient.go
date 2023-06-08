package model

import (
	"context"
	"log"

	"github.com/koksmat-com/koksmat/db"
	"go.mongodb.org/mongo-driver/bson"
)

type RecipientType struct {
	Id                   string   `json:"Id"`
	Guid                 string   `json:"Guid"`
	Alias                string   `json:"Alias"`
	RecipientTypeDetails string   `json:"RecipientTypeDetails"`
	EmailAddresses       []string `json:"EmailAddresses"`
	DisplayName          string   `json:"DisplayName"`
	DistinguishedName    string   `json:"DistinguishedName"`
}

func SyncRecipients(database string) {

	// Connect to MongoDB
	client := db.Connect()

	// Open an aggregation cursor
	coll := client.Database(database).Collection("inputdata")

	// Requires the MongoDB Go Driver
	// https://go.mongodb.org/mongo-driver
	ctx := context.TODO()

	_, err := coll.Aggregate(ctx, bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"domain", "exchange"},
					{"type", "recipients"},
				},
			},
		},
		bson.D{{"$limit", 1}},
		bson.D{{"$limit", 10000000000}},
		bson.D{{"$unwind", bson.D{{"path", "$data"}}}},
		bson.D{{"$sort", bson.D{{"datetime", -1}}}},
		bson.D{{"$unwind", bson.D{{"path", "$data"}}}},
		bson.D{{"$replaceRoot", bson.D{{"newRoot", "$data"}}}},
		bson.D{{"$unwind", bson.D{{"path", "$emailaddresses"}}}},
		bson.D{
			{"$project",
				bson.D{
					{"emailaddresses", 1},
					{"guid", 1},
					{"displayname", 1},
					{"recipienttypedetails", 1},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"address2",
						bson.D{
							{"$split",
								bson.A{
									"$emailaddresses",
									":",
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{

					{"recipienttype", "$recipienttypedetails"},
					{"address", bson.D{{"$toLower", bson.D{{"$last", "$address2"}}}}},
					{"type", bson.D{{"$first", "$address2"}}},
				},
			},
		},
		bson.D{
			{"$match",
				bson.D{
					{"$or",
						bson.A{
							bson.D{{"type", "SMTP"}},
							bson.D{{"type", "smtp"}},
						},
					},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{

					{"guid", 1},
					{"recipienttype", 1},
					{"displayname", 1},
					{"address", 1},
					{"type", 1},
				},
			},
		},
		bson.D{{"$out", "smtp_addresses"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted smtpaddresses")
	if err != nil {
		log.Fatal(err)
	}
}
