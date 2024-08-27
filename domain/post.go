package domain

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreatePostResponse struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

type CreatePostRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
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
