package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type recipentData[K any] struct {
	DateTime time.Time
	Domain   string
	Type     string
	Data     []K
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
func Save[K any](domain string, subject string, recipients []K) {
	connectionString := goDotEnvVariable("DATABASEURL")

	credential := options.Credential{
		Username: goDotEnvVariable("DATABASEADMIN"),
		Password: goDotEnvVariable("DATABASEPASSWORD"),
	}
	clientOpts := options.Client().ApplyURI(connectionString).SetAuth(credential).SetDirect(true)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}

	recipientData := recipentData[K]{}
	recipientData.DateTime = time.Now()
	recipientData.Domain = domain
	recipientData.Type = subject
	recipientData.Data = recipients

	log.Println("Inserting", len(recipients), subject)

	databaseName := goDotEnvVariable("DATABASE")
	_, insertError := client.Database(databaseName).Collection("inputdata").InsertOne(context.TODO(), recipientData)
	if insertError != nil {
		panic(insertError)
	}
	log.Println("Successfully inserted", len(recipients), subject)
}
