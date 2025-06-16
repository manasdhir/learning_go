package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	Id      int32    `xml:"id"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

func main() {
	data := []byte(`<user><id>1</id><name>XYZ</name><email>abc@i.com</email></user>`)
	var user User
	err := xml.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Id: %d, Name:%s, Email: %s\n", user.Id, user.Name, user.Email)
}
