package repository

import (
	"blogging-platform-api/domain/dto"
	"blogging-platform-api/utils"

	"gorm.io/gorm"
)

type (
	ITagRepository interface {
		Create(name string) (dto.Tag, error)
	}

	TagRepository struct {
		db *gorm.DB
	}
)

func NewTagRepository(db *gorm.DB) ITagRepository {
	return &TagRepository{
		db: db,
	}
}

func (r *TagRepository) Create(name string) (dto.Tag, error) {
	var tag dto.Tag

	if err := r.db.Attrs(dto.Tag{CreateTime: utils.GenerateCurrentTimestamp()}).FirstOrCreate(&tag, dto.Tag{Name: name}).Error; err != nil {
		return tag, err
	}

	return tag, nil
}
