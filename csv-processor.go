package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	FName string `json:"First Name [Required]"`
	LName string `json:"Last Name [Required]"`
	Email string `json:"Email Address [Required]"`
}

func main() {
	file, err := os.Open("./user-report-google-admin")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	var users Users

	json.Unmarshal(bytes, &users)

	csvfile, err := os.OpenFile("output.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}

	defer csvfile.Close()

	logfile, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}

	defer logfile.Close()

	log.SetOutput(logfile)

	newwriter := csv.NewWriter(csvfile)

	for i := 0; i < len(users.Users); i++ {
		var temparray []string
		temparray = append(temparray, users.Users[i].FName, users.Users[i].LName, users.Users[i].Email)
		log.Print(temparray)
		newwriter.Write(temparray)
		newwriter.Flush()
	}
	fmt.Println(len(users.Users))

}
