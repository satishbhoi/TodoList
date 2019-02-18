package handler

import (
	"net/http"

	"../model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) AddTodo(c echo.Context) (err error) {
	u := &model.User{
		ID: bson.ObjectIdHex(getID(c)),
	}
	p := &model.TodoDetail{
		ID:     bson.NewObjectId(),
		UserId: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	if p.TaskName == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid  fields"}
	}

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("TodoList").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("TodoList").C("todolist").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) ListTodo(c echo.Context) (err error) {
	userID := getID(c)
	// Defaults

	// Retrieve posts from database
	todoData := []*model.TodoDetail{}
	db := h.DB.Clone()
	if err = db.DB("TodoList").C("todolist").
		Find(bson.M{"UserId": userID}).
		All(&todoData); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, todoData)
}

func (h *Handler) UpdateTodo(c echo.Context) (err error) {
	userID := getID(c)
	id := c.Param("id")
	p := &model.TodoDetail{}
	if err = c.Bind(p); err != nil {
		return
	}
	p.UserId = userID
	// Add a follower to user

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("TodoList").C("todolist").
		UpdateId(bson.ObjectIdHex(id), p); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
			//return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handler) StatusTodo(c echo.Context) (err error) {
	id := c.Param("id")
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("TodoList").C("todolist").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"Status": 1}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
			//return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusOK, "Task completed")
}
