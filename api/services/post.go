package services

import (
	"context"
	"gorm.io/gorm"
	"rowing-market-api/api/models"
)

type PostService struct {
	Ctx         context.Context
	PostManager models.PostManager
}

func GetPostService(ctx context.Context, db *gorm.DB) PostService {
	postManager := models.GetPostManager(ctx, db)
	return PostService{
		ctx,
		postManager,
	}
}

func (ps PostService) SavePost(post models.Post) (int64, error) {

	result, err := ps.PostManager.SavePost(post)
	if err != nil {
		return 0, err
	}

	return result, nil
}
