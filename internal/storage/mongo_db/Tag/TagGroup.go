package tag

import (
	"context"

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

}
func (t *TagGroup) GetAll(ctx context.Context, chatID int) ([]string, error) {

	result := []string{}

	return result, nil
}
