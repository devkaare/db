package main

import (
	"log"
    "math"
    "math/rand"
	"learn-go-json/db"
)

func main() {
	//log.SetFlags(log.Lshortfile)
	for i := 0; i <= 10; i++ {
		randomNum := rand.Intn(math.MaxInt32)
		foo := db.User{"foo", "bar@gmail.com", randomNum}
		db.WriteToFile(foo)
	}

	randUserId := 1836542667
	log.Println("Checking Id:", randUserId)
	user := db.GetUserById(randUserId)
	if user.Username != "" {
		log.Printf("Found user:\nUsername: %v, Email: %v, Id: %v\n", user.Username, user.Email, user.Id)
	} else {
		log.Println("No user was found with Id:", randUserId)
	}
}
