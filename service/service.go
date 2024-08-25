package service

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/repository"
	"blogging-platform-api/repository/cache"
	"blogging-platform-api/utils/pagination"
)

type (
	IService interface {
		GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error)
		Get() (*domain.StructName, error)
	}

	Service struct {
		cache      cache.ICache
		repository repository.IRepository
	}
)

func NewService(
	cache cache.ICache,
	repository repository.IRepository,
) IService {
	return &Service{
		cache:      cache,
		repository: repository,
	}
}

func (s *Service) GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error) {
	resp, err := s.repository.GetWithPagination(pageable)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) Get() (*domain.StructName, error) {
	data, _ := s.cache.Get()
	if data != nil {
		return data, nil
	}

	resp, err := s.repository.GetByUUID("uuid")
	if err != nil {
		return nil, err
	}

	if err = s.cache.Set(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
