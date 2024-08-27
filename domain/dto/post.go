package dto

import "blogging-platform-api/domain"

type Post struct {
	ID         int
	Title      string
	Content    string
	Category   string
	CreateTime int64 `gorm:"default:0"`
	UpdateTime int64 `gorm:"default:0"`
}

func (d *Post) TableName() string {
	return domain.TABLE_POST_TAB
}
