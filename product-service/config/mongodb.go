package config

import (
	"context"
	"github.com/Ahmad-mufied/aws-academy-group-2/product-service/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"time"
)

var DB *mongo.Client

func InitMongo() {

	// Check if the environment variable is set
	if Viper.Get("MONGO_URI") == "" {
		log.Fatal("MONGO_URI is not set")
	}

	log.Println("Connecting to mongo " + "...")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	counts := 0

	for {
		uri := Viper.Get("MONGO_URI").(string)
		connect, err := OpenDB(ctx, uri)

		if counts >= 10 {
			log.Fatal("Failed to connect to mongodb")
		}

		counts++

		if err != nil {
			log.Printf("Failed to connect to mongodb, trying again in 5 seconds : %v, count: %d/10", err, counts)
			time.Sleep(5 * time.Second)
			LoadEnv()
			continue
		}

		err = connect.Ping(ctx, nil)
		if err != nil {
			log.Printf("Failed to ping mongodb, trying again in 5 seconds : %v, count: %d/10", err, counts)
			time.Sleep(5 * time.Second)
			LoadEnv()
			continue
		}

		DB = connect
		log.Println("Connected to mongodb")
		break
	}
}

func OpenDB(ctx context.Context, uri string) (*mongo.Client, error) {
	var MaxPoolSize uint64 = 100
	var MinPoolSize uint64 = 5
	var MaxConnIdleTime time.Duration = 30 * time.Second
	var MaxConnecting uint64 = 10

	// Create custom registry with UUID codec
	registry := bson.NewRegistryBuilder().
		RegisterCodec(
			reflect.TypeOf(uuid.UUID{}),
			&utils.UUIDCodec{},
		).
		Build()

	clientOptions := options.Client().ApplyURI(uri).SetRegistry(registry)
	clientOptions.MaxPoolSize = &MaxPoolSize
	clientOptions.MinPoolSize = &MinPoolSize
	clientOptions.MaxConnIdleTime = &MaxConnIdleTime
	clientOptions.MaxConnecting = &MaxConnecting

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = connect.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return connect, nil
}
