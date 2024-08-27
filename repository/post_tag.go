package repository

import (
	"blogging-platform-api/domain/dto"

	"gorm.io/gorm"
)

type (
	IPostTagRepository interface {
		Create(payload dto.PostTag) error
	}

	PostTagRepository struct {
		db *gorm.DB
	}
)

func NewPostTagRepository(db *gorm.DB) IPostTagRepository {
	return &PostTagRepository{
		db: db,
	}
}

func (r *PostTagRepository) Create(payload dto.PostTag) error {
	return r.db.Create(&payload).Error
}
