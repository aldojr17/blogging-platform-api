package repository

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/utils/pagination"

	"gorm.io/gorm"
)

type (
	IRepository interface {
		GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error)
		Create(payload domain.StructName) error
		GetByUUID(uuid string) (*domain.StructName, error)
	}

	Repository struct {
		db *gorm.DB
	}
)

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error) {
	var count int64
	var err error

	arguments := []interface{}{
		pageable.SearchParams()["<search query>"],
		pageable.FilterParams()["<filter by query>"],
	}

	initArgumentsIndex := len(arguments)

	chainMethod := r.db.Model(domain.StructName{})

	if arguments[0].(string) != "%%" {
		chainMethod = chainMethod.Where("name ILIKE ?", arguments[0].(string))
	}

	// limit pagination
	limit := int(1000 * ((pageable.GetPage() / (1000 / pageable.GetLimit())) + 1))

	err = r.db.Select("count(*)").Table("(?) as <table name>", chainMethod.Session(&gorm.Session{}).Select("*").Limit(limit)).Scan(&count).Error

	if err != nil {
		return pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).Pageable([]interface{}{}), err
	}

	if count == 0 {
		return pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), 0).Pageable([]interface{}{}), err
	}

	paginator := pagination.NewPaginator(pageable.GetPage(), pageable.GetLimit(), int(count))
	arguments = append(arguments, pageable.SortByFunc(), paginator.PerPageNums, paginator.Offset())

	var products []*domain.StructName

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

func (r *Repository) Create(payload domain.StructName) error {
	return r.db.Create(&payload).Error
}

func (r *Repository) GetByUUID(uuid string) (*domain.StructName, error) {
	data := new(domain.StructName)
	if err := r.db.Where("id", uuid).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
