package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	//"github.com/tylerb/graceful"
	"gopkg.in/mgo.v2"
)

type Configuration struct {
	MongoDBHosts string
	AuthDatabase string
	AuthUserName string
	AuthPassword string
}

func main() {

	configFile, _ := os.Open("conf.json")
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	fmt.Println()

	e := echo.New()

	file, _ := os.Create("todo.log")

	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Output = file
	e.Use(middleware.LoggerWithConfig(loggerConfig))

	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{configuration.MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: configuration.AuthDatabase,
		Username: configuration.AuthUserName,
		Password: configuration.AuthPassword,
	}

	//db, err := mgo.Dial("localhost:8000")
	db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		e.Logger.Fatal(err)

	}

	if err = db.Copy().DB("TodoList").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)

	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/tasks", h.AddTodo)
	e.POST("/fetchTasks", h.ListTodo)
	e.POST("/updateTask/:id", h.UpdateTodo)
	e.POST("/completeTask/:id", h.StatusTodo)

	e.Server.Addr = ":6000"

	//graceful.ListenAndServe(e.Server, 5*time.Second)

	e.Logger.Fatal(e.Start(":6000"))

}
