package services

import (
	"github.com/kodainn/repository/models"
	"github.com/kodainn/repository/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return []models.Article{}, err
	}

	return articleList, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	newArticle, err := repositories.SelectArticleDetail(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}
