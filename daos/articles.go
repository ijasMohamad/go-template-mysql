package daos

import (
	"context"
	"database/sql"
	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Article Mutations....

func CreateArticle(article models.Article, ctx context.Context) (models.Article, error) {
	return CreateArticleTx(article, ctx, nil)
}

func CreateArticleTx(article models.Article, ctx context.Context, tx *sql.Tx) (models.Article, error) {
	contextExecutor := getContextExecutor(tx)
	err := article.Insert(ctx, contextExecutor, boil.Infer())
	return article, err
}

func UpdateArticle(article models.Article, ctx context.Context) (models.Article, error) {
	return UpdateArticleTx(article, ctx, nil)
}

func UpdateArticleTx(article models.Article, ctx context.Context, tx *sql.Tx) (models.Article, error) {
	contextExecutor := getContextExecutor(tx)
	_, err := article.Update(ctx, contextExecutor, boil.Infer())
	return article, err
}

func DeleteArticle(article models.Article, ctx context.Context) (int64, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := article.Delete(ctx, contextExecutor)
	return rowsAffected, err
}

// Article Queries....

func FindArticleById(articleID int, ctx context.Context) (*models.Article, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindArticle(ctx, contextExecutor, articleID)
}

func FindAllArticlesWithCount(ctx context.Context) (models.ArticleSlice, int, error) {
	contextExecutor := getContextExecutor(nil)
	articles, err := models.Articles().All(ctx, contextExecutor)
	if err != nil {
		return models.ArticleSlice{}, 0, err
	}
     count, err := models.Articles().Count(ctx, contextExecutor)
     if err != nil {
          return models.ArticleSlice{}, 0, err
     }
	return articles, int(count), nil
}
