package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kodainn/repository/models"
	"github.com/kodainn/repository/repositories"
	"github.com/kodainn/repository/repositories/testdata"
)

func TestSelectArticleDetail(t *testing.T) {

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: testdata.ArticleTestData[0],
		},
		{
			testTitle: "subtest2",
			expected: testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testtest",
		UserName: "saki",
	}

	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != article.Title {
		t.Errorf("new article id is expected %s but got %s\n", newArticle.Title, article.Title)
	}
	if newArticle.Contents != article.Contents {
		t.Errorf("new article id is expected %s but got %s\n", newArticle.Contents, article.Contents)
	}
	if newArticle.UserName != article.UserName {
		t.Errorf("new article id is expected %s but got %s\n", newArticle.UserName, article.UserName)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get after data")
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}

	t.Cleanup(func() {
		const sqlStr = `
			update articles set nice = nice - 1 where article_id = ?
		`

		testDB.Exec(sqlStr, articleID)
	})
}
