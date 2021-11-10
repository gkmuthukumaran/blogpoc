package db

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var dbname string

func SetupDB(name string) error {
	var err error
	db, err = bolt.Open("blog.db", 0600, nil)
	fmt.Println("testSetup", err)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte(name))
		dbname = name
		fmt.Println("testSetup2", err)
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists([]byte("USER"))
		fmt.Println("testSetup3", err)
		if err != nil {
			return err
		}
		_, err = root.CreateBucketIfNotExists([]byte("BLOG"))
		fmt.Println("testSetup4", err)
		if err != nil {
			return err
		}
		return nil
	})
	fmt.Println("testSetup5", err)
	if err != nil {
		return fmt.Errorf("could not set up buckets, %v", err)
	}

	initialiseUserDetails()

	fmt.Println("DB Setup Done")
	return nil
}

func Closedb() {
	if db != nil {
		db.Close()
	}

}

func inittestdb() {

	SetupDB("TESTDB")
}
