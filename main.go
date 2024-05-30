package main

import (
	"learn-go-json/db"
	"log"
	"math"
	"math/rand"
)

func main() {
	db.FillCacheForDB()
	defer db.SaveCacheForDB()

	//log.SetFlags(log.Lshortfile)

	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(math.MaxInt32)
		foo := db.User{"foobar", "foobar@foo.bar", randomNum}
		db.WriteToDB(foo)
        //log.Printf("Wrote number %v, to database\n", i)
	}

    randUserId := 36713240

	log.Println("Checking for user with Id:", randUserId)
	user := db.GetUserById(randUserId)
	if user.Username != "" {
		log.Printf("Found user: Username: %v, Email: %v, Id: %v\n", user.Username, user.Email, user.Id)
	} else {
		log.Println("No user was found with Id:", randUserId)
	}

    log.Println("Trying to delete user with Id:", randUserId)
	if user.Username != "" {
        log.Println("Deleted user with Id:", randUserId)
        db.DeleteUserById(randUserId)
	} else {
		log.Println("No user was found with Id:", randUserId)
	}
}
