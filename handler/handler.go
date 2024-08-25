package handler

import (
	"blogging-platform-api/domain"
	"blogging-platform-api/initialize"
	"blogging-platform-api/repository"
	"blogging-platform-api/repository/cache"
	"blogging-platform-api/service"
	"blogging-platform-api/utils/pagination"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.IService
}

func NewHandler(app *initialize.Application) *Handler {
	return &Handler{
		service: service.NewService(
			cache.NewCache(app.Redis, app.Config.Redis.GetDefaultTTL()),
			repository.NewRepository(app.Database),
		),
	}
}

func (h *Handler) Get(c *gin.Context) {
	data, err := h.service.Get()
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOK(c, data, "Success")
}

func (h *Handler) Post(c *gin.Context) {
	var payload domain.StructName

	if err := payload.Validate(c); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	data, err := h.service.Get()
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseCreated(c, data, "Success")
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
