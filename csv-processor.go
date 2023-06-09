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
	// open the user report generated by Google Admin
	file, err := os.Open("./user-report-google-admin")
	if err != nil {
		fmt.Println(err)
	}
	
	// once file is no longer in use, call back to close the file
	defer file.Close()
	
	
	bytes, _ := ioutil.ReadAll(file)

	var users Users
	
	json.Unmarshal(bytes, &users)

	// create output file for the seperated results
	csvfile, err := os.OpenFile("output.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}

	// once file is no longer in use, call back to close the file
	defer csvfile.Close()

	newwriter := csv.NewWriter(csvfile)
	
	// for every user up until the last user on the list
	for i := 0; i < len(users.Users); i++ {
		// initiate array for newwriter function to write to CSV
		var temparray []string
		temparray = append(temparray, users.Users[i].FName, users.Users[i].LName, users.Users[i].Email)
		// the writer will write the array objects over x columns on a single row
		// where x == number of objects in the array
		newwriter.Write(temparray)
		newwriter.Flush()
	}
}
