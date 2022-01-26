package data

import (
	models "Mutant-Api/Models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type statsMongoData struct {
	database *mongo.Collection
}

func newStatsMongoData() *statsMongoData {
	fmt.Println("ingreso newStatsMongoData")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	collection := client.Database("PruebaMutant").Collection("dna")
	return &statsMongoData{
		collection,
	}
}
func (repo *statsMongoData) AddStats(statsDto models.MutantStatsDTO) error {

	_, err := repo.database.ReplaceOne(context.TODO(), bson.M{}, statsDto)
	if err != nil {
		return err
	}
	return nil

}

func (repo *statsMongoData) GetStats() (models.MutantStatsDTO, error) {
	fmt.Println("ingreso a repo")
	ctx := context.Background()
	var statsDTO models.MutantStatsDTO
	err := repo.database.FindOne(ctx, bson.M{}).Decode(&statsDTO)
	if err != nil {
		return models.MutantStatsDTO{}, err
	}
	fmt.Println(statsDTO)
	return statsDTO, err
}
