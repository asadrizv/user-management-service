package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
}

func main() {
	// Post a user and check for error
	fmt.Println("Posting a user")
	user := &User{
		Name:     "John Doe",
		Email:    "zoroo@example.com",
		Password: "password123",
		Country:  "US",
	}
	userJSON, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		fmt.Println("Error posting user")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var usrCreated User
	json.Unmarshal(body, &usrCreated)

	// Find the user by ID
	fmt.Println("Finding the user by ID")
	resp, err = http.Get(fmt.Sprintf("http://localhost:8080/users/%s", usrCreated.ID))
	if err != nil {
		fmt.Println("Error finding user by ID", err)
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var foundUser User
	json.Unmarshal(body, &foundUser)
	if foundUser.Name != user.Name || foundUser.Email != user.Email {
		fmt.Println("Error finding user by ID")
		return
	}

	// Update the user and see if the changes are reflected
	fmt.Println("Updating the user")
	user.ID = usrCreated.ID
	user.Name = "Jane Doe"
	user.Email = "jane.doe@example.com"
	userJSON, _ = json.Marshal(user)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/users/%s", usrCreated.ID), bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error updating user", err)
		return
	}
	defer resp.Body.Close()

	resp, err = http.Get(fmt.Sprintf("http://localhost:8080/users/%s", usrCreated.ID))
	if err != nil {
		fmt.Println("Error getting user", err)
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var updatedUser User
	json.Unmarshal(body, &updatedUser)
	if updatedUser.Name != "Jane Doe" || updatedUser.Email != "jane.doe@example.com" {
		fmt.Println("Error updating user", updatedUser)
		return
	}

	// Test getting a user by email with pagination
	fmt.Println("Testing getting a user by email with pagination and query")
	resp, err = http.Get("http://localhost:8080/users?query=email%3D%27jane.doe@example.com%27&page=10&limit=5")
	if err != nil {
		fmt.Println("Error getting user by email with pagination", err)
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var users []User
	json.Unmarshal(body, &users)
	if len(users) == 0 {
		fmt.Println("Error getting user by email with pagination")
		return
	}

	if users[0].Email != "jane.doe@example.com" {
		fmt.Println("Error getting user by email with pagination")
		return
	}
	fmt.Println("Integration tests passed successfully!")
}
