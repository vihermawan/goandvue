package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rmai99/goandvue/api/controllers"
)

var server = controllers.Server{}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8090")
	// uName, email, pwd := "", "", ""

	// mux := http.NewServeMux()

	// // Signup
	// mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
	// 	r.ParseForm()
	// 	uName = r.FormValue("name")
	// 	email = r.FormValue("email")
	// 	pwd = r.FormValue("password")
	// 	// for key, value := range r.Form {
	// 	// 	fmt.Printf("%s = %s\n", key, value)
	// 	// }
	// 	uNameCheck := helper.IsEmpty(uName)
	// 	emailCheck := helper.IsEmpty(email)
	// 	pwdCheck := helper.IsEmpty(pwd)

	// 	if uNameCheck || emailCheck || pwdCheck {
	// 		fmt.Fprintf(w, "error. there is empty data")
	// 	} else {
	// 		fmt.Fprintf(w, "Registration Succesful")
	// 	}

	// })

	// //Login
	// mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	// })

	// http.ListenAndServe(":8080", mux)

}

// package helpers
//
// func isEmpty(data string) bool {
// 	if len(data) == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
