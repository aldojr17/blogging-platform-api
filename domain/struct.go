package domain

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type StructName struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Query  string `json:"query" binding:"required"`
}

// Change table name for gorm usages (gorm default will use struct name)
func (d *StructName) TableName() string {
	return "api"
}

// Validate request struct
func (v *StructName) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}

	if len(v.Query) > 1000 {
		return fmt.Errorf("invalid query length (maximum 1000)")
	}

	if v.Width < 0 {
		return fmt.Errorf("invalid width")
	}

	if v.Height < 0 {
		return fmt.Errorf("invalid height")
	}

	if v.Width == 0 {
		v.Width = 1024
	}

	if v.Height == 0 {
		v.Height = 1024
	}

	return nil
}
