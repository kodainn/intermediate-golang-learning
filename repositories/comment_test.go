package repositories_test

import (
	"testing"

	"github.com/kodainn/repository/models"
	"github.com/kodainn/repository/repositories"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment articleID %d get ID %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.ArticleID != comment.ArticleID {
		t.Errorf("new comment id is expected %d but got %d\n", newComment.ArticleID, comment.ArticleID)
	}
	if newComment.Message != comment.Message {
		t.Errorf("new comment id is expected %s but got %s\n", newComment.Message, comment.Message)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}
