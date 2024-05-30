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

    randUserId := 970454965

	log.Println("Checking Id:", randUserId)
	user := db.GetUserById(randUserId)
	if user.Username != "" {
		log.Printf("Found user: Username: %v, Email: %v, Id: %v\n", user.Username, user.Email, user.Id)
	} else {
		log.Println("No user was found with Id:", randUserId)
	}
}
