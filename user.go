package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// defining our database connection
var db *gorm.DB
var err error

// define our user struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// initialize our database and creaet tables if they don't exist
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

// get all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	// connection to the database
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprint(w, "New user successfully created")
}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// find user that matches name
	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ? ", name).Find(&user)

	db.Delete(&user)
	fmt.Fprint(w, "user successfully deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ? ", name).Find(&user)
	user.Email = email

	db.Save(&user)
	fmt.Fprintln(w, "SuccessfulLy Updated User")

}
