package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kodainn/repository/api"
)

var (
	dbUser     = "docker"
	dbPassword = "pass"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// package main

// import (
// 	"context"
// 	"fmt"
// )


// func MyFunc1(ctx context.Context) {
// 	fmt.Println("MyFunc1 start")


// 	fmt.Println("MyFunc1 finish")
// }

// func MyFunc2(ctx context.Context) {
// 	fmt.Println("MyFunc2 start")


// 	fmt.Println("MyFunc2 finish")
// }

// func main() {
// 	var ctx context.Context

// 	go MyFunc1(ctx)
// 	go MyFunc2(ctx)

// }
