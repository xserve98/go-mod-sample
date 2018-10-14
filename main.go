package main

import (
	"log"

	"github.com/dyxj/go-mod-sample/noboiler"
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

}
