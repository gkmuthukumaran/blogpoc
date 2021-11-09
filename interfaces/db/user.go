package db

import (
	"encoding/json"
	"fmt"

	"github.com/blogpoc/model"
	"github.com/boltdb/bolt"
)

func InsertUserDetails(user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {

		err := tx.Bucket([]byte("DB")).Bucket([]byte("BLOG")).Put([]byte(user.Username), []byte(data))
		if err != nil {
			return fmt.Errorf("could not insert User: %v", err)
		}
		return nil
	})
	fmt.Println("Added User")
	return err
}

func GetUserDetails(username string) (*model.User, error) {
	var user model.User
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("USER"))

		b.ForEach(func(k, v []byte) error {
			if username == string(k) {
				err := json.Unmarshal(v, &user)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if len(username) == 0 {
			return fmt.Errorf("Data not found!")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
