package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skbhati199/learningspacepro-api/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()

	// Users
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.GET("/user", uc.GetAllUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Teacher All
	tc := controllers.NewTeacherController(getSession())
	r.GET("/teacher/:id", tc.GetTeacher)
	r.GET("/teacher", tc.GetAllTeachers)
	r.POST("/teacher", tc.CreateTeacher)
	r.DELETE("/teacher/:id", tc.DeleteTeacher)

	// Batch All
	bc := controllers.NewBatchController(getSession())
	r.GET("/batch/:id", bc.GetBatch)
	r.GET("/batch", bc.GetAllBatchs)
	r.POST("/batch", bc.CreateBatch)
	r.DELETE("/batch/:id", bc.DeleteBatch)

	// Institude All
	ic := controllers.NewInstitudeController(getSession())
	r.GET("/institude/:id", ic.GetInstitude)
	r.GET("/institude", ic.GetAllInstitudes)
	r.POST("/institude", ic.CreateInstitude)
	r.DELETE("/institude/:id", ic.DeleteInstitude)

	// Chat All
	cc := controllers.NewChatController(getSession())
	r.GET("/chat/:id", cc.GetChat)
	r.GET("/chat", cc.GetAllChatMessages)
	r.POST("/chat", cc.CreateChat)
	r.DELETE("/chat/:id", cc.DeleteChat)

	// Listen and serve on
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		return
	}

	fmt.Println("")
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	fmt.Println("Mongo Database successfully connected ")
	return s
}
