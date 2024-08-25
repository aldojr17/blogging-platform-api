package domain

import (
	"fmt"
	"strings"
)

type PageableRequest struct {
	Filters map[string]interface{}
	Search  map[string]interface{}
	Limit   int
	Page    int
	SortBy  string
	Sort    string
}

func (p *PageableRequest) SearchParams() map[string]interface{} {
	return p.Search
}
func (p *PageableRequest) FilterParams() map[string]interface{} {
	return p.Filters
}
func (p *PageableRequest) GetPage() int {
	return p.Page
}
func (p *PageableRequest) GetLimit() int {
	return p.Limit
}
func (p *PageableRequest) SortByFunc() string {
	if p.SortBy == "" {
		return "name ASC"
	}

	if p.Sort == "asc" {
		return fmt.Sprintf("%s %s", strings.ToLower(p.SortBy), "ASC")
	}

	return fmt.Sprintf("%s %s", strings.ToLower(p.SortBy), "DESC")
}
