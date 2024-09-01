package handler

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/initialize"
	"blogging-platform-api/repository"
	"blogging-platform-api/repository/cache"
	"blogging-platform-api/service"
	"blogging-platform-api/utils/pagination"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.IPostService
}

func NewHandler(app *initialize.Application) *Handler {
	return &Handler{
		service: service.NewPostService(
			cache.NewCache(app.Redis, app.Config.Redis.GetDefaultTTL()),
			repository.NewPostRepository(app.Database),
			repository.NewTagRepository(app.Database),
			repository.NewPostTagRepository(app.Database),
		),
	}
}

func (h *Handler) CreatePost(c *gin.Context) {
	var payload domain.CreatePostRequest

	if err := payload.Validate(c); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	data, err := h.service.CreatePost(payload)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseCreated(c, data, "New post created.")
}

func (h *Handler) GetDetailPost(c *gin.Context) {
	_id := c.Param("id")

	postId, err := strconv.Atoi(_id)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	}

	data, err := h.service.GetDetailPost(postId)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOK(c, data, "Successfully get post detail")
}

func (h *Handler) GetWithPagination(c *gin.Context) {
	data, err := h.service.GetWithPagination(newPageableRequest(c.Request))
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	PaginationSuccessResponse(c, data, "Success")
}

func newPageableRequest(r *http.Request) *domain.PageableRequest {
	p := &domain.PageableRequest{}
	p.Page = pagination.PageFromQueryParam(r)
	p.Limit = pagination.LimitFromQueryParam(r)
	p.SortBy = pagination.SortValueFromQueryParam(r)

	if p.SortBy == "" {
		p.SortBy = "<default sort by>"
	}

	p.Sort = pagination.SortDirectionFromQueryParam(r)
	p.Search = map[string]interface{}{}
	p.Filters = map[string]interface{}{}

	p.Search["<search query>"] = queryLikeParamOrNull(r, "<search query>")
	p.Filters["<filter by query>"] = queryParamOrNull(r, "<filter by query>")

	return p
}
