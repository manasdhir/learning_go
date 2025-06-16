package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type User struct {
	ID    string `xml:"id,attr"`
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type UserDetails struct {
	XMLName   xml.Name `xml:"users"`
	UserArray []User   `xml:"user"`
}

func main() {
	file, err := os.Open("users.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var userDetails UserDetails
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&userDetails)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	for _, user := range userDetails.UserArray {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
