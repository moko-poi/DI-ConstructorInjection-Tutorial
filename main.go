package main

import (
	"Dependencie/bad-dependencie/service"
	"Dependencie/good-dependencie/repository"
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ur := repository.NewUserReposoty(db)
	us := service.NewUserService(ur)

	// us を使ってゴニョ
}
