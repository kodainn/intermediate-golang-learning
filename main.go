// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/kodainn/repository/api"
// )

// var (
// 	dbUser     = "docker"
// 	dbPassword = "pass"
// 	dbDatabase = "sampledb"
// 	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// )

// func main() {
// 	db, err := sql.Open("mysql", dbConn)
// 	if err != nil {
// 		log.Println("fail to connect DB")
// 		return
// 	}

// 	r := api.NewRouter(db)

// 	log.Println("server start at port 8081")
// 	log.Fatal(http.ListenAndServe(":8081", r))
// }

package main

import (
	"fmt"
	"strings"
)

func doubleInt(src int, intCh chan<- int) {
	result := src * 2
	intCh <- result
}

func doubleString(src string, strCh chan<- string) {
	result := strings.Repeat(src, 2)
	strCh <- result
}

func main() {
	ch1, ch2 := make(chan int), make(chan string)
	defer close(ch1)
	defer close(ch2)

	go doubleInt(1, ch1)
	go doubleString("hello", ch2)

	for i := 0; i < 2; i++ {
		select {
		case numResult := <-ch1:
			fmt.Println(numResult)
		case strResult := <-ch2:
			fmt.Println(strResult)
		}
	}
}
