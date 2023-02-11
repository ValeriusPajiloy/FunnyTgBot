package socialrating

import (
	mongodb "tgbot/internal/storage/mongo_db"

	client "tgbot/internal/clients"
)

type SocialRating struct {
	tg   client.Client
	repo *mongodb.Repository
}

func NewSocialRating(tg client.Client, repo *mongodb.Repository) *SocialRating {
	return &SocialRating{
		tg:   tg,
		repo: repo,
	}
}
