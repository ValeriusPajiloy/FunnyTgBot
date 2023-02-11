package tag

import (
	"context"
	"log"
	tag "tgbot/internal/mapping/Tag"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagGroup struct {
	db *mongo.Client
}

func NewTagGroup(db *mongo.Client) *TagGroup {
	return &TagGroup{
		db: db,
	}
}

func (t *TagGroup) Create(ctx context.Context, name string, chatID int) {
	collection := t.db.Database("Tag").Collection("Group")
	result, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: name}, {Key: "chatID", Value: chatID}})
	if err != nil {
		log.Printf("%s", err)
	}
	log.Println(result)
}
func (t *TagGroup) GetAll(ctx context.Context, chatID int) ([]string, error) {
	result := []string{}

	filter := bson.D{{Key: "chatID", Value: chatID}}
	collection := t.db.Database("Tag").Collection("Group")

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var TagGroup tag.TagGroup
		err := cur.Decode(&TagGroup)
		if err != nil {
			return nil, err
		}
		result = append(result, TagGroup.Name)
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
