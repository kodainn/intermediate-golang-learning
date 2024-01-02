package repositories

import (
	"database/sql"

	"github.com/kodainn/repository/models"
)

const (
	articleNumPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPage * (page - 1), articleNumPage)
	if err != nil {
		return []models.Article{}, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return []models.Article{}, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select * from articles where article_id = ?
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}
	var article models.Article
	var CreatedTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &CreatedTime)
	if err != nil {
		return models.Article{}, err
	}

	if CreatedTime.Valid {
		article.CreatedAt = CreatedTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, articleID)
	if err = row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, niceNum + 1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}