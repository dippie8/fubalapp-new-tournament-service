package mongodb

import (
	"context"
	"github.com/dippie8/fubalapp-new-tournament/pkg/initializing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)


type Storage struct {
	Uri string `yaml:"Uri"`
	db  *mongo.Client
	ctx context.Context
}

func NewDB() (s *Storage, err error){
	yamlFile, err := ioutil.ReadFile("parameters.yml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &s)

	client, err := mongo.NewClient(options.Client().ApplyURI(s.Uri))
	if err != nil {
		return nil, err
	}
	s.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(s.ctx)
	if err != nil {
		return nil, err
	}
	s.db = client


	return s, nil
}

func (s Storage) Disconnect() {
	_ = s.db.Disconnect(s.ctx)
}

func (s Storage) GetStandings() ([]initializing.Standing, error) {
	collection := s.db.Database("qlsr").Collection("standings")
	filter := bson.D{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"elo", -1}})
	cursor, err := collection.Find(context.TODO(), filter, findOptions)

	var standings []initializing.Standing

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var s *Standing

		err = cursor.Decode(&s)
		if err != nil {
			return nil, err
		}

		initializingStanding := initializing.Standing{
			Id:     s.Username,
			Win:    s.Win,
			Played: s.Played,
		}

		standings = append(standings, initializingStanding)
	}
	return standings, nil
}

func (s Storage) AddPrize(player string, medal initializing.Medal) error {
	collection := s.db.Database("qlsr").Collection("players")

	var key string
	switch medal {
	case initializing.Gold:
		key = "goldmedals"
	case initializing.Silver:
		key = "silvermedals"
	case initializing.Bronze:
		key = "bronzemedals"
	}

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": player},
		bson.D{
			{"$inc", bson.D{{key, 1}}},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return err
	}
	return nil
}

func (s Storage) ResetStandings() error {
	collection := s.db.Database("qlsr").Collection("standings")

	_, err := collection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{
			{"$set", bson.D{{"played", 0}}},
			{"$set", bson.D{{"win", 0}}},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return err
	}
	return nil
}

