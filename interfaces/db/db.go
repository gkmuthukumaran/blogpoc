package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func setupDB() error {
	var err error
	db, err = bolt.Open("blog.db", 0600, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists([]byte("USER"))
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists([]byte("BLOG"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return nil
}

func Closedb() {
	if db != nil {
		db.Close()
	}

}
