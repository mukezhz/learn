package author

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/mukezhz/learn/tree/main/golang/sqlc/internal/adapter/persistent/sqlc_gen"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/internal/domain/models"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/framework"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/infrastructure"
	"github.com/mukezhz/learn/tree/main/golang/sqlc/pkg/utils"
)

type Repository struct {
	db      *infrastructure.Database
	queries *repo.Queries
	logger  framework.Logger
}

func NewRepository(
	db *infrastructure.Database,
	logger framework.Logger,
) *Repository {
	return &Repository{
		db:      db,
		queries: repo.New(db.Pool),
		logger:  logger,
	}
}

func (r *Repository) FindAuthors(
	ctx context.Context,
) (authors []models.Author, err error) {
	dbAuthors, err := r.queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}

	authors = utils.Map(dbAuthors, func(dbAuthor repo.Author) models.Author {
		return models.Author{
			ID:        dbAuthor.ID,
			Name:      dbAuthor.Name,
			Bio:       dbAuthor.Bio.String,
			CreatedAt: dbAuthor.CreatedAt.Time,
			UpdatedAt: dbAuthor.UpdatedAt.Time,
		}
	})
	return
}

func (r *Repository) AddAuthor(
	ctx context.Context,
	author models.Author,
) error {
	_, err := r.queries.CreateAuthor(ctx, repo.CreateAuthorParams{
		Name: author.Name,
		Bio: pgtype.Text{
			String: author.Bio,
			Valid:  true,
		},
	})
	return err
}

func (r *Repository) FilterAuthors(
	ctx context.Context,
	pagination framework.PaginatedRequest,
	author models.AuthorFilter,
) (authors []models.Author, err error) {
	r.logger.Info("Filtering authors", " author name: ", author.Name, " bio", author.Bio)
	dbAuthors, err := r.queries.FilterAuthors(ctx, repo.FilterAuthorsParams{
		Limit:  int32(pagination.GetLimit()),
		Name:   author.Name,
		Offset: int32(pagination.GetOffset()),
		Bio:    author.Bio,
	})
	if err != nil {
		r.logger.Error("Error filtering authors", err)
		return []models.Author{}, err
	}

	authors = utils.Map(dbAuthors, func(dbAuthor repo.Author) models.Author {
		return models.Author{
			ID:        dbAuthor.ID,
			Name:      dbAuthor.Name,
			Bio:       dbAuthor.Bio.String,
			CreatedAt: dbAuthor.CreatedAt.Time,
			UpdatedAt: dbAuthor.UpdatedAt.Time,
		}
	})
	return
}

func (r *Repository) GetAuthorByID(
	ctx context.Context,
	id int,
) (author models.Author, err error) {
	dbAuthor, err := r.queries.GetAuthor(ctx, int64(id))
	if err != nil {
		return models.Author{}, err
	}

	author = models.Author{
		ID:        dbAuthor.ID,
		Name:      dbAuthor.Name,
		Bio:       dbAuthor.Bio.String,
		CreatedAt: dbAuthor.CreatedAt.Time,
		UpdatedAt: dbAuthor.UpdatedAt.Time,
	}
	dbBooks, err := r.queries.GetBooksByAuthorID(ctx, author.ID)
	if err != nil {
		author.Books = []models.Book{}
	}
	books := utils.Map(dbBooks, func(dbBook repo.Book) models.Book {
		return models.Book{
			ID:            dbBook.ID,
			Title:         dbBook.Title,
			AuthorID:      dbBook.AuthorID,
			PublishedDate: dbBook.PublishedDate.Time,
		}
	})
	r.logger.Info("Books", books)
	author.Books = books
	return
}

func (r *Repository) AddAuthorWithBooks(
	ctx context.Context,
	author models.Author,
) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	dbAuthor, err := r.queries.WithTx(tx).CreateAuthor(ctx, repo.CreateAuthorParams{
		Name: author.Name,
		Bio: pgtype.Text{
			String: author.Bio,
		},
	})
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	_, err = r.queries.WithTx(tx).CreateBooks(ctx,
		utils.Map(author.Books, func(book models.Book) repo.CreateBooksParams {
			return repo.CreateBooksParams{
				Title:    book.Title,
				AuthorID: dbAuthor.ID,
				PublishedDate: pgtype.Date{
					Time:  book.PublishedDate,
					Valid: true,
				},
			}
		}))

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
