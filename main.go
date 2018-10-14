package main

import (
	"log"

	"github.com/dyxj/go-mod-sample/boiler/models"
	"github.com/dyxj/go-mod-sample/noboiler"
	_ "github.com/lib/pq"
)

func main() {
	DB := noboiler.InitDB()
	defer DB.Close()

	userid, err := noboiler.InsertUser("user1")
	if err != nil {
		log.Fatal(err)
	}

	u, err := noboiler.GetUser(userid)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u)

	users, err := models.Users().All(nil, DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
	for _, u := range users {
		log.Println(u)
	}
}
