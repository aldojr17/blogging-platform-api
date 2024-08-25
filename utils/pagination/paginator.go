package pagination

import (
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Pageable interface {
	SearchParams() map[string]interface{}
	FilterParams() map[string]interface{}
	GetPage() int
	GetLimit() int
	SortByFunc() string
}

type Page struct {
	Data        interface{} `json:"data"`
	CurrentPage int         `json:"current_page"`
	Total       int         `json:"total"`
	TotalPage   int         `json:"total_page"`
	Limit       int         `json:"limit"`
}

type Paginator struct {
	PerPageNums int
	MaxPages    int

	nums     int
	pageNums int
	page     int
}

func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}

	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}

	p.pageNums = int(pageNums)
	return p.pageNums
}

func (p *Paginator) Page() int {
	if p.page > p.PageNums() {
		p.page = p.PageNums()
	}

	if p.page <= 0 {
		p.page = 1
	}

	return p.page
}

func (p *Paginator) Offset() int {
	return (p.Page() - 1) * p.PerPageNums
}

func (p *Paginator) Pageable(data interface{}) *Page {
	return &Page{
		CurrentPage: p.page,
		Total:       p.nums,
		TotalPage:   p.pageNums,
		Data:        data,
		Limit:       p.PerPageNums,
	}
}

func NewPaginator(currentPage, limit, nums int) *Paginator {
	p := Paginator{}

	p.page = currentPage
	if limit <= 0 {
		limit = 10
	}

	p.PerPageNums = limit
	p.nums = nums

	return &p
}

func PageFromQueryParam(r *http.Request) int {
	if page, e := strconv.Atoi(r.FormValue("page")); e != nil {
		return 0
	} else {
		if page <= 0 {
			page = 0
		}
		return page
	}
}

func LimitFromQueryParam(r *http.Request) int {
	if limit, e := strconv.Atoi(r.FormValue("limit")); e != nil {
		return 10
	} else {
		if limit <= 0 {
			limit = 10
		}
		return limit
	}
}

func SortValueFromQueryParam(r *http.Request) string {
	return r.FormValue("sortBy")
}

func SortDirectionFromQueryParam(r *http.Request) string {
	if sort := r.FormValue("sort"); strings.ToLower(sort) == "asc" {
		return "asc"
	} else {
		return "desc"
	}
}
