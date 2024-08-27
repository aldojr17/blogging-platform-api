package dto

import "blogging-platform-api/domain"

type PostTag struct {
	ID         int
	PostID     int
	TagID      int
	CreateTime int64 `gorm:"default:0"`
}

func (d *PostTag) TableName() string {
	return domain.TABLE_POST_TAG_TAB
}
