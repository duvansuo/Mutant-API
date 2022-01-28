package data

import (
	"context"
	"fmt"
	"log"
	"time"

	models "github.com/duvansuo/Mutant-API/Models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type statsMongoData struct {
	database *mongo.Collection
}

func newStatsMongoData() *statsMongoData {
	fmt.Println("ingreso ")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Mutant:Mutant123@pruebamutant.v0njz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	fmt.Println("Continuo ")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
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
	ctx := context.Background()
	var statsDTO models.MutantStatsDTO
	err := repo.database.FindOne(ctx, bson.M{}).Decode(&statsDTO)
	if err != nil {
		return models.MutantStatsDTO{}, err
	}
	return statsDTO, err
}
