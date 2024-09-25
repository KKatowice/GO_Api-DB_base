package main

import (
	//"fmt"
	//databases "tryGo_two/database"
	apis "tryGo_two/api"

	"github.com/joho/godotenv"
)

func main() {
	/* usr := databases.GetUsers()
	fmt.Println("db:", string(usr)) */
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}
	apis.ApisTest()
}
