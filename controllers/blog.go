package controllers

import (
	"net/http"

	"github.com/blogpoc/interfaces/db"
	"github.com/blogpoc/model"
	"github.com/blogpoc/utils"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func GetBlogs(c echo.Context) error {
	id := c.Param("id")
	data, err := db.GetBlogDetails(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.BldGnrRsp(http.StatusNotFound, err.Error(), nil))
	}

	respInterface := make([]interface{}, len(data))
	for i, evc := range data {
		mapForm := (structs.New(evc)).Map()
		respInterface[i] = mapForm
	}

	return c.JSON(http.StatusOK, utils.BldGnrRsp(200, "Success", &respInterface))
}

func PostBlog(c echo.Context) error {

	var blog model.Blog

	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, utils.BldGnrRsp(http.StatusBadRequest, err.Error(), nil))
	}
	id := uuid.New().String()
	blog.Id = id
	err := db.InsertBlogDetails(blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BldGnrRsp(500, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.BldGnrRsp(200, "Success", utils.ToInterfaceArrayFromString(id)))

}
