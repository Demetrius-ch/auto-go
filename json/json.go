package main

import (
	"encoding/json"
	"fmt"
)


type User struct {
	Name string `json:"name"`
	Prenom string `json:"prenom"`
	Age int `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	API := `[ {
	"name": "BANGAKE",
	"prenom": "Demetrius",
	"age": 23,
	"email": "lgdemetrius@gmail.com",
	"phone": "+23672051117"},
	{
	"name": "Chandler",
	"prenom": "Le Grand",
	"age": 56,
	"email": "triphelin@gmail.com",
	"phone": "+23675483942"}
	 ]`
	 var users []User
	 err := json.Unmarshal([]byte (API), &users)
	 if err != nil {
		fmt.Println("Error", err)
	 }
	 fmt.Printf("Json: %v\n", users)

	 //------------------------------------
	 
	 var mystrcut []User
	 var user_one User
	 user_one.Name = "Guedalia"
	 user_one.Prenom = "Corneille"
	 user_one.Age = 20
	 user_one.Email = "corneille@gmail.com"
	 user_one.Phone = "72427027"

	 mystrcut = append(mystrcut, user_one)
	 jsonFromSlice, err := json.Marshal(mystrcut)

	 if  err != nil {
		fmt.Println("Error", err)
		
		
	 }
	 fmt.Println(string(jsonFromSlice))

}