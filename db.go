package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type User struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	ID       int    `json:"ID"`
}

var userFile = "users.json"
var userFileCache []User

func DoesFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, err
}

func LoadUsersFromJSON() {
	result, _ := DoesFileExist(userFile)
	if !result {
		os.Create(userFile)
		os.WriteFile(userFile, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(userFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &userFileCache); err != nil {
		log.Fatal(err)
	}

}

func SaveUsersToJSON() {
	newData, err := json.Marshal(userFileCache)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(userFile, newData, 0644); err != nil {
		log.Fatal(err)
	}
}

func InsertUser(user User) {
	userFileCache = append(userFileCache, user)
}

func RetrieveAllUsers() []User {
	return userFileCache
}

func RetrieveUserByID(id int) User {
	var foundUser User
	for _, user := range userFileCache {
		if user.Id == id {
			foundUser = user
            break
		}
	}
	return foundUser
}

func RemoveUserByID(id int) {
	var userIndex int
	for i, user := range userFileCache {
		if user.Id == id {
			userIndex = i
            break
		}
	}

    if userIndex > 0 {
        userFileCache = append(userFileCache[:userIndex], userFileCache[userIndex+1:]...)
    }
}
