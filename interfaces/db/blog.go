package db

import (
	"encoding/json"
	"fmt"

	"github.com/blogpoc/model"
	"github.com/boltdb/bolt"
)

func GetBlogDetails(id string) ([]model.Blog, error) {
	var blogs []model.Blog
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("BLOG"))
		matched := false
		b.ForEach(func(k, v []byte) error {
			if !matched{
				var blog model.Blog
				if id == string(k) {
					matched = true
					err := json.Unmarshal(v,&blog)
					if err != nil {
						return err
					}
					blogs = append(blogs,blog)
				} else if id == "" {
					err := json.Unmarshal(v,&blog)
					if err != nil {
						return err
					}
					blogs = append(blogs,blog)
				}
				return nil
			} else {
				return nil
			}
		})
		if len(blogs) == 0 {
			return fmt.Errorf("Data not found!")
		}
		return nil
	})
	if err != nil{
		return nil, err
	}
return blogs, nil
}		
func InsertBlogDetails(blog model.Blog) error {
	data, err := json.Marshal(blog)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		
		err := tx.Bucket([]byte("DB")).Bucket([]byte("BLOG")).Put([]byte(blog.Id), []byte(data))
		if err != nil {
			return fmt.Errorf("could not insert Blog: %v", err)
		}
		return nil
	})
	fmt.Println("Added Blog")
	return err
}

