package service

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/domain/dto"
	"blogging-platform-api/repository"
	"blogging-platform-api/repository/cache"
	"blogging-platform-api/utils"
	"blogging-platform-api/utils/pagination"
)

type (
	IPostService interface {
		CreatePost(payload domain.CreatePostRequest) (*domain.CreatePostResponse, error)

		GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error)
		Get() (*dto.Post, error)
	}

	PostService struct {
		cache             cache.ICache
		postRepository    repository.IPostRepository
		tagRepository     repository.ITagRepository
		postTagRepository repository.IPostTagRepository
	}
)

func NewPostService(
	cache cache.ICache,
	postRepository repository.IPostRepository,
	tagRepository repository.ITagRepository,
	postTagRepository repository.IPostTagRepository,
) IPostService {
	return &PostService{
		cache:             cache,
		postRepository:    postRepository,
		tagRepository:     tagRepository,
		postTagRepository: postTagRepository,
	}
}

func (s *PostService) CreatePost(payload domain.CreatePostRequest) (*domain.CreatePostResponse, error) {
	var tagList []int

	for _, tag := range payload.Tags {
		result, err := s.tagRepository.Create(tag)
		if err != nil {
			return nil, err
		}

		tagList = append(tagList, result.ID)
	}

	post := dto.Post{
		Title:      payload.Title,
		Content:    payload.Content,
		Category:   payload.Category,
		CreateTime: utils.GenerateCurrentTimestamp(),
		UpdateTime: utils.GenerateCurrentTimestamp(),
	}

	err := s.postRepository.Create(&post)
	if err != nil {
		return nil, err
	}

	for _, tagID := range tagList {
		postTag := dto.PostTag{
			PostID:     post.ID,
			TagID:      tagID,
			CreateTime: utils.GenerateCurrentTimestamp(),
		}

		err := s.postTagRepository.Create(postTag)
		if err != nil {
			return nil, err
		}
	}

	response := domain.CreatePostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		Tags:      payload.Tags,
		CreatedAt: utils.ConvertTimestampToFormattedDate(post.CreateTime),
		UpdatedAt: utils.ConvertTimestampToFormattedDate(post.UpdateTime),
	}

	return &response, nil
}

func (s *PostService) GetWithPagination(pageable pagination.Pageable) (*pagination.Page, error) {
	resp, err := s.postRepository.GetWithPagination(pageable)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PostService) Get() (*dto.Post, error) {
	data, _ := s.cache.Get()
	if data != nil {
		return data, nil
	}

	resp, err := s.postRepository.GetByUUID("uuid")
	if err != nil {
		return nil, err
	}

	if err = s.cache.Set(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
