package socialrating

import (
	"context"
	"log"
	socialrating "tgbot/internal/mapping/SocialRating"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SocialRating struct {
	db *mongo.Client
}

func NewSocialRating(db *mongo.Client) *SocialRating {
	return &SocialRating{
		db: db,
	}
}

func (s *SocialRating) ScanMessage(ctx context.Context, textMessage string, userName string) {
	collection := s.db.Database("Rating").Collection("User")
	filter := bson.D{{Key: "name", Value: userName}}
	found := collection.FindOne(ctx, filter, options.FindOne())
	user := socialrating.User{}
	err := found.Decode(&user)

	if err != nil && err.Error() == "mongo: no documents in result" {
		result, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: userName}, {Key: "rating", Value: 1000}})
		if err != nil {
			log.Printf("%s", err)
		}
		log.Println(result)
	}
}
