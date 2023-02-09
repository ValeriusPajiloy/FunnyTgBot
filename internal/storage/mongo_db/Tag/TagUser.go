package tag

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TagUser struct {
	db *mongo.Client
}

func NewTagUser(db *mongo.Client) *TagUser {
	return &TagUser{
		db: db,
	}
}

func (t *TagUser) Create(ctx context.Context, tag string, nameGroup string, chatID int) {

}
func (t *TagUser) GetAllForGroup(ctx context.Context, nameGroup string, chatID int) ([]string, error) {

	results := []string{}

	return results, nil
}
