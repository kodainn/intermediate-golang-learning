package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kodainn/repository/controllers"
	"github.com/kodainn/repository/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}