package officegraph

import (
	"context"
	"log"

	"github.com/koksmat-com/koksmat/db"
	"go.mongodb.org/mongo-driver/bson"
)

func AggregateEvents() error {
	// Requires the MongoDB Go Driver
	// https://go.mongodb.org/mongo-driver
	ctx := context.TODO()

	// Connect to MongoDB
	client := db.Connect()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Open an aggregation cursor
	coll := client.Database("christianiabpos").Collection("webhook_events")
	_, err := coll.Aggregate(ctx, bson.A{
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"$concat",
								bson.A{
									"$clientstate",
									"|",
									"$changetype",
									"|",
									"$resourcedata.odataid",
									"|",
									"$resourcedata.odataetag",
								},
							},
						},
					},
					{"created_at", bson.D{{"$first", "$created_at"}}},
					{"changetype", bson.D{{"$first", "$changetype"}}},
					{"tenantid", bson.D{{"$first", "$tenantid"}}},
					{"odataid", bson.D{{"$first", "$resourcedata.odataid"}}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		bson.D{{"$out", "webhook_events_condensed"}},
	})

	return err

}
