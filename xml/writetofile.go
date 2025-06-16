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
	userDetails := UserDetails{
		UserArray: []User{
			{ID: "1", Name: "XYZ1", Email: "abc1@i.com"},
			{ID: "2", Name: "XYZ2", Email: "abc2@i.com"},
		},
	}

	file, err := os.Create("users.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	output, err := xml.MarshalIndent(userDetails, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return
	}

	file.WriteString(xml.Header) // adds <?xml version="1.0" encoding="UTF-8"?>
	file.Write(output)           // writes formatted XML to file
	fmt.Println("Data written to users.xml successfully.")
}
