package databases

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"tryGo_two/database/sqlc"

	"github.com/joho/godotenv"
)

func GetUsers() []byte {
	ctx := context.Background()
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
	dbse, err := sql.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(dbse)
	urs, err := queries.ListUsers(ctx)
	if err != nil {
		fmt.Println("Error", err)

	}
	jsonResponse, err := json.Marshal(urs)
	if err != nil {
		fmt.Println("jsonResponse error convert")

	}

	return jsonResponse
}
