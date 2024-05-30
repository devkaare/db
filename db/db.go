package db

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Username string
	Email    string
	Id       int
}

func WriteToFile(user User) {
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal(err)
		// WARNING: Program will crash here if the json file provided on line 8, doesn't contain square brackets already. Check the json file and add [] if it its empty
	}

	users = append(users, user)

	newData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("users.json", newData, 0644); err != nil {
		log.Fatal(err)
	}
}

func ReadFromFile() []User {
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	var user []User
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatal(err)
	}

	return user
}

func GetUserById(id int) User {
	users := ReadFromFile()

	var foundUser User
	for _, user := range users {
		if user.Id == id {
			foundUser = user
		}
	}
	return foundUser
}
