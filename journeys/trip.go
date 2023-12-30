package journeys

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trip struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id" bson:"id"`
	DetailedJourney  `bson:",inline"`
	Cargo            `bson:",inline"`
}

// func SyncCargo() error {

// 	cargos, err := powershell.GetDomains()
// 	if err != nil {
// 		return err
// 	}

// 	for _, cargo := range *cargos {
// 		log.Println(cargo.DomainName)

// 		_, err := db.FindOne[*Cargo](&Cargo{}, bson.D{{"name", cargo.DomainName}})

// 		if err != nil {
// 			newRecord := &Cargo{
// 				Name: cargo.DomainName,
// 			}
// 			mgm.Coll(newRecord).Create(newRecord)
// 			log.Println("new")
// 		}

//		}
//		return nil
//	}
func (trip *Trip) Collection() *mgm.Collection {
	// Create new client

	client, err := mgm.NewClient(options.Client().ApplyURI(config.MongoConnectionString()))
	if err != nil {
		panic(err)
	}
	// Get the model's db
	db := client.Database("map-of-cava")
	// return the model's custom collection
	return mgm.NewCollection(db, "trips")
}

// func GetJourneys() (trips []Trip, err error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	results, err := mgm.Coll(&Trip{}).Find(context.TODO(), bson.M{})
// 	if err != nil {

// 		return nil, err
// 	}
// 	defer results.Close(ctx)
// 	for results.Next(ctx) {
// 		var trip Trip
// 		if err = results.Decode(&trip); err != nil {
// 			return nil, err
// 		}

// 		trips = append(trips, trip)
// 	}
// 	return trips, nil
// }
