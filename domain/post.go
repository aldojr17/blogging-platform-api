package domain

import (
	"blogging-platform-api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePostRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
}

type CreatePostResponse struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

type GetDetailPostResponse struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags" gorm:"-"`
	CreatedAt string   `json:"createdAt" gorm:"column:create_time"`
	UpdatedAt string   `json:"updatedAt" gorm:"column:update_time"`
}

func (d *GetDetailPostResponse) TableName() string {
	return TABLE_POST_TAB
}

func (h *GetDetailPostResponse) AfterFind(tx *gorm.DB) (err error) {
	if err := tx.
		Table(fmt.Sprintf("%s pt", TABLE_POST_TAG_TAB)).
		Joins(fmt.Sprintf("join %s t on pt.tag_id = t.id", TABLE_TAG_TAB)).
		Select("t.name").
		Where("pt.post_id", h.ID).Find(&h.Tags).Error; err != nil {
		return err
	}

	h.CreatedAt = utils.ConvertTimestampToFormattedDate(int64(utils.ConvertToInteger(h.CreatedAt)))
	h.UpdatedAt = utils.ConvertTimestampToFormattedDate(int64(utils.ConvertToInteger(h.UpdatedAt)))

	return
}

func (v *CreatePostRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}

	if len(v.Title) > 128 {
		return fmt.Errorf("invalid title length (maximum 128)")
	}

	if len(v.Category) > 32 {
		return fmt.Errorf("invalid category length (maximum 32)")
	}

	if len(v.Tags) == 0 {
		return fmt.Errorf("invalid tags length (at least 1 tag)")
	}

	return nil
}
