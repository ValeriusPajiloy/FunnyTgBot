package mongodb

import (
	"context"
	tag "tgbot/internal/storage/mongo_db/Tag"

	"go.mongodb.org/mongo-driver/mongo"
)

type TagGroup interface {
	Create(ctx context.Context, name string, chatID int)
	GetAll(ctx context.Context, chatID int) ([]string, error)
}

type TagUser interface {
	Create(ctx context.Context, tag string, nameGroup string, chatID int)
	GetAllForGroup(ctx context.Context, nameGroup string, chatID int) ([]string, error)
}

type Repository struct {
	TagGroup
	TagUser
}

func NewRepository(mongo *mongo.Client) *Repository {
	return &Repository{
		TagGroup: tag.NewTagGroup(mongo),
		TagUser:  tag.NewTagUser(mongo),
	}
}
