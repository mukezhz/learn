package author

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/models"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/responses"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/utils"
)

type Controller struct {
	service *Service
	logger  framework.Logger
}

func NewController(
	service *Service,
	logger framework.Logger,
) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

func (c *Controller) HandleFindAuthors(ctx *gin.Context) {
	message, _ := c.service.FindAuthors(ctx)
	res := utils.Map(message, func(m models.Author) AuthorResponse {
		return AuthorResponse{
			ID:        m.ID,
			Name:      m.Name,
			Bio:       m.Bio,
			CreatedAt: m.CreatedAt.String(),
			UpdatedAt: m.UpdatedAt.String(),
		}
	})
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (c *Controller) HandleAddAuthor(ctx *gin.Context) {
	b := CreateAuthorParams{}
	if err := ctx.ShouldBindJSON(&b); err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.service.AddAuthor(ctx, b.ToModel()); err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.SuccessJSON(ctx, http.StatusCreated, "Author created successfully")
}

func (c *Controller) HandleFilterAuthors(ctx *gin.Context) {
	b := FilterAuthorsParams{}
	if err := ctx.ShouldBindQuery(&b); err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	c.logger.Infof("Filtering authors params %#v", b)

	message, _ := c.service.FilterAuthors(ctx, framework.PaginatedRequest{
		Page:  b.Page,
		Limit: b.Limit,
	}, b.ToModel())
	res := utils.Map(message, func(m models.Author) AuthorResponse {
		return AuthorResponse{
			ID:        m.ID,
			Name:      m.Name,
			Bio:       m.Bio,
			CreatedAt: m.CreatedAt.String(),
			UpdatedAt: m.UpdatedAt.String(),
		}
	})
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (c *Controller) HandleGetAuthorByID(ctx *gin.Context) {
	rawParam := ctx.Param("id")
	id, err := strconv.Atoi(rawParam)
	if err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}
	message, _ := c.service.GetAuthorByID(ctx, id)
	res := AuthorResponse{
		ID:        message.ID,
		Name:      message.Name,
		Bio:       message.Bio,
		CreatedAt: message.CreatedAt.String(),
		UpdatedAt: message.UpdatedAt.String(),
		Books: utils.Map(message.Books, func(m models.Book) BookResponse {
			b := BookResponse{}
			return b.FromModel(m)
		}),
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (c *Controller) HandleAddAuthorWithBooks(ctx *gin.Context) {
	b := CreateAuthorWithBooksParams{}
	if err := ctx.ShouldBindJSON(&b); err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	author := b.ToModel()
	author.Books = utils.Map(b.Books, func(m CreateBookParams) models.Book {
		return m.ToModel()
	})
	if err := c.service.AddAuthorWithBooks(ctx, author); err != nil {
		responses.ErrorJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	responses.SuccessJSON(ctx, http.StatusCreated, "Author with books created successfully")
}
