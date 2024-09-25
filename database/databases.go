package databases

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	jsoniter "github.com/json-iterator/go"

	_ "github.com/go-sql-driver/mysql"

	"tryGo_two/database/sqlc"
)

var (
	dbse    *sql.DB
	queries *sqlc.Queries
	once    sync.Once
	json    jsoniter.API
)

func initDB() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	var err error
	dbse, err = sql.Open("mysql", os.Getenv("DB_URL"))
	dbse.SetMaxOpenConns(1000)
	if err != nil {
		panic(err)
	}
	queries = sqlc.New(dbse)
	fmt.Println("solo 1??")

}

func GetUsers() []byte {
	once.Do(initDB)
	ctx := context.Background()

	urs, err := queries.ListUsers(ctx)
	if err != nil {
		fmt.Println("Error", err)
	}
	jsonResponse, err := json.Marshal(&urs)
	if err != nil {
		fmt.Println("jsonResponse error convert")
	}

	return jsonResponse
}
