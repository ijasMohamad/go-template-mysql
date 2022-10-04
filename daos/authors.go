package daos

import (
	"context"
	"fmt"
	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Author queries...

func FindAuthorById (authorId int, ctx context.Context) (*models.Author, error) {
     var contextExecutor = getContextExecutor(nil)
     author, err := models.FindAuthor(ctx, contextExecutor, authorId)
     return author, err
}

func FindAllAuthorsWithCount (ctx context.Context) (models.AuthorSlice, int, error) {
     var contextExecutor = getContextExecutor(nil)
     authors, err := models.Authors().All(ctx, contextExecutor)
     if err != nil {
          return models.AuthorSlice{}, 0, err
     }
     count, err := models.Authors().Count(ctx, contextExecutor)
     if err != nil {
          return models.AuthorSlice{}, 0, err
     }
     return authors, int(count), err
}

// Author mutations...

func CreateAuthor (author models.Author, ctx context.Context) (models.Author, error) {
     contextExecutor := getContextExecutor(nil)
     err := author.Insert(ctx, contextExecutor, boil.Infer())
     fmt.Println("Error in doas package: ", err)
     return author, err
}

func UpdateAuthor (author models.Author, ctx context.Context) (models.Author, error) {
     contextExecutor := getContextExecutor(nil)
     _, err := author.Update(ctx, contextExecutor, boil.Infer())
     return author, err
}

func DeleteAuthor (author models.Author, ctx context.Context) (int64, error) {
     var contextExecutor = getContextExecutor(nil)
     rowsAffected, err := author.Delete(ctx, contextExecutor)
     return rowsAffected, err
}