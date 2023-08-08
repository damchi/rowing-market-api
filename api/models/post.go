package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"rowing-market-api/pkg/logger"
	mysqlgorm "rowing-market-api/pkg/mysql-gorm"
)

type Post struct {
	gorm.Model
	UserId      uint64 `json:"user_id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Price       uint64 `json:"price"`
}

func (Post) TableName() string {
	return "post"
}

type PostManager struct {
	db      *gorm.DB
	Context context.Context
}

func GetPostManager(ctx context.Context, db *gorm.DB) PostManager {
	if db != nil {
		return PostManager{db, ctx}
	}
	return PostManager{mysqlgorm.GetConnection(), ctx}
}

func (m PostManager) SavePost(post Post) (int64, error) {

	result := m.db.Create(&post)

	if result.Error != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("Save post : %v", result.Error))
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
