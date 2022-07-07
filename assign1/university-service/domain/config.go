package domain

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/redis.v5"
)

var (
	Db          *mongo.Database
	PSQLDB      *sql.DB
	RedisClient *redis.Client
)

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "1234"
	dbname     = "Universities"
	uri        = "mongodb://localhost:27017"
	DefaultTTL = 20 * time.Second
)

func ConnMongoDB() {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}
	Db = client.Database("University")
	fmt.Println("Successfuly Connected to the mongodb")
}

func ConnPostgress() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	PSQLDB = db
	fmt.Println("Successfully connected with pgSql !")

}

func ClosePgSQL() {
	err := PSQLDB.Close()
	if err != nil {
		panic(err)
	}
}

func RedisCli() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected with Redis!")
}
