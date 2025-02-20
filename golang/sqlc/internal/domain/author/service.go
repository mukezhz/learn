package author

import (
	"context"

	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/models"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
)

// Service handles the business logic of the module
type Service struct {
	repo   *Repository
	logger framework.Logger
}

// NewService creates a new instance of TestService
func NewService(
	repo *Repository,
	logger framework.Logger,
) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

// FindAuthors returns a greeting message
func (s *Service) FindAuthors(ctx context.Context) ([]models.Author, error) {
	return s.repo.FindAuthors(ctx)
}

func (s *Service) AddAuthor(ctx context.Context, author models.Author) error {
	return s.repo.AddAuthor(ctx, author)
}

func (s *Service) FilterAuthors(
	ctx context.Context,
	pagination framework.PaginatedRequest,
	author models.AuthorFilter,
) ([]models.Author, error) {
	s.logger.Infoln("Filtering authors", "author", author, " pagination", pagination)
	return s.repo.FilterAuthors(ctx, pagination, author)
}

func (s *Service) GetAuthorByID(ctx context.Context, id int) (models.Author, error) {
	return s.repo.GetAuthorByID(ctx, id)
}

func (s *Service) AddAuthorWithBooks(
	ctx context.Context,
	author models.Author,
) error {
	return s.repo.AddAuthorWithBooks(ctx, author)
}
