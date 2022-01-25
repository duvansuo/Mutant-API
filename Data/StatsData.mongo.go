package data

import (
	models "Mutant-Api/Models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const MUTANT_STATS_KEY = "mutant_stats"

type statsMongoData struct {
	client *mongo.Client
}

// var dnas = getSession().DB("pruebamutant").C("dnas")

// var adn models.Dna

// func getSession() *mgo.Session {
// 	session, err := mgo.Dial("mongodb://127.0.0.1")

// 	if err != nil {
// 		panic(err)
// 	}

// 	return session
// }
func newStatsMongoData() *statsMongoData {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return &statsMongoData{
		client,
	}
}
func (repo *statsMongoData) AddStats(statsDto models.MutantStatsDTO) error {
	collections := repo.client.Database("PruebaMutant").Collection("dna")
	_, err := collections.InsertOne(context.TODO(), statsDto)
	if err != nil {
		return err
	}
	return nil

}

func (repo *statsMongoData) GetStats() (models.MutantStatsDTO, error) {
	ctx := context.Background()
	collections := repo.client.Database("PruebaMutant").Collection("dna")
	var result bson.M
	err := collections.FindOne(ctx, bson.M{}).Decode(&result)

	if err != nil {
		return models.MutantStatsDTO{}, err
	}

	var statsDTO models.MutantStatsDTO
	bsonBytes, _ := bson.Marshal(result)
	json.Unmarshal(bsonBytes, &statsDTO)

	if err != nil {
		return models.MutantStatsDTO{}, err
	}
	return statsDTO, err
}
