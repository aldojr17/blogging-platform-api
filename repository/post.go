package repository

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/domain/dto"
	"blogging-platform-api/utils/pagination"

	"gorm.io/gorm"
)

type (
	IPostRepository interface {
		Create(payload *dto.Post) error
		GetByID(id int) (*domain.GetDetailPostResponse, error)
		GetAllPost(pageable pagination.Pageable) (*pagination.Page, error)
	}

	PostRepository struct {
		db *gorm.DB
	}
)

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(payload *dto.Post) error {
	return r.db.Create(&payload).Error
}

func (r *PostRepository) GetByID(id int) (*domain.GetDetailPostResponse, error) {
	post := new(domain.GetDetailPostResponse)

	if err := r.db.Where("id", id).First(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) GetAllPost(pageable pagination.Pageable) (*pagination.Page, error) {
	var count int64
	var err error

	arguments := []interface{}{
		pageable.SearchParams()[domain.SEARCH_TERM],
	}

	initArgumentsIndex := len(arguments)

	chainMethod := r.db.Model(domain.GetDetailPostResponse{})

	if arguments[0] != nil {
		term := arguments[0].(string)
		chainMethod = chainMethod.Where("title LIKE ?", term).Or("content LIKE ?", term).Or("category LIKE ?", term)
	}

	// limit pagination
	limit := int(1000 * ((pageable.GetPage() / (1000 / pageable.GetLimit())) + 1))

	err = r.db.Select("count(*)").Table("(?) as posts", chainMethod.Session(&gorm.Session{}).Select("*").Limit(limit)).Scan(&count).Error

	if err != nil {
		return pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).Pageable([]interface{}{}), err
	}

	if count == 0 {
		return pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).Pageable([]interface{}{}), err
	}

	paginator := pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), int(count))
	arguments = append(arguments, pageable.SortByFunc(), paginator.PerPageNums, paginator.Offset())

	var products []*domain.GetDetailPostResponse

	err = chainMethod.
		Order(arguments[initArgumentsIndex].(string)).
		Limit(arguments[initArgumentsIndex+1].(int)).
		Offset(arguments[initArgumentsIndex+2].(int)).
		Find(&products).Error

	if err != nil {
		return pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).
			Pageable([]interface{}{}), err
	}

	return paginator.Pageable(products), nil
}
