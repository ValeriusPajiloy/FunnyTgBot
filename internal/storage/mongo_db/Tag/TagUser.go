package tag

import (
	"context"
	"log"
	tag "tgbot/internal/mapping/Tag"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TagUser struct {
	db *mongo.Client
}

func NewTagUser(db *mongo.Client) *TagUser {
	return &TagUser{
		db: db,
	}
}

func (t *TagUser) Create(ctx context.Context, tagStr string, nameGroup string, chatID int) {
	collection := t.db.Database("Tag").Collection("User")
	filter := bson.D{{Key: "tag", Value: tagStr}, {Key: "nameGroup", Value: nameGroup}, {Key: "chatID", Value: chatID}}
	found := collection.FindOne(ctx, filter, options.FindOne())
	TagUser := tag.TagUser{}
	err := found.Decode(&TagUser)
	if err != nil && err.Error() == "mongo: no documents in result" {
		result, err := collection.InsertOne(ctx, bson.D{{Key: "tag", Value: tagStr}, {Key: "nameGroup", Value: nameGroup}, {Key: "chatID", Value: chatID}})
		if err != nil {
			log.Printf("%s", err)
		}
		log.Println(result)
	}
}
func (t *TagUser) GetAllForGroup(ctx context.Context, nameGroup string, chatID int) ([]string, error) {
	result := []string{}

	filter := bson.D{{Key: "chatID", Value: chatID}, {Key: "nameGroup", Value: nameGroup}}
	collection := t.db.Database("Tag").Collection("User")

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var TagUser tag.TagUser
		err := cur.Decode(&TagUser)
		if err != nil {
			return nil, err
		}
		result = append(result, TagUser.Tag)
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
