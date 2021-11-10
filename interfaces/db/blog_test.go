package db

import (
	"testing"

	"github.com/blogpoc/model"
	"github.com/google/uuid"
)

var id = uuid.New().String()

func TestInsertBlogDetails(t *testing.T) {

	if db == nil {
		inittestdb()
	}

	blog := model.Blog{
		Id:      id,
		Title:   "dummy title test",
		Author:  "test author",
		Content: "test content",
	}

	err := InsertBlogDetails(blog)
	if err != nil {
		t.Errorf("insertion faild %v", err.Error())
	}
}

func TestGetBlogDetails(t *testing.T) {

	blog, err := GetBlogDetails(id)
	if err != nil {
		t.Errorf("get faild %v", err.Error())
	}

	if len(blog) != 1 {
		t.Errorf("get faild")
		return
	}
	if blog[0].Title != "dummy title test" {
		t.Errorf("Title not match ")

	}
}

func TestGetAllBlog(t *testing.T) {

	blog, err := GetBlogDetails("")
	if err != nil {
		t.Errorf("get faild %v", err.Error())
	}
	if len(blog) == 0 {
		t.Errorf("get faild")
	}
}
