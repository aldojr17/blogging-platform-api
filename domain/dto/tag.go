package dto

import "blogging-platform-api/domain"

type Tag struct {
	ID         int
	Name       string
	CreateTime int64 `gorm:"default:0"`
}

func (d *Tag) TableName() string {
	return domain.TABLE_TAG_TAB
}
