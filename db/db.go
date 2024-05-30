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
	Id       int    `json:"Id"`
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

func FillCacheForDB() {
	result, _ := DoesFileExist(userFile)
	if !result {
		os.Create(userFile)
		os.WriteFile(userFile, []byte("[]"), 0644)
	}

	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &userFileCache); err != nil {
		log.Fatal(err)
	}

}

func SaveCacheForDB() {
	newData, err := json.Marshal(userFileCache)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(userFile, newData, 0644); err != nil {
		log.Fatal(err)
	}
}

func WriteToDB(user User) {
	userFileCache = append(userFileCache, user)
}

func ReadFromFile() []User {
	return userFileCache
}

func GetUserById(id int) User {
	var foundUser User
	for _, user := range userFileCache {
		if user.Id == id {
			foundUser = user
            break
		}
	}
	return foundUser
}

func DeleteUserById(id int) {
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
